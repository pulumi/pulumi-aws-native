// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::StudioLifecycleConfig
func LookupStudioLifecycleConfig(ctx *pulumi.Context, args *LookupStudioLifecycleConfigArgs, opts ...pulumi.InvokeOption) (*LookupStudioLifecycleConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStudioLifecycleConfigResult
	err := ctx.Invoke("aws-native:sagemaker:getStudioLifecycleConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStudioLifecycleConfigArgs struct {
	// The name of the Amazon SageMaker Studio Lifecycle Configuration.
	StudioLifecycleConfigName string `pulumi:"studioLifecycleConfigName"`
}

type LookupStudioLifecycleConfigResult struct {
	// The Amazon Resource Name (ARN) of the Lifecycle Configuration.
	StudioLifecycleConfigArn *string `pulumi:"studioLifecycleConfigArn"`
}

func LookupStudioLifecycleConfigOutput(ctx *pulumi.Context, args LookupStudioLifecycleConfigOutputArgs, opts ...pulumi.InvokeOption) LookupStudioLifecycleConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStudioLifecycleConfigResult, error) {
			args := v.(LookupStudioLifecycleConfigArgs)
			r, err := LookupStudioLifecycleConfig(ctx, &args, opts...)
			var s LookupStudioLifecycleConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStudioLifecycleConfigResultOutput)
}

type LookupStudioLifecycleConfigOutputArgs struct {
	// The name of the Amazon SageMaker Studio Lifecycle Configuration.
	StudioLifecycleConfigName pulumi.StringInput `pulumi:"studioLifecycleConfigName"`
}

func (LookupStudioLifecycleConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStudioLifecycleConfigArgs)(nil)).Elem()
}

type LookupStudioLifecycleConfigResultOutput struct{ *pulumi.OutputState }

func (LookupStudioLifecycleConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStudioLifecycleConfigResult)(nil)).Elem()
}

func (o LookupStudioLifecycleConfigResultOutput) ToLookupStudioLifecycleConfigResultOutput() LookupStudioLifecycleConfigResultOutput {
	return o
}

func (o LookupStudioLifecycleConfigResultOutput) ToLookupStudioLifecycleConfigResultOutputWithContext(ctx context.Context) LookupStudioLifecycleConfigResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the Lifecycle Configuration.
func (o LookupStudioLifecycleConfigResultOutput) StudioLifecycleConfigArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioLifecycleConfigResult) *string { return v.StudioLifecycleConfigArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStudioLifecycleConfigResultOutput{})
}