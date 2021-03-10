package schema

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
)

func GetInputsFromState(schema CloudFormationSchema, resourceType string, state resource.PropertyMap) (resource.PropertyMap, error) {
	resourceSpec, ok := schema.ResourceTypes[resourceType]
	if !ok {
		return nil, errors.Errorf("unknown resource type %v", resourceType)
	}

	inputs := resource.NewPropertyMapFromMap(map[string]interface{}{})
	for k, v := range state {
		if _, ok := resourceSpec.Properties[string(k)]; ok {
			inputs[k] = v
		}
	}
	return inputs, nil
}
