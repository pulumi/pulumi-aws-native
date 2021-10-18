// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schema

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/mattbaird/jsonpatch"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

// SdkToCfn converts Pulumi-SDK-shaped state to CloudFormation-shaped payload. In particular, SDK properties
// are lowerCamelCase, while CloudFormation is usually (but not always) PascalCase.
func SdkToCfn(res *CloudAPIResource, types map[string]CloudAPIType, properties map[string]interface{}) (map[string]interface{}, error) {
	converter := sdkToCfnConverter{res, types}
	return converter.sdkToCfn(properties)
}

// DiffToPatch converts a Pulumi object diff to a CloudFormation-shaped patch operation slice. Update/add/delete operations are
// mapped to corresponding patch terms, and SDK properties are translated to respective CFN names.
func DiffToPatch(res *CloudAPIResource, types map[string]CloudAPIType, diff *resource.ObjectDiff) ([]jsonpatch.JsonPatchOperation, error) {
	converter := sdkToCfnConverter{res, types}
	return converter.diffToPatch(diff)
}

type sdkToCfnConverter struct {
	spec  *CloudAPIResource
	types map[string]CloudAPIType
}

func (c *sdkToCfnConverter) sdkToCfn(properties map[string]interface{}) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	var err error
	for k, prop := range c.spec.Inputs {
		if v, ok := properties[k]; ok {
			result[ToCfnName(k)], err = c.sdkTypedValueToCfn(&prop.TypeSpec, v)
			if err != nil {
				return nil, err
			}
		}
	}
	for k, attr := range c.spec.Outputs {
		if v, ok := properties[k]; ok {
			result[ToCfnName(k)], err = c.sdkTypedValueToCfn(&attr.TypeSpec, v)
			if err != nil {
				return nil, err
			}
		}
	}
	return result, nil
}

func (c *sdkToCfnConverter) sdkTypedValueToCfn(spec *pschema.TypeSpec, v interface{}) (interface{}, error) {
	if spec.Ref != "" {
		switch spec.Ref {
		case "pulumi.json#/Any":
			return v, nil
		case "pulumi.json#/Asset", "pulumi.json#/Archive":
			switch t := v.(type) {
			case *resource.Asset:
				b, err := t.Bytes()
				if err != nil {
					return nil, err
				}
				return string(b), nil
			}
		}

		typName := strings.TrimPrefix(spec.Ref, "#/types/")
		return c.sdkObjectValueToCfn(typName, v)
	}

	var err error
	switch spec.Type {
	case "array":
		array := v.([]interface{})
		vs := make([]interface{}, len(array))
		for i, item := range array {
			vs[i], err = c.sdkTypedValueToCfn(spec.Items, item)
			if err != nil {
				return nil, err
			}
		}
		return vs, nil
	case "object":
		sourceMap := v.(map[string]interface{})
		vs := map[string]interface{}{}
		for n, item := range sourceMap {
			vs[n], err = c.sdkTypedValueToCfn(spec.AdditionalProperties, item)
			if err != nil {
				return nil, err
			}
		}
		return vs, nil
	default:
		return v, nil
	}
}

func (c *sdkToCfnConverter) sdkObjectValueToCfn(typeName string, value interface{}) (interface{}, error) {
	properties, ok := value.(map[string]interface{})
	if !ok {
		return value, nil
	}

	spec := c.types[typeName]
	result := map[string]interface{}{}
	var err error
	for k, prop := range spec.Properties {
		if v, ok := properties[k]; ok {
			result[ToCfnName(k)], err = c.sdkTypedValueToCfn(&prop.TypeSpec, v)
			if err != nil {
				return nil, err
			}
		}
	}
	return result, nil
}

func (c *sdkToCfnConverter) diffToPatch(diff *resource.ObjectDiff) ([]jsonpatch.JsonPatchOperation, error) {
	var ops []jsonpatch.JsonPatchOperation
	for sdkName, prop := range c.spec.Inputs {
		cfnName := ToCfnName(sdkName)
		key := resource.PropertyKey(sdkName)
		if v, ok := diff.Updates[key]; ok {
			op, err := c.valueToPatch("replace", cfnName, prop, v.New)
			if err != nil {
				return nil, err
			}
			ops = append(ops, op)
		}
		if v, ok := diff.Adds[key]; ok {
			op, err := c.valueToPatch("add", cfnName, prop, v)
			if err != nil {
				return nil, err
			}
			ops = append(ops, op)
		}
		if _, ok := diff.Deletes[key]; ok {
			op := jsonpatch.NewPatch("remove", "/"+cfnName, nil)
			ops = append(ops, op)
		}
	}
	return ops, nil
}

func (c *sdkToCfnConverter) valueToPatch(opName, propName string, prop pschema.PropertySpec, value resource.PropertyValue) (jsonpatch.JsonPatchOperation, error) {
	op := jsonpatch.NewPatch(opName, "/"+propName, nil)
	switch {
	case value.IsNumber() && prop.Type == "integer":
		i := int32(value.NumberValue())
		op.Value = i
	case value.IsNumber():
		op.Value = value.NumberValue()
	case value.IsBool():
		op.Value = value.BoolValue()
	case value.IsString():
		op.Value = value.StringValue()
	default:
		sdkObj := value.MapRepl(nil, nil)
		cfnObj, err := c.sdkTypedValueToCfn(&prop.TypeSpec, sdkObj)
		if err != nil {
			return jsonpatch.JsonPatchOperation{}, err
		}
		jsonBytes, err := json.Marshal(cfnObj)
		contract.AssertNoError(err)
		op.Value = string(jsonBytes)
	}
	return op, nil
}

// CfnToSdk converts CloudFormation-shaped payload to Pulumi-SDK-shaped state. In particular, SDK properties
// are lowerCamelCase, while CloudFormation is usually (but not always) PascalCase.
func CfnToSdk(properties map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for n, v := range properties {
		result[ToSdkName(n)] = cfnValueToSdk(v)
	}
	return result
}

func cfnValueToSdk(value interface{}) interface{} {
	if value == nil {
		return nil
	}
	switch reflect.TypeOf(value).Kind() {
	case reflect.Map:
		valueMap, ok := value.(map[string]interface{})
		if !ok {
			return value
		}
		return CfnToSdk(valueMap)
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(value)
		result := make([]interface{}, s.Len())
		for i := 0; i < s.Len(); i++ {
			result[i] = cfnValueToSdk(s.Index(i).Interface())
		}
		return result
	default:
		return value
	}
}
