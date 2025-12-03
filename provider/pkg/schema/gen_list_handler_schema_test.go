package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	jsschema "github.com/pulumi/jsschema"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
)

func TestGatherListHandlerSchema(t *testing.T) {
	ctx := cfSchemaContext{
		resourceSpec: &jsschema.Schema{
			Extras: map[string]interface{}{
				"handlers": map[string]interface{}{
					"list": map[string]interface{}{
						"handlerSchema": map[string]interface{}{
							"properties": map[string]interface{}{
								"FunctionName": map[string]interface{}{
									"description": "The name of the function.",
									"type":        "string",
									"pattern":     "ignored",
								},
								// This property should be ignored because it lacks description and type.
								"EmptyProp": map[string]interface{}{
									"pattern": "ignored",
								},
							},
							"required": []interface{}{"FunctionName"},
						},
					},
				},
			},
		},
	}

	listSchema := ctx.gatherListHandlerSchema()
	require.NotNil(t, listSchema)
	assert.Equal(t, []string{"FunctionName"}, listSchema.Required)
	assert.Equal(t, map[string]metadata.ListHandlerProperty{
		"FunctionName": {
			Description: "The name of the function.",
			Type:        "string",
		},
	}, listSchema.Properties)
}

func TestGatherListHandlerSchemaMissing(t *testing.T) {
	ctx := cfSchemaContext{
		resourceSpec: &jsschema.Schema{
			Extras: map[string]interface{}{},
		},
	}

	assert.Nil(t, ctx.gatherListHandlerSchema())
}
