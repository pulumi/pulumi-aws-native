package resources

import (
	"testing"

	"github.com/mattbaird/jsonpatch"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/naming"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/slice"
	"github.com/stretchr/testify/assert"
)

func TestCalcPatch(t *testing.T) {
	type args struct {
		oldInputs resource.PropertyMap
		newInputs resource.PropertyMap
		spec      metadata.CloudAPIResource
		types     map[string]metadata.CloudAPIType
	}
	type implementation func(t *testing.T, args args) []jsonpatch.JsonPatchOperation
	typed := func(t *testing.T, args args) []jsonpatch.JsonPatchOperation {
		t.Helper()
		patch, err := CalcPatch(args.oldInputs, args.newInputs, args.spec, args.types)
		assert.NoError(t, err)
		return patch
	}
	untyped := func(t *testing.T, args args) []jsonpatch.JsonPatchOperation {
		t.Helper()
		// Convert all keys to upper camel case to match the behavior of the SdkToCfn function.
		keysToUpperCamel := func(s string) (string, bool) {
			return naming.ToUpperCamel(s), true
		}
		patch, err := CalculateUntypedPatch(ExtensionResourceInputs{
			Properties: args.oldInputs.MapRepl(keysToUpperCamel, nil),
		}, ExtensionResourceInputs{
			Properties: args.newInputs.MapRepl(keysToUpperCamel, nil),
			CreateOnly: slice.Map(args.spec.CreateOnly, naming.ToUpperCamel),
			WriteOnly:  slice.Map(args.spec.WriteOnly, naming.ToUpperCamel),
		})
		assert.NoError(t, err)
		return patch
	}

	for name, run := range map[string]implementation{
		"Typed":   typed,
		"Untyped": untyped,
	} {
		t.Run(name, func(t *testing.T) {
			t.Run("no diff", func(t *testing.T) {
				patch := run(t, args{
					oldInputs: resource.PropertyMap{
						"prop1": resource.NewStringProperty("value1"),
					},
					newInputs: resource.PropertyMap{
						"prop1": resource.NewStringProperty("value1"),
					},
					spec: metadata.CloudAPIResource{
						Inputs: map[string]schema.PropertySpec{
							"prop1": {TypeSpec: schema.TypeSpec{Type: "string"}},
						},
					}})
				assert.Empty(t, patch)
			})
			t.Run("always sends write-only properties", func(t *testing.T) {
				expected := []jsonpatch.JsonPatchOperation{
					{Operation: "add", Path: "/Prop1", Value: "1"},
					{Operation: "replace", Path: "/Prop2", Value: "2b"},
				}
				patch := run(t, args{
					oldInputs: resource.PropertyMap{
						"prop1": resource.NewStringProperty("1"),
						"prop2": resource.NewStringProperty("2a"),
					},
					newInputs: resource.PropertyMap{
						"prop1": resource.NewStringProperty("1"),
						"prop2": resource.NewStringProperty("2b"),
					},
					spec: metadata.CloudAPIResource{
						Inputs: map[string]schema.PropertySpec{
							"prop1": {TypeSpec: schema.TypeSpec{Type: "string"}},
							"prop2": {TypeSpec: schema.TypeSpec{Type: "string"}},
						},
						WriteOnly: []string{"prop1"},
					}})
				assert.ElementsMatch(t, expected, patch)
			})
			t.Run("don't send write-only, create-only properties", func(t *testing.T) {
				expected := []jsonpatch.JsonPatchOperation{
					{Operation: "replace", Path: "/Prop2", Value: "2b"},
				}
				patch := run(t, args{
					oldInputs: resource.PropertyMap{
						"prop1": resource.NewStringProperty("1"),
						"prop2": resource.NewStringProperty("2a"),
					},
					newInputs: resource.PropertyMap{
						"prop1": resource.NewStringProperty("1"),
						"prop2": resource.NewStringProperty("2b"),
					},
					spec: metadata.CloudAPIResource{
						Inputs: map[string]schema.PropertySpec{
							"prop1": {TypeSpec: schema.TypeSpec{Type: "string"}},
							"prop2": {TypeSpec: schema.TypeSpec{Type: "string"}},
						},
						WriteOnly:  []string{"prop1"},
						CreateOnly: []string{"prop1"},
					}})
				assert.ElementsMatch(t, expected, patch)
			})
		})
	}
}
