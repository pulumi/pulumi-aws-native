// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package xray

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This schema provides construct and validation rules for AWS-XRay SamplingRule resource parameters.
type SamplingRule struct {
	pulumi.CustomResourceState

	RuleARN            pulumi.StringOutput           `pulumi:"ruleARN"`
	RuleName           pulumi.StringPtrOutput        `pulumi:"ruleName"`
	SamplingRule       SamplingRuleTypePtrOutput     `pulumi:"samplingRule"`
	SamplingRuleRecord SamplingRuleRecordPtrOutput   `pulumi:"samplingRuleRecord"`
	SamplingRuleUpdate SamplingRuleUpdatePtrOutput   `pulumi:"samplingRuleUpdate"`
	Tags               TagsItemPropertiesArrayOutput `pulumi:"tags"`
}

// NewSamplingRule registers a new resource with the given unique name, arguments, and options.
func NewSamplingRule(ctx *pulumi.Context,
	name string, args *SamplingRuleArgs, opts ...pulumi.ResourceOption) (*SamplingRule, error) {
	if args == nil {
		args = &SamplingRuleArgs{}
	}

	var resource SamplingRule
	err := ctx.RegisterResource("aws-native:xray:SamplingRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSamplingRule gets an existing SamplingRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSamplingRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SamplingRuleState, opts ...pulumi.ResourceOption) (*SamplingRule, error) {
	var resource SamplingRule
	err := ctx.ReadResource("aws-native:xray:SamplingRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SamplingRule resources.
type samplingRuleState struct {
}

type SamplingRuleState struct {
}

func (SamplingRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*samplingRuleState)(nil)).Elem()
}

type samplingRuleArgs struct {
	RuleName           *string              `pulumi:"ruleName"`
	SamplingRule       *SamplingRuleType    `pulumi:"samplingRule"`
	SamplingRuleRecord *SamplingRuleRecord  `pulumi:"samplingRuleRecord"`
	SamplingRuleUpdate *SamplingRuleUpdate  `pulumi:"samplingRuleUpdate"`
	Tags               []TagsItemProperties `pulumi:"tags"`
}

// The set of arguments for constructing a SamplingRule resource.
type SamplingRuleArgs struct {
	RuleName           pulumi.StringPtrInput
	SamplingRule       SamplingRuleTypePtrInput
	SamplingRuleRecord SamplingRuleRecordPtrInput
	SamplingRuleUpdate SamplingRuleUpdatePtrInput
	Tags               TagsItemPropertiesArrayInput
}

func (SamplingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*samplingRuleArgs)(nil)).Elem()
}

type SamplingRuleInput interface {
	pulumi.Input

	ToSamplingRuleOutput() SamplingRuleOutput
	ToSamplingRuleOutputWithContext(ctx context.Context) SamplingRuleOutput
}

func (*SamplingRule) ElementType() reflect.Type {
	return reflect.TypeOf((**SamplingRule)(nil)).Elem()
}

func (i *SamplingRule) ToSamplingRuleOutput() SamplingRuleOutput {
	return i.ToSamplingRuleOutputWithContext(context.Background())
}

func (i *SamplingRule) ToSamplingRuleOutputWithContext(ctx context.Context) SamplingRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamplingRuleOutput)
}

type SamplingRuleOutput struct{ *pulumi.OutputState }

func (SamplingRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SamplingRule)(nil)).Elem()
}

func (o SamplingRuleOutput) ToSamplingRuleOutput() SamplingRuleOutput {
	return o
}

func (o SamplingRuleOutput) ToSamplingRuleOutputWithContext(ctx context.Context) SamplingRuleOutput {
	return o
}

func (o SamplingRuleOutput) RuleARN() pulumi.StringOutput {
	return o.ApplyT(func(v *SamplingRule) pulumi.StringOutput { return v.RuleARN }).(pulumi.StringOutput)
}

func (o SamplingRuleOutput) RuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SamplingRule) pulumi.StringPtrOutput { return v.RuleName }).(pulumi.StringPtrOutput)
}

func (o SamplingRuleOutput) SamplingRule() SamplingRuleTypePtrOutput {
	return o.ApplyT(func(v *SamplingRule) SamplingRuleTypePtrOutput { return v.SamplingRule }).(SamplingRuleTypePtrOutput)
}

func (o SamplingRuleOutput) SamplingRuleRecord() SamplingRuleRecordPtrOutput {
	return o.ApplyT(func(v *SamplingRule) SamplingRuleRecordPtrOutput { return v.SamplingRuleRecord }).(SamplingRuleRecordPtrOutput)
}

func (o SamplingRuleOutput) SamplingRuleUpdate() SamplingRuleUpdatePtrOutput {
	return o.ApplyT(func(v *SamplingRule) SamplingRuleUpdatePtrOutput { return v.SamplingRuleUpdate }).(SamplingRuleUpdatePtrOutput)
}

func (o SamplingRuleOutput) Tags() TagsItemPropertiesArrayOutput {
	return o.ApplyT(func(v *SamplingRule) TagsItemPropertiesArrayOutput { return v.Tags }).(TagsItemPropertiesArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SamplingRuleInput)(nil)).Elem(), &SamplingRule{})
	pulumi.RegisterOutputType(SamplingRuleOutput{})
}