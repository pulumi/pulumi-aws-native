// Copyright 2026, Pulumi Corporation.

package resources

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func TestNormalizeWafv2ByteMatchInputs(t *testing.T) {
	inputs := resource.PropertyMap{
		wafv2RulesProperty: resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				wafv2StatementProperty: resource.NewObjectProperty(resource.PropertyMap{
					wafv2ByteMatchProperty: resource.NewObjectProperty(resource.PropertyMap{
						wafv2SearchStringProperty: resource.NewStringProperty("/login"),
					}),
				}),
			}),
		}),
	}

	NormalizeWafv2ByteMatchInputs(awsNativeWafv2WebAclToken, inputs)

	byteMatch := inputs[wafv2RulesProperty].ArrayValue()[0].ObjectValue()[wafv2StatementProperty].
		ObjectValue()[wafv2ByteMatchProperty].ObjectValue()
	assert.NotContains(t, byteMatch, resource.PropertyKey(wafv2SearchStringProperty))
	assert.Equal(t, "L2xvZ2lu", byteMatch[wafv2SearchStringBase64Prop].StringValue())
}

func TestNormalizeWafv2ByteMatchInputsPreservesExplicitBoth(t *testing.T) {
	byteMatch := resource.PropertyMap{
		wafv2SearchStringProperty:   resource.NewStringProperty("/login"),
		wafv2SearchStringBase64Prop: resource.NewStringProperty("L2xvZ2lu"),
	}
	inputs := resource.PropertyMap{
		wafv2RulesProperty: resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(byteMatch),
		}),
	}

	NormalizeWafv2ByteMatchInputs(awsNativeWafv2WebAclToken, inputs)

	assert.Contains(t, byteMatch, resource.PropertyKey(wafv2SearchStringProperty))
	assert.Contains(t, byteMatch, resource.PropertyKey(wafv2SearchStringBase64Prop))
}

func TestNormalizeWafv2ByteMatchBaseline(t *testing.T) {
	newByteMatch := func() resource.PropertyValue {
		return resource.NewObjectProperty(resource.PropertyMap{
			wafv2SearchStringProperty:   resource.NewStringProperty("/login"),
			wafv2SearchStringBase64Prop: resource.NewStringProperty("L2xvZ2lu"),
		})
	}

	t.Run("keeps base64 when desired uses base64", func(t *testing.T) {
		actual := resource.PropertyMap{wafv2StatementProperty: newByteMatch()}
		desired := resource.PropertyMap{
			wafv2StatementProperty: resource.NewObjectProperty(resource.PropertyMap{
				wafv2SearchStringBase64Prop: resource.NewStringProperty("L2xvZ2lu"),
			}),
		}

		NormalizeWafv2ByteMatchBaseline(awsNativeWafv2WebAclToken, actual, desired)

		byteMatch := actual[wafv2StatementProperty].ObjectValue()
		assert.NotContains(t, byteMatch, resource.PropertyKey(wafv2SearchStringProperty))
		assert.Contains(t, byteMatch, resource.PropertyKey(wafv2SearchStringBase64Prop))
	})

	t.Run("keeps plain string when desired uses plain string", func(t *testing.T) {
		actual := resource.PropertyMap{wafv2StatementProperty: newByteMatch()}
		desired := resource.PropertyMap{
			wafv2StatementProperty: resource.NewObjectProperty(resource.PropertyMap{
				wafv2SearchStringProperty: resource.NewStringProperty("/login"),
			}),
		}

		NormalizeWafv2ByteMatchBaseline(awsNativeWafv2WebAclToken, actual, desired)

		byteMatch := actual[wafv2StatementProperty].ObjectValue()
		assert.Contains(t, byteMatch, resource.PropertyKey(wafv2SearchStringProperty))
		assert.NotContains(t, byteMatch, resource.PropertyKey(wafv2SearchStringBase64Prop))
	})

	t.Run("does not alter a non-equivalent pair", func(t *testing.T) {
		actual := resource.PropertyMap{
			wafv2StatementProperty: resource.NewObjectProperty(resource.PropertyMap{
				wafv2SearchStringProperty:   resource.NewStringProperty("/login"),
				wafv2SearchStringBase64Prop: resource.NewStringProperty("ZGlmZmVyZW50"),
			}),
		}

		NormalizeWafv2ByteMatchBaseline(awsNativeWafv2WebAclToken, actual, resource.PropertyMap{})

		byteMatch := actual[wafv2StatementProperty].ObjectValue()
		assert.Contains(t, byteMatch, resource.PropertyKey(wafv2SearchStringProperty))
		assert.Contains(t, byteMatch, resource.PropertyKey(wafv2SearchStringBase64Prop))
	})
}
