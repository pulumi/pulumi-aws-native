// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotevents

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::IoTEvents::AlarmModel resource creates a alarm model. AWS IoT Events alarms help you monitor your data for changes. The data can be metrics that you measure for your equipment and processes. You can create alarms that send notifications when a threshold is breached. Alarms help you detect issues, streamline maintenance, and optimize performance of your equipment and processes.
//
// Alarms are instances of alarm models. The alarm model specifies what to detect, when to send notifications, who gets notified, and more. You can also specify one or more supported actions that occur when the alarm state changes. AWS IoT Events routes input attributes derived from your data to the appropriate alarms. If the data that you're monitoring is outside the specified range, the alarm is invoked. You can also acknowledge the alarms or set them to the snooze mode.
func LookupAlarmModel(ctx *pulumi.Context, args *LookupAlarmModelArgs, opts ...pulumi.InvokeOption) (*LookupAlarmModelResult, error) {
	var rv LookupAlarmModelResult
	err := ctx.Invoke("aws-native:iotevents:getAlarmModel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAlarmModelArgs struct {
	// The name of the alarm model.
	AlarmModelName string `pulumi:"alarmModelName"`
}

type LookupAlarmModelResult struct {
	AlarmCapabilities *AlarmModelAlarmCapabilities `pulumi:"alarmCapabilities"`
	AlarmEventActions *AlarmModelAlarmEventActions `pulumi:"alarmEventActions"`
	// A brief description of the alarm model.
	AlarmModelDescription *string              `pulumi:"alarmModelDescription"`
	AlarmRule             *AlarmModelAlarmRule `pulumi:"alarmRule"`
	// The ARN of the role that grants permission to AWS IoT Events to perform its operations.
	RoleArn *string `pulumi:"roleArn"`
	// A non-negative integer that reflects the severity level of the alarm.
	Severity *int `pulumi:"severity"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
	Tags []AlarmModelTag `pulumi:"tags"`
}

func LookupAlarmModelOutput(ctx *pulumi.Context, args LookupAlarmModelOutputArgs, opts ...pulumi.InvokeOption) LookupAlarmModelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAlarmModelResult, error) {
			args := v.(LookupAlarmModelArgs)
			r, err := LookupAlarmModel(ctx, &args, opts...)
			var s LookupAlarmModelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAlarmModelResultOutput)
}

type LookupAlarmModelOutputArgs struct {
	// The name of the alarm model.
	AlarmModelName pulumi.StringInput `pulumi:"alarmModelName"`
}

func (LookupAlarmModelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlarmModelArgs)(nil)).Elem()
}

type LookupAlarmModelResultOutput struct{ *pulumi.OutputState }

func (LookupAlarmModelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlarmModelResult)(nil)).Elem()
}

func (o LookupAlarmModelResultOutput) ToLookupAlarmModelResultOutput() LookupAlarmModelResultOutput {
	return o
}

func (o LookupAlarmModelResultOutput) ToLookupAlarmModelResultOutputWithContext(ctx context.Context) LookupAlarmModelResultOutput {
	return o
}

func (o LookupAlarmModelResultOutput) AlarmCapabilities() AlarmModelAlarmCapabilitiesPtrOutput {
	return o.ApplyT(func(v LookupAlarmModelResult) *AlarmModelAlarmCapabilities { return v.AlarmCapabilities }).(AlarmModelAlarmCapabilitiesPtrOutput)
}

func (o LookupAlarmModelResultOutput) AlarmEventActions() AlarmModelAlarmEventActionsPtrOutput {
	return o.ApplyT(func(v LookupAlarmModelResult) *AlarmModelAlarmEventActions { return v.AlarmEventActions }).(AlarmModelAlarmEventActionsPtrOutput)
}

// A brief description of the alarm model.
func (o LookupAlarmModelResultOutput) AlarmModelDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmModelResult) *string { return v.AlarmModelDescription }).(pulumi.StringPtrOutput)
}

func (o LookupAlarmModelResultOutput) AlarmRule() AlarmModelAlarmRulePtrOutput {
	return o.ApplyT(func(v LookupAlarmModelResult) *AlarmModelAlarmRule { return v.AlarmRule }).(AlarmModelAlarmRulePtrOutput)
}

// The ARN of the role that grants permission to AWS IoT Events to perform its operations.
func (o LookupAlarmModelResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmModelResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

// A non-negative integer that reflects the severity level of the alarm.
func (o LookupAlarmModelResultOutput) Severity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAlarmModelResult) *int { return v.Severity }).(pulumi.IntPtrOutput)
}

// An array of key-value pairs to apply to this resource.
//
// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
func (o LookupAlarmModelResultOutput) Tags() AlarmModelTagArrayOutput {
	return o.ApplyT(func(v LookupAlarmModelResult) []AlarmModelTag { return v.Tags }).(AlarmModelTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAlarmModelResultOutput{})
}