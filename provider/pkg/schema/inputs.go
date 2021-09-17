// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func GetInputsFromState(res *CloudAPIResource, state resource.PropertyMap) (resource.PropertyMap, error) {
	inputs := resource.NewPropertyMapFromMap(map[string]interface{}{})
	for n := range res.Inputs {
		k := resource.PropertyKey(n)
		if v, ok := state[k]; ok {
			inputs[k] = v
		}
	}
	return inputs, nil
}
