package schema

import (
	"github.com/pkg/errors"
	"reflect"
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

type sdkToCfnConverter struct {
	schema       CloudFormationSchema
	resourceType string
}

func (c *sdkToCfnConverter) sdkToCfn(properties map[string]interface{}) map[string]interface{} {
	spec := c.schema.ResourceTypes[c.resourceType]
	result := map[string]interface{}{}
	for k, prop := range spec.PropertyTypeSpec.Properties {
		if v, ok := properties[ToPropertyName(k)]; ok {
			result[k] = c.convertTypedValue(prop.Type, prop.ItemType, v)
		}
	}
	for k, prop := range spec.Attributes {
		if v, ok := properties[ToPropertyName(k)]; ok {
			result[k] = c.convertTypedValue(prop.Type, prop.ItemType, v)
		}
	}
	return result
}

func (c *sdkToCfnConverter) convertTypedValue(propType, propItemType string, v interface{}) interface{} {
	switch propType {
	case "List":
		array := v.([]interface{})
		vs := make([]interface{}, len(array))
		for i, item := range array {
			vs[i] = c.convertPropertyValue(propItemType, item)
		}
		return vs
	case "Map":
		sourceMap := v.(map[string]interface{})
		vs := map[string]interface{}{}
		for n, item := range sourceMap {
			vs[n] = c.convertPropertyValue(propItemType, item)
		}
		return vs
	case "":
		return v
	default:
		return c.convertPropertyValue(propType, v)
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
			result[k] = c.convertTypedValue(prop.Type, prop.ItemType, v)
		}
	}
	return result
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
