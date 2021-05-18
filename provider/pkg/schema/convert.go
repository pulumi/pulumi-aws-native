package schema

import (
	"encoding/json"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

// SdkToCfn converts Pulumi-SDK-shaped state to CloudFormation-shaped payload. In particular, SDK properties
// are lowerCamelCase, while CloudFormation is usually (but not always) PascalCase.
func SdkToCfn(schema CloudFormationSchema, resourceType string, properties map[string]interface{}) (map[string]interface{}, error) {
	if _, ok := schema.ResourceTypes[resourceType]; !ok {
		return nil, errors.Errorf("unknown resource type %v", resourceType)
	}
	converter := sdkToCfnConverter{schema, resourceType}
	result := converter.sdkToCfn(properties)
	return result, nil
}

// DiffToPatch converts a Pulumi object diff to a CloudFormation-shaped patch operation slice. Update/add/delete operations are
// mapped to corresponding patch terms, and SDK properties are translated to respective CFN names.
func DiffToPatch(schema CloudFormationSchema, resourceType string, diff *resource.ObjectDiff) ([]*cloudformation.PatchOperation, error) {
	if _, ok := schema.ResourceTypes[resourceType]; !ok {
		return nil, errors.Errorf("unknown resource type %v", resourceType)
	}
	converter := sdkToCfnConverter{schema, resourceType}
	result := converter.diffToPatch(diff)
	return result, nil
}

type sdkToCfnConverter struct {
	schema       CloudFormationSchema
	resourceType string
}

func (c *sdkToCfnConverter) sdkToCfn(properties map[string]interface{}) map[string]interface{} {
	spec := c.schema.ResourceTypes[c.resourceType]
	result := map[string]interface{}{}
	for k, prop := range spec.PropertyTypeSpec.Properties {
		if v, ok := properties[ToPropertyName(k)]; ok {
			result[k] = c.convertTypedValue(prop.Type, prop.PrimitiveType, prop.ItemType, v)
		}
	}
	for k, attr := range spec.Attributes {
		if v, ok := properties[ToPropertyName(k)]; ok {
			result[k] = c.convertTypedValue(attr.Type, attr.PrimitiveType, attr.ItemType, v)
		}
	}
	return result
}

func (c *sdkToCfnConverter) convertTypedValue(typ, primitiveType, itemType string, v interface{}) interface{} {
	if primitiveType != "" {
		return v
	}

	switch typ {
	case "List":
		array := v.([]interface{})
		vs := make([]interface{}, len(array))
		for i, item := range array {
			vs[i] = c.convertPropertyValue(itemType, item)
		}
		return vs
	case "Map":
		sourceMap := v.(map[string]interface{})
		vs := map[string]interface{}{}
		for n, item := range sourceMap {
			vs[n] = c.convertPropertyValue(itemType, item)
		}
		return vs
	default:
		return c.convertPropertyValue(typ, v)
	}
}

func (c *sdkToCfnConverter) convertPropertyValue(propertyType string, value interface{}) interface{} {
	properties, ok := value.(map[string]interface{})
	if !ok {
		return value
	}

	spec := c.schema.ResourceTypes[c.resourceType]
	propertySpec := spec.PropertyTypeSpec
	if propertyType != "" {
		propertyTypeName := c.resourceType + "." + propertyType
		propertySpec, ok = c.schema.PropertyTypes[propertyTypeName]
		if !ok {
			return value
		}
	}

	result := map[string]interface{}{}
	for k, prop := range propertySpec.Properties {
		if v, ok := properties[ToPropertyName(k)]; ok {
			result[k] = c.convertTypedValue(prop.Type, prop.PrimitiveType, prop.ItemType, v)
		}
	}
	return result
}

func (c *sdkToCfnConverter) diffToPatch(diff *resource.ObjectDiff) []*cloudformation.PatchOperation {
	resourceSpec := c.schema.ResourceTypes[c.resourceType]
	var ops []*cloudformation.PatchOperation
	for cfnName, prop := range resourceSpec.Properties {
		sdkName := resource.PropertyKey(ToPropertyName(cfnName))
		if v, ok := diff.Updates[sdkName]; ok {
			op := c.valueToPatch("replace", cfnName, prop, v.New)
			ops = append(ops, op)
		}
		if v, ok := diff.Adds[sdkName]; ok {
			op := c.valueToPatch("add", cfnName, prop, v)
			ops = append(ops, op)
		}
		if _, ok := diff.Deletes[sdkName]; ok {
			op := cloudformation.PatchOperation{
				Op:   aws.String("remove"),
				Path: aws.String("/" + cfnName),
			}
			ops = append(ops, &op)
		}
	}
	return ops
}

func (c *sdkToCfnConverter) valueToPatch(opName, propName string, prop PropertySpec, value resource.PropertyValue) *cloudformation.PatchOperation {
	op := cloudformation.PatchOperation{
		Op:   aws.String(opName),
		Path: aws.String("/" + propName),
	}
	switch {
	case value.IsNumber() && prop.PrimitiveType == "Integer":
		op.IntegerValue = aws.Int64(int64(value.NumberValue()))
	case value.IsNumber():
		op.NumberValue = aws.Float64(value.NumberValue())
	case value.IsBool():
		op.BooleanValue = aws.Bool(value.BoolValue())
	case value.IsString():
		op.StringValue = aws.String(value.StringValue())
	default:
		sdkObj := value.MapRepl(nil, nil)
		cfnObj := c.convertTypedValue(prop.Type, prop.PrimitiveType, prop.ItemType, sdkObj)
		jsonBytes, err := json.Marshal(cfnObj)
		contract.AssertNoError(err)
		op.ObjectValue = aws.String(string(jsonBytes))
	}
	return &op
}

// CfnToSdk converts CloudFormation-shaped payload to Pulumi-SDK-shaped state. In particular, SDK properties
// are lowerCamelCase, while CloudFormation is usually (but not always) PascalCase.
func CfnToSdk(properties map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for n, v := range properties {
		result[ToPropertyName(n)] = cfnValueToSdk(v)
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
