// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package robomaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS::RoboMaker::SimulationApplicationVersion resource creates an AWS RoboMaker SimulationApplicationVersion. This helps you control which code your simulation uses.
func LookupSimulationApplicationVersion(ctx *pulumi.Context, args *LookupSimulationApplicationVersionArgs, opts ...pulumi.InvokeOption) (*LookupSimulationApplicationVersionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSimulationApplicationVersionResult
	err := ctx.Invoke("aws-native:robomaker:getSimulationApplicationVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSimulationApplicationVersionArgs struct {
	// The Amazon Resource Name (ARN) of the simulation application version.
	Arn string `pulumi:"arn"`
}

type LookupSimulationApplicationVersionResult struct {
	// The simulation application version.
	ApplicationVersion *string `pulumi:"applicationVersion"`
	// The Amazon Resource Name (ARN) of the simulation application version.
	Arn *string `pulumi:"arn"`
}

func LookupSimulationApplicationVersionOutput(ctx *pulumi.Context, args LookupSimulationApplicationVersionOutputArgs, opts ...pulumi.InvokeOption) LookupSimulationApplicationVersionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSimulationApplicationVersionResultOutput, error) {
			args := v.(LookupSimulationApplicationVersionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:robomaker:getSimulationApplicationVersion", args, LookupSimulationApplicationVersionResultOutput{}, options).(LookupSimulationApplicationVersionResultOutput), nil
		}).(LookupSimulationApplicationVersionResultOutput)
}

type LookupSimulationApplicationVersionOutputArgs struct {
	// The Amazon Resource Name (ARN) of the simulation application version.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupSimulationApplicationVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSimulationApplicationVersionArgs)(nil)).Elem()
}

type LookupSimulationApplicationVersionResultOutput struct{ *pulumi.OutputState }

func (LookupSimulationApplicationVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSimulationApplicationVersionResult)(nil)).Elem()
}

func (o LookupSimulationApplicationVersionResultOutput) ToLookupSimulationApplicationVersionResultOutput() LookupSimulationApplicationVersionResultOutput {
	return o
}

func (o LookupSimulationApplicationVersionResultOutput) ToLookupSimulationApplicationVersionResultOutputWithContext(ctx context.Context) LookupSimulationApplicationVersionResultOutput {
	return o
}

// The simulation application version.
func (o LookupSimulationApplicationVersionResultOutput) ApplicationVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimulationApplicationVersionResult) *string { return v.ApplicationVersion }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the simulation application version.
func (o LookupSimulationApplicationVersionResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimulationApplicationVersionResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSimulationApplicationVersionResultOutput{})
}
