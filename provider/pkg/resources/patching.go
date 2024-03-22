package resources

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mattbaird/jsonpatch"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/naming"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/wI2L/jsondiff"
)

func CalcPatch(oldInputs resource.PropertyMap, newInputs resource.PropertyMap, spec metadata.CloudAPIResource, types map[string]metadata.CloudAPIType) ([]jsonpatch.JsonPatchOperation, error) {
	diff := oldInputs.Diff(newInputs)

	// Write-only properties can't even be read internally within the CloudControl service so they must be included in
	// patch requests as adds to ensure the updated model validates.
	// If a property is both write-only and create-only, we should not include it in the patch request.
	createOnlyProps := codegen.NewStringSet(spec.CreateOnly...)
	writeOnlyProps := codegen.NewStringSet(spec.WriteOnly...)
	mustSendProps := writeOnlyProps.Subtract(createOnlyProps)
	for _, writeOnlyPropName := range mustSendProps.SortedValues() {
		propKey := resource.PropertyKey(writeOnlyPropName)
		if _, ok := diff.Sames[propKey]; ok {
			delete(diff.Sames, propKey)
			diff.Adds[propKey] = newInputs[propKey]
		}
	}

	return naming.DiffToPatch(&spec, types, diff)
}

func CalculateUntypedPatch(typedOldInputs ExtensionResourceInputs, typedInputs ExtensionResourceInputs) ([]jsonpatch.JsonPatchOperation, error) {
	jsonDiffPatch, err := jsondiff.Compare(typedOldInputs.Properties, typedInputs.Properties)
	if err != nil {
		return nil, fmt.Errorf("failed to compare properties: %w", err)
	}

	createOnlyProps := codegen.NewStringSet(typedInputs.CreateOnly...)
	for _, writeOnlyPropName := range typedInputs.WriteOnly {
		if createOnlyProps.Has(writeOnlyPropName) {
			continue
		}
		newValue, ok := typedInputs.Properties[writeOnlyPropName]
		if !ok {
			continue
		}
		propPath := "/" + writeOnlyPropName
		hasPatch := false
		for _, op := range jsonDiffPatch {
			if op.Path == propPath {
				hasPatch = true
				break
			}
		}
		if !hasPatch {
			jsonDiffPatch = append(jsonDiffPatch, jsondiff.Operation{
				Type:  "add",
				Path:  propPath,
				Value: newValue,
			})
		}
	}

	jsonPatch := make([]jsonpatch.JsonPatchOperation, 0, len(jsonDiffPatch))
	for _, op := range jsonDiffPatch {
		jsonPatch = append(jsonPatch, jsonpatch.JsonPatchOperation{
			Operation: op.Type,
			Path:      op.Path,
			Value:     op.Value,
		})
	}
	slices.SortStableFunc(jsonPatch, func(a, b jsonpatch.JsonPatchOperation) int {
		return strings.Compare(a.Path, b.Path)
	})
	return jsonPatch, nil
}
