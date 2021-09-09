// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func GetInputsFromState(schema CloudFormationSchema, resourceType string, state resource.PropertyMap) (resource.PropertyMap, error) {
	resourceSpec, ok := schema.ResourceTypes[resourceType]
	if !ok {
		return nil, errors.Errorf("unknown resource type %v", resourceType)
	}

	inputs := resource.NewPropertyMapFromMap(map[string]interface{}{})
	for n := range resourceSpec.Properties {
		k := resource.PropertyKey(ToPropertyName(n))
		if v, ok := state[k]; ok {
			inputs[k] = v
		}
	}
	return inputs, nil
}
