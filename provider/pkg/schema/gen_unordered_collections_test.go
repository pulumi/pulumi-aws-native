// Copyright 2026, Pulumi Corporation.

package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"

	jsschema "github.com/pulumi/jsschema"
)

const (
	childrenProperty = "Children"
	valuesProperty   = "Values"
)

func TestReadUnorderedCollections(t *testing.T) {
	tests := map[string]struct {
		schema   *jsschema.Schema
		expected []string
	}{
		"nil schema": {},
		"top-level primitive array": {
			schema: resourceSchema(map[string]*jsschema.Schema{
				"Names": unorderedArraySchema(stringSchema()),
			}),
			expected: []string{"names"},
		},
		"definition reference and SDK casing": {
			schema: &jsschema.Schema{
				Properties: map[string]*jsschema.Schema{
					"SomeARNs": {Reference: "#/definitions/Arns"},
				},
				Definitions: map[string]*jsschema.Schema{
					"Arns": unorderedArraySchema(stringSchema()),
				},
			},
			expected: []string{"someArns"},
		},
		"nested array": {
			schema: resourceSchema(map[string]*jsschema.Schema{
				"Outer": orderedArraySchema(&jsschema.Schema{
					Type: jsschema.PrimitiveTypes{jsschema.ObjectType},
					Properties: map[string]*jsschema.Schema{
						"InnerValues": unorderedArraySchema(stringSchema()),
					},
				}),
			}),
			expected: []string{"outer/*/innerValues"},
		},
		"same definition at multiple paths": {
			schema: &jsschema.Schema{
				Properties: map[string]*jsschema.Schema{
					"First":  {Reference: "#/definitions/Names"},
					"Second": {Reference: "#/definitions/Names"},
				},
				Definitions: map[string]*jsschema.Schema{
					"Names": unorderedArraySchema(stringSchema()),
				},
			},
			expected: []string{"first", "second"},
		},
		"true absent malformed and non-array annotations": {
			schema: resourceSchema(map[string]*jsschema.Schema{
				"True":      arraySchema(true, stringSchema()),
				"Absent":    orderedArraySchema(stringSchema()),
				"Malformed": arraySchema("false", stringSchema()),
				"NotArray": {
					Type:   jsschema.PrimitiveTypes{jsschema.StringType},
					Extras: map[string]interface{}{insertionOrderProperty: false},
				},
			}),
		},
		"ordered parent and unordered child": {
			schema: resourceSchema(map[string]*jsschema.Schema{
				"Parent": orderedArraySchema(&jsschema.Schema{
					Properties: map[string]*jsschema.Schema{
						childrenProperty: unorderedArraySchema(stringSchema()),
					},
				}),
			}),
			expected: []string{"parent/*/children"},
		},
		"unordered parent and ordered child": {
			schema: resourceSchema(map[string]*jsschema.Schema{
				"Parent": unorderedArraySchema(&jsschema.Schema{
					Properties: map[string]*jsschema.Schema{
						childrenProperty: orderedArraySchema(stringSchema()),
					},
				}),
			}),
			expected: []string{"parent"},
		},
		"duplicate composition paths": {
			schema: resourceSchema(map[string]*jsschema.Schema{
				valuesProperty: {
					AnyOf: jsschema.SchemaList{
						unorderedArraySchema(stringSchema()),
						unorderedArraySchema(stringSchema()),
					},
				},
			}),
			expected: []string{"values"},
		},
		"allOf composition is conservatively omitted": {
			schema: resourceSchema(map[string]*jsschema.Schema{
				valuesProperty: {
					Extras: map[string]interface{}{insertionOrderProperty: false},
					AllOf: jsschema.SchemaList{
						orderedArraySchema(stringSchema()),
					},
				},
			}),
		},
		"mixed anyOf and oneOf composition is conservatively omitted": {
			schema: resourceSchema(map[string]*jsschema.Schema{
				valuesProperty: {
					AnyOf: jsschema.SchemaList{unorderedArraySchema(stringSchema())},
					OneOf: jsschema.SchemaList{unorderedArraySchema(stringSchema())},
				},
			}),
		},
		"array alternative without annotation is conservative": {
			schema: resourceSchema(map[string]*jsschema.Schema{
				valuesProperty: {
					OneOf: jsschema.SchemaList{
						unorderedArraySchema(stringSchema()),
						orderedArraySchema(stringSchema()),
					},
				},
			}),
		},
		"reference cycle": {
			schema: &jsschema.Schema{
				Properties: map[string]*jsschema.Schema{
					"Nodes": {Reference: "#/definitions/Node"},
				},
				Definitions: map[string]*jsschema.Schema{
					"Node": {
						Type: jsschema.PrimitiveTypes{jsschema.ObjectType},
						Properties: map[string]*jsschema.Schema{
							childrenProperty: unorderedArraySchema(&jsschema.Schema{
								Reference: "#/definitions/Node",
							}),
						},
					},
				},
			},
			expected: []string{"nodes/children"},
		},
		"multiple item schemas": {
			schema: resourceSchema(map[string]*jsschema.Schema{
				"Outer": {
					Type: jsschema.PrimitiveTypes{jsschema.ArrayType},
					Items: &jsschema.ItemSpec{TupleMode: true, Schemas: jsschema.SchemaList{
						{Properties: map[string]*jsschema.Schema{
							"Inner": unorderedArraySchema(stringSchema()),
						}},
						{Properties: map[string]*jsschema.Schema{
							"Inner": unorderedArraySchema(stringSchema()),
						}},
					}},
				},
			}),
			expected: []string{"outer/*/inner"},
		},
		"escaped local definition": {
			schema: &jsschema.Schema{
				Properties: map[string]*jsschema.Schema{
					valuesProperty: {Reference: "#/definitions/Foo~1Bar"},
				},
				Definitions: map[string]*jsschema.Schema{
					"Foo/Bar": unorderedArraySchema(stringSchema()),
				},
			},
			expected: []string{"values"},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.expected, readUnorderedCollections(tt.schema))
		})
	}
}

func resourceSchema(properties map[string]*jsschema.Schema) *jsschema.Schema {
	return &jsschema.Schema{Properties: properties}
}

func stringSchema() *jsschema.Schema {
	return &jsschema.Schema{Type: jsschema.PrimitiveTypes{jsschema.StringType}}
}

func unorderedArraySchema(items ...*jsschema.Schema) *jsschema.Schema {
	return arraySchema(false, items...)
}

func orderedArraySchema(items ...*jsschema.Schema) *jsschema.Schema {
	return &jsschema.Schema{
		Type:  jsschema.PrimitiveTypes{jsschema.ArrayType},
		Items: &jsschema.ItemSpec{Schemas: items},
	}
}

func arraySchema(insertionOrder interface{}, items ...*jsschema.Schema) *jsschema.Schema {
	return &jsschema.Schema{
		Type:   jsschema.PrimitiveTypes{jsschema.ArrayType},
		Items:  &jsschema.ItemSpec{Schemas: items},
		Extras: map[string]interface{}{insertionOrderProperty: insertionOrder},
	}
}
