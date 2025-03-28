// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package msk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::MSK::ServerlessCluster
func LookupServerlessCluster(ctx *pulumi.Context, args *LookupServerlessClusterArgs, opts ...pulumi.InvokeOption) (*LookupServerlessClusterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServerlessClusterResult
	err := ctx.Invoke("aws-native:msk:getServerlessCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerlessClusterArgs struct {
	// The Amazon Resource Name (ARN) of the MSK cluster.
	Arn string `pulumi:"arn"`
}

type LookupServerlessClusterResult struct {
	// The Amazon Resource Name (ARN) of the MSK cluster.
	Arn *string `pulumi:"arn"`
}

func LookupServerlessClusterOutput(ctx *pulumi.Context, args LookupServerlessClusterOutputArgs, opts ...pulumi.InvokeOption) LookupServerlessClusterResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupServerlessClusterResultOutput, error) {
			args := v.(LookupServerlessClusterArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:msk:getServerlessCluster", args, LookupServerlessClusterResultOutput{}, options).(LookupServerlessClusterResultOutput), nil
		}).(LookupServerlessClusterResultOutput)
}

type LookupServerlessClusterOutputArgs struct {
	// The Amazon Resource Name (ARN) of the MSK cluster.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupServerlessClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerlessClusterArgs)(nil)).Elem()
}

type LookupServerlessClusterResultOutput struct{ *pulumi.OutputState }

func (LookupServerlessClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerlessClusterResult)(nil)).Elem()
}

func (o LookupServerlessClusterResultOutput) ToLookupServerlessClusterResultOutput() LookupServerlessClusterResultOutput {
	return o
}

func (o LookupServerlessClusterResultOutput) ToLookupServerlessClusterResultOutputWithContext(ctx context.Context) LookupServerlessClusterResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the MSK cluster.
func (o LookupServerlessClusterResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerlessClusterResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerlessClusterResultOutput{})
}
