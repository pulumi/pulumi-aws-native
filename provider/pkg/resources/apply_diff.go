// Copyright 2024, Pulumi Corporation.

package resources

import (
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// ApplyDiff produces a new map as a merge of a calculated diff into an existing map of values.
func ApplyDiff(values resource.PropertyMap, diff *resource.ObjectDiff) resource.PropertyMap {
	if diff == nil {
		return values
	}

	result := resource.PropertyMap{}
	for name, value := range values {
		if !diff.Deleted(name) {
			result[name] = value
		}
	}
	for key, value := range diff.Adds {
		result[key] = value
	}
	for key, value := range diff.Updates {
		result[key] = value.New
	}
	return result
}

// applyFilteredDiff recursively applies a suppression-filtered detailed diff.
// Unlike ApplyDiff, it does not replace a retained parent wholesale with
// ValueDiff.New when some nested updates have been suppressed.
func applyFilteredDiff(values resource.PropertyMap, diff *resource.ObjectDiff) resource.PropertyMap {
	if diff == nil {
		return values
	}
	result := clonePropertyMap(values)
	for key := range diff.Deletes {
		delete(result, key)
	}
	for key, value := range diff.Adds {
		result[key] = clonePropertyValue(value)
	}
	for key, valueDiff := range diff.Updates {
		result[key] = applyFilteredValueDiff(result[key], valueDiff)
	}
	return result
}

// applyFilteredValueDiff rebuilds a value from its old value and only the diff
// entries that survived suppression. Replacing the parent with diff.New would
// also reintroduce nested updates that the suppression pass deliberately
// removed, such as an AWS-only reorder next to a real sibling change.
func applyFilteredValueDiff(old resource.PropertyValue, diff resource.ValueDiff) resource.PropertyValue {
	if diff.Object != nil && old.IsObject() {
		return resource.NewObjectProperty(applyFilteredDiff(old.ObjectValue(), diff.Object))
	}
	if diff.Array != nil && old.IsArray() && diff.New.IsArray() {
		oldValues := old.ArrayValue()
		newValues := diff.New.ArrayValue()
		result := make([]resource.PropertyValue, len(newValues))
		for i := range result {
			if i < len(oldValues) {
				result[i] = clonePropertyValue(oldValues[i])
			} else {
				result[i] = clonePropertyValue(newValues[i])
			}
		}
		for index, value := range diff.Array.Adds {
			if index >= 0 && index < len(result) {
				result[index] = clonePropertyValue(value)
			}
		}
		for index, valueDiff := range diff.Array.Updates {
			if index >= 0 && index < len(result) {
				result[index] = applyFilteredValueDiff(result[index], valueDiff)
			}
		}
		return resource.NewArrayProperty(result)
	}
	return clonePropertyValue(diff.New)
}
