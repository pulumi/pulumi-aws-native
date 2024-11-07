package client

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	lambdaTypes "github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func lambdaSetup(t *testing.T) (*gomock.Controller, *lambdaClientImpl, *MockLambdaApi) {
	ctrl := gomock.NewController(t)
	mockLambdaApi := NewMockLambdaApi(ctrl)
	client := &lambdaClientImpl{
		api:                       mockLambdaApi,
		functionActivationTimeout: 2 * time.Second,
	}
	return ctrl, client, mockLambdaApi
}

func TestInvokeAsync_SuccessfulInvocation(t *testing.T) {
	t.Parallel()
	ctrl, client, mockLambdaApi := lambdaSetup(t)
	defer ctrl.Finish()

	ctx := context.Background()
	functionName := "test-function"
	payload := []byte(`{"key": "value"}`)

	mockLambdaApi.EXPECT().Invoke(ctx, gomock.Any()).Return(&lambda.InvokeOutput{
		StatusCode: 202,
	}, nil)

	err := client.InvokeAsync(ctx, functionName, payload)
	assert.NoError(t, err)
}

func TestInvokeAsync_FunctionNotReadyInitiallyButBecomesReady(t *testing.T) {
	t.Parallel()
	ctrl, client, mockLambdaApi := lambdaSetup(t)
	defer ctrl.Finish()

	ctx := context.Background()
	functionName := "test-function"
	payload := []byte(`{"key": "value"}`)

	mockLambdaApi.EXPECT().Invoke(ctx, gomock.Any()).Return(nil, &lambdaTypes.ResourceNotReadyException{})

	mockLambdaApi.EXPECT().GetFunction(gomock.Any(), &lambda.GetFunctionInput{
		FunctionName: aws.String(functionName),
	}, gomock.Any()).Return(&lambda.GetFunctionOutput{
		Configuration: &lambdaTypes.FunctionConfiguration{
			State: lambdaTypes.StateInactive,
		},
	}, nil)
	mockLambdaApi.EXPECT().GetFunction(gomock.Any(), &lambda.GetFunctionInput{
		FunctionName: aws.String(functionName),
	}, gomock.Any()).Return(&lambda.GetFunctionOutput{
		Configuration: &lambdaTypes.FunctionConfiguration{
			State: lambdaTypes.StateActive,
		},
	}, nil)

	mockLambdaApi.EXPECT().Invoke(ctx, gomock.Any()).Return(&lambda.InvokeOutput{
		StatusCode: 202,
	}, nil)

	err := client.InvokeAsync(ctx, functionName, payload)
	assert.NoError(t, err)
}

func TestInvokeAsync_FunctionNotReadyAndFailsToBecomeReady(t *testing.T) {
	t.Parallel()
	ctrl, client, mockLambdaApi := lambdaSetup(t)
	defer ctrl.Finish()

	ctx := context.Background()
	functionName := "test-function"
	payload := []byte(`{"key": "value"}`)

	mockLambdaApi.EXPECT().Invoke(ctx, gomock.Any()).Return(nil, &lambdaTypes.ResourceNotReadyException{})

	mockLambdaApi.EXPECT().GetFunction(gomock.Any(), &lambda.GetFunctionInput{
		FunctionName: aws.String(functionName),
	}, gomock.Any()).Return(&lambda.GetFunctionOutput{
		Configuration: &lambdaTypes.FunctionConfiguration{
			State: lambdaTypes.StateInactive,
		},
	}, nil).AnyTimes()

	err := client.InvokeAsync(ctx, functionName, payload)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to wait for function to become active")
}

func TestInvokeAsync_InvocationFailsWithNon202StatusCode(t *testing.T) {
	t.Parallel()
	ctrl, client, mockLambdaApi := lambdaSetup(t)
	defer ctrl.Finish()

	ctx := context.Background()
	functionName := "test-function"
	payload := []byte(`{"key": "value"}`)

	mockLambdaApi.EXPECT().Invoke(ctx, gomock.Any()).Return(&lambda.InvokeOutput{
		StatusCode: 400,
	}, nil)

	err := client.InvokeAsync(ctx, functionName, payload)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "lambda invocation failed with status code 400")
}

func TestInvokeAsync_FunctionNotReadyInitiallyButBecomesReadyAndThenFails(t *testing.T) {
	t.Parallel()
	ctrl, client, mockLambdaApi := lambdaSetup(t)
	defer ctrl.Finish()

	ctx := context.Background()
	functionName := "test-function"
	payload := []byte(`{"key": "value"}`)

	mockLambdaApi.EXPECT().Invoke(ctx, gomock.Any()).Return(nil, &lambdaTypes.ResourceNotReadyException{})
	mockLambdaApi.EXPECT().Invoke(ctx, gomock.Any()).Return(&lambda.InvokeOutput{
		StatusCode: 400,
	}, nil)

	mockLambdaApi.EXPECT().GetFunction(gomock.Any(), &lambda.GetFunctionInput{
		FunctionName: aws.String(functionName),
	}, gomock.Any()).Return(&lambda.GetFunctionOutput{
		Configuration: &lambdaTypes.FunctionConfiguration{
			State: lambdaTypes.StateInactive,
		},
	}, nil)
	mockLambdaApi.EXPECT().GetFunction(gomock.Any(), &lambda.GetFunctionInput{
		FunctionName: aws.String(functionName),
	}, gomock.Any()).Return(&lambda.GetFunctionOutput{
		Configuration: &lambdaTypes.FunctionConfiguration{
			State: lambdaTypes.StateActive,
		},
	}, nil)

	err := client.InvokeAsync(ctx, functionName, payload)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "lambda invocation failed with status code 400")
}
