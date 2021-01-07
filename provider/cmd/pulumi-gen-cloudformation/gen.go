package main

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/pulumi/pulumi-cloudformation/provider/pkg/schema"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	dotnetgen "github.com/pulumi/pulumi/pkg/v2/codegen/dotnet"
	pschema "github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

// TODO: convert docs. Currently all docs are just links. The linked pages should all have a meta tag somewhere in them
// that points at the (slightly odd) source markdown on GH.

const packageName = "cloudformation"

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
	moduleName, resourceName := resourceTypeComponents[1], resourceTypeComponents[2]

	// Override the name of the Config module.
	if moduleName == "Config" {
		moduleName = "Configuration"
	}
	return packageName + ":" + moduleName + ":" + resourceName + name
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
		properties[name] = pschema.PropertySpec{
			TypeSpec:    ctx.propertyTypeSpec(resourceScope, spec.ItemType, spec.PrimitiveItemType, spec.PrimitiveType, spec.Type),
			Description: spec.Documentation,
		}
		if spec.Required {
			required = append(required, name)
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
		name = strings.Replace(name, ".", "", -1)
		properties[name] = pschema.PropertySpec{
			TypeSpec: ctx.propertyTypeSpec(resourceScope, spec.ItemType, spec.PrimitiveItemType, spec.PrimitiveType, spec.Type),
		}
		required = append(required, name)
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

func (ctx *context) gatherResourceType(pkg *pschema.PackageSpec, name string, spec schema.ResourceSpec) {
	resourceToken := typeToken(name + ".")
	propertiesToken := typeToken(name + ".Properties")
	attributesToken := typeToken(name + ".Attributes")

	inputProperties := map[string]pschema.PropertySpec{
		"logicalId": {
			TypeSpec:    primitiveTypeSpec("String"),
			Description: "An explicit logical ID for the resource",
		},
		"metadata": {
			TypeSpec:    primitiveTypeSpec("Json"),
			Description: "Arbitrary structured data associated with the resource",
		},
		"deletionPolicy": {
			TypeSpec: primitiveTypeSpec("String"),
			Description: "With the deletionPolicy attribute you can preserve or (in some cases) backup a resource " +
				"when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you " +
				"want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the " +
				"resource by default.",
		},
		"updateReplacePolicy": {
			TypeSpec: primitiveTypeSpec("String"),
			Description: "Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing " +
				"physical instance of a resource when it is replaced during a stack update operation.",
		},
		"properties": {
			TypeSpec:    pschema.TypeSpec{Ref: "#/types/" + propertiesToken},
			Description: "The input properties associated with the resource",
		},
	}

	properties := map[string]pschema.PropertySpec{
		"logicalId":  inputProperties["logicalId"],
		"metadata":   inputProperties["metadata"],
		"properties": inputProperties["properties"],
		"attributes": {
			TypeSpec:    pschema.TypeSpec{Ref: "#/types/" + attributesToken},
			Description: "The attributes associated with the resource",
		},
	}

	if supportsCreationPolicy[name] {
		inputProperties["creationPolicy"] = pschema.PropertySpec{
			TypeSpec:    ctx.nestedTypeSpec("", "CreationPolicy"),
			Description: "The creation policy associated with the resource",
		}
		properties["creationPolicy"] = inputProperties["creationPolicy"]
	}
	if supportsUpdatePolicy[name] {
		inputProperties["updatePolicy"] = pschema.PropertySpec{
			TypeSpec:    ctx.nestedTypeSpec(name, "UpdatePolicy"),
			Description: "The update policy associated with the resource",
		}
		properties["updatePolicy"] = inputProperties["updatePolicy"]
	}

	pkg.Types[propertiesToken] = ctx.gatherPropertyType(name, spec.PropertyTypeSpec)
	pkg.Types[attributesToken] = ctx.gatherAttributesType(name, spec.Attributes)
	pkg.Resources[resourceToken] = pschema.ResourceSpec{
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: spec.Documentation,
			Properties:  properties,
			Type:        "object",
			Required:    []string{"properties", "attributes"},
		},
		InputProperties: inputProperties,
		RequiredInputs:  []string{"properties"},
	}
}

func gatherPackage(schema schema.CloudFormationSchema) pschema.PackageSpec {
	p := pschema.PackageSpec{
		Name:        packageName,
		Description: "A Pulumi package for creating and managing Amazon Web Services (AWS) resources via CloudFormation.",
		Keywords: []string{
			"pulumi",
			"aws",
			"cloudformation",
		},
		Homepage:   "https://pulumi.io",
		License:    "Apache-2.0",
		Repository: "https://github.com/pulumi/pulumi-cloudformation",
		Config: pschema.ConfigSpec{
			Variables: map[string]pschema.PropertySpec{
				"region": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "the region to use for deployments",
				},
				"stack": {
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Description: "the name of the stack to use for deployments",
				},
			},
			Required: []string{
				"region",
				"stack",
			},
		},
		Types:     map[string]pschema.ComplexTypeSpec{},
		Resources: map[string]pschema.ResourceSpec{},
		Functions: map[string]pschema.FunctionSpec{},
		Language:  map[string]json.RawMessage{},
	}
	p.Language["csharp"] = rawMessage(map[string]interface{}{
		"packageReferences": map[string]string{
			"Glob":                         "1.1.5",
			"Pulumi":                       "2.*",
			"System.Collections.Immutable": "1.6.0",
		},
		"dictionaryConstructors": true,
	})
	p.Language["go"] = rawMessage(map[string]interface{}{
		"importBasePath": "github.com/pulumi/pulumi-cloudformation/sdk/go/cloudformation",
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
	for name := range creationPolicyTypes {
		ctx.types.Add(name)
	}
	for name := range updatePolicyTypes {
		ctx.types.Add(name)
	}
	for name := range schema.PropertyTypes {
		ctx.types.Add(name)
	}

	// Add creation and update policy types.
	for name, spec := range creationPolicyTypes {
		p.Types[typeToken(name)] = spec
	}
	for name, spec := range updatePolicyTypes {
		p.Types[typeToken(name)] = spec
	}

	// Gather nested property types.
	for name, spec := range schema.PropertyTypes {
		resourceScope := ""
		components := strings.Split(name, ".")
		if len(components) > 1 {
			resourceScope = components[0]
		}
		p.Types[typeToken(name)] = ctx.gatherPropertyType(resourceScope, spec)
	}

	// Gather resource types.
	for name, spec := range schema.ResourceTypes {
		ctx.gatherResourceType(&p, name, spec)
	}

	// Apply manual overrides.
	if workspaceProperties, ok := p.Types[packageName+":WorkSpaces:WorkspaceProperties"]; ok {
		if prop, ok := workspaceProperties.Properties["WorkspaceProperties"]; ok {
			prop.Language = map[string]json.RawMessage{
				"csharp": rawMessage(dotnetgen.CSharpPropertyInfo{
					Name: "Properties",
				}),
			}
			workspaceProperties.Properties["WorkspaceProperties"] = prop
		}
	}
	if assetProperties, ok := p.Types[packageName+":IoTSiteWise:AssetProperties"]; ok {
		if prop, ok := assetProperties.Properties["AssetProperties"]; ok {
			prop.Language = map[string]json.RawMessage{
				"csharp": rawMessage(dotnetgen.CSharpPropertyInfo{
					Name: "Properties",
				}),
			}
			assetProperties.Properties["AssetProperties"] = prop
		}
	}
	if assetModelProperties, ok := p.Types[packageName+":IoTSiteWise:AssetModelProperties"]; ok {
		if prop, ok := assetModelProperties.Properties["AssetModelProperties"]; ok {
			prop.Language = map[string]json.RawMessage{
				"csharp": rawMessage(dotnetgen.CSharpPropertyInfo{
					Name: "Properties",
				}),
			}
			assetModelProperties.Properties["AssetModelProperties"] = prop
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

var supportsCreationPolicy = map[string]bool{
	"AWS::AutoScaling::AutoScalingGroup": true,
	"AWS::EC2::Instance":                 true,
	"AWS::CloudFormation::WaitCondition": true,
}

var creationPolicyTypes = map[string]pschema.ComplexTypeSpec{
	"CreationPolicy": {
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: "The CreationPolicy for a resource.",
			Properties: map[string]pschema.PropertySpec{
				"AutoScalingCreationPolicy": {
					TypeSpec: refTypeSpec("AutoscalingCreationPolicy"),
					Description: "For an Auto Scaling group replacement update, specifies how many instances must\n" +
						"signal success for the update to succeed.",
				},
				"ResourceSignal": {
					TypeSpec: refTypeSpec("ResourceSignal"),
					Description: "When AWS CloudFormation creates the associated resource, configures the number of\n" +
						"required success signals and the length of time that AWS CloudFormation waits for those signals.",
				},
			},
			Type: "object",
		},
	},
	"AutoscalingCreationPolicy": {
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Properties: map[string]pschema.PropertySpec{
				"MinSuccessfulInstancesPercent": {
					TypeSpec: primitiveTypeSpec("Integer"),
					Description: "Specifies the percentage of instances in an Auto Scaling replacement update that must\n" +
						"signal success for the update to succeed. You can specify a value from 0 to 100. AWS\n" +
						"CloudFormation rounds to the nearest tenth of a percent. For example, if you update five\n" +
						"instances with a minimum successful percentage of 50, three instances must signal success. If\n" +
						"an instance doesn't send a signal within the time specified by the Timeout property, AWS\n" +
						"CloudFormation assumes that the instance wasn't created.",
					Default: 100,
				},
			},
			Type: "object",
		},
	},
	"ResourceSignal": {
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Properties: map[string]pschema.PropertySpec{
				"Count": {
					Description: "The number of success signals AWS CloudFormation must receive before it sets the\n" +
						"resource status as CREATE_COMPLETE. If the resource receives a failure signal or doesn't\n" +
						"receive the specified number of signals before the timeout period expires, the resource\n" +
						"creation fails and AWS CloudFormation rolls the stack back.",
					TypeSpec: primitiveTypeSpec("Integer"),
					Default:  1,
				},
				"Timeout": {
					Description: "The length of time that AWS CloudFormation waits for the number of signals that was\n" +
						"specified in the Count property. The timeout period starts after AWS CloudFormation starts\n" +
						"creating the resource, and the timeout expires no sooner than the time you specify but can\n" +
						"occur shortly thereafter. The maximum time that you can specify is 12 hours.\n" +
						"\n" +
						"The value must be in ISO8601 duration format, in the form: \"PT#H#M#S\", where each # is the\n" +
						"number of hours, minutes, and seconds, respectively. For best results, specify a period of\n" +
						"time that gives your instances plenty of time to get up and running. A shorter timeout can\n" +
						"cause a rollback.",
					TypeSpec: primitiveTypeSpec("String"),
					Default:  "PT5M",
				},
			},
			Type: "object",
		},
	},
}

var supportsUpdatePolicy = map[string]bool{
	"AWS::AutoScaling::AutoScalingGroup": true,
	"AWS::ElastiCache::ReplicationGroup": true,
	"AWS::Elasticsearch::Domain":         true,
	"AWS::Lambda::Alias":                 true,
}

var updatePolicyTypes = map[string]pschema.ComplexTypeSpec{
	"AWS::AutoScaling::AutoScalingGroup.UpdatePolicy": {
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: "The UpdatePolicy for AutoScalingGroup resources.",
			Properties: map[string]pschema.PropertySpec{
				"AutoScalingReplacingUpdate": {
					Description: "To specify how AWS CloudFormation handles replacement updates for an Auto Scaling\n" +
						"group, use the AutoScalingReplacingUpdate policy. This policy enables you to specify whether\n" +
						"AWS CloudFormation replaces an Auto Scaling group with a new one or replaces only the\n" +
						"instances in the Auto Scaling group.",
					TypeSpec: refTypeSpec("AWS::AutoScaling::AutoScalingGroup.ReplacingUpdate"),
				},
				"AutoScalingRollingUpdate": {
					Description: "To specify how AWS CloudFormation handles rolling updates for an Auto Scaling group,\n" +
						"use the AutoScalingRollingUpdate policy. Rolling updates enable you to specify whether AWS\n" +
						"CloudFormation updates instances that are in an Auto Scaling group in batches or all at once.\n" +
						"\n" +
						"*Important*: During a rolling update, some Auto Scaling processes might make changes to the\n" +
						"Auto Scaling group before AWS CloudFormation completes the rolling update. These changes might\n" +
						"cause the rolling update to fail. To prevent Auto Scaling from running processes during a\n" +
						"rolling update, use the SuspendProcesses property. For more information, see\n" +
						"[What are some recommended best practices for performing Auto Scaling group rolling updates?](" +
						"https://aws.amazon.com/premiumsupport/knowledge-center/auto-scaling-group-rolling-updates/)\n" +
						"\n" +
						"Be aware that, during stack update rollback operations, CloudFormation uses the UpdatePolicy\n" +
						"configuration specified in the template before the current stack update operation. For\n" +
						"example, suppose you have updated the MaxBatchSize in your stack template's UpdatePolicy from\n" +
						"1 to 10. You then perform a stack update, and that update fails and CloudFormation initiates\n" +
						"an update rollback operation. In such a case, CloudFormation will use 1 as the maximum batch\n" +
						"size, rather than 10. For this reason, we recommend you make changes to the UpdatePolicy\n" +
						"configuration in a stack update separate from, and prior to, any updates to the\n" +
						"AutoScalingGroup resource that are likely to trigger rolling updates.",
					TypeSpec: refTypeSpec("AWS::AutoScaling::AutoScalingGroup.RollingUpdate"),
				},
				"AutoScalingScheduledAction": {
					Description: "To specify how AWS CloudFormation handles updates for the MinSize, MaxSize, and\n" +
						"DesiredCapacity properties when the AWS::AutoScaling::AutoScalingGroup resource has an\n" +
						"associated scheduled action, use the AutoScalingScheduledAction policy.\n" +
						"\n" +
						"With scheduled actions, the group size properties of an Auto Scaling group can change at any\n" +
						"time. When you update a stack with an Auto Scaling group and scheduled action, AWS\n" +
						"CloudFormation always sets the group size property values of your Auto Scaling group to the\n" +
						"values that are defined in the AWS::AutoScaling::AutoScalingGroup resource of your template,\n" +
						"even if a scheduled action is in effect.\n" +
						"\n" +
						"If you do not want AWS CloudFormation to change any of the group size property values when you\n" +
						"have a scheduled action in effect, use the AutoScalingScheduledAction update policy and set\n" +
						"IgnoreUnmodifiedGroupSizeProperties to true to prevent AWS CloudFormation from changing the\n" +
						"MinSize, MaxSize, or DesiredCapacity properties unless you have modified these values in your\n" +
						"template.",
					TypeSpec: refTypeSpec("AWS::AutoScaling::AutoScalingGroup.ScheduledAction"),
				},
			},
			Type: "object",
		},
	},
	"AWS::AutoScaling::AutoScalingGroup.ReplacingUpdate": {
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Properties: map[string]pschema.PropertySpec{
				"WillReplace": {
					Description: "Specifies whether an Auto Scaling group and the instances it contains are replaced\n" +
						"during an update. During replacement, AWS CloudFormation retains the old group until it\n" +
						"finishes creating the new one. If the update fails, AWS CloudFormation can roll back to the\n" +
						"old Auto Scaling group and delete the new Auto Scaling group.\n" +
						"\n" +
						"While AWS CloudFormation creates the new group, it doesn't detach or attach any instances.\n" +
						"After successfully creating the new Auto Scaling group, AWS CloudFormation deletes the old\n" +
						"Auto Scaling group during the cleanup process.\n" +
						"\n" +
						"When you set the WillReplace parameter, remember to specify a matching CreationPolicy. If the\n" +
						"minimum number of instances (specified by the MinSuccessfulInstancesPercent property) don't\n" +
						"signal success within the Timeout period (specified in the CreationPolicy policy), the\n" +
						"replacement update fails and AWS CloudFormation rolls back to the old Auto Scaling group.",
					TypeSpec: primitiveTypeSpec("Boolean"),
				},
			},
			Type: "object",
		},
	},
	"AWS::AutoScaling::AutoScalingGroup.RollingUpdate": {
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Properties: map[string]pschema.PropertySpec{
				"MaxBatchSize": {
					Description: "Specifies the maximum number of instances that AWS CloudFormation updates.",
					TypeSpec:    primitiveTypeSpec("Integer"),
					Default:     1,
				},
				"MinInstancesInService": {
					Description: "Specifies the minimum number of instances that must be in service within the Auto\n" +
						"Scaling group while AWS CloudFormation updates old instances. This value must be less than the\n" +
						"MaxSize of the Auto Scaling group.",
					TypeSpec: primitiveTypeSpec("Integer"),
					Default:  0,
				},
				"MinSuccessfulInstancesPercent": {
					Description: "Specifies the percentage of instances in an Auto Scaling rolling update that must\n" +
						"signal success for an update to succeed. You can specify a value from 0 to 100. AWS\n" +
						"CloudFormation rounds to the nearest tenth of a percent. For example, if you update five\n" +
						"instances with a minimum successful percentage of 50, three instances must signal success.\n" +
						"\n" +
						"If an instance doesn't send a signal within the time specified in the PauseTime property, AWS\n" +
						"CloudFormation assumes that the instance wasn't updated.\n" +
						"\n" +
						"If you specify this property, you must also enable the WaitOnResourceSignals and PauseTime\n" +
						"properties.\n" +
						"\n" +
						"The MinSuccessfulInstancesPercent parameter applies only to instances only for signaling\n" +
						"purpose. To specify the number of instances in your autoscaling group, see the MinSize,\n" +
						"MaxSize, and DesiredCapacity properties fo the AWS::AutoScaling::AutoScalingGroup resource.",
					TypeSpec: primitiveTypeSpec("Integer"),
					Default:  100,
				},
				"PauseTime": {
					Description: "The amount of time that AWS CloudFormation pauses after making a change to a batch of\n" +
						"instances to give those instances time to start software applications. For example, you might\n" +
						"need to specify PauseTime when scaling up the number of instances in an Auto Scaling group.\n" +
						"\n" +
						"If you enable the WaitOnResourceSignals property, PauseTime is the amount of time that AWS\n" +
						"CloudFormation should wait for the Auto Scaling group to receive the required number of valid\n" +
						"signals from added or replaced instances. If the PauseTime is exceeded before the Auto Scaling\n" +
						"group receives the required number of signals, the update fails. For best results, specify a\n" +
						"time period that gives your applications sufficient time to get started. If the update needs\n" +
						"to be rolled back, a short PauseTime can cause the rollback to fail.\n" +
						"\n" +
						"Specify PauseTime in the ISO8601 duration format (in the format PT#H#M#S, where each # is the\n" +
						"number of hours, minutes, and seconds, respectively). The maximum PauseTime is one hour (PT1H).",
					TypeSpec: primitiveTypeSpec("String"),
					Default:  "PT5M",
				},
				"SuspendProcesses": {
					Description: "Specifies the Auto Scaling processes to suspend during a stack update. Suspending\n" +
						"processes prevents Auto Scaling from interfering with a stack update. For example, you can\n" +
						"suspend alarming so that Amazon EC2 Auto Scaling doesn't execute scaling policies associated\n" +
						"with an alarm. For valid values, see the ScalingProcesses.member.N parameter for the\n" +
						"SuspendProcesses action in the Amazon EC2 Auto Scaling API Reference.",
					TypeSpec: pschema.TypeSpec{
						Type:  "array",
						Items: &pschema.TypeSpec{Type: "string"},
					},
				},
				"WaitOnResourceSignals": {
					Description: "Specifies whether the Auto Scaling group waits on signals from new instances during\n" +
						"an update. Use this property to ensure that instances have completed installing and\n" +
						"configuring applications before the Auto Scaling group update proceeds. AWS CloudFormation\n" +
						"suspends the update of an Auto Scaling group after new EC2 instances are launched into the\n" +
						"group. AWS CloudFormation must receive a signal from each new instance within the specified\n" +
						"PauseTime before continuing the update. To signal the Auto Scaling group, use the cfn-signal\n" +
						"helper script or SignalResource API.\n" +
						"\n" +
						"To have instances wait for an Elastic Load Balancing health check before they signal success,\n" +
						"add a health-check verification by using the cfn-init helper script. For an example, see the\n" +
						"verify_instance_health command in the Auto Scaling rolling updates sample template.",
					TypeSpec: primitiveTypeSpec("Boolean"),
					Default:  false,
				},
			},
			Type: "object",
		},
	},
	"AWS::AutoScaling::AutoScalingGroup.ScheduledAction": {
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Properties: map[string]pschema.PropertySpec{
				"IgnoreUnmodifiedGroupSizeProperties": {
					Description: "If true, AWS CloudFormation ignores differences in group size properties between your\n" +
						"current Auto Scaling group and the Auto Scaling group described in the	" +
						"AWS::AutoScaling::AutoScalingGroup resource of your template during a stack update. If you\n" +
						"modify any of the group size property values in your template, AWS CloudFormation uses the\n" +
						"modified values and updates your Auto Scaling group.",
					TypeSpec: primitiveTypeSpec("Boolean"),
					Default:  false,
				},
			},
			Type: "object",
		},
	},
	"AWS::ElastiCache::ReplicationGroup.UpdatePolicy": {
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: "The UpdatePolicy for a ReplicationGroup.",
			Properties: map[string]pschema.PropertySpec{
				"UseOnlineResharding": {
					Description: "To modify a replication group's shards by adding or removing shards, rather than\n" +
						"replacing the entire AWS::ElastiCache::ReplicationGroup resource, use the\n" +
						"UseOnlineResharding update policy.\n" +
						"\n" +
						"If UseOnlineResharding is set to true, you can update the NumNodeGroups and\n" +
						"NodeGroupConfiguration properties of the AWS::ElastiCache::ReplicationGroup resource, and\n" +
						"CloudFormation will update those properties without interruption. When UseOnlineResharding\n" +
						"is set to false, or not specified, updating the NumNodeGroups and NodeGroupConfiguration\n" +
						"properties results in CloudFormation replacing the entire\n" +
						"AWS::ElastiCache::ReplicationGroup resource.\n" +
						"\n" +
						"The UseOnlineResharding update policy has no properties.\n" +
						"\n" +
						"Things to consider when setting the UseOnlineResharding update policy to true:\n" +
						"\n" +
						"- We strongly recommend you perform updates to the NumNodeGroups and\n" +
						"  NodeGroupConfiguration properties as the only updates in a given stack update operation.\n" +
						"  Updating the node group configuration of a replication group is a resource-intensive\n" +
						"  operation. If a stack update fails, CloudFormation does not roll back changes to the node\n" +
						"  group configuration of a replication group. However, CloudFormation will roll back any\n" +
						"  other properties that were changed as part of the failed update operation.\n" +
						"- Any node group updates require identifying all node groups.  If you specify the\n" +
						"  NodeGroupConfiguration property, you must also specify the NodeGroupId for each node group\n" +
						"  configuration in order for CloudFormation to update the number of nodes without\n" +
						"  interruption.  When creating a replication group, if you do not specify an ID for each\n" +
						"  node group, ElastiCache automatically generates an ID for each node group. To update the\n" +
						"  replication group without interruption, use the ElastiCache console\n" +
						"  (https://console.aws.amazon.com/elasticache/) or DescribeReplicationGroups to retrieve the\n" +
						"  IDs for all node groups in the replication group. Then specify the ID for each node group\n" +
						"  in your stack template before attempting to add or remove shards.\n" +
						"\n" +
						"  *Note*: As a best practice, when you create a replication group in a stack template,\n" +
						"  include an ID for each node group you specify.\n" +
						"\n" +
						"  In addition, updating the number of nodes without interruption requires that you have\n" +
						"  accurately specified the PrimaryAvailabilityZone, ReplicaAvailabilityZones, and\n" +
						"  ReplicaCount properties for each NodeGroupConfiguration as well. Again, you can use the\n" +
						"  ElastiCache console (https://console.aws.amazon.com/elasticache/) or\n" +
						"  DescribeReplicationGroups to retrieve the actual values for each node group and compare\n" +
						"  them to the values in your stack template. You can update the property values of the\n" +
						"  node groups as a separate stack update, or as part of the same stack update that changes\n" +
						"  the number of node groups.  When you use an UseOnlineResharding update policy to update\n" +
						"  the number of node groups without interruption, ElastiCache evenly distributes the\n" +
						"  keyspaces between the specified number of slots. This cannot be updated later.\n" +
						"  Therefore, after updating the number of node groups in this way, you should remove the\n" +
						"  value specified for the Slots property of each NodeGroupConfiguration from the stack\n" +
						"  template, as it no longer reflects the actual values in each node group.\n\n" +
						"- Actual node group removal results may vary.  When you specify a NumNodeGroups value that\n" +
						"  is less than the current number of node groups, CloudFormation instructs ElastiCache to\n" +
						"  remove as many node groups as necessary to reach the specified number of nodes. However,\n" +
						"  ElastiCache may not always be able to remove the desired number of node groups. In the\n" +
						"  event ElastiCache cannot remove the desired number of node groups, CloudFormation\n" +
						"  generates a stack event alerting you to this. In cases where ElastiCache cannot remove any\n" +
						"  node groups, the CloudFormation resource update fails.\n" +
						"\n" +
						"For more information on modifying replication groups, see\n" +
						"ModifyReplicationGroupShardConfiguration in the Amazon ElastiCache API Reference.",
					TypeSpec: primitiveTypeSpec("Boolean"),
				},
			},
			Type: "object",
		},
	},
	"AWS::Elasticsearch::Domain.UpdatePolicy": {
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: "The UpdatePolicy for a Domain resource.",
			Properties: map[string]pschema.PropertySpec{
				"EnableVersionUpgrade": {
					Description: "To upgrade an Amazon ES domain to a new version of Elasticsearch rather than " +
						"replacing the entire AWS::Elasticsearch::Domain resource, use the EnableVersionUpgrade update " +
						"policy.\n" +
						"\n" +
						"If EnableVersionUpgrade is set to true, you can update the ElasticsearchVersion property " +
						"of the AWS::Elasticsearch::Domain resource, and CloudFormation will update that property " +
						"without interruption. When EnableVersionUpgrade is set to false, or not specified, " +
						"updating the ElasticsearchVersion property results in CloudFormation replacing the entire " +
						"AWS::Elasticsearch::Domain resource.\n" +
						"\n" +
						"The EnableVersionUpgrade update policy has no properties.\n" +
						"\n" +
						"For more information about upgrading Amazon ES domains, see UpgradeElasticsearchDomain in " +
						"the Amazon Elasticsearch Service Developer Guide.",
					TypeSpec: primitiveTypeSpec("Boolean"),
				},
			},
			Type: "object",
		},
	},
	"AWS::Lambda::Alias.UpdatePolicy": {
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: "The UpdatePolicy for an Alias resource.",
			Properties: map[string]pschema.PropertySpec{
				"CodeDeployLambdaAliasUpdate": {
					Description: "To perform an CodeDeploy deployment when the version changes on an AWS::Lambda::Alias\n" +
						"resource, use the CodeDeployLambdaAliasUpdate update policy.",
					TypeSpec: refTypeSpec("AWS::Lambda::Alias.CodeDeployLambdaAliasUpdate"),
				},
			},
			Type: "object",
		},
	},
	"AWS::Lambda::Alias.CodeDeployLambdaAliasUpdate": {
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Properties: map[string]pschema.PropertySpec{
				"AfterAllowTrafficHook": {
					Description: "The name of the Lambda function to run after traffic routing completes.",
					TypeSpec:    primitiveTypeSpec("String"),
				},
				"ApplicationName": {
					Description: "The name of the CodeDeploy application.",
					TypeSpec:    primitiveTypeSpec("String"),
				},
				"BeforeAllowTrafficHook": {
					Description: "The name of the Lambda function to run before traffic routing starts.",
					TypeSpec:    primitiveTypeSpec("String"),
				},
				"DeploymentGroupName": {
					Description: "The name of the CodeDeploy deployment group. This is where the traffic-shifting\n" +
						"policy is set.",
					TypeSpec: primitiveTypeSpec("String"),
				},
			},
			Type:     "object",
			Required: []string{"ApplicationName", "DeploymentGroupName"},
		},
	},
}

var pseudoParameters = map[string]string{
	"AccountId": "accountId",
	"Partition": "partition",
	"Region":    "region",
	"StackName": "stackName",
	"StackId":   "stackId",
	"UrlSuffix": "urlSuffix",
}
