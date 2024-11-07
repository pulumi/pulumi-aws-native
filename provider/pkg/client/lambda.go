// Copyright 2016-2024, Pulumi Corporation.

package client

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	lambdaTypes "github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/pkg/errors"
)

const DefaultFunctionActivationTimeout = 5 * time.Minute

//go:generate mockgen -package client -source lambda.go -destination mock_lambda.go LambdaClient,LambdaApi
type LambdaClient interface {
	// InvokeAsync invokes the given Lambda function with the 'Event' invocation type. This means that the function
	// will be invoked asynchronously and no response will be returned. If the function is not ready to be invoked yet,
	// this method will wait for the function to become active before invoking it.
	InvokeAsync(ctx context.Context, functionName string, payload []byte) error
}

type LambdaApi interface {
	Invoke(ctx context.Context, params *lambda.InvokeInput, optFns ...func(*lambda.Options)) (*lambda.InvokeOutput, error)
	GetFunction(context.Context, *lambda.GetFunctionInput, ...func(*lambda.Options)) (*lambda.GetFunctionOutput, error)
}

type lambdaClientImpl struct {
	api LambdaApi
	functionActivationTimeout time.Duration
}

func NewLambdaClient(api LambdaApi) LambdaClient {
	return &lambdaClientImpl{
		api: api,
		functionActivationTimeout: DefaultFunctionActivationTimeout,
	}
}

func (c *lambdaClientImpl) InvokeAsync(ctx context.Context, functionName string, payload []byte) error {
	input := lambda.InvokeInput{
		FunctionName: aws.String(functionName),
		Payload:      payload,
		// async invocation
		InvocationType: lambdaTypes.InvocationTypeEvent,
	}

	// fire off an initial invoke. If the function is not ready, we need to wait for it to become ready.
	// this initial invoke will trigger the AWS Lambda service to start the function transition process.
	invokeResp, err := c.api.Invoke(ctx, &input)

	if err != nil {
		// Lambda functions can be in a state where they are not ready to be invoked yet.
		// If we get this error, we need to wait for the function to become active
		var notReadyErr *lambdaTypes.ResourceNotReadyException
		if errors.As(err, &notReadyErr) {
			err := c.waitForFunctionActive(ctx, functionName)
			if err != nil {
				return fmt.Errorf("failed to wait for function to become active: %w", err)
			}
			invokeResp, err = c.api.Invoke(ctx, &input)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	if invokeResp.StatusCode != 202 {
		return fmt.Errorf("lambda invocation failed with status code %d", invokeResp.StatusCode)
	}

	return nil
}

// waitForFunctionActive waits for the function to be in the active state. If the function is not ready
// after 5 minutes, it will return an error.
func (c *lambdaClientImpl) waitForFunctionActive(ctx context.Context, functionName string) error {
	err := lambda.NewFunctionActiveV2Waiter(c.api, func(o *lambda.FunctionActiveV2WaiterOptions) {
		// We already aggressively retry SDK errors with plenty retry attempts and
		// throttling exceptions will be taken care of by the SDK
		o.MinDelay = time.Second
		o.MaxDelay = 5 * time.Second
	}).Wait(ctx, &lambda.GetFunctionInput{
		FunctionName: aws.String(functionName),
	}, c.functionActivationTimeout)

	return err
}
