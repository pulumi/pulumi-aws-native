package schema

import (
	"strings"

	jsschema "github.com/lestrrat-go/jsschema"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

type TagsStyle string

const (
	TagsStyleUnknown       TagsStyle = ""
	TagsStyleUntyped       TagsStyle = "untyped"
	TagsStyleStringMap     TagsStyle = "stringMap"
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
	switch typeSpec.Type {
	case "object":
		if typeSpec.AdditionalProperties != nil && typeSpec.AdditionalProperties.Type == "string" {
			return TagsStyleStringMap
		} else {
			return TagsStyleUnknown
		}
	default:
		return TagsStyleUnknown
	}
}
