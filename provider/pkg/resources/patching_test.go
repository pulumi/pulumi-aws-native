package resources

import (
	"testing"

	"github.com/mattbaird/jsonpatch"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/naming"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			CreateOnly: Map(args.spec.CreateOnly, naming.ToUpperCamel),
			WriteOnly:  Map(args.spec.WriteOnly, naming.ToUpperCamel),
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
			t.Run("no diff with must send props", func(t *testing.T) {
				expected := []jsonpatch.JsonPatchOperation{
					{Operation: "add", Path: "/Prop1", Value: "1"},
				}
				patch := run(t, args{
					oldInputs: resource.PropertyMap{
						"prop1": resource.NewStringProperty("1"),
						"prop2": resource.NewStringProperty("2"),
					},
					newInputs: resource.PropertyMap{
						"prop1": resource.NewStringProperty("1"),
						"prop2": resource.NewStringProperty("2"),
					},
					spec: metadata.CloudAPIResource{
						Inputs: map[string]schema.PropertySpec{
							"prop1": {TypeSpec: schema.TypeSpec{Type: "string"}},
							"prop2": {TypeSpec: schema.TypeSpec{Type: "string"}},
						},
						WriteOnly: []string{"prop1"},
					}})
				assert.Equal(t, expected, patch)
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
				assert.Equal(t, expected, patch)
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
				assert.Equal(t, expected, patch)
			})
		})
	}
}

func TestCalcPatchWithActualBaseline(t *testing.T) {
	t.Run("write-only falls back to old desired input", func(t *testing.T) {
		spec := metadata.CloudAPIResource{
			Inputs: map[string]schema.PropertySpec{
				"password": {TypeSpec: schema.TypeSpec{Type: "string"}},
				"name":     {TypeSpec: schema.TypeSpec{Type: "string"}},
			},
			WriteOnly: []string{"password"},
		}
		patch, err := CalcPatchWithActualBaseline(
			resource.PropertyMap{
				"password": resource.NewStringProperty("old-secret"),
				"name":     resource.NewStringProperty("old-name"),
			},
			resource.PropertyMap{
				"name": resource.NewStringProperty("old-name"),
			},
			resource.PropertyMap{
				"password": resource.NewStringProperty("old-secret"),
				"name":     resource.NewStringProperty("new-name"),
			},
			spec, nil, nil, "", nil)
		require.NoError(t, err)
		assert.Equal(t, []jsonpatch.JsonPatchOperation{
			{Operation: "replace", Path: "/Name", Value: "new-name"},
			{Operation: "add", Path: "/Password", Value: "old-secret"},
		}, patch)
	})

	t.Run("owned map drift is reconciled", func(t *testing.T) {
		spec := metadata.CloudAPIResource{
			Inputs: map[string]schema.PropertySpec{
				"tags": {
					TypeSpec: schema.TypeSpec{
						Type:                 "object",
						AdditionalProperties: &schema.TypeSpec{Type: "string"},
					},
				},
			},
			Outputs: map[string]schema.PropertySpec{
				"tags": {
					TypeSpec: schema.TypeSpec{
						Type:                 "object",
						AdditionalProperties: &schema.TypeSpec{Type: "string"},
					},
				},
			},
		}
		patch, err := CalcPatchWithActualBaseline(
			resource.PropertyMap{"tags": resource.NewObjectProperty(resource.PropertyMap{
				"owner": resource.NewStringProperty("team"),
			})},
			resource.PropertyMap{"tags": resource.NewObjectProperty(resource.PropertyMap{
				"owner": resource.NewStringProperty("team"),
				"extra": resource.NewStringProperty("external"),
			})},
			resource.PropertyMap{"tags": resource.NewObjectProperty(resource.PropertyMap{
				"owner": resource.NewStringProperty("team"),
			})},
			spec, nil, nil, "", nil)
		require.NoError(t, err)
		assert.Equal(t, []jsonpatch.JsonPatchOperation{
			{
				Operation: "replace",
				Path:      "/Tags",
				Value: map[string]interface{}{
					"owner": "team",
				},
			},
		}, patch)
	})

	t.Run("aws managed tag drift is suppressed", func(t *testing.T) {
		spec := metadata.CloudAPIResource{
			TagsProperty: "tags",
			Inputs: map[string]schema.PropertySpec{
				"tags": {
					TypeSpec: schema.TypeSpec{
						Type:                 "object",
						AdditionalProperties: &schema.TypeSpec{Type: "string"},
					},
				},
			},
			Outputs: map[string]schema.PropertySpec{
				"tags": {
					TypeSpec: schema.TypeSpec{
						Type:                 "object",
						AdditionalProperties: &schema.TypeSpec{Type: "string"},
					},
				},
			},
		}
		patch, err := CalcPatchWithActualBaseline(
			resource.PropertyMap{"tags": resource.NewObjectProperty(resource.PropertyMap{
				"owner": resource.NewStringProperty("team"),
			})},
			resource.PropertyMap{"tags": resource.NewObjectProperty(resource.PropertyMap{
				"owner":       resource.NewStringProperty("team"),
				"aws:managed": resource.NewStringProperty("external"),
			})},
			resource.PropertyMap{"tags": resource.NewObjectProperty(resource.PropertyMap{
				"owner": resource.NewStringProperty("team"),
			})},
			spec, nil, nil, "aws-native:test:Resource", NewTransformCache())
		require.NoError(t, err)
		assert.Empty(t, patch)
	})

}

// Map applies the given function to each element of the given slice and returns a new slice with the results.
func Map[T, U any](s []T, f func(T) U) []U {
	r := Prealloc[U](len(s))
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

// Preallocates a slice of type T of a given length. If the input length is 0 then the returned slice will be nil.
func Prealloc[T any](capacity int) []T {
	if capacity == 0 {
		return nil
	}
	return make([]T, 0, capacity)
}
