// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::AutoScaling::ScheduledAction resource specifies an Amazon EC2 Auto Scaling scheduled action so that the Auto Scaling group can change the number of instances available for your application in response to predictable load changes.
func LookupScheduledAction(ctx *pulumi.Context, args *LookupScheduledActionArgs, opts ...pulumi.InvokeOption) (*LookupScheduledActionResult, error) {
	var rv LookupScheduledActionResult
	err := ctx.Invoke("aws-native:autoscaling:getScheduledAction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduledActionArgs struct {
	// The name of the Auto Scaling group.
	AutoScalingGroupName string `pulumi:"autoScalingGroupName"`
	// Auto-generated unique identifier
	ScheduledActionName string `pulumi:"scheduledActionName"`
}

type LookupScheduledActionResult struct {
	// The desired capacity is the initial capacity of the Auto Scaling group after the scheduled action runs and the capacity it attempts to maintain.
	DesiredCapacity *int `pulumi:"desiredCapacity"`
	// The latest scheduled start time to return. If scheduled action names are provided, this parameter is ignored.
	EndTime *string `pulumi:"endTime"`
	// The minimum size of the Auto Scaling group.
	MaxSize *int `pulumi:"maxSize"`
	// The minimum size of the Auto Scaling group.
	MinSize *int `pulumi:"minSize"`
	// The recurring schedule for the action, in Unix cron syntax format. When StartTime and EndTime are specified with Recurrence , they form the boundaries of when the recurring action starts and stops.
	Recurrence *string `pulumi:"recurrence"`
	// Auto-generated unique identifier
	ScheduledActionName *string `pulumi:"scheduledActionName"`
	// The earliest scheduled start time to return. If scheduled action names are provided, this parameter is ignored.
	StartTime *string `pulumi:"startTime"`
	// The time zone for the cron expression.
	TimeZone *string `pulumi:"timeZone"`
}

func LookupScheduledActionOutput(ctx *pulumi.Context, args LookupScheduledActionOutputArgs, opts ...pulumi.InvokeOption) LookupScheduledActionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScheduledActionResult, error) {
			args := v.(LookupScheduledActionArgs)
			r, err := LookupScheduledAction(ctx, &args, opts...)
			var s LookupScheduledActionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScheduledActionResultOutput)
}

type LookupScheduledActionOutputArgs struct {
	// The name of the Auto Scaling group.
	AutoScalingGroupName pulumi.StringInput `pulumi:"autoScalingGroupName"`
	// Auto-generated unique identifier
	ScheduledActionName pulumi.StringInput `pulumi:"scheduledActionName"`
}

func (LookupScheduledActionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledActionArgs)(nil)).Elem()
}

type LookupScheduledActionResultOutput struct{ *pulumi.OutputState }

func (LookupScheduledActionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledActionResult)(nil)).Elem()
}

func (o LookupScheduledActionResultOutput) ToLookupScheduledActionResultOutput() LookupScheduledActionResultOutput {
	return o
}

func (o LookupScheduledActionResultOutput) ToLookupScheduledActionResultOutputWithContext(ctx context.Context) LookupScheduledActionResultOutput {
	return o
}

// The desired capacity is the initial capacity of the Auto Scaling group after the scheduled action runs and the capacity it attempts to maintain.
func (o LookupScheduledActionResultOutput) DesiredCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *int { return v.DesiredCapacity }).(pulumi.IntPtrOutput)
}

// The latest scheduled start time to return. If scheduled action names are provided, this parameter is ignored.
func (o LookupScheduledActionResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

// The minimum size of the Auto Scaling group.
func (o LookupScheduledActionResultOutput) MaxSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *int { return v.MaxSize }).(pulumi.IntPtrOutput)
}

// The minimum size of the Auto Scaling group.
func (o LookupScheduledActionResultOutput) MinSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *int { return v.MinSize }).(pulumi.IntPtrOutput)
}

// The recurring schedule for the action, in Unix cron syntax format. When StartTime and EndTime are specified with Recurrence , they form the boundaries of when the recurring action starts and stops.
func (o LookupScheduledActionResultOutput) Recurrence() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *string { return v.Recurrence }).(pulumi.StringPtrOutput)
}

// Auto-generated unique identifier
func (o LookupScheduledActionResultOutput) ScheduledActionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *string { return v.ScheduledActionName }).(pulumi.StringPtrOutput)
}

// The earliest scheduled start time to return. If scheduled action names are provided, this parameter is ignored.
func (o LookupScheduledActionResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

// The time zone for the cron expression.
func (o LookupScheduledActionResultOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScheduledActionResultOutput{})
}