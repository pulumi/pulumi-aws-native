// Copyright 2016-2024, Pulumi Corporation.

//nolint:goconst // Repeated literals keep table-driven test fixtures readable.
package client

import (
	"context"
	"fmt"
	"testing"
	"testing/synctest"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/mattbaird/jsonpatch"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWaitForResourceOpCompletion_PollingDelays(t *testing.T) {
	for _, tt := range []struct {
		name             string
		hints            []time.Duration // Zero means absent; negative means expired.
		fallbackAttempts []int
		delays           []time.Duration
	}{
		{
			name:             "consecutive fallback polls increase backoff",
			hints:            []time.Duration{0, 0, 0},
			fallbackAttempts: []int{0, 1, 2},
			delays:           []time.Duration{time.Second, 2 * time.Second, 4 * time.Second},
		},
		{
			name:             "AWS hints do not advance fallback backoff",
			hints:            []time.Duration{10 * time.Second, 20 * time.Second, 0, 0},
			fallbackAttempts: []int{0, 1},
			delays:           []time.Duration{10 * time.Second, 20 * time.Second, time.Second, 2 * time.Second},
		},
		{
			name:             "future hint resets an existing fallback sequence",
			hints:            []time.Duration{0, 0, 10 * time.Second, 0, 0},
			fallbackAttempts: []int{0, 1, 0, 1},
			delays:           []time.Duration{time.Second, 2 * time.Second, 10 * time.Second, time.Second, 2 * time.Second},
		},
		{
			name:             "expired hints do not reset fallback backoff",
			hints:            []time.Duration{-time.Second, 0, -time.Second},
			fallbackAttempts: []int{0, 1, 2},
			delays:           []time.Duration{time.Second, 2 * time.Second, 4 * time.Second},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				progress := func(hint time.Duration) *types.ProgressEvent {
					pi := &types.ProgressEvent{
						OperationStatus: types.OperationStatusInProgress,
						RequestToken:    aws.String("request-token"),
					}
					if hint != 0 {
						retryAfter := time.Now().Add(hint)
						pi.RetryAfter = &retryAfter
					}
					return pi
				}
				finalProgress := &types.ProgressEvent{OperationStatus: types.OperationStatusSuccess}
				calls := 0
				lastResponse := time.Now()
				api := &awaiterTestAPI{getResourceRequestStatusWithHooksFunc: func(
					_ context.Context, token string,
				) (*types.ProgressEvent, []types.HookProgressEvent, error) {
					require.Equal(t, "request-token", token)
					require.Less(t, calls, len(tt.delays), "unexpected status poll after SUCCESS")
					assert.Equal(t, tt.delays[calls], time.Since(lastResponse), "delay before poll %d", calls)
					lastResponse = time.Now()
					calls++
					if calls == len(tt.hints) {
						return finalProgress, nil, nil
					}
					return progress(tt.hints[calls]), nil, nil
				}}
				var attempts []int
				awaiter := NewCloudControlAwaiter(api).(*ccAwaiterImpl)
				awaiter.backoff = retry.BackoffDelayerFunc(func(attempt int, _ error) (time.Duration, error) {
					attempts = append(attempts, attempt)
					return time.Second << attempt, nil
				})

				result, err := awaiter.WaitForResourceOpCompletion(context.Background(), progress(tt.hints[0]))

				require.NoError(t, err)
				assert.Same(t, finalProgress, result)
				assert.Equal(t, len(tt.hints), calls, "must keep polling until terminal SUCCESS")
				assert.Equal(t, tt.fallbackAttempts, attempts)
			})
		})
	}
}

func Test_hasFinished(t *testing.T) {
	tests := []struct {
		name           string
		status         types.OperationStatus
		statusMessage  string
		operation      types.Operation
		errorCode      types.HandlerErrorCode
		expectError    string
		expectFinished bool
	}{
		{
			status:         "SUCCESS",
			expectFinished: true,
		},
		{
			status:         "PENDING",
			expectFinished: false,
		},
		{
			status:         "IN_PROGRESS",
			expectFinished: false,
		},
		{
			status:         "FAILED",
			statusMessage:  "failure message",
			operation:      "GET",
			errorCode:      "400",
			expectFinished: true,
			expectError:    fmt.Sprintf("operation %s failed with %q: %s", "GET", "400", "failure message"),
		},
		{
			status:         "BOBOB",
			expectFinished: true,
			expectError:    fmt.Sprintf("unknown status %q: ", "BOBOB"),
		},
		{
			status:         "CANCEL_COMPLETE", // This is a real unhandled status that might be returned
			expectFinished: true,
			expectError:    fmt.Sprintf("unknown status %q: ", "CANCEL_COMPLETE"),
		},
		{
			status:         "CANCEL_IN_PROGRESS", // This is a real unhandled status that might be returned
			expectFinished: true,
			expectError:    fmt.Sprintf("unknown status %q: ", "CANCEL_IN_PROGRESS"),
		},
	}

	for _, tt := range tests {
		t.Run(string(tt.status)+" "+tt.name, func(t *testing.T) {
			pi := types.ProgressEvent{
				OperationStatus: tt.status,
				StatusMessage:   &tt.statusMessage,
				Operation:       tt.operation,
				ErrorCode:       tt.errorCode,
			}
			finished, err := hasFinished(&pi)
			if tt.expectError != "" {
				assert.ErrorContains(t, err, tt.expectError)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectFinished, finished)
		})
	}
}

func TestWaitForResourceOpCompletion_InitialFailedStatusIncludesHookFailureDetails(t *testing.T) {
	requestToken := "request-token"
	statusMessage := "resource operation failed"
	hookTime := time.Date(2026, time.February, 17, 12, 0, 0, 0, time.UTC)
	finalProgress := &types.ProgressEvent{
		OperationStatus: "FAILED",
		Operation:       "CREATE",
		ErrorCode:       "GeneralServiceException",
		StatusMessage:   &statusMessage,
		RequestToken:    &requestToken,
	}

	requestStatusCalls := 0
	mockAPI := &awaiterTestAPI{
		getResourceRequestStatusWithHooksFunc: func(
			_ context.Context,
			token string,
		) (*types.ProgressEvent, []types.HookProgressEvent, error) {
			requestStatusCalls++
			assert.Equal(t, requestToken, token)

			return finalProgress, []types.HookProgressEvent{
				{
					HookStatus:        aws.String("HOOK_COMPLETE_SUCCEEDED"),
					HookTypeName:      aws.String("AWS::Hooks::SucceededHook"),
					HookStatusMessage: aws.String("this should not appear in the error"),
				},
				{
					HookStatus:        aws.String("HOOK_COMPLETE_FAILED"),
					HookTypeName:      aws.String("AWS::Hooks::FailedHook"),
					HookTypeArn:       aws.String("arn:aws:cloudformation:us-west-2:123456789012:type/hook/FailedHook"),
					HookTypeVersionId: aws.String("00000001"),
					HookEventTime:     &hookTime,
					HookStatusMessage: aws.String("hook validation failed"),
				},
			}, nil
		},
	}
	awaiter := NewCloudControlAwaiter(mockAPI)

	pi, err := awaiter.WaitForResourceOpCompletion(context.Background(), &types.ProgressEvent{
		OperationStatus: "FAILED",
		RequestToken:    &requestToken,
	})

	require.Error(t, err)
	require.Equal(t, finalProgress, pi)
	assert.Equal(t, 1, requestStatusCalls)
	assert.ErrorContains(t, err, `operation CREATE failed with "GeneralServiceException": resource operation failed`)
	assert.ErrorContains(t, err, "Hook failures:")
	assert.ErrorContains(t, err, "HookName: AWS::Hooks::FailedHook")
	assert.ErrorContains(t, err, "hook validation failed")
	assert.NotContains(t, err.Error(), "AWS::Hooks::SucceededHook")
}

func TestHasFinishedWithHooks_FailedWithoutHookEvents(t *testing.T) {
	statusMessage := "failure message"
	pi := &types.ProgressEvent{
		OperationStatus: "FAILED",
		Operation:       "UPDATE",
		ErrorCode:       "NotUpdatable",
		StatusMessage:   &statusMessage,
	}

	finished, err := hasFinishedWithHooks(pi, nil)

	require.Error(t, err)
	assert.True(t, finished)
	assert.ErrorContains(t, err, `operation UPDATE failed with "NotUpdatable": failure message`)
	assert.NotContains(t, err.Error(), "Hook failures:")
}

type awaiterTestAPI struct {
	getResourceRequestStatusWithHooksFunc func(
		ctx context.Context,
		requestToken string,
	) (*types.ProgressEvent, []types.HookProgressEvent, error)
}

func (*awaiterTestAPI) CreateResource(_ context.Context, _, _ string) (*types.ProgressEvent, error) {
	panic("unexpected CreateResource call")
}

func (*awaiterTestAPI) UpdateResource(
	_ context.Context,
	_, _ string,
	_ []jsonpatch.JsonPatchOperation,
) (*types.ProgressEvent, error) {
	panic("unexpected UpdateResource call")
}

func (*awaiterTestAPI) DeleteResource(_ context.Context, _, _ string) (*types.ProgressEvent, error) {
	panic("unexpected DeleteResource call")
}

func (*awaiterTestAPI) GetResource(_ context.Context, _, _ string) (map[string]interface{}, error) {
	panic("unexpected GetResource call")
}

func (*awaiterTestAPI) ListResources(
	_ context.Context,
	_ string,
	_ *string,
	_ *string,
	_ *int32,
) ([]types.ResourceDescription, *string, error) {
	panic("unexpected ListResources call")
}

func (*awaiterTestAPI) GetResourceRequestStatus(
	_ context.Context,
	_ string,
) (*types.ProgressEvent, error) {
	panic("unexpected GetResourceRequestStatus call")
}

func (m *awaiterTestAPI) GetResourceRequestStatusWithHooks(
	ctx context.Context,
	requestToken string,
) (*types.ProgressEvent, []types.HookProgressEvent, error) {
	if m.getResourceRequestStatusWithHooksFunc == nil {
		panic("unexpected GetResourceRequestStatusWithHooks call")
	}
	return m.getResourceRequestStatusWithHooksFunc(ctx, requestToken)
}
