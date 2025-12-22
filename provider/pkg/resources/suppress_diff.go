// Copyright 2024, Pulumi Corporation.

package resources

import (
	"fmt"
	"strings"

	"github.com/golang/glog"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/naming"
	"github.com/pulumi/pulumi-go-provider/resourcex"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// SuppressAWSManagedDiffs modifies a diff to remove changes that are AWS-managed
// and should not be reflected in the user's inputs during refresh.
//
// This handles three categories of suppressions:
// 1. Generic: AWS-managed tags (aws:* prefix) that AWS adds automatically
// 2. Resource-specific: State transitions that AWS manages (e.g., EFS replication)
// 3. PropertyTransform: CloudFormation propertyTransform normalization (e.g., case, numeric mappings)
//
// The transformCache parameter holds compiled JSONata expressions for propertyTransform
// evaluation. Pass the provider's cache instance for production use.
func SuppressAWSManagedDiffs(
	resourceToken string,
	spec *metadata.CloudAPIResource,
	diff *resource.ObjectDiff,
	originalInputs resource.PropertyMap,
	transformCache *TransformCache,
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
		diff = suppressPropertyTransformDiffs(resourceToken, spec, diff, originalInputs, transformCache)
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
	cache *TransformCache,
) *resource.ObjectDiff {
	transforms := cache.GetOrCompile(resourceToken, spec)
	if len(transforms) == 0 {
		return diff
	}

	// Convert originalInputs to a map[string]interface{} for JSONata evaluation
	inputsMap := propertyMapToMap(originalInputs)

	// Process updates by walking the pre-computed ValueDiff tree
	for key, valueDiff := range diff.Updates {
		path := string(key)

		if shouldSuppressValueDiff(transforms, path, valueDiff, inputsMap, spec.IrreversibleNames) {
			glog.V(7).Infof("Suppressing diff for %s.%s via propertyTransform", resourceToken, path)
			delete(diff.Updates, key)
		}
	}

	return diff
}

// shouldSuppressValueDiff walks the pre-computed ValueDiff tree to determine if all
// differences can be suppressed via propertyTransform normalization.
//
// This approach leverages the diff structure already computed by PropertyMap.Diff():
// - ObjectDiff contains Adds, Deletes, and Updates for nested objects
// - ArrayDiff contains Adds, Deletes, and Updates for nested arrays
// - Structural changes (adds/deletes) indicate real diffs that cannot be suppressed
// - Only leaf value changes with matching transforms can be suppressed
func shouldSuppressValueDiff(
	transforms []CompiledTransform,
	path string,
	diff resource.ValueDiff,
	inputsMap map[string]interface{},
	irreversibleNames map[string]string,
) bool {
	// Handle nested object diffs
	if diff.Object != nil {
		// Structural changes (new/deleted properties) are real diffs
		if len(diff.Object.Adds) > 0 || len(diff.Object.Deletes) > 0 {
			return false
		}
		// Recursively check all property updates
		for k, childDiff := range diff.Object.Updates {
			childPath := path + "/" + string(k)
			if !shouldSuppressValueDiff(transforms, childPath, childDiff, inputsMap, irreversibleNames) {
				return false
			}
		}
		return true
	}

	// Handle nested array diffs
	if diff.Array != nil {
		// Structural changes (new/deleted elements) are real diffs
		if len(diff.Array.Adds) > 0 || len(diff.Array.Deletes) > 0 {
			return false
		}
		// Recursively check all element updates
		for idx, childDiff := range diff.Array.Updates {
			childPath := fmt.Sprintf("%s/%d", path, idx)
			if !shouldSuppressValueDiff(transforms, childPath, childDiff, inputsMap, irreversibleNames) {
				return false
			}
		}
		return true
	}

	// Leaf value: check if a transform exists and normalizes old/new to equivalent values
	transform := FindTransformForPath(transforms, path)
	if transform != nil {
		oldVal := resourcex.DecodeValue(diff.Old)
		newVal := resourcex.DecodeValue(diff.New)
		return evaluateAndCompare(transform, oldVal, newVal, inputsMap, path, irreversibleNames)
	}

	// No transform for this leaf - cannot suppress
	return false
}

// evaluateAndCompare evaluates the transform and compares old and new values.
func evaluateAndCompare(
	transform *CompiledTransform,
	oldVal, newVal interface{},
	inputsMap map[string]interface{},
	path string,
	irreversibleNames map[string]string,
) bool {
	// Build the context for JSONata evaluation
	// The context should include sibling properties for expressions that reference them
	context := ExtractPropertyContext(inputsMap, path)
	if context == nil {
		context = inputsMap
	}

	// Also add the property values themselves, using CloudFormation naming
	// JSONata expressions reference properties by their CFN names (e.g., "IpProtocol")
	oldContext := buildCfnContext(context, path, oldVal, irreversibleNames)
	newContext := buildCfnContext(context, path, newVal, irreversibleNames)

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
func buildCfnContext(baseContext map[string]interface{}, path string, value interface{}, irreversibleNames map[string]string) map[string]interface{} {
	ctx := make(map[string]interface{})

	// Copy base context (sibling properties)
	for k, v := range baseContext {
		ctx[k] = v
		// Add CloudFormation name variation using the lookup table
		addCfnNameVariations(ctx, k, v, irreversibleNames)
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
		addCfnNameVariations(ctx, propName, value, irreversibleNames)
		return ctx
	}

	// For multi-segment paths, build nested structure
	// Also add the leaf value directly for backward compatibility
	leafName := segments[len(segments)-1]
	ctx[leafName] = value
	addCfnNameVariations(ctx, leafName, value, irreversibleNames)

	// Build nested structure: fileSystemProtection/replicationOverwriteProtection
	// becomes {"FileSystemProtection": {"ReplicationOverwriteProtection": value}}
	buildNestedContext(ctx, segments, value, irreversibleNames)

	return ctx
}

// buildNestedContext builds a nested object structure in the context from path segments.
// For path ["fileSystemProtection", "replicationOverwriteProtection"] with value "DISABLED",
// it adds multiple naming variations to ctx:
//
//	ctx["fileSystemProtection"] = {"replicationOverwriteProtection": "DISABLED"}
//	ctx["FileSystemProtection"] = {"ReplicationOverwriteProtection": "DISABLED"}
func buildNestedContext(ctx map[string]interface{}, segments []string, value interface{}, irreversibleNames map[string]string) {
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
		addCfnNameVariations(nested, propName, current, irreversibleNames)
		current = nested
	}

	// Add the outermost level to context with all name variations
	rootName := segments[0]
	ctx[rootName] = current
	addCfnNameVariations(ctx, rootName, current, irreversibleNames)
}

// addCfnNameVariations adds the CloudFormation name variation of a property to the context.
// CloudFormation uses PascalCase with uppercase acronyms (DBClusterIdentifier, IpProtocol).
// The irreversibleNames lookup table provides the exact CFN name for properties where
// simple PascalCase conversion doesn't work (e.g., dbClusterIdentifier -> DBClusterIdentifier).
func addCfnNameVariations(ctx map[string]interface{}, sdkName string, value interface{}, irreversibleNames map[string]string) {
	// Use ToCfnName which checks the lookup table first, then falls back to PascalCase
	cfnName := naming.ToCfnName(sdkName, irreversibleNames)
	ctx[cfnName] = value
}

// propertyMapToMap converts a Pulumi PropertyMap to a standard Go map for JSONata evaluation.
// Uses DecodeValue (not Mappable) because it unwraps secrets to their underlying values.
// JSONata needs plain values to evaluate transforms; secret-wrapped values would fail comparison.
func propertyMapToMap(pm resource.PropertyMap) map[string]interface{} {
	if pm == nil {
		return nil
	}
	result := make(map[string]interface{}, len(pm))
	for k, v := range pm {
		result[string(k)] = resourcex.DecodeValue(v)
	}
	return result
}
