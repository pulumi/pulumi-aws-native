// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
//go:build nodejs || all
// +build nodejs all

package examples

import (
	"bytes"
	"context"
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/assertpreview"
	"github.com/pulumi/providertest/pulumitest/changesummary"
	"github.com/pulumi/providertest/pulumitest/opttest"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
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

func TestRefreshChanges(t *testing.T) {
	ctx := context.Background()
	config, err := config.LoadDefaultConfig(ctx)
	require.NoError(t, err)
	client := ssm.NewFromConfig(config)

	cwd := getCwd(t)
	options := []opttest.Option{
		opttest.LocalProviderPath("aws-native", filepath.Join(cwd, "..", "bin")),
		opttest.YarnLink("@pulumi/aws-native"),
	}
	test := pulumitest.NewPulumiTest(t, filepath.Join(cwd, "refresh-changes"), options...)
	defer test.Destroy(t)

	upResult := test.Up(t)
	t.Logf("#%v", upResult.Summary)
	paramName := upResult.Outputs["paramName"].Value.(string)

	// Update the value of the parameter
	_, err = client.PutParameter(ctx, &ssm.PutParameterInput{
		Name:      &paramName,
		Value:     aws.String("new-value"),
		Overwrite: aws.Bool(true),
	})
	require.NoError(t, err)

	// Refresh and assert that there is an update
	refreshResult := test.Refresh(t)
	t.Logf("RefreshResult #%s", refreshResult.StdOut)
	changes := *refreshResult.Summary.ResourceChanges
	assert.Equal(t, 1, changes["update"])

	deployment, err := test.ExportStack(t).Deployment.MarshalJSON()
	require.NoError(t, err)
	var stack apitype.DeploymentV3
	err = json.Unmarshal(deployment, &stack)
	require.NoError(t, err)

	// assert that the refresh has updated the value in state
	var value string
	for _, res := range stack.Resources {
		if res.Type == "aws-native:ssm:Parameter" {
			value = res.Inputs["value"].(string)
		}
	}
	assert.Equal(t, "new-value", value)

	previewResult := test.Preview(t)
	t.Logf("PreviewResult #%s", previewResult.StdOut)
	summary := changesummary.ChangeSummary(previewResult.ChangeSummary)
	updateSummary := summary.WhereOpEquals(apitype.OpUpdate)
	assert.Equal(t, 1, len(*updateSummary))
}

func TestUnorderedCollectionsRefresh(t *testing.T) {
	_ = getEnvRegion(t)
	cwd := getCwd(t)
	test := pulumitest.NewPulumiTest(t, filepath.Join(cwd, "unordered-collections"),
		opttest.LocalProviderPath("aws-native", filepath.Join(cwd, "..", "bin")),
		opttest.YarnLink("@pulumi/aws-native"),
	)
	defer test.Destroy(t)

	test.Up(t)
	test.Refresh(t)

	deploymentJSON, err := test.ExportStack(t).Deployment.MarshalJSON()
	require.NoError(t, err)
	var deployment apitype.DeploymentV3
	require.NoError(t, json.Unmarshal(deploymentJSON, &deployment))

	var role, taskDefinition *apitype.ResourceV3
	for i := range deployment.Resources {
		switch deployment.Resources[i].Type {
		case "aws-native:iam:Role":
			role = &deployment.Resources[i]
		case "aws-native:ecs:TaskDefinition":
			taskDefinition = &deployment.Resources[i]
		}
	}
	require.NotNil(t, role, "expected IAM Role in exported deployment")
	require.NotNil(t, taskDefinition, "expected ECS TaskDefinition in exported deployment")

	// IAM returns inline policy names alphabetically, proving that refresh saw
	// AWS order while the program and checkpoint retain the authored order.
	assert.Equal(t, []string{"z-policy", "a-policy"}, rolePolicyNames(t, role.Inputs["policies"]))
	assert.Equal(t, []string{"a-policy", "z-policy"}, rolePolicyNames(t, role.Outputs["policies"]))

	// ECS preserves the authored environment order. Reverse only the program
	// order after refresh to exercise nested unordered Diff against live state.
	assert.Equal(t, []string{"Z_VAR", "A_VAR"}, taskEnvironmentNames(t, taskDefinition.Inputs))
	assert.Equal(t, []string{"Z_VAR", "A_VAR"}, taskEnvironmentNames(t, taskDefinition.Outputs))
	test.SetConfig(t, "reverseEnvironment", "true")

	previewResult := test.Preview(t)
	assertpreview.HasNoChanges(t, previewResult)
}

func rolePolicyNames(t *testing.T, value any) []string {
	t.Helper()
	policies, ok := value.([]interface{})
	require.True(t, ok, "expected policies array, got %T", value)

	names := make([]string, len(policies))
	for i, rawPolicy := range policies {
		policy, ok := rawPolicy.(map[string]interface{})
		require.True(t, ok, "expected policy object, got %T", rawPolicy)
		name, ok := policy["policyName"].(string)
		require.True(t, ok, "expected policyName string, got %T", policy["policyName"])
		names[i] = name
	}
	return names
}

func taskEnvironmentNames(t *testing.T, properties map[string]interface{}) []string {
	t.Helper()
	containers, ok := properties["containerDefinitions"].([]interface{})
	require.True(t, ok, "expected containerDefinitions array, got %T", properties["containerDefinitions"])
	require.Len(t, containers, 1)
	container, ok := containers[0].(map[string]interface{})
	require.True(t, ok, "expected container definition object, got %T", containers[0])
	environment, ok := container["environment"].([]interface{})
	require.True(t, ok, "expected environment array, got %T", container["environment"])

	names := make([]string, len(environment))
	for i, rawVariable := range environment {
		variable, ok := rawVariable.(map[string]interface{})
		require.True(t, ok, "expected environment object, got %T", rawVariable)
		name, ok := variable["name"].(string)
		require.True(t, ok, "expected environment name string, got %T", variable["name"])
		names[i] = name
	}
	return names
}

func TestPreviewStableProperties(t *testing.T) {
	cwd := getCwd(t)
	options := []opttest.Option{
		opttest.LocalProviderPath("aws-native", filepath.Join(cwd, "..", "bin")),
		opttest.YarnLink("@pulumi/aws-native"),
	}
	test := pulumitest.NewPulumiTest(t, filepath.Join(cwd, "stable-outputs-preview"), options...)
	test.SetConfig(t, "lambdaDescription", "Lambda 1")
	defer func() {
		test.Destroy(t)
	}()

	upResult := test.Up(t)
	t.Logf("#%v", upResult.Summary)

	// updating a non-replaceOnChanges property to ensure
	// that the downstream resources doesn't show a replacement
	test.SetConfig(t, "lambdaDescription", "Lambda 2")
	previewResult := test.Preview(t)
	assertpreview.HasNoReplacements(t, previewResult)
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
			Dir:    filepath.Join(getCwd(t), "autonaming-overlay"),
			Stderr: &buf,
			Config: map[string]string{
				"roleName": "myReallyLongRoleNameThatIsLongerThan64CharactersOneTwoThreeFour",
			},
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
