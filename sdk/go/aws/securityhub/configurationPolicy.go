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

// The AWS::SecurityHub::ConfigurationPolicy resource represents the Central Configuration Policy in your account.
type ConfigurationPolicy struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the configuration policy.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The universally unique identifier (UUID) of the configuration policy.
	AwsId pulumi.StringOutput `pulumi:"awsId"`
	// An object that defines how AWS Security Hub is configured. It includes whether Security Hub is enabled or disabled, a list of enabled security standards, a list of enabled or disabled security controls, and a list of custom parameter values for specified controls. If you provide a list of security controls that are enabled in the configuration policy, Security Hub disables all other controls (including newly released controls). If you provide a list of security controls that are disabled in the configuration policy, Security Hub enables all other controls (including newly released controls).
	ConfigurationPolicy ConfigurationPolicyPolicyOutput `pulumi:"configurationPolicy"`
	// The date and time, in UTC and ISO 8601 format.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The description of the configuration policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the configuration policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// Indicates whether the service that the configuration policy applies to is enabled in the policy.
	ServiceEnabled pulumi.BoolOutput `pulumi:"serviceEnabled"`
	// User-defined tags associated with a configuration policy. For more information, see [Tagging AWS Security Hub resources](https://docs.aws.amazon.com/securityhub/latest/userguide/tagging-resources.html) in the *Security Hub user guide* .
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The date and time, in UTC and ISO 8601 format.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewConfigurationPolicy registers a new resource with the given unique name, arguments, and options.
func NewConfigurationPolicy(ctx *pulumi.Context,
	name string, args *ConfigurationPolicyArgs, opts ...pulumi.ResourceOption) (*ConfigurationPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigurationPolicy == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationPolicy'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ConfigurationPolicy
	err := ctx.RegisterResource("aws-native:securityhub:ConfigurationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationPolicy gets an existing ConfigurationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationPolicyState, opts ...pulumi.ResourceOption) (*ConfigurationPolicy, error) {
	var resource ConfigurationPolicy
	err := ctx.ReadResource("aws-native:securityhub:ConfigurationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationPolicy resources.
type configurationPolicyState struct {
}

type ConfigurationPolicyState struct {
}

func (ConfigurationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationPolicyState)(nil)).Elem()
}

type configurationPolicyArgs struct {
	// An object that defines how AWS Security Hub is configured. It includes whether Security Hub is enabled or disabled, a list of enabled security standards, a list of enabled or disabled security controls, and a list of custom parameter values for specified controls. If you provide a list of security controls that are enabled in the configuration policy, Security Hub disables all other controls (including newly released controls). If you provide a list of security controls that are disabled in the configuration policy, Security Hub enables all other controls (including newly released controls).
	ConfigurationPolicy ConfigurationPolicyPolicy `pulumi:"configurationPolicy"`
	// The description of the configuration policy.
	Description *string `pulumi:"description"`
	// The name of the configuration policy.
	Name *string `pulumi:"name"`
	// User-defined tags associated with a configuration policy. For more information, see [Tagging AWS Security Hub resources](https://docs.aws.amazon.com/securityhub/latest/userguide/tagging-resources.html) in the *Security Hub user guide* .
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ConfigurationPolicy resource.
type ConfigurationPolicyArgs struct {
	// An object that defines how AWS Security Hub is configured. It includes whether Security Hub is enabled or disabled, a list of enabled security standards, a list of enabled or disabled security controls, and a list of custom parameter values for specified controls. If you provide a list of security controls that are enabled in the configuration policy, Security Hub disables all other controls (including newly released controls). If you provide a list of security controls that are disabled in the configuration policy, Security Hub enables all other controls (including newly released controls).
	ConfigurationPolicy ConfigurationPolicyPolicyInput
	// The description of the configuration policy.
	Description pulumi.StringPtrInput
	// The name of the configuration policy.
	Name pulumi.StringPtrInput
	// User-defined tags associated with a configuration policy. For more information, see [Tagging AWS Security Hub resources](https://docs.aws.amazon.com/securityhub/latest/userguide/tagging-resources.html) in the *Security Hub user guide* .
	Tags pulumi.StringMapInput
}

func (ConfigurationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationPolicyArgs)(nil)).Elem()
}

type ConfigurationPolicyInput interface {
	pulumi.Input

	ToConfigurationPolicyOutput() ConfigurationPolicyOutput
	ToConfigurationPolicyOutputWithContext(ctx context.Context) ConfigurationPolicyOutput
}

func (*ConfigurationPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationPolicy)(nil)).Elem()
}

func (i *ConfigurationPolicy) ToConfigurationPolicyOutput() ConfigurationPolicyOutput {
	return i.ToConfigurationPolicyOutputWithContext(context.Background())
}

func (i *ConfigurationPolicy) ToConfigurationPolicyOutputWithContext(ctx context.Context) ConfigurationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationPolicyOutput)
}

type ConfigurationPolicyOutput struct{ *pulumi.OutputState }

func (ConfigurationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationPolicy)(nil)).Elem()
}

func (o ConfigurationPolicyOutput) ToConfigurationPolicyOutput() ConfigurationPolicyOutput {
	return o
}

func (o ConfigurationPolicyOutput) ToConfigurationPolicyOutputWithContext(ctx context.Context) ConfigurationPolicyOutput {
	return o
}

// The Amazon Resource Name (ARN) of the configuration policy.
func (o ConfigurationPolicyOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationPolicy) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The universally unique identifier (UUID) of the configuration policy.
func (o ConfigurationPolicyOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationPolicy) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// An object that defines how AWS Security Hub is configured. It includes whether Security Hub is enabled or disabled, a list of enabled security standards, a list of enabled or disabled security controls, and a list of custom parameter values for specified controls. If you provide a list of security controls that are enabled in the configuration policy, Security Hub disables all other controls (including newly released controls). If you provide a list of security controls that are disabled in the configuration policy, Security Hub enables all other controls (including newly released controls).
func (o ConfigurationPolicyOutput) ConfigurationPolicy() ConfigurationPolicyPolicyOutput {
	return o.ApplyT(func(v *ConfigurationPolicy) ConfigurationPolicyPolicyOutput { return v.ConfigurationPolicy }).(ConfigurationPolicyPolicyOutput)
}

// The date and time, in UTC and ISO 8601 format.
func (o ConfigurationPolicyOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationPolicy) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The description of the configuration policy.
func (o ConfigurationPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the configuration policy.
func (o ConfigurationPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Indicates whether the service that the configuration policy applies to is enabled in the policy.
func (o ConfigurationPolicyOutput) ServiceEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ConfigurationPolicy) pulumi.BoolOutput { return v.ServiceEnabled }).(pulumi.BoolOutput)
}

// User-defined tags associated with a configuration policy. For more information, see [Tagging AWS Security Hub resources](https://docs.aws.amazon.com/securityhub/latest/userguide/tagging-resources.html) in the *Security Hub user guide* .
func (o ConfigurationPolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConfigurationPolicy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The date and time, in UTC and ISO 8601 format.
func (o ConfigurationPolicyOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationPolicy) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationPolicyInput)(nil)).Elem(), &ConfigurationPolicy{})
	pulumi.RegisterOutputType(ConfigurationPolicyOutput{})
}