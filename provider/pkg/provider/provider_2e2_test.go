// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

package provider_test

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"testing"

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
