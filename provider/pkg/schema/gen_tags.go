package schema

import (
	"strings"

	jsschema "github.com/lestrrat-go/jsschema"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

type TagsStyle string

const (
	TagsStyleUnknown   TagsStyle = ""
	TagsStyleUntyped   TagsStyle = "untyped"
	TagsStyleStringMap TagsStyle = "stringMap"
	// TagsStyleKeyValueArray is a style where the tags are represented as an array of key-value pairs.
	TagsStyleKeyValueArray TagsStyle = "keyValueArray"
)

func GetTagsProperty(originalSpec *jsschema.Schema) (string, bool) {
	tagging, ok := originalSpec.Root().Extras["tagging"]
	if !ok {
		return "", false
	}
	asInterface, ok := tagging.(map[string]interface{})
	if !ok {
		return "", false
	}
	tagProperty, ok := asInterface["tagProperty"]
	if !ok {
		return "", false
	}
	tagPropertyString, ok := tagProperty.(string)
	if !ok {
		return "", false
	}
	return strings.TrimPrefix(tagPropertyString, "/properties/"), true
}

func (ctx *context) GetTagsStyle(typeSpec *pschema.TypeSpec, originalSpec *jsschema.Schema) TagsStyle {
	if typeSpec == nil || originalSpec == nil {
		return TagsStyleUnknown
	}
	// Check for "Any" ref
	if typeSpec.Ref == "pulumi.json#/Any" {
		return TagsStyleUntyped
	}

	if typeSpec.AdditionalProperties != nil && typeSpec.AdditionalProperties.Type == "string" {
		return TagsStyleStringMap
	}

	if ctx.TagStyleIsKeyValueArray(typeSpec, originalSpec) {
		return TagsStyleKeyValueArray
	}

	return TagsStyleUnknown
}

func (ctx *context) TagStyleIsKeyValueArray(typeSpec *pschema.TypeSpec, originalSpec *jsschema.Schema) bool {
	if typeSpec == nil || typeSpec.Items == nil || typeSpec.Items.Ref == "" {
		return false
	}

	typeToken, hasPrefix := strings.CutPrefix(typeSpec.Items.Ref, "#/types/")
	if !hasPrefix {
		return false
	}

	// We can't include tags which are create-only properties because this has to be added on the tags type but this is a shared type
	if createOnlyProps := readPropSdkNames(ctx.resourceSpec, "createOnlyProperties"); createOnlyProps.Has("tags") {
		return false
	}

	if refType, ok := ctx.pkg.Types[typeToken]; ok {
		keyProp, keyPropExists := refType.Properties["key"]
		valueProp, valuePropExists := refType.Properties["value"]
		// Check if the type has exactly two properties, "key" and "value", both of type "string"
		for k := range refType.Properties {
			if k != "key" && k != "value" {
				return false
			}
		}
		if keyPropExists && valuePropExists && keyProp.Type == "string" && valueProp.Type == "string" {
			return true
		}
	}

	return false
}
