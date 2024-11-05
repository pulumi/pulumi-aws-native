// Copyright 2024, Pulumi Corporation.

package resources

import (
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// CheckpointObject puts inputs in the `__inputs` field of the state.
func CheckpointObject(inputs resource.PropertyMap, outputs map[string]interface{}) resource.PropertyMap {
	object := resource.NewPropertyMapFromMap(outputs)
	object["__inputs"] = resource.MakeSecret(resource.NewObjectProperty(inputs))
	return object
}

// ParseCheckpointObject returns inputs that are saved in the `__inputs` field of the state.
func ParseCheckpointObject(obj resource.PropertyMap) resource.PropertyMap {
	if inputs, ok := obj["__inputs"]; ok {
		return inputs.SecretValue().Element.ObjectValue()
	}

	return nil
}
