// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package location

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Location::Tracker Resource Type
type Tracker struct {
	pulumi.CustomResourceState

	Arn                   pulumi.StringOutput               `pulumi:"arn"`
	CreateTime            pulumi.StringOutput               `pulumi:"createTime"`
	Description           pulumi.StringPtrOutput            `pulumi:"description"`
	KmsKeyId              pulumi.StringPtrOutput            `pulumi:"kmsKeyId"`
	PositionFiltering     TrackerPositionFilteringPtrOutput `pulumi:"positionFiltering"`
	PricingPlan           TrackerPricingPlanPtrOutput       `pulumi:"pricingPlan"`
	PricingPlanDataSource pulumi.StringPtrOutput            `pulumi:"pricingPlanDataSource"`
	TrackerArn            pulumi.StringOutput               `pulumi:"trackerArn"`
	TrackerName           pulumi.StringOutput               `pulumi:"trackerName"`
	UpdateTime            pulumi.StringOutput               `pulumi:"updateTime"`
}

// NewTracker registers a new resource with the given unique name, arguments, and options.
func NewTracker(ctx *pulumi.Context,
	name string, args *TrackerArgs, opts ...pulumi.ResourceOption) (*Tracker, error) {
	if args == nil {
		args = &TrackerArgs{}
	}

	var resource Tracker
	err := ctx.RegisterResource("aws-native:location:Tracker", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTracker gets an existing Tracker resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTracker(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrackerState, opts ...pulumi.ResourceOption) (*Tracker, error) {
	var resource Tracker
	err := ctx.ReadResource("aws-native:location:Tracker", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Tracker resources.
type trackerState struct {
}

type TrackerState struct {
}

func (TrackerState) ElementType() reflect.Type {
	return reflect.TypeOf((*trackerState)(nil)).Elem()
}

type trackerArgs struct {
	Description           *string                   `pulumi:"description"`
	KmsKeyId              *string                   `pulumi:"kmsKeyId"`
	PositionFiltering     *TrackerPositionFiltering `pulumi:"positionFiltering"`
	PricingPlan           *TrackerPricingPlan       `pulumi:"pricingPlan"`
	PricingPlanDataSource *string                   `pulumi:"pricingPlanDataSource"`
	TrackerName           *string                   `pulumi:"trackerName"`
}

// The set of arguments for constructing a Tracker resource.
type TrackerArgs struct {
	Description           pulumi.StringPtrInput
	KmsKeyId              pulumi.StringPtrInput
	PositionFiltering     TrackerPositionFilteringPtrInput
	PricingPlan           TrackerPricingPlanPtrInput
	PricingPlanDataSource pulumi.StringPtrInput
	TrackerName           pulumi.StringPtrInput
}

func (TrackerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trackerArgs)(nil)).Elem()
}

type TrackerInput interface {
	pulumi.Input

	ToTrackerOutput() TrackerOutput
	ToTrackerOutputWithContext(ctx context.Context) TrackerOutput
}

func (*Tracker) ElementType() reflect.Type {
	return reflect.TypeOf((**Tracker)(nil)).Elem()
}

func (i *Tracker) ToTrackerOutput() TrackerOutput {
	return i.ToTrackerOutputWithContext(context.Background())
}

func (i *Tracker) ToTrackerOutputWithContext(ctx context.Context) TrackerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackerOutput)
}

type TrackerOutput struct{ *pulumi.OutputState }

func (TrackerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Tracker)(nil)).Elem()
}

func (o TrackerOutput) ToTrackerOutput() TrackerOutput {
	return o
}

func (o TrackerOutput) ToTrackerOutputWithContext(ctx context.Context) TrackerOutput {
	return o
}

func (o TrackerOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Tracker) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o TrackerOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Tracker) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o TrackerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Tracker) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o TrackerOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Tracker) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

func (o TrackerOutput) PositionFiltering() TrackerPositionFilteringPtrOutput {
	return o.ApplyT(func(v *Tracker) TrackerPositionFilteringPtrOutput { return v.PositionFiltering }).(TrackerPositionFilteringPtrOutput)
}

func (o TrackerOutput) PricingPlan() TrackerPricingPlanPtrOutput {
	return o.ApplyT(func(v *Tracker) TrackerPricingPlanPtrOutput { return v.PricingPlan }).(TrackerPricingPlanPtrOutput)
}

func (o TrackerOutput) PricingPlanDataSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Tracker) pulumi.StringPtrOutput { return v.PricingPlanDataSource }).(pulumi.StringPtrOutput)
}

func (o TrackerOutput) TrackerArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Tracker) pulumi.StringOutput { return v.TrackerArn }).(pulumi.StringOutput)
}

func (o TrackerOutput) TrackerName() pulumi.StringOutput {
	return o.ApplyT(func(v *Tracker) pulumi.StringOutput { return v.TrackerName }).(pulumi.StringOutput)
}

func (o TrackerOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Tracker) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrackerInput)(nil)).Elem(), &Tracker{})
	pulumi.RegisterOutputType(TrackerOutput{})
}