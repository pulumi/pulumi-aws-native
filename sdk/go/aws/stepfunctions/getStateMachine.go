// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package stepfunctions

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for StateMachine
func LookupStateMachine(ctx *pulumi.Context, args *LookupStateMachineArgs, opts ...pulumi.InvokeOption) (*LookupStateMachineResult, error) {
	var rv LookupStateMachineResult
	err := ctx.Invoke("aws-native:stepfunctions:getStateMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStateMachineArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupStateMachineResult struct {
	Arn                  *string                           `pulumi:"arn"`
	DefinitionString     *string                           `pulumi:"definitionString"`
	LoggingConfiguration *StateMachineLoggingConfiguration `pulumi:"loggingConfiguration"`
	Name                 *string                           `pulumi:"name"`
	RoleArn              *string                           `pulumi:"roleArn"`
	Tags                 []StateMachineTagsEntry           `pulumi:"tags"`
	TracingConfiguration *StateMachineTracingConfiguration `pulumi:"tracingConfiguration"`
}

func LookupStateMachineOutput(ctx *pulumi.Context, args LookupStateMachineOutputArgs, opts ...pulumi.InvokeOption) LookupStateMachineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStateMachineResult, error) {
			args := v.(LookupStateMachineArgs)
			r, err := LookupStateMachine(ctx, &args, opts...)
			var s LookupStateMachineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStateMachineResultOutput)
}

type LookupStateMachineOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupStateMachineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStateMachineArgs)(nil)).Elem()
}

type LookupStateMachineResultOutput struct{ *pulumi.OutputState }

func (LookupStateMachineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStateMachineResult)(nil)).Elem()
}

func (o LookupStateMachineResultOutput) ToLookupStateMachineResultOutput() LookupStateMachineResultOutput {
	return o
}

func (o LookupStateMachineResultOutput) ToLookupStateMachineResultOutputWithContext(ctx context.Context) LookupStateMachineResultOutput {
	return o
}

func (o LookupStateMachineResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStateMachineResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupStateMachineResultOutput) DefinitionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStateMachineResult) *string { return v.DefinitionString }).(pulumi.StringPtrOutput)
}

func (o LookupStateMachineResultOutput) LoggingConfiguration() StateMachineLoggingConfigurationPtrOutput {
	return o.ApplyT(func(v LookupStateMachineResult) *StateMachineLoggingConfiguration { return v.LoggingConfiguration }).(StateMachineLoggingConfigurationPtrOutput)
}

func (o LookupStateMachineResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStateMachineResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupStateMachineResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStateMachineResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupStateMachineResultOutput) Tags() StateMachineTagsEntryArrayOutput {
	return o.ApplyT(func(v LookupStateMachineResult) []StateMachineTagsEntry { return v.Tags }).(StateMachineTagsEntryArrayOutput)
}

func (o LookupStateMachineResultOutput) TracingConfiguration() StateMachineTracingConfigurationPtrOutput {
	return o.ApplyT(func(v LookupStateMachineResult) *StateMachineTracingConfiguration { return v.TracingConfiguration }).(StateMachineTracingConfigurationPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStateMachineResultOutput{})
}