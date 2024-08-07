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

// Resource type definition for AWS::NetworkFirewall::FirewallPolicy
type FirewallPolicy struct {
	pulumi.CustomResourceState

	// A description of the firewall policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The traffic filtering behavior of a firewall policy, defined in a collection of stateless and stateful rule groups and other settings.
	FirewallPolicy FirewallPolicyTypeOutput `pulumi:"firewallPolicy"`
	// The Amazon Resource Name (ARN) of the `FirewallPolicy` .
	FirewallPolicyArn pulumi.StringOutput `pulumi:"firewallPolicyArn"`
	// The unique ID of the `FirewallPolicy` resource.
	FirewallPolicyId pulumi.StringOutput `pulumi:"firewallPolicyId"`
	// The descriptive name of the firewall policy. You can't change the name of a firewall policy after you create it.
	FirewallPolicyName pulumi.StringOutput `pulumi:"firewallPolicyName"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewFirewallPolicy registers a new resource with the given unique name, arguments, and options.
func NewFirewallPolicy(ctx *pulumi.Context,
	name string, args *FirewallPolicyArgs, opts ...pulumi.ResourceOption) (*FirewallPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FirewallPolicy == nil {
		return nil, errors.New("invalid value for required argument 'FirewallPolicy'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"firewallPolicyName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallPolicy
	err := ctx.RegisterResource("aws-native:networkfirewall:FirewallPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallPolicy gets an existing FirewallPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicyState, opts ...pulumi.ResourceOption) (*FirewallPolicy, error) {
	var resource FirewallPolicy
	err := ctx.ReadResource("aws-native:networkfirewall:FirewallPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallPolicy resources.
type firewallPolicyState struct {
}

type FirewallPolicyState struct {
}

func (FirewallPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyState)(nil)).Elem()
}

type firewallPolicyArgs struct {
	// A description of the firewall policy.
	Description *string `pulumi:"description"`
	// The traffic filtering behavior of a firewall policy, defined in a collection of stateless and stateful rule groups and other settings.
	FirewallPolicy FirewallPolicyType `pulumi:"firewallPolicy"`
	// The descriptive name of the firewall policy. You can't change the name of a firewall policy after you create it.
	FirewallPolicyName *string `pulumi:"firewallPolicyName"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a FirewallPolicy resource.
type FirewallPolicyArgs struct {
	// A description of the firewall policy.
	Description pulumi.StringPtrInput
	// The traffic filtering behavior of a firewall policy, defined in a collection of stateless and stateful rule groups and other settings.
	FirewallPolicy FirewallPolicyTypeInput
	// The descriptive name of the firewall policy. You can't change the name of a firewall policy after you create it.
	FirewallPolicyName pulumi.StringPtrInput
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags aws.TagArrayInput
}

func (FirewallPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyArgs)(nil)).Elem()
}

type FirewallPolicyInput interface {
	pulumi.Input

	ToFirewallPolicyOutput() FirewallPolicyOutput
	ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput
}

func (*FirewallPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicy)(nil)).Elem()
}

func (i *FirewallPolicy) ToFirewallPolicyOutput() FirewallPolicyOutput {
	return i.ToFirewallPolicyOutputWithContext(context.Background())
}

func (i *FirewallPolicy) ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyOutput)
}

type FirewallPolicyOutput struct{ *pulumi.OutputState }

func (FirewallPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicy)(nil)).Elem()
}

func (o FirewallPolicyOutput) ToFirewallPolicyOutput() FirewallPolicyOutput {
	return o
}

func (o FirewallPolicyOutput) ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput {
	return o
}

// A description of the firewall policy.
func (o FirewallPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The traffic filtering behavior of a firewall policy, defined in a collection of stateless and stateful rule groups and other settings.
func (o FirewallPolicyOutput) FirewallPolicy() FirewallPolicyTypeOutput {
	return o.ApplyT(func(v *FirewallPolicy) FirewallPolicyTypeOutput { return v.FirewallPolicy }).(FirewallPolicyTypeOutput)
}

// The Amazon Resource Name (ARN) of the `FirewallPolicy` .
func (o FirewallPolicyOutput) FirewallPolicyArn() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringOutput { return v.FirewallPolicyArn }).(pulumi.StringOutput)
}

// The unique ID of the `FirewallPolicy` resource.
func (o FirewallPolicyOutput) FirewallPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringOutput { return v.FirewallPolicyId }).(pulumi.StringOutput)
}

// The descriptive name of the firewall policy. You can't change the name of a firewall policy after you create it.
func (o FirewallPolicyOutput) FirewallPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringOutput { return v.FirewallPolicyName }).(pulumi.StringOutput)
}

// An array of key-value pairs to apply to this resource.
//
// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
func (o FirewallPolicyOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *FirewallPolicy) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallPolicyInput)(nil)).Elem(), &FirewallPolicy{})
	pulumi.RegisterOutputType(FirewallPolicyOutput{})
}
