// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package medialive

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::MediaLive::Multiplex
func LookupMultiplex(ctx *pulumi.Context, args *LookupMultiplexArgs, opts ...pulumi.InvokeOption) (*LookupMultiplexResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupMultiplexResult
	err := ctx.Invoke("aws-native:medialive:getMultiplex", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMultiplexArgs struct {
	// The unique id of the multiplex.
	Id string `pulumi:"id"`
}

type LookupMultiplexResult struct {
	// The unique arn of the multiplex.
	Arn *string `pulumi:"arn"`
	// A list of the multiplex output destinations.
	Destinations []MultiplexOutputDestination `pulumi:"destinations"`
	// The unique id of the multiplex.
	Id *string `pulumi:"id"`
	// Configuration for a multiplex event.
	MultiplexSettings *MultiplexSettings `pulumi:"multiplexSettings"`
	// Name of multiplex.
	Name *string `pulumi:"name"`
	// The number of currently healthy pipelines.
	PipelinesRunningCount *int `pulumi:"pipelinesRunningCount"`
	// The number of programs in the multiplex.
	ProgramCount *int `pulumi:"programCount"`
	// The current state of the multiplex.
	State *MultiplexStateEnum `pulumi:"state"`
	// A collection of key-value pairs.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupMultiplexOutput(ctx *pulumi.Context, args LookupMultiplexOutputArgs, opts ...pulumi.InvokeOption) LookupMultiplexResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupMultiplexResultOutput, error) {
			args := v.(LookupMultiplexArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:medialive:getMultiplex", args, LookupMultiplexResultOutput{}, options).(LookupMultiplexResultOutput), nil
		}).(LookupMultiplexResultOutput)
}

type LookupMultiplexOutputArgs struct {
	// The unique id of the multiplex.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupMultiplexOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMultiplexArgs)(nil)).Elem()
}

type LookupMultiplexResultOutput struct{ *pulumi.OutputState }

func (LookupMultiplexResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMultiplexResult)(nil)).Elem()
}

func (o LookupMultiplexResultOutput) ToLookupMultiplexResultOutput() LookupMultiplexResultOutput {
	return o
}

func (o LookupMultiplexResultOutput) ToLookupMultiplexResultOutputWithContext(ctx context.Context) LookupMultiplexResultOutput {
	return o
}

// The unique arn of the multiplex.
func (o LookupMultiplexResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMultiplexResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// A list of the multiplex output destinations.
func (o LookupMultiplexResultOutput) Destinations() MultiplexOutputDestinationArrayOutput {
	return o.ApplyT(func(v LookupMultiplexResult) []MultiplexOutputDestination { return v.Destinations }).(MultiplexOutputDestinationArrayOutput)
}

// The unique id of the multiplex.
func (o LookupMultiplexResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMultiplexResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Configuration for a multiplex event.
func (o LookupMultiplexResultOutput) MultiplexSettings() MultiplexSettingsPtrOutput {
	return o.ApplyT(func(v LookupMultiplexResult) *MultiplexSettings { return v.MultiplexSettings }).(MultiplexSettingsPtrOutput)
}

// Name of multiplex.
func (o LookupMultiplexResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMultiplexResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The number of currently healthy pipelines.
func (o LookupMultiplexResultOutput) PipelinesRunningCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupMultiplexResult) *int { return v.PipelinesRunningCount }).(pulumi.IntPtrOutput)
}

// The number of programs in the multiplex.
func (o LookupMultiplexResultOutput) ProgramCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupMultiplexResult) *int { return v.ProgramCount }).(pulumi.IntPtrOutput)
}

// The current state of the multiplex.
func (o LookupMultiplexResultOutput) State() MultiplexStateEnumPtrOutput {
	return o.ApplyT(func(v LookupMultiplexResult) *MultiplexStateEnum { return v.State }).(MultiplexStateEnumPtrOutput)
}

// A collection of key-value pairs.
func (o LookupMultiplexResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupMultiplexResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMultiplexResultOutput{})
}
