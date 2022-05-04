// Copyright 2016-2022, Pulumi Corporation.

package provider

import (
	"context"
	"errors"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/schema"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

func (p *cfnProvider) mapCFNPropsToResource(_ context.Context, label string, inputs resource.PropertyMap) (resource.
	PropertyMap,
	error) {
	cfType, ok := inputs["cf"]
	if !ok {
		return nil, errors.New("missing required property 'cf'")
	}
	_, tok, err := schema.CFTypeToToken(cfType.StringValue())
	if err != nil {
		return nil, err
	}

	props, ok := inputs["props"]
	if !ok {
		props = resource.NewObjectProperty(resource.PropertyMap{})
	}
	propMap := props.ObjectValue()
	res, ok := p.resourceMap.Resources[tok]
	if !ok || !res.IsSupported {
		return resource.NewPropertyMapFromMap(map[string]interface{}{
			"isSupported": false,
		}), nil
	}

	mappable := propMap.Mappable()
	logging.V(9).Infof("%s[%s].CfnToSdk[%+v]", label, tok, mappable)
	transformed := schema.CfnToSdk(mappable)
	logging.V(9).Infof("%s[%s].CfnToSdk.result[%+v]", label, tok, transformed)
	outputs := codegen.NewStringSet(codegen.SortedKeys(res.Outputs)...)
	ins := codegen.NewStringSet(codegen.SortedKeys(res.Inputs)...)
	return resource.NewPropertyMapFromMap(map[string]interface{} {
		"props": transformed,
		"outputs": codegen.SortedKeys(outputs.Subtract(ins)),
		"isSupported": true,
	}), nil
}
