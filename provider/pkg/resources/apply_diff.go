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
