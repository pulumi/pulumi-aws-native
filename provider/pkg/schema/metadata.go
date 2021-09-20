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
	CfType     string                          `json:"cf"`
	Inputs     map[string]pschema.PropertySpec `json:"inputs"`
	Outputs    map[string]pschema.PropertySpec `json:"outputs"`
	Required   []string                        `json:"required,omitempty"`
	CreateOnly []string                        `json:"createOnly,omitempty"`
}

// CloudAPIType contains metadata for an auxiliary type.
type CloudAPIType struct {
	Properties map[string]pschema.PropertySpec `json:"properties"`
}
