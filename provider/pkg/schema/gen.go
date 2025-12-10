// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	"encoding/json"
	goerrors "errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	jsschema "github.com/pulumi/jsschema"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/autonaming"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/default_tags"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/naming"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/refdb"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/resources"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/schema/docs"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	dotnetgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/maputil"
)

const packageName = "aws-native"
const globalTagToken = "aws-native:index:Tag"
const globalCreateOnlyTagToken = "aws-native:index:CreateOnlyTag"

// The documentation of some resources/properties is incomplete or incorrect in the AWS CloudFormation schema. The CloudFormation docs are more complete and accurate in these cases.
var forceDocumentationAugmentation = map[string]bool{
	// Documentation is truncated
	"AWS::EC2::Volume":      true,
	"AWS::EC2::VPCEndpoint": true,
	"AWS::ECR::Repository.EncryptionConfiguration.EncryptionType": true,
	"AWS::IAM::Role.Policies":                                     true,
	"AWS::IAM::Role.RoleName":                                     true,
}

type RegionInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GatherPackage builds a package spec based on the provided CF JSON schemas.
func GatherPackage(
	supportedResourceTypes []string,
	jsonSchemas []*jsschema.Schema,
	genAll bool,
	semanticsDocument *metadata.SemanticsSpecDocument,
	docsSchema *Docs,
	regionInfo []RegionInfo,
	refDB *refdb.RefDB,
) (*pschema.PackageSpec, *metadata.CloudAPIMetadata, *Reports, error) {
	globalTagType := pschema.ObjectTypeSpec{
		Type:        "object",
		Description: "A set of tags to apply to the resource.",
		Properties: map[string]pschema.PropertySpec{
			"key":   {TypeSpec: primitiveTypeSpec("String"), Description: "The key name of the tag"},
			"value": {TypeSpec: primitiveTypeSpec("String"), Description: "The value of the tag"},
		},
		Required: []string{"key", "value"},
	}
	globalCreateOnlyTagType := pschema.ObjectTypeSpec{
		Type:        "object",
		Description: "A set of tags to apply to the resource.",
		Properties: map[string]pschema.PropertySpec{
			"key":   {TypeSpec: primitiveTypeSpec("String"), Description: "The key name of the tag", ReplaceOnChanges: true},
			"value": {TypeSpec: primitiveTypeSpec("String"), Description: "The value of the tag", ReplaceOnChanges: true},
		},
		Required: []string{"key", "value"},
	}
	p := pschema.PackageSpec{
		Name:        packageName,
		Description: "A native Pulumi package for creating and managing Amazon Web Services (AWS) resources.",
		DisplayName: "AWS Cloud Control",
		Keywords: []string{
			"pulumi",
			"aws",
			"aws-native",
			"cloud control",
			"ccapi",
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
				"autoNaming": {
					Description: "The configuration for automatically naming resources.",
					TypeSpec:    pschema.TypeSpec{Ref: "#/types/aws-native:config:AutoNaming"},
				},
			},
			Required: []string{
				"region",
			},
		},
		Provider: pschema.ResourceSpec{
			ObjectTypeSpec: pschema.ObjectTypeSpec{
				Description: "The provider type for the AWS Cloud Control package. By default, resources use package-wide configuration settings, however an explicit `Provider` instance may be created and passed during resource construction to achieve fine-grained programmatic control over provider settings. See the [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.",
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
					"autoNaming": {
						Description: "The configuration for automatically naming resources.",
						TypeSpec:    pschema.TypeSpec{Ref: "#/types/aws-native:index:ProviderAutoNaming"},
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
				"autoNaming": {
					Description: "The configuration for automatically naming resources.",
					TypeSpec:    pschema.TypeSpec{Ref: "#/types/aws-native:index:ProviderAutoNaming"},
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
			globalCreateOnlyTagToken: {
				ObjectTypeSpec: globalCreateOnlyTagType,
			},
			resources.AutoNamingTypeToken: {
				ObjectTypeSpec: resources.AutoNamingTypeSpec(),
			},
		},
		Resources: map[string]pschema.ResourceSpec{
			metadata.ExtensionResourceToken: resources.ExtensionResourceSpec(),
			metadata.CfnCustomResourceToken: resources.CfnCustomResourceSpec(docs.CfnCustomResourceEmulatorDocs),
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
		"respectSchemaVersion": true,
		"importBasePath":       "github.com/pulumi/pulumi-aws-native/sdk/go/aws",
		"packageImportAliases": map[string]string{
			"github.com/pulumi/pulumi-aws-native/sdk/go/aws/aws-native": "aws",
		},
	})
	p.Language["nodejs"] = rawMessage(map[string]interface{}{
		"respectSchemaVersion": true,
	})
	p.Language["python"] = rawMessage(map[string]interface{}{
		"respectSchemaVersion": true,
		"pyproject": map[string]bool{
			"enabled": true,
		},
	})

	mdata := metadata.CloudAPIMetadata{
		Resources: map[string]metadata.CloudAPIResource{},
		Types: map[string]metadata.CloudAPIType{
			globalTagToken: {
				Type:       "object",
				Properties: globalTagType.Properties,
			},
			globalCreateOnlyTagToken: {
				Type:       "object",
				Properties: globalCreateOnlyTagType.Properties,
			},
		},
		Functions: map[string]metadata.CloudAPIFunction{},
	}

	reports := NewReports()

	supportedResources := codegen.NewStringSet(supportedResourceTypes...)

	var resourceCount int
	for _, jsonSchema := range jsonSchemas {
		cfTypeName := jsonSchema.Extras["typeName"].(string)
		isSupported := supportedResources.Has(cfTypeName)
		if isSupported || genAll {
			resourceName, resourceToken := typeToken(cfTypeName)
			fullMod := moduleName(cfTypeName)
			mod := strings.ToLower(fullMod)
			ctx := &cfSchemaContext{
				pkg:              &p,
				metadata:         &mdata,
				cfTypeName:       cfTypeName,
				mod:              mod,
				resourceName:     resourceName,
				resourceToken:    resourceToken,
				resourceSpec:     jsonSchema,
				visitedTypes:     codegen.NewStringSet(),
				docsVisitedTypes: codegen.NewStringSet(),
				isSupported:      isSupported,
				semantics:        semanticsDocument,
				reports:          reports,
				docs:             docsSchema,
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
	missingCount := 0
	for _, missing := range reports.MissingDocs {
		for _, val := range missing {
			if val == "0" {
				continue
			}
			missingCount += 1
		}
	}
	fmt.Printf("Number of docs augmented: %d\n", reports.DocsUpdated)
	fmt.Printf("Number of docs missing: %d\n", missingCount)
	fmt.Printf("Number of resource types: %d\n", resourceCount)

	// If there are types in the overlays that do not exist in the schema (e.g., enum types), add them.
	for tok, overlayType := range typeOverlays {
		if _, typeExists := p.Types[tok]; !typeExists {
			p.Types[tok] = overlayType
		}
	}

	// Generate the region enum.
	p.Types["aws-native:index/Region:Region"] = generateRegionEnum(regionInfo)

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
		"respectSchemaVersion": true,
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

	// Enrich with metadata about CF Ref Intrinsic function behavior.
	if refDB != nil {
		for resKey, res := range mdata.Resources {
			if r, ok := refDB.Resources[res.CfType]; ok {
				res.CfRef = &r.RefReturns.CfRefBehavior
				mdata.Resources[resKey] = res
			}
		}
	}

	return &p, &mdata, reports, nil
}

// cfSchemaContext holds shared information for a single CF JSON schema.
type cfSchemaContext struct {
	pkg              *pschema.PackageSpec
	metadata         *metadata.CloudAPIMetadata
	cfTypeName       string
	mod              string
	resourceName     string
	resourceToken    string
	resourceSpec     *jsschema.Schema
	visitedTypes     codegen.StringSet
	docsVisitedTypes codegen.StringSet
	isSupported      bool
	semantics        *metadata.SemanticsSpecDocument
	reports          *Reports
	docs             *Docs
}

func (ctx *cfSchemaContext) markCreateOnlyProperties(createOnlyProperties codegen.StringSet, resource *pschema.ResourceSpec) error {
	errs := []error{}
	for _, propPath := range createOnlyProperties.SortedValues() {
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

func (ctx *cfSchemaContext) markCreateOnlyProperty(propPath []string, property *pschema.PropertySpec) error {
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

func (ctx *cfSchemaContext) markCreateOnlyPropertyOnType(propName, typeRef string, remainingPath []string) error {
	ref := strings.TrimPrefix(typeRef, "#/types/")
	propType, ok := ctx.pkg.Types[ref]
	if !ok {
		return errors.New("Could not find referencedType: " + typeRef)
	}

	propertyName := naming.ToSdkName(propName)
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

func (ctx *cfSchemaContext) gatherInvoke() error {
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
		sdkName := naming.ToSdkName(n)
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
	tagsProp, hasTags := GetTagsProperty(ctx.resourceSpec)

	outputs := map[string]pschema.PropertySpec{}
	for k, v := range ctx.resourceSpec.Properties {
		sdkName := naming.ToSdkName(k)
		if nonOutputProperties.Has(k) {
			continue
		}
		p, err := ctx.propertySpec(k, resourceTypeName, v)
		if err != nil {
			return err
		}
		if hasTags && k == tagsProp {
			ctx.ApplyTagsTransformation(k, p, v)
		}
		addUntypedPropDocs(p, ctx.cfTypeName)
		outputs[sdkName] = *p
	}

	if len(outputs) == 0 {
		// If there are no outputs to be returned, (all are write/create properties), no need to emit the function.
		return nil
	}

	ctx.pkg.Functions[getterToken] = pschema.FunctionSpec{
		Description: naming.SanitizeCfnString(ctx.resourceSpec.Description),
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
		identifiers[i] = naming.ToSdkName(v)
	}
	ctx.metadata.Functions[getterToken] = metadata.CloudAPIFunction{
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
	if resourceSpec == nil {
		return codegen.NewStringSet()
	}
	output := codegen.NewStringSet()
	if p, ok := resourceSpec.Extras[listName]; ok {
		pSlice := p.([]interface{})
		for _, v := range pSlice {
			n := ResourceProperty(strings.TrimPrefix(v.(string), "/properties/"))
			output.Add(n.ToSdkName())
		}
	}
	return output
}

// gatherResourceType builds the schema for the resource type in the context.
func (ctx *cfSchemaContext) gatherResourceType() error {
	resourceTypeName := typeName(ctx.cfTypeName)

	writeOnlyProperties := readPropSdkNames(ctx.resourceSpec, "writeOnlyProperties")
	createOnlyProperties := readPropSdkNames(ctx.resourceSpec, "createOnlyProperties")
	readOnlyProperties := codegen.NewStringSet(readPropNames(ctx.resourceSpec, "readOnlyProperties")...)
	tagsProp, hasTags := GetTagsProperty(ctx.resourceSpec)
	requiredProperties := codegen.NewStringSet(ctx.resourceSpec.Required...)
	var tagsStyle default_tags.TagsStyle

	irreversibleNames := map[string]string{}
	inputProperties, requiredInputs := map[string]pschema.PropertySpec{}, codegen.NewStringSet()
	properties, required := map[string]pschema.PropertySpec{}, codegen.NewStringSet()
	props := maputil.SortedKeys(ctx.resourceSpec.Properties)
	for _, prop := range props {
		spec := ctx.resourceSpec.Properties[prop]
		sdkName := naming.ToSdkName(prop)
		originalSdkName := sdkName
		if sdkName == "id" {
			sdkName = "awsId"
		}

		// augment the documentation with the CloudFormation documentation
		// we do this before creating the propertySpec because in the process of
		// creating the propertySpec we change some of the names of properties which
		// makes it impossible to lookup in the docs
		ctx.augmentDocumentation(ctx.cfTypeName, prop, spec)

		propertySpec, err := ctx.propertySpec(prop, resourceTypeName, spec)
		if err != nil {
			return err
		}
		if hasTags && prop == tagsProp {
			tagsStyle = ctx.ApplyTagsTransformation(prop, propertySpec, spec)
		}
		addUntypedPropDocs(propertySpec, ctx.cfTypeName)
		properties[sdkName] = *propertySpec
		if requiredProperties.Has(prop) || readOnlyProperties.Has(prop) {
			required.Add(sdkName)
		}
		if createOnlyProperties.Has(originalSdkName) || !readOnlyProperties.Has(prop) {
			inputProperties[sdkName] = *propertySpec
			if requiredProperties.Has(prop) {
				requiredInputs.Add(sdkName)
			}
		}
		if naming.HasUppercaseAcronym(prop) || naming.ToCfnName(sdkName, nil) != prop {
			irreversibleNames[sdkName] = prop
		}

	}

	autoNamingSpec := autonaming.CreateAutoNamingSpec(inputProperties, resourceTypeName, ctx.resourceSpec.Properties, ctx.semantics.Resources[ctx.resourceToken])
	// If a field can be auto-named, its no longer required.
	if autoNamingSpec != nil {
		delete(requiredInputs, autoNamingSpec.SdkName)
	} else {
		ctx.reports.MissedAutonaming[ctx.resourceToken] = map[string]any{
			"cfTypeName": ctx.cfTypeName,
			"properties": inputProperties,
		}
	}

	ctx.updateDesc(ctx.cfTypeName, "", ctx.resourceSpec)
	var deprecationMessage string
	if !ctx.isSupported {
		deprecationMessage = fmt.Sprintf("%s is not yet supported by AWS Cloud Control, so its creation will currently fail. Please use the classic AWS provider, if possible.", resourceTypeName)
	}
	resourceSpec := pschema.ResourceSpec{
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: naming.SanitizeCfnString(ctx.resourceSpec.Description),
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

	ctx.metadata.Resources[ctx.resourceToken] = metadata.CloudAPIResource{
		CfType:            ctx.cfTypeName,
		Inputs:            inputProperties,
		Outputs:           properties,
		CreateOnly:        createOnlyProperties.SortedValues(),
		Required:          requiredInputs.SortedValues(),
		AutoNamingSpec:    autoNamingSpec,
		WriteOnly:         writeOnlyProperties.SortedValues(),
		ReadOnly:          readPropSdkNames(ctx.resourceSpec, "readOnlyProperties").SortedValues(),
		IrreversibleNames: irreversibleNames,
		TagsProperty:      naming.ToSdkName(tagsProp),
		TagsStyle:         tagsStyle,
		PrimaryIdentifier: ctx.gatherResourcePrimaryIdentifier(),
		ListHandlerSchema: ctx.gatherListHandlerSchema(),
	}
	return nil
}

func (ctx *cfSchemaContext) gatherResourcePrimaryIdentifier() []string {
	primaryIdentifier := readPropNames(ctx.resourceSpec, "primaryIdentifier")
	identifiers := make([]string, len(primaryIdentifier))
	for i, v := range primaryIdentifier {
		identifiers[i] = naming.ToSdkName(v)
	}
	return identifiers
}

func (ctx *cfSchemaContext) gatherListHandlerSchema() *metadata.ListHandlerSchema {
	handlersAny, ok := ctx.resourceSpec.Extras["handlers"]
	if !ok {
		return nil
	}
	handlers, ok := handlersAny.(map[string]interface{})
	if !ok {
		return nil
	}

	listAny, ok := handlers["list"]
	if !ok {
		return nil
	}
	list, ok := listAny.(map[string]interface{})
	if !ok {
		return nil
	}

	handlerSchemaAny, ok := list["handlerSchema"]
	if !ok {
		return nil
	}
	handlerSchema, ok := handlerSchemaAny.(map[string]interface{})
	if !ok {
		return nil
	}

	result := &metadata.ListHandlerSchema{}

	if reqAny, ok := handlerSchema["required"]; ok {
		if reqSlice, ok := reqAny.([]interface{}); ok {
			for _, v := range reqSlice {
				if s, ok := v.(string); ok {
					result.Required = append(result.Required, s)
				}
			}
		}
	}

	if propsAny, ok := handlerSchema["properties"]; ok {
		if propsMap, ok := propsAny.(map[string]interface{}); ok {
			result.Properties = map[string]metadata.ListHandlerProperty{}
			for propName, propVal := range propsMap {
				propMap, ok := propVal.(map[string]interface{})
				if !ok {
					continue
				}

				var prop metadata.ListHandlerProperty
				if desc, ok := propMap["description"].(string); ok {
					prop.Description = desc
				}
				if typ, ok := propMap["type"].(string); ok {
					prop.Type = typ
				}

				if prop.Description == "" && prop.Type == "" {
					if ref, ok := propMap["$ref"].(string); ok {
						refName := parsePropertyNameFromRef(ref)
						if refName != "" {
							if propSchema, ok := ctx.resourceSpec.Properties[refName]; ok && propSchema != nil {
								if prop.Description == "" && propSchema.Description != "" {
									prop.Description = naming.SanitizeCfnString(propSchema.Description)
								}
								if prop.Type == "" {
									prop.Type = firstPrimitiveType(propSchema.Type)
								}
							}
						}
					}
				}

				// Only include properties when at least one field is present.
				if prop.Description == "" && prop.Type == "" {
					continue
				}

				result.Properties[propName] = prop
			}
			if len(result.Properties) == 0 {
				result.Properties = nil
			}
		}
	}

	if len(result.Required) == 0 && len(result.Properties) == 0 {
		return nil
	}

	return result
}

func parsePropertyNameFromRef(ref string) string {
	const propToken = "/properties/"
	idx := strings.Index(ref, propToken)
	if idx == -1 {
		return ""
	}
	return ref[idx+len(propToken):]
}

func firstPrimitiveType(types jsschema.PrimitiveTypes) string {
	for _, t := range types {
		switch t {
		case jsschema.StringType:
			return "string"
		case jsschema.NumberType:
			return "number"
		case jsschema.IntegerType:
			return "integer"
		case jsschema.BooleanType:
			return "boolean"
		case jsschema.ArrayType:
			return "array"
		case jsschema.ObjectType:
			return "object"
		case jsschema.NullType:
			return "null"
		}
	}
	return ""
}

func addUntypedPropDocs(propertySpec *pschema.PropertySpec, cfTypeName string) {
	if propertySpec.Ref == "pulumi.json#/Any" {
		if propertySpec.Description != "" {
			propertySpec.Description += "\n\n"
		}
		propertySpec.Description += fmt.Sprintf("Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `%s` for more information about the expected schema for this property.", cfTypeName)
	}
}

// reportMissingDocs add missing docs to the report
func (ctx *cfSchemaContext) reportMissingDocs(desc, propName string) {
	if desc == "" {
		val, ok := ctx.reports.MissingDocs[ctx.cfTypeName]
		if ok {
			if _, ok := val[propName]; !ok {
				val[propName] = "1"
			}
		} else {
			ctx.reports.MissingDocs[ctx.cfTypeName] = map[string]string{
				propName: "1",
			}
		}
	}
}

func (ctx *cfSchemaContext) propertySpec(propName, resourceTypeName string, spec *jsschema.Schema) (*pschema.PropertySpec, error) {
	typeSpec, err := ctx.propertyTypeSpec(naming.LowerAcronyms(propName), spec)
	if err != nil {
		return nil, err
	}

	ctx.reportMissingDocs(spec.Description, propName)

	propertySpec := pschema.PropertySpec{
		TypeSpec:    *typeSpec,
		Description: naming.SanitizeCfnString(spec.Description),
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

	return &propertySpec, nil
}

// updateDesc updates the schema with a description from the CloudFormation docs if one
// is not already provided or incomplete.
func (ctx *cfSchemaContext) updateDesc(refName, propName string, spec *jsschema.Schema) {
	fullyQualifiedPropertyName := GetFullyQualifiedPropertyName(refName, propName)
	if spec.Description == "" {
		desc, found := ctx.docs.GetPropertyDesc(
			refName,
			propName,
		)
		spec.Description = desc
		// if it has been updated then increment the count
		if spec.Description != "" {
			ctx.reports.DocsUpdated += 1
		}

		if !found {
			if val, ok := ctx.reports.MissingDocs[ctx.cfTypeName]; ok {
				val[propName] = "0"
			} else {
				ctx.reports.MissingDocs[ctx.cfTypeName] = map[string]string{
					propName: "0",
				}
			}
		}
	} else if _, ok := forceDocumentationAugmentation[fullyQualifiedPropertyName]; ok {
		if desc, found := ctx.docs.GetPropertyDesc(refName, propName); found {
			spec.Description = desc
			ctx.reports.DocsUpdated += 1
		} else {
			ctx.reportMissingDocs(desc, propName)
			fmt.Printf("Description augmentation configured for type %s but CloudFormation Docs do not include description\n", fullyQualifiedPropertyName)
		}
	}
}

// augmentDocumentation will augment the CloudFormation JSON schema with documentation from the
// CloudFormation documentation. This will recurse through each schema property and attempt to find
// the corresponding entry in the documentation schema.
func (ctx *cfSchemaContext) augmentDocumentation(referenceName, propName string, spec *jsschema.Schema) {
	// some types are self-referencing so we need to only process them the first time
	visitedType := fmt.Sprintf("%s.%s", referenceName, propName)
	if ctx.docsVisitedTypes.Has(visitedType) {
		return
	}
	ctx.docsVisitedTypes.Add(visitedType)

	// references are handled as a special case because the TypeName of a property is derived
	// from it's reference _not_ from its name, for example
	//    {
	//      "typeName": "AWS::Pipes::Pipe",
	//      "properties": {
	//        "AccessEndpoints": {
	//          "$ref": "#/definitions/AccessEndpoint"
	//        }
	//      }
	//    }
	//
	// The name of this property in the docs schema will be `AWS::Pipes::Pipe.AccessEndpoint`
	if spec.Reference != "" {
		schemaName := strings.TrimPrefix(spec.Reference, "#/definitions/")
		typeSchema, ok := ctx.resourceSpec.Definitions[schemaName]
		if ok {
			refName := fmt.Sprintf("%s.%s", ctx.cfTypeName, schemaName)

			if len(typeSchema.Type) == 1 {
				if typeSchema.PatternProperties != nil && len(typeSchema.PatternProperties) == 1 {
					for _, schema := range typeSchema.PatternProperties {
						ctx.augmentDocumentation(refName, schemaName, schema)
					}
				}

				if typeSchema.Type.Contains(jsschema.ObjectType) {
					for _, name := range maputil.SortedKeys(typeSchema.Properties) {
						value := typeSchema.Properties[name]
						ctx.augmentDocumentation(refName, name, value)
					}
				}
			}
		}
	}

	if len(spec.Properties) > 0 {
		for _, name := range maputil.SortedKeys(spec.Properties) {
			value := spec.Properties[name]
			refName := referenceName
			if refName == ctx.cfTypeName {
				refName = fmt.Sprintf("%s.%s", ctx.cfTypeName, propName)
			}
			ctx.augmentDocumentation(refName, name, value)
		}
	}

	if len(spec.AnyOf) > 0 || len(spec.OneOf) > 0 {
		schemas := FlattenJSSchema(spec)
		for _, sch := range schemas {
			ctx.augmentDocumentation(referenceName, propName, sch)
		}
	}

	if spec.PatternProperties != nil && len(spec.PatternProperties) == 1 {
		for _, schema := range spec.PatternProperties {
			ctx.augmentDocumentation(referenceName, propName, schema)
		}
	}

	if spec.Items != nil && len(spec.Items.Schemas) == 1 {
		arraySpec := spec.Items.Schemas[0]
		if arraySpec.Reference != "" {
			schemaName := strings.TrimPrefix(arraySpec.Reference, "#/definitions/")
			refName := fmt.Sprintf("%s.%s", ctx.cfTypeName, schemaName)
			ctx.augmentDocumentation(refName, schemaName, arraySpec)
			if spec.Description == "" {
				spec.Description = arraySpec.Description
			}
		}
	}

	ctx.updateDesc(referenceName, propName, spec)
}

// propertyTypeSpec converts a JSON type definition to a Pulumi type definition.
func (ctx *cfSchemaContext) propertyTypeSpec(parentName string, propSchema *jsschema.Schema) (*pschema.TypeSpec, error) {
	propSchema = NormaliseTypes(propSchema)

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

		typName = naming.LowerAcronyms(typName)

		tok := fmt.Sprintf("%s:%s:%s", packageName, ctx.mod, typName)

		typeSchema, ok := ctx.resourceSpec.Definitions[schemaName]
		if !ok {
			fmt.Printf("definition %s not found in resource %s\n", schemaName, ctx.cfTypeName)
			return &pschema.TypeSpec{Ref: "pulumi.json#/Any"}, nil
		}
		var mapType *jsschema.Schema
		if typeSchema.AdditionalProperties != nil && typeSchema.AdditionalProperties.Schema != nil {
			mapType = typeSchema.AdditionalProperties.Schema
		} else if typeSchema.PatternProperties != nil && len(typeSchema.PatternProperties) == 1 {
			// TODO: Capture pattern add add to documentation or even provider metadata for early validation feedback.
			// https://github.com/pulumi/pulumi-aws-native/issues/1346
			for _, schema := range typeSchema.PatternProperties {
				mapType = schema
			}
		}

		if mapType != nil {
			valueType, err := ctx.propertyTypeSpec(parentName+"Value", mapType)
			if err != nil {
				return nil, err
			}

			// Where a map type has a value which is either a map or a list, replace the value with "Any" to appease Go codegen
			// TODO: remove once Go codegen issue is solved: https://github.com/pulumi/pulumi/issues/15478
			if valueType.Items != nil || valueType.AdditionalProperties != nil {
				valueType = &pschema.TypeSpec{
					Ref: "pulumi.json#/Any",
				}
			}
			return &pschema.TypeSpec{
				Type:                 "object",
				AdditionalProperties: valueType,
			}, nil
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
						Description: naming.SanitizeCfnString(typeSchema.Description),
						Type:        "object",
						Properties:  specs,
						Required:    requiredSpecs.SortedValues(),
					},
				}
				ctx.metadata.Types[tok] = metadata.CloudAPIType{
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
				Description: naming.SanitizeCfnString(propSchema.Description),
				Type:        "object",
				Properties:  specs,
				Required:    requiredSpecs.SortedValues(),
			},
		}
		ctx.metadata.Types[tok] = metadata.CloudAPIType{
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
	if propSchema.AdditionalProperties != nil && propSchema.AdditionalProperties.Schema != nil {
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

	if propSchema.Type.Contains(jsschema.ArrayType) {
		elementType, err := ctx.propertyTypeSpec(parentName+"Item", propSchema.Items.Schemas[0])
		if err != nil {
			return nil, err
		}
		return &pschema.TypeSpec{
			Type:  "array",
			Items: elementType,
		}, nil
	}

	// All other types.
	if len(propSchema.Type) == 1 {
		typeSpec := parseJsonType(propSchema.Type[0])
		return &typeSpec, nil
	}

	if len(propSchema.Type) > 1 {
		// Avoid creating a union containing an "Any" type - as it can always collapse back to "Any".
		if propSchema.Type.Contains(jsschema.ObjectType) || propSchema.Type.Contains(jsschema.UnspecifiedType) {
			return &pschema.TypeSpec{Ref: "pulumi.json#/Any"}, nil
		}
		specs := make([]pschema.TypeSpec, len(propSchema.Type))
		for i, t := range propSchema.Type {
			specs[i] = parseJsonType(t)
		}
		return &pschema.TypeSpec{OneOf: specs}, nil
	}

	fmt.Printf("failed to generate property types for %+v\n", propSchema)
	return &pschema.TypeSpec{Ref: "pulumi.json#/Any"}, nil
}

// parseJsonType converts a JSON type with no additional information to a Pulumi type.
func parseJsonType(t jsschema.PrimitiveType) pschema.TypeSpec {
	switch t {
	case jsschema.StringType:
		return pschema.TypeSpec{Type: "string"}
	case jsschema.IntegerType:
		return pschema.TypeSpec{Type: "integer"}
	case jsschema.NumberType:
		return pschema.TypeSpec{Type: "number"}
	case jsschema.BooleanType:
		return pschema.TypeSpec{Type: "boolean"}
	case jsschema.ArrayType:
		return pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Ref: "pulumi.json#/Any"}}
	case jsschema.ObjectType:
		return pschema.TypeSpec{Ref: "pulumi.json#/Any"}
	default:
		return pschema.TypeSpec{Ref: "pulumi.json#/Any"}
	}
}

func (ctx *cfSchemaContext) genProperties(parentName string, typeSchema *jsschema.Schema) (map[string]pschema.PropertySpec, codegen.StringSet, map[string]string, error) {
	specs := map[string]pschema.PropertySpec{}
	requiredSpecs := codegen.NewStringSet()
	irreversibleNames := map[string]string{}
	for _, name := range maputil.SortedKeys(typeSchema.Properties) {
		value := typeSchema.Properties[name]
		sdkName := naming.ToSdkName(name)

		typeSpec, err := ctx.propertyTypeSpec(parentName+naming.LowerAcronyms(name), value)
		if err != nil {
			return nil, nil, nil, errors.Wrapf(err, "property %s", name)
		}
		ctx.reportMissingDocs(value.Description, name)
		propertySpec := pschema.PropertySpec{
			Description: naming.SanitizeCfnString(value.Description),
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

		if naming.HasUppercaseAcronym(name) {
			irreversibleNames[sdkName] = name
		}
	}
	for _, name := range typeSchema.Required {
		sdkName := naming.ToSdkName(name)
		if _, has := specs[sdkName]; has {
			requiredSpecs.Add(sdkName)
		}
	}
	return specs, requiredSpecs, irreversibleNames, nil
}

// genEnumType generates the enum type for a given schema.
func (ctx *cfSchemaContext) genEnumType(enumName string, propSchema *jsschema.Schema) (*pschema.TypeSpec, error) {
	if len(propSchema.Type) == 0 {
		return nil, nil
	}
	if propSchema.Type[0] != jsschema.StringType {
		return nil, nil
	}
	typName := naming.LowerAcronyms(enumName)
	if !strings.HasPrefix(enumName, ctx.resourceName) {
		typName = ctx.resourceName + enumName
	}
	switch ctx.mod + ":" + typName {
	case "networkfirewall:RuleGroupType":
		typName = "RuleGroupTypeEnum" // Go SDK name conflict vs. RuleGroup resource
	}

	switch ctx.cfTypeName + ":" + enumName {
	case "AWS::QBusiness::Index:Status":
		typName = "QBusinessIndexStatus"
	case "AWS::ARCZonalShift::AutoshiftObserverNotificationStatus:AutoshiftObserverNotificationStatus":
		typName = "AutoshiftObserverNotificationStatusEnum" // enumName matches typName which causes a conflict
	}

	tok := fmt.Sprintf("%s:%s:%s", packageName, ctx.mod, typName)

	// AWS CloudFormation schema incorrectly omits REPLICATING from the enum,
	// even though the description mentions it and AWS API returns it during replication.
	// This is a workaround for the upstream schema bug. Match on the fully qualified
	// type token to ensure we only modify this specific enum.
	if tok == "aws-native:efs:FileSystemProtectionReplicationOverwriteProtection" {
		hasReplicating := false
		for _, val := range propSchema.Enum {
			if val == "REPLICATING" {
				hasReplicating = true
				break
			}
		}
		if !hasReplicating {
			propSchema.Enum = append(propSchema.Enum, "REPLICATING")
		}
	}

	enumSpec := &pschema.ComplexTypeSpec{
		Enum: []pschema.EnumValueSpec{},
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: naming.SanitizeCfnString(propSchema.Description),
			Type:        "string",
		},
	}

	values := codegen.NewStringSet()
	for _, val := range propSchema.Enum {
		str := naming.LowerAcronyms(naming.ToUpperCamel(val.(string)))
		if values.Has(str) {
			continue
		}

		// Special case for when there is an enum which has an empty string value
		// TODO[pulumi/aws-native#2121]
		if (typName == "ChannelPreset" || typName == "SoftwarePackageVersionSbomValidationStatus") && str == "" {
			str = "Empty"
		}

		// Enum values cannot end in `*Input` or `*Output`. Special casing these two instances
		if (typName == "FlowNodeType" || typName == "FlowVersionFlowNodeType") && (str == "Input" || str == "Output") {
			str = str + "Type"
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
	ctx.metadata.Types[tok] = metadata.CloudAPIType{Type: enumSpec.Type}

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

	return naming.LowerAcronyms(module)
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
	return naming.LowerAcronyms(name)
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

func GatherSemantics(schemaDir string) (metadata.SemanticsSpecDocument, error) {
	var semanticsDocument metadata.SemanticsSpecDocument
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

func GetFullyQualifiedPropertyName(refName, propName string) string {
	if propName == "" {
		return refName
	}
	return fmt.Sprintf("%s.%s", refName, propName)
}
