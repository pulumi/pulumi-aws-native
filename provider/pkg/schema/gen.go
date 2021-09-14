// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"strings"

	jsschema "github.com/lestrrat-go/jsschema"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	dotnetgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

const packageName = "aws-native"

// GatherPackage builds a package spec based on the provided CF JSON schemas.
func GatherPackage(supportedResourceTypes []string, jsonSchemas []jsschema.Schema) (*pschema.PackageSpec, error) {
	p := pschema.PackageSpec{
		Name:        packageName,
		Description: "A Pulumi package for creating and managing Amazon Web Services (AWS) resources via CloudFormation.",
		Keywords: []string{
			"pulumi",
			"aws",
			"aws-native",
		},
		Homepage:   "https://pulumi.io",
		License:    "Apache-2.0",
		Repository: "https://github.com/pulumi/pulumi-aws-native",
		Config: pschema.ConfigSpec{
			Variables: map[string]pschema.PropertySpec{
				"region": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: `The region where AWS operations will take place. Examples are "us-east-1", "us-west-2", etc.`,
				},
			},
			Required: []string{
				"region",
			},
		},
		Provider: pschema.ResourceSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Description: "The provider type for the native AWS package.",
				Type:        "object",
			},
			InputProperties: map[string]pschema.PropertySpec{
				"region": {
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"AWS_REGION",
							"AWS_DEFAULT_REGION",
						},
					},
					Description: `The region where AWS operations will take place. Examples are "us-east-1", "us-west-2", etc.`,
					TypeSpec:    pschema.TypeSpec{Type: "string"},
				},
			},
		},
		Types:     map[string]pschema.ComplexTypeSpec{},
		Resources: map[string]pschema.ResourceSpec{},
		Functions: map[string]pschema.FunctionSpec{},
		Language:  map[string]pschema.RawMessage{},
	}
	csharpNamespaces := map[string]string{
		"aws-native": "AwsNative",
	}
	p.Language["go"] = rawMessage(map[string]interface{}{
		"importBasePath": "github.com/pulumi/pulumi-aws-native/sdk/go/aws",
		"packageImportAliases": map[string]string{
			"github.com/pulumi/pulumi-aws-native/sdk/go/aws/aws-native": "aws",
		},
	})
	p.Language["nodejs"] = rawMessage(map[string]interface{}{
		"dependencies": map[string]string{
			"@pulumi/pulumi":    "^3.0.0",
			"shell-quote":       "^1.6.1",
			"tmp":               "^0.0.33",
			"@types/tmp":        "^0.0.33",
			"glob":              "^7.1.2",
			"@types/glob":       "^5.0.35",
			"node-fetch":        "^2.3.0",
			"@types/node-fetch": "^2.1.4",
		},
		"devDependencies": map[string]string{
			"mocha":              "^5.2.0",
			"@types/mocha":       "^5.2.5",
			"@types/shell-quote": "^1.6.0",
		},
	})
	p.Language["python"] = rawMessage(map[string]interface{}{
		"requires": map[string]string{
			"pulumi": ">=3.0.0,<4.0.0",
		},
	})

	p.Provider = pschema.ResourceSpec{
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Properties: p.Config.Variables,
			Type:       "object",
			Required:   p.Config.Required,
		},
		InputProperties: p.Config.Variables,
		RequiredInputs:  p.Config.Required,
	}

	supportedResources := codegen.NewStringSet(supportedResourceTypes...)

	var resourceCount int
	for _, jsonSchema := range jsonSchemas {
		resourceName := jsonSchema.Extras["typeName"].(string)
		if supportedResources.Has(resourceName) {
			ctx := &context{
				pkg:           &p,
				resourceName:  resourceName,
				resourceToken: typeToken(resourceName),
				resourceSpec:  &jsonSchema,
				visitedTypes:  codegen.NewStringSet(),
			}
			err := ctx.gatherResourceType()
			if err != nil {
				return nil, err
			}
			resourceCount++
		}
	}
	fmt.Printf("Number of resource types: %d\n", resourceCount)

	// Add getters for CFN pseudo parameters.
	for name, property := range pseudoParameters {
		p.Functions[packageName+":index:get"+name] = pschema.FunctionSpec{
			Outputs: &pschema.ObjectTypeSpec{
				Properties: map[string]pschema.PropertySpec{
					property: {TypeSpec: primitiveTypeSpec("String")},
				},
				Required: []string{property},
			},
		}
	}

	p.Language["csharp"] = rawMessage(map[string]interface{}{
		"packageReferences": map[string]string{
			"Pulumi": "3.*",
		},
		"namespaces": csharpNamespaces,
	})

	// Add CFN intrinsics.
	p.Functions[packageName+":index:getAzs"] = pschema.FunctionSpec{
		Inputs: &pschema.ObjectTypeSpec{
			Properties: map[string]pschema.PropertySpec{
				"region": {TypeSpec: primitiveTypeSpec("String")},
			},
		},
		Outputs: &pschema.ObjectTypeSpec{
			Properties: map[string]pschema.PropertySpec{
				"azs": {
					TypeSpec: pschema.TypeSpec{
						Type:  "array",
						Items: &pschema.TypeSpec{Type: "string"},
					},
				},
			},
			Required: []string{"azs"},
		},
	}
	p.Functions[packageName+":index:cidr"] = pschema.FunctionSpec{
		Inputs: &pschema.ObjectTypeSpec{
			Properties: map[string]pschema.PropertySpec{
				"ipBlock":  {TypeSpec: primitiveTypeSpec("String")},
				"count":    {TypeSpec: primitiveTypeSpec("Integer")},
				"cidrBits": {TypeSpec: primitiveTypeSpec("Integer")},
			},
			Required: []string{"ipBlock", "count", "cidrBits"},
		},
		Outputs: &pschema.ObjectTypeSpec{
			Properties: map[string]pschema.PropertySpec{
				"subnets": {
					TypeSpec: pschema.TypeSpec{
						Type:  "array",
						Items: &pschema.TypeSpec{Type: "string"},
					},
				},
			},
			Required: []string{"subnets"},
		},
	}
	p.Functions[packageName+":index:getSsmParameterString"] = pschema.FunctionSpec{
		Inputs: &pschema.ObjectTypeSpec{
			Properties: map[string]pschema.PropertySpec{
				"name": {TypeSpec: primitiveTypeSpec("String")},
			},
			Required: []string{"name"},
		},
		Outputs: &pschema.ObjectTypeSpec{
			Properties: map[string]pschema.PropertySpec{
				"value": {TypeSpec: primitiveTypeSpec("String")},
			},
			Required: []string{"value"},
		},
	}
	p.Functions[packageName+":index:getSsmParameterList"] = pschema.FunctionSpec{
		Inputs: &pschema.ObjectTypeSpec{
			Properties: map[string]pschema.PropertySpec{
				"name": {TypeSpec: primitiveTypeSpec("String")},
			},
			Required: []string{"name"},
		},
		Outputs: &pschema.ObjectTypeSpec{
			Properties: map[string]pschema.PropertySpec{
				"value": {
					TypeSpec: pschema.TypeSpec{
						Type:  "array",
						Items: &pschema.TypeSpec{Type: "string"},
					},
				},
			},
			Required: []string{"value"},
		},
	}
	p.Functions[packageName+":index:importValue"] = pschema.FunctionSpec{
		Inputs: &pschema.ObjectTypeSpec{
			Properties: map[string]pschema.PropertySpec{
				"name": {TypeSpec: primitiveTypeSpec("String")},
			},
			Required: []string{"name"},
		},
		Outputs: &pschema.ObjectTypeSpec{
			Properties: map[string]pschema.PropertySpec{
				"value": {TypeSpec: pschema.TypeSpec{Ref: "pulumi.json#/Any"}},
			},
		},
	}

	return &p, nil
}

// context holds shared information for a single CF JSON schema.
type context struct {
	pkg           *pschema.PackageSpec
	resourceName  string
	resourceToken string
	resourceSpec  *jsschema.Schema
	visitedTypes  codegen.StringSet
}

// gatherResourceType builds the schema for the resource type in the context.
func (ctx *context) gatherResourceType() error {
	resourceTypeName := typeName(ctx.resourceName)

	readOnlyProperties := codegen.NewStringSet()
	if rop, ok := ctx.resourceSpec.Extras["readOnlyProperties"].([]interface{}); ok {
		for _, propRef := range rop {
			propName := strings.TrimPrefix(propRef.(string), "/properties/")
			readOnlyProperties.Add(propName)
		}
	}

	inputProperties, requiredInputs := map[string]pschema.PropertySpec{}, codegen.NewStringSet()
	properties, required := map[string]pschema.PropertySpec{}, codegen.NewStringSet()
	for prop, spec := range ctx.resourceSpec.Properties {
		sdkName := ToPropertyName(prop)
		if sdkName == "id" {
			continue
		}
		typeSpec, err := ctx.propertyTypeSpec(spec)
		if err != nil {
			return err
		}
		propertySpec := pschema.PropertySpec{
			TypeSpec:    *typeSpec,
			Description: spec.Description,
		}
		if resourceTypeName == prop {
			propertySpec.Language = map[string]pschema.RawMessage{
				"csharp": rawMessage(dotnetgen.CSharpPropertyInfo{
					Name: prop + "Value",
				}),
			}
		}
		properties[sdkName] = propertySpec
		if !readOnlyProperties.Has(prop) {
			inputProperties[sdkName] = propertySpec
		}
	}

	for _, prop := range ctx.resourceSpec.Required {
		sdkName := ToPropertyName(prop)
		if _, has := properties[sdkName]; has {
			required.Add(sdkName)
		}
		if _, has := inputProperties[sdkName]; has {
			requiredInputs.Add(sdkName)
		}
	}
	for prop := range readOnlyProperties {
		sdkName := ToPropertyName(prop)
		if _, has := properties[sdkName]; has {
			required.Add(sdkName)
		}
	}

	ctx.pkg.Resources[ctx.resourceToken] = pschema.ResourceSpec{
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: ctx.resourceSpec.Description,
			Properties:  properties,
			Type:        "object",
			Required:    required.SortedValues(),
		},
		InputProperties: inputProperties,
		RequiredInputs:  requiredInputs.SortedValues(),
	}
	return nil
}

// propertyTypeSpec converts a JSON type definition to a Pulumi type definition.
func (ctx *context) propertyTypeSpec(propSchema *jsschema.Schema) (*pschema.TypeSpec, error) {
	// References to other type definitions.
	if propSchema.Reference != "" {
		schemaName := strings.TrimPrefix(propSchema.Reference, "#/definitions/")
		tok := fmt.Sprintf(`%s%s`, ctx.resourceToken, schemaName)

		typeSchema, ok := ctx.resourceSpec.Definitions[schemaName]
		if !ok {
			return nil, errors.Errorf("definition %s not found in resource %s", schemaName, ctx.resourceName)
		}
		baseType := "object"
		if len(typeSchema.Type) > 0 {
			baseType = typeSchema.Type[0].String()
		}
		if baseType == "object" {
			if !ctx.visitedTypes.Has(tok) {
				ctx.visitedTypes.Add(tok)
				specs, requiredSpecs, err := ctx.genProperties(typeSchema)
				if err != nil {
					return nil, err
				}

				ctx.pkg.Types[tok] = pschema.ComplexTypeSpec{
					ObjectTypeSpec: pschema.ObjectTypeSpec{
						Description: typeSchema.Description,
						Type:        "object",
						Properties:  specs,
						Required:    requiredSpecs.SortedValues(),
					},
				}
			}

			referencedTypeName := fmt.Sprintf("#/types/%s", tok)
			return &pschema.TypeSpec{Ref: referencedTypeName}, nil
		} else {
			return ctx.propertyTypeSpec(typeSchema)
		}
	}

	// Union types.
	if len(propSchema.AnyOf) > 0 {
		var types []pschema.TypeSpec
		for _, sch := range propSchema.AnyOf {
			typ, err := ctx.propertyTypeSpec(sch)
			if err != nil {
				return nil, err
			}
			types = append(types, *typ)
		}
		return &pschema.TypeSpec{
			OneOf: types,
		}, nil
	}

	// All other types.
	if len(propSchema.Type) > 0 {
		switch propSchema.Type[0] {
		case jsschema.IntegerType:
			return &pschema.TypeSpec{Type: "integer"}, nil
		case jsschema.StringType:
			return &pschema.TypeSpec{Type: "string"}, nil
		case jsschema.BooleanType:
			return &pschema.TypeSpec{Type: "boolean"}, nil
		case jsschema.NumberType:
			return &pschema.TypeSpec{Type: "number"}, nil
		case jsschema.ObjectType:
			return &pschema.TypeSpec{Ref: "pulumi.json#/Any"}, nil
		case jsschema.ArrayType:
			elementType, err := ctx.propertyTypeSpec(propSchema.Items.Schemas[0])
			if err != nil {
				return nil, err
			}
			return &pschema.TypeSpec{
				Type:  "array",
				Items: elementType,
			}, nil
		}
	}

	return nil, errors.New("failed to generate property types")
}

func (ctx *context) genProperties(typeSchema *jsschema.Schema) (map[string]pschema.PropertySpec, codegen.StringSet, error) {
	specs := map[string]pschema.PropertySpec{}
	requiredSpecs := codegen.NewStringSet()
	for _, name := range codegen.SortedKeys(typeSchema.Properties) {
		value := typeSchema.Properties[name]
		sdkName := ToPropertyName(name)

		typeSpec, err := ctx.propertyTypeSpec(value)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "property %s", name)
		}
		specs[sdkName] = pschema.PropertySpec{
			Description: value.Description,
			TypeSpec:    *typeSpec,
		}
	}
	for _, name := range typeSchema.Required {
		sdkName := ToPropertyName(name)
		if _, has := specs[sdkName]; has {
			requiredSpecs.Add(sdkName)
		}
	}
	return specs, requiredSpecs, nil
}

var pseudoParameters = map[string]string{
	"AccountId": "accountId",
	"Partition": "partition",
	"Region":    "region",
	"UrlSuffix": "urlSuffix",
}

func rawMessage(v interface{}) pschema.RawMessage {
	bytes, err := json.Marshal(v)
	contract.Assert(err == nil)
	return bytes
}

func moduleName(resourceType string) string {
	resourceTypeComponents := strings.Split(resourceType, "::")
	module := resourceTypeComponents[1]

	// Override the name of the Config module.
	if module == "Config" {
		module = "Configuration"
	}

	return module
}

func typeToken(typ string) string {
	resourceTypeComponents := strings.Split(typ, "::")
	resourceName := resourceTypeComponents[2]
	module := strings.ToLower(moduleName(typ))
	contract.Assertf(len(resourceTypeComponents) == 3, "expected three parts in type %q", resourceTypeComponents)
	return fmt.Sprintf("%s:%s:%s", packageName, module, resourceName)
}

func typeName(typ string) string {
	resourceTypeComponents := strings.Split(typ, "::")
	contract.Assertf(len(resourceTypeComponents) == 3, "expected three parts in type %q", resourceTypeComponents)
	return resourceTypeComponents[2]
}

func primitiveTypeSpec(primitiveType string) pschema.TypeSpec {
	switch primitiveType {
	case "Json":
		return pschema.TypeSpec{
			Type: "string",
			OneOf: []pschema.TypeSpec{
				{Ref: "pulumi.json#/Json"},
				{Type: "string"},
			},
		}
	case "String", "Timestamp":
		return pschema.TypeSpec{Type: "string"}
	case "Double":
		return pschema.TypeSpec{Type: "number"}
	case "Long", "Integer":
		return pschema.TypeSpec{Type: "integer"}
	case "Boolean":
		return pschema.TypeSpec{Type: "boolean"}
	default:
		contract.Failf("unexpected primitive type '%v'", primitiveType)
		return pschema.TypeSpec{Ref: "pulumi.json#/Any"}
	}
}
