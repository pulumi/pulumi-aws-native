// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kinesisvideo

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type Definition for AWS::KinesisVideo::SignalingChannel
func LookupSignalingChannel(ctx *pulumi.Context, args *LookupSignalingChannelArgs, opts ...pulumi.InvokeOption) (*LookupSignalingChannelResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSignalingChannelResult
	err := ctx.Invoke("aws-native:kinesisvideo:getSignalingChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSignalingChannelArgs struct {
	// The name of the Kinesis Video Signaling Channel.
	Name string `pulumi:"name"`
}

type LookupSignalingChannelResult struct {
	// The Amazon Resource Name (ARN) of the Kinesis Video Signaling Channel.
	Arn *string `pulumi:"arn"`
	// The period of time a signaling channel retains undelivered messages before they are discarded.
	MessageTtlSeconds *int `pulumi:"messageTtlSeconds"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
	// The type of the Kinesis Video Signaling Channel to create. Currently, SINGLE_MASTER is the only supported channel type.
	Type *SignalingChannelType `pulumi:"type"`
}

func LookupSignalingChannelOutput(ctx *pulumi.Context, args LookupSignalingChannelOutputArgs, opts ...pulumi.InvokeOption) LookupSignalingChannelResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSignalingChannelResultOutput, error) {
			args := v.(LookupSignalingChannelArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:kinesisvideo:getSignalingChannel", args, LookupSignalingChannelResultOutput{}, options).(LookupSignalingChannelResultOutput), nil
		}).(LookupSignalingChannelResultOutput)
}

type LookupSignalingChannelOutputArgs struct {
	// The name of the Kinesis Video Signaling Channel.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupSignalingChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalingChannelArgs)(nil)).Elem()
}

type LookupSignalingChannelResultOutput struct{ *pulumi.OutputState }

func (LookupSignalingChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalingChannelResult)(nil)).Elem()
}

func (o LookupSignalingChannelResultOutput) ToLookupSignalingChannelResultOutput() LookupSignalingChannelResultOutput {
	return o
}

func (o LookupSignalingChannelResultOutput) ToLookupSignalingChannelResultOutputWithContext(ctx context.Context) LookupSignalingChannelResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the Kinesis Video Signaling Channel.
func (o LookupSignalingChannelResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalingChannelResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The period of time a signaling channel retains undelivered messages before they are discarded.
func (o LookupSignalingChannelResultOutput) MessageTtlSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSignalingChannelResult) *int { return v.MessageTtlSeconds }).(pulumi.IntPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupSignalingChannelResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupSignalingChannelResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// The type of the Kinesis Video Signaling Channel to create. Currently, SINGLE_MASTER is the only supported channel type.
func (o LookupSignalingChannelResultOutput) Type() SignalingChannelTypePtrOutput {
	return o.ApplyT(func(v LookupSignalingChannelResult) *SignalingChannelType { return v.Type }).(SignalingChannelTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSignalingChannelResultOutput{})
}
