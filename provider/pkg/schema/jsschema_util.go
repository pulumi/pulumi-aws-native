// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	jsschema "github.com/lestrrat-go/jsschema"
)

// FlattenJSSchema recursively flattens a schema containing AnyOf or OneOf into a list of schemas
// Note: this only performs a shallow copy, don't use in situations where schema fields may be mutated.
// TODO(https://github.com/pulumi/pulumi-aws-native/issues/992): this should flatten AllOf as well.
func FlattenJSSchema(sch *jsschema.Schema) jsschema.SchemaList {
	if len(sch.AnyOf) == 0 && len(sch.OneOf) == 0 {
		return jsschema.SchemaList{sch}
	}

	schemas := make(jsschema.SchemaList, 0, len(sch.AnyOf)+len(sch.OneOf))

	// build a flattened list of subschemas
	for _, s := range sch.AnyOf {
		schemas = append(schemas, FlattenJSSchema(s)...)
	}
	for _, s := range sch.OneOf {
		schemas = append(schemas, FlattenJSSchema(s)...)
	}

	// merge each with the base schema and clear the AnyOf/OneOf fields
	for i, s := range schemas {
		merged := jsschema.New()
		MergeJSSchema(merged, sch)
		merged.AnyOf = nil
		merged.OneOf = nil
		MergeJSSchema(merged, s)
		schemas[i] = merged
	}
	return schemas

}

// MergeJSSchema overlays the contents of src into dest to merge the values of the fields
// Note: This only does a shallow copy of src fields, don't use where schemas maybe mutated
func MergeJSSchema(dest *jsschema.Schema, src *jsschema.Schema) {
	if src.ID != "" {
		dest.ID = src.ID
	}

	if src.Title != "" {
		dest.Title = src.Title
	}

	if src.Description != "" {
		dest.Description = src.Description
	}

	if src.Default != nil {
		dest.Default = src.Default
	}

	if len(src.Type) != 0 {
		dest.Type = append(src.Type, dest.Type...)
	}

	if src.SchemaRef != "" {
		dest.SchemaRef = src.SchemaRef
	}

	if len(src.Definitions) != 0 {
		dest.Definitions = mergeMaps(dest.Definitions, src.Definitions)
	}

	if src.Reference != "" {
		dest.Reference = src.Reference
	}

	if src.Format != "" {
		dest.Format = src.Format
	}

	if src.MultipleOf.Initialized {
		dest.MultipleOf = src.MultipleOf
	}

	if src.Minimum.Initialized {
		dest.Minimum = src.Minimum
	}

	if src.Maximum.Initialized {
		dest.Maximum = src.Maximum
	}

	if src.ExclusiveMinimum.Initialized {
		dest.ExclusiveMinimum = src.ExclusiveMinimum
	}

	if src.ExclusiveMaximum.Initialized {
		dest.ExclusiveMaximum = src.ExclusiveMaximum
	}

	if src.MaxLength.Initialized {
		dest.MaxLength = src.MaxLength
	}

	if src.MinLength.Initialized {
		dest.MinLength = src.MinLength
	}

	if src.Pattern != nil {
		dest.Pattern = src.Pattern
	}

	if src.AdditionalItems != nil {
		dest.AdditionalItems = src.AdditionalItems
	}

	if src.Items != nil {
		dest.Items = src.Items
	}

	if src.MinItems.Initialized {
		dest.MinItems = src.MinItems
	}

	if src.MaxItems.Initialized {
		dest.MaxItems = src.MaxItems
	}

	if src.UniqueItems.Initialized {
		dest.UniqueItems = src.UniqueItems
	}

	if src.MaxProperties.Initialized {
		dest.MaxProperties = src.MaxProperties
	}

	if src.MinProperties.Initialized {
		dest.MinProperties = src.MinProperties
	}

	if len(src.Required) != 0 {
		dest.Required = append(src.Required, dest.Required...)
	}

	if len(src.Dependencies.Names) != 0 {
		dest.Dependencies = jsschema.DependencyMap{
			Names:   mergeMaps(dest.Dependencies.Names, src.Dependencies.Names),
			Schemas: mergeMaps(dest.Dependencies.Schemas, src.Dependencies.Schemas),
		}
	}

	if len(src.Properties) != 0 {
		dest.Properties = mergeMaps(dest.Properties, src.Properties)
	}

	if src.AdditionalProperties != nil && src.AdditionalProperties.Schema != nil {
		if dest.AdditionalProperties == nil || dest.AdditionalProperties.Schema == nil {
			dest.AdditionalProperties = src.AdditionalProperties
		} else {
			merged := jsschema.AdditionalProperties{Schema: jsschema.New()}
			MergeJSSchema(merged.Schema, dest.AdditionalProperties.Schema)
			MergeJSSchema(merged.Schema, src.AdditionalProperties.Schema)
			dest.AdditionalProperties = &merged
		}
	}

	if len(src.PatternProperties) != 0 {
		dest.PatternProperties = mergeMaps(dest.PatternProperties, src.PatternProperties)
	}

	if len(src.Enum) != 0 {
		dest.Enum = append(src.Enum, dest.Enum...)
	}

	if len(src.AllOf) != 0 {
		dest.AllOf = append(src.AllOf, dest.AllOf...)
	}

	if len(src.AnyOf) != 0 {
		dest.AnyOf = append(src.AnyOf, dest.AnyOf...)
	}

	if len(src.OneOf) != 0 {
		dest.OneOf = append(src.OneOf, dest.OneOf...)
	}

	if src.Not != nil {
		dest.Not = src.Not
	}

	if len(src.Extras) != 0 {
		dest.Extras = mergeMaps(dest.Extras, src.Extras)
	}
}

func mergeMaps[K comparable, V any](m1, m2 map[K]V) map[K]V {
	merged := map[K]V{}
	for k, v := range m1 {
		merged[k] = v
	}
	for k, v := range m2 {
		merged[k] = v
	}
	return merged
}
