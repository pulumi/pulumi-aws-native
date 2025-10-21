// Copyright 2024, Pulumi Corporation.

package resources

import (
	"context"
	"fmt"
	"time"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/autonaming"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/client"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/default_tags"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi-go-provider/resourcex"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

type AutoNaming struct {
	PropertyName string
	MinLength    int
	MaxLength    int
}

type ExtensionResourceInputs struct {
	Type         string
	Properties   map[string]any
	CreateOnly   []string
	WriteOnly    []string
	TagsProperty string
	TagsStyle    default_tags.TagsStyle
	AutoNaming   *AutoNaming
}

type extensionResource struct {
	client client.CloudControlClient
}

// Check extensionResource implements CustomResource
var _ CustomResource = (*extensionResource)(nil)

func NewExtensionResource(client client.CloudControlClient) *extensionResource {
	return &extensionResource{
		client: client,
	}
}

func (r *extensionResource) Check(ctx context.Context, urn resource.URN,
	engineAutonaming autonaming.EngineAutoNamingConfig, inputs, state resource.PropertyMap,
	defaultTags map[string]string) (resource.PropertyMap, []ValidationFailure, error) {
	var typedInputs ExtensionResourceInputs
	_, err := resourcex.Unmarshal(&typedInputs, inputs, resourcex.UnmarshalOptions{})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal inputs: %w", err)
	}

	var failures []ValidationFailure
	typedInputs, failures = ApplyDefaults(typedInputs)

	if len(failures) > 0 {
		return nil, failures, nil
	}

	if typedInputs.AutoNaming != nil {
		if err = autonaming.ApplyAutoNaming(&metadata.AutoNamingSpec{
			SdkName:   typedInputs.AutoNaming.PropertyName,
			MinLength: typedInputs.AutoNaming.MinLength,
			MaxLength: typedInputs.AutoNaming.MaxLength,
		}, urn, engineAutonaming, nil, state, inputs); err != nil {
			return nil, nil, fmt.Errorf("failed to apply auto-naming: %w", err)
		}
	}

	// Merge default tags into the inputs if the resource supports tags and the user has not overridden them.
	if len(defaultTags) > 0 && typedInputs.TagsProperty != "" && typedInputs.TagsStyle != default_tags.TagsStyleUnknown {
		tagsKey := resource.PropertyKey(typedInputs.TagsProperty)
		inputProperties := resource.NewPropertyMapFromMap(typedInputs.Properties)
		tagsStyle := typedInputs.TagsStyle
		// When sending raw payloads to the CloudControl API we need to use the upper case version of the tags.
		if tagsStyle.IsKeyValueArray() {
			tagsStyle = default_tags.TagsStyleKeyValueArrayUpperCase
		}
		val, err := default_tags.MergeDefaultTags(inputProperties[tagsKey], defaultTags, tagsStyle)
		if err != nil {
			return nil, nil, err
		}
		inputProperties[tagsKey] = val
		inputs[resource.PropertyKey("properties")] = resource.PropertyValue{V: inputProperties}
	}

	return inputs, nil, nil
}

func ApplyDefaults(typedInputs ExtensionResourceInputs) (ExtensionResourceInputs, []ValidationFailure) {
	if !typedInputs.TagsStyle.IsValid() {
		return typedInputs, []ValidationFailure{
			{
				Path:   "tagsStyle",
				Reason: "tagsStyle is invalid, must be one of: stringMap, keyValueArray, none",
			},
		}
	}

	// If the tags style is unknown (unset) then we'll default to keyValueArray which is the most common style.
	if typedInputs.TagsProperty != "" && typedInputs.TagsStyle == default_tags.TagsStyleUnknown {
		typedInputs.TagsStyle = default_tags.TagsStyleKeyValueArray
	}

	// If the tags property is unset then we'll default to "Tags" which is the most common property name.
	// Except if the tags style is explicitly set to none, in which case we'll ignore the tags property.
	if typedInputs.TagsProperty == "" && typedInputs.TagsStyle != default_tags.TagsStyleNone && typedInputs.TagsStyle != default_tags.TagsStyleUnknown {
		typedInputs.TagsProperty = "Tags"
	}

	return typedInputs, nil
}

func (r *extensionResource) Create(ctx context.Context, urn resource.URN, inputs resource.PropertyMap, timeout time.Duration) (identifier *string, outputs resource.PropertyMap, err error) {
	var typedInputs ExtensionResourceInputs
	_, err = resourcex.Unmarshal(&typedInputs, inputs, resourcex.UnmarshalOptions{})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal inputs: %w", err)
	}

	id, resourceState, err := r.client.Create(ctx, urn, typedInputs.Type, typedInputs.Properties)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create resource: %w", err)
	}

	rawOutputs := typedInputs.toOutputs(resourceState)
	return id, CheckpointObject(inputs, rawOutputs), nil
}

func (r *extensionResource) Read(ctx context.Context, urn resource.URN, id string, oldInputs, oldState resource.PropertyMap) (outputs resource.PropertyMap, inputs resource.PropertyMap, exists bool, err error) {
	if len(oldState) == 0 {
		// We can't yet support import because the type would not be known.
		return nil, nil, false, fmt.Errorf("ExtensionResource import not implemented")
	}
	var typedInputs ExtensionResourceInputs
	_, err = resourcex.Unmarshal(&typedInputs, oldInputs, resourcex.UnmarshalOptions{})
	if err != nil {
		return nil, nil, false, fmt.Errorf("failed to unmarshal old inputs: %w", err)
	}

	var oldOutputsMap resource.PropertyMap
	if oldOutputsValue, ok := oldState[resource.PropertyKey("outputs")]; ok && oldOutputsValue.IsObject() {
		oldOutputsMap = oldOutputsValue.ObjectValue()
	}

	resourceState, exists, err := r.client.Read(ctx, typedInputs.Type, id)
	if err != nil {
		return nil, nil, false, fmt.Errorf("failed to read resource: %w", err)
	}
	if !exists {
		return nil, nil, false, nil
	}

	// Re-add write-only properties before calculating the inputs.
	resourceState = typedInputs.AddWriteOnlyProps(resourceState)
	oldProperties := resource.NewPropertyMapFromMap(typedInputs.Properties)
	newOutputs := resource.NewPropertyMapFromMap(resourceState)
	outputsDiff := oldOutputsMap.Diff(newOutputs)
	newProperties := ApplyDiff(oldProperties, outputsDiff)

	typedInputs.Properties = resourcex.Decode(newProperties)
	newInputs := oldInputs.Copy()
	newInputs[resource.PropertyKey("properties")] = resource.PropertyValue{V: newProperties}
	rawState := typedInputs.toOutputs(resourceState)
	return CheckpointObject(newInputs, rawState), newInputs, true, nil
}

func (r *extensionResource) Update(ctx context.Context, urn resource.URN, id string, inputs, oldInputs, state resource.PropertyMap, timeout time.Duration) (resource.PropertyMap, error) {
	var typedOldInputs ExtensionResourceInputs
	_, err := resourcex.Unmarshal(&typedOldInputs, oldInputs, resourcex.UnmarshalOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal old inputs: %w", err)
	}

	var typedInputs ExtensionResourceInputs
	_, err = resourcex.Unmarshal(&typedInputs, inputs, resourcex.UnmarshalOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal new inputs: %w", err)
	}

	if typedOldInputs.Type != typedInputs.Type {
		return nil, fmt.Errorf("changing the type of an extension resource is not supported")
	}

	jsonPatch, err := CalculateUntypedPatch(typedOldInputs, typedInputs)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate patch: %w", err)
	}

	resourceState, err := r.client.Update(ctx, urn, typedInputs.Type, id, jsonPatch)
	if err != nil {
		return nil, fmt.Errorf("failed to update resource: %w", err)
	}

	rawState := typedInputs.toOutputs(resourceState)
	return CheckpointObject(inputs, rawState), nil
}

func (r *extensionResource) Delete(ctx context.Context, urn resource.URN, id string, inputs, state resource.PropertyMap, timeout time.Duration) error {
	var typedInputs ExtensionResourceInputs
	_, err := resourcex.Unmarshal(&typedInputs, inputs, resourcex.UnmarshalOptions{})
	if err != nil {
		return fmt.Errorf("failed to unmarshal inputs: %w", err)
	}

	err = r.client.Delete(ctx, urn, typedInputs.Type, id)
	if err != nil {
		return fmt.Errorf("failed to delete resource: %w", err)
	}

	return nil
}

const AutoNamingTypeToken = "aws-native:index:AutoNaming"

func AutoNamingTypeSpec() pschema.ObjectTypeSpec {
	return pschema.ObjectTypeSpec{
		Description: "Auto-naming specification for the resource.",
		Type:        "object",
		Properties: map[string]pschema.PropertySpec{
			"propertyName": {
				Description: "The name of the property in the Cloud Control payload that is used to set the name of the resource.",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"minLength": {
				Description: "The minimum length of the name.",
				TypeSpec:    pschema.TypeSpec{Type: "integer"},
			},
			"maxLength": {
				Description: "The maximum length of the name.",
				TypeSpec:    pschema.TypeSpec{Type: "integer"},
			},
		},
	}
}

func ExtensionResourceSpec() pschema.ResourceSpec {
	return pschema.ResourceSpec{
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: "A special resource that enables deploying CloudFormation Extensions (third-party resources). An extension has to be pre-registered in your AWS account in order to use this resource.",
			Properties: map[string]pschema.PropertySpec{
				"outputs": {
					Description: "Dictionary of the extension resource attributes.",
					TypeSpec: pschema.TypeSpec{
						Type: "object",
						AdditionalProperties: &pschema.TypeSpec{
							Ref: "pulumi.json#/Any",
						},
					},
				},
			},
			Required: []string{"outputs"},
		},
		InputProperties: map[string]pschema.PropertySpec{
			"type": {
				Description: "CloudFormation type name. This has three parts, each separated by two colons. For AWS resources this starts with `AWS::` e.g. `AWS::Logs::LogGroup`. Third party resources should use a namespace prefix e.g. `MyCompany::MyService::MyResource`.",
				TypeSpec: pschema.TypeSpec{
					Type: "string",
				},
				ReplaceOnChanges: true,
			},
			"properties": {
				Description: "Property bag containing the properties for the resource. These should be defined using the casing expected by the CloudControl API as these values are sent exact as provided.",
				TypeSpec: pschema.TypeSpec{
					Type: "object",
					AdditionalProperties: &pschema.TypeSpec{
						Ref: "pulumi.json#/Any",
					},
				},
			},
			"createOnly": {
				Description: "Property names as defined by `createOnlyProperties` in the CloudFormation schema. Create-only properties can't be set during updates, so will not be included in patches even if they are also marked as write-only, and will cause an error if attempted to be updated. Therefore any property here should also be included in the `replaceOnChanges` resource option too.\nIn the CloudFormation schema these are fully qualified property paths (e.g. `/properties/AccessToken`) whereas here we only include the top-level property name (e.g. `AccessToken`).",
				TypeSpec: pschema.TypeSpec{
					Type: "array",
					Items: &pschema.TypeSpec{
						Type: "string",
					},
				},
			},
			"writeOnly": {
				Description: "Property names as defined by `writeOnlyProperties` in the CloudFormation schema. Write-only properties are not returned during read operations and have to be included in all update operations as CloudControl itself can't read their previous values.\nIn the CloudFormation schema these are fully qualified property paths (e.g. `/properties/AccessToken`) whereas here we only include the top-level property name (e.g. `AccessToken`).",
				TypeSpec: pschema.TypeSpec{
					Type: "array",
					Items: &pschema.TypeSpec{
						Type: "string",
					},
				},
			},
			"tagsProperty": {
				Description: "Optional name of the property containing the tags. Defaults to \"Tags\" if the `tagsStyle` is set to either \"stringMap\" or \"keyValueArray\". This is used to apply default tags to the resource and can be ignored if not using default tags.",
				TypeSpec: pschema.TypeSpec{
					Type: "string",
				},
			},
			"tagsStyle": {
				Description: "Optional style of tags this resource uses. Valid values are \"stringMap\", \"keyValueArray\" or \"none\". Defaults to `keyValueArray` if `tagsProperty` is set. This is used to apply default tags to the resource and can be ignored if not using default tags.",
				TypeSpec: pschema.TypeSpec{
					Type: "string",
				},
			},
			"autoNaming": {
				Description: "Optional auto-naming specification for the resource.\nIf provided and the name is not specified manually, the provider will automatically generate a name based on the Pulumi resource name and a random suffix.",
				TypeSpec:    pschema.TypeSpec{Ref: fmt.Sprint("#/types/", AutoNamingTypeToken)},
			},
		},
		RequiredInputs: []string{"type", "properties"},
	}
}

func (r *ExtensionResourceInputs) toOutputs(resourceState map[string]interface{}) map[string]any {
	return map[string]interface{}{
		"outputs": r.AddWriteOnlyProps(resourceState),
	}
}

// AddWriteOnlyProps fills in the write-only properties from the old state as they won't be returned when reading.
func (r *ExtensionResourceInputs) AddWriteOnlyProps(resourceState map[string]interface{}) map[string]interface{} {
	if len(r.WriteOnly) == 0 {
		return resourceState
	}
	appended := make(map[string]interface{}, len(resourceState)+len(r.WriteOnly))
	for k, v := range resourceState {
		appended[k] = v
	}
	for _, writeOnlyProp := range r.WriteOnly {
		if _, ok := appended[writeOnlyProp]; !ok {
			oldValue, ok := r.Properties[writeOnlyProp]
			if ok {
				appended[writeOnlyProp] = oldValue
			}
		}
	}
	return appended
}

// extension resources have outputs returned in an "outputs" property
// since the outputs can be arbitrary we just mark the entire thing as computed
func (r *extensionResource) PreviewCustomResourceOutputs() resource.PropertyMap {
	return resource.PropertyMap{
		"outputs": resource.MakeComputed(resource.NewStringProperty("")),
	}
}
