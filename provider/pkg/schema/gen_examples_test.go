// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	jsschema "github.com/pulumi/jsschema"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

type PropertyTypeSpecTestCase struct {
	json          string
	expected      pschema.TypeSpec
	expectedTypes map[string]pschema.ComplexTypeSpec
}

func TestPropertyTypeSpec(t *testing.T) {
	test := func(tt PropertyTypeSpecTestCase) func(t *testing.T) {
		return func(t *testing.T) {
			t.Helper()
			ctx := cfSchemaContext{
				reports: NewReports(),
				pkg: &pschema.PackageSpec{
					Types: map[string]pschema.ComplexTypeSpec{},
				},
				visitedTypes: codegen.NewStringSet(),
				metadata: &metadata.CloudAPIMetadata{
					Types: map[string]metadata.CloudAPIType{},
				},
				resourceSpec: &jsschema.Schema{
					Definitions: map[string]*jsschema.Schema{
						"Obj": {Type: jsschema.PrimitiveTypes{jsschema.ObjectType}},
						"OneOf": {
							OneOf: jsschema.SchemaList{
								&jsschema.Schema{Type: jsschema.PrimitiveTypes{jsschema.NumberType}},
								&jsschema.Schema{Type: jsschema.PrimitiveTypes{jsschema.StringType}},
							},
						},
						"ObjLike1": {
							Properties: map[string]*jsschema.Schema{
								"foo": jsschema.New(),
							},
						},
						"ObjLike2": {
							PatternProperties: map[string]*jsschema.Schema{
								".+": jsschema.New(),
							},
						},
					},
				},
			}

			schema := jsschema.New()
			err := schema.UnmarshalJSON([]byte(tt.json))
			require.NoError(t, err)

			actual, err := ctx.propertyTypeSpec("Foo", schema)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, *actual)
			if len(tt.expectedTypes) > 0 {
				assert.Equal(t, tt.expectedTypes, ctx.pkg.Types)
			}
		}
	}

	t.Run("oneOf-strings-with-base-type", test(PropertyTypeSpecTestCase{
		json: `{
				"type": "string",
				"oneOf": [
					{"format": "date-time"},
					{"format": "timestamp"}
				]
			}`,
		expected: pschema.TypeSpec{Type: "string"},
	}))
	t.Run("anyOf-strings-with-base-type", test(PropertyTypeSpecTestCase{
		json: `{
				"type": "string",
				"anyOf": [
					{"format": "date-time"},
					{"format": "timestamp"}
				]
			}`,
		expected: pschema.TypeSpec{Type: "string"},
	}))
	t.Run("oneOf-strings-no-base-type", test(PropertyTypeSpecTestCase{
		json: `{
				"oneOf": [
					{"type": "string", "format": "date-time"},
					{"type": "string", "format": "timestamp"}
				]
			}`,
		expected: pschema.TypeSpec{Type: "string"},
	}))
	t.Run("anyOf-strings-no-base-type", test(PropertyTypeSpecTestCase{
		json: `{
				"anyOf": [
					{"type": "string", "format": "date-time"},
					{"type": "string", "format": "timestamp"}
				]
			}`,
		expected: pschema.TypeSpec{Type: "string"},
	}))
	t.Run("oneOf-primatives", test(PropertyTypeSpecTestCase{
		json: `{
				"oneOf": [
					{"type": "string"},
					{"type": "number"}
				]
			}`,
		expected: pschema.TypeSpec{
			OneOf: []pschema.TypeSpec{{Type: "string"}, {Type: "number"}},
		},
	}))
	t.Run("anyOf-primatives", test(PropertyTypeSpecTestCase{
		json: `{
				"anyOf": [
					{"type": "string"},
					{"type": "number"}
				]
			}`,
		expected: pschema.TypeSpec{
			OneOf: []pschema.TypeSpec{{Type: "string"}, {Type: "number"}},
		},
	}))
	t.Run("oneOf-objects", test(PropertyTypeSpecTestCase{
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
	}))
	t.Run("anyOf-mixed-types-no-base-type", test(PropertyTypeSpecTestCase{
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
	}))
	t.Run("oneOf-object-properties", test(PropertyTypeSpecTestCase{
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
	}))
	t.Run("ref-to-untyped-object", test(PropertyTypeSpecTestCase{
		json: `{
				"$ref": "#/definitions/Obj"
			 }`,
		expected: pschema.TypeSpec{Ref: "#/types/aws-native::Obj"},
	}))
	t.Run("ref-to-object-with-named-properties", test(PropertyTypeSpecTestCase{
		json: `{
				"$ref": "#/definitions/ObjLike1"
			 }`,
		expected: pschema.TypeSpec{Ref: "#/types/aws-native::ObjLike1"},
	}))
	t.Run("ref-to-object-with-patternProperties", test(PropertyTypeSpecTestCase{
		json: `{
				"$ref": "#/definitions/ObjLike2"
			 }`,
		expected: pschema.TypeSpec{
			Type: "object",
			AdditionalProperties: &pschema.TypeSpec{
				Ref: "pulumi.json#/Any",
			},
		},
	}))
	t.Run("ref-to-oneOf", test(PropertyTypeSpecTestCase{
		json: `{
				"$ref": "#/definitions/OneOf"
			 }`,
		expected: pschema.TypeSpec{
			OneOf: []pschema.TypeSpec{
				{Type: "number"},
				{Type: "string"},
			},
		},
	}))
	t.Run("string-and-number", test(PropertyTypeSpecTestCase{
		json: `{
				"type": ["string","number"]
			}`,
		expected: pschema.TypeSpec{
			OneOf: []pschema.TypeSpec{{Type: "string"}, {Type: "number"}},
		},
	}))
	t.Run("inferred-array", test(PropertyTypeSpecTestCase{
		json: `{
				"items": { "type": "string" }
			}`,
		expected: pschema.TypeSpec{
			Type: "array",
			Items: &pschema.TypeSpec{
				Type: "string",
			},
		},
	}))
	t.Run("collapse-oneOf-any", test(PropertyTypeSpecTestCase{
		json: `{
				"type": ["string","object"]
			}`,
		expected: pschema.TypeSpec{
			Ref: "pulumi.json#/Any",
		},
	}))
	t.Run("string-enum", test(PropertyTypeSpecTestCase{
		json: `{
				"type": "string",
				"enum": ["UseHTTP1Thing", "use_HTTP2_thing"]
			}`,
		expected: pschema.TypeSpec{
			Ref: "#/types/aws-native::Foo",
		},
		expectedTypes: map[string]pschema.ComplexTypeSpec{
			"aws-native::Foo": {
				ObjectTypeSpec: pschema.ObjectTypeSpec{
					Type: "string",
				},
				Enum: []pschema.EnumValueSpec{
					{Name: "UseHttp1Thing", Value: "UseHTTP1Thing"},
					{Name: "UseHttp2Thing", Value: "use_HTTP2_thing"},
				},
			},
		},
	}))
}

func TestEnumType(t *testing.T) {
	cases := []struct {
		name           string
		schema         *jsschema.Schema
		expectedType   string
		expectedValues map[string]string
		cfTypeName     string
		resourceName   string
		mod            string
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
		{
			name:         "FileSystemProtectionReplicationOverwriteProtection",
			cfTypeName:   "AWS::EFS::FileSystem",
			resourceName: "FileSystemProtection",
			mod:          "efs",
			schema: &jsschema.Schema{
				Type:        jsschema.PrimitiveTypes{jsschema.StringType},
				Enum:        []interface{}{"DISABLED", "ENABLED"}, // CloudFormation schema values only
				Description: "The status including REPLICATING state",
			},
			expectedType: "aws-native:efs:FileSystemProtectionReplicationOverwriteProtection",
			expectedValues: map[string]string{
				"Disabled":    "DISABLED",
				"Enabled":     "ENABLED",
				"Replicating": "REPLICATING", // Expected after fix
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctx := cfSchemaContext{
				pkg: &pschema.PackageSpec{
					Types: map[string]pschema.ComplexTypeSpec{},
				},
				metadata: &metadata.CloudAPIMetadata{
					Types: map[string]metadata.CloudAPIType{},
				},
				cfTypeName:   tt.cfTypeName,
				resourceName: tt.resourceName,
				mod:          tt.mod,
			}
			out, err := (&ctx).genEnumType(tt.name, tt.schema)
			assert.NoError(t, err)
			assert.Equal(t, "#/types/"+tt.expectedType, out.Ref)
			if assert.Contains(t, ctx.pkg.Types, tt.expectedType) {
				v := ctx.pkg.Types[tt.expectedType]
				actualValues := map[string]string{}
				for _, v := range v.Enum {
					actualValues[v.Name] = v.Value.(string)
				}
				assert.Equal(t, tt.expectedValues, actualValues)
			}
		})
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
				"foo": {},
				"obj": {TypeSpec: pschema.TypeSpec{Ref: "#/types/obj"}},
				"arr": {TypeSpec: pschema.TypeSpec{
					Items: &pschema.TypeSpec{Ref: "#/types/obj"},
				}},
			},
		},
	}
	baseTypes := map[string]pschema.ComplexTypeSpec{
		"obj": {
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Properties: map[string]pschema.PropertySpec{
					"bar": {},
					"baz": {TypeSpec: pschema.TypeSpec{Ref: "#/types/obj2"}},
				},
			},
		},
		"obj2": {
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Properties: map[string]pschema.PropertySpec{
					"zz": {},
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
		ctx := cfSchemaContext{
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
