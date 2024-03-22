package autonaming

import (
	"testing"

	jsschema "github.com/pulumi/jsschema"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"
)

func TestCreateAutoNamingSpec(t *testing.T) {
	type args struct {
		resourceTypeName     string
		inputProperties      map[string]schema.PropertySpec
		jsonSchemaProperties map[string]*jsschema.Schema
		semanticsSpec        metadata.SemanticsSpec
	}
	create := func(args args) *metadata.AutoNamingSpec {
		return CreateAutoNamingSpec(args.inputProperties, args.resourceTypeName, args.jsonSchemaProperties, args.semanticsSpec)
	}
	t.Run("literal name", func(t *testing.T) {
		expected := &metadata.AutoNamingSpec{
			SdkName: "name",
		}
		spec := create(args{
			inputProperties: map[string]schema.PropertySpec{
				"name": {TypeSpec: schema.TypeSpec{Type: "string"}},
			},
			jsonSchemaProperties: map[string]*jsschema.Schema{
				"Name": {},
			},
		})
		assert.Equal(t, expected, spec)
	})

	t.Run("skips if not input", func(t *testing.T) {
		spec := create(args{
			jsonSchemaProperties: map[string]*jsschema.Schema{
				"Name": {},
			},
		})
		assert.Nil(t, spec)
	})

	t.Run("literal name preferred", func(t *testing.T) {
		expected := &metadata.AutoNamingSpec{
			SdkName: "name",
		}
		spec := create(args{
			inputProperties: map[string]schema.PropertySpec{
				"name":         {TypeSpec: schema.TypeSpec{Type: "string"}},
				"resourceName": {TypeSpec: schema.TypeSpec{Type: "string"}},
			},
			jsonSchemaProperties: map[string]*jsschema.Schema{
				"Name":         {},
				"ResourceName": {},
			},
		})
		assert.Equal(t, expected, spec)
	})

	t.Run("resource plus name", func(t *testing.T) {
		expected := &metadata.AutoNamingSpec{
			SdkName: "resourceName",
		}
		spec := create(args{
			resourceTypeName: "Resource",
			inputProperties: map[string]schema.PropertySpec{
				"resourceName": {TypeSpec: schema.TypeSpec{Type: "string"}},
			},
			jsonSchemaProperties: map[string]*jsschema.Schema{
				"ResourceName": {},
			},
		})
		assert.Equal(t, expected, spec)
	})

	t.Run("resource plus name preferred over partial", func(t *testing.T) {
		expected := &metadata.AutoNamingSpec{
			SdkName: "fooBarName",
		}
		spec := create(args{
			resourceTypeName: "FooBar",
			inputProperties: map[string]schema.PropertySpec{
				"fooBarName": {TypeSpec: schema.TypeSpec{Type: "string"}},
				"barName":    {TypeSpec: schema.TypeSpec{Type: "string"}},
			},
			jsonSchemaProperties: map[string]*jsschema.Schema{
				"FooBarName": {},
				"BarName":    {},
			},
		})
		assert.Equal(t, expected, spec)
	})

	t.Run("partial overlap with type name", func(t *testing.T) {
		expected := &metadata.AutoNamingSpec{
			SdkName: "barName",
		}
		spec := create(args{
			resourceTypeName: "FooBar",
			inputProperties: map[string]schema.PropertySpec{
				"barName": {TypeSpec: schema.TypeSpec{Type: "string"}},
			},
			jsonSchemaProperties: map[string]*jsschema.Schema{
				"BarName": {},
			},
		})
		assert.Equal(t, expected, spec)
	})

	t.Run("skips when multiple names overlap", func(t *testing.T) {
		spec := create(args{
			resourceTypeName: "FooBarBaz",
			inputProperties: map[string]schema.PropertySpec{
				"bazName":    {TypeSpec: schema.TypeSpec{Type: "string"}},
				"barBazName": {TypeSpec: schema.TypeSpec{Type: "string"}},
			},
			jsonSchemaProperties: map[string]*jsschema.Schema{
				"BazName":    {},
				"BarBazName": {},
			},
		})
		assert.Nil(t, spec)
	})

	t.Run("handles oddly cased acryonyms", func(t *testing.T) {
		expected := &metadata.AutoNamingSpec{
			SdkName: "tlsInspectionConfigurationName",
		}
		spec := create(args{
			resourceTypeName: "TlsInspectionConfiguration",
			inputProperties: map[string]schema.PropertySpec{
				"tlsInspectionConfigurationName": {TypeSpec: schema.TypeSpec{Type: "string"}},
			},
			jsonSchemaProperties: map[string]*jsschema.Schema{
				"TLSInspectionConfigurationName": {},
			},
		})
		assert.Equal(t, expected, spec)
	})
}
