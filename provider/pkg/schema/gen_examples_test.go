// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	//	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	jsschema "github.com/lestrrat-go/jsschema"
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
	}

	for _, tt := range cases {
		// t.Log("JSON: ", tt.json) // uncomment to debug
		schema := jsschema.New()
		err := schema.UnmarshalJSON([]byte(tt.json))
		assert.Nil(t, err)
		ctx := context{}
		actual, _ := ctx.propertyTypeSpec("Foo", schema)
		assert.Equal(t, tt.expected, *actual)
	}
}
