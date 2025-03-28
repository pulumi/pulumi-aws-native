// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package chatbot

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::Chatbot::MicrosoftTeamsChannelConfiguration.
type MicrosoftTeamsChannelConfiguration struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the configuration
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the configuration
	ConfigurationName pulumi.StringOutput `pulumi:"configurationName"`
	// ARNs of Custom Actions to associate with notifications in the provided chat channel.
	CustomizationResourceArns pulumi.StringArrayOutput `pulumi:"customizationResourceArns"`
	// The list of IAM policy ARNs that are applied as channel guardrails. The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.
	GuardrailPolicies pulumi.StringArrayOutput `pulumi:"guardrailPolicies"`
	// The ARN of the IAM role that defines the permissions for AWS Chatbot
	IamRoleArn pulumi.StringOutput `pulumi:"iamRoleArn"`
	// Specifies the logging level for this configuration:ERROR,INFO or NONE. This property affects the log entries pushed to Amazon CloudWatch logs
	LoggingLevel pulumi.StringPtrOutput `pulumi:"loggingLevel"`
	// ARNs of SNS topics which delivers notifications to AWS Chatbot, for example CloudWatch alarm notifications.
	SnsTopicArns pulumi.StringArrayOutput `pulumi:"snsTopicArns"`
	// The tags to add to the configuration
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// The id of the Microsoft Teams team
	TeamId pulumi.StringOutput `pulumi:"teamId"`
	// The id of the Microsoft Teams channel
	TeamsChannelId pulumi.StringOutput `pulumi:"teamsChannelId"`
	// The name of the Microsoft Teams channel
	TeamsChannelName pulumi.StringPtrOutput `pulumi:"teamsChannelName"`
	// The id of the Microsoft Teams tenant
	TeamsTenantId pulumi.StringOutput `pulumi:"teamsTenantId"`
	// Enables use of a user role requirement in your chat configuration
	UserRoleRequired pulumi.BoolPtrOutput `pulumi:"userRoleRequired"`
}

// NewMicrosoftTeamsChannelConfiguration registers a new resource with the given unique name, arguments, and options.
func NewMicrosoftTeamsChannelConfiguration(ctx *pulumi.Context,
	name string, args *MicrosoftTeamsChannelConfigurationArgs, opts ...pulumi.ResourceOption) (*MicrosoftTeamsChannelConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IamRoleArn == nil {
		return nil, errors.New("invalid value for required argument 'IamRoleArn'")
	}
	if args.TeamId == nil {
		return nil, errors.New("invalid value for required argument 'TeamId'")
	}
	if args.TeamsChannelId == nil {
		return nil, errors.New("invalid value for required argument 'TeamsChannelId'")
	}
	if args.TeamsTenantId == nil {
		return nil, errors.New("invalid value for required argument 'TeamsTenantId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"configurationName",
		"teamId",
		"teamsTenantId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MicrosoftTeamsChannelConfiguration
	err := ctx.RegisterResource("aws-native:chatbot:MicrosoftTeamsChannelConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMicrosoftTeamsChannelConfiguration gets an existing MicrosoftTeamsChannelConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMicrosoftTeamsChannelConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MicrosoftTeamsChannelConfigurationState, opts ...pulumi.ResourceOption) (*MicrosoftTeamsChannelConfiguration, error) {
	var resource MicrosoftTeamsChannelConfiguration
	err := ctx.ReadResource("aws-native:chatbot:MicrosoftTeamsChannelConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MicrosoftTeamsChannelConfiguration resources.
type microsoftTeamsChannelConfigurationState struct {
}

type MicrosoftTeamsChannelConfigurationState struct {
}

func (MicrosoftTeamsChannelConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*microsoftTeamsChannelConfigurationState)(nil)).Elem()
}

type microsoftTeamsChannelConfigurationArgs struct {
	// The name of the configuration
	ConfigurationName *string `pulumi:"configurationName"`
	// ARNs of Custom Actions to associate with notifications in the provided chat channel.
	CustomizationResourceArns []string `pulumi:"customizationResourceArns"`
	// The list of IAM policy ARNs that are applied as channel guardrails. The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.
	GuardrailPolicies []string `pulumi:"guardrailPolicies"`
	// The ARN of the IAM role that defines the permissions for AWS Chatbot
	IamRoleArn string `pulumi:"iamRoleArn"`
	// Specifies the logging level for this configuration:ERROR,INFO or NONE. This property affects the log entries pushed to Amazon CloudWatch logs
	LoggingLevel *string `pulumi:"loggingLevel"`
	// ARNs of SNS topics which delivers notifications to AWS Chatbot, for example CloudWatch alarm notifications.
	SnsTopicArns []string `pulumi:"snsTopicArns"`
	// The tags to add to the configuration
	Tags []aws.Tag `pulumi:"tags"`
	// The id of the Microsoft Teams team
	TeamId string `pulumi:"teamId"`
	// The id of the Microsoft Teams channel
	TeamsChannelId string `pulumi:"teamsChannelId"`
	// The name of the Microsoft Teams channel
	TeamsChannelName *string `pulumi:"teamsChannelName"`
	// The id of the Microsoft Teams tenant
	TeamsTenantId string `pulumi:"teamsTenantId"`
	// Enables use of a user role requirement in your chat configuration
	UserRoleRequired *bool `pulumi:"userRoleRequired"`
}

// The set of arguments for constructing a MicrosoftTeamsChannelConfiguration resource.
type MicrosoftTeamsChannelConfigurationArgs struct {
	// The name of the configuration
	ConfigurationName pulumi.StringPtrInput
	// ARNs of Custom Actions to associate with notifications in the provided chat channel.
	CustomizationResourceArns pulumi.StringArrayInput
	// The list of IAM policy ARNs that are applied as channel guardrails. The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.
	GuardrailPolicies pulumi.StringArrayInput
	// The ARN of the IAM role that defines the permissions for AWS Chatbot
	IamRoleArn pulumi.StringInput
	// Specifies the logging level for this configuration:ERROR,INFO or NONE. This property affects the log entries pushed to Amazon CloudWatch logs
	LoggingLevel pulumi.StringPtrInput
	// ARNs of SNS topics which delivers notifications to AWS Chatbot, for example CloudWatch alarm notifications.
	SnsTopicArns pulumi.StringArrayInput
	// The tags to add to the configuration
	Tags aws.TagArrayInput
	// The id of the Microsoft Teams team
	TeamId pulumi.StringInput
	// The id of the Microsoft Teams channel
	TeamsChannelId pulumi.StringInput
	// The name of the Microsoft Teams channel
	TeamsChannelName pulumi.StringPtrInput
	// The id of the Microsoft Teams tenant
	TeamsTenantId pulumi.StringInput
	// Enables use of a user role requirement in your chat configuration
	UserRoleRequired pulumi.BoolPtrInput
}

func (MicrosoftTeamsChannelConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*microsoftTeamsChannelConfigurationArgs)(nil)).Elem()
}

type MicrosoftTeamsChannelConfigurationInput interface {
	pulumi.Input

	ToMicrosoftTeamsChannelConfigurationOutput() MicrosoftTeamsChannelConfigurationOutput
	ToMicrosoftTeamsChannelConfigurationOutputWithContext(ctx context.Context) MicrosoftTeamsChannelConfigurationOutput
}

func (*MicrosoftTeamsChannelConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**MicrosoftTeamsChannelConfiguration)(nil)).Elem()
}

func (i *MicrosoftTeamsChannelConfiguration) ToMicrosoftTeamsChannelConfigurationOutput() MicrosoftTeamsChannelConfigurationOutput {
	return i.ToMicrosoftTeamsChannelConfigurationOutputWithContext(context.Background())
}

func (i *MicrosoftTeamsChannelConfiguration) ToMicrosoftTeamsChannelConfigurationOutputWithContext(ctx context.Context) MicrosoftTeamsChannelConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MicrosoftTeamsChannelConfigurationOutput)
}

type MicrosoftTeamsChannelConfigurationOutput struct{ *pulumi.OutputState }

func (MicrosoftTeamsChannelConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MicrosoftTeamsChannelConfiguration)(nil)).Elem()
}

func (o MicrosoftTeamsChannelConfigurationOutput) ToMicrosoftTeamsChannelConfigurationOutput() MicrosoftTeamsChannelConfigurationOutput {
	return o
}

func (o MicrosoftTeamsChannelConfigurationOutput) ToMicrosoftTeamsChannelConfigurationOutputWithContext(ctx context.Context) MicrosoftTeamsChannelConfigurationOutput {
	return o
}

// Amazon Resource Name (ARN) of the configuration
func (o MicrosoftTeamsChannelConfigurationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftTeamsChannelConfiguration) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The name of the configuration
func (o MicrosoftTeamsChannelConfigurationOutput) ConfigurationName() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftTeamsChannelConfiguration) pulumi.StringOutput { return v.ConfigurationName }).(pulumi.StringOutput)
}

// ARNs of Custom Actions to associate with notifications in the provided chat channel.
func (o MicrosoftTeamsChannelConfigurationOutput) CustomizationResourceArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MicrosoftTeamsChannelConfiguration) pulumi.StringArrayOutput {
		return v.CustomizationResourceArns
	}).(pulumi.StringArrayOutput)
}

// The list of IAM policy ARNs that are applied as channel guardrails. The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.
func (o MicrosoftTeamsChannelConfigurationOutput) GuardrailPolicies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MicrosoftTeamsChannelConfiguration) pulumi.StringArrayOutput { return v.GuardrailPolicies }).(pulumi.StringArrayOutput)
}

// The ARN of the IAM role that defines the permissions for AWS Chatbot
func (o MicrosoftTeamsChannelConfigurationOutput) IamRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftTeamsChannelConfiguration) pulumi.StringOutput { return v.IamRoleArn }).(pulumi.StringOutput)
}

// Specifies the logging level for this configuration:ERROR,INFO or NONE. This property affects the log entries pushed to Amazon CloudWatch logs
func (o MicrosoftTeamsChannelConfigurationOutput) LoggingLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MicrosoftTeamsChannelConfiguration) pulumi.StringPtrOutput { return v.LoggingLevel }).(pulumi.StringPtrOutput)
}

// ARNs of SNS topics which delivers notifications to AWS Chatbot, for example CloudWatch alarm notifications.
func (o MicrosoftTeamsChannelConfigurationOutput) SnsTopicArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MicrosoftTeamsChannelConfiguration) pulumi.StringArrayOutput { return v.SnsTopicArns }).(pulumi.StringArrayOutput)
}

// The tags to add to the configuration
func (o MicrosoftTeamsChannelConfigurationOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *MicrosoftTeamsChannelConfiguration) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

// The id of the Microsoft Teams team
func (o MicrosoftTeamsChannelConfigurationOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftTeamsChannelConfiguration) pulumi.StringOutput { return v.TeamId }).(pulumi.StringOutput)
}

// The id of the Microsoft Teams channel
func (o MicrosoftTeamsChannelConfigurationOutput) TeamsChannelId() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftTeamsChannelConfiguration) pulumi.StringOutput { return v.TeamsChannelId }).(pulumi.StringOutput)
}

// The name of the Microsoft Teams channel
func (o MicrosoftTeamsChannelConfigurationOutput) TeamsChannelName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MicrosoftTeamsChannelConfiguration) pulumi.StringPtrOutput { return v.TeamsChannelName }).(pulumi.StringPtrOutput)
}

// The id of the Microsoft Teams tenant
func (o MicrosoftTeamsChannelConfigurationOutput) TeamsTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftTeamsChannelConfiguration) pulumi.StringOutput { return v.TeamsTenantId }).(pulumi.StringOutput)
}

// Enables use of a user role requirement in your chat configuration
func (o MicrosoftTeamsChannelConfigurationOutput) UserRoleRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MicrosoftTeamsChannelConfiguration) pulumi.BoolPtrOutput { return v.UserRoleRequired }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MicrosoftTeamsChannelConfigurationInput)(nil)).Elem(), &MicrosoftTeamsChannelConfiguration{})
	pulumi.RegisterOutputType(MicrosoftTeamsChannelConfigurationOutput{})
}
