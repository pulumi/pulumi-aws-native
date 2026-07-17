//nolint:goconst // Repeated literals keep schema fixtures readable.
package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	jsschema "github.com/pulumi/jsschema"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
)

func TestGatherListHandlerSchema(t *testing.T) {
	ctx := cfSchemaContext{
		resourceSpec: &jsschema.Schema{
			Properties: map[string]*jsschema.Schema{
				"FunctionName": {
					Type:        jsschema.PrimitiveTypes{jsschema.StringType},
					Description: "The name of the function from base schema.",
				},
			},
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
							},
							"required": []interface{}{"FunctionName"},
						},
					},
				},
			},
		},
	}

	listSchema, err := ctx.gatherListHandlerSchema()
	require.NoError(t, err)
	require.NotNil(t, listSchema)
	assert.Equal(t, []string{"FunctionName"}, listSchema.Required)
	assert.Equal(t, map[string]metadata.ListHandlerProperty{
		"FunctionName": {
			Description: "The name of the function.",
			Type:        "string",
		},
	}, listSchema.Properties)
}

func TestGatherListHandlerSchemaWithRefFallback(t *testing.T) {
	ctx := cfSchemaContext{
		resourceSpec: &jsschema.Schema{
			Properties: map[string]*jsschema.Schema{
				"ApiId": {
					Type: jsschema.PrimitiveTypes{jsschema.StringType},
				},
			},
			Extras: map[string]interface{}{
				"handlers": map[string]interface{}{
					"list": map[string]interface{}{
						"handlerSchema": map[string]interface{}{
							"properties": map[string]interface{}{
								"ApiId": map[string]interface{}{
									"$ref": "resource-schema.json#/properties/ApiId",
								},
							},
							"required": []interface{}{"ApiId"},
						},
					},
				},
			},
		},
	}

	listSchema, err := ctx.gatherListHandlerSchema()
	require.NoError(t, err)
	require.NotNil(t, listSchema)
	assert.Equal(t, []string{"ApiId"}, listSchema.Required)
	assert.Equal(t, map[string]metadata.ListHandlerProperty{
		"ApiId": {
			Type: "string",
		},
	}, listSchema.Properties)
}

func TestGatherListHandlerSchemaWithRefThroughDefinition(t *testing.T) {
	ctx := cfSchemaContext{
		resourceSpec: &jsschema.Schema{
			Properties: map[string]*jsschema.Schema{
				"DomainName": {
					Reference: "#/definitions/DomainName",
				},
			},
			Definitions: map[string]*jsschema.Schema{
				"DomainName": {
					Type:        jsschema.PrimitiveTypes{jsschema.StringType},
					Description: "The unique name of the domain.",
				},
			},
			Extras: map[string]interface{}{
				"handlers": map[string]interface{}{
					"list": map[string]interface{}{
						"handlerSchema": map[string]interface{}{
							"properties": map[string]interface{}{
								"DomainName": map[string]interface{}{
									"$ref": "resource-schema.json#/properties/DomainName",
								},
							},
							"required": []interface{}{"DomainName"},
						},
					},
				},
			},
		},
	}

	listSchema, err := ctx.gatherListHandlerSchema()
	require.NoError(t, err)
	require.NotNil(t, listSchema)
	assert.Equal(t, []string{"DomainName"}, listSchema.Required)
	assert.Equal(t, map[string]metadata.ListHandlerProperty{
		"DomainName": {
			Description: "The unique name of the domain.",
			Type:        "string",
		},
	}, listSchema.Properties)
}

func TestGatherListHandlerSchemaRequiredPropertyFromResourceSchema(t *testing.T) {
	ctx := cfSchemaContext{
		resourceSpec: &jsschema.Schema{
			Properties: map[string]*jsschema.Schema{
				"MultiplexId": {
					Type:        jsschema.PrimitiveTypes{jsschema.StringType},
					Description: "The ID of the multiplex that the program belongs to.",
				},
			},
			Extras: map[string]interface{}{
				"handlers": map[string]interface{}{
					"list": map[string]interface{}{
						"handlerSchema": map[string]interface{}{
							"properties": map[string]interface{}{
								"Arn": map[string]interface{}{
									"$ref": "resource-schema.json#/properties/MultiplexId",
								},
							},
							"required": []interface{}{"MultiplexId"},
						},
					},
				},
			},
		},
	}

	listSchema, err := ctx.gatherListHandlerSchema()
	require.NoError(t, err)
	require.NotNil(t, listSchema)
	assert.Equal(t, []string{"MultiplexId"}, listSchema.Required)
	assert.Equal(t, map[string]metadata.ListHandlerProperty{
		"MultiplexId": {
			Description: "The ID of the multiplex that the program belongs to.",
			Type:        "string",
		},
	}, listSchema.Properties)
}

func TestGatherListHandlerSchemaWithRefDifferentName(t *testing.T) {
	ctx := cfSchemaContext{
		resourceSpec: &jsschema.Schema{
			Properties: map[string]*jsschema.Schema{
				"ActualName": {
					Type:        jsschema.PrimitiveTypes{jsschema.StringType},
					Description: "Description from base property.",
				},
			},
			Extras: map[string]interface{}{
				"handlers": map[string]interface{}{
					"list": map[string]interface{}{
						"handlerSchema": map[string]interface{}{
							"properties": map[string]interface{}{
								"AliasName": map[string]interface{}{
									"$ref": "resource-schema.json#/properties/ActualName",
								},
							},
							"required": []interface{}{"AliasName"},
						},
					},
				},
			},
		},
	}

	listSchema, err := ctx.gatherListHandlerSchema()
	require.NoError(t, err)
	require.NotNil(t, listSchema)
	assert.Equal(t, []string{"AliasName"}, listSchema.Required)
	assert.Equal(t, map[string]metadata.ListHandlerProperty{
		"AliasName": {
			Description: "Description from base property.",
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

	listSchema, err := ctx.gatherListHandlerSchema()
	require.NoError(t, err)
	assert.Nil(t, listSchema)
}

func TestGatherListHandlerSchemaFailsOnMalformedHandlerSchema(t *testing.T) {
	ctx := cfSchemaContext{
		cfTypeName: "AWS::Test::Resource",
		resourceSpec: &jsschema.Schema{
			Extras: map[string]interface{}{
				"handlers": map[string]interface{}{
					"list": map[string]interface{}{
						"handlerSchema": map[string]interface{}{
							"properties": map[string]interface{}{
								"ScopeName": map[string]interface{}{
									"pattern": "missing type or ref",
								},
							},
						},
					},
				},
			},
		},
	}

	_, err := ctx.gatherListHandlerSchema()
	require.ErrorContains(t, err, `list handler property "ScopeName" for AWS::Test::Resource has no schema`)
}

func TestGatherListInputsEmptyWhenNoHandlerSchema(t *testing.T) {
	ctx := cfSchemaContext{
		resourceSpec: &jsschema.Schema{
			Extras: map[string]interface{}{
				"handlers": map[string]interface{}{
					"list": map[string]interface{}{},
				},
			},
		},
	}

	listInputs, err := ctx.gatherListInputs()
	require.NoError(t, err)
	require.NotNil(t, listInputs)
	assert.Equal(t, &pschema.ObjectTypeSpec{
		Type:       "object",
		Properties: map[string]pschema.PropertySpec{},
	}, listInputs)
}

func TestGatherListInputsFromHandlerSchema(t *testing.T) {
	ctx := cfSchemaContext{
		cfTypeName: "AWS::ApiGateway::Stage",
		resourceSpec: &jsschema.Schema{
			Extras: map[string]interface{}{
				"handlers": map[string]interface{}{
					"list": map[string]interface{}{
						"handlerSchema": map[string]interface{}{
							"properties": map[string]interface{}{
								"RestApiId": map[string]interface{}{
									"description": "The REST API ID.",
									"type":        "string",
								},
								"IncludeValues": map[string]interface{}{
									"type": "boolean",
								},
							},
							"required": []interface{}{"RestApiId"},
						},
					},
				},
			},
		},
	}

	listInputs, err := ctx.gatherListInputs()
	require.NoError(t, err)
	require.NotNil(t, listInputs)
	assert.Equal(t, "object", listInputs.Type)
	assert.Equal(t, []string{"restApiId"}, listInputs.Required)
	assert.Equal(t, map[string]pschema.PropertySpec{
		"includeValues": {
			TypeSpec: pschema.TypeSpec{Type: "boolean"},
		},
		"restApiId": {
			Description: "The REST API ID.",
			TypeSpec:    pschema.TypeSpec{Type: "string"},
		},
	}, listInputs.Properties)
}

func TestGatherListInputsMissingWhenNoListHandler(t *testing.T) {
	ctx := cfSchemaContext{
		resourceSpec: &jsschema.Schema{
			Extras: map[string]interface{}{},
		},
	}

	listInputs, err := ctx.gatherListInputs()
	require.NoError(t, err)
	assert.Nil(t, listInputs)
}

func TestGatherListInputsFailsOnSdkNameCollision(t *testing.T) {
	ctx := cfSchemaContext{
		cfTypeName: "AWS::Test::Resource",
		resourceSpec: &jsschema.Schema{
			Extras: map[string]interface{}{
				"handlers": map[string]interface{}{
					"list": map[string]interface{}{
						"handlerSchema": map[string]interface{}{
							"properties": map[string]interface{}{
								"URL": map[string]interface{}{
									"type": "string",
								},
								"Url": map[string]interface{}{
									"type": "string",
								},
							},
						},
					},
				},
			},
		},
	}

	_, err := ctx.gatherListInputs()
	require.ErrorContains(t, err, "both map to SDK name")
}

func TestGatherListInputsFailsWhenRequiredPropertyHasNoSchema(t *testing.T) {
	ctx := cfSchemaContext{
		cfTypeName: "AWS::Test::Resource",
		resourceSpec: &jsschema.Schema{
			Extras: map[string]interface{}{
				"handlers": map[string]interface{}{
					"list": map[string]interface{}{
						"handlerSchema": map[string]interface{}{
							"properties": map[string]interface{}{
								"ScopeName": map[string]interface{}{
									"$ref": "resource-schema.json#/properties/ScopeName",
								},
							},
							"required": []interface{}{"ScopeName"},
						},
					},
				},
			},
		},
	}

	_, err := ctx.gatherListInputs()
	require.ErrorContains(t, err, `list handler property "ScopeName" for AWS::Test::Resource has no schema`)
}
