// Copyright 2016-2021, Pulumi Corporation.

package schema

import pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"

// CloudAPIMetadata is a collection of all resources and types in the AWS Cloud Control API.
type CloudAPIMetadata struct {
	Resources map[string]CloudAPIResource `json:"resources"`
	Types     map[string]CloudAPIType     `json:"types"`
}

// CloudAPIResource contains metadata for a single AWS Resource.
type CloudAPIResource struct {
	CfType          string                          `json:"cf"`
	Inputs          map[string]pschema.PropertySpec `json:"inputs"`
	Outputs         map[string]pschema.PropertySpec `json:"outputs"`
	AutoNamingSpec  *AutoNamingSpec                 `json:"autoNamingSpec,omitempty"`
	Required        []string                        `json:"required,omitempty"`
	CreateOnly      []string                        `json:"createOnly,omitempty"`
}

type AutoNamingSpec struct {
	SdkName   string `json:"sdkName"`
	MinLength int    `json:"minLength,omitempty"`
	MaxLength int    `json:"maxLength,omitempty"`
}

// CloudAPIType contains metadata for an auxiliary type.
type CloudAPIType struct {
	Type       string                          `json:"type"`
	Properties map[string]pschema.PropertySpec `json:"properties,omitempty"`
}

// ExtensionResourceToken is a Pulumi token for the resource to deploy
// custom third-party CloudFormation types.
const ExtensionResourceToken = "aws-native:index:ExtensionResource"
