package resources

import (
	"github.com/mattbaird/jsonpatch"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/naming"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
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
