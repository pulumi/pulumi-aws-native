// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	"testing"
	"testing/quick"

	jsschema "github.com/pulumi/jsschema"
	"github.com/stretchr/testify/assert"
)

func TestFlattenJSSchema(t *testing.T) {

	anyOfemptyOuter := func(as []AlmostJSSchema) bool {
		for i, s := range as {
			as[i] = withoutAnyOrAll(s)
		}
		ss := asJSSchemaArray(as)
		outer := newishJSSchema()
		outer.OneOf = nil
		outer.AnyOf = ss

		flat := FlattenJSSchema(outer)

		result := true
		for i, s := range ss {
			if len(s.Extras) == 0 {
				s.Extras = nil // Flatten drops empty Extras map to nil
			}
			result = result && assert.Equal(t, s, flat[i])
		}
		return result
	}

	anyOfemptyInner := func(as AlmostJSSchema) bool {
		clean := withoutAnyOrAll(as)
		outer := asJSSchema(clean)
		outer.AnyOf = jsschema.SchemaList{newishJSSchema()}

		flat := FlattenJSSchema(outer)

		expect := jsschema.SchemaList{asJSSchema(clean)}
		if len(expect[0].Extras) == 0 {
			expect[0].Extras = nil // Flatten drops empty Extras map to nil
		}
		return assert.Equal(t, expect, flat)
	}

	oneOfemptyOuter := func(as []AlmostJSSchema) bool {
		for i, s := range as {
			as[i] = withoutAnyOrAll(s)
		}
		ss := asJSSchemaArray(as)
		outer := newishJSSchema()
		outer.OneOf = ss
		outer.AnyOf = nil

		flat := FlattenJSSchema(outer)

		result := true
		for i, s := range ss {
			if len(s.Extras) == 0 {
				s.Extras = nil // Flatten drops empty Extras map to nil
			}
			result = result && assert.Equal(t, s, flat[i])
		}
		return result
	}

	oneOfemptyInner := func(as AlmostJSSchema) bool {
		clean := withoutAnyOrAll(as)
		outer := asJSSchema(clean)
		outer.OneOf = jsschema.SchemaList{newishJSSchema()}

		flat := FlattenJSSchema(outer)

		expect := jsschema.SchemaList{asJSSchema(clean)}
		if len(expect[0].Extras) == 0 {
			expect[0].Extras = nil // Flatten drops empty Extras map to nil
		}
		return assert.Equal(t, expect, flat)
	}

	t.Run("anyOf/emptyOuter", func(t *testing.T) {
		if err := quick.Check(anyOfemptyOuter, nil); err != nil {
			t.Error(err)
		}
	})

	t.Run("anyOf/emptyInner", func(t *testing.T) {
		if err := quick.Check(anyOfemptyInner, nil); err != nil {
			t.Error(err)
		}
	})

	t.Run("oneOf/emptyOuter", func(t *testing.T) {
		if err := quick.Check(oneOfemptyOuter, nil); err != nil {
			t.Error(err)
		}
	})

	t.Run("oneOf/emptyInner", func(t *testing.T) {
		if err := quick.Check(oneOfemptyInner, nil); err != nil {
			t.Error(err)
		}
	})

}

func TestMergeJSSchema(t *testing.T) {

	intoEmpty := func(as AlmostJSSchema) bool {
		s1 := asJSSchema(as)
		s := newishJSSchema()

		MergeJSSchema(s, s1)

		return assert.Equal(t, s1, s)
	}

	addEmpty := func(as AlmostJSSchema) bool {
		s1 := asJSSchema(as)
		s2 := asJSSchema(as)

		MergeJSSchema(s2, jsschema.New())

		return assert.Equal(t, s1, s2)
	}

	isAssociative := func(as1 AlmostJSSchema, as2 AlmostJSSchema, as3 AlmostJSSchema) bool {
		s1 := asJSSchema(as1)
		s2 := asJSSchema(as2)
		s3 := asJSSchema(as3)

		sa := asJSSchema(as1)
		sb := asJSSchema(as2)
		sc := asJSSchema(as3)

		MergeJSSchema(s1, s2)
		MergeJSSchema(s1, s3)

		MergeJSSchema(sb, sc)
		MergeJSSchema(sa, sb)

		return assert.Equal(t, s1, sa)
	}

	t.Run("intoEmpty", func(t *testing.T) {
		if err := quick.Check(intoEmpty, nil); err != nil {
			t.Error(err)
		}
	})

	t.Run("addEmpty", func(t *testing.T) {
		if err := quick.Check(addEmpty, nil); err != nil {
			t.Error(err)
		}
	})

	t.Run("isAssociative", func(t *testing.T) {
		if err := quick.Check(isAssociative, nil); err != nil {
			t.Error(err)
		}
	})

}

// jsschema.New doesn't initialize Extras, but schemas read from JSON always do
func newishJSSchema() *jsschema.Schema {
	s := jsschema.New()
	s.Extras = map[string]interface{}{}
	return s
}

func withoutAnyOrAll(s AlmostJSSchema) AlmostJSSchema {
	clean := s
	clean.AnyOf = nil
	clean.AllOf = nil
	return clean
}

func asJSSchema(almost AlmostJSSchema) *jsschema.Schema {
	schema := jsschema.New()

	schema.ID = almost.ID
	schema.Title = almost.Title
	schema.Description = almost.Description
	schema.Default = almost.Default
	schema.Type = almost.Type
	schema.SchemaRef = almost.SchemaRef
	schema.Definitions = asJSSchemaMap(almost.Definitions)
	schema.Reference = almost.Reference
	schema.Format = almost.Format

	schema.MultipleOf = almost.MultipleOf
	schema.Minimum = almost.Minimum
	schema.Maximum = almost.Maximum
	schema.ExclusiveMinimum = almost.ExclusiveMinimum
	schema.ExclusiveMaximum = almost.ExclusiveMaximum

	schema.MaxLength = almost.MaxLength
	schema.MinLength = almost.MinLength
	schema.Pattern = almost.Pattern

	schema.AdditionalItems = almost.AdditionalItems
	schema.Items = almost.Items
	schema.MinItems = almost.MinItems
	schema.MaxItems = almost.MaxItems
	schema.UniqueItems = almost.UniqueItems

	schema.MaxProperties = almost.MaxProperties
	schema.MinProperties = almost.MinProperties
	schema.Required = almost.Required
	schema.Dependencies = almost.Dependencies
	schema.Properties = asJSSchemaMap(almost.Properties)
	schema.AdditionalProperties = almost.AdditionalProperties
	schema.PatternProperties = asJSSchemaMap(almost.PatternProperties)

	schema.Enum = make([]interface{}, len(almost.Enum))

	for i, v := range almost.Enum {
		schema.Enum[i] = v
	}

	schema.AllOf = asJSSchemaArray(almost.AllOf)
	schema.AnyOf = asJSSchemaArray(almost.AnyOf)
	schema.OneOf = asJSSchemaArray(almost.OneOf)

	if almost.Not != nil {
		schema.Not = asJSSchema(*almost.Not)
	}

	schema.Extras = make(map[string]interface{}, len(almost.Extras))

	for k, v := range almost.Extras {
		schema.Extras[k] = v
	}

	//run it through json to standardize a bit
	json, err := schema.MarshalJSON()
	if err != nil {
		panic(err)
	}

	schema2 := jsschema.New()
	err = schema2.UnmarshalJSON(json)
	if err != nil {
		panic(err)
	}

	return schema2
}

func asJSSchemaMap[K comparable](m map[K]AlmostJSSchema) map[K]*jsschema.Schema {
	result := make(map[K]*jsschema.Schema, len(m))
	for k, v := range m {
		result[k] = asJSSchema(v)
	}
	return result
}

func asJSSchemaArray(arr []AlmostJSSchema) []*jsschema.Schema {
	result := make([]*jsschema.Schema, len(arr))
	for i, v := range arr {
		result[i] = asJSSchema(v)
	}
	return result
}

type AlmostJSSchema struct {
	ID          string
	Title       string
	Description string
	Default     struct{ Foo string }
	Type        jsschema.PrimitiveTypes
	SchemaRef   string
	Definitions map[string]AlmostJSSchema
	Reference   string
	Format      jsschema.Format

	// NumericValidations
	MultipleOf       jsschema.Number
	Minimum          jsschema.Number
	Maximum          jsschema.Number
	ExclusiveMinimum jsschema.Bool
	ExclusiveMaximum jsschema.Bool

	// StringValidation
	MaxLength jsschema.Integer
	MinLength jsschema.Integer
	Pattern   string

	// ArrayValidations
	AdditionalItems *jsschema.AdditionalItems
	Items           *jsschema.ItemSpec
	MinItems        jsschema.Integer
	MaxItems        jsschema.Integer
	UniqueItems     jsschema.Bool

	// ObjectValidations
	MaxProperties        jsschema.Integer
	MinProperties        jsschema.Integer
	Required             []string
	Dependencies         jsschema.DependencyMap
	Properties           map[string]AlmostJSSchema
	AdditionalProperties *jsschema.AdditionalProperties
	PatternProperties    map[string]AlmostJSSchema

	Enum   []struct{ Foo string }
	AllOf  []AlmostJSSchema
	AnyOf  []AlmostJSSchema
	OneOf  []AlmostJSSchema
	Not    *AlmostJSSchema
	Extras map[string]struct{ Foo string }
}
