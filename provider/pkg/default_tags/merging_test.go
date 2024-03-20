package default_tags

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

func TestDefaultTags(t *testing.T) {
	defaultTags := map[string]string{
		"defaultTag": "defaultTagValue",
	}

	t.Run("string map nil resource tags", func(t *testing.T) {
		var tags resource.PropertyValue
		expected := resource.NewPropertyValue(map[string]interface{}{
			"defaultTag": "defaultTagValue",
		})
		actual, err := MergeDefaultTags(tags, defaultTags, TagsStyleStringMap)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("string map empty resource tags", func(t *testing.T) {
		tags := resource.NewPropertyValue(resource.NewObjectProperty(map[resource.PropertyKey]resource.PropertyValue{}))
		expected := resource.PropertyValue(resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
			"defaultTag": "defaultTagValue",
		})))
		actual, err := MergeDefaultTags(tags, defaultTags, TagsStyleStringMap)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("string map nil default tags", func(t *testing.T) {
		var defaultTags map[string]string
		tags := resource.NewPropertyValue(map[string]interface{}{
			"localTag": "localTagValue",
		})
		expected := resource.NewPropertyValue(map[string]interface{}{
			"localTag": "localTagValue",
		})
		actual, err := MergeDefaultTags(tags, defaultTags, TagsStyleStringMap)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("string map resource overrides default", func(t *testing.T) {
		tags := resource.NewPropertyValue(map[string]interface{}{
			"defaultTag": "localTagValue",
		})
		expected := resource.NewPropertyValue(map[string]interface{}{
			"defaultTag": "localTagValue",
		})
		actual, err := MergeDefaultTags(tags, defaultTags, TagsStyleStringMap)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("string map merge tags", func(t *testing.T) {
		tags := resource.NewPropertyValue(map[string]interface{}{
			"localTag": "localTagValue",
		})
		expected := resource.NewPropertyValue(map[string]interface{}{
			"defaultTag": "defaultTagValue",
			"localTag":   "localTagValue",
		})
		actual, err := MergeDefaultTags(tags, defaultTags, TagsStyleStringMap)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("key value array nil resource tags", func(t *testing.T) {
		var tags resource.PropertyValue
		expected := resource.NewPropertyValue([]interface{}{
			map[string]interface{}{
				"key":   "defaultTag",
				"value": "defaultTagValue",
			},
		})
		actual, err := MergeDefaultTags(tags, defaultTags, TagsStyleKeyValueArray)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("key value array empty resource tags", func(t *testing.T) {
		tags := resource.NewPropertyValue([]interface{}{})
		expected := resource.NewPropertyValue([]interface{}{
			map[string]interface{}{
				"key":   "defaultTag",
				"value": "defaultTagValue",
			},
		})
		actual, err := MergeDefaultTags(tags, defaultTags, TagsStyleKeyValueArray)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("key value array nil default tags", func(t *testing.T) {
		var defaultTags map[string]string
		tags := resource.NewPropertyValue([]interface{}{
			map[string]interface{}{
				"key":   "localTag",
				"value": "localTagValue",
			},
		})
		expected := resource.NewPropertyValue([]interface{}{
			map[string]interface{}{
				"key":   "localTag",
				"value": "localTagValue",
			},
		})
		actual, err := MergeDefaultTags(tags, defaultTags, TagsStyleKeyValueArray)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("key value array resource overrides default", func(t *testing.T) {
		tags := resource.NewPropertyValue([]interface{}{
			map[string]interface{}{
				"key":   "defaultTag",
				"value": "localTagValue",
			},
		})
		expected := resource.NewPropertyValue([]interface{}{
			map[string]interface{}{
				"key":   "defaultTag",
				"value": "localTagValue",
			},
		})
		actual, err := MergeDefaultTags(tags, defaultTags, TagsStyleKeyValueArray)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("key value array merge tags", func(t *testing.T) {
		tags := resource.NewPropertyValue([]interface{}{
			map[string]interface{}{
				"key":   "localTag",
				"value": "localTagValue",
			},
		})
		expected := resource.NewPropertyValue([]interface{}{
			map[string]interface{}{
				"key":   "localTag",
				"value": "localTagValue",
			},
			map[string]interface{}{
				"key":   "defaultTag",
				"value": "defaultTagValue",
			},
		})
		actual, err := MergeDefaultTags(tags, defaultTags, TagsStyleKeyValueArray)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("key value array uppercase", func(t *testing.T) {
		tags := resource.NewPropertyValue([]interface{}{
			map[string]interface{}{
				"Key":   "localTag",
				"Value": "localTagValue",
			},
		})
		expected := resource.NewPropertyValue([]interface{}{
			map[string]interface{}{
				"Key":   "localTag",
				"Value": "localTagValue",
			},
			map[string]interface{}{
				"Key":   "defaultTag",
				"Value": "defaultTagValue",
			},
		})
		actual, err := MergeDefaultTags(tags, defaultTags, TagsStyleKeyValueArrayUpperCase)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}
