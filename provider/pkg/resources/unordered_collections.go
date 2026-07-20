// Copyright 2026, Pulumi Corporation.

package resources

import (
	"sort"
	"strconv"
	"strings"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
)

// Bound the total candidate comparisons across nested unordered collections.
// Collections that exceed the budget conservatively retain their positional diff.
const maxUnorderedComparisons = 256 * 256

// unorderedCollectionComparer removes only diff updates whose unordered
// collection values can be paired completely without transferring state between
// elements. It owns path matching and a shared work budget for nested matches.
type unorderedCollectionComparer struct {
	resourceToken string
	unordered     pathSet
	// writeOnlyGuardedPaths contains unordered roots with write-only children.
	// CloudControl omits those children, so positional fallback cannot prove
	// which hidden value belongs to an element after reordering.
	writeOnlyGuardedPaths pathSet
	remainingComparisons  int
}

// suppressUnorderedCollectionDiffs removes proven order-only updates from diff.
// Every unproven or unsupported comparison is retained unchanged.
func suppressUnorderedCollectionDiffs(
	resourceToken string,
	spec *metadata.CloudAPIResource,
	diff *resource.ObjectDiff,
) *resource.ObjectDiff {
	if diff == nil || len(spec.UnorderedCollections) == 0 {
		return diff
	}

	var writeOnlyGuardedPaths []string
	for _, unorderedPath := range spec.UnorderedCollections {
		for _, writeOnlyPath := range spec.WriteOnly {
			if metadataPathStrictlyBelow(writeOnlyPath, unorderedPath) {
				writeOnlyGuardedPaths = append(writeOnlyGuardedPaths, unorderedPath)
				break
			}
		}
	}
	comparer := unorderedCollectionComparer{
		resourceToken:         resourceToken,
		unordered:             newPathSet(spec.UnorderedCollections),
		writeOnlyGuardedPaths: newPathSet(writeOnlyGuardedPaths),
		remainingComparisons:  maxUnorderedComparisons,
	}
	return comparer.filterObjectDiff(diff, "")
}

// filterObjectDiff recursively removes proven order-only updates. Returning nil
// means every update below this object was suppressed and no structural change
// remains.
func (c *unorderedCollectionComparer) filterObjectDiff(
	diff *resource.ObjectDiff,
	path string,
) *resource.ObjectDiff {
	if diff == nil {
		return nil
	}
	updateKeys := make([]resource.PropertyKey, 0, len(diff.Updates))
	for key := range diff.Updates {
		updateKeys = append(updateKeys, key)
	}
	sort.Slice(updateKeys, func(i, j int) bool { return updateKeys[i] < updateKeys[j] })
	for _, key := range updateKeys {
		childPath := joinPath(path, string(key))
		filtered, retain := c.filterValueDiff(diff.Updates[key], childPath)
		if retain {
			diff.Updates[key] = filtered
		} else {
			delete(diff.Updates, key)
		}
	}
	if !diff.AnyChanges() {
		return nil
	}
	return diff
}

// filterValueDiff filters unordered updates within one value diff. The boolean
// reports whether the resulting update must be retained.
func (c *unorderedCollectionComparer) filterValueDiff(
	diff resource.ValueDiff,
	path string,
) (resource.ValueDiff, bool) {
	oldValue, oldSecret, oldKnown := unwrapComparableValue(diff.Old)
	newValue, newSecret, newKnown := unwrapComparableValue(diff.New)
	if c.unordered.matches(path) && oldKnown && newKnown && oldSecret == newSecret &&
		oldValue.IsArray() && newValue.IsArray() {
		if c.unorderedArraysProvablyEquivalent(oldValue.ArrayValue(), newValue.ArrayValue(), path, path) {
			return resource.ValueDiff{}, false
		}
		// An unordered root is atomic. Existing positional suppressors must not
		// reinterpret a comparison that could not prove equivalence.
		return diff, true
	}

	if diff.Object != nil {
		diff.Object = c.filterObjectDiff(diff.Object, path)
		return diff, diff.Object != nil
	}
	if diff.Array != nil {
		updateIndexes := make([]int, 0, len(diff.Array.Updates))
		for index := range diff.Array.Updates {
			updateIndexes = append(updateIndexes, index)
		}
		sort.Ints(updateIndexes)
		for _, index := range updateIndexes {
			childPath := joinPath(path, strconv.Itoa(index))
			filtered, retain := c.filterValueDiff(diff.Array.Updates[index], childPath)
			if retain {
				diff.Array.Updates[index] = filtered
			} else {
				delete(diff.Array.Updates, index)
			}
		}
		if len(diff.Array.Adds)+len(diff.Array.Deletes)+len(diff.Array.Updates) == 0 {
			return resource.ValueDiff{}, false
		}
	}
	return diff, true
}

// valuesProvablyEquivalent reports whether two values have proven semantic
// equivalence. Unknowns, unsupported wrappers, normalization failures, and real
// differences all return false so the existing positional diff is retained.
func (c *unorderedCollectionComparer) valuesProvablyEquivalent(
	oldValue, newValue resource.PropertyValue,
	oldPath, newPath string,
) bool {
	if containsUnknowns(oldValue) || containsUnknowns(newValue) ||
		oldValue.ContainsSecrets() || newValue.ContainsSecrets() {
		return false
	}

	oldValue, _, oldKnown := unwrapComparableValue(oldValue)
	newValue, _, newKnown := unwrapComparableValue(newValue)
	if !oldKnown || !newKnown {
		return false
	}
	if oldValue.DeepEquals(newValue) {
		return true
	}
	if c.isIAMPolicyPath(oldPath) && c.isIAMPolicyPath(newPath) {
		return policyDocumentsSemanticallyEqual(oldValue, newValue)
	}

	switch {
	case oldValue.IsObject() && newValue.IsObject():
		oldObject, newObject := oldValue.ObjectValue(), newValue.ObjectValue()
		if len(oldObject) != len(newObject) {
			return false
		}
		keys := make([]resource.PropertyKey, 0, len(oldObject))
		for key := range oldObject {
			keys = append(keys, key)
		}
		sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
		for _, key := range keys {
			oldChild := oldObject[key]
			newChild, ok := newObject[key]
			if !ok || !c.valuesProvablyEquivalent(oldChild, newChild,
				joinPath(oldPath, string(key)), joinPath(newPath, string(key))) {
				return false
			}
		}
		return true
	case oldValue.IsArray() && newValue.IsArray():
		if c.unordered.matches(oldPath) && c.unordered.matches(newPath) {
			return c.unorderedArraysProvablyEquivalent(oldValue.ArrayValue(), newValue.ArrayValue(), oldPath, newPath)
		}
		oldArray, newArray := oldValue.ArrayValue(), newValue.ArrayValue()
		if len(oldArray) != len(newArray) {
			return false
		}
		for i := range oldArray {
			if !c.valuesProvablyEquivalent(oldArray[i], newArray[i],
				joinPath(oldPath, strconv.Itoa(i)), joinPath(newPath, strconv.Itoa(i))) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

// unorderedArraysProvablyEquivalent reports whether a complete one-to-one
// pairing exists between two unordered arrays. The pairing preserves duplicate
// counts and uses only element pairs whose equivalence was independently proven.
func (c *unorderedCollectionComparer) unorderedArraysProvablyEquivalent(
	oldValues, newValues []resource.PropertyValue,
	oldPath, newPath string,
) bool {
	if len(oldValues) != len(newValues) ||
		c.writeOnlyGuardedPaths.matches(oldPath) || c.writeOnlyGuardedPaths.matches(newPath) {
		return false
	}

	comparisons := len(oldValues) * len(newValues)
	if comparisons > c.remainingComparisons {
		return false
	}
	c.remainingComparisons -= comparisons

	equivalent := make([][]bool, len(oldValues))
	for i, oldValue := range oldValues {
		equivalent[i] = make([]bool, len(newValues))
		for j, newValue := range newValues {
			equivalent[i][j] = c.valuesProvablyEquivalent(oldValue, newValue,
				joinPath(oldPath, strconv.Itoa(i)), joinPath(newPath, strconv.Itoa(j)))
		}
	}
	return hasPerfectMatching(equivalent)
}

// hasPerfectMatching reports whether every old element can be assigned to a
// distinct new element along a proven-equivalence edge.
func hasPerfectMatching(equivalent [][]bool) bool {
	matchedOld := make([]int, len(equivalent))
	for i := range matchedOld {
		matchedOld[i] = -1
	}
	var augment func(int, []bool) bool
	augment = func(oldIndex int, seen []bool) bool {
		for newIndex, isEquivalent := range equivalent[oldIndex] {
			if !isEquivalent || seen[newIndex] {
				continue
			}
			seen[newIndex] = true
			if matchedOld[newIndex] == -1 || augment(matchedOld[newIndex], seen) {
				matchedOld[newIndex] = oldIndex
				return true
			}
		}
		return false
	}
	for oldIndex := range equivalent {
		if !augment(oldIndex, make([]bool, len(equivalent))) {
			return false
		}
	}
	return true
}

// containsUnknowns differs from PropertyValue.ContainsUnknowns by descending
// into the element of known output wrappers.
func containsUnknowns(value resource.PropertyValue) bool {
	switch {
	case value.IsComputed():
		return true
	case value.IsOutput():
		output := value.OutputValue()
		return !output.Known || containsUnknowns(output.Element)
	case value.IsSecret():
		return containsUnknowns(value.SecretValue().Element)
	case value.IsArray():
		for _, element := range value.ArrayValue() {
			if containsUnknowns(element) {
				return true
			}
		}
	case value.IsObject():
		for _, element := range value.ObjectValue() {
			if containsUnknowns(element) {
				return true
			}
		}
	case value.IsResourceReference():
		return containsUnknowns(value.ResourceReferenceValue().ID)
	}
	return false
}

// unwrapComparableValue removes collection-level secret and known-output
// wrappers. It returns the unwrapped value, whether any wrapper was secret, and
// whether the value is known enough to compare.
func unwrapComparableValue(value resource.PropertyValue) (resource.PropertyValue, bool, bool) {
	secret := false
	for {
		switch {
		case value.IsSecret():
			secret = true
			value = value.SecretValue().Element
		case value.IsOutput():
			output := value.OutputValue()
			if !output.Known {
				return value, secret || output.Secret, false
			}
			secret = secret || output.Secret
			value = output.Element
		default:
			return value, secret, true
		}
	}
}

// isIAMPolicyPath reports whether path has IAM policy-document semantics for
// either the native or bridged IAM Role token.
func (c *unorderedCollectionComparer) isIAMPolicyPath(path string) bool {
	return (c.resourceToken == awsNativeIAMRoleToken || c.resourceToken == awsIAMRoleToken) &&
		isIAMRolePolicyDocumentPath(path)
}

// metadataPathStrictlyBelow reports whether path is a proper descendant of
// parent, treating wildcard segments on either side as compatible.
func metadataPathStrictlyBelow(path, parent string) bool {
	pathSegments := strings.Split(path, "/")
	parentSegments := strings.Split(parent, "/")
	if len(pathSegments) <= len(parentSegments) {
		return false
	}
	for i, parentSegment := range parentSegments {
		pathSegment := pathSegments[i]
		if parentSegment != "*" && pathSegment != "*" && parentSegment != pathSegment {
			return false
		}
	}
	return true
}
