package schema

import (
	"strings"

	jsschema "github.com/pulumi/jsschema"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

type TagsStyle string

const (
	// TagsStyleUnknown indicates we can't identify the style of tags.
	TagsStyleUnknown TagsStyle = ""
	// TagsStyleUntyped is a style where the tags are represented as "Any" - without a schema.
	TagsStyleUntyped TagsStyle = "untyped"
	// TagsStyleStringMap is a style where the tags are represented as a map of strings.
	TagsStyleStringMap TagsStyle = "stringMap"
	// TagsStyleKeyValueArray is a style where the tags are represented as an array of key-value pairs.
	TagsStyleKeyValueArray TagsStyle = "keyValueArray"
	// TagsStyleKeyValueCreateOnlyArray is a style where the tags are represented as an array of key-value pairs, but
	// the tags are create-only.
	TagsStyleKeyValueCreateOnlyArray TagsStyle = "keyValueCreateOnlyArray"
	// TagsStyleKeyValueArrayWithExtraProperties is a style where the tags are represented as an array of key-value pairs
	// but can have extra properties.
	TagsStyleKeyValueArrayWithExtraProperties TagsStyle = "keyValueArrayWithExtraProperties"
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

func (ctx *context) GetTagsStyle(propName string, typeSpec *pschema.TypeSpec) TagsStyle {
	if typeSpec == nil {
		return TagsStyleUnknown
	}
	// Check for "Any" ref
	if typeSpec.Ref == "pulumi.json#/Any" {
		return TagsStyleUntyped
	}

	if typeSpec.AdditionalProperties != nil && typeSpec.AdditionalProperties.Type == "string" {
		return TagsStyleStringMap
	}

	if ctx.tagStyleIsKeyValueArray(propName, typeSpec, false /* includeExtraProperties */) {
		if ctx.isPropCreateOnly(propName) {
			return TagsStyleKeyValueCreateOnlyArray
		}
		return TagsStyleKeyValueArray
	}

	if ctx.tagStyleIsKeyValueArray(propName, typeSpec, true /* includeExtraProperties */) {
		return TagsStyleKeyValueArrayWithExtraProperties
	}

	return TagsStyleUnknown
}

func (ctx *context) tagStyleIsKeyValueArray(propName string, typeSpec *pschema.TypeSpec, allowExtraProperties bool) bool {
	if typeSpec == nil || typeSpec.Items == nil || typeSpec.Items.Ref == "" {
		return false
	}

	typeToken, hasPrefix := strings.CutPrefix(typeSpec.Items.Ref, "#/types/")
	if !hasPrefix {
		return false
	}

	if refType, ok := ctx.pkg.Types[typeToken]; ok {
		if hasKeyValueStringProperties(&refType, allowExtraProperties) {
			return true
		}
	}

	return false
}

func (ctx *context) isPropCreateOnly(propName string) bool {
	createOnlyProps := readPropSdkNames(ctx.resourceSpec, "createOnlyProperties")
	return createOnlyProps.Has(ToSdkName(propName))
}

func hasKeyValueStringProperties(typeSpec *pschema.ComplexTypeSpec, allowExtraProperties bool) bool {
	if typeSpec == nil || typeSpec.Properties == nil {
		return false
	}
	keyProp, keyPropExists := typeSpec.Properties["key"]
	valueProp, valuePropExists := typeSpec.Properties["value"]
	// Check if the type has exactly two properties, "key" and "value", both of type "string"
	hasKeyValueStrings := keyPropExists && valuePropExists && keyProp.Type == "string" && valueProp.Type == "string"
	if !hasKeyValueStrings {
		return false
	}
	return len(typeSpec.Properties) == 2 || allowExtraProperties
}
