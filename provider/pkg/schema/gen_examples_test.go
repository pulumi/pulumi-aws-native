// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	jsschema "github.com/lestrrat-go/jsschema"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

type PropertyTypeSpecTestCase struct {
	json     string
	expected pschema.TypeSpec
}

func TestPropertyTypeSpec(t *testing.T) {
	cases := []PropertyTypeSpecTestCase{
		{
			json: `{
				"type": "string",
				"oneOf": [
					{"format": "date-time"},
					{"format": "timestamp"}
				]
			}`,
			expected: pschema.TypeSpec{Type: "string"},
		},
		{
			json: `{
				"type": "string",
				"anyOf": [
					{"format": "date-time"},
					{"format": "timestamp"}
				]
			}`,
			expected: pschema.TypeSpec{Type: "string"},
		},
		{
			json: `{
				"oneOf": [
					{"type": "string", "format": "date-time"},
					{"type": "string", "format": "timestamp"}
				]
			}`,
			expected: pschema.TypeSpec{Type: "string"},
		},
		{
			json: `{
				"anyOf": [
					{"type": "string", "format": "date-time"},
					{"type": "string", "format": "timestamp"}
				]
			}`,
			expected: pschema.TypeSpec{Type: "string"},
		},
		{
			json: `{
				"oneOf": [
					{"type": "string"},
					{"type": "number"}
				]
			}`,
			expected: pschema.TypeSpec{
				OneOf: []pschema.TypeSpec{{Type: "string"}, {Type: "number"}},
			},
		},
		{
			json: `{
				"anyOf": [
					{"type": "string"},
					{"type": "number"}
				]
			}`,
			expected: pschema.TypeSpec{
				OneOf: []pschema.TypeSpec{{Type: "string"}, {Type: "number"}},
			},
		},
		{
			json: `{
				"anyOf": [
					{"type": "object", "properties": { "A": { "type": "number" } } },
					{"type": "object", "properties": { "A": { "type": "string" } } }
				]
			}`,
			expected: pschema.TypeSpec{
				OneOf: []pschema.TypeSpec{
					{Ref: "#/types/aws-native::Foo0Properties"},
					{Ref: "#/types/aws-native::Foo1Properties"},
				},
			},
		},
		{
			json: `{
				"anyOf": [
					{"type": "object", "properties": { "A": { "type": "number" } } },
					{"type": "object" },
					{"type": "string" }
				]
			}`,
			expected: pschema.TypeSpec{
				OneOf: []pschema.TypeSpec{
					{Ref: "#/types/aws-native::FooProperties"},
					{Ref: "pulumi.json#/Any"},
					{Type: "string"},
				},
			},
		},
		{
			json: `{
				"type": "object",
				"oneOf": [
					 {"properties": { "A": { "type": "number" } } },
					 {"properties": { "B": { "type": "number" } } }
				 ]
			 }`,
			expected: pschema.TypeSpec{
				OneOf: []pschema.TypeSpec{
					{Ref: "#/types/aws-native::Foo0Properties"},
					{Ref: "#/types/aws-native::Foo1Properties"},
				},
			},
		},
		{
			json: `{
				"$ref": "#/definitions/Obj"
			 }`,
			expected: pschema.TypeSpec{Ref: "#/types/aws-native::Obj"},
		},
		{
			json: `{
				"$ref": "#/definitions/ObjLike1"
			 }`,
			expected: pschema.TypeSpec{Ref: "#/types/aws-native::ObjLike1"},
		},
		{
			json: `{
				"$ref": "#/definitions/ObjLike2"
			 }`,
			expected: pschema.TypeSpec{Ref: "#/types/aws-native::ObjLike2"},
		},
		{
			json: `{
				"$ref": "#/definitions/OneOf"
			 }`,
			expected: pschema.TypeSpec{
				OneOf: []pschema.TypeSpec{
					{Type: "number"},
					{Type: "string"},
				},
			},
		},
	}

	ctx := context{
		pkg: &pschema.PackageSpec{
			Types: map[string]pschema.ComplexTypeSpec{},
		},
		visitedTypes: codegen.NewStringSet(),
		metadata: &CloudAPIMetadata{
			Types: map[string]CloudAPIType{},
		},
		resourceSpec: &jsschema.Schema{
			Definitions: map[string]*jsschema.Schema{
				"Obj": &jsschema.Schema{Type: jsschema.PrimitiveTypes{jsschema.ObjectType}},
				"OneOf": &jsschema.Schema{
					OneOf: jsschema.SchemaList{
						&jsschema.Schema{Type: jsschema.PrimitiveTypes{jsschema.NumberType}},
						&jsschema.Schema{Type: jsschema.PrimitiveTypes{jsschema.StringType}},
					},
				},
				"ObjLike1": &jsschema.Schema{
					Properties: map[string]*jsschema.Schema{
						"foo": jsschema.New(),
					},
				},
				"ObjLike2": &jsschema.Schema{
					PatternProperties: map[*regexp.Regexp]*jsschema.Schema{
						regexp.MustCompile(".+"): jsschema.New(),
					},
				},
			},
		},
	}

	for _, tt := range cases {
		schema := jsschema.New()
		err := schema.UnmarshalJSON([]byte(tt.json))
		assert.Nil(t, err)

		actual, err := ctx.propertyTypeSpec("Foo", schema)
		if assert.Nil(t, err) {
			assert.Equal(t, tt.expected, *actual)
		}
	}
}

func TestEnumType(t *testing.T) {
	cases := []struct {
		name           string
		schema         *jsschema.Schema
		expectedType   string
		expectedValues map[string]string
	}{
		{
			name: "SomeHTTPEnum",
			schema: &jsschema.Schema{
				Type: jsschema.PrimitiveTypes{jsschema.StringType},
				Enum: []interface{}{"UseHTTP1Thing", "use_HTTP2_thing"},
			},
			expectedType: "aws-native::SomeHttpEnum",
			expectedValues: map[string]string{
				"UseHttp1Thing": "UseHTTP1Thing",
				"UseHttp2Thing": "use_HTTP2_thing",
			},
		},
	}

	for _, tt := range cases {

		ctx := context{
			pkg: &pschema.PackageSpec{
				Types: map[string]pschema.ComplexTypeSpec{},
			},
			metadata: &CloudAPIMetadata{
				Types: map[string]CloudAPIType{},
			},
		}
		out, err := (&ctx).genEnumType(tt.name, tt.schema)
		assert.NoError(t, err)
		assert.Equal(t, "#/types/"+tt.expectedType, out.Ref)
		if assert.Contains(t, ctx.pkg.Types, tt.expectedType) {
			v, _ := ctx.pkg.Types[tt.expectedType]
			actualValues := map[string]string{}
			for _, v := range v.Enum {
				actualValues[v.Name] = v.Value.(string)
			}
			assert.Equal(t, tt.expectedValues, actualValues)
		}
	}
}

func TestModuleName(t *testing.T) {
	cases := map[string]string{
		"aws::SomeEC2Thing::SomeDHCPOptions": "SomeEc2Thing",
	}

	for input, expected := range cases {
		assert.Equal(t, expected, moduleName(input))
	}
}

func TestTypeName(t *testing.T) {
	cases := map[string]string{
		"aws::FOO::SomeDHCPOptions": "SomeDhcpOptions",
	}

	for input, expected := range cases {
		assert.Equal(t, expected, typeName(input))
	}
}
