// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssmcontacts

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SSMContacts::ContactChannel
func LookupContactChannel(ctx *pulumi.Context, args *LookupContactChannelArgs, opts ...pulumi.InvokeOption) (*LookupContactChannelResult, error) {
	var rv LookupContactChannelResult
	err := ctx.Invoke("aws-native:ssmcontacts:getContactChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContactChannelArgs struct {
	// The Amazon Resource Name (ARN) of the engagement to a contact channel.
	Arn string `pulumi:"arn"`
}

type LookupContactChannelResult struct {
	// The Amazon Resource Name (ARN) of the engagement to a contact channel.
	Arn *string `pulumi:"arn"`
	// The details that SSM Incident Manager uses when trying to engage the contact channel.
	ChannelAddress *string `pulumi:"channelAddress"`
	// The device name. String of 6 to 50 alphabetical, numeric, dash, and underscore characters.
	ChannelName *string `pulumi:"channelName"`
	// If you want to activate the channel at a later time, you can choose to defer activation. SSM Incident Manager can't engage your contact channel until it has been activated.
	DeferActivation *bool `pulumi:"deferActivation"`
}

func LookupContactChannelOutput(ctx *pulumi.Context, args LookupContactChannelOutputArgs, opts ...pulumi.InvokeOption) LookupContactChannelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContactChannelResult, error) {
			args := v.(LookupContactChannelArgs)
			r, err := LookupContactChannel(ctx, &args, opts...)
			var s LookupContactChannelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContactChannelResultOutput)
}

type LookupContactChannelOutputArgs struct {
	// The Amazon Resource Name (ARN) of the engagement to a contact channel.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupContactChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContactChannelArgs)(nil)).Elem()
}

type LookupContactChannelResultOutput struct{ *pulumi.OutputState }

func (LookupContactChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContactChannelResult)(nil)).Elem()
}

func (o LookupContactChannelResultOutput) ToLookupContactChannelResultOutput() LookupContactChannelResultOutput {
	return o
}

func (o LookupContactChannelResultOutput) ToLookupContactChannelResultOutputWithContext(ctx context.Context) LookupContactChannelResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the engagement to a contact channel.
func (o LookupContactChannelResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContactChannelResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The details that SSM Incident Manager uses when trying to engage the contact channel.
func (o LookupContactChannelResultOutput) ChannelAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContactChannelResult) *string { return v.ChannelAddress }).(pulumi.StringPtrOutput)
}

// The device name. String of 6 to 50 alphabetical, numeric, dash, and underscore characters.
func (o LookupContactChannelResultOutput) ChannelName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContactChannelResult) *string { return v.ChannelName }).(pulumi.StringPtrOutput)
}

// If you want to activate the channel at a later time, you can choose to defer activation. SSM Incident Manager can't engage your contact channel until it has been activated.
func (o LookupContactChannelResultOutput) DeferActivation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupContactChannelResult) *bool { return v.DeferActivation }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContactChannelResultOutput{})
}