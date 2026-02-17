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
	CfType         string                          `json:"cf"`
	Inputs         map[string]pschema.PropertySpec `json:"inputs"`
	Outputs        map[string]pschema.PropertySpec `json:"outputs"`
	AutoNamingSpec *AutoNamingSpec                 `json:"autoNamingSpec,omitempty"`
	Required       []string                        `json:"required,omitempty"`
	CreateOnly     []string                        `json:"createOnly,omitempty"`
	// ReadOnly properties are properties that only exist as output values
	// The properties can be top level properties or can be nested within object or array
	// properties.
	// '/' is used to denote a nested object property
	// '/*/' is used to denote a nested array property
	// e.g.
	// - ReadOnly: [
	//     'arn', // top level
	//     'someObjectProp/arn', // the arn value of someObjectProp is an output (but not the other properties of someObjectProp)
	//     'someArrayProp/*/someObjectProp/arn'
	//   ]
	//
	// NOTE: The values in this property will have been converted from CFN casing to pulumi sdk casing
	ReadOnly          []string               `json:"readOnly,omitempty"`
	WriteOnly         []string               `json:"writeOnly,omitempty"`
	IrreversibleNames map[string]string      `json:"irreversibleNames,omitempty"`
	TagsProperty      string                 `json:"tagsProperty,omitempty"`
	TagsStyle         default_tags.TagsStyle `json:"tagsStyle,omitempty"`

	// Describes the behavior of the CF Ref intrinsic for this resource.
	CfRef *CfRefBehavior `json:"cfRef,omitempty"`

	// PrimaryIdentifier is a list of Pulumi property names that together uniquely identify a resource.
	//
	// If more than one property is given, the values may be joined with the "|" character to form a specific
	// resource identifier value suitable for use with the Cloud Control API.
	//
	// See also https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-identifier.html
	PrimaryIdentifier []string `json:"primaryIdentifier,omitempty"`

	// ListHandlerSchema contains a minimal subset of the CloudFormation list handler schema for a resource.
	ListHandlerSchema *ListHandlerSchema `json:"listHandlerSchema,omitempty"`

	// PropertyTransforms maps SDK property paths to JSONata expressions for drift detection.
	// CloudFormation schemas include these transforms to define how property values should be
	// normalized during comparison (e.g., case normalization, numeric to string mappings).
	//
	// Paths use lowerCamelCase with "/" separators and "*" for array elements.
	// Example: {"securityGroupEgress/*/ipProtocol": "$lowercase(IpProtocol)"}
	//
	// The expressions are evaluated using JSONata (https://jsonata.org/).
	PropertyTransforms map[string]string `json:"propertyTransforms,omitempty"`
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

// ListHandlerSchema captures a limited view of a resource's list handler schema.
type ListHandlerSchema struct {
	Properties map[string]ListHandlerProperty `json:"properties,omitempty"`
	Required   []string                       `json:"required,omitempty"`
}

// ListHandlerProperty contains a subset of the CFN property schema: description and type.
type ListHandlerProperty struct {
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
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
