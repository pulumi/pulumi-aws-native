// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	"github.com/mattbaird/jsonpatch"
	"sort"
	"testing"

	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

func TestCfnToSdk(t *testing.T) {
	actual := CfnToSdk(cfnPayload)
	assert.Equal(t, sdkState, actual)
}

func TestSdkToCfn(t *testing.T) {
	res := sampleSchema.Resources["aws-native:ecs:Service"]
	actual := SdkToCfn(&res, sampleSchema.Types, sdkState)
	assert.Equal(t, cfnPayload, actual)
}

func TestDiffToPatch(t *testing.T) {
	diff := resource.ObjectDiff{
		Updates: map[resource.PropertyKey]resource.ValueDiff{
			"desiredCount":         {New: resource.NewNumberProperty(2)},
			"enableECSManagedTags": {New: resource.NewBoolProperty(true)},
			"loadBalancers": {
				New: resource.NewArrayProperty([]resource.PropertyValue{
					resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
						"containerName": "my-app",
						"containerPort": 80,
					})),
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
	res := sampleSchema.Resources["aws-native:ecs:Service"]
	actual, err := DiffToPatch(&res, sampleSchema.Types, &diff)
	assert.NoError(t, err)
	sort.SliceStable(actual, func(i, j int) bool {
		return actual[i].Path < actual[j].Path
	})
	assert.Equal(t, expected, actual)
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
	"enableECSManagedTags": false,
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

var sampleSchema = &CloudAPIMetadata{
	Types: map[string]CloudAPIType{
		"aws-native:ecs:AwsVpcConfiguration": {
			Properties: map[string]pschema.PropertySpec{
				"assignPublicIp": {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"securityGroups": {
					TypeSpec: pschema.TypeSpec{
						Type: "array",
						Items: &pschema.TypeSpec{Type: "string"},
					},
				},
				"subnets": {
					TypeSpec: pschema.TypeSpec{
						Type: "array",
						Items: &pschema.TypeSpec{Type: "string"},
					},
				},
			},
		},
		"aws-native:ecs:DeploymentCircuitBreaker": {
			Properties: map[string]pschema.PropertySpec{
				"enable":   {TypeSpec: pschema.TypeSpec{Type: "boolean"}},
				"rollback": {TypeSpec: pschema.TypeSpec{Type: "boolean"}},
			},
		},
		"aws-native:ecs:DeploymentConfiguration": {
			Properties: map[string]pschema.PropertySpec{
				"deploymentCircuitBreaker": {
					TypeSpec: pschema.TypeSpec{
						Ref: "#/types/aws-native:ecs:DeploymentCircuitBreaker",
					},
				},
				"maximumPercent":           {TypeSpec: pschema.TypeSpec{Type: "integer"}},
				"minimumHealthyPercent":    {TypeSpec: pschema.TypeSpec{Type: "integer"}},
			},
		},
		"aws-native:ecs:LoadBalancer": {
			Properties: map[string]pschema.PropertySpec{
				"containerName":    {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"containerPort":    {TypeSpec: pschema.TypeSpec{Type: "integer"}},
				"loadBalancerName": {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"targetGroupArn":   {TypeSpec: pschema.TypeSpec{Type: "string"}},
			},
		},
		"aws-native:ecs:NetworkConfiguration": {
			Properties: map[string]pschema.PropertySpec{
				"awsvpcConfiguration": {
					TypeSpec: pschema.TypeSpec{
						Ref: "#/types/aws-native:ecs:AwsVpcConfiguration",
					},
				},
			},
		},
	},
	Resources: map[string]CloudAPIResource{
		"aws-native:ecs:Service": {
			Inputs: map[string]pschema.PropertySpec{
				"serviceArn":              {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"cluster":                 {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"deploymentConfiguration": {
					TypeSpec: pschema.TypeSpec{
						Ref: "#/types/aws-native:ecs:DeploymentConfiguration",
					},
				},
				"desiredCount":            {TypeSpec: pschema.TypeSpec{Type: "integer"}},
				"enableECSManagedTags":    {TypeSpec: pschema.TypeSpec{Type: "boolean"}},
				"launchType":              {TypeSpec: pschema.TypeSpec{Type: "string"}},
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
				"platformVersion":      {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"role":                 {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"schedulingStrategy":   {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"serviceName":          {TypeSpec: pschema.TypeSpec{Type: "string"}},
				"taskDefinition":       {TypeSpec: pschema.TypeSpec{Type: "string"}},
			},
			Outputs: map[string]pschema.PropertySpec{
				"name": {TypeSpec: pschema.TypeSpec{Type: "string"}},
			},
		},
	},
}
