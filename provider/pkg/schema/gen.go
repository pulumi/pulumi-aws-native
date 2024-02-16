// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	"encoding/json"
	goerrors "errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	jsschema "github.com/lestrrat-go/jsschema"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	dotnetgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

const packageName = "aws-native"
const globalTagToken = "aws-native:index:Tag"

// GatherPackage builds a package spec based on the provided CF JSON schemas.
func GatherPackage(supportedResourceTypes []string, jsonSchemas []jsschema.Schema,
	genAll bool, semanticsDocument *SemanticsSpecDocument) (*pschema.PackageSpec, *CloudAPIMetadata, *Reports, error) {
	globalTagType := pschema.ObjectTypeSpec{
		Type:        "object",
		Description: "A set of tags to apply to the resource.",
		Properties: map[string]pschema.PropertySpec{
			"key":   {TypeSpec: primitiveTypeSpec("String"), Description: "The key name of the tag"},
			"value": {TypeSpec: primitiveTypeSpec("String"), Description: "The value of the tag"},
		},
		Required: []string{"key", "value"},
	}
	p := pschema.PackageSpec{
		Name:        packageName,
		Description: "A native Pulumi package for creating and managing Amazon Web Services (AWS) resources.",
		DisplayName: "AWS Native",
		Keywords: []string{
			"pulumi",
			"aws",
			"aws-native",
			"category/cloud",
			"kind/native",
		},
		Homepage:   "https://pulumi.com",
		License:    "Apache-2.0",
		Publisher:  "Pulumi",
		Repository: "https://github.com/pulumi/pulumi-aws-native",
		Config: pschema.ConfigSpec{
			Variables: map[string]pschema.PropertySpec{
				"accessKey": {
					Description: "The access key for API operations. You can retrieve this from the ‘Security & Credentials’ section of the AWS console.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
					Secret:      true,
				},
				"allowedAccountIds": {
					Description: "List of allowed AWS account IDs to prevent you from mistakenly using an incorrect one. Conflicts with `forbiddenAccountIds`.",
					TypeSpec:    pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Type: "string"}},
				},
				"assumeRole": {
					Description: "Configuration for retrieving temporary credentials from the STS service.",
					TypeSpec:    pschema.TypeSpec{Ref: "#/types/aws-native:config:AssumeRole"},
				},
				"roleArn": {
					Description: "The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role for Cloud Control API to use when performing this resource operation. Note, this is a unique feature for server side security enforcement, not to be confused with assumeRole, which is used to obtain temporary client credentials. If you do not specify a role, Cloud Control API uses a temporary session created using your AWS user credentials instead.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
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
					Secret:      true,
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
					Secret:      true,
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
					"allowedAccountIds": {
						Description: "List of allowed AWS account IDs to prevent you from mistakenly using an incorrect one. Conflicts with `forbiddenAccountIds`.",
						TypeSpec:    pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Type: "string"}},
					},
					"assumeRole": {
						Description: "Configuration for retrieving temporary credentials from the STS service.",
						TypeSpec:    pschema.TypeSpec{Ref: "#/types/aws-native:index:ProviderAssumeRole"},
					},
					"roleArn": {
						Description: "The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role for Cloud Control API to use when performing this resource operation. Note, this is a unique feature for server side security enforcement, not to be confused with assumeRole, which is used to obtain temporary client credentials. If you do not specify a role, Cloud Control API uses a temporary session created using your AWS user credentials instead.",
						TypeSpec:    pschema.TypeSpec{Type: "string"},
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
				},
			},
			InputProperties: map[string]pschema.PropertySpec{
				"accessKey": {
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
				"roleArn": {
					Description: "The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role for Cloud Control API to use when performing this resource operation. Note, this is a unique feature for server side security enforcement, not to be confused with assumeRole, which is used to obtain temporary client credentials. If you do not specify a role, Cloud Control API uses a temporary session created using your AWS user credentials instead.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
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
					Description: "Session token for validating temporary credentials. Typically provided after successful identity federation or Multi-Factor Authentication (MFA) login. With MFA login, this is the session token provided afterward, not the 6 digit MFA code used to get temporary credentials.",
					Secret:      true,
					TypeSpec:    pschema.TypeSpec{Type: "string"},
				},
			},
			RequiredInputs: []string{
				"region",
			},
		},
		Types: map[string]pschema.ComplexTypeSpec{
			globalTagToken: {
				ObjectTypeSpec: globalTagType,
			},
		},
		Resources: map[string]pschema.ResourceSpec{
			ExtensionResourceToken: {
				ObjectTypeSpec: pschema.ObjectTypeSpec{
					Description: "A special resource that enables deploying CloudFormation Extensions (third-party resources). An extension has to be pre-registered in your AWS account in order to use this resource.",
					Properties: map[string]pschema.PropertySpec{
						"outputs": {
							Description: "Dictionary of the extension resource attributes.",
							TypeSpec: pschema.TypeSpec{
								Type: "object",
								AdditionalProperties: &pschema.TypeSpec{
									Ref: "pulumi.json#/Any",
								},
							},
						},
					},
					Required: []string{"outputs"},
				},
				InputProperties: map[string]pschema.PropertySpec{
					"type": {
						Description: "CloudFormation type name.",
						TypeSpec: pschema.TypeSpec{
							Type: "string",
						},
					},
					"properties": {
						Description: "Dictionary of the extension resource properties.",
						TypeSpec: pschema.TypeSpec{
							Type: "object",
							AdditionalProperties: &pschema.TypeSpec{
								Ref: "pulumi.json#/Any",
							},
						},
					},
				},
				RequiredInputs: []string{"type", "properties"},
			},
		},
		Functions: map[string]pschema.FunctionSpec{},
		Language:  map[string]pschema.RawMessage{},
	}
	csharpNamespaces := map[string]string{
		"aws-native": "AwsNative",
	}
	javaPackages := map[string]string{
		"aws-native": "awsnative",
	}
	p.Language["java"] = rawMessage(map[string]interface{}{
		"packages": javaPackages,
	})
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
		"pyproject": map[string]bool{
			"enabled": true,
		},
	})

	metadata := CloudAPIMetadata{
		Resources: map[string]CloudAPIResource{
			ExtensionResourceToken: {
				Inputs:     p.Resources[ExtensionResourceToken].InputProperties,
				Outputs:    p.Resources[ExtensionResourceToken].Properties,
				CreateOnly: []string{"type", "properties"},
			},
		},
		Types: map[string]CloudAPIType{
			globalTagToken: {
				Type:       "object",
				Properties: globalTagType.Properties,
			},
		},
		Functions: map[string]CloudAPIFunction{},
	}

	reports := &Reports{
		UnexpectedTagsShapes: map[string]interface{}{},
	}

	supportedResources := codegen.NewStringSet(supportedResourceTypes...)

	var resourceCount int
	for _, jsonSchema := range jsonSchemas {
		cfTypeName := jsonSchema.Extras["typeName"].(string)
		isSupported := supportedResources.Has(cfTypeName)
		if isSupported || genAll {
			resourceName, resourceToken := typeToken(cfTypeName)
			fullMod := moduleName(cfTypeName)
			mod := strings.ToLower(fullMod)
			ctx := &context{
				pkg:           &p,
				metadata:      &metadata,
				cfTypeName:    cfTypeName,
				mod:           mod,
				resourceName:  resourceName,
				resourceToken: resourceToken,
				resourceSpec:  &jsonSchema,
				visitedTypes:  codegen.NewStringSet(),
				isSupported:   isSupported,
				semantics:     semanticsDocument,
				reports:       reports,
			}
			err := ctx.gatherResourceType()
			if err != nil {
				return nil, nil, nil, err
			}
			err = ctx.gatherInvoke()
			if err != nil {
				return nil, nil, nil, err
			}
			csharpNamespaces[mod] = fullMod
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

	p.Functions[packageName+":index:getPartition"] = pschema.FunctionSpec{
		Outputs: &pschema.ObjectTypeSpec{
			Properties: map[string]pschema.PropertySpec{
				"partition": {
					TypeSpec:    primitiveTypeSpec("String"),
					Description: "Identifier of the current partition (e.g., `aws` in AWS Commercial, `aws-cn` in AWS China).",
				},
				"dnsSuffix": {
					TypeSpec:    primitiveTypeSpec("String"),
					Description: "Base DNS domain name for the current partition (e.g., `amazonaws.com` in AWS Commercial, `amazonaws.com.cn` in AWS China).",
				},
			},
			Required: []string{"partition", "dnsSuffix"},
		},
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

	return &p, &metadata, reports, nil
}

// context holds shared information for a single CF JSON schema.
type context struct {
	pkg           *pschema.PackageSpec
	metadata      *CloudAPIMetadata
	cfTypeName    string
	mod           string
	resourceName  string
	resourceToken string
	resourceSpec  *jsschema.Schema
	visitedTypes  codegen.StringSet
	isSupported   bool
	semantics     *SemanticsSpecDocument
	reports       *Reports
}

func (ctx *context) markCreateOnlyProperties(createOnlyProperties codegen.StringSet, resource *pschema.ResourceSpec) error {
	errs := []error{}
	for propPath, _ := range createOnlyProperties {
		// each path in createOnlyProperties is delimited with "/"
		path := strings.Split(propPath, "/")
		prop, ok := resource.Properties[path[0]]
		if !ok {
			errs = append(errs, errors.Errorf("Could not mark createOnlyProperty %s in %s as replaceOnChanges: property not found on Resource", propPath, ctx.cfTypeName))
			continue
		}
		err := ctx.markCreateOnlyProperty(path[1:], &prop)
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "Could not mark createOnlyProperty %s in %s as replaceOnChanges", propPath, ctx.cfTypeName))
		}
		resource.Properties[path[0]] = prop
	}
	if len(errs) > 0 {
		return goerrors.Join(errs...)
	}
	return nil
}

func (ctx *context) markCreateOnlyProperty(propPath []string, property *pschema.PropertySpec) error {
	// base case, just set the flag
	if len(propPath) == 0 {
		property.ReplaceOnChanges = true
		return nil
	}

	// Strip explicit traversal of array from the path
	if propPath[0] == "*" {
		return ctx.markCreateOnlyProperty(propPath[1:], property)
	}

	// Default to next path component is ref to object, but exclude global types
	if len(property.Ref) != 0 && !strings.HasPrefix(property.Ref, "#/types/aws-native:index") {
		return ctx.markCreateOnlyPropertyOnType(propPath[0], property.Ref, propPath[1:])
	}

	// Next try looking at it as an array of objects
	if property.Items != nil && len(property.Items.Ref) != 0 && !strings.HasPrefix(property.Items.Ref, "#/types/aws-native:index") {
		typeRef := property.Items.Ref
		if len(typeRef) == 0 {
			return errors.Errorf("Property does not reference an array: %v", property)
		}
		return ctx.markCreateOnlyPropertyOnType(propPath[0], typeRef, propPath[1:])
	}

	// Give up
	return errors.Errorf("Property is not a Ref or Array[Ref], can't traverse: %v", property)
}

func (ctx *context) markCreateOnlyPropertyOnType(propName, typeRef string, remainingPath []string) error {
	ref := strings.TrimPrefix(typeRef, "#/types/")
	propType, ok := ctx.pkg.Types[ref]
	if !ok {
		return errors.Errorf("Could not find referencedType: " + typeRef)
	}

	propertyName := ToSdkName(propName)
	innerProperty, ok := propType.Properties[propertyName]
	if !ok {
		return errors.Errorf("Type %s does not have property named '%s'", typeRef, propertyName)
	}
	err := ctx.markCreateOnlyProperty(remainingPath, &innerProperty)
	if err != nil {
		return err
	}
	propType.Properties[propertyName] = innerProperty
	ctx.pkg.Types[ref] = propType
	return nil
}

func (ctx *context) gatherInvoke() error {
	resourceTypeName := typeName(ctx.cfTypeName)
	_, getterToken := getterToken(ctx.cfTypeName)

	// Skip resource that causes naming clashes
	switch ctx.cfTypeName {
	case "AWS::MediaConnect::FlowOutput":
		return nil
	}

	primaryIdentifier := readPropNames(ctx.resourceSpec, "primaryIdentifier")

	inputs := map[string]pschema.PropertySpec{}
	inputNames := make([]string, len(primaryIdentifier))
	for i, r := range primaryIdentifier {
		n := strings.TrimPrefix(r, "/properties/")
		sdkName := ToSdkName(n)
		inputNames[i] = sdkName
		s, ok := ctx.resourceSpec.Properties[n]
		if !ok {
			fmt.Printf("Unable to find primary identifier %q for resource %s, skipping\n", n, resourceTypeName)
			return nil
		}

		p, err := ctx.propertySpec(n, resourceTypeName, s)
		if err != nil {
			return err
		}
		inputs[sdkName] = *p
	}

	writeOnlyProperties := readPropNames(ctx.resourceSpec, "writeOnlyProperties")
	createOnlyProperties := readPropNames(ctx.resourceSpec, "createOnlyProperties")
	nonOutputProperties := codegen.NewStringSet(append(writeOnlyProperties, createOnlyProperties...)...)

	outputs := map[string]pschema.PropertySpec{}
	for k, v := range ctx.resourceSpec.Properties {
		sdkName := ToSdkName(k)
		if nonOutputProperties.Has(k) {
			continue
		}
		p, err := ctx.propertySpec(k, resourceTypeName, v)
		if err != nil {
			return err
		}
		outputs[sdkName] = *p
	}

	if len(outputs) == 0 {
		// If there are no outputs to be returned, (all are write/create properties), no need to emit the function.
		return nil
	}

	ctx.pkg.Functions[getterToken] = pschema.FunctionSpec{
		Description: ctx.resourceSpec.Description,
		Inputs: &pschema.ObjectTypeSpec{
			Properties: inputs,
			Required:   inputNames,
		},
		Outputs: &pschema.ObjectTypeSpec{
			Properties: outputs,
		},
	}

	identifiers := make([]string, len(primaryIdentifier))
	for i, v := range primaryIdentifier {
		identifiers[i] = ToSdkName(v)
	}
	ctx.metadata.Functions[getterToken] = CloudAPIFunction{
		CfType:      ctx.cfTypeName,
		Identifiers: identifiers,
	}

	return nil
}

func readPropNames(resourceSpec *jsschema.Schema, listName string) []string {
	if p, ok := resourceSpec.Extras[listName]; ok {
		pSlice := p.([]interface{})
		props := make([]string, len(pSlice))
		for i, v := range pSlice {
			n := strings.TrimPrefix(v.(string), "/properties/")
			props[i] = n
		}
		return props
	}
	return make([]string, 0)
}

func readPropSdkNames(resourceSpec *jsschema.Schema, listName string) codegen.StringSet {
	output := codegen.NewStringSet()
	if p, ok := resourceSpec.Extras[listName]; ok {
		pSlice := p.([]interface{})
		for _, v := range pSlice {
			n := strings.TrimPrefix(v.(string), "/properties/")
			output.Add(ToSdkName(n))
		}
	}
	return output
}

// gatherResourceType builds the schema for the resource type in the context.
func (ctx *context) gatherResourceType() error {
	resourceTypeName := typeName(ctx.cfTypeName)

	writeOnlyProperties := readPropSdkNames(ctx.resourceSpec, "writeOnlyProperties")
	createOnlyProperties := readPropSdkNames(ctx.resourceSpec, "createOnlyProperties")
	readOnlyProperties := codegen.NewStringSet(readPropNames(ctx.resourceSpec, "readOnlyProperties")...)

	irreversibleNames := map[string]string{}
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
		if createOnlyProperties.Has(sdkName) || !readOnlyProperties.Has(prop) {
			inputProperties[sdkName] = *propertySpec
		}
		if HasUppercaseAcronym(prop) {
			irreversibleNames[sdkName] = prop
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

	autoNamingSpec := ctx.createAutoNamingSpec(inputProperties, resourceTypeName, properties)
	// If a field can be auto-named, its no longer required.
	if autoNamingSpec != nil {
		delete(requiredInputs, autoNamingSpec.SdkName)
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
	resourceSpec := pschema.ResourceSpec{
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

	err := ctx.markCreateOnlyProperties(createOnlyProperties, &resourceSpec)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	ctx.pkg.Resources[ctx.resourceToken] = resourceSpec

	ctx.metadata.Resources[ctx.resourceToken] = CloudAPIResource{
		CfType:            ctx.cfTypeName,
		Inputs:            inputProperties,
		Outputs:           properties,
		CreateOnly:        createOnlyProperties.SortedValues(),
		Required:          requiredInputs.SortedValues(),
		AutoNamingSpec:    autoNamingSpec,
		WriteOnly:         writeOnlyProperties.SortedValues(),
		IrreversibleNames: irreversibleNames,
	}
	return nil
}

func (ctx *context) createAutoNamingSpec(inputProperties map[string]pschema.PropertySpec, resourceTypeName string, properties map[string]pschema.PropertySpec) *AutoNamingSpec {
	// ** Autonaming **
	var autoNameSpec *AutoNamingSpec

	semanticsSpec := ctx.semantics.Resources[ctx.resourceToken]

	lookForField := func(fieldName string) func(map[string]pschema.PropertySpec) {
		return func(properties map[string]pschema.PropertySpec) {
			sdkName := ToSdkName(fieldName)
			if propSpec, has := inputProperties[sdkName]; has && propSpec.Type == "string" {
				autoNameSpec = &AutoNamingSpec{
					SdkName: sdkName,
				}
				spec, ok := ctx.resourceSpec.Properties[fieldName]
				if ok {
					autoNameSpec.MinLength = spec.MinLength.Val
					autoNameSpec.MaxLength = spec.MaxLength.Val
				}
				namingTriviaSpec, ok := semanticsSpec.NamingTriviaSpec[sdkName]
				if ok {
					autoNameSpec.TriviaSpec = &namingTriviaSpec
				}
			}
		}
	}

	autoNamePropLookupFuncs := []func(map[string]pschema.PropertySpec){
		lookForField("Name"),
		lookForField(resourceTypeName + "Name"),
	}

	for _, fn := range autoNamePropLookupFuncs {
		fn(properties)
		if autoNameSpec != nil {
			break
		}
	}

	return autoNameSpec
}

func (ctx *context) propertySpec(propName, resourceTypeName string, spec *jsschema.Schema) (*pschema.PropertySpec, error) {
	typeSpec, err := ctx.propertyTypeSpec(lowerAcronyms(propName), spec)
	if err != nil {
		return nil, err
	}
	propertySpec := pschema.PropertySpec{
		TypeSpec:    *typeSpec,
		Description: spec.Description,
	}

	// If the property name is the same as the resource type name, we need to rename the property's C# name
	// to avoid 'member names cannot be the same as their enclosing type'. We normalize the property and resource
	// type names to lowercase to avoid case sensitivity issues.
	if strings.ToLower(resourceTypeName) == strings.ToLower(propName) {
		propertySpec.Language = map[string]pschema.RawMessage{
			"csharp": rawMessage(dotnetgen.CSharpPropertyInfo{
				Name: propName + "Value",
			}),
		}
	}

	if propertySpec.Ref == "pulumi.json#/Any" {
		if propertySpec.Description != "" {
			propertySpec.Description += "\n\n"
		}
		propertySpec.Description += fmt.Sprintf("Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `%s` for more information about the expected schema for this property.", ctx.cfTypeName)
	}

	if tagsProp, hasCustomTagsProp := GetTagsProperty(spec); propName == "Tags" || (hasCustomTagsProp && propName == tagsProp) {
		switch ctx.GetTagsStyle(typeSpec, spec) {
		case TagsStyleUntyped:
		case TagsStyleStringMap:
			// Nothing to do
		case TagsStyleKeyValueArray:
			// Swap referenced type to shared definition and remove custom type.
			oldRef := propertySpec.TypeSpec.Items.Ref
			propertySpec.TypeSpec.Items.Ref = "#/types/" + globalTagToken
			delete(ctx.pkg.Types, oldRef)
		default: // Unknown
			ctx.reports.UnexpectedTagsShapes[ctx.resourceToken] = spec
		}
	}

	return &propertySpec, nil
}

// propertyTypeSpec converts a JSON type definition to a Pulumi type definition.
func (ctx *context) propertyTypeSpec(parentName string, propSchema *jsschema.Schema) (*pschema.TypeSpec, error) {
	// References to other type definitions.
	if propSchema.Reference != "" {
		schemaName := strings.TrimPrefix(propSchema.Reference, "#/definitions/")
		typName := schemaName

		// For the fully qualified name, we prepend the resource name to the type name if it's not
		// already there. This avoids name collisions of types from different modules. However, we
		// don't do this if it would result in a name that already exists.
		// Preserve a previous special case, see https://github.com/pulumi/pulumi-aws-native/issues/644
		if parentName == "LoggingClusterLogging" && typName == "ClusterLogging" {
			typName = "ClusterLoggingEnabledTypes"
		} else if !strings.HasPrefix(schemaName, ctx.resourceName) {
			// Create a full type name by turning Foo into ResourceFoo.
			fullTypName := fmt.Sprintf("%s%s", ctx.resourceName, schemaName)
			if _, ok := ctx.resourceSpec.Definitions[fullTypName]; !ok {
				typName = fullTypName
			} else {
				fmt.Printf("Not expanding type name from %s to %s because a type of that name already exists\n",
					schemaName, fullTypName)
			}
		}

		typName = lowerAcronyms(typName)

		tok := fmt.Sprintf("%s:%s:%s", packageName, ctx.mod, typName)

		typeSchema, ok := ctx.resourceSpec.Definitions[schemaName]
		if !ok {
			return nil, errors.Errorf("definition %s not found in resource %s", schemaName, ctx.cfTypeName)
		}
		baseType := ""
		if len(typeSchema.Type) > 0 {
			baseType = typeSchema.Type[0].String()
		} else if len(typeSchema.Properties) > 0 || len(typeSchema.PatternProperties) > 0 {
			baseType = "object"
		}
		if baseType == "object" {
			if !ctx.visitedTypes.Has(tok) {
				ctx.visitedTypes.Add(tok)
				specs, requiredSpecs, irreversibleNames, err := ctx.genProperties(typName, typeSchema)
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
					Type:              "object",
					Properties:        specs,
					IrreversibleNames: irreversibleNames,
				}
			}

			referencedTypeName := fmt.Sprintf("#/types/%s", tok)
			return &pschema.TypeSpec{Ref: referencedTypeName}, nil
		} else {
			return ctx.propertyTypeSpec(typName, typeSchema)
		}
	}

	// Inline properties.
	if len(propSchema.Properties) > 0 {
		typName := parentName + "Properties"
		tok := fmt.Sprintf("%s:%s:%s", packageName, ctx.mod, typName)
		specs, requiredSpecs, irreversibleNames, err := ctx.genProperties(typName, propSchema)
		if err != nil {
			return nil, err
		}

		ctx.pkg.Types[tok] = pschema.ComplexTypeSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Description: propSchema.Description,
				Type:        "object",
				Properties:  specs,
				Required:    requiredSpecs.SortedValues(),
			},
		}
		ctx.metadata.Types[tok] = CloudAPIType{
			Type:              "object",
			Properties:        specs,
			IrreversibleNames: irreversibleNames,
		}
		referencedTypeName := fmt.Sprintf("#/types/%s", tok)
		return &pschema.TypeSpec{Ref: referencedTypeName}, nil
	}

	// Union types
	if len(propSchema.AnyOf) > 0 || len(propSchema.OneOf) > 0 {
		schemas := FlattenJSSchema(propSchema)
		types := make([]pschema.TypeSpec, 0, len(schemas))
		typeMap := make(map[string]bool, len(schemas))
		for i, sch := range schemas {
			typ, err := ctx.propertyTypeSpec(fmt.Sprintf("%s%d", parentName, i), sch)
			if err != nil {
				return nil, err
			}

			// stringify to create a key for identifying duplicates
			keybytes, err := json.Marshal(*typ)
			if err != nil {
				return nil, err
			}
			key := string(keybytes)

			if _, ok := typeMap[key]; !ok {
				types = append(types, *typ)
				typeMap[key] = true
			}
		}

		//simplify names when there is only one arm with a named type
		//TODO(https://github.com/pulumi/pulumi-aws-native/issues/994):
		// don't generate the complicated type names at all if we don't need them
		complexArms := []int{}
		for i, t := range types {
			if strings.HasSuffix(t.Ref, "Properties") || strings.HasSuffix(t.Ref, fmt.Sprintf("%s%d", parentName, i)) {
				complexArms = append(complexArms, i)
			}
		}
		if len(complexArms) == 1 {
			i := complexArms[0]
			typ, err := ctx.propertyTypeSpec(parentName, schemas[i])
			if err != nil {
				return nil, err
			}
			types[i] = *typ
		}

		if len(types) == 1 {
			return &types[0], nil
		} else {
			return &pschema.TypeSpec{OneOf: types}, nil
		}
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

	var mapType *jsschema.Schema
	if propSchema.AdditionalProperties != nil {
		mapType = propSchema.AdditionalProperties.Schema
	} else if propSchema.PatternProperties != nil && len(propSchema.PatternProperties) == 1 {
		// TODO: Capture pattern add add to documentation or even provider metadata for early validation feedback.
		// https://github.com/pulumi/pulumi-aws-native/issues/1346
		for _, schema := range propSchema.PatternProperties {
			mapType = schema
		}
	}

	if mapType != nil {
		valueType, err := ctx.propertyTypeSpec(parentName+"Value", mapType)
		if err != nil {
			return nil, err
		}
		return &pschema.TypeSpec{
			Type:                 "object",
			AdditionalProperties: valueType,
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
			elementType, err := ctx.propertyTypeSpec(parentName+"Item", propSchema.Items.Schemas[0])
			if err != nil {
				return nil, err
			}
			return &pschema.TypeSpec{
				Type:  "array",
				Items: elementType,
			}, nil
		default:
			return &pschema.TypeSpec{Ref: "pulumi.json#/Any"}, nil
		}
	}

	fmt.Printf("failed to generate property types for %+v\n", propSchema)
	return &pschema.TypeSpec{Ref: "pulumi.json#/Any"}, nil
}

func (ctx *context) genProperties(parentName string, typeSchema *jsschema.Schema) (map[string]pschema.PropertySpec, codegen.StringSet, map[string]string, error) {
	specs := map[string]pschema.PropertySpec{}
	requiredSpecs := codegen.NewStringSet()
	irreversibleNames := map[string]string{}
	for _, name := range codegen.SortedKeys(typeSchema.Properties) {
		value := typeSchema.Properties[name]
		sdkName := ToSdkName(name)

		typeSpec, err := ctx.propertyTypeSpec(parentName+lowerAcronyms(name), value)
		if err != nil {
			return nil, nil, nil, errors.Wrapf(err, "property %s", name)
		}
		propertySpec := pschema.PropertySpec{
			Description: value.Description,
			TypeSpec:    *typeSpec,
		}
		// TODO: temporary workaround to get the 0.1.0 out, let's find a better solution later.
		if name == "ClusterLogging" {
			propertySpec.Language = map[string]pschema.RawMessage{
				"csharp": rawMessage(dotnetgen.CSharpPropertyInfo{
					Name: "ClusterLoggingValue",
				}),
			}
		}
		specs[sdkName] = propertySpec

		if HasUppercaseAcronym(name) {
			irreversibleNames[sdkName] = name
		}
	}
	for _, name := range typeSchema.Required {
		sdkName := ToSdkName(name)
		if _, has := specs[sdkName]; has {
			requiredSpecs.Add(sdkName)
		}
	}
	return specs, requiredSpecs, irreversibleNames, nil
}

// genEnumType generates the enum type for a given schema.
func (ctx *context) genEnumType(enumName string, propSchema *jsschema.Schema) (*pschema.TypeSpec, error) {
	if len(propSchema.Type) == 0 {
		return nil, nil
	}
	if propSchema.Type[0] != jsschema.StringType {
		return nil, nil
	}
	typName := lowerAcronyms(enumName)
	if !strings.HasPrefix(enumName, ctx.resourceName) {
		typName = ctx.resourceName + enumName
	}
	switch ctx.mod + ":" + typName {
	case "networkfirewall:RuleGroupType":
		typName = "RuleGroupTypeEnum" // Go SDK name conflict vs. RuleGroup resource
	}
	tok := fmt.Sprintf("%s:%s:%s", packageName, ctx.mod, typName)

	enumSpec := &pschema.ComplexTypeSpec{
		Enum: []pschema.EnumValueSpec{},
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: propSchema.Description,
			Type:        "string",
		},
	}

	values := codegen.NewStringSet()
	for _, val := range propSchema.Enum {
		str := lowerAcronyms(ToUpperCamel(val.(string)))
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
			switch tok {
			case "aws-native:mediaconnect:FlowSourceProtocol":
				// FlowSourceProtocol is defined in two resources and one is a subset of the other.
				// They seem to be the same type. Pick the longer one here to avoid losing values.
				if len(enumSpec.Enum) < len(other.Enum) {
					enumSpec = &other
				}
			default:
				return nil, errors.Errorf("duplicate enum %q: %+v vs. %+v", tok, enumSpec.Enum, other.Enum)
			}
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
	contract.Assertf(len(resourceTypeComponents) == 3, "expected three parts in type %q", resourceTypeComponents)
	module := resourceTypeComponents[1]

	// Override the name of the Config module.
	if module == "Config" {
		module = "Configuration"
	}

	return lowerAcronyms(module)
}

func typeToken(typ string) (string, string) {
	resourceName := typeName(typ)
	module := strings.ToLower(moduleName(typ))

	return resourceName, fmt.Sprintf("%s:%s:%s", packageName, module, resourceName)
}

func getterToken(typ string) (string, string) {
	resourceName := typeName(typ)
	module := strings.ToLower(moduleName(typ))

	return resourceName, fmt.Sprintf("%s:%s:get%s", packageName, module, resourceName)
}

func typeName(typ string) string {
	resourceTypeComponents := strings.Split(typ, "::")
	contract.Assertf(len(resourceTypeComponents) == 3, "expected three parts in type %q", resourceTypeComponents)
	name := resourceTypeComponents[2]
	// Override name to avoid duplicate types due to "Output" suffix
	// See https://github.com/pulumi/pulumi/issues/8018
	if trimmed, hadSuffix := strings.CutSuffix(name, "Output"); hadSuffix {
		// Skip renaming existing FlowOutput type.
		if typ == "AWS::MediaConnect::FlowOutput" {
			return name
		}
		name = trimmed + "OutputResource"
	}
	return lowerAcronyms(name)
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

func GatherSemantics(schemaDir string) (SemanticsSpecDocument, error) {
	var semanticsDocument SemanticsSpecDocument
	semanticsPath := filepath.Join(schemaDir, "semantics.json")
	semanticsBytes, err := os.ReadFile(semanticsPath)
	if err != nil {
		return semanticsDocument, err
	}
	err = json.Unmarshal(semanticsBytes, &semanticsDocument)
	if err != nil {
		return semanticsDocument, err
	}
	return semanticsDocument, nil
}
