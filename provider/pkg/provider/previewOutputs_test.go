package provider

import (
	"testing"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

func TestPrevierOutputs(t *testing.T) {
	t.Run("Nested output value", func(t *testing.T) {
		result := PreviewOutputs(
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"name": "my-access-point",
				"objectLambdaConfiguration": map[string]interface{}{
					"allowedFeatures":          []string{"GetObject-Range"},
					"cloudWatchMetricsEnabled": true,
				},
			}),
			map[string]metadata.CloudAPIType{
				"aws-native:s3objectlambda:AccessPointAlias": {
					Properties: map[string]schema.PropertySpec{
						"Status": {
							TypeSpec: schema.TypeSpec{Type: "string"},
						},
						"Value": {
							TypeSpec: schema.TypeSpec{Type: "string"},
						},
					},
				},
				"aws-native:s3objectlambda:AccessPointObjectLambdaConfiguration": {
					Properties: map[string]schema.PropertySpec{
						"cloudWatchMetricsEnabled": {
							TypeSpec: schema.TypeSpec{Type: "boolean"},
						},
						"allowedFeatures": {
							TypeSpec: schema.TypeSpec{
								Type:  "array",
								Items: &schema.TypeSpec{Type: "string"},
							},
						},
					},
				},
			},
			map[string]schema.PropertySpec{
				"alias": {
					TypeSpec: schema.TypeSpec{
						Ref: "#/types/aws-native:s3objectlambda:AccessPointAlias",
					},
				},
				"objectLambdaConfiguration": {
					TypeSpec: schema.TypeSpec{
						Ref: "#/types/aws-native:s3objectlambda:AccessPointObjectLambdaConfiguration",
					},
				},
			},
			[]string{"alias", "alias/Status", "alias/Value"},
		)
		assert.Equal(t, resource.PropertyMap{
			"alias": resource.MakeComputed(resource.NewStringProperty("")),
			"objectLambdaConfiguration": resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
				"allowedFeatures":          []string{"GetObject-Range"},
				"cloudWatchMetricsEnabled": true,
			})),
		}, result)
	})

	t.Run("Mixed input and output types", func(t *testing.T) {
		result := PreviewOutputs(
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"name": "my-access-point",
				"objectLambdaConfiguration": map[string]interface{}{
					"allowedFeatures": []string{"GetObject-Range"},
				},
			}),
			map[string]metadata.CloudAPIType{
				"aws-native:s3objectlambda:AccessPointObjectLambdaConfiguration": {
					Properties: map[string]schema.PropertySpec{
						"cloudWatchMetricsEnabled": {
							TypeSpec: schema.TypeSpec{Type: "boolean"},
						},
						"allowedFeatures": {
							TypeSpec: schema.TypeSpec{
								Type:  "array",
								Items: &schema.TypeSpec{Type: "string"},
							},
						},
					},
				},
			},
			map[string]schema.PropertySpec{
				"objectLambdaConfiguration": {
					TypeSpec: schema.TypeSpec{
						Ref: "#/types/aws-native:s3objectlambda:AccessPointObjectLambdaConfiguration",
					},
				},
			},
			[]string{
				"objectLambdaConfiguration/CloudWatchMetricsEnabled",
			},
		)
		assert.Equal(t, resource.PropertyMap{
			"objectLambdaConfiguration": resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
				"allowedFeatures":          []string{"GetObject-Range"},
				"cloudWatchMetricsEnabled": resource.MakeComputed(resource.NewStringProperty("")),
			})),
		}, result)
	})

	t.Run("Array with mixed input output types", func(t *testing.T) {
		result := PreviewOutputs(
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"subscribers": []map[string]interface{}{
					{
						"address": "address",
					},
				},
			}),
			map[string]metadata.CloudAPIType{
				"aws-native:ce:AnomalySubscriptionSubscriber": {
					Type: "object",
					Properties: map[string]schema.PropertySpec{
						"address": {
							TypeSpec: schema.TypeSpec{Type: "string"},
						},
						"status": {
							TypeSpec: schema.TypeSpec{Type: "string"},
						},
						"type": {
							TypeSpec: schema.TypeSpec{Type: "string"},
						},
					},
				},
			},
			map[string]schema.PropertySpec{
				"subscribers": {
					TypeSpec: schema.TypeSpec{
						Type: "array",
						Items: &schema.TypeSpec{
							Ref: "#/types/aws-native:ce:AnomalySubscriptionSubscriber",
						},
					},
				},
			},
			[]string{
				"subscribers/*/Status",
			},
		)
		assert.Equal(t, resource.PropertyMap{
			"subscribers": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
					"address": "address",
					"status":  resource.MakeComputed(resource.NewStringProperty("")),
				})),
			}),
		}, result)
	})

	t.Run("Custom resource", func(t *testing.T) {
		result := PreviewOutputs(
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"name": "my-access-point",
				"objectLambdaConfiguration": map[string]interface{}{
					"allowedFeatures":          []string{"GetObject-Range"},
					"cloudWatchMetricsEnabled": true,
					"supportingAccessPoint":    "arn:aws:s3:us-west-2:123456789012:accesspoint/my-supporting-ap",
				},
			}),
			nil,
			nil,
			[]string{},
		)
		assert.Equal(t, resource.PropertyMap{
			"outputs": resource.MakeComputed(resource.NewStringProperty("")),
		}, result)
	})
}

func Test_updatePreviewWithOutputs(t *testing.T) {
	t.Run("Top level stable output", func(t *testing.T) {
		result := populateStableOutputs(
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"input": "value",
				"arn":   resource.MakeComputed(resource.NewStringProperty("")),
			}),
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"input": "value",
				"arn":   "arn:aws:s3-object-lambda:us-west-2:123456789012:accesspoint/my-access-point",
			}),
			[]string{
				"Arn",
			},
			"accesspoint",
		)
		assert.Equal(t, resource.PropertyMap{
			"input": resource.NewStringProperty("value"),
			"arn":   resource.NewStringProperty("arn:aws:s3-object-lambda:us-west-2:123456789012:accesspoint/my-access-point"),
		}, result)
	})

	t.Run("Nested output value resourceNameArn", func(t *testing.T) {
		result := populateStableOutputs(
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"objectLambdaConfiguration": map[string]interface{}{
					"allowedFeatures": []string{"GetObject-Range"},
					"accessPointArn":  resource.MakeComputed(resource.NewStringProperty("")),
				},
				"alias": map[string]interface{}{
					"status": resource.MakeComputed(resource.NewStringProperty("")),
				},
			}),
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"objectLambdaConfiguration": map[string]interface{}{
					"allowedFeatures": []string{"GetObject-Range"},
					"accessPointArn":  "arn:aws:s3-object-lambda:us-west-2:123456789012:accesspoint/my-access-point",
				},
				"alias": map[string]interface{}{
					"status": "ACTIVE",
				},
			}),
			[]string{
				"objectLambdaConfiguration/AccessPointArn",
				"alias",
				"alias/Status",
			},
			"AccessPoint",
		)
		assert.Equal(t, resource.PropertyMap{
			"alias": resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
				"status": resource.MakeComputed(resource.NewStringProperty("")),
			})),
			"objectLambdaConfiguration": resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
				"allowedFeatures": []string{"GetObject-Range"},
				"accessPointArn":  "arn:aws:s3-object-lambda:us-west-2:123456789012:accesspoint/my-access-point",
			})),
		}, result)
	})

	t.Run("Nested output value arn isn't stable", func(t *testing.T) {
		result := populateStableOutputs(
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"objectLambdaConfiguration": map[string]interface{}{
					"arn": resource.MakeComputed(resource.NewStringProperty("")),
				},
			}),
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"objectLambdaConfiguration": map[string]interface{}{
					"arn": "arn:aws:s3-object-lambda:us-west-2:123456789012:accesspoint/my-access-point",
				},
			}),
			[]string{
				"objectLambdaConfiguration/Arn",
			},
			"AccessPoint",
		)
		assert.Equal(t, resource.PropertyMap{
			"objectLambdaConfiguration": resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
				"arn": resource.MakeComputed(resource.NewStringProperty("")),
			})),
		}, result)
	})

	t.Run("Nested output value id isn't stable", func(t *testing.T) {
		result := populateStableOutputs(
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"objectLambdaConfiguration": map[string]interface{}{
					"id": resource.MakeComputed(resource.NewStringProperty("")),
				},
			}),
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"objectLambdaConfiguration": map[string]interface{}{
					"id": "my-access-point",
				},
			}),
			[]string{
				"objectLambdaConfiguration/Id",
			},
			"AccessPoint",
		)
		assert.Equal(t, resource.PropertyMap{
			"objectLambdaConfiguration": resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
				"id": resource.MakeComputed(resource.NewStringProperty("")),
			})),
		}, result)
	})

	t.Run("Nested output value resourceNameId", func(t *testing.T) {
		result := populateStableOutputs(
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"objectLambdaConfiguration": map[string]interface{}{
					"accessPointId": resource.MakeComputed(resource.NewStringProperty("")),
				},
			}),
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"objectLambdaConfiguration": map[string]interface{}{
					"accessPointId": "my-access-point",
				},
			}),
			[]string{
				"objectLambdaConfiguration/AccessPointId",
			},
			"AccessPoint",
		)
		assert.Equal(t, resource.PropertyMap{
			"objectLambdaConfiguration": resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
				"accessPointId": "my-access-point",
			})),
		}, result)
	})

	t.Run("Nested output value resourceNameName", func(t *testing.T) {
		result := populateStableOutputs(
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"objectLambdaConfiguration": map[string]interface{}{
					"accessPointName": resource.MakeComputed(resource.NewStringProperty("")),
				},
			}),
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"objectLambdaConfiguration": map[string]interface{}{
					"accessPointName": "my-access-point",
				},
			}),
			[]string{
				"objectLambdaConfiguration/AccessPointName",
			},
			"AccessPoint",
		)
		assert.Equal(t, resource.PropertyMap{
			"objectLambdaConfiguration": resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
				"accessPointName": "my-access-point",
			})),
		}, result)
	})

	t.Run("Nested output value name isn't  stable", func(t *testing.T) {
		result := populateStableOutputs(
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"objectLambdaConfiguration": map[string]interface{}{
					"name": resource.MakeComputed(resource.NewStringProperty("")),
				},
			}),
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"objectLambdaConfiguration": map[string]interface{}{
					"name": "my-access-point",
				},
			}),
			[]string{
				"objectLambdaConfiguration/Name",
			},
			"AccessPoint",
		)
		assert.Equal(t, resource.PropertyMap{
			"objectLambdaConfiguration": resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
				"name": resource.MakeComputed(resource.NewStringProperty("")),
			})),
		}, result)
	})

	t.Run("only readonly properties can be stable", func(t *testing.T) {
		result := populateStableOutputs(
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"name": resource.MakeComputed(resource.NewStringProperty("")),
			}),
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"name": "my-access-point",
			}),
			[]string{},
			"AccessPoint",
		)
		assert.Equal(t, resource.PropertyMap{
			"name": resource.MakeComputed(resource.NewStringProperty("")),
		}, result)
	})

	t.Run("Array properties ignored", func(t *testing.T) {
		result := populateStableOutputs(
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"objectLambdaConfiguration": map[string]interface{}{
					"arrayValue": []map[string]interface{}{
						{
							"allowedFeatures": []string{"GetObject-Range"},
							"arn":             resource.MakeComputed(resource.NewStringProperty("")),
						},
					},
				},
			}),
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"objectLambdaConfiguration": map[string]interface{}{
					"arrayValue": []map[string]interface{}{
						{
							"allowedFeatures": []string{"GetObject-Range"},
							"arn":             "arn:aws:s3-object-lambda:us-west-2:123456789012:accesspoint/my-access-point",
						},
					},
				},
			}),
			[]string{
				"objectLambdaConfiguration/ArrayValue/*/Arn",
			},
			"AccessPoint",
		)
		assert.Equal(t, resource.PropertyMap{
			"objectLambdaConfiguration": resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
				"arrayValue": []map[string]interface{}{
					{
						"allowedFeatures": []string{"GetObject-Range"},
						"arn":             resource.MakeComputed(resource.NewStringProperty("")),
					},
				},
			})),
		}, result)

	})
}
