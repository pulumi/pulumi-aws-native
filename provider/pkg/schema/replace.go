package schema

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
)

func getItemTypeReplaces(schema CloudFormationSchema, itemType, resourceType, path string, diff resource.ValueDiff) ([]string, error) {
	propertyTypeName := resourceType + "." + itemType
	propertyTypeSpec, ok := schema.PropertyTypes[propertyTypeName]
	if !ok {
		propertyTypeSpec, ok = schema.PropertyTypes[itemType]
		if !ok {
			return nil, errors.Errorf("could not find property type %v in schema", propertyTypeName)
		}
	}
	return getPropertiesReplaces(schema, propertyTypeSpec, resourceType, path, diff.Object)
}

func getPropertyReplaces(schema CloudFormationSchema, spec PropertySpec, resourceType, path string, diff resource.ValueDiff) ([]string, error) {
	// If this property is immutable, it will force a replacement. We can stop here.
	if spec.UpdateType == "Immutable" {
		return []string{path}, nil
	}

	// Check the type of the property.
	if spec.PrimitiveType != "" {
		// Nothing to do.
		return nil, nil
	}
	switch spec.Type {
	case "List":
		if diff.Array == nil || spec.PrimitiveItemType != "" {
			return nil, nil
		}
		var paths []string
		for i, itemDiff := range diff.Array.Updates {
			itemPath := fmt.Sprintf("%s[%d]", path, i)
			ps, err := getItemTypeReplaces(schema, spec.ItemType, resourceType, itemPath, itemDiff)
			if err != nil {
				return nil, err
			}
			paths = append(paths, ps...)
		}
		return paths, nil
	case "Map":
		if diff.Object == nil || spec.PrimitiveItemType != "" {
			return nil, nil
		}
		var paths []string
		for k, itemDiff := range diff.Object.Updates {
			itemPath := fmt.Sprintf("%s.%s", path, k)
			ps, err := getItemTypeReplaces(schema, spec.ItemType, resourceType, itemPath, itemDiff)
			if err != nil {
				return nil, err
			}
			paths = append(paths, ps...)
		}
		return paths, nil
	case "":
		return nil, errors.Errorf("schema is missing type information for %v", path)
	default:
		return getItemTypeReplaces(schema, spec.Type, resourceType, path, diff)
	}
}

func getPropertiesReplaces(schema CloudFormationSchema, spec PropertyTypeSpec, resourceType, path string, diff *resource.ObjectDiff) ([]string, error) {
	if diff == nil {
		return nil, nil
	}

	var paths []string
	for k, itemDiff := range diff.Updates {
		var propertyPath string
		if path == "" {
			propertyPath = string(k)
		} else {
			propertyPath = fmt.Sprintf("%s.%s", path, k)
		}

		propertySpec, ok := spec.Properties[string(k)]
		if !ok {
			return nil, errors.Errorf("unknown property '%v'", path)
		}
		ps, err := getPropertyReplaces(schema, propertySpec, resourceType, propertyPath, itemDiff)
		if err != nil {
			return nil, err
		}
		paths = append(paths, ps...)
	}
	return paths, nil
}

func makeEmptyCopy(v resource.PropertyValue) resource.PropertyValue {
	switch {
	case v.IsArray():
		array := v.ArrayValue()
		copy := make([]resource.PropertyValue, len(array))
		for i, e := range array {
			copy[i] = makeEmptyCopy(e)
		}
		return resource.NewArrayProperty(copy)
	case v.IsObject():
		copy := make(resource.PropertyMap)
		for k, e := range v.ObjectValue() {
			copy[k] = makeEmptyCopy(e)
		}
		return resource.NewObjectProperty(copy)
	default:
		return resource.PropertyValue{}
	}
}

func preprocessDiff(diff resource.ValueDiff) {
	switch {
	case diff.Array != nil:
		if diff.Array.Updates == nil {
			diff.Array.Updates = make(map[int]resource.ValueDiff)
		}
		for _, v := range diff.Array.Updates {
			preprocessDiff(v)
		}
		for i, v := range diff.Array.Adds {
			if d := makeEmptyCopy(v).Diff(v); d != nil {
				diff.Array.Updates[i] = *d
			}
		}
		for i, v := range diff.Array.Deletes {
			if d := v.Diff(makeEmptyCopy(v)); d != nil {
				diff.Array.Updates[i] = *d
			}
		}
	case diff.Object != nil:
		if diff.Object.Updates == nil {
			diff.Object.Updates = make(map[resource.PropertyKey]resource.ValueDiff)
		}
		for _, v := range diff.Object.Updates {
			preprocessDiff(v)
		}
		for k, v := range diff.Object.Adds {
			if d := makeEmptyCopy(v).Diff(v); d != nil {
				diff.Object.Updates[k] = *d
			}
		}
		for k, v := range diff.Object.Deletes {
			if d := v.Diff(makeEmptyCopy(v)); d != nil {
				diff.Object.Updates[k] = *d
			}
		}
	}
}

func GetResourceReplaces(schema CloudFormationSchema, resourceType string, diff resource.ValueDiff) ([]string, error) {
	// Fetch the schema for the resource type.
	resourceSpec, ok := schema.ResourceTypes[resourceType]
	if !ok {
		return nil, errors.Errorf("unknown resource type %v", resourceType)
	}

	// Preprocess the diff to turn all changes into updates.
	switch {
	case diff.Object != nil:
		preprocessDiff(diff)
	case diff.Old.HasValue():
		if d := diff.Old.Diff(makeEmptyCopy(diff.Old)); d != nil {
			diff = *d
		}
	case diff.New.HasValue():
		if d := diff.New.Diff(makeEmptyCopy(diff.New)); d != nil {
			diff = *d
		}
	}

	return getPropertiesReplaces(schema, resourceSpec.PropertyTypeSpec, resourceType, "", diff.Object)
}
