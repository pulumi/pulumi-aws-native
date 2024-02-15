package schema

import (
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
