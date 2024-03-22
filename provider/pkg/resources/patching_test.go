package resources

import (
	"testing"

	"github.com/mattbaird/jsonpatch"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

func TestCalcPatch(t *testing.T) {
	type args struct {
		oldInputs resource.PropertyMap
		newInputs resource.PropertyMap
		spec      metadata.CloudAPIResource
		types     map[string]metadata.CloudAPIType
	}
	run := func(t *testing.T, args args) []jsonpatch.JsonPatchOperation {
		t.Helper()
		patch, err := CalcPatch(args.oldInputs, args.newInputs, args.spec, args.types)
		assert.NoError(t, err)
		return patch
	}
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
}
