// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::Lambda::EventInvokeConfig resource configures options for asynchronous invocation on a version or an alias.
func LookupEventInvokeConfig(ctx *pulumi.Context, args *LookupEventInvokeConfigArgs, opts ...pulumi.InvokeOption) (*LookupEventInvokeConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEventInvokeConfigResult
	err := ctx.Invoke("aws-native:lambda:getEventInvokeConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventInvokeConfigArgs struct {
	// The name of the Lambda function.
	FunctionName string `pulumi:"functionName"`
	// The identifier of a version or alias.
	Qualifier string `pulumi:"qualifier"`
}

type LookupEventInvokeConfigResult struct {
	// A destination for events after they have been sent to a function for processing.
	//
	// **Destinations** - *Function* - The Amazon Resource Name (ARN) of a Lambda function.
	// - *Queue* - The ARN of a standard SQS queue.
	// - *Bucket* - The ARN of an Amazon S3 bucket.
	// - *Topic* - The ARN of a standard SNS topic.
	// - *Event Bus* - The ARN of an Amazon EventBridge event bus.
	//
	// > S3 buckets are supported only for on-failure destinations. To retain records of successful invocations, use another destination type.
	DestinationConfig *EventInvokeConfigDestinationConfig `pulumi:"destinationConfig"`
	// The maximum age of a request that Lambda sends to a function for processing.
	MaximumEventAgeInSeconds *int `pulumi:"maximumEventAgeInSeconds"`
	// The maximum number of times to retry when the function returns an error.
	MaximumRetryAttempts *int `pulumi:"maximumRetryAttempts"`
}

func LookupEventInvokeConfigOutput(ctx *pulumi.Context, args LookupEventInvokeConfigOutputArgs, opts ...pulumi.InvokeOption) LookupEventInvokeConfigResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupEventInvokeConfigResultOutput, error) {
			args := v.(LookupEventInvokeConfigArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:lambda:getEventInvokeConfig", args, LookupEventInvokeConfigResultOutput{}, options).(LookupEventInvokeConfigResultOutput), nil
		}).(LookupEventInvokeConfigResultOutput)
}

type LookupEventInvokeConfigOutputArgs struct {
	// The name of the Lambda function.
	FunctionName pulumi.StringInput `pulumi:"functionName"`
	// The identifier of a version or alias.
	Qualifier pulumi.StringInput `pulumi:"qualifier"`
}

func (LookupEventInvokeConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventInvokeConfigArgs)(nil)).Elem()
}

type LookupEventInvokeConfigResultOutput struct{ *pulumi.OutputState }

func (LookupEventInvokeConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventInvokeConfigResult)(nil)).Elem()
}

func (o LookupEventInvokeConfigResultOutput) ToLookupEventInvokeConfigResultOutput() LookupEventInvokeConfigResultOutput {
	return o
}

func (o LookupEventInvokeConfigResultOutput) ToLookupEventInvokeConfigResultOutputWithContext(ctx context.Context) LookupEventInvokeConfigResultOutput {
	return o
}

// A destination for events after they have been sent to a function for processing.
//
// **Destinations** - *Function* - The Amazon Resource Name (ARN) of a Lambda function.
// - *Queue* - The ARN of a standard SQS queue.
// - *Bucket* - The ARN of an Amazon S3 bucket.
// - *Topic* - The ARN of a standard SNS topic.
// - *Event Bus* - The ARN of an Amazon EventBridge event bus.
//
// > S3 buckets are supported only for on-failure destinations. To retain records of successful invocations, use another destination type.
func (o LookupEventInvokeConfigResultOutput) DestinationConfig() EventInvokeConfigDestinationConfigPtrOutput {
	return o.ApplyT(func(v LookupEventInvokeConfigResult) *EventInvokeConfigDestinationConfig { return v.DestinationConfig }).(EventInvokeConfigDestinationConfigPtrOutput)
}

// The maximum age of a request that Lambda sends to a function for processing.
func (o LookupEventInvokeConfigResultOutput) MaximumEventAgeInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEventInvokeConfigResult) *int { return v.MaximumEventAgeInSeconds }).(pulumi.IntPtrOutput)
}

// The maximum number of times to retry when the function returns an error.
func (o LookupEventInvokeConfigResultOutput) MaximumRetryAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEventInvokeConfigResult) *int { return v.MaximumRetryAttempts }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventInvokeConfigResultOutput{})
}
