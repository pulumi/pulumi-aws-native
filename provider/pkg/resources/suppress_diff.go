// Copyright 2024, Pulumi Corporation.

package resources

import (
	"fmt"
	"strings"

	"github.com/golang/glog"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/naming"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// SuppressAWSManagedDiffs modifies a diff to remove changes that are AWS-managed
// and should not be reflected in the user's inputs during refresh.
//
// This handles three categories of suppressions:
// 1. Generic: AWS-managed tags (aws:* prefix) that AWS adds automatically
// 2. Resource-specific: State transitions that AWS manages (e.g., EFS replication)
// 3. PropertyTransform: CloudFormation propertyTransform normalization (e.g., case, numeric mappings)
func SuppressAWSManagedDiffs(
	resourceToken string,
	spec *metadata.CloudAPIResource,
	diff *resource.ObjectDiff,
	originalInputs resource.PropertyMap,
) *resource.ObjectDiff {
	if diff == nil {
		return nil
	}

	// 1. Generic: Suppress aws:* prefixed tag additions
	if spec.TagsProperty != "" {
		diff = suppressAWSManagedTagAdditions(spec.TagsProperty, diff, originalInputs)
	}

	// 2. Resource-specific suppressions
	diff = suppressResourceSpecificChanges(resourceToken, diff)

	// 3. PropertyTransform-based suppressions
	if len(spec.PropertyTransforms) > 0 {
		diff = suppressPropertyTransformDiffs(resourceToken, spec, diff, originalInputs)
	}

	return diff
}

// suppressAWSManagedTagAdditions filters out aws:* prefixed tags from additions
// to the tags property. AWS adds these automatically (e.g., aws:elasticfilesystem:default-backup,
// aws:servicecatalog:applicationName) and users cannot manage them.
func suppressAWSManagedTagAdditions(
	tagsProperty string,
	diff *resource.ObjectDiff,
	originalInputs resource.PropertyMap,
) *resource.ObjectDiff {
	tagsKey := resource.PropertyKey(tagsProperty)

	// Check if tags are being added
	if addedTags, isAdd := diff.Adds[tagsKey]; isAdd {
		filtered := filterAWSPrefixedTags(addedTags, originalInputs[tagsKey])
		if filtered.IsNull() {
			delete(diff.Adds, tagsKey)
		} else {
			diff.Adds[tagsKey] = filtered
		}
	}

	// Check if tags are being updated
	if updatedTags, isUpdate := diff.Updates[tagsKey]; isUpdate {
		filtered := filterAWSPrefixedTags(updatedTags.New, originalInputs[tagsKey])
		if filtered.DeepEquals(updatedTags.Old) {
			// After filtering, old and new are the same - no real change
			delete(diff.Updates, tagsKey)
		} else {
			diff.Updates[tagsKey] = resource.ValueDiff{
				Old: updatedTags.Old,
				New: filtered,
			}
		}
	}

	return diff
}

// filterAWSPrefixedTags removes aws:* prefixed tags that weren't in originalTags.
// Handles keyValueArray style tags ([]{"key": "...", "value": "..."}).
func filterAWSPrefixedTags(newTags, originalTags resource.PropertyValue) resource.PropertyValue {
	// String map style tags
	if newTags.IsObject() {
		originalKeys := make(map[string]bool)
		if originalTags.IsObject() {
			for k := range originalTags.ObjectValue() {
				originalKeys[string(k)] = true
			}
		}

		filtered := resource.PropertyMap{}
		for k, v := range newTags.ObjectValue() {
			keyStr := string(k)
			if !strings.HasPrefix(keyStr, "aws:") || originalKeys[keyStr] {
				filtered[k] = v
			}
		}
		if len(filtered) == 0 {
			return resource.NewNullProperty()
		}
		return resource.NewPropertyValue(filtered)
	}

	if !newTags.IsArray() {
		return newTags
	}

	// Build set of original tag keys
	originalKeys := make(map[string]bool)
	if originalTags.IsArray() {
		for _, tag := range originalTags.ArrayValue() {
			if key := getTagKey(tag); key != "" {
				originalKeys[key] = true
			}
		}
	}

	// Filter out aws:* tags not in original
	var filtered []resource.PropertyValue
	for _, tag := range newTags.ArrayValue() {
		keyStr := getTagKey(tag)
		if keyStr == "" {
			// Not a recognized tag format, keep it
			filtered = append(filtered, tag)
			continue
		}

		// Keep if: not aws:* prefixed, OR was in original inputs
		if !strings.HasPrefix(keyStr, "aws:") || originalKeys[keyStr] {
			filtered = append(filtered, tag)
		}
	}

	if len(filtered) == 0 {
		return resource.NewNullProperty()
	}
	return resource.NewArrayProperty(filtered)
}

// getTagKey extracts the key from a tag object, handling both lowercase "key"
// and uppercase "Key" property names used by different AWS resources.
func getTagKey(tag resource.PropertyValue) string {
	if !tag.IsObject() {
		return ""
	}
	obj := tag.ObjectValue()

	// Try lowercase "key" first (most common)
	if key, ok := obj["key"]; ok && key.IsString() {
		return key.StringValue()
	}
	// Try uppercase "Key"
	if key, ok := obj["Key"]; ok && key.IsString() {
		return key.StringValue()
	}
	return ""
}

// suppressResourceSpecificChanges handles resource-specific AWS-managed state transitions.
// NOTE: Resource-specific suppressions should be handled via CloudFormation propertyTransform
// expressions in the schema. This function is kept for any edge cases that cannot be
// expressed as JSONata transforms. Currently there are no such cases.
func suppressResourceSpecificChanges(resourceToken string, diff *resource.ObjectDiff) *resource.ObjectDiff {
	_ = resourceToken // Reserved for future resource-specific cases
	return diff
}

// suppressPropertyTransformDiffs uses CloudFormation propertyTransform specifications
// to suppress diffs where values are semantically equivalent after transformation.
//
// PropertyTransforms define JSONata expressions that normalize property values for
// comparison. For example, EC2 SecurityGroup transforms numeric IP protocols (6, 17)
// to their string equivalents (tcp, udp), and RDS transforms identifiers to lowercase.
func suppressPropertyTransformDiffs(
	resourceToken string,
	spec *metadata.CloudAPIResource,
	diff *resource.ObjectDiff,
	originalInputs resource.PropertyMap,
) *resource.ObjectDiff {
	transforms := GlobalTransformCache.GetOrCompile(resourceToken, spec)
	if len(transforms) == 0 {
		return diff
	}

	// Convert originalInputs to a map[string]interface{} for JSONata evaluation
	inputsMap := propertyMapToMap(originalInputs)

	// Process updates - these are changes to existing properties
	for key, valueDiff := range diff.Updates {
		path := string(key)
		oldVal := propertyValueToInterface(valueDiff.Old)
		newVal := propertyValueToInterface(valueDiff.New)

		if shouldSuppressPropertyDiff(transforms, path, oldVal, newVal, inputsMap, spec.IrreversibleNames) {
			glog.V(7).Infof("Suppressing diff for %s.%s via propertyTransform", resourceToken, path)
			delete(diff.Updates, key)
		}
	}

	// Process adds - new properties being added
	for key, value := range diff.Adds {
		path := string(key)
		newVal := propertyValueToInterface(value)

		// For adds, check if the transform would normalize the value to match
		// what we expect (often nil/empty maps to a default)
		if shouldSuppressPropertyAdd(transforms, path, newVal, inputsMap, spec.IrreversibleNames) {
			glog.V(7).Infof("Suppressing add for %s.%s via propertyTransform", resourceToken, path)
			delete(diff.Adds, key)
		}
	}

	return diff
}

// shouldSuppressPropertyDiff checks if a property diff should be suppressed
// based on propertyTransform normalization.
func shouldSuppressPropertyDiff(
	transforms []CompiledTransform,
	path string,
	oldVal, newVal interface{},
	inputsMap map[string]interface{},
	irreversibleNames map[string]string,
) bool {
	// Check for direct match first
	transform := FindTransformForPath(transforms, path)
	if transform != nil {
		return evaluateAndCompare(transform, oldVal, newVal, inputsMap, path)
	}

	// Handle nested objects and arrays
	oldMap, oldIsMap := oldVal.(map[string]interface{})
	newMap, newIsMap := newVal.(map[string]interface{})
	if oldIsMap && newIsMap {
		return shouldSuppressObjectDiff(transforms, path, oldMap, newMap, inputsMap, irreversibleNames)
	}

	oldArr, oldIsArr := oldVal.([]interface{})
	newArr, newIsArr := newVal.([]interface{})
	if oldIsArr && newIsArr {
		return shouldSuppressArrayDiff(transforms, path, oldArr, newArr, inputsMap, irreversibleNames)
	}

	// For primitive values without transforms, check if they're equal
	// This is needed when recursively comparing objects/arrays where
	// only some properties have transforms
	return ValuesEquivalent(oldVal, newVal)
}

// shouldSuppressObjectDiff recursively checks if all differences in an object
// should be suppressed based on transforms.
func shouldSuppressObjectDiff(
	transforms []CompiledTransform,
	basePath string,
	oldObj, newObj map[string]interface{},
	inputsMap map[string]interface{},
	irreversibleNames map[string]string,
) bool {
	// All keys must be equivalent after transforms
	allKeys := make(map[string]bool)
	for k := range oldObj {
		allKeys[k] = true
	}
	for k := range newObj {
		allKeys[k] = true
	}

	for k := range allKeys {
		propPath := basePath + "/" + k
		oldPropVal := oldObj[k]
		newPropVal := newObj[k]

		if oldPropVal == nil && newPropVal == nil {
			continue
		}

		// If only one side has the value, it's a real diff
		if (oldPropVal == nil) != (newPropVal == nil) {
			return false
		}

		if !shouldSuppressPropertyDiff(transforms, propPath, oldPropVal, newPropVal, inputsMap, irreversibleNames) {
			return false
		}
	}
	return true
}

// shouldSuppressArrayDiff checks if array differences should be suppressed.
func shouldSuppressArrayDiff(
	transforms []CompiledTransform,
	basePath string,
	oldArr, newArr []interface{},
	inputsMap map[string]interface{},
	irreversibleNames map[string]string,
) bool {
	if len(oldArr) != len(newArr) {
		return false
	}

	for i := range oldArr {
		elemPath := fmt.Sprintf("%s/%d", basePath, i)
		if !shouldSuppressPropertyDiff(transforms, elemPath, oldArr[i], newArr[i], inputsMap, irreversibleNames) {
			return false
		}
	}
	return true
}

// shouldSuppressPropertyAdd checks if a property addition should be suppressed.
// Currently disabled because propertyTransforms describe value transformations for
// existing properties, not AWS-managed defaults added by the Cloud Control API.
//
// Future use cases may include normalizing default values that AWS adds to resources
// (e.g., if AWS adds a normalized "enabled" field that matches what a transform would
// produce from user input). If such cases are discovered, implement logic here to
// compare the added value against the transform output.
func shouldSuppressPropertyAdd(
	transforms []CompiledTransform,
	path string,
	newVal interface{},
	inputsMap map[string]interface{},
	irreversibleNames map[string]string,
) bool {
	return false
}

// evaluateAndCompare evaluates the transform and compares old and new values.
func evaluateAndCompare(
	transform *CompiledTransform,
	oldVal, newVal interface{},
	inputsMap map[string]interface{},
	path string,
) bool {
	// Build the context for JSONata evaluation
	// The context should include sibling properties for expressions that reference them
	context := ExtractPropertyContext(inputsMap, path)
	if context == nil {
		context = inputsMap
	}

	// Also add the property values themselves, using CloudFormation naming
	// JSONata expressions reference properties by their CFN names (e.g., "IpProtocol")
	oldContext := buildCfnContext(context, path, oldVal)
	newContext := buildCfnContext(context, path, newVal)

	// Evaluate transform on old value
	transformedOld, err := EvaluateTransform(*transform, oldVal, oldContext)
	if err != nil {
		glog.V(9).Infof("Transform evaluation failed for old value at %s: %v", path, err)
		return false
	}

	// Evaluate transform on new value
	transformedNew, err := EvaluateTransform(*transform, newVal, newContext)
	if err != nil {
		glog.V(9).Infof("Transform evaluation failed for new value at %s: %v", path, err)
		return false
	}

	// Compare transformed values
	return ValuesEquivalent(transformedOld, transformedNew)
}

// buildCfnContext creates a context map for JSONata evaluation with CloudFormation naming.
// JSONata expressions in propertyTransform use CloudFormation property names (PascalCase),
// often with uppercase acronyms like "DBClusterIdentifier" or "IpProtocol".
//
// For nested paths like "fileSystemProtection/replicationOverwriteProtection", we build
// a nested structure that JSONata can traverse:
//
//	{
//	  "FileSystemProtection": {
//	    "ReplicationOverwriteProtection": value
//	  }
//	}
//
// This allows JSONata expressions like "$uppercase(FileSystemProtection.ReplicationOverwriteProtection)"
// to work correctly.
func buildCfnContext(baseContext map[string]interface{}, path string, value interface{}) map[string]interface{} {
	ctx := make(map[string]interface{})

	// Copy base context (sibling properties)
	for k, v := range baseContext {
		ctx[k] = v
		// Add multiple name variations for CFN compatibility
		addCfnNameVariations(ctx, k, v)
	}

	// Build nested structure from path segments
	segments := strings.Split(path, "/")
	if len(segments) == 0 {
		return ctx
	}

	// For single-segment paths (e.g., "ipProtocol"), add flat key with variations
	if len(segments) == 1 {
		propName := segments[0]
		ctx[propName] = value
		addCfnNameVariations(ctx, propName, value)
		return ctx
	}

	// For multi-segment paths, build nested structure
	// Also add the leaf value directly for backward compatibility
	leafName := segments[len(segments)-1]
	ctx[leafName] = value
	addCfnNameVariations(ctx, leafName, value)

	// Build nested structure: fileSystemProtection/replicationOverwriteProtection
	// becomes {"FileSystemProtection": {"ReplicationOverwriteProtection": value}}
	buildNestedContext(ctx, segments, value)

	return ctx
}

// buildNestedContext builds a nested object structure in the context from path segments.
// For path ["fileSystemProtection", "replicationOverwriteProtection"] with value "DISABLED",
// it adds multiple naming variations to ctx:
//
//	ctx["fileSystemProtection"] = {"replicationOverwriteProtection": "DISABLED"}
//	ctx["FileSystemProtection"] = {"ReplicationOverwriteProtection": "DISABLED"}
func buildNestedContext(ctx map[string]interface{}, segments []string, value interface{}) {
	if len(segments) < 2 {
		return
	}

	// Build the nested structure from innermost to outermost
	// Start with the leaf value
	current := value

	// Work backwards through segments (except the first)
	// to build the nested map structure
	for i := len(segments) - 1; i >= 1; i-- {
		propName := segments[i]
		nested := make(map[string]interface{})
		nested[propName] = current
		addCfnNameVariations(nested, propName, current)
		current = nested
	}

	// Add the outermost level to context with all name variations
	rootName := segments[0]
	ctx[rootName] = current
	addCfnNameVariations(ctx, rootName, current)
}

// addCfnNameVariations adds multiple naming variations of a property to the context.
// CloudFormation uses PascalCase with uppercase acronyms (DBClusterIdentifier, IpProtocol).
func addCfnNameVariations(ctx map[string]interface{}, sdkName string, value interface{}) {
	// Basic PascalCase
	pascalName := naming.ToPascalCase(sdkName)
	ctx[pascalName] = value

	// Handle common acronym patterns that CloudFormation uses
	// These are the most common ones in propertyTransform expressions
	// We use replaceCamelCaseAcronyms to properly handle word boundaries
	acronymName := replaceCamelCaseAcronyms(pascalName)
	if acronymName != pascalName {
		ctx[acronymName] = value
	}
}

// replaceCamelCaseAcronyms replaces PascalCase acronyms with their uppercase versions,
// respecting CamelCase word boundaries. For example:
// - "DbClusterIdentifier" -> "DBClusterIdentifier" (Db->DB, but Id in Identifier is not replaced)
// - "IpProtocol" -> "IPProtocol"
// - "ResourceId" -> "ResourceID"
func replaceCamelCaseAcronyms(s string) string {
	// Common acronyms that CloudFormation uses uppercase
	acronyms := []struct {
		lower string
		upper string
	}{
		// Order matters: longer acronyms first to avoid partial matches
		{"Https", "HTTPS"},
		{"Http", "HTTP"},
		{"Arn", "ARN"},
		{"Kms", "KMS"},
		{"Vpc", "VPC"},
		{"Ec2", "EC2"},
		{"Ssl", "SSL"},
		{"Tls", "TLS"},
		{"Cpu", "CPU"},
		{"Ram", "RAM"},
		{"Api", "API"},
		{"Uri", "URI"},
		{"Url", "URL"},
		{"Dns", "DNS"},
		{"Ssh", "SSH"},
		{"Efs", "EFS"},
		{"Rds", "RDS"},
		{"Sns", "SNS"},
		{"Sqs", "SQS"},
		{"Iam", "IAM"},
		{"Db", "DB"},
		{"Ip", "IP"},
		{"Id", "ID"},
	}

	result := s
	for _, acr := range acronyms {
		result = replaceAcronymAtWordBoundary(result, acr.lower, acr.upper)
	}
	return result
}

// replaceAcronymAtWordBoundary replaces an acronym only at CamelCase word boundaries.
// A word boundary in CamelCase is:
// - Start of string
// - After a lowercase letter (e.g., "resourceId" - the 'I' starts a new word)
// The acronym must be followed by:
// - End of string
// - An uppercase letter (next word)
// - A non-letter (number, etc.)
func replaceAcronymAtWordBoundary(s, lower, upper string) string {
	if len(lower) == 0 {
		return s
	}

	var result strings.Builder
	result.Grow(len(s))

	i := 0
	for i < len(s) {
		// Check if we can match the acronym at this position
		if i+len(lower) <= len(s) && s[i:i+len(lower)] == lower {
			// Check if this is a valid word boundary position
			// The acronym should start at a word boundary (start or after lowercase)
			atWordStart := i == 0 || isLower(s[i-1])

			// The acronym should end at a word boundary
			afterIdx := i + len(lower)
			atWordEnd := afterIdx >= len(s) || isUpper(s[afterIdx]) || !isLetter(s[afterIdx])

			if atWordStart && atWordEnd {
				result.WriteString(upper)
				i += len(lower)
				continue
			}
		}
		result.WriteByte(s[i])
		i++
	}

	return result.String()
}

func isLower(b byte) bool {
	return b >= 'a' && b <= 'z'
}

func isUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func isLetter(b byte) bool {
	return isLower(b) || isUpper(b)
}

// propertyMapToMap converts a Pulumi PropertyMap to a standard Go map.
func propertyMapToMap(pm resource.PropertyMap) map[string]interface{} {
	if pm == nil {
		return nil
	}
	result := make(map[string]interface{}, len(pm))
	for k, v := range pm {
		result[string(k)] = propertyValueToInterface(v)
	}
	return result
}

// propertyValueToInterface converts a Pulumi PropertyValue to a standard Go interface.
func propertyValueToInterface(pv resource.PropertyValue) interface{} {
	if pv.IsNull() {
		return nil
	}
	if pv.IsBool() {
		return pv.BoolValue()
	}
	if pv.IsNumber() {
		return pv.NumberValue()
	}
	if pv.IsString() {
		return pv.StringValue()
	}
	if pv.IsArray() {
		arr := pv.ArrayValue()
		result := make([]interface{}, len(arr))
		for i, v := range arr {
			result[i] = propertyValueToInterface(v)
		}
		return result
	}
	if pv.IsObject() {
		return propertyMapToMap(pv.ObjectValue())
	}
	if pv.IsComputed() {
		return propertyValueToInterface(pv.Input().Element)
	}
	if pv.IsSecret() {
		return propertyValueToInterface(pv.SecretValue().Element)
	}
	return nil
}
