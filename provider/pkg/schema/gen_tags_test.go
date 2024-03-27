package schema

import (
	"encoding/json"
	"os"
	"path"
	"testing"

	jsschema "github.com/pulumi/jsschema"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/default_tags"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTagsStyle(t *testing.T) {
	t.Run("defaults to unknown when typeSpec is nil", func(t *testing.T) {
		ctx := &cfSchemaContext{}
		assert.Equal(t, default_tags.TagsStyleUnknown, ctx.getTagsStyle("Tags", nil))
	})
	t.Run("untyped style", func(t *testing.T) {
		ctx := &cfSchemaContext{}
		typeSpec := &schema.TypeSpec{
			Ref: "pulumi.json#/Any",
		}
		assert.Equal(t, default_tags.TagsStyleUntyped, ctx.getTagsStyle("Tags", typeSpec))
	})
	t.Run("string map style", func(t *testing.T) {
		ctx := &cfSchemaContext{}
		typeSpec := &schema.TypeSpec{
			AdditionalProperties: &schema.TypeSpec{
				Type: "string",
			},
		}
		assert.Equal(t, default_tags.TagsStyleStringMap, ctx.getTagsStyle("Tags", typeSpec))
	})
	t.Run("key value array style", func(t *testing.T) {
		ctx := &cfSchemaContext{
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
		assert.Equal(t, default_tags.TagsStyleKeyValueArray, ctx.getTagsStyle("Tags", typeSpec))
	})
	t.Run("not key value array style with extra field", func(t *testing.T) {
		ctx := &cfSchemaContext{
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
		assert.Equal(t, default_tags.TagsStyleKeyValueArrayWithExtraProperties, ctx.getTagsStyle("Tags", typeSpec))
	})
	t.Run("key value create-only array style if causes replacement", func(t *testing.T) {
		ctx := &cfSchemaContext{
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
		assert.Equal(t, default_tags.TagsStyleKeyValueArrayCreateOnly, ctx.getTagsStyle("Tags", typeSpec))
	})
	t.Run("key value array with alternate type style", func(t *testing.T) {
		ctx := &cfSchemaContext{
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
					"pulumi:types:input:common:ComponentResourceOptions:AltTagsEntry": {
						ObjectTypeSpec: schema.ObjectTypeSpec{
							Properties: map[string]schema.PropertySpec{
								"tagKey":   {TypeSpec: schema.TypeSpec{Type: "string"}},
								"tagValue": {TypeSpec: schema.TypeSpec{Type: "string"}},
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
				OneOf: []schema.TypeSpec{
					{
						Ref: "#/types/pulumi:types:input:common:ComponentResourceOptions:AltTagsEntry",
					},
					{
						Ref: "#/types/pulumi:types:input:common:ComponentResourceOptions:TagsEntry",
					},
				},
			},
		}
		assert.Equal(t, default_tags.TagsStyleKeyValueArrayWithAlternateType, ctx.getTagsStyle("Tags", typeSpec))
	})
}

func TestNoUnexpectedTagsShapes(t *testing.T) {
	path := path.Join("..", "..", "..", "reports", "unexpectedTagsShapes.json")
	bytes, err := os.ReadFile(path)
	require.NoError(t, err)
	var unexpectedTagsShapes map[string]interface{}
	err = json.Unmarshal(bytes, &unexpectedTagsShapes)
	require.NoError(t, err)
	assert.Empty(t, unexpectedTagsShapes, "reports/unexpectedTagsShapes.json should be empty, which means that we need to update getTagsStyle in gen_tags.go to cover new tags variations.")
}
