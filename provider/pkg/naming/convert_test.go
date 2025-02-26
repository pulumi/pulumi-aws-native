// Copyright 2016-2021, Pulumi Corporation.

package naming

import (
	"sort"
	"testing"

	"github.com/mattbaird/jsonpatch"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCfnToSdk(t *testing.T) {
	actual := CfnToSdk(cfnPayload)
	assert.Equal(t, sdkState, actual)
}

func TestSdkToCfn(t *testing.T) {
	res := sampleSchema.Resources["aws-native:ecs:Service"]
	actual, _ := SdkToCfn(&res, sampleSchema.Types, sdkState)
	assert.Equal(t, cfnPayload, actual)

	// error when value does not match type
	badSdkState := sdkState
	badSdkState["loadBalancers"] = map[string]interface{}{
		"containerName":  "my-app",
		"containerPort":  80,
		"targetGroupArn": "arn:aws:elasticloadbalancing:us-west-2:12345:targetgroup/app-tg-a56e4c3/51317265a5dc6a1f",
	}
	_, err := SdkToCfn(&res, sampleSchema.Types, badSdkState)
	cErr := &ConversionError{}
	assert.ErrorAs(t, err, &cErr, "Should catch conversions of invalid types")
}

func TestSdkToCfnEnumTypeRef(t *testing.T) {
	res := metadata.CloudAPIResource{
		Inputs: map[string]pschema.PropertySpec{
			"foo": {
				TypeSpec: pschema.TypeSpec{
					Ref: "#/types/bar",
				},
			},
		},
	}
	// Enums just look like strings in the metadata
	types := map[string]metadata.CloudAPIType{
		"bar": {Type: "string"},
	}
	state := map[string]interface{}{"foo": "BBB"}
	expected := map[string]interface{}{"Foo": "BBB"}
	actual, err := SdkToCfn(&res, types, state)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestSdkToCfnOneOf(t *testing.T) {
	res := metadata.CloudAPIResource{
		Inputs: map[string]pschema.PropertySpec{
			"foo": {
				TypeSpec: pschema.TypeSpec{
					OneOf: []pschema.TypeSpec{
						{
							Type:  "array",
							Items: &pschema.TypeSpec{Type: "string"},
						},
						{
							Type:                 "object",
							AdditionalProperties: &pschema.TypeSpec{Type: "number"},
						},
					},
				},
			},
		},
	}
	state := map[string]interface{}{"foo": map[string]interface{}{"bar": 1, "baz": 2}}
	expected := map[string]interface{}{"Foo": map[string]interface{}{"bar": 1, "baz": 2}}
	actual, err := SdkToCfn(&res, sampleSchema.Types, state)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

// Check that ambiguous OneOf without a discriminator will pick the largest match.
func TestSdkToCfnOneOfAmbiguous(t *testing.T) {
	res := metadata.CloudAPIResource{
		Inputs: map[string]pschema.PropertySpec{
			"configuration": {TypeSpec: pschema.TypeSpec{
				OneOf: []pschema.TypeSpec{
					{Ref: "#/types/aws-native:datazone:DataSourceConfigurationInput0Properties"},
					{Ref: "#/types/aws-native:datazone:DataSourceConfigurationInput1Properties"},
				},
			}},
		},
	}
	types := map[string]metadata.CloudAPIType{
		"aws-native:datazone:DataSourceConfigurationInput0Properties": metadata.CloudAPIType{
			Type: "object",
			Properties: map[string]pschema.PropertySpec{
				"glueRunConfiguration": pschema.PropertySpec{
					TypeSpec: pschema.TypeSpec{
						Ref: "#/types/aws-native:datazone:DataSourceGlueRunConfigurationInput",
					},
				},
			},
		},
		"aws-native:datazone:DataSourceGlueRunConfigurationInput": metadata.CloudAPIType{
			Type: "object",
			Properties: map[string]pschema.PropertySpec{
				"dataAccessRole": pschema.PropertySpec{TypeSpec: pschema.TypeSpec{Type: "string"}},
			},
		},
		"aws-native:datazone:DataSourceConfigurationInput1Properties": metadata.CloudAPIType{
			Type: "object",
			Properties: map[string]pschema.PropertySpec{"redshiftRunConfiguration": pschema.PropertySpec{
				TypeSpec: pschema.TypeSpec{
					Ref: "#/types/aws-native:datazone:DataSourceRedshiftRunConfigurationInput",
				},
			}},
		},
		"aws-native:datazone:DataSourceRedshiftRunConfigurationInput": metadata.CloudAPIType{
			Type: "object",
			Properties: map[string]pschema.PropertySpec{
				"dataAccessRole": pschema.PropertySpec{TypeSpec: pschema.TypeSpec{Type: "string"}},
			},
		},
	}

	result, err := SdkToCfn(&res, types, map[string]any{
		"configuration": map[string]any{
			"redshiftRunConfiguration": map[string]any{
				"dataAccessRole": "myrole",
			},
		},
	})
	require.NoError(t, err)

	actualConfig := result["Configuration"].(map[string]any)
	actualRedshiftConfig := actualConfig["RedshiftRunConfiguration"].(map[string]any)
	actualDataAccessRole := actualRedshiftConfig["DataAccessRole"].(string)
	require.Equal(t, "myrole", actualDataAccessRole)
}

func TestDiffToPatch(t *testing.T) {
	test := func(t *testing.T, diff resource.ObjectDiff, expected []jsonpatch.JsonPatchOperation) {
		res := sampleSchema.Resources["aws-native:ecs:Service"]
		actual, err := DiffToPatch(&res, sampleSchema.Types, &diff)
		assert.NoError(t, err)
		sort.SliceStable(actual, func(i, j int) bool {
			return actual[i].Path < actual[j].Path
		})
		assert.Equal(t, expected, actual)

	}

	t.Run("add-delete-replace", func(t *testing.T) {
		diff := resource.ObjectDiff{
			Updates: map[resource.PropertyKey]resource.ValueDiff{
				"desiredCount":         {New: resource.NewNumberProperty(2)},
				"enableEcsManagedTags": {New: resource.NewBoolProperty(true)},
				"loadBalancers": {
					New: resource.NewArrayProperty([]resource.PropertyValue{
						resource.NewObjectProperty(resource.PropertyMap{
							resource.PropertyKey("containerName"): resource.NewStringProperty("my-app"),
							resource.PropertyKey("containerPort"): resource.NewNumberProperty(80),
						}),
					}),
				},
			},
			Adds: map[resource.PropertyKey]resource.PropertyValue{
				"launchType": resource.NewStringProperty("FARGATE"),
			},
			Deletes: map[resource.PropertyKey]resource.PropertyValue{
				"platformVersion": resource.NewStringProperty("LATEST"),
			},
		}
		two := int32(2)
		expected := []jsonpatch.JsonPatchOperation{
			jsonpatch.NewPatch("replace", "/DesiredCount", two),
			jsonpatch.NewPatch("replace", "/EnableECSManagedTags", true),
			jsonpatch.NewPatch("add", "/LaunchType", "FARGATE"),
			jsonpatch.NewPatch("replace", "/LoadBalancers",
				[]interface{}{
					map[string]interface{}{
						"ContainerName": "my-app",
						"ContainerPort": 80.,
					},
				}),
			jsonpatch.NewPatch("remove", "/PlatformVersion", nil),
		}
		test(t, diff, expected)
	})

	t.Run("secrets", func(t *testing.T) {
		diff := resource.ObjectDiff{
			Updates: map[resource.PropertyKey]resource.ValueDiff{
				"cluster": {New: resource.MakeSecret(resource.NewStringProperty("mycluster"))},
				"loadBalancers": {
					New: resource.NewArrayProperty([]resource.PropertyValue{
						resource.NewObjectProperty(resource.PropertyMap{
							resource.PropertyKey("targetGroupArn"): resource.MakeSecret(resource.NewStringProperty("arn:mytargetgroup")),
						}),
					}),
				},
			},
			Adds: map[resource.PropertyKey]resource.PropertyValue{
				"schedulingStrategy": resource.MakeSecret(resource.NewStringProperty("REPLICA")),
			},
		}
		expected := []jsonpatch.JsonPatchOperation{
			jsonpatch.NewPatch("replace", "/Cluster", "mycluster"),
			jsonpatch.NewPatch("replace", "/LoadBalancers",
				[]interface{}{
					map[string]interface{}{
						"TargetGroupArn": "arn:mytargetgroup",
					},
				}),
			jsonpatch.NewPatch("add", "/SchedulingStrategy", "REPLICA"),
		}
		test(t, diff, expected)
	})
}

var cfnPayload = map[string]interface{}{
	"Cluster": "arn:aws:ecs:us-west-2:12345:cluster/cloud-api-cluster",
	"DeploymentConfiguration": map[string]interface{}{
		"DeploymentCircuitBreaker": map[string]interface{}{
			"Enable":   false,
			"Rollback": false,
		},
		"MaximumPercent":        200,
		"MinimumHealthyPercent": 100,
	},
	"DesiredCount":         3,
	"EnableECSManagedTags": false,
	"LaunchType":           "FARGATE",
	"LoadBalancers": []interface{}{map[string]interface{}{
		"ContainerName":  "my-app",
		"ContainerPort":  80,
		"TargetGroupArn": "arn:aws:elasticloadbalancing:us-west-2:12345:targetgroup/app-tg-a56e4c3/51317265a5dc6a1f",
	}},
	"Name": "app-svc-cloud-api",
	"NetworkConfiguration": map[string]interface{}{
		"AwsvpcConfiguration": map[string]interface{}{
			"AssignPublicIp": "ENABLED",
			"SecurityGroups": []interface{}{
				"sg-0e8188f55b67b53f3",
			},
			"Subnets": []interface{}{
				"subnet-0016572b",
				"subnet-43f43a1e",
				"subnet-c7d926bf",
				"subnet-d7e7fe9c",
			},
		},
	},
	"PlatformVersion":    "LATEST",
	"Role":               "arn:aws:iam::12345:role/aws-service-role/ecs.amazonaws.com/AWSServiceRoleForECS",
	"SchedulingStrategy": "REPLICA",
	"ServiceArn":         "arn:aws:ecs:us-west-2:12345:service/cloud-api-cluster/app-svc-cloud-api",
	"ServiceName":        "app-svc-cloud-api",
	"TaskDefinition":     "arn:aws:ecs:us-west-2:12345:task-definition/fargate-task-definition:131",
}

var sdkState = map[string]interface{}{
	"cluster": "arn:aws:ecs:us-west-2:12345:cluster/cloud-api-cluster",
	"deploymentConfiguration": map[string]interface{}{
		"deploymentCircuitBreaker": map[string]interface{}{
			"enable":   false,
			"rollback": false,
		},
		"maximumPercent":        200,
		"minimumHealthyPercent": 100,
	},
	"desiredCount":         3,
	"enableEcsManagedTags": false,
	"launchType":           "FARGATE",
	"loadBalancers": []interface{}{map[string]interface{}{
		"containerName":  "my-app",
		"containerPort":  80,
		"targetGroupArn": "arn:aws:elasticloadbalancing:us-west-2:12345:targetgroup/app-tg-a56e4c3/51317265a5dc6a1f",
	}},
	"name": "app-svc-cloud-api",
	"networkConfiguration": map[string]interface{}{
		"awsvpcConfiguration": map[string]interface{}{
			"assignPublicIp": "ENABLED",
			"securityGroups": []interface{}{
				"sg-0e8188f55b67b53f3",
			},
			"subnets": []interface{}{
				"subnet-0016572b",
				"subnet-43f43a1e",
				"subnet-c7d926bf",
				"subnet-d7e7fe9c",
			},
		},
	},
	"platformVersion":    "LATEST",
	"role":               "arn:aws:iam::12345:role/aws-service-role/ecs.amazonaws.com/AWSServiceRoleForECS",
	"schedulingStrategy": "REPLICA",
	"serviceArn":         "arn:aws:ecs:us-west-2:12345:service/cloud-api-cluster/app-svc-cloud-api",
	"serviceName":        "app-svc-cloud-api",
	"taskDefinition":     "arn:aws:ecs:us-west-2:12345:task-definition/fargate-task-definition:131",
}

var sampleSchema = &metadata.CloudAPIMetadata{
	Types: map[string]metadata.CloudAPIType{
		"aws-native:ecs:AwsVpcConfiguration": {
			Type: "object",
			Properties: map[string]pschema.PropertySpec{
				"assignPublicIp": {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"securityGroups": {
					TypeSpec: pschema.TypeSpec{
						Type:  "array",
						Items: &pschema.TypeSpec{Type: "string"},
					},
				},
				"subnets": {
					TypeSpec: pschema.TypeSpec{
						Type:  "array",
						Items: &pschema.TypeSpec{Type: "string"},
					},
				},
			},
		},
		"aws-native:ecs:DeploymentCircuitBreaker": {
			Type: "object",
			Properties: map[string]pschema.PropertySpec{
				"enable":   {TypeSpec: pschema.TypeSpec{Type: "boolean"}},
				"rollback": {TypeSpec: pschema.TypeSpec{Type: "boolean"}},
			},
		},
		"aws-native:ecs:DeploymentConfiguration": {
			Type: "object",
			Properties: map[string]pschema.PropertySpec{
				"deploymentCircuitBreaker": {
					TypeSpec: pschema.TypeSpec{
						Ref: "#/types/aws-native:ecs:DeploymentCircuitBreaker",
					},
				},
				"maximumPercent":        {TypeSpec: pschema.TypeSpec{Type: "integer"}},
				"minimumHealthyPercent": {TypeSpec: pschema.TypeSpec{Type: "integer"}},
			},
		},
		"aws-native:ecs:LoadBalancer": {
			Type: "object",
			Properties: map[string]pschema.PropertySpec{
				"containerName":    {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"containerPort":    {TypeSpec: pschema.TypeSpec{Type: "integer"}},
				"loadBalancerName": {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"targetGroupArn":   {TypeSpec: pschema.TypeSpec{Type: "string"}},
			},
		},
		"aws-native:ecs:NetworkConfiguration": {
			Type: "object",
			Properties: map[string]pschema.PropertySpec{
				"awsvpcConfiguration": {
					TypeSpec: pschema.TypeSpec{
						Ref: "#/types/aws-native:ecs:AwsVpcConfiguration",
					},
				},
			},
		},
	},
	Resources: map[string]metadata.CloudAPIResource{
		"aws-native:ecs:Service": {
			Inputs: map[string]pschema.PropertySpec{
				"serviceArn": {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"cluster":    {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"deploymentConfiguration": {
					TypeSpec: pschema.TypeSpec{
						Ref: "#/types/aws-native:ecs:DeploymentConfiguration",
					},
				},
				"desiredCount":         {TypeSpec: pschema.TypeSpec{Type: "integer"}},
				"enableEcsManagedTags": {TypeSpec: pschema.TypeSpec{Type: "boolean"}},
				"launchType":           {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"loadBalancers": {
					TypeSpec: pschema.TypeSpec{
						Type: "array",
						Items: &pschema.TypeSpec{
							Ref: "#/types/aws-native:ecs:LoadBalancer",
						},
					},
				},
				"networkConfiguration": {
					TypeSpec: pschema.TypeSpec{
						Ref: "#/types/aws-native:ecs:NetworkConfiguration",
					},
				},
				"platformVersion":    {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"role":               {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"schedulingStrategy": {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"serviceName":        {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"taskDefinition":     {TypeSpec: pschema.TypeSpec{Type: "string"}},
			},
			Outputs: map[string]pschema.PropertySpec{
				"name": {TypeSpec: pschema.TypeSpec{Type: "string"}},
			},
			IrreversibleNames: map[string]string{
				"enableEcsManagedTags": "EnableECSManagedTags",
			},
		},
	},
}

func TestSanitizeCfnString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "No special characters",
			input:    "Hello, World!",
			expected: "Hello, World!",
		},
		{
			name:     "Replace left and right double quotation marks",
			input:    "“Hello, World!”",
			expected: "\"Hello, World!\"",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := SanitizeCfnString(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestToStringifiedMap(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name:     "Nil input",
			input:    nil,
			expected: nil,
		},
		{
			name:     "Empty map",
			input:    map[string]interface{}{},
			expected: map[string]interface{}{},
		},
		{
			name: "Map with primitive values",
			input: map[string]interface{}{
				"string": "value",
				"int":    42,
				"float":  3.14,
				"bool":   true,
			},
			expected: map[string]interface{}{
				"string": "value",
				"int":    "42",
				"float":  "3.14",
				"bool":   "true",
			},
		},
		{
			name: "Map with nested map",
			input: map[string]interface{}{
				"nested": map[string]interface{}{
					"key": "value",
					"num": 123,
				},
			},
			expected: map[string]interface{}{
				"nested": map[string]interface{}{
					"key": "value",
					"num": "123",
				},
			},
		},
		{
			name: "Map with array",
			input: map[string]interface{}{
				"array": []interface{}{"a", 1, 2.5, false},
			},
			expected: map[string]interface{}{
				"array": []interface{}{"a", "1", "2.5", "false"},
			},
		},
		{
			name: "Map with mixed nested structures",
			input: map[string]interface{}{
				"level1": map[string]interface{}{
					"level2": []interface{}{
						map[string]interface{}{
							"key1": "value1",
							"key2": 2,
						},
						3.14,
						"string",
					},
					"anotherKey": true,
					"arrayOfMaps": []interface{}{
						map[string]interface{}{
							"key1": "value1",
							"key2": 2,
						},
						map[string]interface{}{
							"key3": "value3",
							"key4": 4,
						},
					},
				},
			},
			expected: map[string]interface{}{
				"level1": map[string]interface{}{
					"level2": []interface{}{
						map[string]interface{}{
							"key1": "value1",
							"key2": "2",
						},
						"3.14",
						"string",
					},
					"anotherKey": "true",
					"arrayOfMaps": []interface{}{
						map[string]interface{}{
							"key1": "value1",
							"key2": "2",
						},
						map[string]interface{}{
							"key3": "value3",
							"key4": "4",
						},
					},
				},
			},
		},
		{
			name: "Map with arbitrary keys and deeply nested structures",
			input: map[string]interface{}{
				"level1": map[interface{}]interface{}{
					123: "numberKey",
					true: map[string]interface{}{
						"nestedKey": []interface{}{
							map[string]interface{}{
								"key1": "value1",
								"key2": 2,
							},
							3.14,
							"string",
						},
					},
					"anotherKey": false,
				},
			},
			expected: map[string]interface{}{
				"level1": map[interface{}]interface{}{
					123: "numberKey",
					true: map[string]interface{}{
						"nestedKey": []interface{}{
							map[string]interface{}{
								"key1": "value1",
								"key2": "2",
							},
							"3.14",
							"string",
						},
					},
					"anotherKey": "false",
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := ToStringifiedMap(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
