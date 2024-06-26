// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkfirewall

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource type definition for AWS::NetworkFirewall::RuleGroup
type RuleGroup struct {
	pulumi.CustomResourceState

	// The maximum operating resources that this rule group can use. You can't change a rule group's capacity setting after you create the rule group. When you update a rule group, you are limited to this capacity. When you reference a rule group from a firewall policy, Network Firewall reserves this capacity for the rule group.
	Capacity pulumi.IntOutput `pulumi:"capacity"`
	// A description of the rule group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// An object that defines the rule group rules.
	RuleGroup RuleGroupTypePtrOutput `pulumi:"ruleGroup"`
	// The Amazon Resource Name (ARN) of the `RuleGroup` .
	RuleGroupArn pulumi.StringOutput `pulumi:"ruleGroupArn"`
	// The unique ID of the `RuleGroup` resource.
	RuleGroupId pulumi.StringOutput `pulumi:"ruleGroupId"`
	// The descriptive name of the rule group. You can't change the name of a rule group after you create it.
	RuleGroupName pulumi.StringOutput `pulumi:"ruleGroupName"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// Indicates whether the rule group is stateless or stateful. If the rule group is stateless, it contains
	// stateless rules. If it is stateful, it contains stateful rules.
	Type RuleGroupTypeEnumOutput `pulumi:"type"`
}

// NewRuleGroup registers a new resource with the given unique name, arguments, and options.
func NewRuleGroup(ctx *pulumi.Context,
	name string, args *RuleGroupArgs, opts ...pulumi.ResourceOption) (*RuleGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Capacity == nil {
		return nil, errors.New("invalid value for required argument 'Capacity'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"capacity",
		"ruleGroupName",
		"type",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RuleGroup
	err := ctx.RegisterResource("aws-native:networkfirewall:RuleGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuleGroup gets an existing RuleGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuleGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleGroupState, opts ...pulumi.ResourceOption) (*RuleGroup, error) {
	var resource RuleGroup
	err := ctx.ReadResource("aws-native:networkfirewall:RuleGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuleGroup resources.
type ruleGroupState struct {
}

type RuleGroupState struct {
}

func (RuleGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleGroupState)(nil)).Elem()
}

type ruleGroupArgs struct {
	// The maximum operating resources that this rule group can use. You can't change a rule group's capacity setting after you create the rule group. When you update a rule group, you are limited to this capacity. When you reference a rule group from a firewall policy, Network Firewall reserves this capacity for the rule group.
	Capacity int `pulumi:"capacity"`
	// A description of the rule group.
	Description *string `pulumi:"description"`
	// An object that defines the rule group rules.
	RuleGroup *RuleGroupType `pulumi:"ruleGroup"`
	// The descriptive name of the rule group. You can't change the name of a rule group after you create it.
	RuleGroupName *string `pulumi:"ruleGroupName"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags []aws.Tag `pulumi:"tags"`
	// Indicates whether the rule group is stateless or stateful. If the rule group is stateless, it contains
	// stateless rules. If it is stateful, it contains stateful rules.
	Type RuleGroupTypeEnum `pulumi:"type"`
}

// The set of arguments for constructing a RuleGroup resource.
type RuleGroupArgs struct {
	// The maximum operating resources that this rule group can use. You can't change a rule group's capacity setting after you create the rule group. When you update a rule group, you are limited to this capacity. When you reference a rule group from a firewall policy, Network Firewall reserves this capacity for the rule group.
	Capacity pulumi.IntInput
	// A description of the rule group.
	Description pulumi.StringPtrInput
	// An object that defines the rule group rules.
	RuleGroup RuleGroupTypePtrInput
	// The descriptive name of the rule group. You can't change the name of a rule group after you create it.
	RuleGroupName pulumi.StringPtrInput
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags aws.TagArrayInput
	// Indicates whether the rule group is stateless or stateful. If the rule group is stateless, it contains
	// stateless rules. If it is stateful, it contains stateful rules.
	Type RuleGroupTypeEnumInput
}

func (RuleGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleGroupArgs)(nil)).Elem()
}

type RuleGroupInput interface {
	pulumi.Input

	ToRuleGroupOutput() RuleGroupOutput
	ToRuleGroupOutputWithContext(ctx context.Context) RuleGroupOutput
}

func (*RuleGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleGroup)(nil)).Elem()
}

func (i *RuleGroup) ToRuleGroupOutput() RuleGroupOutput {
	return i.ToRuleGroupOutputWithContext(context.Background())
}

func (i *RuleGroup) ToRuleGroupOutputWithContext(ctx context.Context) RuleGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleGroupOutput)
}

type RuleGroupOutput struct{ *pulumi.OutputState }

func (RuleGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleGroup)(nil)).Elem()
}

func (o RuleGroupOutput) ToRuleGroupOutput() RuleGroupOutput {
	return o
}

func (o RuleGroupOutput) ToRuleGroupOutputWithContext(ctx context.Context) RuleGroupOutput {
	return o
}

// The maximum operating resources that this rule group can use. You can't change a rule group's capacity setting after you create the rule group. When you update a rule group, you are limited to this capacity. When you reference a rule group from a firewall policy, Network Firewall reserves this capacity for the rule group.
func (o RuleGroupOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v *RuleGroup) pulumi.IntOutput { return v.Capacity }).(pulumi.IntOutput)
}

// A description of the rule group.
func (o RuleGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuleGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// An object that defines the rule group rules.
func (o RuleGroupOutput) RuleGroup() RuleGroupTypePtrOutput {
	return o.ApplyT(func(v *RuleGroup) RuleGroupTypePtrOutput { return v.RuleGroup }).(RuleGroupTypePtrOutput)
}

// The Amazon Resource Name (ARN) of the `RuleGroup` .
func (o RuleGroupOutput) RuleGroupArn() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleGroup) pulumi.StringOutput { return v.RuleGroupArn }).(pulumi.StringOutput)
}

// The unique ID of the `RuleGroup` resource.
func (o RuleGroupOutput) RuleGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleGroup) pulumi.StringOutput { return v.RuleGroupId }).(pulumi.StringOutput)
}

// The descriptive name of the rule group. You can't change the name of a rule group after you create it.
func (o RuleGroupOutput) RuleGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleGroup) pulumi.StringOutput { return v.RuleGroupName }).(pulumi.StringOutput)
}

// An array of key-value pairs to apply to this resource.
//
// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
func (o RuleGroupOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *RuleGroup) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

// Indicates whether the rule group is stateless or stateful. If the rule group is stateless, it contains
// stateless rules. If it is stateful, it contains stateful rules.
func (o RuleGroupOutput) Type() RuleGroupTypeEnumOutput {
	return o.ApplyT(func(v *RuleGroup) RuleGroupTypeEnumOutput { return v.Type }).(RuleGroupTypeEnumOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleGroupInput)(nil)).Elem(), &RuleGroup{})
	pulumi.RegisterOutputType(RuleGroupOutput{})
}
