package client

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
)

// HookError represents an error with CloudFormation hook information
type HookError struct {
	baseMessage string
	errorCode   string
	hookErrors  []string
}

func NewHookError(operation string, errorCode string, statusMessage string) *HookError {
	return &HookError{
		baseMessage: fmt.Sprintf("operation %s failed with %q: %s", operation, errorCode, statusMessage),
		errorCode:   errorCode,
		hookErrors:  []string{},
	}
}

func (e *HookError) WithHookEvent(hookEvent types.HookProgressEvent) *HookError {
	hookMessage := fmt.Sprintf("HookName: %s, HookArn: %s, HookVersion: %s, Time: %s, HookMessage: %s",
		aws.ToString(hookEvent.HookTypeName),
		aws.ToString(hookEvent.HookTypeArn),
		aws.ToString(hookEvent.HookTypeVersionId),
		hookEvent.HookEventTime.Format(time.RFC3339),
		aws.ToString(hookEvent.HookStatusMessage))

	e.hookErrors = append(e.hookErrors, hookMessage)
	return e
}

func (e *HookError) Error() string {
	msg := e.baseMessage
	if len(e.hookErrors) > 0 {
		msg += ". Hook failures: " + strings.Join(e.hookErrors, "; ")
	}
	return msg
}
