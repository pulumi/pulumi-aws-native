// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	"strings"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func GetInputsFromState(res *metadata.CloudAPIResource, state resource.PropertyMap) (resource.PropertyMap, error) {
	inputs := resource.NewPropertyMapFromMap(map[string]interface{}{})
	for n := range res.Inputs {
		k := resource.PropertyKey(n)
		if v, ok := state[k]; ok {
			inputs[k] = v
		}
	}
	return inputs, nil
}

// getPropertiesWithNestedReadOnly returns a set of top-level property names
// that have nested read-only paths in the readOnly array.
// For example, if readOnly contains "replicationConfiguration/destinations/*/status",
// this returns {"replicationConfiguration": true}
func getPropertiesWithNestedReadOnly(readOnly []string) map[string]bool {
	result := make(map[string]bool)
	for _, path := range readOnly {
		parts := strings.Split(path, "/")
		// Only include if it's a nested path (has more than one part)
		if len(parts) > 1 {
			result[parts[0]] = true
		}
	}
	return result
}

// FilterEffectivelyReadOnlyProperties removes properties from refreshed inputs that:
// 1. Were NOT in the original user inputs
// 2. Have nested read-only sub-properties
//
// This prevents the provider from trying to "remove" properties that appeared
// in AWS state (like replicationConfiguration on an EFS destination) but weren't
// specified by the user.
func FilterEffectivelyReadOnlyProperties(
	refreshedInputs resource.PropertyMap,
	originalInputs resource.PropertyMap,
	readOnly []string,
) resource.PropertyMap {
	propsWithNestedReadOnly := getPropertiesWithNestedReadOnly(readOnly)

	result := make(resource.PropertyMap)
	for k, v := range refreshedInputs {
		propName := string(k)

		// If user specified this property originally, always keep it
		if _, wasOriginal := originalInputs[k]; wasOriginal {
			result[k] = v
			continue
		}

		// If this property has nested read-only paths and user didn't specify it,
		// exclude it from inputs (don't try to manage it)
		if propsWithNestedReadOnly[propName] {
			continue
		}

		result[k] = v
	}
	return result
}

// getNestedReadOnlyPaths returns a map of top-level property name to list of nested paths
// For example, if readOnly contains "fileSystemProtection/replicationOverwriteProtection",
// this returns {"fileSystemProtection": ["replicationOverwriteProtection"]}
func getNestedReadOnlyPaths(readOnly []string) map[string][]string {
	result := make(map[string][]string)
	for _, path := range readOnly {
		parts := strings.Split(path, "/")
		if len(parts) > 1 {
			topLevel := parts[0]
			nestedPath := strings.Join(parts[1:], "/")
			result[topLevel] = append(result[topLevel], nestedPath)
		}
	}
	return result
}

// ApplyReadOnlyNestedProperties copies read-only nested property values from old state
// to new inputs. This ensures that when AWS changes a nested property to a read-only state
// (like changing replicationOverwriteProtection from DISABLED to REPLICATING), the user's
// program value is overridden with the AWS state value to prevent failed updates.
func ApplyReadOnlyNestedProperties(
	newInputs resource.PropertyMap,
	oldInputs resource.PropertyMap,
	readOnly []string,
) resource.PropertyMap {
	nestedReadOnlyPaths := getNestedReadOnlyPaths(readOnly)

	result := newInputs.Copy()
	for topLevel, nestedPaths := range nestedReadOnlyPaths {
		topKey := resource.PropertyKey(topLevel)

		// Skip if old state doesn't have this property
		oldVal, hasOld := oldInputs[topKey]
		if !hasOld {
			continue
		}

		// Skip if new inputs don't have this property
		newVal, hasNew := result[topKey]
		if !hasNew {
			continue
		}

		// Both old and new have the property - apply nested read-only values
		for _, nestedPath := range nestedPaths {
			result[topKey] = applyNestedReadOnlyValue(newVal, oldVal, nestedPath)
		}
	}

	return result
}

// applyNestedReadOnlyValue recursively applies the old value at the given path to new value
func applyNestedReadOnlyValue(newVal, oldVal resource.PropertyValue, path string) resource.PropertyValue {
	parts := strings.SplitN(path, "/", 2)
	key := parts[0]

	// Handle array wildcard (*)
	if key == "*" {
		if !newVal.IsArray() || !oldVal.IsArray() {
			return newVal
		}
		// For arrays, we don't try to match - just return as-is
		// This is a simplification; full implementation would need index matching
		return newVal
	}

	// Handle object property
	if !newVal.IsObject() || !oldVal.IsObject() {
		return newVal
	}

	newObj := newVal.ObjectValue()
	oldObj := oldVal.ObjectValue()
	propKey := resource.PropertyKey(key)

	oldPropVal, hasOld := oldObj[propKey]
	if !hasOld {
		return newVal
	}

	// If this is the leaf of the path, replace the value
	if len(parts) == 1 {
		resultObj := newObj.Copy()
		resultObj[propKey] = oldPropVal
		return resource.NewObjectProperty(resultObj)
	}

	// Otherwise, recurse
	newPropVal, hasNew := newObj[propKey]
	if !hasNew {
		return newVal
	}

	resultObj := newObj.Copy()
	resultObj[propKey] = applyNestedReadOnlyValue(newPropVal, oldPropVal, parts[1])
	return resource.NewObjectProperty(resultObj)
}
