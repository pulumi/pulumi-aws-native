// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	jsschema "github.com/pulumi/jsschema"
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
					PatternProperties: map[string]*jsschema.Schema{
						".+": jsschema.New(),
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

type MarkCreateOnlyPropertiesTestCase struct {
	createOnlyProperties codegen.StringSet
	resourceSpec         pschema.ResourceSpec
	types                map[string]pschema.ComplexTypeSpec

	expectedErr      string
	expectedResource pschema.ResourceSpec
	expectedTypes    map[string]pschema.ComplexTypeSpec
}

func TestMarkCreateOnlyProperties(t *testing.T) {

	baseResourceSpec := pschema.ResourceSpec{
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Properties: map[string]pschema.PropertySpec{
				"foo": pschema.PropertySpec{},
				"obj": pschema.PropertySpec{TypeSpec: pschema.TypeSpec{Ref: "#/types/obj"}},
				"arr": pschema.PropertySpec{TypeSpec: pschema.TypeSpec{
					Items: &pschema.TypeSpec{Ref: "#/types/obj"},
				}},
			},
		},
	}
	baseTypes := map[string]pschema.ComplexTypeSpec{
		"obj": pschema.ComplexTypeSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Properties: map[string]pschema.PropertySpec{
					"bar": pschema.PropertySpec{},
					"baz": pschema.PropertySpec{TypeSpec: pschema.TypeSpec{Ref: "#/types/obj2"}},
				},
			},
		},
		"obj2": pschema.ComplexTypeSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Properties: map[string]pschema.PropertySpec{
					"zz": pschema.PropertySpec{},
				},
			},
		},
	}

	baseResourceSpecJson, _ := json.Marshal(baseResourceSpec)
	baseTypesJson, _ := json.Marshal(baseTypes)

	newBaseResourceSpec := func() pschema.ResourceSpec {
		r := pschema.ResourceSpec{}
		_ = json.Unmarshal(baseResourceSpecJson, &r)
		return r
	}
	newBaseTypes := func() map[string]pschema.ComplexTypeSpec {
		r := map[string]pschema.ComplexTypeSpec{}
		_ = json.Unmarshal(baseTypesJson, &r)
		return r
	}

	resourceWithReplaceOnChangesFoo := newBaseResourceSpec()
	resourceWithReplaceOnChangesFoo.Properties["foo"] = pschema.PropertySpec{ReplaceOnChanges: true}

	resourceWithReplaceOnChangesObj := newBaseResourceSpec()
	modifiedObj := resourceWithReplaceOnChangesObj.Properties["obj"]
	modifiedObj.ReplaceOnChanges = true
	resourceWithReplaceOnChangesObj.Properties["obj"] = modifiedObj

	resourceWithReplaceOnChangesArr := newBaseResourceSpec()
	modifiedArr := resourceWithReplaceOnChangesObj.Properties["arr"]
	modifiedArr.ReplaceOnChanges = true
	resourceWithReplaceOnChangesArr.Properties["arr"] = modifiedArr

	typesWithReplaceOnChangesBar := newBaseTypes()
	modifedObjT := typesWithReplaceOnChangesBar["obj"]
	modifedObjT.ObjectTypeSpec.Properties["bar"] = pschema.PropertySpec{ReplaceOnChanges: true}
	typesWithReplaceOnChangesBar["obj"] = modifedObjT

	typesWithReplaceOnChangesZZ := newBaseTypes()
	modifedObj2T := typesWithReplaceOnChangesZZ["obj2"]
	modifedObj2T.ObjectTypeSpec.Properties["zz"] = pschema.PropertySpec{ReplaceOnChanges: true}
	typesWithReplaceOnChangesZZ["obj2"] = modifedObj2T

	cases := []MarkCreateOnlyPropertiesTestCase{
		{
			createOnlyProperties: codegen.NewStringSet("a"),
			resourceSpec:         newBaseResourceSpec(),
			types:                newBaseTypes(),
			expectedErr:          "Could not mark createOnlyProperty a in  as replaceOnChanges: property not found on Resource",
		},
		{
			createOnlyProperties: codegen.NewStringSet("obj/quxx"),
			resourceSpec:         newBaseResourceSpec(),
			types:                newBaseTypes(),
			expectedErr:          "Could not mark createOnlyProperty obj/quxx in  as replaceOnChanges: Type #/types/obj does not have property named 'quxx'",
		},
		{
			createOnlyProperties: codegen.NewStringSet("foo/quxx"),
			resourceSpec:         newBaseResourceSpec(),
			types:                newBaseTypes(),
			expectedErr:          "Could not mark createOnlyProperty foo/quxx in  as replaceOnChanges: Property is not a Ref or Array[Ref], can't traverse",
		},
		{
			createOnlyProperties: codegen.NewStringSet("foo"),
			resourceSpec:         newBaseResourceSpec(),
			types:                newBaseTypes(),
			expectedResource:     resourceWithReplaceOnChangesFoo,
			expectedTypes:        newBaseTypes(),
		},
		{
			createOnlyProperties: codegen.NewStringSet("obj"),
			resourceSpec:         newBaseResourceSpec(),
			types:                newBaseTypes(),
			expectedResource:     resourceWithReplaceOnChangesObj,
			expectedTypes:        newBaseTypes(),
		},
		{
			createOnlyProperties: codegen.NewStringSet("obj/bar"),
			resourceSpec:         newBaseResourceSpec(),
			types:                newBaseTypes(),
			expectedResource:     newBaseResourceSpec(),
			expectedTypes:        typesWithReplaceOnChangesBar,
		},
		{
			createOnlyProperties: codegen.NewStringSet("obj/baz/zz"),
			resourceSpec:         newBaseResourceSpec(),
			types:                newBaseTypes(),
			expectedResource:     newBaseResourceSpec(),
			expectedTypes:        typesWithReplaceOnChangesZZ,
		},
		{
			createOnlyProperties: codegen.NewStringSet("arr/bar"),
			resourceSpec:         newBaseResourceSpec(),
			types:                newBaseTypes(),
			expectedResource:     newBaseResourceSpec(),
			expectedTypes:        typesWithReplaceOnChangesBar,
		},
		{
			createOnlyProperties: codegen.NewStringSet("arr/*/bar"),
			resourceSpec:         newBaseResourceSpec(),
			types:                newBaseTypes(),
			expectedResource:     newBaseResourceSpec(),
			expectedTypes:        typesWithReplaceOnChangesBar,
		},
		{
			createOnlyProperties: codegen.NewStringSet("arr"),
			resourceSpec:         newBaseResourceSpec(),
			types:                newBaseTypes(),
			expectedResource:     resourceWithReplaceOnChangesArr,
			expectedTypes:        newBaseTypes(),
		},
		{
			createOnlyProperties: codegen.NewStringSet("arr/*"),
			resourceSpec:         newBaseResourceSpec(),
			types:                newBaseTypes(),
			expectedResource:     resourceWithReplaceOnChangesArr,
			expectedTypes:        newBaseTypes(),
		},
	}

	for _, tt := range cases {
		ctx := context{
			pkg: &pschema.PackageSpec{
				Types: tt.types,
			},
		}

		err := ctx.markCreateOnlyProperties(tt.createOnlyProperties, &tt.resourceSpec)

		testCaseInfo := fmt.Sprintf("Testpaths: %v", tt.createOnlyProperties)

		if len(tt.expectedErr) > 0 {
			assert.ErrorContains(t, err, tt.expectedErr, testCaseInfo)
		} else if assert.NoError(t, err, testCaseInfo) {
			assert.Equal(t, tt.expectedResource, tt.resourceSpec, testCaseInfo)
			assert.Equal(t, tt.expectedTypes, tt.types, testCaseInfo)
		}
	}
}
