// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"math"
	"strings"
)

type ValidationFailure struct {
	Path   string
	Reason string
}

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
	case "string":
		if !property.IsString() {
			return []ValidationFailure{{Path: path, Reason: fmt.Sprintf("%v must be a string", path)}}, nil
		}
	case "integer":
		if !property.IsNumber() || math.Trunc(property.NumberValue()) != property.NumberValue() {
			return []ValidationFailure{{Path: path, Reason: fmt.Sprintf("%v must be an integer", path)}}, nil
		}
	case "number":
		if !property.IsNumber() {
			return []ValidationFailure{{Path: path, Reason: fmt.Sprintf("%v must be a number", path)}}, nil
		}
	case "boolean":
		if !property.IsBool() {
			return []ValidationFailure{{Path: path, Reason: fmt.Sprintf("%v must be a bool", path)}}, nil
		}
	}
	return nil, nil
}

func validateProperty(types map[string]CloudAPIType, required codegen.StringSet, spec *pschema.TypeSpec, path string, property resource.PropertyValue) ([]ValidationFailure, error) {
	// If the property is secret, inspect its element.
	for property.IsSecret() {
		property = property.SecretValue().Element
	}

	// If the property value is missing and the property is required, issue a "missing required property" error.
	if required.Has(path) && property.IsNull() {
		return []ValidationFailure{{Path: path, Reason: fmt.Sprintf("missing required property %v", path)}}, nil
	}

	// If the property value is unknown or missing, we're done.
	if property.IsComputed() || property.IsOutput() || property.IsNull() {
		return nil, nil
	}

	if spec.Ref != "" {
		if spec.Ref == "pulumi.json#/Any" {
			return nil, nil
		}

		typName := strings.TrimPrefix(spec.Ref, "#/types/")
		if !property.IsObject() {
			return []ValidationFailure{{Path: path, Reason: fmt.Sprintf("%v must be an object", path)}}, nil
		}
		typeSpec, ok := types[typName]
		if !ok {
			return nil, errors.Errorf("could not find property type %v in schema", typName)
		}
		return validateProperties(types, required, typeSpec.Properties, path, property.ObjectValue())
	}

	// Check the type of the property.
	switch spec.Type {
	case "array":
		if !property.IsArray() {
			return []ValidationFailure{{Path: path, Reason: fmt.Sprintf("%v must be an array", path)}}, nil
		}
		var failures []ValidationFailure
		for i, item := range property.ArrayValue() {
			itemPath := fmt.Sprintf("%s[%d]", path, i)
			fs, err := validateProperty(types, required, spec.Items, itemPath, item)
			if err != nil {
				return nil, err
			}
			failures = append(failures, fs...)
		}
		return failures, nil
	case "object":
		if !property.IsObject() {
			return []ValidationFailure{{Path: path, Reason: fmt.Sprintf("%v must be an object", path)}}, nil
		}
		var failures []ValidationFailure
		for k, item := range property.ObjectValue() {
			itemPath := fmt.Sprintf("%s.%s", path, k)
			fs, err := validateProperty(types, required, spec.AdditionalProperties, itemPath, item)
			if err != nil {
				return nil, err
			}
			failures = append(failures, fs...)
		}
		return failures, nil
	default:
		return validatePrimitive(spec.Type, path, property)
	}
}

func validateProperties(types map[string]CloudAPIType, required codegen.StringSet, propertySpecs map[string]pschema.PropertySpec, path string, properties resource.PropertyMap) ([]ValidationFailure, error) {
	// Do a schema-directed check first.
	var failures []ValidationFailure
	remainingProperties := properties.Mappable()
	for k, propertySpec := range propertySpecs {
		var propertyPath string
		if path == "" {
			propertyPath = k
		} else {
			propertyPath = fmt.Sprintf("%s/%s", path, k)
		}
		fs, err := validateProperty(types, required, &propertySpec.TypeSpec, propertyPath, properties[resource.PropertyKey(k)])
		if err != nil {
			return nil, err
		}
		delete(remainingProperties, k)
		failures = append(failures, fs...)
	}

	// Now check for unexpected properties.
	for k := range remainingProperties {
		var propertyPath string
		if path == "" {
			propertyPath = string(k)
		} else {
			propertyPath = fmt.Sprintf("%s.%s", path, k)
		}
		failures = append(failures, ValidationFailure{Path: propertyPath, Reason: fmt.Sprintf("unknown property %v", propertyPath)})
	}

	return failures, nil
}

func ValidateResource(res *CloudAPIResource, types map[string]CloudAPIType, properties resource.PropertyMap) ([]ValidationFailure, error) {
	required := codegen.NewStringSet(res.Required...)
	return validateProperties(types, required, res.Inputs, "", properties)
}
