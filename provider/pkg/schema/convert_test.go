// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	"sort"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

func TestCfnToSdk(t *testing.T) {
	actual := CfnToSdk(cfnPayload)
	assert.Equal(t, sdkState, actual)
}

func TestSdkToCfn(t *testing.T) {
	actual, err := SdkToCfn(sampleSchema, "AWS::ECS::Service", sdkState)
	assert.NoError(t, err)
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
	expected := []types.PatchOperation{
		{Op: types.Op("replace"), Path: aws.String("/DesiredCount"), IntegerValue: &two},
		{Op: types.Op("replace"), Path: aws.String("/EnableECSManagedTags"), BooleanValue: aws.Bool(true)},
		{Op: types.Op("add"), Path: aws.String("/LaunchType"), StringValue: aws.String("FARGATE")},
		{
			Op:          types.Op("replace"),
			Path:        aws.String("/LoadBalancers"),
			ObjectValue: aws.String("[{\"ContainerName\":\"my-app\",\"ContainerPort\":80}]"),
		},
		{Op: types.Op("remove"), Path: aws.String("/PlatformVersion")},
	}
	actual, err := DiffToPatch(sampleSchema, "AWS::ECS::Service", &diff)
	assert.NoError(t, err)
	sort.SliceStable(actual, func(i, j int) bool {
		return *actual[i].Path < *actual[j].Path
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

var sampleSchema = CloudFormationSchema{
	PropertyTypes: map[string]PropertyTypeSpec{
		"AWS::ECS::Service.AwsVpcConfiguration": {
			Properties: map[string]PropertySpec{
				"AssignPublicIp": {PrimitiveType: "String"},
				"SecurityGroups": {
					Type:              "List",
					PrimitiveItemType: "String",
				},
				"Subnets": {
					Type:              "List",
					PrimitiveItemType: "String",
				},
			},
		},
		"AWS::ECS::Service.DeploymentCircuitBreaker": {
			Properties: map[string]PropertySpec{
				"Enable":   {PrimitiveType: "Boolean"},
				"Rollback": {PrimitiveType: "Boolean"},
			},
		},
		"AWS::ECS::Service.DeploymentConfiguration": {
			Properties: map[string]PropertySpec{
				"DeploymentCircuitBreaker": {Type: "DeploymentCircuitBreaker"},
				"MaximumPercent":           {PrimitiveType: "Integer"},
				"MinimumHealthyPercent":    {PrimitiveType: "Integer"},
			},
		},
		"AWS::ECS::Service.LoadBalancer": {
			Properties: map[string]PropertySpec{
				"ContainerName":    {PrimitiveType: "String"},
				"ContainerPort":    {PrimitiveType: "Integer"},
				"LoadBalancerName": {PrimitiveType: "String"},
				"TargetGroupArn":   {PrimitiveType: "String"},
			},
		},
		"AWS::ECS::Service.NetworkConfiguration": {
			Properties: map[string]PropertySpec{
				"AwsvpcConfiguration": {Type: "AwsVpcConfiguration"},
			},
		},
	},
	ResourceTypes: map[string]ResourceSpec{
		"AWS::ECS::Service": {
			PropertyTypeSpec: PropertyTypeSpec{
				Properties: map[string]PropertySpec{
					"ServiceArn":              {PrimitiveType: "String"},
					"Cluster":                 {PrimitiveType: "String"},
					"DeploymentConfiguration": {Type: "DeploymentConfiguration"},
					"DesiredCount":            {PrimitiveType: "Integer"},
					"EnableECSManagedTags":    {PrimitiveType: "Boolean"},
					"LaunchType":              {PrimitiveType: "String"},
					"LoadBalancers": {
						Type:     "List",
						ItemType: "LoadBalancer",
					},
					"NetworkConfiguration": {Type: "NetworkConfiguration"},
					"PlatformVersion":      {PrimitiveType: "String"},
					"Role":                 {PrimitiveType: "String"},
					"SchedulingStrategy":   {PrimitiveType: "String"},
					"ServiceName":          {PrimitiveType: "String"},
					"TaskDefinition":       {PrimitiveType: "String"},
				},
			},
			Attributes: map[string]AttributeSpec{
				"Name": {PrimitiveType: "String"},
			},
		},
	},
}
