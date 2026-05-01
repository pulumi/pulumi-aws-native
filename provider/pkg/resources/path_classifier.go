package resources

import (
	"strconv"
	"strings"

	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
)

const (
	anyRef     = "pulumi.json#/Any"
	objectType = "object"
	arrayType  = "array"
)

type PathKind int

const (
	ConcreteField PathKind = iota
	MapKey
	ArrayElement
	UnknownPath
)

type PathInfo struct {
	Path       string
	Kind       PathKind
	Input      bool
	Output     bool
	Required   bool
	ReadOnly   bool
	WriteOnly  bool
	CreateOnly bool
}

type PathClassifier struct {
	res        *metadata.CloudAPIResource
	types      map[string]metadata.CloudAPIType
	readOnly   pathSet
	writeOnly  pathSet
	createOnly pathSet
	required   pathSet
}

func NewPathClassifier(res *metadata.CloudAPIResource, types map[string]metadata.CloudAPIType) *PathClassifier {
	c := &PathClassifier{
		res:        res,
		types:      types,
		readOnly:   newPathSet(res.ReadOnly),
		writeOnly:  newPathSet(res.WriteOnly),
		createOnly: newPathSet(res.CreateOnly),
		required:   newPathSet(res.Required),
	}
	c.addNestedRequired()
	return c
}

func (c *PathClassifier) Classify(path string) (PathInfo, bool) {
	info, ok := c.classify(path)
	if !ok {
		return PathInfo{Path: path, Kind: UnknownPath}, false
	}
	info.ReadOnly = c.readOnly.matches(path)
	info.WriteOnly = c.writeOnly.matches(path)
	info.CreateOnly = c.createOnly.matches(path)
	info.Required = c.required.matches(path)
	return info, true
}

func (c *PathClassifier) ProjectActualInputs(outputs resource.PropertyMap) resource.PropertyMap {
	result := resource.PropertyMap{}
	for name, input := range c.res.Inputs {
		key := resource.PropertyKey(name)
		value, ok := outputs[key]
		if !ok || c.readOnly.matches(name) || c.writeOnly.matches(name) {
			continue
		}
		if projected, ok := c.projectValue(name, input.TypeSpec, value); ok {
			result[key] = projected
		}
	}
	return result
}

func (c *PathClassifier) ActualInputBaseline(
	oldDesired, actualInputs, newDesired resource.PropertyMap,
) resource.PropertyMap {
	result := clonePropertyMap(actualInputs)
	c.addWriteOnlyFallbacks(result, oldDesired)
	c.pruneUnownedComputed(result, oldDesired, newDesired, "")
	return result
}

func SuppressBaselineDiffs(
	resourceToken string,
	spec *metadata.CloudAPIResource,
	oldDesired, baseline resource.PropertyMap,
	transformCache *TransformCache,
) resource.PropertyMap {
	diff := oldDesired.Diff(baseline)
	diff = SuppressAWSManagedDiffsWithContext(resourceToken, spec, diff, oldDesired, oldDesired, baseline, transformCache)
	return ApplyDiff(oldDesired, diff)
}

func GetPath(m resource.PropertyMap, path string) (resource.PropertyValue, bool) {
	if path == "" {
		return resource.NewObjectProperty(m), true
	}
	return slashPathForValue(resource.NewObjectProperty(m), path).Get(resource.NewObjectProperty(m))
}

func SetPath(m resource.PropertyMap, path string, v resource.PropertyValue) {
	_, _ = slashPathForValue(resource.NewObjectProperty(m), path).Add(resource.NewObjectProperty(m), v)
}

func DeletePath(m resource.PropertyMap, path string) {
	_ = slashPathForValue(resource.NewObjectProperty(m), path).Delete(resource.NewObjectProperty(m))
}

func (c *PathClassifier) classify(path string) (PathInfo, bool) {
	segments := strings.Split(path, "/")
	if len(segments) == 0 || segments[0] == "" {
		return PathInfo{}, false
	}
	input, isInput := c.res.Inputs[segments[0]]
	output, isOutput := c.res.Outputs[segments[0]]
	if !isInput && !isOutput {
		return PathInfo{Path: path, Kind: UnknownPath}, false
	}
	if len(segments) == 1 {
		return PathInfo{Path: path, Kind: ConcreteField, Input: isInput, Output: isOutput}, true
	}
	var typ *pschema.TypeSpec
	if isInput {
		typ = &input.TypeSpec
	} else {
		typ = &output.TypeSpec
	}
	return c.classifyType(path, segments[1:], typ, isInput, isOutput)
}

func (c *PathClassifier) classifyType(
	path string, segments []string, typ *pschema.TypeSpec, input, output bool,
) (PathInfo, bool) {
	if typ == nil {
		return PathInfo{Path: path, Kind: UnknownPath, Input: input, Output: output}, false
	}
	if typ.Ref != "" && typ.Ref != anyRef {
		typeName := strings.TrimPrefix(typ.Ref, "#/types/")
		typeSpec, ok := c.types[typeName]
		if !ok {
			return PathInfo{Path: path, Kind: UnknownPath, Input: input, Output: output}, false
		}
		prop, ok := typeSpec.Properties[segments[0]]
		if !ok {
			return PathInfo{Path: path, Kind: UnknownPath, Input: input, Output: output}, false
		}
		if len(segments) == 1 {
			return PathInfo{Path: path, Kind: ConcreteField, Input: input, Output: output}, true
		}
		return c.classifyType(path, segments[1:], &prop.TypeSpec, input, output)
	}
	if typ.Type == objectType && typ.AdditionalProperties != nil {
		return PathInfo{Path: path, Kind: MapKey, Input: input, Output: output}, true
	}
	if typ.Type == arrayType && typ.Items != nil {
		if len(segments) == 1 {
			return PathInfo{Path: path, Kind: ArrayElement, Input: input, Output: output}, true
		}
		return c.classifyType(path, segments[1:], typ.Items, input, output)
	}
	return PathInfo{Path: path, Kind: UnknownPath, Input: input, Output: output}, false
}

func (c *PathClassifier) projectValue(path string, typ pschema.TypeSpec, value resource.PropertyValue) (resource.PropertyValue, bool) {
	if typ.Ref != "" && typ.Ref != anyRef {
		if !value.IsObject() {
			return value, true
		}
		typeName := strings.TrimPrefix(typ.Ref, "#/types/")
		typeSpec, ok := c.types[typeName]
		if !ok {
			return value, true
		}
		result := resource.PropertyMap{}
		for name, prop := range typeSpec.Properties {
			childPath := joinPath(path, name)
			if c.readOnly.matches(childPath) || c.writeOnly.matches(childPath) {
				continue
			}
			child, ok := value.ObjectValue()[resource.PropertyKey(name)]
			if !ok {
				continue
			}
			if projected, ok := c.projectValue(childPath, prop.TypeSpec, child); ok {
				result[resource.PropertyKey(name)] = projected
			}
		}
		return resource.NewObjectProperty(result), true
	}
	if typ.Type == arrayType && typ.Items != nil && value.IsArray() {
		values := value.ArrayValue()
		result := make([]resource.PropertyValue, 0, len(values))
		for i, item := range values {
			childPath := joinPath(path, strconv.Itoa(i))
			if projected, ok := c.projectValue(childPath, *typ.Items, item); ok {
				result = append(result, projected)
			}
		}
		return resource.NewArrayProperty(result), true
	}
	return value, true
}

func (c *PathClassifier) addNestedRequired() {
	for name, input := range c.res.Inputs {
		c.addNestedRequiredForType(name, &input.TypeSpec)
	}
}

func (c *PathClassifier) addNestedRequiredForType(path string, typ *pschema.TypeSpec) {
	if typ == nil || typ.Ref == "" || typ.Ref == anyRef {
		if typ != nil && typ.Type == arrayType {
			c.addNestedRequiredForType(joinPath(path, "*"), typ.Items)
		}
		return
	}
	typeName := strings.TrimPrefix(typ.Ref, "#/types/")
	typeSpec, ok := c.types[typeName]
	if !ok {
		return
	}
	for _, required := range typeSpec.Required {
		c.required.add(joinPath(path, required))
	}
	for name, prop := range typeSpec.Properties {
		c.addNestedRequiredForType(joinPath(path, name), &prop.TypeSpec)
	}
}

func (c *PathClassifier) addWriteOnlyFallbacks(result, oldDesired resource.PropertyMap) {
	for _, path := range c.writeOnly.paths {
		if c.createOnly.matches(path) {
			continue
		}
		for _, concretePath := range ExpandMatchingPaths(oldDesired, path) {
			if value, ok := GetPath(oldDesired, concretePath); ok {
				SetPath(result, concretePath, value)
			}
		}
	}
}

func (c *PathClassifier) AddWriteOnlyFallbacks(result, oldDesired resource.PropertyMap) {
	c.addWriteOnlyFallbacks(result, oldDesired)
}

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

func (c *PathClassifier) pruneUnownedComputed(
	actual, oldDesired, newDesired resource.PropertyMap, prefix string,
) {
	for key, value := range actual {
		path := joinPath(prefix, string(key))
		info, ok := c.Classify(path)
		if ok && info.Kind == ConcreteField && info.Input && info.Output &&
			!info.Required && !info.ReadOnly && !info.WriteOnly &&
			!hasPath(oldDesired, path) && !hasPath(newDesired, path) {
			delete(actual, key)
			continue
		}
		if value.IsObject() {
			if isSchemaObject(c, path) {
				c.pruneUnownedComputed(value.ObjectValue(), oldDesired, newDesired, path)
			}
			if len(value.ObjectValue()) == 0 && !hasPath(oldDesired, path) && !hasPath(newDesired, path) {
				delete(actual, key)
			}
		}
		if value.IsArray() {
			for i, item := range value.ArrayValue() {
				if item.IsObject() {
					c.pruneUnownedComputed(
						item.ObjectValue(), oldDesired, newDesired, joinPath(path, strconv.Itoa(i)))
				}
			}
		}
	}
}

func isSchemaObject(c *PathClassifier, path string) bool {
	info, ok := c.Classify(path)
	return ok && info.Kind == ConcreteField
}

func hasPath(m resource.PropertyMap, path string) bool {
	_, ok := GetPath(m, path)
	return ok
}

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

func slashPath(path string) resource.PropertyPath {
	segments := strings.Split(path, "/")
	result := make(resource.PropertyPath, 0, len(segments))
	for _, segment := range segments {
		result = append(result, segment)
	}
	return result
}

func clonePropertyMap(m resource.PropertyMap) resource.PropertyMap {
	result := resource.PropertyMap{}
	for k, v := range m {
		result[k] = clonePropertyValue(v)
	}
	return result
}

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

func joinPath(parent, child string) string {
	if parent == "" {
		return child
	}
	return parent + "/" + child
}

type pathSet struct {
	paths []string
}

func newPathSet(paths []string) pathSet {
	return pathSet{paths: append([]string(nil), paths...)}
}

func (s *pathSet) add(path string) {
	s.paths = append(s.paths, path)
}

func (s pathSet) matches(path string) bool {
	for _, pattern := range s.paths {
		if pathMatches(pattern, path) {
			return true
		}
	}
	return false
}

func pathMatches(pattern, path string) bool {
	return slashPath(pattern).Contains(slashPath(path)) && len(strings.Split(pattern, "/")) == len(strings.Split(path, "/"))
}
