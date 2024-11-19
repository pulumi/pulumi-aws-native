package main

import (
	"encoding/json"
	"os"
)

// Collected metadata about the Ref intrinsic function.
//
// In particular, the values that Ref intrinsic function returns are not consistent across resources. Documentation has
// some information about what this function would return for a particular resource but it is not easily
// machine-readable.
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html
type db struct {
	// Indexed by CloudFormation Resource ID
	Resources map[string]resourceInfo `json:"resources"`
}

type resourceInfo struct {
	RefReturns refReturnsInfo `json:"refReturns"`
}

// One and only one of [Property], [Properties], and [NotSupported] must be set.
type refReturnsInfo struct {
	// If set, indicates that Ref will return the value of the given Resource property.
	Property string `json:"property,omitempty"`

	// If set, indicates that Ref will return a string value obtained by joining several Resource properties with a
	// delimiter, typically "|".
	Properties []string `json:"properties,omitempty"`

	// Delimiter, typically "|". See [Properties].
	Delimiter string `json:"delimiter,omitempty"`

	// If set, Ref is not supported for this resource.
	NotSupported bool `json:"notSupported,omitempty"`
}

func (data *db) store(file string) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(file, bytes, 0600)
}

func (data *db) load(file string) error {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, data); err != nil {
		return err
	}
	if data.Resources == nil {
		data.Resources = map[string]resourceInfo{}
	}
	return nil
}
