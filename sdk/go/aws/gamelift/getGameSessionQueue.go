// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::GameLift::GameSessionQueue
func LookupGameSessionQueue(ctx *pulumi.Context, args *LookupGameSessionQueueArgs, opts ...pulumi.InvokeOption) (*LookupGameSessionQueueResult, error) {
	var rv LookupGameSessionQueueResult
	err := ctx.Invoke("aws-native:gamelift:getGameSessionQueue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGameSessionQueueArgs struct {
	Id string `pulumi:"id"`
}

type LookupGameSessionQueueResult struct {
	Arn                   *string                                `pulumi:"arn"`
	CustomEventData       *string                                `pulumi:"customEventData"`
	Destinations          []GameSessionQueueDestination          `pulumi:"destinations"`
	FilterConfiguration   *GameSessionQueueFilterConfiguration   `pulumi:"filterConfiguration"`
	Id                    *string                                `pulumi:"id"`
	NotificationTarget    *string                                `pulumi:"notificationTarget"`
	PlayerLatencyPolicies []GameSessionQueuePlayerLatencyPolicy  `pulumi:"playerLatencyPolicies"`
	PriorityConfiguration *GameSessionQueuePriorityConfiguration `pulumi:"priorityConfiguration"`
	Tags                  []GameSessionQueueTag                  `pulumi:"tags"`
	TimeoutInSeconds      *int                                   `pulumi:"timeoutInSeconds"`
}

func LookupGameSessionQueueOutput(ctx *pulumi.Context, args LookupGameSessionQueueOutputArgs, opts ...pulumi.InvokeOption) LookupGameSessionQueueResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGameSessionQueueResult, error) {
			args := v.(LookupGameSessionQueueArgs)
			r, err := LookupGameSessionQueue(ctx, &args, opts...)
			var s LookupGameSessionQueueResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGameSessionQueueResultOutput)
}

type LookupGameSessionQueueOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupGameSessionQueueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGameSessionQueueArgs)(nil)).Elem()
}

type LookupGameSessionQueueResultOutput struct{ *pulumi.OutputState }

func (LookupGameSessionQueueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGameSessionQueueResult)(nil)).Elem()
}

func (o LookupGameSessionQueueResultOutput) ToLookupGameSessionQueueResultOutput() LookupGameSessionQueueResultOutput {
	return o
}

func (o LookupGameSessionQueueResultOutput) ToLookupGameSessionQueueResultOutputWithContext(ctx context.Context) LookupGameSessionQueueResultOutput {
	return o
}

func (o LookupGameSessionQueueResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGameSessionQueueResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupGameSessionQueueResultOutput) CustomEventData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGameSessionQueueResult) *string { return v.CustomEventData }).(pulumi.StringPtrOutput)
}

func (o LookupGameSessionQueueResultOutput) Destinations() GameSessionQueueDestinationArrayOutput {
	return o.ApplyT(func(v LookupGameSessionQueueResult) []GameSessionQueueDestination { return v.Destinations }).(GameSessionQueueDestinationArrayOutput)
}

func (o LookupGameSessionQueueResultOutput) FilterConfiguration() GameSessionQueueFilterConfigurationPtrOutput {
	return o.ApplyT(func(v LookupGameSessionQueueResult) *GameSessionQueueFilterConfiguration {
		return v.FilterConfiguration
	}).(GameSessionQueueFilterConfigurationPtrOutput)
}

func (o LookupGameSessionQueueResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGameSessionQueueResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupGameSessionQueueResultOutput) NotificationTarget() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGameSessionQueueResult) *string { return v.NotificationTarget }).(pulumi.StringPtrOutput)
}

func (o LookupGameSessionQueueResultOutput) PlayerLatencyPolicies() GameSessionQueuePlayerLatencyPolicyArrayOutput {
	return o.ApplyT(func(v LookupGameSessionQueueResult) []GameSessionQueuePlayerLatencyPolicy {
		return v.PlayerLatencyPolicies
	}).(GameSessionQueuePlayerLatencyPolicyArrayOutput)
}

func (o LookupGameSessionQueueResultOutput) PriorityConfiguration() GameSessionQueuePriorityConfigurationPtrOutput {
	return o.ApplyT(func(v LookupGameSessionQueueResult) *GameSessionQueuePriorityConfiguration {
		return v.PriorityConfiguration
	}).(GameSessionQueuePriorityConfigurationPtrOutput)
}

func (o LookupGameSessionQueueResultOutput) Tags() GameSessionQueueTagArrayOutput {
	return o.ApplyT(func(v LookupGameSessionQueueResult) []GameSessionQueueTag { return v.Tags }).(GameSessionQueueTagArrayOutput)
}

func (o LookupGameSessionQueueResultOutput) TimeoutInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupGameSessionQueueResult) *int { return v.TimeoutInSeconds }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGameSessionQueueResultOutput{})
}