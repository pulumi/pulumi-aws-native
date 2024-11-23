// Copyright 2016-2021, Pulumi Corporation.

package metadata

import (
	"github.com/pulumi/pulumi-aws-native/provider/pkg/default_tags"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

// CloudAPIMetadata is a collection of all resources and types in the AWS Cloud Control API.
type CloudAPIMetadata struct {
	Resources map[string]CloudAPIResource `json:"resources"`
	Types     map[string]CloudAPIType     `json:"types"`
	Functions map[string]CloudAPIFunction `json:"functions"`
}

// CloudAPIResource contains metadata for a single AWS Resource.
type CloudAPIResource struct {
	CfType            string                          `json:"cf"`
	Inputs            map[string]pschema.PropertySpec `json:"inputs"`
	Outputs           map[string]pschema.PropertySpec `json:"outputs"`
	AutoNamingSpec    *AutoNamingSpec                 `json:"autoNamingSpec,omitempty"`
	Required          []string                        `json:"required,omitempty"`
	CreateOnly        []string                        `json:"createOnly,omitempty"`
	WriteOnly         []string                        `json:"writeOnly,omitempty"`
	IrreversibleNames map[string]string               `json:"irreversibleNames,omitempty"`
	TagsProperty      string                          `json:"tagsProperty,omitempty"`
	TagsStyle         default_tags.TagsStyle          `json:"tagsStyle,omitempty"`

	// Describes the behavior of the CF Ref intrinsic for this resource.
	CfRef *CfRefBehavior `json:"cfRef,omitempty"`
}

type AutoNamingSpec struct {
	SdkName    string            `json:"sdkName"`
	MinLength  int               `json:"minLength,omitempty"`
	MaxLength  int               `json:"maxLength,omitempty"`
	TriviaSpec *NamingTriviaSpec `json:"namingTriviaSpec,omitempty"`
}

// CloudAPIType contains metadata for an auxiliary type.
type CloudAPIType struct {
	Type              string                          `json:"type"`
	Properties        map[string]pschema.PropertySpec `json:"properties,omitempty"`
	IrreversibleNames map[string]string               `json:"irreversibleNames,omitempty"`
}

type CloudAPIFunction struct {
	CfType      string   `json:"cf"`
	Identifiers []string `json:"ids"`
}

// ExtensionResourceToken is a Pulumi token for the resource to deploy
// custom third-party CloudFormation types.
const ExtensionResourceToken = "aws-native:index:ExtensionResource"

// CfnCustomResourceToken is a Pulumi token for the resource to deploy
// CloudFormation custom resources.
const CfnCustomResourceToken = "aws-native:cloudformation:CustomResourceEmulator"

// Describes the behavior of CloudFormation Ref intrinsic for a given resource.
//
// One and only one of [Property], [Properties], [NotSupported], [NotSupportedYet] must be set.
type CfRefBehavior struct {
	// If set, indicates that Ref will return the value of the given Resource property directly.
	Property string `json:"property,omitempty"`

	// If set, indicates that Ref will return a string value obtained by joining several Resource properties with a
	// delimiter, typically "|".
	//
	// Usually these properties are strings, but they can also be objects, in which case their values are
	// JSON-encoded. See AWS::LakeFormation::PrincipalPermissions for an example of this.
	Properties []string `json:"properties,omitempty"`

	// Delimiter, typically "|". See [Properties].
	Delimiter string `json:"delimiter,omitempty"`

	// If set, Ref is not supported for this resource in CF.
	NotSupported bool `json:"notSupported,omitempty"`

	// If set, Ref is supported in CF but this metadata is not yet available in the Pulumi provider but might be
	// added in a later version.
	NotSupportedYet bool `json:"notSupportedYet,omitempty"`
}
