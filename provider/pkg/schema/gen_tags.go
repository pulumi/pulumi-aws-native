package schema

import (
	"strings"

	jsschema "github.com/pulumi/jsschema"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

type TagsStyle string

const (
	tagStyleKeyValueArrayPrefix = "keyValueArray"
	// TagsStyleUnknown indicates we can't identify the style of tags.
	TagsStyleUnknown TagsStyle = ""
	// TagsStyleUntyped indicates the resource has no tags
	TagsStyleNone TagsStyle = "none"
	// TagsStyleUntyped is a style where the tags are represented as "Any" - without a schema.
	TagsStyleUntyped TagsStyle = "untyped"
	// TagsStyleStringMap is a style where the tags are represented as a map of strings.
	TagsStyleStringMap TagsStyle = "stringMap"
	// TagsStyleKeyValueArray is a style where the tags are represented as an array of key-value pairs.
	TagsStyleKeyValueArray TagsStyle = tagStyleKeyValueArrayPrefix
	// TagsStyleKeyValueArrayCreateOnly is a style where the tags are represented as an array of key-value pairs, but
	// the tags are create-only.
	TagsStyleKeyValueArrayCreateOnly TagsStyle = tagStyleKeyValueArrayPrefix + "CreateOnly"
	// TagsStyleKeyValueArrayWithExtraProperties is a style where the tags are represented as an array of key-value pairs
	// but can have extra properties.
	TagsStyleKeyValueArrayWithExtraProperties TagsStyle = tagStyleKeyValueArrayPrefix + "WithExtraProperties"
	// TagsStyleKeyValueArrayWithAlternateType is a style where the tags are represented as an array of key-value pairs
	// but the value can also be of a different type.
	TagsStyleKeyValueArrayWithAlternateType TagsStyle = tagStyleKeyValueArrayPrefix + "WithAlternateType"
)

func (ts TagsStyle) IsStringMap() bool {
	return ts == TagsStyleStringMap
}

func (ts TagsStyle) IsKeyValueArray() bool {
	return strings.HasPrefix(string(ts), tagStyleKeyValueArrayPrefix)
}

func (ctx *context) ApplyTagsTransformation(propertySpec *pschema.PropertySpec, spec *jsschema.Schema, tagsStyle TagsStyle) {
	switch tagsStyle {
	case TagsStyleUntyped:
	case TagsStyleStringMap:
		// Nothing to do
	case TagsStyleKeyValueArray:
		// Swap referenced type to shared definition and remove custom type.
		propertySpec.TypeSpec.Items.Ref = "#/types/" + globalTagToken
	case TagsStyleKeyValueArrayCreateOnly:
		// Swap referenced type to shared definition and remove custom type.
		propertySpec.TypeSpec.Items.Ref = "#/types/" + globalCreateOnlyTagToken
	case TagsStyleKeyValueArrayWithExtraProperties:
		// Keep custom type
	case TagsStyleKeyValueArrayWithAlternateType:
		// Keep custom type
	default: // Unknown
		ctx.reports.UnexpectedTagsShapes[ctx.resourceToken] = spec
	}
}

func GetTagsProperty(originalSpec *jsschema.Schema) (string, bool) {
	fromTaggingMetadata, ok := func() (string, bool) {
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
		// Some property paths have a leading #, some don't
		return strings.TrimPrefix(strings.TrimPrefix(tagPropertyString, "#"), "/properties/"), true
	}()
	if ok {
		return fromTaggingMetadata, true
	}
	if originalSpec.Properties != nil && originalSpec.Properties["Tags"] != nil {
		return "Tags", true
	}
	return "", false
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

	if isKeyValueArray, style := ctx.tagStyleIsKeyValueArray(propName, typeSpec); isKeyValueArray {
		if style == TagsStyleKeyValueArray && ctx.isPropCreateOnly(propName) {
			return TagsStyleKeyValueArrayCreateOnly
		}
		if style != TagsStyleUnknown {
			return style
		}
	}

	return TagsStyleUnknown
}

func (ctx *context) tagStyleIsKeyValueArray(propName string, typeSpec *pschema.TypeSpec) (bool, TagsStyle) {
	if typeSpec == nil || typeSpec.Items == nil {
		return false, TagsStyleUnknown
	}

	if typeSpec.Items.Ref != "" {
		typeToken, hasPrefix := strings.CutPrefix(typeSpec.Items.Ref, "#/types/")
		if !hasPrefix {
			return false, TagsStyleUnknown
		}

		if refType, ok := ctx.pkg.Types[typeToken]; ok {
			if isKeyValue, hasAdditions := hasKeyValueStringPropertiesAndAdditions(&refType); isKeyValue {
				if hasAdditions {
					return true, TagsStyleKeyValueArrayWithExtraProperties
				}
				return true, TagsStyleKeyValueArray
			}
		}
	}

	if typeSpec.Items.OneOf != nil {
		for _, item := range typeSpec.Items.OneOf {
			typeToken := strings.TrimPrefix(item.Ref, "#/types/")
			if typeToken != "" {
				if refType, ok := ctx.pkg.Types[typeToken]; ok {
					if isKeyValue, _ := hasKeyValueStringPropertiesAndAdditions(&refType); isKeyValue {
						return true, TagsStyleKeyValueArrayWithAlternateType
					}
				}
			}
		}
	}

	return false, TagsStyleUnknown
}

func (ctx *context) isPropCreateOnly(propName string) bool {
	createOnlyProps := readPropSdkNames(ctx.resourceSpec, "createOnlyProperties")
	return createOnlyProps.Has(ToSdkName(propName))
}

func hasKeyValueStringPropertiesAndAdditions(typeSpec *pschema.ComplexTypeSpec) (hasKeyValueStringProperties bool, hasExtraProperties bool) {
	if typeSpec == nil || typeSpec.Properties == nil {
		return false, false
	}
	keyProp, keyPropExists := typeSpec.Properties["key"]
	valueProp, valuePropExists := typeSpec.Properties["value"]
	// Check if the type has exactly two properties, "key" and "value", both of type "string"
	hasKeyValueStrings := keyPropExists && valuePropExists && keyProp.Type == "string" && valueProp.Type == "string"
	if !hasKeyValueStrings {
		return false, false
	}
	return true, len(typeSpec.Properties) != 2
}
