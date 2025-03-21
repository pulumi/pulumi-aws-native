// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apprunner

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes an AWS App Runner automatic configuration resource that enables automatic scaling of instances used to process web requests. You can share an auto scaling configuration across multiple services.
func LookupAutoScalingConfiguration(ctx *pulumi.Context, args *LookupAutoScalingConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupAutoScalingConfigurationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAutoScalingConfigurationResult
	err := ctx.Invoke("aws-native:apprunner:getAutoScalingConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAutoScalingConfigurationArgs struct {
	// The Amazon Resource Name (ARN) of this auto scaling configuration.
	AutoScalingConfigurationArn string `pulumi:"autoScalingConfigurationArn"`
}

type LookupAutoScalingConfigurationResult struct {
	// The Amazon Resource Name (ARN) of this auto scaling configuration.
	AutoScalingConfigurationArn *string `pulumi:"autoScalingConfigurationArn"`
	// The revision of this auto scaling configuration. It's unique among all the active configurations ("Status": "ACTIVE") that share the same AutoScalingConfigurationName.
	AutoScalingConfigurationRevision *int `pulumi:"autoScalingConfigurationRevision"`
	// It's set to true for the configuration with the highest Revision among all configurations that share the same AutoScalingConfigurationName. It's set to false otherwise. App Runner temporarily doubles the number of provisioned instances during deployments, to maintain the same capacity for both old and new code.
	Latest *bool `pulumi:"latest"`
}

func LookupAutoScalingConfigurationOutput(ctx *pulumi.Context, args LookupAutoScalingConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupAutoScalingConfigurationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAutoScalingConfigurationResultOutput, error) {
			args := v.(LookupAutoScalingConfigurationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:apprunner:getAutoScalingConfiguration", args, LookupAutoScalingConfigurationResultOutput{}, options).(LookupAutoScalingConfigurationResultOutput), nil
		}).(LookupAutoScalingConfigurationResultOutput)
}

type LookupAutoScalingConfigurationOutputArgs struct {
	// The Amazon Resource Name (ARN) of this auto scaling configuration.
	AutoScalingConfigurationArn pulumi.StringInput `pulumi:"autoScalingConfigurationArn"`
}

func (LookupAutoScalingConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutoScalingConfigurationArgs)(nil)).Elem()
}

type LookupAutoScalingConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupAutoScalingConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutoScalingConfigurationResult)(nil)).Elem()
}

func (o LookupAutoScalingConfigurationResultOutput) ToLookupAutoScalingConfigurationResultOutput() LookupAutoScalingConfigurationResultOutput {
	return o
}

func (o LookupAutoScalingConfigurationResultOutput) ToLookupAutoScalingConfigurationResultOutputWithContext(ctx context.Context) LookupAutoScalingConfigurationResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of this auto scaling configuration.
func (o LookupAutoScalingConfigurationResultOutput) AutoScalingConfigurationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoScalingConfigurationResult) *string { return v.AutoScalingConfigurationArn }).(pulumi.StringPtrOutput)
}

// The revision of this auto scaling configuration. It's unique among all the active configurations ("Status": "ACTIVE") that share the same AutoScalingConfigurationName.
func (o LookupAutoScalingConfigurationResultOutput) AutoScalingConfigurationRevision() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAutoScalingConfigurationResult) *int { return v.AutoScalingConfigurationRevision }).(pulumi.IntPtrOutput)
}

// It's set to true for the configuration with the highest Revision among all configurations that share the same AutoScalingConfigurationName. It's set to false otherwise. App Runner temporarily doubles the number of provisioned instances during deployments, to maintain the same capacity for both old and new code.
func (o LookupAutoScalingConfigurationResultOutput) Latest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAutoScalingConfigurationResult) *bool { return v.Latest }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAutoScalingConfigurationResultOutput{})
}
