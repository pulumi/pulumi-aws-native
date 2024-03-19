package schema

import (
	"strings"

	jsschema "github.com/pulumi/jsschema"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/default_tags"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/naming"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func (ctx *cfSchemaContext) ApplyTagsTransformation(propName string, propertySpec *pschema.PropertySpec, spec *jsschema.Schema) default_tags.TagsStyle {
	tagsStyle := ctx.getTagsStyle(propName, &propertySpec.TypeSpec)
	switch tagsStyle {
	case default_tags.TagsStyleUntyped:
	case default_tags.TagsStyleStringMap:
		// Nothing to do
	case default_tags.TagsStyleKeyValueArray:
		// Swap referenced type to shared definition and remove custom type.
		propertySpec.TypeSpec.Items.Ref = "#/types/" + globalTagToken
	case default_tags.TagsStyleKeyValueArrayCreateOnly:
		// Swap referenced type to shared definition and remove custom type.
		propertySpec.TypeSpec.Items.Ref = "#/types/" + globalCreateOnlyTagToken
	case default_tags.TagsStyleKeyValueArrayWithExtraProperties:
		// Keep custom type
	case default_tags.TagsStyleKeyValueArrayWithAlternateType:
		// Keep custom type
	default: // Unknown
		ctx.reports.UnexpectedTagsShapes[ctx.resourceToken] = spec
	}
	return tagsStyle
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

func (ctx *cfSchemaContext) getTagsStyle(propName string, typeSpec *pschema.TypeSpec) default_tags.TagsStyle {
	if typeSpec == nil {
		return default_tags.TagsStyleUnknown
	}
	// Check for "Any" ref
	if typeSpec.Ref == "pulumi.json#/Any" {
		return default_tags.TagsStyleUntyped
	}

	if typeSpec.AdditionalProperties != nil && typeSpec.AdditionalProperties.Type == "string" {
		return default_tags.TagsStyleStringMap
	}

	if isKeyValueArray, style := ctx.tagStyleIsKeyValueArray(propName, typeSpec); isKeyValueArray {
		if style == default_tags.TagsStyleKeyValueArray && ctx.isPropCreateOnly(propName) {
			return default_tags.TagsStyleKeyValueArrayCreateOnly
		}
		if style != default_tags.TagsStyleUnknown {
			return style
		}
	}

	return default_tags.TagsStyleUnknown
}

func (ctx *cfSchemaContext) tagStyleIsKeyValueArray(propName string, typeSpec *pschema.TypeSpec) (bool, default_tags.TagsStyle) {
	if typeSpec == nil || typeSpec.Items == nil {
		return false, default_tags.TagsStyleUnknown
	}

	if typeSpec.Items.Ref != "" {
		typeToken, hasPrefix := strings.CutPrefix(typeSpec.Items.Ref, "#/types/")
		if !hasPrefix {
			return false, default_tags.TagsStyleUnknown
		}

		if refType, ok := ctx.pkg.Types[typeToken]; ok {
			if isKeyValue, hasAdditions := hasKeyValueStringPropertiesAndAdditions(&refType); isKeyValue {
				if hasAdditions {
					return true, default_tags.TagsStyleKeyValueArrayWithExtraProperties
				}
				return true, default_tags.TagsStyleKeyValueArray
			}
		}
	}

	if typeSpec.Items.OneOf != nil {
		for _, item := range typeSpec.Items.OneOf {
			typeToken := strings.TrimPrefix(item.Ref, "#/types/")
			if typeToken != "" {
				if refType, ok := ctx.pkg.Types[typeToken]; ok {
					if isKeyValue, _ := hasKeyValueStringPropertiesAndAdditions(&refType); isKeyValue {
						return true, default_tags.TagsStyleKeyValueArrayWithAlternateType
					}
				}
			}
		}
	}

	return false, default_tags.TagsStyleUnknown
}

func (ctx *cfSchemaContext) isPropCreateOnly(propName string) bool {
	createOnlyProps := readPropSdkNames(ctx.resourceSpec, "createOnlyProperties")
	return createOnlyProps.Has(naming.ToSdkName(propName))
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
