// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package configuration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Config::RemediationConfiguration
func LookupRemediationConfiguration(ctx *pulumi.Context, args *LookupRemediationConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupRemediationConfigurationResult, error) {
	var rv LookupRemediationConfigurationResult
	err := ctx.Invoke("aws-native:configuration:getRemediationConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRemediationConfigurationArgs struct {
	Id string `pulumi:"id"`
}

type LookupRemediationConfigurationResult struct {
	Automatic                *bool                                      `pulumi:"automatic"`
	ExecutionControls        *RemediationConfigurationExecutionControls `pulumi:"executionControls"`
	Id                       *string                                    `pulumi:"id"`
	MaximumAutomaticAttempts *int                                       `pulumi:"maximumAutomaticAttempts"`
	Parameters               interface{}                                `pulumi:"parameters"`
	ResourceType             *string                                    `pulumi:"resourceType"`
	RetryAttemptSeconds      *int                                       `pulumi:"retryAttemptSeconds"`
	TargetId                 *string                                    `pulumi:"targetId"`
	TargetType               *string                                    `pulumi:"targetType"`
	TargetVersion            *string                                    `pulumi:"targetVersion"`
}

func LookupRemediationConfigurationOutput(ctx *pulumi.Context, args LookupRemediationConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupRemediationConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemediationConfigurationResult, error) {
			args := v.(LookupRemediationConfigurationArgs)
			r, err := LookupRemediationConfiguration(ctx, &args, opts...)
			var s LookupRemediationConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemediationConfigurationResultOutput)
}

type LookupRemediationConfigurationOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupRemediationConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemediationConfigurationArgs)(nil)).Elem()
}

type LookupRemediationConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupRemediationConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemediationConfigurationResult)(nil)).Elem()
}

func (o LookupRemediationConfigurationResultOutput) ToLookupRemediationConfigurationResultOutput() LookupRemediationConfigurationResultOutput {
	return o
}

func (o LookupRemediationConfigurationResultOutput) ToLookupRemediationConfigurationResultOutputWithContext(ctx context.Context) LookupRemediationConfigurationResultOutput {
	return o
}

func (o LookupRemediationConfigurationResultOutput) Automatic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemediationConfigurationResult) *bool { return v.Automatic }).(pulumi.BoolPtrOutput)
}

func (o LookupRemediationConfigurationResultOutput) ExecutionControls() RemediationConfigurationExecutionControlsPtrOutput {
	return o.ApplyT(func(v LookupRemediationConfigurationResult) *RemediationConfigurationExecutionControls {
		return v.ExecutionControls
	}).(RemediationConfigurationExecutionControlsPtrOutput)
}

func (o LookupRemediationConfigurationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationConfigurationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupRemediationConfigurationResultOutput) MaximumAutomaticAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemediationConfigurationResult) *int { return v.MaximumAutomaticAttempts }).(pulumi.IntPtrOutput)
}

func (o LookupRemediationConfigurationResultOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupRemediationConfigurationResult) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

func (o LookupRemediationConfigurationResultOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationConfigurationResult) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o LookupRemediationConfigurationResultOutput) RetryAttemptSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemediationConfigurationResult) *int { return v.RetryAttemptSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemediationConfigurationResultOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationConfigurationResult) *string { return v.TargetId }).(pulumi.StringPtrOutput)
}

func (o LookupRemediationConfigurationResultOutput) TargetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationConfigurationResult) *string { return v.TargetType }).(pulumi.StringPtrOutput)
}

func (o LookupRemediationConfigurationResultOutput) TargetVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemediationConfigurationResult) *string { return v.TargetVersion }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemediationConfigurationResultOutput{})
}