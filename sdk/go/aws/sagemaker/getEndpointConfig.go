// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::EndpointConfig
func LookupEndpointConfig(ctx *pulumi.Context, args *LookupEndpointConfigArgs, opts ...pulumi.InvokeOption) (*LookupEndpointConfigResult, error) {
	var rv LookupEndpointConfigResult
	err := ctx.Invoke("aws-native:sagemaker:getEndpointConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEndpointConfigArgs struct {
	Id string `pulumi:"id"`
}

type LookupEndpointConfigResult struct {
	Id   *string             `pulumi:"id"`
	Tags []EndpointConfigTag `pulumi:"tags"`
}

func LookupEndpointConfigOutput(ctx *pulumi.Context, args LookupEndpointConfigOutputArgs, opts ...pulumi.InvokeOption) LookupEndpointConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEndpointConfigResult, error) {
			args := v.(LookupEndpointConfigArgs)
			r, err := LookupEndpointConfig(ctx, &args, opts...)
			var s LookupEndpointConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEndpointConfigResultOutput)
}

type LookupEndpointConfigOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupEndpointConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointConfigArgs)(nil)).Elem()
}

type LookupEndpointConfigResultOutput struct{ *pulumi.OutputState }

func (LookupEndpointConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointConfigResult)(nil)).Elem()
}

func (o LookupEndpointConfigResultOutput) ToLookupEndpointConfigResultOutput() LookupEndpointConfigResultOutput {
	return o
}

func (o LookupEndpointConfigResultOutput) ToLookupEndpointConfigResultOutputWithContext(ctx context.Context) LookupEndpointConfigResultOutput {
	return o
}

func (o LookupEndpointConfigResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointConfigResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointConfigResultOutput) Tags() EndpointConfigTagArrayOutput {
	return o.ApplyT(func(v LookupEndpointConfigResult) []EndpointConfigTag { return v.Tags }).(EndpointConfigTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEndpointConfigResultOutput{})
}