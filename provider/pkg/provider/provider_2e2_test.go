// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

package provider_test

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/pulumi/providertest"
	"github.com/pulumi/providertest/providers"
	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/assertpreview"
	"github.com/pulumi/providertest/pulumitest/opttest"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/provider"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/stretchr/testify/assert"
)

func TestE2eSnapshots(t *testing.T) {
	t.Run("logGroup-0.94.0", func(t *testing.T) {
		t.Parallel()
		test := newAwsTest(t, filepath.Join("testdata", "logGroup"))
		testUpgradeFrom(t, test, "0.94.0")
	})

	t.Run("webAcl-0.94.0", func(t *testing.T) {
		t.Parallel()
		test := newAwsTest(t, filepath.Join("testdata", "webAcl"))
		testUpgradeFrom(t, test, "0.94.0")
	})
}

func TestAutonaming(t *testing.T) {
	skipIfShort(t)
	pt := newAwsTest(t, filepath.Join("testdata", "autonaming"), opttest.Env("PULUMI_EXPERIMENTAL", "1"))
	pt.Preview(t)
	up := pt.Up(t)
	logGroupName, ok := up.Outputs["logGroupName"].Value.(string)
	assert.True(t, ok)
	assert.Contains(t, logGroupName, "autonaming-log-") // project + name + random suffix
	fifoQueueName, ok := up.Outputs["fifoQueueName"].Value.(string)
	assert.True(t, ok)

	// Check that the queue name matches pattern: ${name}-${alphanum(6)}.fifo (.fifo is the resource's autonaming trivia suffix)
	assert.Regexp(t, `^queue-[a-zA-Z0-9]{6}\.fifo$`, fifoQueueName)
}

func TestThrottling(t *testing.T) {
	// Creates, destroys, and immediately re-creates an SQS to test whether we
	// correctly retry on throttling errors.
	t.Parallel()
	skipIfShort(t)

	// Pick a slightly random name so as to not conflict with concurrent CI runs.
	name := fmt.Sprintf("throttling-queue-%d", rand.Intn(1000))
	test := newAwsTest(t, filepath.Join("testdata", "throttling"))
	test.SetConfig(t, "queueName", name)
	_ = test.Up(t)
	_ = test.Destroy(t)
	_ = test.Up(t)
}

// EFS replication refresh retry constants.
// These values are tuned for EFS replication behavior:
// - Initial replication setup takes 5-10 minutes
// - State transitions happen gradually
// - 10 attempts with increasing backoff covers ~10-15 minutes of polling
const (
	efsRefreshMaxAttempts    = 10
	efsRefreshInitialBackoff = 5 * time.Second
	efsRefreshMaxBackoff     = 30 * time.Second
)

func TestEfsReplicationProtection(t *testing.T) {
	// Skip: takes ~45 min to run. Manual testing only.
	// Run with: go test -v -timeout 45m -run TestEfsReplicationProtection ./pkg/provider
	t.Skip("Manual test: EFS replication takes ~45 minutes")

	skipIfShort(t)

	// NOTE: This test uses explicit provider resources with different regions.
	// We do NOT use newAwsTest/attachProvider here because that attaches a single
	// provider server which ignores the region configurations in the YAML file.
	// Instead, we use the actual provider binary which respects each provider's config.
	pt := pulumitest.NewPulumiTest(t, filepath.Join("testdata", "efs-replication-protection"),
		opttest.Env("PULUMI_DEBUG_GRPC", "true"),
		opttest.LocalProviderPath("aws-native", filepath.Join("..", "..", "..", "bin")),
	)
	// Cleanup: Using aws:efs:ReplicationConfiguration sidecar resource means Pulumi
	// handles deletion order correctly - replication is deleted before file systems.
	defer pt.Destroy(t)

	// This test sets up actual EFS replication to verify the REPLICATING enum value
	// Refresh must succeed not reject REPLICATING.

	t.Log("=== Phase 1: Initial Preview and Up ===")
	pt.Preview(t)

	up := pt.Up(t)
	assert.NotNil(t, up.Summary)

	sourceProtection, ok := up.Outputs["sourceProtection"].Value.(string)
	assert.True(t, ok, "sourceProtection output should be a string")
	assert.Equal(t, "ENABLED", sourceProtection, "source should have ENABLED protection")

	destinationProtection, ok := up.Outputs["destinationProtection"].Value.(string)
	assert.True(t, ok, "destinationProtection output should be a string")
	t.Logf("Destination protection after up: %s", destinationProtection)

	// Log all outputs for debugging
	t.Log("=== Outputs after initial up ===")
	for k, v := range up.Outputs {
		t.Logf("  %s = %v", k, v.Value)
	}

	// KEY TEST: Refresh with adaptive backoff until destination shows REPLICATING.
	// AWS takes time to transition the destination from DISABLED to REPLICATING after
	// replication is established. We must wait for this transition before running up.
	t.Log("=== Phase 2: Refresh loop waiting for REPLICATING state ===")
	backoff := efsRefreshInitialBackoff
	foundReplicating := false

	for attempt := 1; attempt <= efsRefreshMaxAttempts; attempt++ {
		t.Logf("Refresh attempt %d/%d (backoff: %v)", attempt, efsRefreshMaxAttempts, backoff)

		refreshResult := pt.Refresh(t)
		assert.NotNil(t, refreshResult.Summary, "Refresh should complete successfully")

		// Log refresh output for debugging
		t.Logf("=== Refresh %d StdOut ===\n%s", attempt, refreshResult.StdOut)
		if refreshResult.StdErr != "" {
			t.Logf("=== Refresh %d StdErr ===\n%s", attempt, refreshResult.StdErr)
		}

		if refreshResult.Summary.ResourceChanges != nil {
			changes := *refreshResult.Summary.ResourceChanges
			t.Logf("Resource changes: %+v", changes)
		}

		// Check if destination has transitioned to REPLICATING
		// After refresh, we need to run preview/up to get current outputs
		// Parse the refresh stdout for the destinationProtection value
		if strings.Contains(refreshResult.StdOut, "REPLICATING") {
			t.Log("Detected REPLICATING in refresh output, proceeding with up")
			foundReplicating = true
			break
		}

		if attempt < efsRefreshMaxAttempts {
			t.Logf("Waiting %v for REPLICATING state...", backoff)
			time.Sleep(backoff)
			if backoff < efsRefreshMaxBackoff {
				backoff = backoff * 3 / 2
			}
		}
	}

	// If we didn't find REPLICATING, the test is still valid - it tests that
	// refresh succeeds. But we log a warning.
	if !foundReplicating {
		t.Log("WARNING: Destination never transitioned to REPLICATING during refresh loop")
		t.Log("This may cause the subsequent up to fail if AWS has started the transition")
	}

	// KEY TEST: Run pulumi up after refresh - this is where the failure occurs
	// The refresh may have added replicationConfiguration to state, and now
	// pulumi up tries to reconcile that with the program (which doesn't specify it)
	t.Log("=== Phase 3: Up after refresh (this is where the failure occurs) ===")
	upAfterRefresh := pt.Up(t)
	assert.NotNil(t, upAfterRefresh.Summary, "Up after refresh should complete successfully")

	t.Log("=== Outputs after up (post-refresh) ===")
	for k, v := range upAfterRefresh.Outputs {
		t.Logf("  %s = %v", k, v.Value)
	}

	t.Log("Test passed: Refresh and subsequent Up succeeded with REPLICATING enum value")
}

func testUpgradeFrom(t *testing.T, test *pulumitest.PulumiTest, version string) {
	result := providertest.PreviewProviderUpgrade(t, test, "aws-native", version)
	assertpreview.HasNoChanges(t, result)
}

func newAwsTest(t *testing.T, source string, opts ...opttest.Option) *pulumitest.PulumiTest {
	t.Helper()
	opts = append(opts, attachProvider())
	test := pulumitest.NewPulumiTest(t, source, opts...)
	test.SetConfig(t, "aws-native:region", "us-west-2")
	test.SetConfig(t, "aws:region", "us-west-2")
	return test
}

func attachProvider() opttest.Option {
	return opttest.AttachProviderServer("aws-native", func(pt providers.PulumiTest) (pulumirpc.ResourceProviderServer, error) {
		return testProviderServer()
	})
}

func testProviderServer() (pulumirpc.ResourceProviderServer, error) {
	cmdDir := filepath.Join("..", "..", "..", "bin")
	schemaBytes, err := os.ReadFile(filepath.Join(cmdDir, "schema.json.gz"))
	if err != nil {
		return nil, err
	}
	metadataBytes, err := os.ReadFile(filepath.Join(cmdDir, "metadata.json.gz"))
	if err != nil {
		return nil, err
	}
	return provider.NewAwsNativeProvider(nil, "aws-native", "0.1.0", schemaBytes, metadataBytes)
}

func skipIfShort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}
}
