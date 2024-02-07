// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

package provider_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/providertest"
	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/assertpreview"
	"github.com/pulumi/providertest/pulumitest/opttest"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/provider"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

func TestE2eSnapshots(t *testing.T) {
	t.Run("logGroup-0.94.0", func(t *testing.T) {
		t.Parallel()
		test := newAwsTest(t, filepath.Join("testdata", "logGroup"))
		testUpgradeFrom(test, "0.94.0")
	})

	t.Run("webAcl-0.94.0", func(t *testing.T) {
		t.Parallel()
		test := newAwsTest(t, filepath.Join("testdata", "webAcl"))
		testUpgradeFrom(test, "0.94.0")
	})
}

func testUpgradeFrom(test *pulumitest.PulumiTest, version string) {
	test.T().Helper()
	result := providertest.PreviewProviderUpgrade(test, "aws-native", version)
	assertpreview.HasNoChanges(test.T(), result)
}

func newAwsTest(t *testing.T, source string, opts ...opttest.Option) *pulumitest.PulumiTest {
	t.Helper()
	opts = append(opts, attachProvider())
	test := pulumitest.NewPulumiTest(t, source, opts...)
	test.SetConfig("aws-native:region", "us-west-2")
	test.SetConfig("aws:region", "us-west-2")
	return test
}

func attachProvider() opttest.Option {
	return opttest.AttachProviderServer("aws-native", func() (pulumirpc.ResourceProviderServer, error) {
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
	})
}
