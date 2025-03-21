// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package supportapp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An AWS Support App resource that creates, updates, lists and deletes Slack channel configurations.
func LookupSlackChannelConfiguration(ctx *pulumi.Context, args *LookupSlackChannelConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupSlackChannelConfigurationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSlackChannelConfigurationResult
	err := ctx.Invoke("aws-native:supportapp:getSlackChannelConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSlackChannelConfigurationArgs struct {
	// The channel ID in Slack, which identifies a channel within a workspace.
	ChannelId string `pulumi:"channelId"`
	// The team ID in Slack, which uniquely identifies a workspace.
	TeamId string `pulumi:"teamId"`
}

type LookupSlackChannelConfigurationResult struct {
	// The channel name in Slack.
	ChannelName *string `pulumi:"channelName"`
	// The Amazon Resource Name (ARN) of an IAM role that grants the AWS Support App access to perform operations for AWS services.
	ChannelRoleArn *string `pulumi:"channelRoleArn"`
	// Whether to notify when a correspondence is added to a case.
	NotifyOnAddCorrespondenceToCase *bool `pulumi:"notifyOnAddCorrespondenceToCase"`
	// The severity level of a support case that a customer wants to get notified for.
	NotifyOnCaseSeverity *SlackChannelConfigurationNotifyOnCaseSeverity `pulumi:"notifyOnCaseSeverity"`
	// Whether to notify when a case is created or reopened.
	NotifyOnCreateOrReopenCase *bool `pulumi:"notifyOnCreateOrReopenCase"`
	// Whether to notify when a case is resolved.
	NotifyOnResolveCase *bool `pulumi:"notifyOnResolveCase"`
}

func LookupSlackChannelConfigurationOutput(ctx *pulumi.Context, args LookupSlackChannelConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupSlackChannelConfigurationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSlackChannelConfigurationResultOutput, error) {
			args := v.(LookupSlackChannelConfigurationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:supportapp:getSlackChannelConfiguration", args, LookupSlackChannelConfigurationResultOutput{}, options).(LookupSlackChannelConfigurationResultOutput), nil
		}).(LookupSlackChannelConfigurationResultOutput)
}

type LookupSlackChannelConfigurationOutputArgs struct {
	// The channel ID in Slack, which identifies a channel within a workspace.
	ChannelId pulumi.StringInput `pulumi:"channelId"`
	// The team ID in Slack, which uniquely identifies a workspace.
	TeamId pulumi.StringInput `pulumi:"teamId"`
}

func (LookupSlackChannelConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSlackChannelConfigurationArgs)(nil)).Elem()
}

type LookupSlackChannelConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupSlackChannelConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSlackChannelConfigurationResult)(nil)).Elem()
}

func (o LookupSlackChannelConfigurationResultOutput) ToLookupSlackChannelConfigurationResultOutput() LookupSlackChannelConfigurationResultOutput {
	return o
}

func (o LookupSlackChannelConfigurationResultOutput) ToLookupSlackChannelConfigurationResultOutputWithContext(ctx context.Context) LookupSlackChannelConfigurationResultOutput {
	return o
}

// The channel name in Slack.
func (o LookupSlackChannelConfigurationResultOutput) ChannelName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSlackChannelConfigurationResult) *string { return v.ChannelName }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of an IAM role that grants the AWS Support App access to perform operations for AWS services.
func (o LookupSlackChannelConfigurationResultOutput) ChannelRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSlackChannelConfigurationResult) *string { return v.ChannelRoleArn }).(pulumi.StringPtrOutput)
}

// Whether to notify when a correspondence is added to a case.
func (o LookupSlackChannelConfigurationResultOutput) NotifyOnAddCorrespondenceToCase() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSlackChannelConfigurationResult) *bool { return v.NotifyOnAddCorrespondenceToCase }).(pulumi.BoolPtrOutput)
}

// The severity level of a support case that a customer wants to get notified for.
func (o LookupSlackChannelConfigurationResultOutput) NotifyOnCaseSeverity() SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput {
	return o.ApplyT(func(v LookupSlackChannelConfigurationResult) *SlackChannelConfigurationNotifyOnCaseSeverity {
		return v.NotifyOnCaseSeverity
	}).(SlackChannelConfigurationNotifyOnCaseSeverityPtrOutput)
}

// Whether to notify when a case is created or reopened.
func (o LookupSlackChannelConfigurationResultOutput) NotifyOnCreateOrReopenCase() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSlackChannelConfigurationResult) *bool { return v.NotifyOnCreateOrReopenCase }).(pulumi.BoolPtrOutput)
}

// Whether to notify when a case is resolved.
func (o LookupSlackChannelConfigurationResultOutput) NotifyOnResolveCase() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSlackChannelConfigurationResult) *bool { return v.NotifyOnResolveCase }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSlackChannelConfigurationResultOutput{})
}
