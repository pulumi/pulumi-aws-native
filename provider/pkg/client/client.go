// Copyright 2016-2024, Pulumi Corporation.

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/aws/smithy-go"
	"github.com/golang/glog"
	"github.com/mattbaird/jsonpatch"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

// Logger provides a way to emit diagnostic messages with resource context.
type Logger interface {
	LogStatus(ctx context.Context, sev diag.Severity, urn resource.URN, msg string) error
}

// CloudControlApiClient providers CRUD operations around Cloud Control API, with the mechanics of API calls abstracted away.
// For instance, it serializes and deserializes wire data and follows the protocol of long-running operations.
//
//go:generate mockgen -package client -source client.go -destination mock_client.go CloudControlApiClient
type CloudControlClient interface {
	// Create creates a resource of the specified type with the desired state.
	// It awaits the operation until completion and returns a map of output property values.
	Create(ctx context.Context, urn resource.URN, typeName string, desiredState map[string]any) (identifier *string, resourceState map[string]any, err error)

	// Read returns the current state of the specified resource. It deserializes
	// the response from the service into a map of untyped values.
	// If the resource does not exist, no error is returned but the flag exists is set to false.
	Read(ctx context.Context, typeName, identifier string) (resourceState map[string]interface{}, exists bool, err error)

	// Update updates a resource of the specified type with the specified changeset.
	// It awaits the operation until completion and returns a map of output property values.
	Update(ctx context.Context, urn resource.URN, typeName, identifier string, patches []jsonpatch.JsonPatchOperation) (map[string]interface{}, error)

	// Delete deletes a resource of the specified type with the given identifier.
	// It awaits the operation until completion.
	Delete(ctx context.Context, urn resource.URN, typeName, identifier string) error
}

type clientImpl struct {
	api     CloudControlApiClient
	awaiter CloudControlAwaiter
	logger  Logger
}

func NewCloudControlClient(cctl *cloudcontrol.Client, roleArn *string, logger Logger) CloudControlClient {
	api := NewCloudControlApiClient(cctl, roleArn)
	return &clientImpl{
		api:     api,
		awaiter: NewCloudControlAwaiter(api),
		logger:  logger,
	}
}

func (c *clientImpl) Create(ctx context.Context, urn resource.URN, typeName string, desiredState map[string]any) (identifier *string, resourceState map[string]any, err error) {
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
	if waitErr != nil {
		if pi == nil || pi.Identifier == nil {
			return nil, nil, fmt.Errorf("creating resource (await): %w", waitErr)
		}
		if pi.ErrorCode == types.HandlerErrorCodeAlreadyExists {
			// Already Exists is a special case because it's the scenario when we can't proceed to resource read
			// to validate its existence. We should return immediately as a hard error.
			return nil, nil, waitErr
		}
	}
	if pi.Identifier == nil {
		return nil, nil, errors.New("received nil identifier while awaiting completion")
	}

	// Pick the getResource method. If there was an error from awaiting completion, simply try once. If there were
	// no errors however, try multiple times to tolerate occasional GetResource NotFound errors that arise due to
	// eventual consistency.
	getResource := c.api.GetResource
	if waitErr == nil {
		getResource = c.getResourceRetryNotFound
	}

	// Read the state - even if there was a creation error but the progress event contains a resource ID.
	// Retrieve the resource state from AWS.
	// Note that we do so even if creation hasn't succeeded but the identifier is assigned.
	resourceState, err = getResource(ctx, typeName, *pi.Identifier)
	if err != nil {
		if waitErr != nil {
			// Both wait and read fail. Provisioning failed entirely, return the wait error as more informative.
			return nil, nil, fmt.Errorf("creating resource (await): %w", waitErr)
		}
		// Creation succeeded but reading failed - return the read error.
		return nil, nil, fmt.Errorf("reading resource state: %w", err)
	}

	return pi.Identifier, resourceState, waitErr
}

func (c *clientImpl) Read(ctx context.Context, typeName, identifier string) (resourceState map[string]interface{}, exists bool, err error) {
	resourceState, err = c.api.GetResource(ctx, typeName, identifier)
	if err != nil {
		var oe *smithy.OperationError
		if errors.As(err, &oe) {
			if re, ok := oe.Unwrap().(*http.ResponseError); ok {
				statusCode := re.HTTPStatusCode()
				errorMessage := re.Error()
				isHttpNotFound := statusCode == 404
				isResourceNotFound := statusCode == 400 && strings.Contains(errorMessage, "ResourceNotFoundException")
				if isHttpNotFound || isResourceNotFound {
					return nil, false, nil
				}
			}
		}
		return nil, false, err
	}

	return resourceState, true, nil
}

func (c *clientImpl) Update(ctx context.Context, urn resource.URN, typeName, identifier string, patches []jsonpatch.JsonPatchOperation) (map[string]interface{}, error) {
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

func (c *clientImpl) Delete(ctx context.Context, urn resource.URN, typeName, identifier string) error {
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

// Wraps c.api.GetResource with Not Found error retry logic to try to compensate for eventual consistency issues for
// newly created resources.
func (c *clientImpl) getResourceRetryNotFound(
	ctx context.Context,
	typeName, identifier string,
) (map[string]interface{}, error) {
	maxAttempts, retryBackoff := c.getResourceRetryNotFoundRetrySettings()
	var lastError error
	for attempt := 0; attempt < maxAttempts; attempt++ {
		result, err := c.api.GetResource(ctx, typeName, identifier)
		var notFound *types.ResourceNotFoundException
		switch {
		case err == nil:
			return result, nil
		case !errors.As(err, &notFound):
			return nil, err
		}
		lastError = err

		delay, err := retryBackoff.BackoffDelay(attempt, nil)
		contract.AssertNoErrorf(err, "BackoffDelay should not fail")

		glog.V(9).Infof("CloudControl GetResource(%q, %q) failed with ResourceNotFoundException:"+
			" attempt #%d, retrying in %v", typeName, identifier, attempt, delay)

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(delay):
			continue
		}
	}
	return nil, lastError
}

func (*clientImpl) getResourceRetryNotFoundRetrySettings() (int, *retry.ExponentialJitterBackoff) {
	retryBackoff := retry.NewExponentialJitterBackoff(5 * time.Second)
	maxAttempts := 5
	return maxAttempts, retryBackoff
}
