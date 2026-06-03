package resources

import (
	"strconv"
	"strings"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// GetPath reads a slash-delimited SDK path from a property map.
//
// For example, GetPath(inputs, "code/imageUri") reads:
//
//	code: {
//	  imageUri: "repo/image:tag"
//	}
//
// Array indices are written as decimal path segments, such as
// "defaultActions/0/authenticateOidcConfig/clientSecret".
func GetPath(m resource.PropertyMap, path string) (resource.PropertyValue, bool) {
	if path == "" {
		return resource.NewObjectProperty(m), true
	}
	return slashPathForValue(resource.NewObjectProperty(m), path).Get(resource.NewObjectProperty(m))
}

// SetPath writes a slash-delimited SDK path into a property map.
//
// For example, SetPath(outputs, "code/imageUri", value) writes nested object
// state under the top-level "code" property. Numeric segments are treated as
// array indices when the existing value at that point is an array.
func SetPath(m resource.PropertyMap, path string, v resource.PropertyValue) {
	_, _ = slashPathForValue(resource.NewObjectProperty(m), path).Add(resource.NewObjectProperty(m), v)
}

// SetPathWithShape writes a slash-delimited SDK path into a property map,
// using shape as a guide for array indices that do not already exist in m.
func SetPathWithShape(m resource.PropertyMap, shape resource.PropertyMap, path string, v resource.PropertyValue) {
	_, _ = slashPathForValueWithShape(
		resource.NewObjectProperty(m), resource.NewObjectProperty(shape), path,
	).Add(resource.NewObjectProperty(m), v)
}

// DeletePath removes a slash-delimited SDK path from a property map.
//
// For example, DeletePath(inputs, "code/imageUri") removes only the nested
// "imageUri" value from "code" and leaves sibling fields intact.
func DeletePath(m resource.PropertyMap, path string) {
	_ = slashPathForValue(resource.NewObjectProperty(m), path).Delete(resource.NewObjectProperty(m))
}

// ExpandMatchingPaths expands a metadata path pattern against concrete values.
//
// For example, given:
//
//	defaultActions: [
//	  {authenticateOidcConfig: {clientSecret: "s0"}},
//	  {authenticateOidcConfig: {clientSecret: "s1"}},
//	]
//
// ExpandMatchingPaths(inputs,
// "defaultActions/*/authenticateOidcConfig/clientSecret") returns:
//
//	[]string{
//	  "defaultActions/0/authenticateOidcConfig/clientSecret",
//	  "defaultActions/1/authenticateOidcConfig/clientSecret",
//	}
func ExpandMatchingPaths(m resource.PropertyMap, pattern string) []string {
	if !strings.Contains(pattern, "*") {
		if _, ok := GetPath(m, pattern); ok {
			return []string{pattern}
		}
		return nil
	}
	var result []string
	expandMatchingPaths(resource.NewObjectProperty(m), strings.Split(pattern, "/"), nil, &result)
	return result
}

// expandMatchingPaths recursively walks arrays and objects for wildcard path
// expansion.
func expandMatchingPaths(v resource.PropertyValue, segments []string, prefix []string, result *[]string) {
	if len(segments) == 0 {
		*result = append(*result, strings.Join(prefix, "/"))
		return
	}
	segment := segments[0]
	if segment == "*" {
		if v.IsArray() {
			for i, child := range v.ArrayValue() {
				expandMatchingPaths(child, segments[1:], append(prefix, strconv.Itoa(i)), result)
			}
		}
		if v.IsObject() {
			for key, child := range v.ObjectValue() {
				expandMatchingPaths(child, segments[1:], append(prefix, string(key)), result)
			}
		}
		return
	}
	child, ok := getPathChild(v, segment)
	if !ok {
		return
	}
	expandMatchingPaths(child, segments[1:], append(prefix, segment), result)
}

// slashPathForValue converts a slash-delimited SDK path into Pulumi's
// PropertyPath, interpreting numeric segments as array indices only when the
// current value is an array.
func slashPathForValue(root resource.PropertyValue, path string) resource.PropertyPath {
	segments := strings.Split(path, "/")
	result := make(resource.PropertyPath, 0, len(segments))
	current := root
	for _, segment := range segments {
		if current.IsArray() {
			i, err := strconv.Atoi(segment)
			if err == nil {
				result = append(result, i)
				if i >= 0 && i < len(current.ArrayValue()) {
					current = current.ArrayValue()[i]
				} else {
					current = resource.PropertyValue{}
				}
				continue
			}
		}
		result = append(result, segment)
		if current.IsObject() {
			if child, ok := current.ObjectValue()[resource.PropertyKey(segment)]; ok {
				current = child
				continue
			}
		}
		current = resource.PropertyValue{}
	}
	return result
}

// slashPathForValueWithShape is like slashPathForValue, but treats numeric
// segments as array indices when either the destination or guide value at that
// point is an array.
func slashPathForValueWithShape(root, shape resource.PropertyValue, path string) resource.PropertyPath {
	segments := strings.Split(path, "/")
	result := make(resource.PropertyPath, 0, len(segments))
	current := root
	currentShape := shape
	for _, segment := range segments {
		if current.IsArray() || currentShape.IsArray() {
			i, err := strconv.Atoi(segment)
			if err == nil {
				result = append(result, i)
				if current.IsArray() && i >= 0 && i < len(current.ArrayValue()) {
					current = current.ArrayValue()[i]
				} else {
					current = resource.PropertyValue{}
				}
				if currentShape.IsArray() && i >= 0 && i < len(currentShape.ArrayValue()) {
					currentShape = currentShape.ArrayValue()[i]
				} else {
					currentShape = resource.PropertyValue{}
				}
				continue
			}
		}
		result = append(result, segment)
		if current.IsObject() {
			if child, ok := current.ObjectValue()[resource.PropertyKey(segment)]; ok {
				current = child
			} else {
				current = resource.PropertyValue{}
			}
		} else {
			current = resource.PropertyValue{}
		}
		if currentShape.IsObject() {
			if child, ok := currentShape.ObjectValue()[resource.PropertyKey(segment)]; ok {
				currentShape = child
			} else {
				currentShape = resource.PropertyValue{}
			}
		} else {
			currentShape = resource.PropertyValue{}
		}
	}
	return result
}

// getPathChild reads one child segment from an object or array property value.
func getPathChild(v resource.PropertyValue, segment string) (resource.PropertyValue, bool) {
	if v.IsObject() {
		child, ok := v.ObjectValue()[resource.PropertyKey(segment)]
		return child, ok
	}
	if v.IsArray() {
		i, err := strconv.Atoi(segment)
		if err != nil || i < 0 || i >= len(v.ArrayValue()) {
			return resource.PropertyValue{}, false
		}
		return v.ArrayValue()[i], true
	}
	return resource.PropertyValue{}, false
}

// clonePropertyMap deep-copies a PropertyMap for mutation.
func clonePropertyMap(m resource.PropertyMap) resource.PropertyMap {
	result := resource.PropertyMap{}
	for k, v := range m {
		result[k] = clonePropertyValue(v)
	}
	return result
}

// clonePropertyValue deep-copies object, array, and secret property values.
func clonePropertyValue(v resource.PropertyValue) resource.PropertyValue {
	switch {
	case v.IsObject():
		return resource.NewObjectProperty(clonePropertyMap(v.ObjectValue()))
	case v.IsArray():
		values := v.ArrayValue()
		cloned := make([]resource.PropertyValue, len(values))
		for i, item := range values {
			cloned[i] = clonePropertyValue(item)
		}
		return resource.NewArrayProperty(cloned)
	case v.IsSecret():
		return resource.MakeSecret(clonePropertyValue(v.SecretValue().Element))
	default:
		return v
	}
}

// joinPath appends a child segment to an existing slash-delimited path.
func joinPath(parent, child string) string {
	if parent == "" {
		return child
	}
	return parent + "/" + child
}
