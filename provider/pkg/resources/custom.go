// Copyright 2024, Pulumi Corporation.

package resources

import (
	"context"
	"time"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/autonaming"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

//go:generate mockgen -package resources -source custom.go -destination mock_custom_resource.go CustomResource
type CustomResource interface {
	// Check validates and transforms the inputs of the resource.
	Check(ctx context.Context, urn resource.URN, engineAutonaming autonaming.EngineAutoNamingConfig,
		inputs, state resource.PropertyMap, defaultTags map[string]string) (resource.PropertyMap, []ValidationFailure, error)
	// Create creates a new resource in the cloud provider and returns its unique identifier and outputs.
	Create(ctx context.Context, urn resource.URN, inputs resource.PropertyMap, timeout time.Duration) (identifier *string, outputs resource.PropertyMap, err error)
	// Read returns the outputs and the updated inputs of the resource.
	Read(ctx context.Context, urn resource.URN, id string, oldInputs, oldOutputs resource.PropertyMap) (outputs resource.PropertyMap, inputs resource.PropertyMap, exists bool, err error)
	// Update applies the diff of the inputs to the resource and returns the updated outputs.
	Update(ctx context.Context, urn resource.URN, id string, inputs, oldInputs, state resource.PropertyMap, timeout time.Duration) (resource.PropertyMap, error)
	// Delete removes the resource from the cloud provider.
	Delete(ctx context.Context, urn resource.URN, id string, inputs, state resource.PropertyMap, timeout time.Duration) error
	// PreviewCustomResourceOutputs returns the outputs of the resource based on the inputs and the output properties in the resource schema.
	PreviewCustomResourceOutputs() resource.PropertyMap
}
