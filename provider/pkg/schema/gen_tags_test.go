package schema

import (
	"testing"

	jsschema "github.com/lestrrat-go/jsschema"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"
)

func TestGetTagsStyle(t *testing.T) {
	t.Run("defaults to unknown when typeSpec is nil", func(t *testing.T) {
		ctx := &context{}
		assert.Equal(t, TagsStyleUnknown, ctx.GetTagsStyle("Tags", nil))
	})
	t.Run("untyped style", func(t *testing.T) {
		ctx := &context{}
		typeSpec := &schema.TypeSpec{
			Ref: "pulumi.json#/Any",
		}
		assert.Equal(t, TagsStyleUntyped, ctx.GetTagsStyle("Tags", typeSpec))
	})
	t.Run("string map style", func(t *testing.T) {
		ctx := &context{}
		typeSpec := &schema.TypeSpec{
			AdditionalProperties: &schema.TypeSpec{
				Type: "string",
			},
		}
		assert.Equal(t, TagsStyleStringMap, ctx.GetTagsStyle("Tags", typeSpec))
	})
	t.Run("key value array style", func(t *testing.T) {
		ctx := &context{
			pkg: &schema.PackageSpec{
				Types: map[string]schema.ComplexTypeSpec{
					"pulumi:types:input:common:ComponentResourceOptions:TagsEntry": {
						ObjectTypeSpec: schema.ObjectTypeSpec{
							Properties: map[string]schema.PropertySpec{
								"key":   {TypeSpec: schema.TypeSpec{Type: "string"}},
								"value": {TypeSpec: schema.TypeSpec{Type: "string"}},
							},
						},
					},
				},
			},
		}
		typeSpec := &schema.TypeSpec{
			Items: &schema.TypeSpec{
				Ref: "#/types/pulumi:types:input:common:ComponentResourceOptions:TagsEntry",
			},
		}
		assert.Equal(t, TagsStyleKeyValueArray, ctx.GetTagsStyle("Tags", typeSpec))
	})
	t.Run("not key value array style with extra field", func(t *testing.T) {
		ctx := &context{
			pkg: &schema.PackageSpec{
				Types: map[string]schema.ComplexTypeSpec{
					"pulumi:types:input:common:ComponentResourceOptions:TagsEntry": {
						ObjectTypeSpec: schema.ObjectTypeSpec{
							Properties: map[string]schema.PropertySpec{
								"key":   {TypeSpec: schema.TypeSpec{Type: "string"}},
								"value": {TypeSpec: schema.TypeSpec{Type: "string"}},
								"extra": {TypeSpec: schema.TypeSpec{Type: "string"}},
							},
						},
					},
				},
			},
		}
		typeSpec := &schema.TypeSpec{
			Items: &schema.TypeSpec{
				Ref: "#/types/pulumi:types:input:common:ComponentResourceOptions:TagsEntry",
			},
		}
		assert.NotEqual(t, TagsStyleKeyValueArray, ctx.GetTagsStyle("Tags", typeSpec))
	})
	t.Run("not key value array style if causes replacement", func(t *testing.T) {
		ctx := &context{
			pkg: &schema.PackageSpec{
				Types: map[string]schema.ComplexTypeSpec{
					"pulumi:types:input:common:ComponentResourceOptions:TagsEntry": {
						ObjectTypeSpec: schema.ObjectTypeSpec{
							Properties: map[string]schema.PropertySpec{
								"key":   {TypeSpec: schema.TypeSpec{Type: "string"}},
								"value": {TypeSpec: schema.TypeSpec{Type: "string"}},
							},
						},
					},
				},
			},
			resourceSpec: &jsschema.Schema{
				Extras: map[string]interface{}{
					"createOnlyProperties": []interface{}{"/properties/Tags"},
				},
			},
		}
		typeSpec := &schema.TypeSpec{
			Items: &schema.TypeSpec{
				Ref: "#/types/pulumi:types:input:common:ComponentResourceOptions:TagsEntry",
			},
		}
		assert.NotEqual(t, TagsStyleKeyValueArray, ctx.GetTagsStyle("Tags", typeSpec))
	})
}
