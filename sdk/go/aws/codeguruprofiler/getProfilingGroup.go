// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codeguruprofiler

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource schema represents the Profiling Group resource in the Amazon CodeGuru Profiler service.
func LookupProfilingGroup(ctx *pulumi.Context, args *LookupProfilingGroupArgs, opts ...pulumi.InvokeOption) (*LookupProfilingGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProfilingGroupResult
	err := ctx.Invoke("aws-native:codeguruprofiler:getProfilingGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProfilingGroupArgs struct {
	// The name of the profiling group.
	ProfilingGroupName string `pulumi:"profilingGroupName"`
}

type LookupProfilingGroupResult struct {
	// The agent permissions attached to this profiling group.
	AgentPermissions *AgentPermissionsProperties `pulumi:"agentPermissions"`
	// Configuration for Notification Channels for Anomaly Detection feature in CodeGuru Profiler which enables customers to detect anomalies in the application profile for those methods that represent the highest proportion of CPU time or latency
	AnomalyDetectionNotificationConfiguration []ProfilingGroupChannel `pulumi:"anomalyDetectionNotificationConfiguration"`
	// The Amazon Resource Name (ARN) of the specified profiling group.
	Arn *string `pulumi:"arn"`
	// The tags associated with a profiling group.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupProfilingGroupOutput(ctx *pulumi.Context, args LookupProfilingGroupOutputArgs, opts ...pulumi.InvokeOption) LookupProfilingGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupProfilingGroupResultOutput, error) {
			args := v.(LookupProfilingGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:codeguruprofiler:getProfilingGroup", args, LookupProfilingGroupResultOutput{}, options).(LookupProfilingGroupResultOutput), nil
		}).(LookupProfilingGroupResultOutput)
}

type LookupProfilingGroupOutputArgs struct {
	// The name of the profiling group.
	ProfilingGroupName pulumi.StringInput `pulumi:"profilingGroupName"`
}

func (LookupProfilingGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfilingGroupArgs)(nil)).Elem()
}

type LookupProfilingGroupResultOutput struct{ *pulumi.OutputState }

func (LookupProfilingGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfilingGroupResult)(nil)).Elem()
}

func (o LookupProfilingGroupResultOutput) ToLookupProfilingGroupResultOutput() LookupProfilingGroupResultOutput {
	return o
}

func (o LookupProfilingGroupResultOutput) ToLookupProfilingGroupResultOutputWithContext(ctx context.Context) LookupProfilingGroupResultOutput {
	return o
}

// The agent permissions attached to this profiling group.
func (o LookupProfilingGroupResultOutput) AgentPermissions() AgentPermissionsPropertiesPtrOutput {
	return o.ApplyT(func(v LookupProfilingGroupResult) *AgentPermissionsProperties { return v.AgentPermissions }).(AgentPermissionsPropertiesPtrOutput)
}

// Configuration for Notification Channels for Anomaly Detection feature in CodeGuru Profiler which enables customers to detect anomalies in the application profile for those methods that represent the highest proportion of CPU time or latency
func (o LookupProfilingGroupResultOutput) AnomalyDetectionNotificationConfiguration() ProfilingGroupChannelArrayOutput {
	return o.ApplyT(func(v LookupProfilingGroupResult) []ProfilingGroupChannel {
		return v.AnomalyDetectionNotificationConfiguration
	}).(ProfilingGroupChannelArrayOutput)
}

// The Amazon Resource Name (ARN) of the specified profiling group.
func (o LookupProfilingGroupResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfilingGroupResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The tags associated with a profiling group.
func (o LookupProfilingGroupResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupProfilingGroupResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProfilingGroupResultOutput{})
}
