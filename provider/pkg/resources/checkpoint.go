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
// Consecutive secret wrappers are collapsed so that checkpointed state never
// grows an extra layer of secret nesting per operation.
func CheckpointPropertyMap(inputs resource.PropertyMap, outputs resource.PropertyMap) resource.PropertyMap {
	props := resource.PropertyMap{}
	for key, value := range outputs {
		props[key] = CollapseConsecutiveSecrets(value)
	}

	props["__inputs"] = resource.MakeSecret(CollapseConsecutiveSecrets(resource.NewObjectProperty(inputs)))
	return props
}

// ParseCheckpointObject returns inputs that are saved in the `__inputs` field of the state.
// Consecutive secret wrappers are collapsed so that states written by older
// versions with nested secrets are healed when they are next read.
func ParseCheckpointObject(obj resource.PropertyMap) resource.PropertyMap {
	if inputs, ok := obj["__inputs"]; ok {
		return CollapseConsecutiveSecrets(inputs).SecretValue().Element.ObjectValue()
	}

	return nil
}

// CollapseConsecutiveSecrets returns a value in which every chain of directly
// nested secrets is replaced by a single secret wrapper. Secret(Secret(v)) and
// Secret(v) mark the same value secret, but each extra wrapper doubles the
// JSON-escaped size of the encrypted state entry. The value is rebuilt rather
// than mutated, so callers can safely pass values that alias other maps.
func CollapseConsecutiveSecrets(v resource.PropertyValue) resource.PropertyValue {
	switch {
	case v.IsSecret():
		element := CollapseConsecutiveSecrets(v.SecretValue().Element)
		if element.IsSecret() {
			return element
		}
		return resource.MakeSecret(element)
	case v.IsObject():
		obj := v.ObjectValue()
		result := make(resource.PropertyMap, len(obj))
		for key, element := range obj {
			result[key] = CollapseConsecutiveSecrets(element)
		}
		return resource.NewObjectProperty(result)
	case v.IsArray():
		arr := v.ArrayValue()
		result := make([]resource.PropertyValue, len(arr))
		for i, element := range arr {
			result[i] = CollapseConsecutiveSecrets(element)
		}
		return resource.NewArrayProperty(result)
	case v.IsOutput():
		output := v.OutputValue()
		output.Element = CollapseConsecutiveSecrets(output.Element)
		return resource.NewOutputProperty(output)
	default:
		return v
	}
}
