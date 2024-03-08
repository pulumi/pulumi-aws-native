// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/mattbaird/jsonpatch"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
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
	if diff == nil {
		return []jsonpatch.JsonPatchOperation{}, nil
	}
	converter := sdkToCfnConverter{res, types}
	return converter.diffToPatch(diff)
}

var quotationReplacer = strings.NewReplacer("“", "\"", "”", "\"")

// SanitizeCfnString ensures that a string from CFN docs meets the requirements for Pulumi schema strings.
func SanitizeCfnString(str string) string {
	return quotationReplacer.Replace(str)
}

type sdkToCfnConverter struct {
	spec  *CloudAPIResource
	types map[string]CloudAPIType
}

func (c *sdkToCfnConverter) sdkToCfn(properties map[string]interface{}) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	for k, prop := range c.spec.Inputs {
		if v, ok := properties[k]; ok {
			converted, err := c.sdkTypedValueToCfn(&prop.TypeSpec, v)
			if err != nil {
				return nil, err
			}
			result[ToCfnName(k, c.spec.IrreversibleNames)] = converted
		}
	}
	for k, attr := range c.spec.Outputs {
		if v, ok := properties[k]; ok {
			converted, err := c.sdkTypedValueToCfn(&attr.TypeSpec, v)
			if err != nil {
				return nil, err
			}
			result[ToCfnName(k, c.spec.IrreversibleNames)] = converted
		}
	}
	return result, nil
}

func (c *sdkToCfnConverter) sdkTypedValueToCfn(spec *pschema.TypeSpec, v interface{}) (interface{}, error) {
	if spec.Ref != "" {
		if spec.Ref == "pulumi.json#/Any" {
			return v, nil
		}

		typName := strings.TrimPrefix(spec.Ref, "#/types/")
		typSpec := c.types[typName]
		switch typSpec.Type {
		case "object":
			return c.sdkObjectValueToCfn(typName, typSpec, v)
		case "string":
			if _, ok := v.(string); ok {
				return v, nil
			}
		default:
			return nil, &ConversionError{spec.Type, v}
		}
	}

	if spec.OneOf != nil {
		for _, item := range spec.OneOf {
			converted, err := c.sdkTypedValueToCfn(&item, v)
			// return the first successful conversion
			if err == nil {
				return converted, nil
			}
		}
		return nil, &ConversionError{fmt.Sprintf("%v", *spec), v}
	}

	switch spec.Type {
	case "array":
		if array, ok := v.([]interface{}); ok {
			vs := make([]interface{}, len(array))
			for i, item := range array {
				converted, err := c.sdkTypedValueToCfn(spec.Items, item)
				if err != nil {
					return nil, err
				}
				vs[i] = converted
			}
			return vs, nil
		}
	case "object":
		if sourceMap, ok := v.(map[string]interface{}); ok {
			vs := map[string]interface{}{}
			for n, item := range sourceMap {
				converted, err := c.sdkTypedValueToCfn(spec.AdditionalProperties, item)
				if err != nil {
					return nil, err
				}
				vs[n] = converted
			}
			return vs, nil
		}
	case "string":
		if _, ok := v.(string); ok {
			return v, nil
		}
	case "number":
		if isNumberLike(v) {
			return v, nil
		}
	case "integer":
		if isNumberLike(v) {
			return v, nil
		}
	case "boolean":
		if _, ok := v.(bool); ok {
			return v, nil
		}
	default:
		return nil, fmt.Errorf("Unrecognized type: " + spec.Type)
	}
	return nil, &ConversionError{spec.Type, v}
}

func isNumberLike(v interface{}) bool {
	switch v.(type) {
	case int, uint, int32, uint32, int64, uint64, float32, float64:
		return true
	}
	return false
}

func (c *sdkToCfnConverter) sdkObjectValueToCfn(typeName string, spec CloudAPIType, value interface{}) (interface{}, error) {
	properties, ok := value.(map[string]interface{})
	if !ok {
		return nil, &ConversionError{typeName, value}
	}

	result := map[string]interface{}{}
	for k, prop := range spec.Properties {
		if v, ok := properties[k]; ok {
			converted, err := c.sdkTypedValueToCfn(&prop.TypeSpec, v)
			if err != nil {
				return nil, err
			}
			result[ToCfnName(k, spec.IrreversibleNames)] = converted
		}
	}
	return result, nil
}

func (c *sdkToCfnConverter) diffToPatch(diff *resource.ObjectDiff) ([]jsonpatch.JsonPatchOperation, error) {
	var ops []jsonpatch.JsonPatchOperation
	for sdkName, prop := range c.spec.Inputs {
		cfnName := ToCfnName(sdkName, c.spec.IrreversibleNames)
		key := resource.PropertyKey(sdkName)
		if v, ok := diff.Updates[key]; ok {
			// Check if properties are write-only, and use an "add" if so. This is because the old values of write only
			// properties don't show up when applying the diff within CC, so we need to do an "add".
			if c.isPropWriteOnly(sdkName) {
				op, err := c.valueToPatch("add", cfnName, prop, v.New)
				if err != nil {
					return nil, err
				}
				ops = append(ops, *op)
			} else {
				op, err := c.valueToPatch("replace", cfnName, prop, v.New)
				if err != nil {
					return nil, err
				}
				ops = append(ops, *op)
			}
		}
		if v, ok := diff.Adds[key]; ok {
			op, err := c.valueToPatch("add", cfnName, prop, v)
			if err != nil {
				return nil, err
			}
			ops = append(ops, *op)
		}
		if _, ok := diff.Deletes[key]; ok {
			op := jsonpatch.NewPatch("remove", "/"+cfnName, nil)
			ops = append(ops, op)
		}
	}
	return ops, nil
}

func (c *sdkToCfnConverter) valueToPatch(opName, propName string, prop pschema.PropertySpec, value resource.PropertyValue) (*jsonpatch.JsonPatchOperation, error) {
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
			return nil, err
		}
		op.Value = cfnObj
	}
	return &op, nil
}

func (c *sdkToCfnConverter) isPropWriteOnly(cfnName string) bool {
	for _, createOnlyProp := range c.spec.WriteOnly {
		if createOnlyProp == cfnName {
			return true
		}
	}
	return false
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

type ConversionError struct {
	Type  string
	Value interface{}
}

func (e *ConversionError) Error() string {
	return fmt.Sprintf("Could not convert as type %s: %v", e.Type, e.Value)
}
