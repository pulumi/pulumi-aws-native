// Copyright 2016-2024, Pulumi Corporation.

package client

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/aws/smithy-go"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/mattbaird/jsonpatch"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testURN = resource.URN("urn:pulumi:stack::project::type::name")

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

func TestClientCreate(t *testing.T) {
	ctx := context.TODO()
	typeName := "exampleType"
	desiredState := map[string]interface{}{"input1": "value1"}

	// Mock API implementation
	mockAPI := &mockAPI{}
	client := &clientImpl{
		api:     mockAPI,
		awaiter: mockAPI,
		retryer: func() aws.Retryer { return aws.NopRetryer{} },
	}

	t.Run("Resource creation success", func(t *testing.T) {
		resourceID := "exampleID"
		resourceState := map[string]interface{}{"output1": "outvalue1"}
		mockAPI.CreateResourceFunc = func(ctx context.Context, cfType, desiredState string) (*types.ProgressEvent, error) {
			return &types.ProgressEvent{
				OperationStatus: types.OperationStatusSuccess,
				Identifier:      &resourceID,
			}, nil
		}
		mockAPI.WaitForResourceOpCompletionFunc = func(ctx context.Context, pi *types.ProgressEvent) (*types.ProgressEvent, error) {
			return pi, nil
		}
		mockAPI.GetResourceFunc = func(ctx context.Context, typeName, identifier string) (map[string]interface{}, error) {
			if identifier != resourceID {
				return nil, errors.New("unexpected identifier")
			}
			return resourceState, nil
		}

		id, outputs, err := client.Create(ctx, testURN, typeName, desiredState)

		assert.NoError(t, err)
		assert.Equal(t, &resourceID, id)
		assert.Equal(t, resourceState, outputs)
	})

	t.Run("Resource creation failure", func(t *testing.T) {
		mockAPI.CreateResourceFunc = func(ctx context.Context, cfType, desiredState string) (*types.ProgressEvent, error) {
			return nil, errors.New("creation failed")
		}

		id, outputs, err := client.Create(ctx, testURN, typeName, desiredState)

		assert.Equal(t, "creating resource: creation failed", err.Error())
		assert.Nil(t, id)
		assert.Nil(t, outputs)
	})

	t.Run("Resource await failure no state", func(t *testing.T) {
		resourceID := "exampleID"
		mockAPI.CreateResourceFunc = func(ctx context.Context, cfType, desiredState string) (*types.ProgressEvent, error) {
			return &types.ProgressEvent{
				OperationStatus: types.OperationStatusSuccess,
				Identifier:      &resourceID,
			}, nil
		}
		mockAPI.WaitForResourceOpCompletionFunc = func(ctx context.Context, pi *types.ProgressEvent) (*types.ProgressEvent, error) {
			return nil, errors.New("creation failed")
		}

		id, outputs, err := client.Create(ctx, testURN, typeName, desiredState)

		assert.Equal(t, "creating resource: creation failed", err.Error())
		assert.Nil(t, id)
		assert.Nil(t, outputs)
	})

	t.Run("Resource await returns no identifier", func(t *testing.T) {
		resourceID := "exampleID"
		mockAPI.CreateResourceFunc = func(ctx context.Context, cfType, desiredState string) (*types.ProgressEvent, error) {
			return &types.ProgressEvent{
				OperationStatus: types.OperationStatusSuccess,
				Identifier:      &resourceID,
			}, nil
		}
		mockAPI.WaitForResourceOpCompletionFunc = func(ctx context.Context, pi *types.ProgressEvent) (*types.ProgressEvent, error) {
			return &types.ProgressEvent{Identifier: nil}, nil
		}

		id, outputs, err := client.Create(ctx, testURN, typeName, desiredState)

		assert.Equal(t, "received nil identifier while awaiting completion", err.Error())
		assert.Nil(t, id)
		assert.Nil(t, outputs)
	})

	t.Run("Resource await and read errors, but with identifier", func(t *testing.T) {
		resourceID := "exampleID"
		mockAPI.CreateResourceFunc = func(ctx context.Context, cfType, desiredState string) (*types.ProgressEvent, error) {
			return &types.ProgressEvent{
				OperationStatus: types.OperationStatusSuccess,
				Identifier:      &resourceID,
			}, nil
		}
		mockAPI.WaitForResourceOpCompletionFunc = func(ctx context.Context, pi *types.ProgressEvent) (*types.ProgressEvent, error) {
			return &types.ProgressEvent{Identifier: &resourceID}, errors.New("await failed")
		}
		mockAPI.GetResourceFunc = func(ctx context.Context, typeName, identifier string) (map[string]interface{}, error) {
			if identifier != resourceID {
				return nil, errors.New("unexpected identifier")
			}
			return nil, errors.New("read error")
		}

		id, outputs, err := client.Create(ctx, testURN, typeName, desiredState)

		assert.Equal(t, "creating resource (await): await failed", err.Error())
		assert.Nil(t, id)
		assert.Nil(t, outputs)
	})

	t.Run("Resource read error", func(t *testing.T) {
		resourceID := "exampleID"
		mockAPI.CreateResourceFunc = func(ctx context.Context, cfType, desiredState string) (*types.ProgressEvent, error) {
			return &types.ProgressEvent{
				OperationStatus: types.OperationStatusSuccess,
				Identifier:      &resourceID,
			}, nil
		}
		mockAPI.WaitForResourceOpCompletionFunc = func(ctx context.Context, pi *types.ProgressEvent) (*types.ProgressEvent, error) {
			return &types.ProgressEvent{Identifier: &resourceID}, nil
		}
		mockAPI.GetResourceFunc = func(ctx context.Context, typeName, identifier string) (map[string]interface{}, error) {
			if identifier != resourceID {
				return nil, errors.New("unexpected identifier")
			}
			return nil, errors.New("read error")
		}

		id, outputs, err := client.Create(ctx, testURN, typeName, desiredState)

		assert.Equal(t, "reading resource state: read error", err.Error())
		assert.Nil(t, id)
		assert.Nil(t, outputs)
	})

	t.Run("Resource creation succeeds, including outputs, but await returned an error", func(t *testing.T) {
		resourceID := "exampleID"
		resourceState := map[string]interface{}{"output1": "outvalue1"}
		mockAPI.CreateResourceFunc = func(ctx context.Context, cfType, desiredState string) (*types.ProgressEvent, error) {
			return &types.ProgressEvent{
				OperationStatus: types.OperationStatusSuccess,
				Identifier:      &resourceID,
			}, nil
		}
		mockAPI.WaitForResourceOpCompletionFunc = func(ctx context.Context, pi *types.ProgressEvent) (*types.ProgressEvent, error) {
			return &types.ProgressEvent{Identifier: &resourceID}, errors.New("await failed")
		}
		mockAPI.GetResourceFunc = func(ctx context.Context, typeName, identifier string) (map[string]interface{}, error) {
			if identifier != resourceID {
				return nil, errors.New("unexpected identifier")
			}
			return resourceState, nil
		}

		id, outputs, err := client.Create(ctx, testURN, typeName, desiredState)

		assert.Equal(t, "await failed", err.Error())
		assert.Equal(t, &resourceID, id)
		assert.Equal(t, resourceState, outputs)
	})

	t.Run("AlreadyExists returns no resource state despite its existence", func(t *testing.T) {
		resourceID := "exampleID"
		mockAPI.CreateResourceFunc = func(ctx context.Context, cfType, desiredState string) (*types.ProgressEvent, error) {
			return &types.ProgressEvent{
				OperationStatus: types.OperationStatusSuccess,
				Identifier:      &resourceID,
			}, nil
		}
		mockAPI.WaitForResourceOpCompletionFunc = func(ctx context.Context, pi *types.ProgressEvent) (*types.ProgressEvent, error) {
			return &types.ProgressEvent{Identifier: &resourceID, ErrorCode: "AlreadyExists"}, errors.New("resource with same id alteady exists")
		}

		id, outputs, err := client.Create(ctx, testURN, typeName, desiredState)

		assert.Equal(t, "resource with same id alteady exists", err.Error())
		assert.Nil(t, id)
		assert.Nil(t, outputs)
	})

	t.Run("GetResource is retried if it fails with a ResourceNotFoundException", func(t *testing.T) {
		resourceID := "exampleID"
		mockAPI.CreateResourceFunc = func(ctx context.Context, cfType, desiredState string) (*types.ProgressEvent, error) {
			return &types.ProgressEvent{
				OperationStatus: types.OperationStatusSuccess,
				Identifier:      &resourceID,
			}, nil
		}
		mockAPI.WaitForResourceOpCompletionFunc = func(ctx context.Context, pi *types.ProgressEvent) (*types.ProgressEvent, error) {
			return pi, nil
		}
		getResourceInvocation := 0
		mockAPI.GetResourceFunc = func(ctx context.Context, typeName, identifier string) (map[string]interface{}, error) {
			switch getResourceInvocation {
			case 0:
				getResourceInvocation++
				return nil, &smithy.OperationError{
					ServiceID:     "CloudControl",
					OperationName: "GetResource",
					Err: &awshttp.ResponseError{
						RequestID: "9068dc0f-1124-4fc7-a120-39956230b810",
						ResponseError: &smithyhttp.ResponseError{
							Response: &smithyhttp.Response{
								Response: &http.Response{
									StatusCode: 400,
								},
							},
							Err: &types.ResourceNotFoundException{},
						},
					},
				}
			default:
				return map[string]interface{}{"output1": "outvalue1"}, nil
			}
		}

		id, outputs, err := client.Create(ctx, testURN, typeName, desiredState)

		assert.NoError(t, err)
		t.Logf("ID=%v", id)
		t.Logf("Outputs=%v", outputs)
	})
}

// Mock API implementation
type mockAPI struct {
	GetResourceFunc                 func(ctx context.Context, typeName, identifier string) (map[string]interface{}, error)
	CreateResourceFunc              func(ctx context.Context, cfType, desiredState string) (*types.ProgressEvent, error)
	WaitForResourceOpCompletionFunc func(ctx context.Context, pi *types.ProgressEvent) (*types.ProgressEvent, error)
}

func (m *mockAPI) GetResource(ctx context.Context, typeName, identifier string) (map[string]interface{}, error) {
	return m.GetResourceFunc(ctx, typeName, identifier)
}

func (m *mockAPI) CreateResource(ctx context.Context, cfType, desiredState string) (*types.ProgressEvent, error) {
	return m.CreateResourceFunc(ctx, cfType, desiredState)
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

func (m *mockAPI) WaitForResourceOpCompletion(ctx context.Context, pi *types.ProgressEvent) (*types.ProgressEvent, error) {
	return m.WaitForResourceOpCompletionFunc(ctx, pi)
}

func TestGetResourceRetryNotFoundRetrySettings(t *testing.T) {
	ci := &clientImpl{}
	var totalDelay time.Duration
	attempts, backoff := ci.getResourceRetryNotFoundRetrySettings()
	for attempt := 0; attempt < attempts; attempt++ {
		delay, err := backoff.BackoffDelay(attempt, nil)
		require.NoError(t, err)
		t.Logf("attempt %d delay is %v", attempt, delay)
		totalDelay += delay
	}
	// Typical test output (though it is random):
	//
	// client_test.go:378: attempt 0 delay is 745.506998ms
	// client_test.go:378: attempt 1 delay is 1.229581817s
	// client_test.go:378: attempt 2 delay is 931.993865ms
	// client_test.go:378: attempt 3 delay is 5s
	// client_test.go:378: attempt 4 delay is 5s
	//
	// The worst case wait is 25 seconds.
	require.LessOrEqual(t, totalDelay, 25*time.Second)
}
