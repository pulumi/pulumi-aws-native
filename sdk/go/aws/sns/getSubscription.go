// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SNS::Subscription
func LookupSubscription(ctx *pulumi.Context, args *LookupSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupSubscriptionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSubscriptionResult
	err := ctx.Invoke("aws-native:sns:getSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubscriptionArgs struct {
	// Arn of the subscription
	Arn string `pulumi:"arn"`
}

type LookupSubscriptionResult struct {
	// Arn of the subscription
	Arn *string `pulumi:"arn"`
	// The delivery policy JSON assigned to the subscription. Enables the subscriber to define the message delivery retry strategy in the case of an HTTP/S endpoint subscribed to the topic.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SNS::Subscription` for more information about the expected schema for this property.
	DeliveryPolicy interface{} `pulumi:"deliveryPolicy"`
	// The filter policy JSON assigned to the subscription. Enables the subscriber to filter out unwanted messages.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SNS::Subscription` for more information about the expected schema for this property.
	FilterPolicy interface{} `pulumi:"filterPolicy"`
	// This attribute lets you choose the filtering scope by using one of the following string value types: MessageAttributes (default) and MessageBody.
	FilterPolicyScope *string `pulumi:"filterPolicyScope"`
	// When set to true, enables raw message delivery. Raw messages don't contain any JSON formatting and can be sent to Amazon SQS and HTTP/S endpoints.
	RawMessageDelivery *bool `pulumi:"rawMessageDelivery"`
	// When specified, sends undeliverable messages to the specified Amazon SQS dead-letter queue. Messages that can't be delivered due to client errors are held in the dead-letter queue for further analysis or reprocessing.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SNS::Subscription` for more information about the expected schema for this property.
	RedrivePolicy interface{} `pulumi:"redrivePolicy"`
	// Specifies whether Amazon SNS resends the notification to the subscription when a message's attribute changes.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SNS::Subscription` for more information about the expected schema for this property.
	ReplayPolicy interface{} `pulumi:"replayPolicy"`
	// This property applies only to Amazon Data Firehose delivery stream subscriptions.
	SubscriptionRoleArn *string `pulumi:"subscriptionRoleArn"`
}

func LookupSubscriptionOutput(ctx *pulumi.Context, args LookupSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSubscriptionResultOutput, error) {
			args := v.(LookupSubscriptionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:sns:getSubscription", args, LookupSubscriptionResultOutput{}, options).(LookupSubscriptionResultOutput), nil
		}).(LookupSubscriptionResultOutput)
}

type LookupSubscriptionOutputArgs struct {
	// Arn of the subscription
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionArgs)(nil)).Elem()
}

type LookupSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionResult)(nil)).Elem()
}

func (o LookupSubscriptionResultOutput) ToLookupSubscriptionResultOutput() LookupSubscriptionResultOutput {
	return o
}

func (o LookupSubscriptionResultOutput) ToLookupSubscriptionResultOutputWithContext(ctx context.Context) LookupSubscriptionResultOutput {
	return o
}

// Arn of the subscription
func (o LookupSubscriptionResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The delivery policy JSON assigned to the subscription. Enables the subscriber to define the message delivery retry strategy in the case of an HTTP/S endpoint subscribed to the topic.
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SNS::Subscription` for more information about the expected schema for this property.
func (o LookupSubscriptionResultOutput) DeliveryPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) interface{} { return v.DeliveryPolicy }).(pulumi.AnyOutput)
}

// The filter policy JSON assigned to the subscription. Enables the subscriber to filter out unwanted messages.
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SNS::Subscription` for more information about the expected schema for this property.
func (o LookupSubscriptionResultOutput) FilterPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) interface{} { return v.FilterPolicy }).(pulumi.AnyOutput)
}

// This attribute lets you choose the filtering scope by using one of the following string value types: MessageAttributes (default) and MessageBody.
func (o LookupSubscriptionResultOutput) FilterPolicyScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.FilterPolicyScope }).(pulumi.StringPtrOutput)
}

// When set to true, enables raw message delivery. Raw messages don't contain any JSON formatting and can be sent to Amazon SQS and HTTP/S endpoints.
func (o LookupSubscriptionResultOutput) RawMessageDelivery() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *bool { return v.RawMessageDelivery }).(pulumi.BoolPtrOutput)
}

// When specified, sends undeliverable messages to the specified Amazon SQS dead-letter queue. Messages that can't be delivered due to client errors are held in the dead-letter queue for further analysis or reprocessing.
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SNS::Subscription` for more information about the expected schema for this property.
func (o LookupSubscriptionResultOutput) RedrivePolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) interface{} { return v.RedrivePolicy }).(pulumi.AnyOutput)
}

// Specifies whether Amazon SNS resends the notification to the subscription when a message's attribute changes.
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SNS::Subscription` for more information about the expected schema for this property.
func (o LookupSubscriptionResultOutput) ReplayPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) interface{} { return v.ReplayPolicy }).(pulumi.AnyOutput)
}

// This property applies only to Amazon Data Firehose delivery stream subscriptions.
func (o LookupSubscriptionResultOutput) SubscriptionRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.SubscriptionRoleArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubscriptionResultOutput{})
}
