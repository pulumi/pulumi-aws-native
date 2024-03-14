// Copyright 2016-2024, Pulumi Corporation.

package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/mattbaird/jsonpatch"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/schema"
)

type Client interface {
	Create(ctx context.Context, typeName string, desiredState map[string]any) (identifier *string, resourceState map[string]any, err error)
	Read(ctx context.Context, typeName, identifier string) (map[string]interface{}, error)
	Update(ctx context.Context, typeName, identifier string, patches []jsonpatch.JsonPatchOperation) (map[string]interface{}, error)
	Delete(ctx context.Context, typeName, identifier string) error
}

type clientImpl struct {
	api     CloudControlApiClient
	awaiter CloudControlAwaiter
}

func NewClient(cctl *cloudcontrol.Client, roleArn *string) Client {
	api := NewCloudControlApiClient(cctl, roleArn)
	return &clientImpl{
		api:     api,
		awaiter: NewCloudControlAwaiter(api),
	}
}

func (c *clientImpl) Create(ctx context.Context, typeName string, desiredState map[string]any) (identifier *string, resourceState map[string]any, err error) {
	// Serialize inputs as a desired state JSON.
	jsonBytes, err := json.Marshal(desiredState)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal as JSON: %w", err)
	}
	payload := string(jsonBytes)

	res, err := c.api.CreateResource(ctx, typeName, payload)
	if err != nil {
		return nil, nil, fmt.Errorf("creating resource: %w", err)
	}

	pi, waitErr := c.awaiter.WaitForResourceOpCompletion(ctx, res)

	// Read the state - even if there was a creation error but the progress event contains a resource ID.
	var id string
	var outputs map[string]interface{}
	var readErr error
	if pi != nil && pi.Identifier != nil {
		// Retrieve the resource state from AWS.
		// Note that we do so even if creation hasn't succeeded but the identifier is assigned.
		id = *pi.Identifier
		resourceState, err := c.api.GetResource(ctx, typeName, id)
		if err != nil {
			readErr = fmt.Errorf("reading resource state: %w", err)
		} else {
			outputs = schema.CfnToSdk(resourceState)
		}
	}

	if waitErr != nil {
		if id == "" {
			return nil, nil, waitErr
		}

		if readErr != nil {
			return nil, nil, fmt.Errorf("resource partially created but read failed. read error: %v, create error: %w", readErr, waitErr)
		}

		// Resource was created but failed to fully initialize.
		// If it has some state, return a partial error.
		return &id, outputs, waitErr
	}
	if pi.Identifier == nil {
		return nil, nil, errors.New("received nil identifier while reading resource state")
	}
	if readErr != nil {
		return nil, nil, fmt.Errorf("reading resource state: %w", readErr)
	}

	return &id, outputs, nil
}

func (c *clientImpl) Read(ctx context.Context, typeName, identifier string) (map[string]interface{}, error) {
	return c.api.GetResource(ctx, typeName, identifier)
}

func (c *clientImpl) Update(ctx context.Context, typeName, identifier string, patches []jsonpatch.JsonPatchOperation) (map[string]interface{}, error) {
	res, err := c.api.UpdateResource(ctx, typeName, identifier, patches)
	if err != nil {
		return nil, err
	}

	if _, err = c.awaiter.WaitForResourceOpCompletion(ctx, res); err != nil {
		return nil, err
	}

	resourceState, err := c.api.GetResource(ctx, typeName, identifier)
	if err != nil {
		return nil, fmt.Errorf("reading resource state: %w", err)
	}

	return resourceState, nil
}

func (c *clientImpl) Delete(ctx context.Context, typeName, identifier string) error {
	res, err := c.api.DeleteResource(ctx, typeName, identifier)
	if err != nil {
		return err
	}

	pi, err := c.awaiter.WaitForResourceOpCompletion(ctx, res)

	if err != nil && pi != nil {
		errorCode := pi.ErrorCode
		if errorCode == types.HandlerErrorCodeNotFound {
			// NotFound means that the resource was already deleted, so the operation can succeed.
			return nil
		}
	}

	return err
}
