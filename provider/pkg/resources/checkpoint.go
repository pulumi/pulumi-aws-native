// Copyright 2024, Pulumi Corporation.

package resources

import (
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// CheckpointObject puts inputs in the `__inputs` field of the state.
func CheckpointObject(inputs resource.PropertyMap, outputs map[string]interface{}) resource.PropertyMap {
	return CheckpointPropertyMap(inputs, resource.NewPropertyMapFromMap(outputs))
}

// CheckpointPropertyMap puts inputs in the `__inputs` field of the state.
func CheckpointPropertyMap(inputs resource.PropertyMap, outputs resource.PropertyMap) resource.PropertyMap {
	var props resource.PropertyMap
	if outputs == nil {
		props = resource.PropertyMap{}
	} else {
		props = outputs.Copy()
	}

	props["__inputs"] = resource.MakeSecret(resource.NewObjectProperty(inputs))
	return props
}

// ParseCheckpointObject returns inputs that are saved in the `__inputs` field of the state.
func ParseCheckpointObject(obj resource.PropertyMap) resource.PropertyMap {
	if inputs, ok := obj["__inputs"]; ok {
		return inputs.SecretValue().Element.ObjectValue()
	}

	return nil
}
