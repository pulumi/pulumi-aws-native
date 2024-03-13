// Copyright 2016-2024, Pulumi Corporation.

package awsclient

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

// CloudControlClient is an interface that wraps around the AWS Cloud Control API.
type CloudControlClient interface {
	// CreateResource creates a resource of the specified type with the desired state.
	// It returns a ProgressEvent which is the initial progress returned directly from the API call,
	// without awaiting any long-running operations.
	CreateResource(ctx context.Context, cfType, desiredState string) (*types.ProgressEvent, error)

	// UpdateResource updates a resource of the specified type with the desired state.
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
}

// NewCloudControlClient creates a wrapper around the AWS SDK's Cloud Control client.
func NewCloudControlClient(cctl *cloudcontrol.Client, roleArn *string) CloudControlClient {
	return &ccClientImpl{
		cctl:    cctl,
		roleArn: roleArn,
	}
}

type ccClientImpl struct {
	cctl    *cloudcontrol.Client
	roleArn *string
}

func (client *ccClientImpl) CreateResource(ctx context.Context, cfType, desiredState string) (*types.ProgressEvent, error) {
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

func (client *ccClientImpl) UpdateResource(ctx context.Context, cfType, id string, patches []jsonpatch.JsonPatchOperation) (*types.ProgressEvent, error) {
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

func (client *ccClientImpl) DeleteResource(ctx context.Context, cfType, id string) (*types.ProgressEvent, error) {
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

func (client *ccClientImpl) GetResource(ctx context.Context, typeName, identifier string) (map[string]interface{}, error) {
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

func (client *ccClientImpl) GetResourceRequestStatus(ctx context.Context, requestToken string) (*types.ProgressEvent, error) {
	res, err := client.cctl.GetResourceRequestStatus(ctx, &cloudcontrol.GetResourceRequestStatusInput{
		RequestToken: aws.String(requestToken),
	})
	if err != nil {
		return nil, err
	}
	return res.ProgressEvent, nil
}
