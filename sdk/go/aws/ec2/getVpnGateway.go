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

// Specifies a virtual private gateway. A virtual private gateway is the endpoint on the VPC side of your VPN connection. You can create a virtual private gateway before creating the VPC itself.
//
//	For more information, see [](https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the *User Guide*.
func LookupVpnGateway(ctx *pulumi.Context, args *LookupVpnGatewayArgs, opts ...pulumi.InvokeOption) (*LookupVpnGatewayResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpnGatewayResult
	err := ctx.Invoke("aws-native:ec2:getVpnGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpnGatewayArgs struct {
	// The ID of the VPN gateway.
	VpnGatewayId string `pulumi:"vpnGatewayId"`
}

type LookupVpnGatewayResult struct {
	// Any tags assigned to the virtual private gateway.
	Tags []aws.Tag `pulumi:"tags"`
	// The ID of the VPN gateway.
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
}

func LookupVpnGatewayOutput(ctx *pulumi.Context, args LookupVpnGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupVpnGatewayResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVpnGatewayResultOutput, error) {
			args := v.(LookupVpnGatewayArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:ec2:getVpnGateway", args, LookupVpnGatewayResultOutput{}, options).(LookupVpnGatewayResultOutput), nil
		}).(LookupVpnGatewayResultOutput)
}

type LookupVpnGatewayOutputArgs struct {
	// The ID of the VPN gateway.
	VpnGatewayId pulumi.StringInput `pulumi:"vpnGatewayId"`
}

func (LookupVpnGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnGatewayArgs)(nil)).Elem()
}

type LookupVpnGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupVpnGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnGatewayResult)(nil)).Elem()
}

func (o LookupVpnGatewayResultOutput) ToLookupVpnGatewayResultOutput() LookupVpnGatewayResultOutput {
	return o
}

func (o LookupVpnGatewayResultOutput) ToLookupVpnGatewayResultOutputWithContext(ctx context.Context) LookupVpnGatewayResultOutput {
	return o
}

// Any tags assigned to the virtual private gateway.
func (o LookupVpnGatewayResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// The ID of the VPN gateway.
func (o LookupVpnGatewayResultOutput) VpnGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) *string { return v.VpnGatewayId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpnGatewayResultOutput{})
}
