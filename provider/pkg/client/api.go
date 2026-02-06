// Copyright 2016-2024, Pulumi Corporation.

package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
	"github.com/mattbaird/jsonpatch"
)

// CloudControlApiClient is an interface that wraps around the AWS Cloud Control API.
type CloudControlApiClient interface {
	// CreateResource creates a resource of the specified type with the desired state.
	// It returns a ProgressEvent which is the initial progress returned directly from the API call,
	// without awaiting any long-running operations.
	CreateResource(ctx context.Context, cfType, desiredState string) (*types.ProgressEvent, error)

	// UpdateResource updates a resource of the specified type with the specified changeset.
	// It returns a ProgressEvent which is the initial progress returned directly from the API call,
	// without awaiting any long-running operations.
	// The changes to be applied are expressed as a list of JSON patch operations.
	UpdateResource(ctx context.Context, cfType, id string, patches []jsonpatch.JsonPatchOperation) (*types.ProgressEvent, error)

	// DeleteResource deletes a resource of the specified type with the given identifier.
	// It returns a ProgressEvent which is the initial progress returned directly from the API call,
	// without awaiting any long-running operations.
	DeleteResource(ctx context.Context, cfType, id string) (*types.ProgressEvent, error)

	// GetResource returns information about the current state of the specified resource. It deserializes
	// the response from the service into a map of untyped values.
	GetResource(ctx context.Context, typeName, identifier string) (map[string]interface{}, error)

	// GetResourceRequestStatus returns the current status of a resource operation request.
	GetResourceRequestStatus(ctx context.Context, requestToken string) (*types.ProgressEvent, error)

	// GetResourceRequestStatusWithHooks returns the current status of a resource operation request with hook information.
	GetResourceRequestStatusWithHooks(ctx context.Context, requestToken string) (*types.ProgressEvent, []types.HookProgressEvent, error)
}

// NewCloudControlApiClient creates a wrapper around the AWS SDK's Cloud Control client.
func NewCloudControlApiClient(cctl *cloudcontrol.Client, roleArn *string) CloudControlApiClient {
	return &ccApiClientImpl{
		cctl:    cctl,
		roleArn: roleArn,
	}
}

type ccApiClientImpl struct {
	cctl    *cloudcontrol.Client
	roleArn *string
}

func (client *ccApiClientImpl) CreateResource(ctx context.Context, cfType, desiredState string) (*types.ProgressEvent, error) {
	clientToken := uuid.New().String()
	res, err := client.cctl.CreateResource(ctx, &cloudcontrol.CreateResourceInput{
		ClientToken:  aws.String(clientToken),
		TypeName:     aws.String(cfType),
		DesiredState: aws.String(desiredState),
		RoleArn:      client.roleArn,
	})
	if err != nil {
		return nil, err
	}
	return res.ProgressEvent, nil
}

func (client *ccApiClientImpl) UpdateResource(ctx context.Context, cfType, id string, patches []jsonpatch.JsonPatchOperation) (*types.ProgressEvent, error) {
	doc, err := json.Marshal(patches)
	if err != nil {
		return nil, fmt.Errorf("serializing patch as json: %w", err)
	}

	docAsString := string(doc)
	clientToken := uuid.New().String()
	res, err := client.cctl.UpdateResource(ctx, &cloudcontrol.UpdateResourceInput{
		ClientToken:   aws.String(clientToken),
		TypeName:      aws.String(cfType),
		Identifier:    aws.String(id),
		PatchDocument: &docAsString,
		RoleArn:       client.roleArn,
	})
	if err != nil {
		return nil, err
	}
	return res.ProgressEvent, nil
}

func (client *ccApiClientImpl) DeleteResource(ctx context.Context, cfType, id string) (*types.ProgressEvent, error) {
	clientToken := uuid.New().String()
	res, err := client.cctl.DeleteResource(ctx, &cloudcontrol.DeleteResourceInput{
		ClientToken: aws.String(clientToken),
		TypeName:    aws.String(cfType),
		Identifier:  aws.String(id),
		RoleArn:     client.roleArn,
	})
	if err != nil {
		return nil, err
	}
	return res.ProgressEvent, nil
}

func (client *ccApiClientImpl) GetResource(ctx context.Context, typeName, identifier string) (map[string]interface{}, error) {
	getRes, err := client.cctl.GetResource(ctx, &cloudcontrol.GetResourceInput{
		TypeName:   aws.String(typeName),
		Identifier: aws.String(identifier),
		RoleArn:    client.roleArn,
	})
	if err != nil {
		return nil, err
	}
	if getRes.ResourceDescription.Properties == nil {
		return nil, errors.New("received nil properties")
	}

	properties := *getRes.ResourceDescription.Properties
	var outputs map[string]interface{}
	if err = json.Unmarshal([]byte(properties), &outputs); err != nil {
		return nil, err
	}

	return outputs, nil
}

func (client *ccApiClientImpl) GetResourceRequestStatus(ctx context.Context, requestToken string) (*types.ProgressEvent, error) {
	res, err := client.cctl.GetResourceRequestStatus(ctx, &cloudcontrol.GetResourceRequestStatusInput{
		RequestToken: aws.String(requestToken),
	})
	if err != nil {
		return nil, err
	}
	return res.ProgressEvent, nil
}

func (client *ccApiClientImpl) GetResourceRequestStatusWithHooks(ctx context.Context, requestToken string) (*types.ProgressEvent, []types.HookProgressEvent, error) {
	res, err := client.cctl.GetResourceRequestStatus(ctx, &cloudcontrol.GetResourceRequestStatusInput{
		RequestToken: aws.String(requestToken),
	})
	if err != nil {
		return nil, nil, err
	}
	return res.ProgressEvent, res.HooksProgressEvent, nil
}
