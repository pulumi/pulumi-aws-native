// Copyright 2016-2024, Pulumi Corporation.

package client

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/stretchr/testify/assert"
)

func TestHookError_WithHookEvent_NilHookEventTime(t *testing.T) {
	err := NewHookError("CREATE", "GeneralServiceException", "request failed").
		WithHookEvent(types.HookProgressEvent{
			HookTypeName:      aws.String("My::Hook"),
			HookTypeArn:       aws.String("arn:aws:cloudformation:us-west-2:123456789012:type/hook/My-Hook"),
			HookTypeVersionId: aws.String("00000001"),
			HookStatusMessage: aws.String("hook failed without event timestamp"),
		})

	assert.ErrorContains(t, err, `operation CREATE failed with "GeneralServiceException": request failed`)
	assert.ErrorContains(t, err, "Hook failures:")
	assert.ErrorContains(t, err, "HookName: My::Hook")
	assert.ErrorContains(t, err, "Time: N/A")
}
