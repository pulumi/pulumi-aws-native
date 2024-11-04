// Copyright 2024, Pulumi Corporation.

package resources

import (
	"context"
	"time"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

type CustomResource interface {
	// Check validates and transforms the inputs of the resource.
	Check(ctx context.Context, urn resource.URN, randomSeed []byte, inputs, state resource.PropertyMap, defaultTags map[string]string) (resource.PropertyMap, []ValidationFailure, error)
	// Create creates a new resource in the cloud provider and returns its unique identifier and outputs.
	Create(ctx context.Context, urn resource.URN, inputs resource.PropertyMap, timeout time.Duration) (identifier *string, outputs resource.PropertyMap, err error)
	// Read returns the outputs and the updated inputs of the resource.
	Read(ctx context.Context, urn resource.URN, id string, oldInputs, oldOutputs resource.PropertyMap) (outputs resource.PropertyMap, inputs resource.PropertyMap, exists bool, err error)
	// Update applies the diff of the inputs to the resource and returns the updated outputs.
	Update(ctx context.Context, urn resource.URN, id string, inputs, oldInputs resource.PropertyMap, timeout time.Duration) (resource.PropertyMap, error)
	// Delete removes the resource from the cloud provider.
	Delete(ctx context.Context, urn resource.URN, id string, inputs resource.PropertyMap, timeout time.Duration) error
}
