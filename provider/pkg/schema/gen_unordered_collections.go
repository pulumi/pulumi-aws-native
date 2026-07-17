// Copyright 2026, Pulumi Corporation.

package schema

import (
	"maps"
	"slices"
	"sort"
	"strings"

	jsschema "github.com/pulumi/jsschema"
)

const insertionOrderProperty = "insertionOrder"

// CloudFormation resource schemas put unordered-array semantics on the array
// node itself. insertionOrder is a CloudFormation extension rather than a field
// modeled by jsschema, so it is available through Schema.Extras. For example:
//
//	"Policies": {
//	  "type": "array",
//	  "insertionOrder": false,
//	  "items": {"$ref": "#/definitions/Policy"}
//	}
//
// Pulumi package schemas and SDK array types have no corresponding annotation.
// The provider therefore records the source semantics in CloudAPIResource as
// resource-relative SDK-name paths. The example above produces "policies".
//
// Nested arrays may be reached through both array items and definitions:
//
//	"ContainerDefinitions": {
//	  "type": "array",
//	  "items": {"$ref": "#/definitions/ContainerDefinition"}
//	}
//	"ContainerDefinition": {
//	  "properties": {
//	    "Environment": {"type": "array", "insertionOrder": false}
//	  }
//	}
//
// That shape produces "containerDefinitions/*/environment". At runtime the
// wildcard matches concrete diff paths such as
// "containerDefinitions/0/environment". Walking from each resource property,
// rather than collecting annotations directly from definitions, is necessary
// because one definition can be referenced at several resource paths.
//
// unorderedSchemaBranch represents one effective schema alternative at a
// resource path. refs is the reference stack for that path; it prevents cycles
// without preventing the same definition from being visited at another path.
type unorderedSchemaBranch struct {
	schema *jsschema.Schema
	refs   map[string]struct{}
}

// readUnorderedCollections walks the resource's reachable property graph and
// returns SDK-name paths for arrays that every applicable CloudFormation schema
// branch explicitly declares unordered. The sorted result is persisted in
// CloudAPIResource.UnorderedCollections for refresh, Diff, and patch handling.
func readUnorderedCollections(resourceSpec *jsschema.Schema) []string {
	if resourceSpec == nil {
		return nil
	}

	paths := map[string]struct{}{}
	for _, name := range slices.Sorted(maps.Keys(resourceSpec.Properties)) {
		walkUnorderedCollections(resourceSpec, []unorderedSchemaBranch{{
			schema: resourceSpec.Properties[name],
			refs:   map[string]struct{}{},
		}}, []string{name}, paths)
	}

	if len(paths) == 0 {
		return nil
	}
	result := make([]string, 0, len(paths))
	for path := range paths {
		result = append(result, path)
	}
	sort.Strings(result)
	return result
}

// walkUnorderedCollections carries the resource-relative CloudFormation path
// while traversing object properties and array items. Array item traversal adds
// "*" because metadata describes all elements rather than one concrete index.
func walkUnorderedCollections(
	root *jsschema.Schema,
	branches []unorderedSchemaBranch,
	path []string,
	paths map[string]struct{},
) {
	branches = expandUnorderedAlternatives(root, branches)
	if len(branches) == 0 {
		return
	}

	arrayBranches := make([]unorderedSchemaBranch, 0, len(branches))
	for _, branch := range branches {
		if schemaIsArray(branch.schema) {
			arrayBranches = append(arrayBranches, branch)
		}
	}
	if len(arrayBranches) > 0 && everyArrayBranchIsUnordered(arrayBranches) {
		paths[ResourceProperty(strings.Join(path, "/")).ToSdkName()] = struct{}{}
	}

	properties := map[string][]unorderedSchemaBranch{}
	for _, branch := range branches {
		for name, property := range branch.schema.Properties {
			properties[name] = append(properties[name], unorderedSchemaBranch{
				schema: property,
				refs:   branch.refs,
			})
		}
	}
	for _, name := range slices.Sorted(maps.Keys(properties)) {
		walkUnorderedCollections(root, properties[name], appendPath(path, name), paths)
	}

	items := []unorderedSchemaBranch{}
	for _, branch := range arrayBranches {
		if branch.schema.Items == nil {
			continue
		}
		for _, item := range branch.schema.Items.Schemas {
			items = append(items, unorderedSchemaBranch{schema: item, refs: branch.refs})
		}
	}
	if len(items) > 0 {
		walkUnorderedCollections(root, items, appendPath(path, "*"), paths)
	}
}

// expandUnorderedAlternatives resolves local references and ordinary oneOf or
// anyOf alternatives. Ambiguous composition is omitted conservatively: allOf
// requires merging constraints, and combining oneOf with anyOf requires a
// Cartesian expansion that is not warranted by current CloudFormation schemas.
func expandUnorderedAlternatives(
	root *jsschema.Schema,
	branches []unorderedSchemaBranch,
) []unorderedSchemaBranch {
	result := []unorderedSchemaBranch{}
	for _, branch := range branches {
		result = append(result, expandUnorderedAlternative(root, branch)...)
	}
	return result
}

// expandUnorderedAlternative turns one schema node into the effective branches
// that must agree before unordered semantics can be emitted for its path.
func expandUnorderedAlternative(
	root *jsschema.Schema,
	branch unorderedSchemaBranch,
) []unorderedSchemaBranch {
	if branch.schema == nil {
		return nil
	}

	if branch.schema.Reference != "" {
		return expandUnorderedReference(root, branch)
	}

	if len(branch.schema.AllOf) > 0 ||
		(len(branch.schema.AnyOf) > 0 && len(branch.schema.OneOf) > 0) {
		return nil
	}
	if len(branch.schema.AnyOf) == 0 && len(branch.schema.OneOf) == 0 {
		return []unorderedSchemaBranch{branch}
	}

	alternatives := branch.schema.AnyOf
	if len(alternatives) == 0 {
		alternatives = branch.schema.OneOf
	}
	base := cloneJSSchema(branch.schema)
	base.AnyOf = nil
	base.OneOf = nil

	result := []unorderedSchemaBranch{}
	for _, alternative := range alternatives {
		merged := cloneJSSchema(base)
		MergeJSSchema(merged, alternative)
		result = append(result, expandUnorderedAlternative(root, unorderedSchemaBranch{
			schema: merged,
			refs:   branch.refs,
		})...)
	}
	return result
}

// expandUnorderedReference overlays a referenced definition with any fields
// next to the $ref, then continues walking with path-local cycle protection.
func expandUnorderedReference(
	root *jsschema.Schema,
	branch unorderedSchemaBranch,
) []unorderedSchemaBranch {
	ref := branch.schema.Reference
	withoutRef := cloneJSSchema(branch.schema)
	withoutRef.Reference = ""
	if _, seen := branch.refs[ref]; seen {
		return expandUnorderedAlternative(root, unorderedSchemaBranch{
			schema: withoutRef,
			refs:   branch.refs,
		})
	}

	target, ok := resolveLocalDefinition(root, ref)
	if !ok {
		return expandUnorderedAlternative(root, unorderedSchemaBranch{
			schema: withoutRef,
			refs:   branch.refs,
		})
	}
	merged := cloneJSSchema(target)
	MergeJSSchema(merged, withoutRef)
	refs := cloneRefStack(branch.refs)
	refs[ref] = struct{}{}
	return expandUnorderedAlternative(root, unorderedSchemaBranch{schema: merged, refs: refs})
}

func cloneJSSchema(schema *jsschema.Schema) *jsschema.Schema {
	result := jsschema.New()
	MergeJSSchema(result, schema)
	return result
}

// resolveLocalDefinition resolves the local definition references used by
// CloudFormation resource schemas, including JSON Pointer escaping.
func resolveLocalDefinition(root *jsschema.Schema, ref string) (*jsschema.Schema, bool) {
	const prefix = "#/definitions/"
	if !strings.HasPrefix(ref, prefix) {
		return nil, false
	}
	name := strings.TrimPrefix(ref, prefix)
	name = strings.ReplaceAll(name, "~1", "/")
	name = strings.ReplaceAll(name, "~0", "~")
	definition, ok := root.Definitions[name]
	return definition, ok
}

// schemaIsArray recognizes both explicit array types and the item-bearing
// shape that the main generator also normalizes to an array.
func schemaIsArray(schema *jsschema.Schema) bool {
	return schema != nil && (schema.Type.Contains(jsschema.ArrayType) ||
		(schema.Items != nil && len(schema.Items.Schemas) > 0))
}

// everyArrayBranchIsUnordered requires explicit insertionOrder: false on each
// applicable array alternative. Missing, malformed, or true values are ordered.
func everyArrayBranchIsUnordered(branches []unorderedSchemaBranch) bool {
	for _, branch := range branches {
		value, ok := branch.schema.Extras[insertionOrderProperty].(bool)
		if !ok || value {
			return false
		}
	}
	return true
}

func cloneRefStack(refs map[string]struct{}) map[string]struct{} {
	result := make(map[string]struct{}, len(refs)+1)
	for ref := range refs {
		result[ref] = struct{}{}
	}
	return result
}

func appendPath(path []string, segment string) []string {
	result := make([]string, len(path), len(path)+1)
	copy(result, path)
	return append(result, segment)
}
