// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityhub

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::SecurityHub::AutomationRuleV2
type AutomationRuleV2 struct {
	pulumi.CustomResourceState

	// A list of actions to be performed when the rule criteria is met
	Actions AutomationRuleV2AutomationRulesActionV2ArrayOutput `pulumi:"actions"`
	// The timestamp when the V2 automation rule was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The filtering type and configuration of the automation rule.
	Criteria AutomationRuleV2CriteriaOutput `pulumi:"criteria"`
	// A description of the automation rule
	Description pulumi.StringOutput `pulumi:"description"`
	// The ARN of the automation rule
	RuleArn pulumi.StringOutput `pulumi:"ruleArn"`
	// The ID of the automation rule
	RuleId pulumi.StringOutput `pulumi:"ruleId"`
	// The name of the automation rule
	RuleName pulumi.StringOutput `pulumi:"ruleName"`
	// The value for the rule priority
	RuleOrder pulumi.Float64Output `pulumi:"ruleOrder"`
	// The status of the automation rule
	RuleStatus AutomationRuleV2RuleStatusPtrOutput `pulumi:"ruleStatus"`
	// A list of key-value pairs associated with the V2 automation rule.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The timestamp when the V2 automation rule was updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewAutomationRuleV2 registers a new resource with the given unique name, arguments, and options.
func NewAutomationRuleV2(ctx *pulumi.Context,
	name string, args *AutomationRuleV2Args, opts ...pulumi.ResourceOption) (*AutomationRuleV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	if args.Criteria == nil {
		return nil, errors.New("invalid value for required argument 'Criteria'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.RuleName == nil {
		return nil, errors.New("invalid value for required argument 'RuleName'")
	}
	if args.RuleOrder == nil {
		return nil, errors.New("invalid value for required argument 'RuleOrder'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AutomationRuleV2
	err := ctx.RegisterResource("aws-native:securityhub:AutomationRuleV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutomationRuleV2 gets an existing AutomationRuleV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutomationRuleV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutomationRuleV2State, opts ...pulumi.ResourceOption) (*AutomationRuleV2, error) {
	var resource AutomationRuleV2
	err := ctx.ReadResource("aws-native:securityhub:AutomationRuleV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutomationRuleV2 resources.
type automationRuleV2State struct {
}

type AutomationRuleV2State struct {
}

func (AutomationRuleV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*automationRuleV2State)(nil)).Elem()
}

type automationRuleV2Args struct {
	// A list of actions to be performed when the rule criteria is met
	Actions []AutomationRuleV2AutomationRulesActionV2 `pulumi:"actions"`
	// The filtering type and configuration of the automation rule.
	Criteria AutomationRuleV2Criteria `pulumi:"criteria"`
	// A description of the automation rule
	Description string `pulumi:"description"`
	// The name of the automation rule
	RuleName string `pulumi:"ruleName"`
	// The value for the rule priority
	RuleOrder float64 `pulumi:"ruleOrder"`
	// The status of the automation rule
	RuleStatus *AutomationRuleV2RuleStatus `pulumi:"ruleStatus"`
	// A list of key-value pairs associated with the V2 automation rule.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AutomationRuleV2 resource.
type AutomationRuleV2Args struct {
	// A list of actions to be performed when the rule criteria is met
	Actions AutomationRuleV2AutomationRulesActionV2ArrayInput
	// The filtering type and configuration of the automation rule.
	Criteria AutomationRuleV2CriteriaInput
	// A description of the automation rule
	Description pulumi.StringInput
	// The name of the automation rule
	RuleName pulumi.StringInput
	// The value for the rule priority
	RuleOrder pulumi.Float64Input
	// The status of the automation rule
	RuleStatus AutomationRuleV2RuleStatusPtrInput
	// A list of key-value pairs associated with the V2 automation rule.
	Tags pulumi.StringMapInput
}

func (AutomationRuleV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*automationRuleV2Args)(nil)).Elem()
}

type AutomationRuleV2Input interface {
	pulumi.Input

	ToAutomationRuleV2Output() AutomationRuleV2Output
	ToAutomationRuleV2OutputWithContext(ctx context.Context) AutomationRuleV2Output
}

func (*AutomationRuleV2) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomationRuleV2)(nil)).Elem()
}

func (i *AutomationRuleV2) ToAutomationRuleV2Output() AutomationRuleV2Output {
	return i.ToAutomationRuleV2OutputWithContext(context.Background())
}

func (i *AutomationRuleV2) ToAutomationRuleV2OutputWithContext(ctx context.Context) AutomationRuleV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleV2Output)
}

type AutomationRuleV2Output struct{ *pulumi.OutputState }

func (AutomationRuleV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomationRuleV2)(nil)).Elem()
}

func (o AutomationRuleV2Output) ToAutomationRuleV2Output() AutomationRuleV2Output {
	return o
}

func (o AutomationRuleV2Output) ToAutomationRuleV2OutputWithContext(ctx context.Context) AutomationRuleV2Output {
	return o
}

// A list of actions to be performed when the rule criteria is met
func (o AutomationRuleV2Output) Actions() AutomationRuleV2AutomationRulesActionV2ArrayOutput {
	return o.ApplyT(func(v *AutomationRuleV2) AutomationRuleV2AutomationRulesActionV2ArrayOutput { return v.Actions }).(AutomationRuleV2AutomationRulesActionV2ArrayOutput)
}

// The timestamp when the V2 automation rule was created.
func (o AutomationRuleV2Output) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationRuleV2) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The filtering type and configuration of the automation rule.
func (o AutomationRuleV2Output) Criteria() AutomationRuleV2CriteriaOutput {
	return o.ApplyT(func(v *AutomationRuleV2) AutomationRuleV2CriteriaOutput { return v.Criteria }).(AutomationRuleV2CriteriaOutput)
}

// A description of the automation rule
func (o AutomationRuleV2Output) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationRuleV2) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The ARN of the automation rule
func (o AutomationRuleV2Output) RuleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationRuleV2) pulumi.StringOutput { return v.RuleArn }).(pulumi.StringOutput)
}

// The ID of the automation rule
func (o AutomationRuleV2Output) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationRuleV2) pulumi.StringOutput { return v.RuleId }).(pulumi.StringOutput)
}

// The name of the automation rule
func (o AutomationRuleV2Output) RuleName() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationRuleV2) pulumi.StringOutput { return v.RuleName }).(pulumi.StringOutput)
}

// The value for the rule priority
func (o AutomationRuleV2Output) RuleOrder() pulumi.Float64Output {
	return o.ApplyT(func(v *AutomationRuleV2) pulumi.Float64Output { return v.RuleOrder }).(pulumi.Float64Output)
}

// The status of the automation rule
func (o AutomationRuleV2Output) RuleStatus() AutomationRuleV2RuleStatusPtrOutput {
	return o.ApplyT(func(v *AutomationRuleV2) AutomationRuleV2RuleStatusPtrOutput { return v.RuleStatus }).(AutomationRuleV2RuleStatusPtrOutput)
}

// A list of key-value pairs associated with the V2 automation rule.
func (o AutomationRuleV2Output) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AutomationRuleV2) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The timestamp when the V2 automation rule was updated.
func (o AutomationRuleV2Output) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationRuleV2) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRuleV2Input)(nil)).Elem(), &AutomationRuleV2{})
	pulumi.RegisterOutputType(AutomationRuleV2Output{})
}
