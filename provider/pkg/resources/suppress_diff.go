// Copyright 2024, Pulumi Corporation.

package resources

import (
	"strings"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// SuppressAWSManagedDiffs modifies a diff to remove changes that are AWS-managed
// and should not be reflected in the user's inputs during refresh.
//
// This handles two categories of suppressions:
// 1. Generic: AWS-managed tags (aws:* prefix) that AWS adds automatically
// 2. Resource-specific: State transitions that AWS manages (e.g., EFS replication)
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
func suppressResourceSpecificChanges(resourceToken string, diff *resource.ObjectDiff) *resource.ObjectDiff {
	switch resourceToken {
	case "aws-native:efs:FileSystem":
		diff = suppressEFSReplicationProtectionTransition(diff)
	}
	return diff
}

// suppressEFSReplicationProtectionTransition prevents the diff from updating
// fileSystemProtection.replicationOverwriteProtection when AWS has transitioned
// it to REPLICATING (which happens when the filesystem becomes a replication destination).
//
// When an EFS filesystem is used as a replication destination, AWS automatically
// transitions replicationOverwriteProtection from DISABLED to REPLICATING. This is
// an AWS-managed state that the user cannot change back while replication is active.
func suppressEFSReplicationProtectionTransition(diff *resource.ObjectDiff) *resource.ObjectDiff {
	fspKey := resource.PropertyKey("fileSystemProtection")

	updated, isUpdate := diff.Updates[fspKey]
	if !isUpdate || !updated.New.IsObject() {
		return diff
	}

	newFsp := updated.New.ObjectValue()
	rop, hasRop := newFsp["replicationOverwriteProtection"]
	if !hasRop || !rop.IsString() || rop.StringValue() != "REPLICATING" {
		return diff
	}

	// AWS has set it to REPLICATING - suppress this change
	delete(diff.Updates, fspKey)
	return diff
}
