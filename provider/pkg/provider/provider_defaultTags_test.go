package provider

import (
	"testing"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

func TestDefaultTags(t *testing.T) {
	defaultTags := map[string]string{
		"defaultTag": "defaultTagValue",
	}
	t.Run("string map", func(t *testing.T) {
		tags := resource.PropertyValue(resource.NewObjectProperty(map[resource.PropertyKey]resource.PropertyValue{}))
		expected := resource.PropertyValue(resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
			"defaultTag": "defaultTagValue",
		})))
		actual, err := mergeDefaultTags(tags, defaultTags, schema.TagsStyleStringMap)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("string map no input", func(t *testing.T) {
		var tags resource.PropertyValue
		expected := resource.PropertyValue(resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
			"defaultTag": "defaultTagValue",
		})))
		actual, err := mergeDefaultTags(tags, defaultTags, schema.TagsStyleStringMap)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("key value array", func(t *testing.T) {
		tags := resource.PropertyValue(resource.NewArrayProperty([]resource.PropertyValue{}))
		expected := resource.NewPropertyValue([]interface{}{
			map[string]interface{}{
				"key":   "defaultTag",
				"value": "defaultTagValue",
			},
		})
		actual, err := mergeDefaultTags(tags, defaultTags, schema.TagsStyleKeyValueArray)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}
