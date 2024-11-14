// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
//go:build nodejs || all
// +build nodejs all

package examples

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/assertpreview"
	"github.com/pulumi/providertest/pulumitest/opttest"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimpleTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "simple-ts"),
		})

	integration.ProgramTest(t, &test)
}

func TestGetTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "get-ts"),
		})

	integration.ProgramTest(t, &test)
}

func TestCustomResourceEmulator(t *testing.T) {
	crossTest := func(t *testing.T, outputs auto.OutputMap) {
		require.Contains(t, outputs, "cloudformationAmiId")
		cloudformationAmiId := outputs["cloudformationAmiId"].Value.(string)
		require.NotEmpty(t, cloudformationAmiId)

		require.Contains(t, outputs, "emulatorAmiId")
		emulatorAmiId := outputs["emulatorAmiId"].Value.(string)
		assert.Equal(t, cloudformationAmiId, emulatorAmiId)
	}

	cwd := getCwd(t)
	options := []opttest.Option{
		opttest.LocalProviderPath("aws-native", filepath.Join(cwd, "..", "bin")),
		opttest.YarnLink("@pulumi/aws-native"),
	}
	test := pulumitest.NewPulumiTest(t, filepath.Join(cwd, "cfn-custom-resource"), options...)
	test.SetConfig(t, "amiRegion", "us-west-2")

	previewResult := test.Preview(t)
	t.Logf("#%v", previewResult.ChangeSummary)

	upResult := test.Up(t)
	t.Logf("#%v", upResult.Summary)
	crossTest(t, upResult.Outputs)

	previewResult = test.Preview(t)
	assertpreview.HasNoChanges(t, previewResult)

	test.SetConfig(t, "amiRegion", "us-east-1")
	upResult = test.Up(t)
	t.Logf("#%v", upResult.Summary)
	crossTest(t, upResult.Outputs)

	previewResult = test.Preview(t)
	assertpreview.HasNoChanges(t, previewResult)

	test.Destroy(t)
}

func TestVpcCidrs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "cidr-ts"),
		})

	integration.ProgramTest(t, &test)
}

func TestUpdate(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "update", "step1"),
			EditDirs: []integration.EditDir{
				{
					Dir:      filepath.Join(getCwd(t), "update", "step2"),
					Additive: true,
					ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
						// Check that the name of the updated secret is correct.
						secretValue := stackInfo.Outputs["secretValue"].(string)
						if secretValue != "secretbuzz" {
							t.Errorf("Expected secretValue to be 'secretbuzz', got %q (%[1]T)", secretValue)
						}
					},
				},
			},
		})

	integration.ProgramTest(t, &test)
}

func TestLambdaUpdate(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "lambda-update", "step1"),
			EditDirs: []integration.EditDir{
				{
					Dir:      filepath.Join(getCwd(t), "lambda-update", "step2"),
					Additive: true,
				},
			},
		})

	integration.ProgramTest(t, &test)
}

func TestNamingConventions(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "aws-native-naming-conventions"),
			ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
				// Check that the name of the created queue is correct.
				fifoName := stackInfo.Outputs["fifoName"].(string)
				if !(fifoName[len(fifoName)-5:] == ".fifo") {
					t.Errorf("Expected fifo queue name to end with '.fifo', got '%s'", fifoName)
				}
				standardName := stackInfo.Outputs["standardName"].(string)
				if standardName[len(standardName)-5:] == ".fifo" {
					t.Errorf("Expected standard queue name to not end with '.fifo', got '%s'", standardName)
				}
			},
		})

	integration.ProgramTest(t, &test)
}

func TestAutoNamingOverlay(t *testing.T) {
	var buf bytes.Buffer
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:           filepath.Join(getCwd(t), "autonaming-overlay"),
			Stderr:        &buf,
			ExpectFailure: true,
			ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
				assert.Contains(t, buf.String(), "is too large to fix max length constraint of 64")
			},
		})

	integration.ProgramTest(t, &test)
}

func TestAutoNamingOverlayWithConfig(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "autonaming-overlay"),
			Config: map[string]string{
				"aws-native:autoNaming": `{"autoTrim": true, "randomSuffixMinLength": 7}`,
				"roleName":              "myReallyLongRoleNameThatIsLongerThan64CharactersOneTwoThreeFour",
			},
			ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
				roleName := stackInfo.Outputs["roleName"].(string)
				assert.Equal(t, 64, len(roleName))
				assert.Regexp(t, "myReallyLongRoleNameThatIsLon64CharactersOneTwoThreeFour-[a-z0-9]{7}", roleName)
			},
		})

	integration.ProgramTest(t, &test)
}

func TestParallelTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "parallel-ts"),
		})

	integration.ProgramTest(t, &test)
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@pulumi/aws-native",
		},
	})

	return baseJS
}
