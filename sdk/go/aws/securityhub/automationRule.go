// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityhub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::SecurityHub::AutomationRule resource represents the Automation Rule in your account. One rule resource is created for each Automation Rule in which you configure rule criteria and actions.
type AutomationRule struct {
	pulumi.CustomResourceState

	Actions AutomationRulesActionArrayOutput `pulumi:"actions"`
	// The date and time when Automation Rule was created
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The identifier by which created the rule
	CreatedBy pulumi.StringOutput `pulumi:"createdBy"`
	// The rule criteria for evaluating findings
	Criteria AutomationRulesFindingFiltersPtrOutput `pulumi:"criteria"`
	// Rule description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// If Rule is a terminal rule
	IsTerminal pulumi.BoolPtrOutput `pulumi:"isTerminal"`
	// An Automation Rule Arn is automatically created
	RuleArn pulumi.StringOutput `pulumi:"ruleArn"`
	// Rule name
	RuleName pulumi.StringPtrOutput `pulumi:"ruleName"`
	// Rule order value
	RuleOrder pulumi.IntPtrOutput `pulumi:"ruleOrder"`
	// Status of the Rule upon creation
	RuleStatus AutomationRuleRuleStatusPtrOutput `pulumi:"ruleStatus"`
	Tags       AutomationRuleTagsPtrOutput       `pulumi:"tags"`
	// The date and time when Automation Rule was last updated
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewAutomationRule registers a new resource with the given unique name, arguments, and options.
func NewAutomationRule(ctx *pulumi.Context,
	name string, args *AutomationRuleArgs, opts ...pulumi.ResourceOption) (*AutomationRule, error) {
	if args == nil {
		args = &AutomationRuleArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AutomationRule
	err := ctx.RegisterResource("aws-native:securityhub:AutomationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutomationRule gets an existing AutomationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutomationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutomationRuleState, opts ...pulumi.ResourceOption) (*AutomationRule, error) {
	var resource AutomationRule
	err := ctx.ReadResource("aws-native:securityhub:AutomationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutomationRule resources.
type automationRuleState struct {
}

type AutomationRuleState struct {
}

func (AutomationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*automationRuleState)(nil)).Elem()
}

type automationRuleArgs struct {
	Actions []AutomationRulesAction `pulumi:"actions"`
	// The rule criteria for evaluating findings
	Criteria *AutomationRulesFindingFilters `pulumi:"criteria"`
	// Rule description
	Description *string `pulumi:"description"`
	// If Rule is a terminal rule
	IsTerminal *bool `pulumi:"isTerminal"`
	// Rule name
	RuleName *string `pulumi:"ruleName"`
	// Rule order value
	RuleOrder *int `pulumi:"ruleOrder"`
	// Status of the Rule upon creation
	RuleStatus *AutomationRuleRuleStatus `pulumi:"ruleStatus"`
	Tags       *AutomationRuleTags       `pulumi:"tags"`
}

// The set of arguments for constructing a AutomationRule resource.
type AutomationRuleArgs struct {
	Actions AutomationRulesActionArrayInput
	// The rule criteria for evaluating findings
	Criteria AutomationRulesFindingFiltersPtrInput
	// Rule description
	Description pulumi.StringPtrInput
	// If Rule is a terminal rule
	IsTerminal pulumi.BoolPtrInput
	// Rule name
	RuleName pulumi.StringPtrInput
	// Rule order value
	RuleOrder pulumi.IntPtrInput
	// Status of the Rule upon creation
	RuleStatus AutomationRuleRuleStatusPtrInput
	Tags       AutomationRuleTagsPtrInput
}

func (AutomationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*automationRuleArgs)(nil)).Elem()
}

type AutomationRuleInput interface {
	pulumi.Input

	ToAutomationRuleOutput() AutomationRuleOutput
	ToAutomationRuleOutputWithContext(ctx context.Context) AutomationRuleOutput
}

func (*AutomationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomationRule)(nil)).Elem()
}

func (i *AutomationRule) ToAutomationRuleOutput() AutomationRuleOutput {
	return i.ToAutomationRuleOutputWithContext(context.Background())
}

func (i *AutomationRule) ToAutomationRuleOutputWithContext(ctx context.Context) AutomationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleOutput)
}

type AutomationRuleOutput struct{ *pulumi.OutputState }

func (AutomationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomationRule)(nil)).Elem()
}

func (o AutomationRuleOutput) ToAutomationRuleOutput() AutomationRuleOutput {
	return o
}

func (o AutomationRuleOutput) ToAutomationRuleOutputWithContext(ctx context.Context) AutomationRuleOutput {
	return o
}

func (o AutomationRuleOutput) Actions() AutomationRulesActionArrayOutput {
	return o.ApplyT(func(v *AutomationRule) AutomationRulesActionArrayOutput { return v.Actions }).(AutomationRulesActionArrayOutput)
}

// The date and time when Automation Rule was created
func (o AutomationRuleOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The identifier by which created the rule
func (o AutomationRuleOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.StringOutput { return v.CreatedBy }).(pulumi.StringOutput)
}

// The rule criteria for evaluating findings
func (o AutomationRuleOutput) Criteria() AutomationRulesFindingFiltersPtrOutput {
	return o.ApplyT(func(v *AutomationRule) AutomationRulesFindingFiltersPtrOutput { return v.Criteria }).(AutomationRulesFindingFiltersPtrOutput)
}

// Rule description
func (o AutomationRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// If Rule is a terminal rule
func (o AutomationRuleOutput) IsTerminal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.BoolPtrOutput { return v.IsTerminal }).(pulumi.BoolPtrOutput)
}

// An Automation Rule Arn is automatically created
func (o AutomationRuleOutput) RuleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.StringOutput { return v.RuleArn }).(pulumi.StringOutput)
}

// Rule name
func (o AutomationRuleOutput) RuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.StringPtrOutput { return v.RuleName }).(pulumi.StringPtrOutput)
}

// Rule order value
func (o AutomationRuleOutput) RuleOrder() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.IntPtrOutput { return v.RuleOrder }).(pulumi.IntPtrOutput)
}

// Status of the Rule upon creation
func (o AutomationRuleOutput) RuleStatus() AutomationRuleRuleStatusPtrOutput {
	return o.ApplyT(func(v *AutomationRule) AutomationRuleRuleStatusPtrOutput { return v.RuleStatus }).(AutomationRuleRuleStatusPtrOutput)
}

func (o AutomationRuleOutput) Tags() AutomationRuleTagsPtrOutput {
	return o.ApplyT(func(v *AutomationRule) AutomationRuleTagsPtrOutput { return v.Tags }).(AutomationRuleTagsPtrOutput)
}

// The date and time when Automation Rule was last updated
func (o AutomationRuleOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationRule) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRuleInput)(nil)).Elem(), &AutomationRule{})
	pulumi.RegisterOutputType(AutomationRuleOutput{})
}