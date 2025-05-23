// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schema

import (
	"sort"
	"strings"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/naming"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func configToProvider(config pschema.ComplexTypeSpec) pschema.ComplexTypeSpec {
	// The Provider schema is the Config schema with an additional Language block for each Property.
	for k, v := range config.Properties {
		v.Language = map[string]pschema.RawMessage{
			"python": rawMessage(map[string]interface{}{
				"mapCase": false,
			}),
		}
		config.Properties[k] = v
	}
	return config
}

var assumeRole = pschema.ComplexTypeSpec{
	ObjectTypeSpec: pschema.ObjectTypeSpec{
		Description: "The configuration for a Provider to assume a role.",
		Properties: map[string]pschema.PropertySpec{
			"durationSeconds": {
				Description: "Number of seconds to restrict the assume role session duration.",
				TypeSpec:    pschema.TypeSpec{Type: "integer"},
			},
			"externalId": {
				Description: "External identifier to use when assuming the role.",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"policy": {
				Description: "IAM Policy JSON describing further restricting permissions for the IAM Role being assumed.",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"policyArns": {
				Description: "Set of Amazon Resource Names (ARNs) of IAM Policies describing further restricting permissions for the role.",
				TypeSpec:    pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Type: "string"}},
			},
			"roleArn": {
				Description: "Amazon Resource Name (ARN) of the IAM Role to assume.",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"sessionName": {
				Description: "Session name to use when assuming the role.",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"tags": {
				Description: "Map of assume role session tags.",
				TypeSpec: pschema.TypeSpec{
					Type:                 "object",
					AdditionalProperties: &pschema.TypeSpec{Type: "string"}},
			},
			"transitiveTagKeys": {
				Description: "A list of keys for session tags that you want to set as transitive. If you set a tag key as transitive, the corresponding key and value passes to subsequent sessions in a role chain.",
				TypeSpec:    pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Type: "string"}},
			},
		},
		Type: "object",
	},
}

var defaultTags = pschema.ComplexTypeSpec{
	ObjectTypeSpec: pschema.ObjectTypeSpec{
		Description: "The configuration with resource tag settings to apply across all resources handled by this provider. This is designed to replace redundant per-resource `tags` configurations. Provider tags can be overridden with new values, but not excluded from specific resources. To override provider tag values, use the `tags` argument within a resource to configure new tag values for matching keys.",
		Properties: map[string]pschema.PropertySpec{
			"tags": {
				Description: "A group of tags to set across all resources.",
				TypeSpec: pschema.TypeSpec{
					Type:                 "object",
					AdditionalProperties: &pschema.TypeSpec{Type: "string"}},
			},
		},
		Type: "object",
	},
}

var endpoints = pschema.ComplexTypeSpec{
	ObjectTypeSpec: pschema.ObjectTypeSpec{
		Description: "The configuration for for customizing service endpoints.",
		Properties: map[string]pschema.PropertySpec{
			"cloudcontrol": {
				Description: "Override the default endpoint for AWS CloudControl",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"cloudformation": {
				Description: "Override the default endpoint for AWS CloudFormation",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"ec2": {
				Description: "Override the default endpoint for AWS Elastic Compute Cloud (EC2)",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"ssm": {
				Description: "Override the default endpoint for AWS Systems Manager",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"sts": {
				Description: "Override the default endpoint for AWS Security Token Service (STS)",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
		},
		Type: "object",
	},
}

var ignoreTags = pschema.ComplexTypeSpec{
	ObjectTypeSpec: pschema.ObjectTypeSpec{
		Description: "The configuration with resource tag settings to ignore across all resources handled by this provider (except any individual service tag resources such as `ec2.Tag`) for situations where external systems are managing certain resource tags.",
		Properties: map[string]pschema.PropertySpec{
			"keyPrefixes": {
				Description: "List of exact resource tag keys to ignore across all resources handled by this provider. This configuration prevents Pulumi from returning the tag in any `tags` attributes and displaying any configuration difference for the tag value. If any resource configuration still has this tag key configured in the `tags` argument, it will display a perpetual difference until the tag is removed from the argument or `ignoreChanges` is also used.",
				TypeSpec:    pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Type: "string"}},
			},
			"keys": {
				Description: "List of resource tag key prefixes to ignore across all resources handled by this provider. This configuration prevents Pulumi from returning any tag key matching the prefixes in any `tags` attributes and displaying any configuration difference for those tag values. If any resource configuration still has a tag matching one of the prefixes configured in the `tags` argument, it will display a perpetual difference until the tag is removed from the argument or `ignoreChanges` is also used.",
				TypeSpec:    pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Type: "string"}},
			},
		},
		Type: "object",
	},
}

var autoName = pschema.ComplexTypeSpec{
	ObjectTypeSpec: pschema.ObjectTypeSpec{
		Description: "The configuration for automatically naming resources.",
		Properties: map[string]pschema.PropertySpec{
			"autoTrim": {
				Description: "Automatically trim the auto-generated name to meet the maximum length constraint.",
				TypeSpec:    pschema.TypeSpec{Type: "boolean"},
			},
			"randomSuffixMinLength": {
				Description: "The minimum length of the random suffix to append to the auto-generated name.",
				Default:     1,
				TypeSpec:    pschema.TypeSpec{Type: "integer"},
			},
		},
		Type: "object",
	},
}

func generateRegionEnum(regions []RegionInfo) pschema.ComplexTypeSpec {
	enums := make([]pschema.EnumValueSpec, len(regions))
	for i, region := range regions {
		sections := strings.Split(region.Name, "-")
		for i, section := range sections {
			sections[i] = naming.ToPascalCase(section)
		}
		capitalizedName := strings.Join(sections, "")

		enums[i] = pschema.EnumValueSpec{
			Name:        capitalizedName,
			Value:       region.Name,
			Description: region.Description,
		}
	}

	sort.SliceStable(enums, func(i, j int) bool {
		return enums[i].Name < enums[j].Name
	})

	return pschema.ComplexTypeSpec{
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: "A Region represents any valid Amazon region that may be targeted with deployments.",
			Type:        "string",
		},
		Enum: enums,
	}
}

// typeOverlays augment the types defined by the schema.
var typeOverlays = map[string]pschema.ComplexTypeSpec{
	"aws-native:config:AssumeRole":         assumeRole,
	"aws-native:index:ProviderAssumeRole":  configToProvider(assumeRole),
	"aws-native:config:DefaultTags":        defaultTags,
	"aws-native:index:ProviderDefaultTags": configToProvider(defaultTags),
	"aws-native:config:Endpoints":          endpoints,
	"aws-native:index:ProviderEndpoint":    configToProvider(endpoints),
	"aws-native:config:IgnoreTags":         ignoreTags,
	"aws-native:index:ProviderIgnoreTags":  configToProvider(ignoreTags),
	"aws-native:config:AutoNaming":         autoName,
	"aws-native:index:ProviderAutoNaming":  configToProvider(autoName),
}
