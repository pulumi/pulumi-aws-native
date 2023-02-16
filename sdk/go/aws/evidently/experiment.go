// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package evidently

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Evidently::Experiment.
type Experiment struct {
	pulumi.CustomResourceState

	Arn               pulumi.StringOutput                   `pulumi:"arn"`
	Description       pulumi.StringPtrOutput                `pulumi:"description"`
	MetricGoals       ExperimentMetricGoalObjectArrayOutput `pulumi:"metricGoals"`
	Name              pulumi.StringOutput                   `pulumi:"name"`
	OnlineAbConfig    ExperimentOnlineAbConfigObjectOutput  `pulumi:"onlineAbConfig"`
	Project           pulumi.StringOutput                   `pulumi:"project"`
	RandomizationSalt pulumi.StringPtrOutput                `pulumi:"randomizationSalt"`
	RemoveSegment     pulumi.BoolPtrOutput                  `pulumi:"removeSegment"`
	// Start Experiment. Default is False
	RunningStatus ExperimentRunningStatusObjectPtrOutput `pulumi:"runningStatus"`
	SamplingRate  pulumi.IntPtrOutput                    `pulumi:"samplingRate"`
	Segment       pulumi.StringPtrOutput                 `pulumi:"segment"`
	// An array of key-value pairs to apply to this resource.
	Tags       ExperimentTagArrayOutput             `pulumi:"tags"`
	Treatments ExperimentTreatmentObjectArrayOutput `pulumi:"treatments"`
}

// NewExperiment registers a new resource with the given unique name, arguments, and options.
func NewExperiment(ctx *pulumi.Context,
	name string, args *ExperimentArgs, opts ...pulumi.ResourceOption) (*Experiment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MetricGoals == nil {
		return nil, errors.New("invalid value for required argument 'MetricGoals'")
	}
	if args.OnlineAbConfig == nil {
		return nil, errors.New("invalid value for required argument 'OnlineAbConfig'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Treatments == nil {
		return nil, errors.New("invalid value for required argument 'Treatments'")
	}
	var resource Experiment
	err := ctx.RegisterResource("aws-native:evidently:Experiment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExperiment gets an existing Experiment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExperiment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExperimentState, opts ...pulumi.ResourceOption) (*Experiment, error) {
	var resource Experiment
	err := ctx.ReadResource("aws-native:evidently:Experiment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Experiment resources.
type experimentState struct {
}

type ExperimentState struct {
}

func (ExperimentState) ElementType() reflect.Type {
	return reflect.TypeOf((*experimentState)(nil)).Elem()
}

type experimentArgs struct {
	Description       *string                        `pulumi:"description"`
	MetricGoals       []ExperimentMetricGoalObject   `pulumi:"metricGoals"`
	Name              *string                        `pulumi:"name"`
	OnlineAbConfig    ExperimentOnlineAbConfigObject `pulumi:"onlineAbConfig"`
	Project           string                         `pulumi:"project"`
	RandomizationSalt *string                        `pulumi:"randomizationSalt"`
	RemoveSegment     *bool                          `pulumi:"removeSegment"`
	// Start Experiment. Default is False
	RunningStatus *ExperimentRunningStatusObject `pulumi:"runningStatus"`
	SamplingRate  *int                           `pulumi:"samplingRate"`
	Segment       *string                        `pulumi:"segment"`
	// An array of key-value pairs to apply to this resource.
	Tags       []ExperimentTag             `pulumi:"tags"`
	Treatments []ExperimentTreatmentObject `pulumi:"treatments"`
}

// The set of arguments for constructing a Experiment resource.
type ExperimentArgs struct {
	Description       pulumi.StringPtrInput
	MetricGoals       ExperimentMetricGoalObjectArrayInput
	Name              pulumi.StringPtrInput
	OnlineAbConfig    ExperimentOnlineAbConfigObjectInput
	Project           pulumi.StringInput
	RandomizationSalt pulumi.StringPtrInput
	RemoveSegment     pulumi.BoolPtrInput
	// Start Experiment. Default is False
	RunningStatus ExperimentRunningStatusObjectPtrInput
	SamplingRate  pulumi.IntPtrInput
	Segment       pulumi.StringPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags       ExperimentTagArrayInput
	Treatments ExperimentTreatmentObjectArrayInput
}

func (ExperimentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*experimentArgs)(nil)).Elem()
}

type ExperimentInput interface {
	pulumi.Input

	ToExperimentOutput() ExperimentOutput
	ToExperimentOutputWithContext(ctx context.Context) ExperimentOutput
}

func (*Experiment) ElementType() reflect.Type {
	return reflect.TypeOf((**Experiment)(nil)).Elem()
}

func (i *Experiment) ToExperimentOutput() ExperimentOutput {
	return i.ToExperimentOutputWithContext(context.Background())
}

func (i *Experiment) ToExperimentOutputWithContext(ctx context.Context) ExperimentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentOutput)
}

type ExperimentOutput struct{ *pulumi.OutputState }

func (ExperimentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Experiment)(nil)).Elem()
}

func (o ExperimentOutput) ToExperimentOutput() ExperimentOutput {
	return o
}

func (o ExperimentOutput) ToExperimentOutputWithContext(ctx context.Context) ExperimentOutput {
	return o
}

func (o ExperimentOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o ExperimentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ExperimentOutput) MetricGoals() ExperimentMetricGoalObjectArrayOutput {
	return o.ApplyT(func(v *Experiment) ExperimentMetricGoalObjectArrayOutput { return v.MetricGoals }).(ExperimentMetricGoalObjectArrayOutput)
}

func (o ExperimentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ExperimentOutput) OnlineAbConfig() ExperimentOnlineAbConfigObjectOutput {
	return o.ApplyT(func(v *Experiment) ExperimentOnlineAbConfigObjectOutput { return v.OnlineAbConfig }).(ExperimentOnlineAbConfigObjectOutput)
}

func (o ExperimentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ExperimentOutput) RandomizationSalt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringPtrOutput { return v.RandomizationSalt }).(pulumi.StringPtrOutput)
}

func (o ExperimentOutput) RemoveSegment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Experiment) pulumi.BoolPtrOutput { return v.RemoveSegment }).(pulumi.BoolPtrOutput)
}

// Start Experiment. Default is False
func (o ExperimentOutput) RunningStatus() ExperimentRunningStatusObjectPtrOutput {
	return o.ApplyT(func(v *Experiment) ExperimentRunningStatusObjectPtrOutput { return v.RunningStatus }).(ExperimentRunningStatusObjectPtrOutput)
}

func (o ExperimentOutput) SamplingRate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Experiment) pulumi.IntPtrOutput { return v.SamplingRate }).(pulumi.IntPtrOutput)
}

func (o ExperimentOutput) Segment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringPtrOutput { return v.Segment }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o ExperimentOutput) Tags() ExperimentTagArrayOutput {
	return o.ApplyT(func(v *Experiment) ExperimentTagArrayOutput { return v.Tags }).(ExperimentTagArrayOutput)
}

func (o ExperimentOutput) Treatments() ExperimentTreatmentObjectArrayOutput {
	return o.ApplyT(func(v *Experiment) ExperimentTreatmentObjectArrayOutput { return v.Treatments }).(ExperimentTreatmentObjectArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExperimentInput)(nil)).Elem(), &Experiment{})
	pulumi.RegisterOutputType(ExperimentOutput{})
}