// Copyright 2016-2024, Pulumi Corporation.

package client

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/golang/glog"
	"github.com/pkg/errors"
)

// CloudControlAwaiter provides a mechanism to poll long-running Cloud Control operations until they resolve
// to completion or failure.
type CloudControlAwaiter interface {
	// WaitForResourceOpCompletion keeps polling Cloud Control API until a long-running operation defined by
	// the input progress event resolved to completion or failre. It then returns the last progress event
	// that signifies the final status of the resource operation.
	WaitForResourceOpCompletion(ctx context.Context, pi *types.ProgressEvent) (*types.ProgressEvent, error)
}

func NewCloudControlAwaiter(client CloudControlApiClient) CloudControlAwaiter {
	return &ccAwaiterImpl{
		client: client,
	}
}

type ccAwaiterImpl struct {
	client CloudControlApiClient
}

func (a *ccAwaiterImpl) WaitForResourceOpCompletion(ctx context.Context, pi *types.ProgressEvent) (*types.ProgressEvent, error) {
	retryBackoff := retry.NewExponentialJitterBackoff(30 * time.Second)
	i := 0
	for {
		status := pi.OperationStatus
		identifier := ""
		if pi.Identifier != nil {
			identifier = *pi.Identifier
		}
		glog.V(9).Infof("waiting for resource %q: attempt #%d status %q", identifier, i, status)

		var (
			finished   bool
			err        error
			hookEvents []types.HookProgressEvent
		)
		if status == "FAILED" && pi.RequestToken != nil {
			pi, hookEvents, err = a.client.GetResourceRequestStatusWithHooks(ctx, *pi.RequestToken)
			if err != nil {
				return nil, errors.Wrap(err, "getting resource request status")
			}
		}

		finished, err = hasFinishedWithHooks(pi, hookEvents)
		if finished || err != nil {
			return pi, err
		}

		var pause time.Duration
		if pi.RetryAfter != nil && pi.RetryAfter.After(time.Now()) {
			pause = time.Until(*pi.RetryAfter)
		} else {
			pause, err = retryBackoff.BackoffDelay(i, err)
			if err != nil {
				return nil, err
			}
		}
		glog.V(9).Infof("resource operation is in progress, pausing for %v", pause)
		time.Sleep(pause)

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default: // Continue to wait
		}

		if pi.RequestToken == nil {
			return nil, errors.New("missing request token")
		}

		pi, hookEvents, err = a.client.GetResourceRequestStatusWithHooks(ctx, *pi.RequestToken)
		if err != nil {
			return nil, errors.Wrap(err, "getting resource request status")
		}

		finished, err = hasFinishedWithHooks(pi, hookEvents)
		if finished || err != nil {
			return pi, err
		}
		i += 1
	}
}

func hasFinished(pi *types.ProgressEvent) (bool, error) {
	return hasFinishedWithHooks(pi, nil)
}

func hasFinishedWithHooks(pi *types.ProgressEvent, hookEvents []types.HookProgressEvent) (bool, error) {
	status := pi.OperationStatus
	switch status {
	case "SUCCESS":
		return true, nil
	case "FAILED":
		statusMessage := ""
		if pi.StatusMessage != nil {
			statusMessage = *pi.StatusMessage
		}

		hookErr := NewHookError(string(pi.Operation), string(pi.ErrorCode), statusMessage)

		// Add hook information if available
		if len(hookEvents) > 0 {
			for _, hookEvent := range hookEvents {
				hookStatus := aws.ToString(hookEvent.HookStatus)
				// Check for failed hook statuses
				if hookStatus == "HOOK_COMPLETE_FAILED" || hookStatus == "HOOK_FAILED" {
					hookErr.WithHookEvent(hookEvent)
				}
			}
		}

		return true, hookErr
	case "IN_PROGRESS", "PENDING":
		return false, nil
	default:
		return true, errors.Errorf("unknown status %q: %+v", status, pi)
	}
}
