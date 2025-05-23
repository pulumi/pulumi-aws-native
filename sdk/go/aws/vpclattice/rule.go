// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpclattice

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a listener rule. Each listener has a default rule for checking connection requests, but you can define additional rules. Each rule consists of a priority, one or more actions, and one or more conditions.
type Rule struct {
	pulumi.CustomResourceState

	// Describes the action for a rule.
	Action RuleActionOutput `pulumi:"action"`
	// The Amazon Resource Name (ARN) of the rule.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ID of the listener.
	AwsId pulumi.StringOutput `pulumi:"awsId"`
	// The ID or ARN of the listener.
	ListenerIdentifier pulumi.StringPtrOutput `pulumi:"listenerIdentifier"`
	// The rule match.
	Match RuleMatchOutput `pulumi:"match"`
	// The name of the rule. The name must be unique within the listener. The valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
	//
	// If you don't specify a name, CloudFormation generates one. However, if you specify a name, and later want to replace the resource, you must specify a new name.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The priority assigned to the rule. Each rule for a specific listener must have a unique priority. The lower the priority number the higher the priority.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// The ID or ARN of the service.
	ServiceIdentifier pulumi.StringPtrOutput `pulumi:"serviceIdentifier"`
	// The tags for the rule.
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewRule registers a new resource with the given unique name, arguments, and options.
func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOption) (*Rule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.Match == nil {
		return nil, errors.New("invalid value for required argument 'Match'")
	}
	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"listenerIdentifier",
		"name",
		"serviceIdentifier",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Rule
	err := ctx.RegisterResource("aws-native:vpclattice:Rule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRule gets an existing Rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleState, opts ...pulumi.ResourceOption) (*Rule, error) {
	var resource Rule
	err := ctx.ReadResource("aws-native:vpclattice:Rule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rule resources.
type ruleState struct {
}

type RuleState struct {
}

func (RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleState)(nil)).Elem()
}

type ruleArgs struct {
	// Describes the action for a rule.
	Action RuleAction `pulumi:"action"`
	// The ID or ARN of the listener.
	ListenerIdentifier *string `pulumi:"listenerIdentifier"`
	// The rule match.
	Match RuleMatch `pulumi:"match"`
	// The name of the rule. The name must be unique within the listener. The valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
	//
	// If you don't specify a name, CloudFormation generates one. However, if you specify a name, and later want to replace the resource, you must specify a new name.
	Name *string `pulumi:"name"`
	// The priority assigned to the rule. Each rule for a specific listener must have a unique priority. The lower the priority number the higher the priority.
	Priority int `pulumi:"priority"`
	// The ID or ARN of the service.
	ServiceIdentifier *string `pulumi:"serviceIdentifier"`
	// The tags for the rule.
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	// Describes the action for a rule.
	Action RuleActionInput
	// The ID or ARN of the listener.
	ListenerIdentifier pulumi.StringPtrInput
	// The rule match.
	Match RuleMatchInput
	// The name of the rule. The name must be unique within the listener. The valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
	//
	// If you don't specify a name, CloudFormation generates one. However, if you specify a name, and later want to replace the resource, you must specify a new name.
	Name pulumi.StringPtrInput
	// The priority assigned to the rule. Each rule for a specific listener must have a unique priority. The lower the priority number the higher the priority.
	Priority pulumi.IntInput
	// The ID or ARN of the service.
	ServiceIdentifier pulumi.StringPtrInput
	// The tags for the rule.
	Tags aws.TagArrayInput
}

func (RuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleArgs)(nil)).Elem()
}

type RuleInput interface {
	pulumi.Input

	ToRuleOutput() RuleOutput
	ToRuleOutputWithContext(ctx context.Context) RuleOutput
}

func (*Rule) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil)).Elem()
}

func (i *Rule) ToRuleOutput() RuleOutput {
	return i.ToRuleOutputWithContext(context.Background())
}

func (i *Rule) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleOutput)
}

type RuleOutput struct{ *pulumi.OutputState }

func (RuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil)).Elem()
}

func (o RuleOutput) ToRuleOutput() RuleOutput {
	return o
}

func (o RuleOutput) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return o
}

// Describes the action for a rule.
func (o RuleOutput) Action() RuleActionOutput {
	return o.ApplyT(func(v *Rule) RuleActionOutput { return v.Action }).(RuleActionOutput)
}

// The Amazon Resource Name (ARN) of the rule.
func (o RuleOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The ID of the listener.
func (o RuleOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// The ID or ARN of the listener.
func (o RuleOutput) ListenerIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.ListenerIdentifier }).(pulumi.StringPtrOutput)
}

// The rule match.
func (o RuleOutput) Match() RuleMatchOutput {
	return o.ApplyT(func(v *Rule) RuleMatchOutput { return v.Match }).(RuleMatchOutput)
}

// The name of the rule. The name must be unique within the listener. The valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
//
// If you don't specify a name, CloudFormation generates one. However, if you specify a name, and later want to replace the resource, you must specify a new name.
func (o RuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The priority assigned to the rule. Each rule for a specific listener must have a unique priority. The lower the priority number the higher the priority.
func (o RuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *Rule) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

// The ID or ARN of the service.
func (o RuleOutput) ServiceIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.ServiceIdentifier }).(pulumi.StringPtrOutput)
}

// The tags for the rule.
func (o RuleOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *Rule) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleInput)(nil)).Elem(), &Rule{})
	pulumi.RegisterOutputType(RuleOutput{})
}
