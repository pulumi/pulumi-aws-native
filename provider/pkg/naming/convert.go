// Copyright 2016-2021, Pulumi Corporation.

package naming

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/golang/glog"
	"github.com/mattbaird/jsonpatch"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi-go-provider/resourcex"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

// SdkToCfn converts Pulumi-SDK-shaped state to CloudFormation-shaped payload. In particular, SDK properties
// are lowerCamelCase, while CloudFormation is usually (but not always) PascalCase.
func SdkToCfn(res *metadata.CloudAPIResource, types map[string]metadata.CloudAPIType, properties map[string]interface{}) (map[string]interface{}, error) {
	converter := sdkToCfnConverter{res, types}
	return converter.sdkToCfn(properties)
}

// DiffToPatch converts a Pulumi object diff to a CloudFormation-shaped patch operation slice. Update/add/delete operations are
// mapped to corresponding patch terms, and SDK properties are translated to respective CFN names.
func DiffToPatch(res *metadata.CloudAPIResource, types map[string]metadata.CloudAPIType, diff *resource.ObjectDiff) ([]jsonpatch.JsonPatchOperation, error) {
	if diff == nil {
		return []jsonpatch.JsonPatchOperation{}, nil
	}
	converter := sdkToCfnConverter{res, types}
	return converter.diffToPatch(diff)
}

var schemaStringReplacer = strings.NewReplacer(
	// Replace smart quotes with regular quotes
	"“", "\"", "”", "\"",
	// Line separator
	"\u2028", "\n",
	// Paragraph separator
	"\u2029", "\n")

// SanitizeCfnString ensures that a string from CFN docs meets the requirements for Pulumi schema strings.
func SanitizeCfnString(str string) string {
	return schemaStringReplacer.Replace(str)
}

type sdkToCfnConverter struct {
	spec  *metadata.CloudAPIResource
	types map[string]metadata.CloudAPIType
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
		contract.Assertf(len(spec.OneOf) >= 1, "at least one union case is required")

		results := []any{}
		errs := []error{}

		for _, item := range spec.OneOf {
			converted, err := c.sdkTypedValueToCfn(&item, v)
			if err != nil {
				errs = append(errs, err)
			} else {
				results = append(results, converted)
			}
		}

		switch {
		case len(results) == 0:
			glog.V(9).Infof("conversion error: all union variants failed: %v", errors.Join(errs...))
			return nil, &ConversionError{fmt.Sprintf("%v", *spec), v}
		case len(results) == 1:
			return results[0], nil
		default:
			// Properties such as aws-native:datazone:DataSource configuration do not specify a
			// discriminator yet. Without a discriminator in case of multiple successful results we could
			// instead pick the largest by length.

			sizeOf := func(x any) int {
				switch x := x.(type) {
				case map[string]any:
					return len(x)
				case []any:
					return len(x)
				case nil:
					return 0
				default:
					return 1
				}
			}

			var bestResult any
			bestResultSize := -1
			for _, r := range results {
				s := sizeOf(r)
				if s > bestResultSize {
					bestResult = r
					bestResultSize = s
				}
			}
			return bestResult, nil
		}

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
		return nil, errors.New("Unrecognized type: " + spec.Type)
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

func (c *sdkToCfnConverter) sdkObjectValueToCfn(typeName string, spec metadata.CloudAPIType, value interface{}) (interface{}, error) {
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
	// Sort keys to ensure deterministic ordering of patch operations.
	sortedKeys := make([]string, 0, len(c.spec.Inputs))
	for sdkName := range c.spec.Inputs {
		sortedKeys = append(sortedKeys, string(sdkName))
	}
	slices.Sort(sortedKeys)
	for _, sdkName := range sortedKeys {
		prop := c.spec.Inputs[sdkName]
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
	case value.IsSecret():
		return c.valueToPatch(opName, propName, prop, value.SecretValue().Element)
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
		cfnObj, err := c.sdkTypedValueToCfn(&prop.TypeSpec, resourcex.DecodeValue(value))
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

// ToStringifiedMap creates a copy of the input map with all deeply nested primitive values converted to strings.
func ToStringifiedMap[K comparable](value map[K]interface{}) map[K]interface{} {
	if value == nil {
		return nil
	}

	result := map[K]interface{}{}
	for k, v := range value {
		result[k] = primitivesToString(v)
	}
	return result
}

func primitivesToString(value interface{}) interface{} {
	if value == nil {
		return nil
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Map:
		mapValue := reflect.MakeMap(reflect.TypeOf(value))
		iter := reflect.ValueOf(value).MapRange()
		for iter.Next() {
			mapValue.SetMapIndex(iter.Key(), reflect.ValueOf(primitivesToString(iter.Value().Interface())))
		}

		return mapValue.Interface()
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(value)
		result := make([]interface{}, s.Len())
		for i := 0; i < s.Len(); i++ {
			result[i] = primitivesToString(s.Index(i).Interface())
		}
		return result
	default:
		return fmt.Sprintf("%v", value)
	}
}

type ConversionError struct {
	Type  string
	Value interface{}
}

func (e *ConversionError) Error() string {
	return fmt.Sprintf("Could not convert as type %s: %v", e.Type, e.Value)
}
