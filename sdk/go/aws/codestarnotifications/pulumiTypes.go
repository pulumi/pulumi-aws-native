// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codestarnotifications

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type NotificationRuleTarget struct {
	// The Amazon Resource Name (ARN) of the  topic or  client.
	TargetAddress string `pulumi:"targetAddress"`
	// The target type. Can be an Amazon Simple Notification Service topic or  client.
	//
	// - Amazon Simple Notification Service topics are specified as `SNS` .
	// - clients are specified as `AWSChatbotSlack` .
	// - clients for Microsoft Teams are specified as `AWSChatbotMicrosoftTeams` .
	TargetType string `pulumi:"targetType"`
}

// NotificationRuleTargetInput is an input type that accepts NotificationRuleTargetArgs and NotificationRuleTargetOutput values.
// You can construct a concrete instance of `NotificationRuleTargetInput` via:
//
//	NotificationRuleTargetArgs{...}
type NotificationRuleTargetInput interface {
	pulumi.Input

	ToNotificationRuleTargetOutput() NotificationRuleTargetOutput
	ToNotificationRuleTargetOutputWithContext(context.Context) NotificationRuleTargetOutput
}

type NotificationRuleTargetArgs struct {
	// The Amazon Resource Name (ARN) of the  topic or  client.
	TargetAddress pulumi.StringInput `pulumi:"targetAddress"`
	// The target type. Can be an Amazon Simple Notification Service topic or  client.
	//
	// - Amazon Simple Notification Service topics are specified as `SNS` .
	// - clients are specified as `AWSChatbotSlack` .
	// - clients for Microsoft Teams are specified as `AWSChatbotMicrosoftTeams` .
	TargetType pulumi.StringInput `pulumi:"targetType"`
}

func (NotificationRuleTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationRuleTarget)(nil)).Elem()
}

func (i NotificationRuleTargetArgs) ToNotificationRuleTargetOutput() NotificationRuleTargetOutput {
	return i.ToNotificationRuleTargetOutputWithContext(context.Background())
}

func (i NotificationRuleTargetArgs) ToNotificationRuleTargetOutputWithContext(ctx context.Context) NotificationRuleTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationRuleTargetOutput)
}

// NotificationRuleTargetArrayInput is an input type that accepts NotificationRuleTargetArray and NotificationRuleTargetArrayOutput values.
// You can construct a concrete instance of `NotificationRuleTargetArrayInput` via:
//
//	NotificationRuleTargetArray{ NotificationRuleTargetArgs{...} }
type NotificationRuleTargetArrayInput interface {
	pulumi.Input

	ToNotificationRuleTargetArrayOutput() NotificationRuleTargetArrayOutput
	ToNotificationRuleTargetArrayOutputWithContext(context.Context) NotificationRuleTargetArrayOutput
}

type NotificationRuleTargetArray []NotificationRuleTargetInput

func (NotificationRuleTargetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NotificationRuleTarget)(nil)).Elem()
}

func (i NotificationRuleTargetArray) ToNotificationRuleTargetArrayOutput() NotificationRuleTargetArrayOutput {
	return i.ToNotificationRuleTargetArrayOutputWithContext(context.Background())
}

func (i NotificationRuleTargetArray) ToNotificationRuleTargetArrayOutputWithContext(ctx context.Context) NotificationRuleTargetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationRuleTargetArrayOutput)
}

type NotificationRuleTargetOutput struct{ *pulumi.OutputState }

func (NotificationRuleTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationRuleTarget)(nil)).Elem()
}

func (o NotificationRuleTargetOutput) ToNotificationRuleTargetOutput() NotificationRuleTargetOutput {
	return o
}

func (o NotificationRuleTargetOutput) ToNotificationRuleTargetOutputWithContext(ctx context.Context) NotificationRuleTargetOutput {
	return o
}

// The Amazon Resource Name (ARN) of the  topic or  client.
func (o NotificationRuleTargetOutput) TargetAddress() pulumi.StringOutput {
	return o.ApplyT(func(v NotificationRuleTarget) string { return v.TargetAddress }).(pulumi.StringOutput)
}

// The target type. Can be an Amazon Simple Notification Service topic or  client.
//
// - Amazon Simple Notification Service topics are specified as `SNS` .
// - clients are specified as `AWSChatbotSlack` .
// - clients for Microsoft Teams are specified as `AWSChatbotMicrosoftTeams` .
func (o NotificationRuleTargetOutput) TargetType() pulumi.StringOutput {
	return o.ApplyT(func(v NotificationRuleTarget) string { return v.TargetType }).(pulumi.StringOutput)
}

type NotificationRuleTargetArrayOutput struct{ *pulumi.OutputState }

func (NotificationRuleTargetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NotificationRuleTarget)(nil)).Elem()
}

func (o NotificationRuleTargetArrayOutput) ToNotificationRuleTargetArrayOutput() NotificationRuleTargetArrayOutput {
	return o
}

func (o NotificationRuleTargetArrayOutput) ToNotificationRuleTargetArrayOutputWithContext(ctx context.Context) NotificationRuleTargetArrayOutput {
	return o
}

func (o NotificationRuleTargetArrayOutput) Index(i pulumi.IntInput) NotificationRuleTargetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NotificationRuleTarget {
		return vs[0].([]NotificationRuleTarget)[vs[1].(int)]
	}).(NotificationRuleTargetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationRuleTargetInput)(nil)).Elem(), NotificationRuleTargetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationRuleTargetArrayInput)(nil)).Elem(), NotificationRuleTargetArray{})
	pulumi.RegisterOutputType(NotificationRuleTargetOutput{})
	pulumi.RegisterOutputType(NotificationRuleTargetArrayOutput{})
}
