// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Schema for EC2 VPN Gateway
func LookupVPNGateway(ctx *pulumi.Context, args *LookupVPNGatewayArgs, opts ...pulumi.InvokeOption) (*LookupVPNGatewayResult, error) {
	var rv LookupVPNGatewayResult
	err := ctx.Invoke("aws-native:ec2:getVPNGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVPNGatewayArgs struct {
	// VPN Gateway ID generated by service
	VPNGatewayId string `pulumi:"vPNGatewayId"`
}

type LookupVPNGatewayResult struct {
	// Any tags assigned to the virtual private gateway.
	Tags []VPNGatewayTag `pulumi:"tags"`
	// VPN Gateway ID generated by service
	VPNGatewayId *string `pulumi:"vPNGatewayId"`
}

func LookupVPNGatewayOutput(ctx *pulumi.Context, args LookupVPNGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupVPNGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVPNGatewayResult, error) {
			args := v.(LookupVPNGatewayArgs)
			r, err := LookupVPNGateway(ctx, &args, opts...)
			var s LookupVPNGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVPNGatewayResultOutput)
}

type LookupVPNGatewayOutputArgs struct {
	// VPN Gateway ID generated by service
	VPNGatewayId pulumi.StringInput `pulumi:"vPNGatewayId"`
}

func (LookupVPNGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVPNGatewayArgs)(nil)).Elem()
}

type LookupVPNGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupVPNGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVPNGatewayResult)(nil)).Elem()
}

func (o LookupVPNGatewayResultOutput) ToLookupVPNGatewayResultOutput() LookupVPNGatewayResultOutput {
	return o
}

func (o LookupVPNGatewayResultOutput) ToLookupVPNGatewayResultOutputWithContext(ctx context.Context) LookupVPNGatewayResultOutput {
	return o
}

// Any tags assigned to the virtual private gateway.
func (o LookupVPNGatewayResultOutput) Tags() VPNGatewayTagArrayOutput {
	return o.ApplyT(func(v LookupVPNGatewayResult) []VPNGatewayTag { return v.Tags }).(VPNGatewayTagArrayOutput)
}

// VPN Gateway ID generated by service
func (o LookupVPNGatewayResultOutput) VPNGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVPNGatewayResult) *string { return v.VPNGatewayId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVPNGatewayResultOutput{})
}