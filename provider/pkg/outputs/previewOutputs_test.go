package outputs

import (
	"testing"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

func TestPreviewOutputs(t *testing.T) {
	t.Run("Without prior state", func(t *testing.T) {
		result := PreviewOutputs(
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"name": "my-resource",
			}),
			map[string]metadata.CloudAPIType{},
			map[string]schema.PropertySpec{
				"name": schema.PropertySpec{
					TypeSpec: schema.TypeSpec{Type: "string"},
				},
				"arn": schema.PropertySpec{
					TypeSpec: schema.TypeSpec{Type: "string"},
				},
			},
			[]string{
				"arn",
			},
			"AccessPoint",
			nil,
		)
		assert.Equal(t, resource.PropertyMap{
			"name": resource.NewStringProperty("my-resource"),
			"arn":  resource.MakeComputed(resource.NewStringProperty("")),
		}, result)
	})

	t.Run("With prior state", func(t *testing.T) {
		priorOutputs := resource.NewPropertyMapFromMap(map[string]interface{}{
			"arn": "arnvalue",
		})
		result := PreviewOutputs(
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"name": "my-resource",
			}),
			map[string]metadata.CloudAPIType{},
			map[string]schema.PropertySpec{
				"name": schema.PropertySpec{
					TypeSpec: schema.TypeSpec{Type: "string"},
				},
				"arn": schema.PropertySpec{
					TypeSpec: schema.TypeSpec{Type: "string"},
				},
			},
			[]string{
				"arn",
			},
			"AccessPoint",
			&priorOutputs,
		)
		assert.Equal(t, resource.PropertyMap{
			"name": resource.NewStringProperty("my-resource"),
			"arn":  resource.NewStringProperty("arnvalue"),
		}, result)
	})
}

func Test_previewResourceOutputs(t *testing.T) {
	t.Run("Nested output value", func(t *testing.T) {
		result := previewResourceOutputs(
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
						"status": {
							TypeSpec: schema.TypeSpec{Type: "string"},
						},
						"value": {
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
			[]string{"alias", "alias/status", "alias/value"},
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
		result := previewResourceOutputs(
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
				"objectLambdaConfiguration/cloudWatchMetricsEnabled",
			},
		)
		assert.Equal(t, resource.PropertyMap{
			"objectLambdaConfiguration": resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
				"allowedFeatures":          []string{"GetObject-Range"},
				"cloudWatchMetricsEnabled": resource.MakeComputed(resource.NewStringProperty("")),
			})),
		}, result)
	})

	t.Run("with additionalProperties", func(t *testing.T) {
		result := previewResourceOutputs(
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"name": "my-access-point",
				"objectLambdaConfiguration": map[string]interface{}{
					"allowedFeatures": "GetObject-Range",
				},
			}),
			map[string]metadata.CloudAPIType{},
			map[string]schema.PropertySpec{
				"objectLambdaConfiguration": {
					TypeSpec: schema.TypeSpec{
						AdditionalProperties: &schema.TypeSpec{
							Type: "string",
						},
					},
				},
			},
			[]string{},
		)
		assert.Equal(t, resource.PropertyMap{
			"objectLambdaConfiguration": resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
				"allowedFeatures": "GetObject-Range",
			})),
		}, result)
	})

	t.Run("with additionalProperties Ref", func(t *testing.T) {
		result := previewResourceOutputs(
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
						AdditionalProperties: &schema.TypeSpec{
							Ref: "#/types/aws-native:s3objectlambda:AccessPointObjectLambdaConfiguration",
						},
					},
				},
			},
			[]string{
				"objectLambdaConfiguration/cloudWatchMetricsEnabled",
			},
		)
		assert.Equal(t, resource.PropertyMap{
			"objectLambdaConfiguration": resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
				"allowedFeatures":          []string{"GetObject-Range"},
				"cloudWatchMetricsEnabled": resource.MakeComputed(resource.NewStringProperty("")),
			})),
		}, result)
	})

	t.Run("with additionalProperties Items", func(t *testing.T) {
		result := previewResourceOutputs(
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
								AdditionalProperties: &schema.TypeSpec{
									Type:  "array",
									Items: &schema.TypeSpec{Type: "string"},
								},
							},
						},
					},
				},
			},
			map[string]schema.PropertySpec{
				"objectLambdaConfiguration": {
					TypeSpec: schema.TypeSpec{
						AdditionalProperties: &schema.TypeSpec{
							Ref: "#/types/aws-native:s3objectlambda:AccessPointObjectLambdaConfiguration",
						},
					},
				},
			},
			[]string{
				"objectLambdaConfiguration/cloudWatchMetricsEnabled",
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
		result := previewResourceOutputs(
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
				"subscribers/*/status",
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

	t.Run("with secret values", func(t *testing.T) {
		result := previewResourceOutputs(
			resource.PropertyMap{
				"subscribers": resource.NewSecretProperty(&resource.Secret{Element: resource.NewArrayProperty(
					[]resource.PropertyValue{
						resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
							"address": "address",
						})),
					},
				)}),
			},
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
				"subscribers/*/status",
			},
		)
		assert.Equal(t, resource.PropertyMap{
			"subscribers": resource.NewSecretProperty(&resource.Secret{Element: resource.NewArrayProperty(
				[]resource.PropertyValue{
					resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
						"address": "address",
						"status":  resource.MakeComputed(resource.NewStringProperty("")),
					})),
				},
			)}),
		}, result)
	})

	t.Run("with output values", func(t *testing.T) {
		result := previewResourceOutputs(
			resource.PropertyMap{
				"subscribers": resource.NewOutputProperty(resource.Output{Element: resource.NewArrayProperty(
					[]resource.PropertyValue{
						resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
							"address": "address",
						})),
					},
				)}),
			},
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
				"subscribers/*/status",
			},
		)
		assert.Equal(t, resource.PropertyMap{
			"subscribers": resource.NewOutputProperty(resource.Output{Element: resource.NewArrayProperty(
				[]resource.PropertyValue{
					resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
						"address": "address",
						"status":  resource.MakeComputed(resource.NewStringProperty("")),
					})),
				},
			)}),
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
				"arn",
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
				"objectLambdaConfiguration/accessPointArn",
				"alias",
				"alias/status",
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
				"objectLambdaConfiguration/arn",
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
				"objectLambdaConfiguration/id",
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
				"objectLambdaConfiguration/accessPointId",
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
				"objectLambdaConfiguration/accessPointName",
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
				"objectLambdaConfiguration/name",
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
				"objectLambdaConfiguration/arrayValue/*/arn",
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
