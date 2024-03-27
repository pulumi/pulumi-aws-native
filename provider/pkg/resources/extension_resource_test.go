// Copyright 2024, Pulumi Corporation.

package resources

import (
	"testing"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/default_tags"
	"github.com/stretchr/testify/assert"
)

func TestExtensionResourceDefaults(t *testing.T) {
	t.Run("Defaults no change", func(t *testing.T) {
		r, failures := ApplyDefaults(ExtensionResourceInputs{})
		assert.Empty(t, failures)
		assert.Equal(t, ExtensionResourceInputs{}, r)
	})
	t.Run("Only tag prop", func(t *testing.T) {
		r, failures := ApplyDefaults(ExtensionResourceInputs{
			TagsProperty: "CustomTags",
		})
		assert.Empty(t, failures)
		assert.Equal(t, ExtensionResourceInputs{
			TagsProperty: "CustomTags",
			TagsStyle:    default_tags.TagsStyleKeyValueArray,
		}, r)
	})
	t.Run("Only tag style", func(t *testing.T) {
		r, failures := ApplyDefaults(ExtensionResourceInputs{
			TagsStyle: default_tags.TagsStyleStringMap,
		})
		assert.Empty(t, failures)
		assert.Equal(t, ExtensionResourceInputs{
			TagsProperty: "Tags",
			TagsStyle:    default_tags.TagsStyleStringMap,
		}, r)
	})
	t.Run("None tag style", func(t *testing.T) {
		r, failures := ApplyDefaults(ExtensionResourceInputs{
			TagsStyle: default_tags.TagsStyleNone,
		})
		assert.Empty(t, failures)
		assert.Equal(t, ExtensionResourceInputs{
			TagsStyle: default_tags.TagsStyleNone,
		}, r)
	})
	t.Run("Both tag prop and style", func(t *testing.T) {
		r, failures := ApplyDefaults(ExtensionResourceInputs{
			TagsProperty: "CustomTags",
			TagsStyle:    default_tags.TagsStyleStringMap,
		})
		assert.Empty(t, failures)
		assert.Equal(t, ExtensionResourceInputs{
			TagsProperty: "CustomTags",
			TagsStyle:    default_tags.TagsStyleStringMap,
		}, r)
	})
	t.Run("Invalid tag style", func(t *testing.T) {
		_, failures := ApplyDefaults(ExtensionResourceInputs{
			TagsStyle: "invalid",
		})
		assert.Equal(t, []ValidationFailure{{Path: "tagsStyle", Reason: "tagsStyle is invalid, must be one of: stringMap, keyValueArray, none"}}, failures)
	})
}
