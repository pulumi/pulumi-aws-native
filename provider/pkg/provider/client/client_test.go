// Copyright 2016-2024, Pulumi Corporation.

package client

import (
	"context"
	"errors"
	"net/http"
	"testing"

	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/aws/smithy-go"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/mattbaird/jsonpatch"
	"github.com/stretchr/testify/assert"
)

func TestClientRead(t *testing.T) {
	ctx := context.TODO()
	typeName := "exampleType"
	identifier := "exampleIdentifier"

	// Mock API implementation
	mockAPI := &mockAPI{}
	client := &clientImpl{
		api: mockAPI,
	}

	t.Run("Resource found", func(t *testing.T) {
		resourceState := map[string]interface{}{"key": "value"}
		mockAPI.GetResourceFunc = func(ctx context.Context, typeName, identifier string) (map[string]interface{}, error) {
			return resourceState, nil
		}

		state, exists, err := client.Read(ctx, typeName, identifier)

		assert.NoError(t, err)
		assert.True(t, exists)
		assert.Equal(t, resourceState, state)
	})

	t.Run("Resource not found 404", func(t *testing.T) {
		mockAPI.GetResourceFunc = func(ctx context.Context, typeName, identifier string) (map[string]interface{}, error) {
			return nil, &smithy.OperationError{
				Err: &awshttp.ResponseError{
					ResponseError: &smithyhttp.ResponseError{
						Response: &smithyhttp.Response{
							Response: &http.Response{
								StatusCode: 404,
							},
						},
					},
				},
			}
		}

		state, exists, err := client.Read(ctx, typeName, identifier)

		assert.NoError(t, err)
		assert.False(t, exists)
		assert.Nil(t, state)
	})

	t.Run("Resource not found ResourceNotFoundException", func(t *testing.T) {
		mockAPI.GetResourceFunc = func(ctx context.Context, typeName, identifier string) (map[string]interface{}, error) {
			return nil, &smithy.OperationError{
				Err: &awshttp.ResponseError{
					ResponseError: &smithyhttp.ResponseError{
						Response: &smithyhttp.Response{
							Response: &http.Response{
								StatusCode: 400,
							},
						},
						Err: errors.New("Oh no ResourceNotFoundException happened!"),
					},
				},
			}
		}

		state, exists, err := client.Read(ctx, typeName, identifier)

		assert.NoError(t, err)
		assert.False(t, exists)
		assert.Nil(t, state)
	})

	t.Run("Other error", func(t *testing.T) {
		expectedErr := errors.New("some error")
		mockAPI.GetResourceFunc = func(ctx context.Context, typeName, identifier string) (map[string]interface{}, error) {
			return nil, expectedErr
		}

		state, exists, err := client.Read(ctx, typeName, identifier)

		assert.Equal(t, expectedErr, err)
		assert.False(t, exists)
		assert.Nil(t, state)
	})
}

// Mock API implementation
type mockAPI struct {
	GetResourceFunc func(ctx context.Context, typeName, identifier string) (map[string]interface{}, error)
}

func (m *mockAPI) GetResource(ctx context.Context, typeName, identifier string) (map[string]interface{}, error) {
	return m.GetResourceFunc(ctx, typeName, identifier)
}

func (m *mockAPI) CreateResource(ctx context.Context, cfType, desiredState string) (*types.ProgressEvent, error) {
	panic("not implemented")
}

// UpdateResource updates a resource of the specified type with the specified changeset.
// It returns a ProgressEvent which is the initial progress returned directly from the API call,
// without awaiting any long-running operations.
// The changes to be applied are expressed as a list of JSON patch operations.
func (m *mockAPI) UpdateResource(ctx context.Context, cfType, id string, patches []jsonpatch.JsonPatchOperation) (*types.ProgressEvent, error) {
	panic("not implemented")
}

// DeleteResource deletes a resource of the specified type with the given identifier.
// It returns a ProgressEvent which is the initial progress returned directly from the API call,
// without awaiting any long-running operations.
func (m *mockAPI) DeleteResource(ctx context.Context, cfType, id string) (*types.ProgressEvent, error) {
	panic("not implemented")
}

// GetResourceRequestStatus returns the current status of a resource operation request.
func (m *mockAPI) GetResourceRequestStatus(ctx context.Context, requestToken string) (*types.ProgressEvent, error) {
	panic("not implemented")
}
