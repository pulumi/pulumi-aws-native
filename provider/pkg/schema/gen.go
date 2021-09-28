// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	"encoding/json"
	"fmt"
	"strings"

	jsschema "github.com/lestrrat-go/jsschema"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	dotnetgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

const packageName = "aws-native"

// GatherPackage builds a package spec based on the provided CF JSON schemas.
func GatherPackage(supportedResourceTypes []string, jsonSchemas []jsschema.Schema,
	genAll bool) (*pschema.PackageSpec, *CloudAPIMetadata, error) {
	p := pschema.PackageSpec{
		Name:        packageName,
		Description: "A native Pulumi package for creating and managing Amazon Web Services (AWS) resources.",
		Keywords: []string{
			"pulumi",
			"aws",
			"aws-native",
		},
		Homepage:   "https://pulumi.com",
		License:    "Apache-2.0",
		Repository: "https://github.com/pulumi/pulumi-aws-native",
		Config: pschema.ConfigSpec{
			Variables: map[string]pschema.PropertySpec{
				"accessKey": {
					Description: "The access key for API operations. You can retrieve this from the ‘Security & Credentials’ section of the AWS console.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
				},
				"allowedAccountIds": {
					Description: "List of allowed AWS account IDs to prevent you from mistakenly using an incorrect one. Conflicts with `forbiddenAccountIds`.",
					TypeSpec:    pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Type: "string"}},
				},
				"assumeRole": {
					Description: "Configuration for retrieving temporary credentials from the STS service.",
					TypeSpec:    pschema.TypeSpec{Ref: "#/types/aws-native:config:AssumeRole"},
				},
				"defaultTags": {
					Description: "Configuration block with resource tag settings to apply across all resources handled by this provider. This is designed to replace redundant per-resource `tags` configurations. Provider tags can be overridden with new values, but not excluded from specific resources. To override provider tag values, use the `tags` argument within a resource to configure new tag values for matching keys.",
					TypeSpec:    pschema.TypeSpec{Ref: "#/types/aws-native:config:DefaultTags"},
				},
				"endpoints": {
					Description: "Configuration block for customizing service endpoints.",
					TypeSpec:    pschema.TypeSpec{Ref: "#/types/aws-native:config:Endpoints"},
				},
				"forbiddenAccountIds": {
					TypeSpec:    pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Type: "string"}},
					Description: "List of forbidden AWS account IDs to prevent you from mistakenly using the wrong one (and potentially end up destroying a live environment). Conflicts with `allowedAccountIds`.",
				},
				"ignoreTags": {
					Description: "Configuration block with resource tag settings to ignore across all resources handled by this provider (except any individual service tag resources such as `ec2.Tag`) for situations where external systems are managing certain resource tags.",
					TypeSpec:    pschema.TypeSpec{Ref: "#/types/aws-native:config:IgnoreTags"},
				},
				"insecure": {
					Description: "Explicitly allow the provider to perform \"insecure\" SSL requests. If omitted,default value is `false`.",
					TypeSpec:    pschema.TypeSpec{Type: "boolean"},
				},
				"maxRetries": {
					Description: "The maximum number of times an AWS API request is being executed. If the API request still fails, an error is thrown.",
					TypeSpec:    pschema.TypeSpec{Type: "integer"},
				},
				"profile": {
					Description: "The profile for API operations. If not set, the default profile created with `aws configure` will be used.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
				},
				"region": {
					Description: "The region where AWS operations will take place. Examples are `us-east-1`, `us-west-2`, etc.",
					TypeSpec: pschema.TypeSpec{
						Type: "string",
						Ref:  "#/types/aws-native:index/region:Region",
					},
				},
				"s3ForcePathStyle": {
					Description: "Set this to true to force the request to use path-style addressing, i.e., `http://s3.amazonaws.com/BUCKET/KEY`. By default, the S3 client will use virtual hosted bucket addressing when possible (`http://BUCKET.s3.amazonaws.com/KEY`). Specific to the Amazon S3 service.",
					TypeSpec:    pschema.TypeSpec{Type: "boolean"},
				},
				"secretKey": {
					Description: "The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
				},
				"sharedCredentialsFile": {
					Description: "The path to the shared credentials file. If not set this defaults to `~/.aws/credentials`.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
				},
				"skipCredentialsValidation": {
					Default:     true,
					Description: "Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS available/implemented.",
					TypeSpec:    pschema.TypeSpec{Type: "boolean"},
				},
				"skipGetEc2Platforms": {
					Default:     true,
					Description: "Skip getting the supported EC2 platforms. Used by users that don't have `ec2:DescribeAccountAttributes` permissions.",
					TypeSpec:    pschema.TypeSpec{Type: "boolean"},
				},
				"skipMetadataApiCheck": {
					Default:     true,
					Description: "Skip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint. Setting to true prevents Pulumi from authenticating via the Metadata API. You may need to use other authentication methods like static credentials, configuration variables, or environment variables.",
					TypeSpec:    pschema.TypeSpec{Type: "boolean"},
				},
				"skipRegionValidation": {
					Default:     true,
					Description: "Skip static validation of region name. Used by users of alternative AWS-like APIs or users with access to regions that are not public.",
					TypeSpec:    pschema.TypeSpec{Type: "boolean"},
				},
				"skipRequestingAccountId": {
					Description: "Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.",
					TypeSpec:    pschema.TypeSpec{Type: "boolean"},
				},
				"token": {
					Description: "Session token for validating temporary credentials. Typically provided after successful identity federation or Multi-Factor Authentication (MFA) login. With MFA login, this is the session token provided afterward, not the 6 digit MFA code used to get temporary credentials.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
				},
			},
			Required: []string{
				"region",
			},
		},
		Provider: pschema.ResourceSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Description: "The provider type for the AWS native package. By default, resources use package-wide configuration settings, however an explicit `Provider` instance may be created and passed during resource construction to achieve fine-grained programmatic control over provider settings. See the [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.",
				Properties: map[string]pschema.PropertySpec{
					"accessKey": {
						Description: "The access key for API operations. You can retrieve this from the ‘Security & Credentials’ section of the AWS console.",
						TypeSpec:    pschema.TypeSpec{Type: "string"},
					},
					"allowedAccountIds": {
						Description: "List of allowed AWS account IDs to prevent you from mistakenly using an incorrect one. Conflicts with `forbiddenAccountIds`.",
						TypeSpec:    pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Type: "string"}},
					},
					"assumeRole": {
						Description: "Configuration for retrieving temporary credentials from the STS service.",
						TypeSpec:    pschema.TypeSpec{Ref: "#/types/aws-native:index:ProviderAssumeRole"},
					},
					"defaultTags": {
						Description: "Configuration block with resource tag settings to apply across all resources handled by this provider. This is designed to replace redundant per-resource `tags` configurations. Provider tags can be overridden with new values, but not excluded from specific resources. To override provider tag values, use the `tags` argument within a resource to configure new tag values for matching keys.",
						TypeSpec:    pschema.TypeSpec{Ref: "#/types/aws-native:index:ProviderDefaultTags"},
					},
					"endpoints": {
						Description: "Configuration block for customizing service endpoints.",
						TypeSpec:    pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Ref: "#/types/aws-native:index:ProviderEndpoint"}},
					},
					"forbiddenAccountIds": {
						Description: "List of forbidden AWS account IDs to prevent you from mistakenly using the wrong one (and potentially end up destroying a live environment). Conflicts with `allowedAccountIds`.",
						TypeSpec:    pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Type: "string"}},
					},
					"ignoreTags": {
						Description: "Configuration block with resource tag settings to ignore across all resources handled by this provider (except any individual service tag resources such as `ec2.Tag`) for situations where external systems are managing certain resource tags.",
						TypeSpec:    pschema.TypeSpec{Ref: "#/types/aws-native:index:ProviderIgnoreTags"},
					},
					"insecure": {
						Description: "Explicitly allow the provider to perform \"insecure\" SSL requests. If omitted,default value is `false`.",
						TypeSpec:    pschema.TypeSpec{Type: "boolean"},
					},
					"maxRetries": {
						Description: "The maximum number of times an AWS API request is being executed. If the API request still fails, an error is thrown.",
						TypeSpec:    pschema.TypeSpec{Type: "integer"},
					},
					"profile": {
						Description: "The profile for API operations. If not set, the default profile created with `aws configure` will be used.",
						TypeSpec:    pschema.TypeSpec{Type: "string"},
					},
					"region": {
						Description: "The region where AWS operations will take place. Examples are `us-east-1`, `us-west-2`, etc.",
						TypeSpec: pschema.TypeSpec{
							Type: "string",
							Ref:  "#/types/aws-native:index/region:Region",
						},
					},
					"s3ForcePathStyle": {
						Description: "Set this to true to force the request to use path-style addressing, i.e., `http://s3.amazonaws.com/BUCKET/KEY`. By default, the S3 client will use virtual hosted bucket addressing when possible (`http://BUCKET.s3.amazonaws.com/KEY`). Specific to the Amazon S3 service.",
						TypeSpec:    pschema.TypeSpec{Type: "boolean"},
					},
					"secretKey": {
						Description: "The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.",
						TypeSpec:    pschema.TypeSpec{Type: "string"},
					},
					"sharedCredentialsFile": {
						Description: "The path to the shared credentials file. If not set this defaults to `~/.aws/credentials`.",
						TypeSpec:    pschema.TypeSpec{Type: "string"},
					},
					"skipCredentialsValidation": {
						Description: "Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS available/implemented.",
						TypeSpec:    pschema.TypeSpec{Type: "boolean"},
					},
					"skipGetEc2Platforms": {
						Description: "Skip getting the supported EC2 platforms. Used by users that don't have `ec2:DescribeAccountAttributes` permissions.",
						TypeSpec:    pschema.TypeSpec{Type: "boolean"},
					},
					"skipMetadataApiCheck": {
						Description: "Skip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint. Setting to true prevents Pulumi from authenticating via the Metadata API. You may need to use other authentication methods like static credentials, configuration variables, or environment variables.",
						TypeSpec:    pschema.TypeSpec{Type: "boolean"},
					},
					"skipRegionValidation": {
						Description: "Skip static validation of region name. Used by users of alternative AWS-like APIs or users with access to regions that are not public.",
						TypeSpec:    pschema.TypeSpec{Type: "boolean"},
					},
					"skipRequestingAccountId": {
						Description: "Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.",
						TypeSpec:    pschema.TypeSpec{Type: "boolean"},
					},
					"token": {
						DefaultInfo: &pschema.DefaultSpec{
							Environment: []string{
								"AWS_SESSION_TOKEN",
							},
						},
						Description: "Session token for validating temporary credentials. Typically provided after successful identity federation or Multi-Factor Authentication (MFA) login. With MFA login, this is the session token provided afterward, not the 6 digit MFA code used to get temporary credentials.",
						TypeSpec:    pschema.TypeSpec{Type: "string"},
					},
				},
			},
			InputProperties: map[string]pschema.PropertySpec{
				"accessKey": {
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"AWS_ACCESS_KEY_ID",
						},
					},
					Description: "The access key for API operations. You can retrieve this from the ‘Security & Credentials’ section of the AWS console.",
					Secret:      true,
					TypeSpec:    pschema.TypeSpec{Type: "string"},
				},
				"allowedAccountIds": {
					Description: "List of allowed AWS account IDs to prevent you from mistakenly using an incorrect one. Conflicts with `forbiddenAccountIds`.",
					TypeSpec:    pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Type: "string"}},
				},
				"assumeRole": {
					Description: "Configuration for retrieving temporary credentials from the STS service.",
					TypeSpec:    pschema.TypeSpec{Ref: "#/types/aws-native:index:ProviderAssumeRole"},
				},
				"defaultTags": {
					Description: "Configuration block with resource tag settings to apply across all resources handled by this provider. This is designed to replace redundant per-resource `tags` configurations. Provider tags can be overridden with new values, but not excluded from specific resources. To override provider tag values, use the `tags` argument within a resource to configure new tag values for matching keys.",
					TypeSpec:    pschema.TypeSpec{Ref: "#/types/aws-native:index:ProviderDefaultTags"},
				},
				"endpoints": {
					Description: "Configuration block for customizing service endpoints.",
					TypeSpec:    pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Ref: "#/types/aws-native:index:ProviderEndpoint"}},
				},
				"forbiddenAccountIds": {
					Description: "List of forbidden AWS account IDs to prevent you from mistakenly using the wrong one (and potentially end up destroying a live environment). Conflicts with `allowedAccountIds`.",
					TypeSpec:    pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Type: "string"}},
				},
				"ignoreTags": {
					Description: "Configuration block with resource tag settings to ignore across all resources handled by this provider (except any individual service tag resources such as `ec2.Tag`) for situations where external systems are managing certain resource tags.",
					TypeSpec:    pschema.TypeSpec{Ref: "#/types/aws-native:index:ProviderIgnoreTags"},
				},
				"insecure": {
					Description: "Explicitly allow the provider to perform \"insecure\" SSL requests. If omitted,default value is `false`.",
					TypeSpec:    pschema.TypeSpec{Type: "boolean"},
				},
				"maxRetries": {
					Description: "The maximum number of times an AWS API request is being executed. If the API request still fails, an error is thrown.",
					TypeSpec:    pschema.TypeSpec{Type: "integer"},
				},
				"profile": {
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"AWS_PROFILE",
						},
					},
					Description: "The profile for API operations. If not set, the default profile created with `aws configure` will be used.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
				},
				"region": {
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"AWS_REGION",
							"AWS_DEFAULT_REGION",
						},
					},
					Description: "The region where AWS operations will take place. Examples are `us-east-1`, `us-west-2`, etc.",
					TypeSpec: pschema.TypeSpec{
						Type: "string",
						Ref:  "#/types/aws-native:index/region:Region",
					},
				},
				"s3ForcePathStyle": {
					Description: "Set this to true to force the request to use path-style addressing, i.e., `http://s3.amazonaws.com/BUCKET/KEY`. By default, the S3 client will use virtual hosted bucket addressing when possible (`http://BUCKET.s3.amazonaws.com/KEY`). Specific to the Amazon S3 service.",
					TypeSpec:    pschema.TypeSpec{Type: "boolean"},
				},
				"secretKey": {
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"AWS_SECRET_ACCESS_KEY",
						},
					},
					Description: "The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.",
					Secret:      true,
					TypeSpec:    pschema.TypeSpec{Type: "string"},
				},
				"sharedCredentialsFile": {
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"AWS_SHARED_CREDENTIALS_FILE",
						},
					},
					Description: "The path to the shared credentials file. If not set this defaults to `~/.aws/credentials`.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
				},
				"skipCredentialsValidation": {
					Default:     true,
					Description: "Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS available/implemented.",
					TypeSpec:    pschema.TypeSpec{Type: "boolean"},
				},
				"skipGetEc2Platforms": {
					Default:     true,
					Description: "Skip getting the supported EC2 platforms. Used by users that don't have `ec2:DescribeAccountAttributes` permissions.",
					TypeSpec:    pschema.TypeSpec{Type: "boolean"},
				},
				"skipMetadataApiCheck": {
					Default:     true,
					Description: "Skip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint. Setting to true prevents Pulumi from authenticating via the Metadata API. You may need to use other authentication methods like static credentials, configuration variables, or environment variables.",
					TypeSpec:    pschema.TypeSpec{Type: "boolean"},
				},
				"skipRegionValidation": {
					Default:     true,
					Description: "Skip static validation of region name. Used by users of alternative AWS-like APIs or users with access to regions that are not public.",
					TypeSpec:    pschema.TypeSpec{Type: "boolean"},
				},
				"skipRequestingAccountId": {
					Description: "Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.",
					TypeSpec:    pschema.TypeSpec{Type: "boolean"},
				},
				"token": {
					DefaultInfo: &pschema.DefaultSpec{
						Environment: []string{
							"AWS_SESSION_TOKEN",
						},
					},
					Description: "Session token for validating temporary credentials. Typically provided after successful identity federation or Multi-Factor Authentication (MFA) login. With MFA login, this is the session token provided afterward, not the 6 digit MFA code used to get temporary credentials.",
					Secret:      true,
					TypeSpec:    pschema.TypeSpec{Type: "string"},
				},
			},
			RequiredInputs: []string{
				"region",
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

	metadata := CloudAPIMetadata{
		Resources: map[string]CloudAPIResource{},
		Types:     map[string]CloudAPIType{},
	}

	supportedResources := codegen.NewStringSet(supportedResourceTypes...)

	var resourceCount int
	for _, jsonSchema := range jsonSchemas {
		resourceName := jsonSchema.Extras["typeName"].(string)
		isSupported := supportedResources.Has(resourceName)
		if isSupported || genAll {
			ctx := &context{
				pkg:           &p,
				metadata:      &metadata,
				resourceName:  resourceName,
				resourceToken: typeToken(resourceName),
				resourceSpec:  &jsonSchema,
				visitedTypes:  codegen.NewStringSet(),
				isSupported:   isSupported,
			}
			err := ctx.gatherResourceType()
			if err != nil {
				return nil, nil, err
			}
			module := moduleName(resourceName)
			csharpNamespaces[strings.ToLower(module)] = module
			resourceCount++
		}
	}
	fmt.Printf("Number of resource types: %d\n", resourceCount)

	// If there are types in the overlays that do not exist in the schema (e.g., enum types), add them.
	for tok, overlayType := range typeOverlays {
		if _, typeExists := p.Types[tok]; !typeExists {
			p.Types[tok] = overlayType
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

	return &p, &metadata, nil
}

// context holds shared information for a single CF JSON schema.
type context struct {
	pkg           *pschema.PackageSpec
	metadata      *CloudAPIMetadata
	resourceName  string
	resourceToken string
	resourceSpec  *jsschema.Schema
	visitedTypes  codegen.StringSet
	isSupported   bool
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
		sdkName := ToSdkName(prop)
		if sdkName == "id" {
			continue
		}
		propertySpec, err := ctx.propertySpec(prop, resourceTypeName, spec)
		if err != nil {
			return err
		}
		properties[sdkName] = *propertySpec
		if !readOnlyProperties.Has(prop) {
			inputProperties[sdkName] = *propertySpec
		}
	}

	for _, prop := range ctx.resourceSpec.Required {
		sdkName := ToSdkName(prop)
		if _, has := properties[sdkName]; has {
			required.Add(sdkName)
		}
		if _, has := inputProperties[sdkName]; has {
			requiredInputs.Add(sdkName)
		}
	}
	for prop := range readOnlyProperties {
		sdkName := ToSdkName(prop)
		if _, has := properties[sdkName]; has {
			required.Add(sdkName)
		}
	}

	var deprecationMessage string
	if !ctx.isSupported {
		deprecationMessage = fmt.Sprintf("%s is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.", resourceTypeName)
	}
	ctx.pkg.Resources[ctx.resourceToken] = pschema.ResourceSpec{
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: ctx.resourceSpec.Description,
			Properties:  properties,
			Type:        "object",
			Required:    required.SortedValues(),
		},
		InputProperties:    inputProperties,
		RequiredInputs:     requiredInputs.SortedValues(),
		DeprecationMessage: deprecationMessage,
	}

	createOnlyProperties := codegen.NewStringSet()
	if rop, ok := ctx.resourceSpec.Extras["createOnlyProperties"].([]interface{}); ok {
		for _, propRef := range rop {
			cfName := strings.TrimPrefix(propRef.(string), "/properties/")
			createOnlyProperties.Add(ToSdkName(cfName))
		}
	}
	ctx.metadata.Resources[ctx.resourceToken] = CloudAPIResource{
		CfType:     ctx.resourceName,
		Inputs:     inputProperties,
		Outputs:    properties,
		CreateOnly: createOnlyProperties.SortedValues(),
		Required:   requiredInputs.SortedValues(),
	}
	return nil
}

func (ctx *context) propertySpec(propName, resourceTypeName string, spec *jsschema.Schema) (*pschema.PropertySpec, error) {
	typeSpec, err := ctx.propertyTypeSpec(propName, spec)
	if err != nil {
		return nil, err
	}
	propertySpec := pschema.PropertySpec{
		TypeSpec:    *typeSpec,
		Description: spec.Description,
	}
	if resourceTypeName == propName {
		propertySpec.Language = map[string]pschema.RawMessage{
			"csharp": rawMessage(dotnetgen.CSharpPropertyInfo{
				Name: propName + "Value",
			}),
		}
	}
	return &propertySpec, nil
}

// propertyTypeSpec converts a JSON type definition to a Pulumi type definition.
func (ctx *context) propertyTypeSpec(parentName string, propSchema *jsschema.Schema) (*pschema.TypeSpec, error) {
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
				specs, requiredSpecs, err := ctx.genProperties(schemaName, typeSchema)
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
				ctx.metadata.Types[tok] = CloudAPIType{
					Type:       "object",
					Properties: specs,
				}
			}

			referencedTypeName := fmt.Sprintf("#/types/%s", tok)
			return &pschema.TypeSpec{Ref: referencedTypeName}, nil
		} else {
			return ctx.propertyTypeSpec(schemaName, typeSchema)
		}
	}

	// Union types.
	if len(propSchema.AnyOf) > 0 {
		var types []pschema.TypeSpec
		for _, sch := range propSchema.AnyOf {
			typ, err := ctx.propertyTypeSpec(parentName, sch)
			if err != nil {
				return nil, err
			}
			types = append(types, *typ)
		}
		return &pschema.TypeSpec{
			OneOf: types,
		}, nil
	}

	if len(propSchema.Enum) > 0 {
		enum, err := ctx.genEnumType(parentName, propSchema)
		if err != nil {
			return nil, err
		}
		if enum != nil {
			return enum, nil
		}
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
			elementType, err := ctx.propertyTypeSpec(parentName+"Item", propSchema.Items.Schemas[0])
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

func (ctx *context) genProperties(parentName string, typeSchema *jsschema.Schema) (map[string]pschema.PropertySpec, codegen.StringSet, error) {
	specs := map[string]pschema.PropertySpec{}
	requiredSpecs := codegen.NewStringSet()
	for _, name := range codegen.SortedKeys(typeSchema.Properties) {
		value := typeSchema.Properties[name]
		sdkName := ToSdkName(name)

		typeSpec, err := ctx.propertyTypeSpec(parentName+name, value)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "property %s", name)
		}
		specs[sdkName] = pschema.PropertySpec{
			Description: value.Description,
			TypeSpec:    *typeSpec,
		}
	}
	for _, name := range typeSchema.Required {
		sdkName := ToSdkName(name)
		if _, has := specs[sdkName]; has {
			requiredSpecs.Add(sdkName)
		}
	}
	return specs, requiredSpecs, nil
}

// genEnumType generates the enum type for a given schema.
func (ctx *context) genEnumType(enumName string, propSchema *jsschema.Schema) (*pschema.TypeSpec, error) {
	if len(propSchema.Type) == 0 {
		return nil, nil
	}
	if propSchema.Type[0] != jsschema.StringType {
		return nil, nil
	}
	tok := ctx.resourceToken + enumName

	enumSpec := &pschema.ComplexTypeSpec{
		Enum: []pschema.EnumValueSpec{},
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: propSchema.Description,
			Type:        "string",
		},
	}

	values := codegen.NewStringSet()
	for _, val := range propSchema.Enum {
		str := ToUpperCamel(val.(string))
		if values.Has(str) {
			continue
		}
		values.Add(str)
		enumVal := pschema.EnumValueSpec{
			Value: val,
			Name:  str,
		}
		enumSpec.Enum = append(enumSpec.Enum, enumVal)
	}

	// Make sure that the type name we composed doesn't clash with another type
	// already defined in the schema earlier. The same enum does show up in multiple
	// places of specs, so we want to error only if they a) have the same name
	// b) the list of values does not match.
	if other, ok := ctx.pkg.Types[tok]; ok {
		same := len(enumSpec.Enum) == len(other.Enum)
		for _, val := range other.Enum {
			same = same && values.Has(val.Name)
		}
		if !same {
			return nil, errors.Errorf("duplicate enum %q: %+v vs. %+v", tok, enumSpec.Enum, other.Enum)
		}
	}
	ctx.pkg.Types[tok] = *enumSpec
	ctx.metadata.Types[tok] = CloudAPIType{Type: enumSpec.Type}

	referencedTypeName := fmt.Sprintf("#/types/%s", tok)
	return &pschema.TypeSpec{
		Ref: referencedTypeName,
	}, nil
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

	// Override name to avoid duplicate
	// See https://github.com/pulumi/pulumi/issues/8018
	switch typ {
	case "AWS::KinesisAnalytics::ApplicationOutput", "AWS::KinesisAnalyticsV2::ApplicationOutput":
		resourceName = strings.Replace(resourceName, "ApplicationOutput", "ApplicationOutputResource", 1)
	}

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
