// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package evidently

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Evidently::Experiment.
func LookupExperiment(ctx *pulumi.Context, args *LookupExperimentArgs, opts ...pulumi.InvokeOption) (*LookupExperimentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupExperimentResult
	err := ctx.Invoke("aws-native:evidently:getExperiment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExperimentArgs struct {
	// The ARN of the experiment. For example, `arn:aws:evidently:us-west-2:0123455678912:project/myProject/experiment/myExperiment`
	Arn string `pulumi:"arn"`
}

type LookupExperimentResult struct {
	// The ARN of the experiment. For example, `arn:aws:evidently:us-west-2:0123455678912:project/myProject/experiment/myExperiment`
	Arn *string `pulumi:"arn"`
	// An optional description of the experiment.
	Description *string `pulumi:"description"`
	// An array of structures that defines the metrics used for the experiment, and whether a higher or lower value for each metric is the goal. You can use up to three metrics in an experiment.
	MetricGoals []ExperimentMetricGoalObject `pulumi:"metricGoals"`
	// A structure that contains the configuration of which variation to use as the "control" version. The "control" version is used for comparison with other variations. This structure also specifies how much experiment traffic is allocated to each variation.
	OnlineAbConfig *ExperimentOnlineAbConfigObject `pulumi:"onlineAbConfig"`
	// When Evidently assigns a particular user session to an experiment, it must use a randomization ID to determine which variation the user session is served. This randomization ID is a combination of the entity ID and `randomizationSalt` . If you omit `randomizationSalt` , Evidently uses the experiment name as the `randomizationSalt` .
	RandomizationSalt *string `pulumi:"randomizationSalt"`
	// Set this to `true` to remove the segment that is associated with this experiment. You can't use this parameter if the experiment is currently running.
	RemoveSegment *bool `pulumi:"removeSegment"`
	// Start Experiment. Default is False
	RunningStatus *ExperimentRunningStatusObject `pulumi:"runningStatus"`
	// The portion of the available audience that you want to allocate to this experiment, in thousandths of a percent. The available audience is the total audience minus the audience that you have allocated to overrides or current launches of this feature.
	//
	// This is represented in thousandths of a percent. For example, specify 10,000 to allocate 10% of the available audience.
	SamplingRate *int `pulumi:"samplingRate"`
	// Specifies an audience *segment* to use in the experiment. When a segment is used in an experiment, only user sessions that match the segment pattern are used in the experiment.
	//
	// For more information, see [Segment rule pattern syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-segments.html#CloudWatch-Evidently-segments-syntax) .
	Segment *string `pulumi:"segment"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
	// An array of structures that describe the configuration of each feature variation used in the experiment.
	Treatments []ExperimentTreatmentObject `pulumi:"treatments"`
}

func LookupExperimentOutput(ctx *pulumi.Context, args LookupExperimentOutputArgs, opts ...pulumi.InvokeOption) LookupExperimentResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupExperimentResultOutput, error) {
			args := v.(LookupExperimentArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:evidently:getExperiment", args, LookupExperimentResultOutput{}, options).(LookupExperimentResultOutput), nil
		}).(LookupExperimentResultOutput)
}

type LookupExperimentOutputArgs struct {
	// The ARN of the experiment. For example, `arn:aws:evidently:us-west-2:0123455678912:project/myProject/experiment/myExperiment`
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupExperimentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExperimentArgs)(nil)).Elem()
}

type LookupExperimentResultOutput struct{ *pulumi.OutputState }

func (LookupExperimentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExperimentResult)(nil)).Elem()
}

func (o LookupExperimentResultOutput) ToLookupExperimentResultOutput() LookupExperimentResultOutput {
	return o
}

func (o LookupExperimentResultOutput) ToLookupExperimentResultOutputWithContext(ctx context.Context) LookupExperimentResultOutput {
	return o
}

// The ARN of the experiment. For example, `arn:aws:evidently:us-west-2:0123455678912:project/myProject/experiment/myExperiment`
func (o LookupExperimentResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// An optional description of the experiment.
func (o LookupExperimentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// An array of structures that defines the metrics used for the experiment, and whether a higher or lower value for each metric is the goal. You can use up to three metrics in an experiment.
func (o LookupExperimentResultOutput) MetricGoals() ExperimentMetricGoalObjectArrayOutput {
	return o.ApplyT(func(v LookupExperimentResult) []ExperimentMetricGoalObject { return v.MetricGoals }).(ExperimentMetricGoalObjectArrayOutput)
}

// A structure that contains the configuration of which variation to use as the "control" version. The "control" version is used for comparison with other variations. This structure also specifies how much experiment traffic is allocated to each variation.
func (o LookupExperimentResultOutput) OnlineAbConfig() ExperimentOnlineAbConfigObjectPtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *ExperimentOnlineAbConfigObject { return v.OnlineAbConfig }).(ExperimentOnlineAbConfigObjectPtrOutput)
}

// When Evidently assigns a particular user session to an experiment, it must use a randomization ID to determine which variation the user session is served. This randomization ID is a combination of the entity ID and `randomizationSalt` . If you omit `randomizationSalt` , Evidently uses the experiment name as the `randomizationSalt` .
func (o LookupExperimentResultOutput) RandomizationSalt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *string { return v.RandomizationSalt }).(pulumi.StringPtrOutput)
}

// Set this to `true` to remove the segment that is associated with this experiment. You can't use this parameter if the experiment is currently running.
func (o LookupExperimentResultOutput) RemoveSegment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *bool { return v.RemoveSegment }).(pulumi.BoolPtrOutput)
}

// Start Experiment. Default is False
func (o LookupExperimentResultOutput) RunningStatus() ExperimentRunningStatusObjectPtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *ExperimentRunningStatusObject { return v.RunningStatus }).(ExperimentRunningStatusObjectPtrOutput)
}

// The portion of the available audience that you want to allocate to this experiment, in thousandths of a percent. The available audience is the total audience minus the audience that you have allocated to overrides or current launches of this feature.
//
// This is represented in thousandths of a percent. For example, specify 10,000 to allocate 10% of the available audience.
func (o LookupExperimentResultOutput) SamplingRate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *int { return v.SamplingRate }).(pulumi.IntPtrOutput)
}

// Specifies an audience *segment* to use in the experiment. When a segment is used in an experiment, only user sessions that match the segment pattern are used in the experiment.
//
// For more information, see [Segment rule pattern syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-segments.html#CloudWatch-Evidently-segments-syntax) .
func (o LookupExperimentResultOutput) Segment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *string { return v.Segment }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupExperimentResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupExperimentResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// An array of structures that describe the configuration of each feature variation used in the experiment.
func (o LookupExperimentResultOutput) Treatments() ExperimentTreatmentObjectArrayOutput {
	return o.ApplyT(func(v LookupExperimentResult) []ExperimentTreatmentObject { return v.Treatments }).(ExperimentTreatmentObjectArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExperimentResultOutput{})
}
