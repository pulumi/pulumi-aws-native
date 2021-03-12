package main

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/schema"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	dotnetgen "github.com/pulumi/pulumi/pkg/v2/codegen/dotnet"
	pschema "github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

// TODO: convert docs. Currently all docs are just links. The linked pages should all have a meta tag somewhere in them
// that points at the (slightly odd) source markdown on GH.

const packageName = "aws-native"

type context struct {
	types codegen.StringSet
}

func rawMessage(v interface{}) json.RawMessage {
	bytes, err := json.Marshal(v)
	contract.Assert(err == nil)
	return bytes
}

func typeToken(typ string) string {
	components := strings.Split(typ, ".")
	if len(components) == 1 {
		return packageName + ":index:" + components[0]
	}

	resourceType, name := components[0], components[1]
	resourceTypeComponents := strings.Split(resourceType, "::")
	contract.Assertf(len(resourceTypeComponents) == 3, "expected three parts in type %q", resourceTypeComponents)
	moduleName, resourceName := resourceTypeComponents[1], resourceTypeComponents[2]

	// Override the name of the Config module.
	if moduleName == "Config" {
		moduleName = "Configuration"
	}
	return packageName + ":" + moduleName + ":" + resourceName + name
}

func typeName(typ string) string {
	components := strings.Split(typ, ".")
	if len(components) == 1 {
		return components[0]
	}

	resourceType := components[0]
	resourceTypeComponents := strings.Split(resourceType, "::")
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

func refTypeSpec(typ string) pschema.TypeSpec {
	return pschema.TypeSpec{Ref: "#/types/" + typeToken(typ)}
}

func (ctx *context) nestedTypeSpec(resourceScope, typ string) pschema.TypeSpec {
	qualifiedName := resourceScope + "." + typ
	if !ctx.types.Has(qualifiedName) {
		if !ctx.types.Has(typ) {
			return pschema.TypeSpec{Ref: "pulumi.json#/Any"}
		}
		qualifiedName = typ
	}
	return refTypeSpec(qualifiedName)
}

func (ctx *context) itemTypeSpec(resourceScope, itemType, primitiveItemType string) pschema.TypeSpec {
	if primitiveItemType != "" {
		return primitiveTypeSpec(primitiveItemType)
	}
	return ctx.nestedTypeSpec(resourceScope, itemType)
}

func (ctx *context) propertyTypeSpec(resourceScope, itemType, primitiveItemType, primitiveType, typ string) pschema.TypeSpec {
	if primitiveType != "" {
		return primitiveTypeSpec(primitiveType)
	}

	switch typ {
	case "List":
		elementType := ctx.itemTypeSpec(resourceScope, itemType, primitiveItemType)
		return pschema.TypeSpec{
			Type:  "array",
			Items: &elementType,
		}
	case "Map":
		elementType := ctx.itemTypeSpec(resourceScope, itemType, primitiveItemType)
		return pschema.TypeSpec{
			Type:                 "object",
			AdditionalProperties: &elementType,
		}
	case "":
		// TODO: We hit this on a bunch of properties:
		// AWS::MediaPackage::PackagingConfiguration.SpekeKeyProvider
		// AWS::DataBrew::Recipe.Parameters
		// AWS::ECR::Repository.RepositoryPolicyText
		// AWS::SageMaker::ModelPackageGroup.ModelPackageGroupPolicy
		// AWS::KMS::Key.KeyPolicy
		// Make special cases for those?
		return pschema.TypeSpec{Ref: "pulumi.json#/Any"}
	default:
		return ctx.nestedTypeSpec(resourceScope, typ)
	}
}

func (ctx *context) gatherPropertyType(resourceScope string, spec schema.PropertyTypeSpec) pschema.ComplexTypeSpec {
	properties, required := map[string]pschema.PropertySpec{}, []string{}
	for name, spec := range spec.Properties {
		sdkName := schema.ToPropertyName(name)
		properties[sdkName] = pschema.PropertySpec{
			TypeSpec:    ctx.propertyTypeSpec(resourceScope, spec.ItemType, spec.PrimitiveItemType, spec.PrimitiveType, spec.Type),
			Description: spec.Documentation,
		}
		if spec.Required {
			required = append(required, sdkName)
		}
	}
	sort.Strings(required)

	return pschema.ComplexTypeSpec{
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: spec.Documentation,
			Properties:  properties,
			Type:        "object",
			Required:    required,
		},
	}
}

func (ctx *context) gatherAttributesType(resourceScope string, attributes map[string]schema.AttributeSpec) pschema.ComplexTypeSpec {
	properties, required := map[string]pschema.PropertySpec{}, []string{}
	for name, spec := range attributes {
		sdkName := schema.ToPropertyName(strings.Replace(name, ".", "", -1))
		properties[sdkName] = pschema.PropertySpec{
			TypeSpec: ctx.propertyTypeSpec(resourceScope, spec.ItemType, spec.PrimitiveItemType, spec.PrimitiveType, spec.Type),
		}
		required = append(required, sdkName)
	}
	sort.Strings(required)

	return pschema.ComplexTypeSpec{
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Properties: properties,
			Type:       "object",
			Required:   required,
		},
	}
}

func (ctx *context) gatherResourceType(pkg *pschema.PackageSpec, resourceName string, resourceSpec schema.ResourceSpec) {
	resourceToken := typeToken(resourceName + ".")
	resourceTypeName := typeName(resourceName + ".")

	inputProperties, requiredInputs := map[string]pschema.PropertySpec{}, []string{}
	properties, required := map[string]pschema.PropertySpec{}, []string{}
	for prop, spec := range resourceSpec.Properties {
		propertySpec := pschema.PropertySpec{
			TypeSpec:    ctx.propertyTypeSpec(resourceName, spec.ItemType, spec.PrimitiveItemType, spec.PrimitiveType, spec.Type),
			Description: spec.Documentation,
		}
		if resourceTypeName == prop {
			propertySpec.Language = map[string]json.RawMessage{
				"csharp": rawMessage(dotnetgen.CSharpPropertyInfo{
					Name: prop + "Value",
				}),
			}
		}
		sdkName := schema.ToPropertyName(prop)
		properties[sdkName] = propertySpec
		inputProperties[sdkName] = propertySpec
		if spec.Required {
			required = append(required, sdkName)
			requiredInputs = append(requiredInputs, sdkName)
		}
	}
	sort.Strings(required)

	for attr, spec := range resourceSpec.Attributes {
		attr = strings.Replace(attr, ".", "", -1)
		propertySpec := pschema.PropertySpec{
			TypeSpec: ctx.propertyTypeSpec(resourceName, spec.ItemType, spec.PrimitiveItemType, spec.PrimitiveType, spec.Type),
		}
		if resourceTypeName == attr {
			propertySpec.Language = map[string]json.RawMessage{
				"csharp": rawMessage(dotnetgen.CSharpPropertyInfo{
					Name: attr + "Value",
				}),
			}
		}
		sdkName := schema.ToPropertyName(attr)
		properties[sdkName] = propertySpec
		required = append(required, sdkName)
	}
	sort.Strings(required)

	pkg.Resources[resourceToken] = pschema.ResourceSpec{
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: resourceSpec.Documentation,
			Properties:  properties,
			Type:        "object",
			Required:    required,
		},
		InputProperties: inputProperties,
		RequiredInputs:  requiredInputs,
	}
}

func gatherPackage(schema schema.CloudFormationSchema, supportedResourceTypes []string) pschema.PackageSpec {
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
					Description: "the region to use for deployments",
				},
			},
			Required: []string{
				"region",
			},
		},
		Types:     map[string]pschema.ComplexTypeSpec{},
		Resources: map[string]pschema.ResourceSpec{},
		Functions: map[string]pschema.FunctionSpec{},
		Language:  map[string]json.RawMessage{},
	}
	csharpNamespaces := map[string]string{
		"aws-native": "AwsNative",
	}
	p.Language["csharp"] = rawMessage(map[string]interface{}{
		"packageReferences": map[string]string{
			"Pulumi":                       "2.*",
			"System.Collections.Immutable": "1.6.0",
		},
		"namespaces": csharpNamespaces,
	})
	p.Language["go"] = rawMessage(map[string]interface{}{
		"importBasePath": "github.com/pulumi/pulumi-aws-native/sdk/go/aws",
		"packageImportAliases": map[string]string{
			"github.com/pulumi/pulumi-aws-native/sdk/go/aws/aws-native": "aws",
		},
	})
	p.Language["nodejs"] = rawMessage(map[string]interface{}{
		"dependencies": map[string]string{
			"@pulumi/pulumi":    "^2.0.0",
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
			"pulumi":   ">=2.0.0,<3.0.0",
			"requests": ">=2.21.0,<2.22.0",
			"pyyaml":   ">=5.1,<5.2",
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

	ctx := &context{types: codegen.StringSet{}}

	// Prime the type set.
	for name := range schema.PropertyTypes {
		ctx.types.Add(name)
	}

	supportedResources := codegen.NewStringSet(supportedResourceTypes...)

	// Gather nested property types.
	for name, spec := range schema.PropertyTypes {
		resourceScope := ""
		components := strings.Split(name, ".")
		if len(components) > 1 {
			resourceScope = components[0]
		}
		if supportedResources.Has(resourceScope) {
			p.Types[typeToken(name)] = ctx.gatherPropertyType(resourceScope, spec)
		}
	}

	// Gather resource types.
	for name, spec := range schema.ResourceTypes {
		if supportedResources.Has(name) {
			ctx.gatherResourceType(&p, name, spec)
		}
	}

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

	// Add CFN instrinsics.
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

	return p
}
var pseudoParameters = map[string]string{
	"AccountId": "accountId",
	"Partition": "partition",
	"Region":    "region",
	"UrlSuffix": "urlSuffix",
}
