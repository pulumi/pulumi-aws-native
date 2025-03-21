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

// Allocates an internet gateway for use with a VPC. After creating the Internet gateway, you then attach it to a VPC.
func LookupInternetGateway(ctx *pulumi.Context, args *LookupInternetGatewayArgs, opts ...pulumi.InvokeOption) (*LookupInternetGatewayResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInternetGatewayResult
	err := ctx.Invoke("aws-native:ec2:getInternetGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInternetGatewayArgs struct {
	// The ID of the internet gateway.
	InternetGatewayId string `pulumi:"internetGatewayId"`
}

type LookupInternetGatewayResult struct {
	// The ID of the internet gateway.
	InternetGatewayId *string `pulumi:"internetGatewayId"`
	// Any tags to assign to the internet gateway.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupInternetGatewayOutput(ctx *pulumi.Context, args LookupInternetGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupInternetGatewayResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupInternetGatewayResultOutput, error) {
			args := v.(LookupInternetGatewayArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:ec2:getInternetGateway", args, LookupInternetGatewayResultOutput{}, options).(LookupInternetGatewayResultOutput), nil
		}).(LookupInternetGatewayResultOutput)
}

type LookupInternetGatewayOutputArgs struct {
	// The ID of the internet gateway.
	InternetGatewayId pulumi.StringInput `pulumi:"internetGatewayId"`
}

func (LookupInternetGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInternetGatewayArgs)(nil)).Elem()
}

type LookupInternetGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupInternetGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInternetGatewayResult)(nil)).Elem()
}

func (o LookupInternetGatewayResultOutput) ToLookupInternetGatewayResultOutput() LookupInternetGatewayResultOutput {
	return o
}

func (o LookupInternetGatewayResultOutput) ToLookupInternetGatewayResultOutputWithContext(ctx context.Context) LookupInternetGatewayResultOutput {
	return o
}

// The ID of the internet gateway.
func (o LookupInternetGatewayResultOutput) InternetGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInternetGatewayResult) *string { return v.InternetGatewayId }).(pulumi.StringPtrOutput)
}

// Any tags to assign to the internet gateway.
func (o LookupInternetGatewayResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupInternetGatewayResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInternetGatewayResultOutput{})
}
