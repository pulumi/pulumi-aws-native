// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/assertpreview"
	"github.com/pulumi/providertest/pulumitest/opttest"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getEnvRegion(t *testing.T) string {
	envRegion := os.Getenv("AWS_REGION")
	if envRegion == "" {
		t.Skipf("Skipping test due to missing AWS_REGION environment variable")
	}

	return envRegion
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions(t *testing.T) integration.ProgramTestOptions {
	envRegion := getEnvRegion(t)
	binPath, err := filepath.Abs("../bin")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Using binPath %s\n", binPath)
	return integration.ProgramTestOptions{
		Config: map[string]string{
			"aws-native:region": envRegion,
		},
		Quick: true,
		LocalProviders: []integration.LocalDependency{
			{
				Package: "aws-native",
				Path:    binPath,
			},
		},
	}
}

// TestUnorderedCollectionsRefresh runs a YAML program so it does not depend on
// any per-language runtime toolchain. It verifies that refresh keeps the
// authored ordering of CloudFormation unordered collections and that a
// program-side reorder produces a no-change preview.
func TestUnorderedCollectionsRefresh(t *testing.T) {
	_ = getEnvRegion(t)
	cwd := getCwd(t)
	test := pulumitest.NewPulumiTest(t, filepath.Join(cwd, "unordered-collections"),
		opttest.LocalProviderPath("aws-native", filepath.Join(cwd, "..", "bin")),
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
	test.SetConfig(t, "environmentOrderIndex", "1")

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
