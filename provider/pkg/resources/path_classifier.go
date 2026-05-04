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

// PathKind describes how a slash-delimited SDK property path maps to schema.
type PathKind int

const (
	ConcreteField PathKind = iota
	MapKey
	ArrayElement
	UnknownPath
)

// PathInfo is the schema and metadata classification for one concrete path.
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

// PathClassifier projects CloudControl output state into writable input shape
// and decides which input-shaped actual paths should be treated as user-owned.
//
// Paths use SDK names with slash separators. For example, Lambda image URI
// metadata is represented as "code/imageUri", and array metadata can use
// wildcards such as "defaultActions/*/authenticateOidcConfig/clientSecret".
type PathClassifier struct {
	res        *metadata.CloudAPIResource
	types      map[string]metadata.CloudAPIType
	readOnly   pathSet
	writeOnly  pathSet
	createOnly pathSet
	required   pathSet
}

// NewPathClassifier creates a classifier for a Cloud API resource and its
// referenced object types.
//
// The classifier combines top-level resource metadata with nested type metadata.
// For example, if a resource input refers to "#/types/aws-native:test:Settings"
// and that type marks "name" as required, the classifier treats
// "settings/name" as required when deciding whether a returned value is
// optional-computed style.
func NewPathClassifier(res *metadata.CloudAPIResource, types map[string]metadata.CloudAPIType) *PathClassifier {
	c := &PathClassifier{
		res:        res,
		types:      types,
		readOnly:   newPathSet(res.ReadOnly),
		writeOnly:  newPathSet(res.WriteOnly),
		createOnly: newPathSet(res.CreateOnly),
		required:   newPathSet(res.Required),
	}
	return c
}

// Classify returns schema and metadata information for a concrete SDK path.
//
// For example, Classify("code/s3Bucket") returns a ConcreteField for a nested
// object property, while Classify("tags/owner") returns a MapKey when "tags" is
// modeled as an object with additionalProperties.
func (c *PathClassifier) Classify(path string) (PathInfo, bool) {
	info, ok := c.classify(path)
	if !ok {
		return PathInfo{Path: path, Kind: UnknownPath}, false
	}
	info.ReadOnly = c.readOnly.matches(path)
	info.WriteOnly = c.writeOnly.matches(path)
	info.CreateOnly = c.createOnly.matches(path)
	info.Required = info.Required || c.required.matches(path)
	return info, true
}

// ProjectWritableOutputState converts output state to writable input shape.
//
// Read-only and write-only paths are excluded. Read-only values are not valid
// desired inputs, and write-only values cannot be read back from CloudControl.
// User-owned write-only values are restored later by ActualInputBaseline from
// checkpointed old desired inputs.
//
// This projection is not ownership-filtered and should not be returned as
// refreshed inputs for an existing resource. Use ActualInputBaselineFromOutputs
// when old desired inputs are available.
//
// For example, if outputs contain:
//
//	code: {s3Bucket: "b", imageUri: "secret", resolvedImageUri: "sha"}
//
// with "code/imageUri" marked write-only and "code/resolvedImageUri" marked
// read-only, the projected input value is:
//
//	code: {s3Bucket: "b"}
func (c *PathClassifier) ProjectWritableOutputState(outputs resource.PropertyMap) resource.PropertyMap {
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

// ActualInputBaselineFromOutputs builds the engine old-input baseline directly
// from output state.
//
// For example, if AWS returns optional input/output field backupTarget but the
// old desired inputs omitted it, the returned baseline still omits backupTarget.
// If old desired inputs contained managed field settings/name and AWS now
// returns a drifted value, the returned baseline contains the drifted value so
// the next engine diff can repair it.
func (c *PathClassifier) ActualInputBaselineFromOutputs(
	oldDesired, outputs, newDesired resource.PropertyMap,
) resource.PropertyMap {
	return c.actualInputBaseline(oldDesired, c.ProjectWritableOutputState(outputs), newDesired)
}

// actualInputBaseline builds the old-input baseline the engine should compare
// with the next program inputs.
//
// It starts from projected actual state, restores user-owned write-only values
// from oldDesired, and prunes optional-computed fields that are present in
// actual state but absent from both oldDesired and newDesired. For maps and
// arrays, ownership is at the containing property, so extra map keys or array
// elements remain visible when the containing property was owned.
func (c *PathClassifier) actualInputBaseline(
	oldDesired, projectedActual, newDesired resource.PropertyMap,
) resource.PropertyMap {
	result := clonePropertyMap(projectedActual)
	c.addWriteOnlyFallbacks(result, oldDesired, true)
	c.pruneUnownedComputed(result, oldDesired, newDesired, "")
	return result
}

// SuppressBaselineDiffs applies refresh-time diff suppression to a candidate
// actual baseline before it is persisted to __inputs.
//
// For example, if oldDesired has tags {"owner": "team"} and baseline has tags
// {"owner": "team", "aws:managed": "x"}, the returned baseline keeps only the
// user-owned "owner" tag while preserving the AWS-managed tag in output state.
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

// classify does the schema-only part of Classify before metadata path sets are
// applied.
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

// classifyType walks a type spec using remaining slash path segments.
//
// For object refs it resolves named properties, for additionalProperties it
// returns MapKey, and for arrays it treats the index segment as ArrayElement
// before continuing into the item type.
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
		required := stringInSlice(typeSpec.Required, segments[0])
		if len(segments) == 1 {
			return PathInfo{Path: path, Kind: ConcreteField, Input: input, Output: output, Required: required}, true
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

// projectValue recursively projects one output value through an input type
// shape, dropping read-only and write-only child paths along the way.
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

// addWriteOnlyFallbacks restores user-owned write-only values from old desired
// inputs into an actual baseline.
//
// CloudControl cannot return write-only values, so excluding them from
// ProjectWritableOutputState does not lose ownership: values set on create
// remain in __inputs and are added back here before Diff or Update patch
// generation.
func (c *PathClassifier) addWriteOnlyFallbacks(
	result, oldDesired resource.PropertyMap, includeCreateOnly bool,
) {
	for _, path := range c.writeOnly.paths {
		if !includeCreateOnly && c.createOnly.matches(path) {
			continue
		}
		for _, concretePath := range ExpandMatchingPaths(oldDesired, path) {
			if value, ok := GetPath(oldDesired, concretePath); ok {
				SetPathWithShape(result, oldDesired, concretePath, value)
			}
		}
	}
}

// AddWriteOnlyOutputFallbacks restores user-owned write-only values into
// output state, including create-only paths.
//
// This is used before output checkpointing so write-only values that the user
// previously supplied remain available in state even though AWS does not return
// them.
func (c *PathClassifier) AddWriteOnlyOutputFallbacks(result, oldDesired resource.PropertyMap) {
	c.addWriteOnlyFallbacks(result, oldDesired, true)
}

// pruneUnownedComputed removes actual values for optional-computed-style fields
// that the user did not previously own and is not starting to own now.
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

// isSchemaObject reports whether path names a schema-declared field rather than
// a freeform map key or array element.
func isSchemaObject(c *PathClassifier, path string) bool {
	info, ok := c.Classify(path)
	return ok && info.Kind == ConcreteField
}

// hasPath reports whether a slash-delimited path exists in a property map.
func hasPath(m resource.PropertyMap, path string) bool {
	_, ok := GetPath(m, path)
	return ok
}

func stringInSlice(values []string, value string) bool {
	for _, candidate := range values {
		if candidate == value {
			return true
		}
	}
	return false
}
