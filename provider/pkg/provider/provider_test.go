package provider

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/stretchr/testify/assert"
)

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
