// Copyright 2016-2024, Pulumi Corporation.

package refdb

import (
	"encoding/json"
	"os"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
)

// Collected metadata about the Ref intrinsic function.
//
// In particular, the values that Ref intrinsic function returns are not consistent across resources. Documentation has
// some information about what this function would return for a particular resource but it is not easily
// machine-readable.
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html
type RefDB struct {
	// Indexed by CloudFormation Resource ID
	Resources map[string]RefDBResource `json:"resources"`
}

// See [RefDB].
type RefDBResource struct {
	RefReturns RefDBRecord `json:"refReturns"`
}

// See [RefDB].
type RefDBRecord struct {
	metadata.CfRefBehavior

	// If set, indicates the data point was assigned by a heuristic rule. This annotation can be removed once it is
	// confirmed the metadata correctly describes the behavior.
	Heuristic string `json:"heuristic,omitempty"`
}

func (data *RefDB) StoreJSONFile(file string) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(file, bytes, 0600)
}

func (data *RefDB) LoadJSONFile(file string) error {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, data); err != nil {
		return err
	}
	if data.Resources == nil {
		data.Resources = map[string]RefDBResource{}
	}
	return nil
}
