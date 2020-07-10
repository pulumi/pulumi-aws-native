package schema

import (
	"fmt"
	"math"
	"time"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
)

type ValidationFailure struct {
	Path   string
	Reason string
}

// TODO(pdg): should this just be time.RFC3339Nano?
const timestampFormat = "2006-01-02T15:04:05.999Z07:00"

func validatePrimitive(primitiveType string, path string, property resource.PropertyValue) ([]ValidationFailure, error) {
	// If the property is secret, inspect its element.
	for property.IsSecret() {
		property = property.SecretValue().Element
	}

	// If the property value is unknown we're done.
	if property.IsComputed() || property.IsOutput() {
		return nil, nil
	}

	switch primitiveType {
	case "String":
		if !property.IsString() {
			return []ValidationFailure{{Path: path, Reason: fmt.Sprintf("%v must be a string", path)}}, nil
		}
	case "Long", "Integer":
		if !property.IsNumber() || math.Trunc(property.NumberValue()) != property.NumberValue() {
			return []ValidationFailure{{Path: path, Reason: fmt.Sprintf("%v must be an integer", path)}}, nil
		}
	case "Double":
		if !property.IsNumber() {
			return []ValidationFailure{{Path: path, Reason: fmt.Sprintf("%v must be a number", path)}}, nil
		}
	case "Boolean":
		if !property.IsBool() {
			return []ValidationFailure{{Path: path, Reason: fmt.Sprintf("%v must be a bool", path)}}, nil
		}
	case "Json":
		// Nothing to do here.
	case "Timestamp":
		if !property.IsString() {
			return []ValidationFailure{{Path: path, Reason: fmt.Sprintf("%v must be a valid Timestamp", path)}}, nil
		}
		if _, err := time.Parse(timestampFormat, property.StringValue()); err != nil {
			return []ValidationFailure{{Path: path, Reason: fmt.Sprintf("%v must be a valid Timestamp (%v)", path, err)}}, nil
		}
	default:
		return nil, errors.Errorf("unexpected primitive type '%v'", primitiveType)
	}
	return nil, nil
}

func validateItemType(schema CloudFormationSchema, primitiveItemType, itemType, resourceType, path string, item resource.PropertyValue) ([]ValidationFailure, error) {
	// If the property is secret, inspect its element.
	for item.IsSecret() {
		item = item.SecretValue().Element
	}

	if primitiveItemType != "" {
		return validatePrimitive(primitiveItemType, path, item)
	}

	if !item.IsObject() {
		return []ValidationFailure{{Path: path, Reason: fmt.Sprintf("%v must be an object", path)}}, nil
	}
	propertyTypeName := resourceType + "." + itemType
	propertyTypeSpec, ok := schema.PropertyTypes[propertyTypeName]
	if !ok {
		propertyTypeSpec, ok = schema.PropertyTypes[itemType]
		if !ok {
			return nil, errors.Errorf("could not find property type %v in schema", propertyTypeName)
		}
	}
	return validateProperties(schema, propertyTypeSpec, resourceType, path, item.ObjectValue())
}

func validateProperty(schema CloudFormationSchema, spec PropertySpec, resourceType, path string, property resource.PropertyValue) ([]ValidationFailure, error) {
	// If the property is secret, inspect its element.
	for property.IsSecret() {
		property = property.SecretValue().Element
	}

	// If the property value is missing and the property is required, issue a "missing required property" error.
	if spec.Required && property.IsNull() {
		return []ValidationFailure{{Path: path, Reason: fmt.Sprintf("missing required property %v", path)}}, nil
	}

	// If the property value is unknown or missing, we're done.
	if property.IsComputed() || property.IsOutput() || property.IsNull() {
		return nil, nil
	}

	// Check the type of the property.
	if spec.PrimitiveType != "" {
		return validatePrimitive(spec.PrimitiveType, path, property)
	}
	switch spec.Type {
	case "List":
		if !property.IsArray() {
			return []ValidationFailure{{Path: path, Reason: fmt.Sprintf("%v must be a list", path)}}, nil
		}
		var failures []ValidationFailure
		for i, item := range property.ArrayValue() {
			itemPath := fmt.Sprintf("%s[%d]", path, i)
			fs, err := validateItemType(schema, spec.PrimitiveItemType, spec.ItemType, resourceType, itemPath, item)
			if err != nil {
				return nil, err
			}
			failures = append(failures, fs...)
		}
		return failures, nil
	case "Map":
		if !property.IsObject() {
			return []ValidationFailure{{Path: path, Reason: fmt.Sprintf("%v must be an object", path)}}, nil
		}
		var failures []ValidationFailure
		for k, item := range property.ObjectValue() {
			itemPath := fmt.Sprintf("%s.%s", path, k)
			fs, err := validateItemType(schema, spec.PrimitiveItemType, spec.ItemType, resourceType, itemPath, item)
			if err != nil {
				return nil, err
			}
			failures = append(failures, fs...)
		}
		return failures, nil
	case "":
		return nil, errors.Errorf("schema is missing type information for %v", path)
	default:
		if !property.IsObject() {
			return []ValidationFailure{{Path: path, Reason: fmt.Sprintf("%v must be an object", path)}}, nil
		}
		propertyTypeName := fmt.Sprintf("%s.%s", resourceType, spec.Type)
		propertyTypeSpec, ok := schema.PropertyTypes[propertyTypeName]
		if !ok {
			propertyTypeSpec, ok = schema.PropertyTypes[spec.Type]
			if !ok {
				return nil, errors.Errorf("could not find property type %v in schema", propertyTypeName)
			}
		}
		return validateProperties(schema, propertyTypeSpec, resourceType, path, property.ObjectValue())
	}
}

func validateProperties(schema CloudFormationSchema, spec PropertyTypeSpec, resourceType, path string, properties resource.PropertyMap) ([]ValidationFailure, error) {
	// Do a schema-directed check first.
	var failures []ValidationFailure
	for k, propertySpec := range spec.Properties {
		var propertyPath string
		if path == "" {
			propertyPath = k
		} else {
			propertyPath = fmt.Sprintf("%s.%s", path, k)
		}
		fs, err := validateProperty(schema, propertySpec, resourceType, propertyPath, properties[resource.PropertyKey(k)])
		if err != nil {
			return nil, err
		}
		failures = append(failures, fs...)
	}

	// Now check for unexpected properties.
	for k := range properties {
		if _, has := spec.Properties[string(k)]; !has {
			var propertyPath string
			if path == "" {
				propertyPath = string(k)
			} else {
				propertyPath = fmt.Sprintf("%s.%s", path, k)
			}
			failures = append(failures, ValidationFailure{Path: propertyPath, Reason: fmt.Sprintf("unknown property %v", propertyPath)})
		}
	}

	return failures, nil
}

func ValidateResource(schema CloudFormationSchema, resourceType string, properties resource.PropertyMap) ([]ValidationFailure, error) {
	// Fetch the schema for the resource type.
	resourceSpec, ok := schema.ResourceTypes[resourceType]
	if !ok {
		return nil, errors.Errorf("unknown resource type %v", resourceType)
	}

	return validateProperties(schema, resourceSpec.PropertyTypeSpec, resourceType, "", properties)
}
