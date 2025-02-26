// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connect

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Connect::Queue
func LookupQueue(ctx *pulumi.Context, args *LookupQueueArgs, opts ...pulumi.InvokeOption) (*LookupQueueResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupQueueResult
	err := ctx.Invoke("aws-native:connect:getQueue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupQueueArgs struct {
	// The Amazon Resource Name (ARN) for the queue.
	QueueArn string `pulumi:"queueArn"`
}

type LookupQueueResult struct {
	// The description of the queue.
	Description *string `pulumi:"description"`
	// The identifier for the hours of operation.
	HoursOfOperationArn *string `pulumi:"hoursOfOperationArn"`
	// The identifier of the Amazon Connect instance.
	InstanceArn *string `pulumi:"instanceArn"`
	// The maximum number of contacts that can be in the queue before it is considered full.
	MaxContacts *int `pulumi:"maxContacts"`
	// The name of the queue.
	Name *string `pulumi:"name"`
	// The outbound caller ID name, number, and outbound whisper flow.
	OutboundCallerConfig *QueueOutboundCallerConfig `pulumi:"outboundCallerConfig"`
	// The outbound email address ID.
	OutboundEmailConfig *QueueOutboundEmailConfig `pulumi:"outboundEmailConfig"`
	// The Amazon Resource Name (ARN) for the queue.
	QueueArn *string `pulumi:"queueArn"`
	// The quick connects available to agents who are working the queue.
	QuickConnectArns []string `pulumi:"quickConnectArns"`
	// The status of the queue.
	Status *QueueStatus `pulumi:"status"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
	// The type of queue.
	Type *QueueType `pulumi:"type"`
}

func LookupQueueOutput(ctx *pulumi.Context, args LookupQueueOutputArgs, opts ...pulumi.InvokeOption) LookupQueueResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupQueueResultOutput, error) {
			args := v.(LookupQueueArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:connect:getQueue", args, LookupQueueResultOutput{}, options).(LookupQueueResultOutput), nil
		}).(LookupQueueResultOutput)
}

type LookupQueueOutputArgs struct {
	// The Amazon Resource Name (ARN) for the queue.
	QueueArn pulumi.StringInput `pulumi:"queueArn"`
}

func (LookupQueueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueueArgs)(nil)).Elem()
}

type LookupQueueResultOutput struct{ *pulumi.OutputState }

func (LookupQueueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueueResult)(nil)).Elem()
}

func (o LookupQueueResultOutput) ToLookupQueueResultOutput() LookupQueueResultOutput {
	return o
}

func (o LookupQueueResultOutput) ToLookupQueueResultOutputWithContext(ctx context.Context) LookupQueueResultOutput {
	return o
}

// The description of the queue.
func (o LookupQueueResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The identifier for the hours of operation.
func (o LookupQueueResultOutput) HoursOfOperationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *string { return v.HoursOfOperationArn }).(pulumi.StringPtrOutput)
}

// The identifier of the Amazon Connect instance.
func (o LookupQueueResultOutput) InstanceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *string { return v.InstanceArn }).(pulumi.StringPtrOutput)
}

// The maximum number of contacts that can be in the queue before it is considered full.
func (o LookupQueueResultOutput) MaxContacts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *int { return v.MaxContacts }).(pulumi.IntPtrOutput)
}

// The name of the queue.
func (o LookupQueueResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The outbound caller ID name, number, and outbound whisper flow.
func (o LookupQueueResultOutput) OutboundCallerConfig() QueueOutboundCallerConfigPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *QueueOutboundCallerConfig { return v.OutboundCallerConfig }).(QueueOutboundCallerConfigPtrOutput)
}

// The outbound email address ID.
func (o LookupQueueResultOutput) OutboundEmailConfig() QueueOutboundEmailConfigPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *QueueOutboundEmailConfig { return v.OutboundEmailConfig }).(QueueOutboundEmailConfigPtrOutput)
}

// The Amazon Resource Name (ARN) for the queue.
func (o LookupQueueResultOutput) QueueArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *string { return v.QueueArn }).(pulumi.StringPtrOutput)
}

// The quick connects available to agents who are working the queue.
func (o LookupQueueResultOutput) QuickConnectArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupQueueResult) []string { return v.QuickConnectArns }).(pulumi.StringArrayOutput)
}

// The status of the queue.
func (o LookupQueueResultOutput) Status() QueueStatusPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *QueueStatus { return v.Status }).(QueueStatusPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupQueueResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupQueueResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// The type of queue.
func (o LookupQueueResultOutput) Type() QueueTypePtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *QueueType { return v.Type }).(QueueTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQueueResultOutput{})
}
