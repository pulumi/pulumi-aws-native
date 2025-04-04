// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for Carrier Gateway which describes the Carrier Gateway resource
func LookupCarrierGateway(ctx *pulumi.Context, args *LookupCarrierGatewayArgs, opts ...pulumi.InvokeOption) (*LookupCarrierGatewayResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCarrierGatewayResult
	err := ctx.Invoke("aws-native:ec2:getCarrierGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCarrierGatewayArgs struct {
	// The ID of the carrier gateway.
	CarrierGatewayId string `pulumi:"carrierGatewayId"`
}

type LookupCarrierGatewayResult struct {
	// The ID of the carrier gateway.
	CarrierGatewayId *string `pulumi:"carrierGatewayId"`
	// The ID of the owner.
	OwnerId *string `pulumi:"ownerId"`
	// The state of the carrier gateway.
	State *string `pulumi:"state"`
	// The tags for the carrier gateway.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupCarrierGatewayOutput(ctx *pulumi.Context, args LookupCarrierGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupCarrierGatewayResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCarrierGatewayResultOutput, error) {
			args := v.(LookupCarrierGatewayArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:ec2:getCarrierGateway", args, LookupCarrierGatewayResultOutput{}, options).(LookupCarrierGatewayResultOutput), nil
		}).(LookupCarrierGatewayResultOutput)
}

type LookupCarrierGatewayOutputArgs struct {
	// The ID of the carrier gateway.
	CarrierGatewayId pulumi.StringInput `pulumi:"carrierGatewayId"`
}

func (LookupCarrierGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCarrierGatewayArgs)(nil)).Elem()
}

type LookupCarrierGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupCarrierGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCarrierGatewayResult)(nil)).Elem()
}

func (o LookupCarrierGatewayResultOutput) ToLookupCarrierGatewayResultOutput() LookupCarrierGatewayResultOutput {
	return o
}

func (o LookupCarrierGatewayResultOutput) ToLookupCarrierGatewayResultOutputWithContext(ctx context.Context) LookupCarrierGatewayResultOutput {
	return o
}

// The ID of the carrier gateway.
func (o LookupCarrierGatewayResultOutput) CarrierGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCarrierGatewayResult) *string { return v.CarrierGatewayId }).(pulumi.StringPtrOutput)
}

// The ID of the owner.
func (o LookupCarrierGatewayResultOutput) OwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCarrierGatewayResult) *string { return v.OwnerId }).(pulumi.StringPtrOutput)
}

// The state of the carrier gateway.
func (o LookupCarrierGatewayResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCarrierGatewayResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// The tags for the carrier gateway.
func (o LookupCarrierGatewayResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupCarrierGatewayResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCarrierGatewayResultOutput{})
}
