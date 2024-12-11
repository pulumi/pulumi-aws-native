package provider

import (
	"fmt"
	"slices"
	"strings"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/naming"
	rSchema "github.com/pulumi/pulumi-aws-native/provider/pkg/schema"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// This calculates the outputs of a resource based on the inputs and the output
// properties in the resource schema.
//
// For example, if there is a property "someProperty" that is both an input and
// an output, the underlying type could have readonly properties meaning that part
// of the type is an output only value. Those will not exist as input values
// and will be marked as computed
func previewResourceOutputs(
	inputs resource.PropertyMap,
	types map[string]metadata.CloudAPIType,
	outputs map[string]schema.PropertySpec,
	readOnly []string,
) resource.PropertyMap {
	result := resource.PropertyMap{}
	for name, prop := range outputs {
		key := resource.PropertyKey(name)
		if inputValue, ok := inputs[key]; ok {
			result[key] = previewOutputValue(inputValue, types, &prop.TypeSpec, filterAndReturnNested(readOnly, name))
			// otherwise we could have one of two cases
			// 1. the property is an input & output property, but does not have an input value
			// 2. the property is a readonly (output only) property
			// In either case, the property could be computed so we should default to marking it as computed
		} else if isReadOnly(readOnly, name) {
			result[key] = resource.MakeComputed(resource.NewStringProperty(""))
		}
	}
	return result
}

// populateStableOutputs updates the preview outputs with the outputs from
// the prior state, but only in certain cases.
//
// There is no way (based on the schema) to tell which output properties are stable
// (i.e. they will only change if the resource is replaced). Because of this we
// cannot blanket copy all outputs from the prior state. The ones that are not stable
// should remain computed since they could be changing.
//
// Instead, we use some heuristics to determine which outputs are stable and which are not.
// We are making the assumption that Name, Id, and Arn are stable outputs. Sometimes resources
// will have `name`, `id`, and `arn` properties, and sometimes they will have `resourceName`, `resourceId`,
// and `resourceArn` properties.
func populateStableOutputs(
	outputsFromInputs resource.PropertyMap,
	outputsFromPriorState resource.PropertyMap,
	readOnly []string,
	resourceTypeName tokens.TypeName,
) resource.PropertyMap {
	for _, readOnlyProp := range readOnly {
		// For now skip outputs that are part of an array property.
		// We have no way of knowing which item in the outputs array corresponds
		// to which item from the inputs array (since inputs could be changing).
		if strings.Contains(readOnlyProp, "/*/") {
			continue
		}
		if isStableOutput(readOnlyProp, resourceTypeName) {
			// nested property. Sometimes the resource Arn is in a nested property
			if strings.Contains(readOnlyProp, "/") {
				props := strings.Split(readOnlyProp, "/")
				current := props[0]
				key := resource.PropertyKey(current)
				if outputFromInput, ok := outputsFromInputs[key]; ok {
					if outputFromState, ok := outputsFromPriorState[key]; ok {
						outputsFromInputs[key] = resource.NewObjectProperty(populateStableOutputs(
							outputFromInput.ObjectValue(),
							outputFromState.ObjectValue(),
							[]string{strings.Join(props[1:], "/")},
							resourceTypeName,
						))
					}
				}
			} else {
				key := resource.PropertyKey(readOnlyProp)
				if output, ok := outputsFromPriorState[key]; ok {
					outputsFromInputs[key] = output
				}
			}
		}
	}
	return outputsFromInputs
}

// isStableOutput determines if a property is a stable output or not.
// This uses a very conservative heuristic and does not cover all cases.
//
// e.g. sometimes the arn property does not contain the full resource type name
// - `aws-native:sagemaker:ModelExplainabilityJobDefinition` has a property `jobDefinitionArn`
// - `aws-native:securityhub:PolicyAssociation` has a property `associationIdentifier`
//   - TODO[pulumi/aws-native#1892]: in some cases this property is the `primaryIdentifier`. Could we use that as another heuristic?
//     a readonly property that is also a primary identifier? It doesn't catch all cases, but would catch more
func isStableOutput(propName string, resourceTypeName tokens.TypeName) bool {
	typeName := naming.ToSdkName(resourceTypeName.String())
	stableOutputsNameOnly := []string{
		fmt.Sprintf("%sName", typeName),
		fmt.Sprintf("%sId", typeName),
		fmt.Sprintf("%sArn", typeName),
	}
	// These are the most common properties that are stable outputs
	// and are used to determine if an output is stable or not
	topLevelStableOutputs := slices.Concat([]string{
		"name",
		"id",
		"arn",
	}, stableOutputsNameOnly)

	// we can't handle arrays because we don't know which item in the array
	// the value corresponds to
	if rSchema.ResourceProperty(propName).IsArrayProperty() {
		return false
	}

	// e.g. aws-native:s3:StorageLens has a storageLensConfiguration/StorageLensArn readOnly property
	if strings.Contains(propName, "/") {
		parts := strings.Split(propName, "/")
		name := parts[len(parts)-1]
		// If the property is a nested property then only consider it stable if
		// the property contains the resource type name. There are a lot of cases where
		// an object has a property called `id` or `arn` that is not  the resource id or arn
		return slices.Contains(stableOutputsNameOnly, name)
	}
	return slices.Contains(topLevelStableOutputs, propName)
}

func isReadOnly(readOnly []string, key string) bool {
	for _, prop := range readOnly {
		if prop == key {
			return true
		}
	}
	return false
}

// filterAndReturnNested will filter the readOnly nested properties and
// return the nested properties if found
// e.g.
//   - if the readOnly properties are ["foo/bar", "foo/bar/baz", "somethingElse"]
//     and the key is "foo"
//     then the result will be ["bar", "bar/baz"]
func filterAndReturnNested(readOnly []string, key string) []string {
	result := []string{}
	for _, prop := range readOnly {
		if strings.HasPrefix(prop, key+"/*/") {
			result = append(result, strings.TrimPrefix(prop, key+"/*/"))
		} else if strings.HasPrefix(prop, key+"/") {
			result = append(result, strings.TrimPrefix(prop, key+"/"))
		}
	}
	return result
}

// previewOutputValue exists to recurse through nested objects and populate computed values
// There are cases where the input and output types are the same, but some properties of the type
// are only output values.
func previewOutputValue(
	inputValue resource.PropertyValue,
	types map[string]metadata.CloudAPIType,
	prop *schema.TypeSpec,
	readOnly []string,
) resource.PropertyValue {
	switch {
	case inputValue.IsSecret():
		return resource.NewSecretProperty(&resource.Secret{
			Element: previewOutputValue(inputValue.SecretValue().Element, types, prop, readOnly),
		})
	case inputValue.IsOutput():
		return resource.NewOutputProperty(resource.Output{
			Element:      previewOutputValue(inputValue.OutputValue().Element, types, prop, readOnly),
			Known:        inputValue.OutputValue().Known,
			Secret:       inputValue.OutputValue().Secret,
			Dependencies: inputValue.OutputValue().Dependencies,
		})
	case inputValue.IsArray() && (prop.Type == "array" || prop.Items != nil):
		var items []resource.PropertyValue
		for _, item := range inputValue.ArrayValue() {
			items = append(items, previewOutputValue(item, types, prop.Items, readOnly))
		}
		return resource.NewArrayProperty(items)
	case inputValue.IsObject() && strings.HasPrefix(prop.Ref, "#/types/"):
		typeName := strings.TrimPrefix(prop.Ref, "#/types/")
		if t, ok := types[typeName]; ok {
			v := previewResourceOutputs(inputValue.ObjectValue(), types, t.Properties, readOnly)
			return resource.NewObjectProperty(v)
		}
	case inputValue.IsObject() && prop.AdditionalProperties != nil && (prop.AdditionalProperties.Ref != "" || prop.AdditionalProperties.Items != nil):
		return previewOutputValue(inputValue, types, prop.AdditionalProperties, readOnly)
	case inputValue.IsObject() && prop.AdditionalProperties != nil && prop.AdditionalProperties.Ref == "":
		inputObject := inputValue.ObjectValue()
		result := resource.PropertyMap{}
		for name, value := range inputObject {
			p := value
			result[name] = previewOutputValue(p, types, prop.AdditionalProperties, readOnly)
		}
		return resource.NewObjectProperty(result)
	}
	// fallback to assuming input is the same as the output
	return inputValue
}
