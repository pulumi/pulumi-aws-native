// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package configuration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Config::ConfigRule
func LookupConfigRule(ctx *pulumi.Context, args *LookupConfigRuleArgs, opts ...pulumi.InvokeOption) (*LookupConfigRuleResult, error) {
	var rv LookupConfigRuleResult
	err := ctx.Invoke("aws-native:configuration:getConfigRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigRuleArgs struct {
	ConfigRuleId string `pulumi:"configRuleId"`
}

type LookupConfigRuleResult struct {
	Arn                       *string           `pulumi:"arn"`
	ComplianceType            *string           `pulumi:"complianceType"`
	ConfigRuleId              *string           `pulumi:"configRuleId"`
	Description               *string           `pulumi:"description"`
	InputParameters           interface{}       `pulumi:"inputParameters"`
	MaximumExecutionFrequency *string           `pulumi:"maximumExecutionFrequency"`
	Scope                     *ConfigRuleScope  `pulumi:"scope"`
	Source                    *ConfigRuleSource `pulumi:"source"`
}

func LookupConfigRuleOutput(ctx *pulumi.Context, args LookupConfigRuleOutputArgs, opts ...pulumi.InvokeOption) LookupConfigRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigRuleResult, error) {
			args := v.(LookupConfigRuleArgs)
			r, err := LookupConfigRule(ctx, &args, opts...)
			var s LookupConfigRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigRuleResultOutput)
}

type LookupConfigRuleOutputArgs struct {
	ConfigRuleId pulumi.StringInput `pulumi:"configRuleId"`
}

func (LookupConfigRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigRuleArgs)(nil)).Elem()
}

type LookupConfigRuleResultOutput struct{ *pulumi.OutputState }

func (LookupConfigRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigRuleResult)(nil)).Elem()
}

func (o LookupConfigRuleResultOutput) ToLookupConfigRuleResultOutput() LookupConfigRuleResultOutput {
	return o
}

func (o LookupConfigRuleResultOutput) ToLookupConfigRuleResultOutputWithContext(ctx context.Context) LookupConfigRuleResultOutput {
	return o
}

func (o LookupConfigRuleResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigRuleResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupConfigRuleResultOutput) ComplianceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigRuleResult) *string { return v.ComplianceType }).(pulumi.StringPtrOutput)
}

func (o LookupConfigRuleResultOutput) ConfigRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigRuleResult) *string { return v.ConfigRuleId }).(pulumi.StringPtrOutput)
}

func (o LookupConfigRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupConfigRuleResultOutput) InputParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupConfigRuleResult) interface{} { return v.InputParameters }).(pulumi.AnyOutput)
}

func (o LookupConfigRuleResultOutput) MaximumExecutionFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigRuleResult) *string { return v.MaximumExecutionFrequency }).(pulumi.StringPtrOutput)
}

func (o LookupConfigRuleResultOutput) Scope() ConfigRuleScopePtrOutput {
	return o.ApplyT(func(v LookupConfigRuleResult) *ConfigRuleScope { return v.Scope }).(ConfigRuleScopePtrOutput)
}

func (o LookupConfigRuleResultOutput) Source() ConfigRuleSourcePtrOutput {
	return o.ApplyT(func(v LookupConfigRuleResult) *ConfigRuleSource { return v.Source }).(ConfigRuleSourcePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigRuleResultOutput{})
}