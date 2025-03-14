// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the AWS::QuickSight::RefreshSchedule Resource Type.
type RefreshSchedule struct {
	pulumi.CustomResourceState

	// <p>The Amazon Resource Name (ARN) of the data source.</p>
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The AWS account ID of the account that you are creating a schedule in.
	AwsAccountId pulumi.StringPtrOutput `pulumi:"awsAccountId"`
	// The ID of the dataset that you are creating a refresh schedule for.
	DataSetId pulumi.StringPtrOutput `pulumi:"dataSetId"`
	// The refresh schedule of a dataset.
	Schedule RefreshScheduleMapPtrOutput `pulumi:"schedule"`
}

// NewRefreshSchedule registers a new resource with the given unique name, arguments, and options.
func NewRefreshSchedule(ctx *pulumi.Context,
	name string, args *RefreshScheduleArgs, opts ...pulumi.ResourceOption) (*RefreshSchedule, error) {
	if args == nil {
		args = &RefreshScheduleArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"awsAccountId",
		"dataSetId",
		"schedule.scheduleId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RefreshSchedule
	err := ctx.RegisterResource("aws-native:quicksight:RefreshSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRefreshSchedule gets an existing RefreshSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRefreshSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RefreshScheduleState, opts ...pulumi.ResourceOption) (*RefreshSchedule, error) {
	var resource RefreshSchedule
	err := ctx.ReadResource("aws-native:quicksight:RefreshSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RefreshSchedule resources.
type refreshScheduleState struct {
}

type RefreshScheduleState struct {
}

func (RefreshScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*refreshScheduleState)(nil)).Elem()
}

type refreshScheduleArgs struct {
	// The AWS account ID of the account that you are creating a schedule in.
	AwsAccountId *string `pulumi:"awsAccountId"`
	// The ID of the dataset that you are creating a refresh schedule for.
	DataSetId *string `pulumi:"dataSetId"`
	// The refresh schedule of a dataset.
	Schedule *RefreshScheduleMap `pulumi:"schedule"`
}

// The set of arguments for constructing a RefreshSchedule resource.
type RefreshScheduleArgs struct {
	// The AWS account ID of the account that you are creating a schedule in.
	AwsAccountId pulumi.StringPtrInput
	// The ID of the dataset that you are creating a refresh schedule for.
	DataSetId pulumi.StringPtrInput
	// The refresh schedule of a dataset.
	Schedule RefreshScheduleMapPtrInput
}

func (RefreshScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*refreshScheduleArgs)(nil)).Elem()
}

type RefreshScheduleInput interface {
	pulumi.Input

	ToRefreshScheduleOutput() RefreshScheduleOutput
	ToRefreshScheduleOutputWithContext(ctx context.Context) RefreshScheduleOutput
}

func (*RefreshSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((**RefreshSchedule)(nil)).Elem()
}

func (i *RefreshSchedule) ToRefreshScheduleOutput() RefreshScheduleOutput {
	return i.ToRefreshScheduleOutputWithContext(context.Background())
}

func (i *RefreshSchedule) ToRefreshScheduleOutputWithContext(ctx context.Context) RefreshScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RefreshScheduleOutput)
}

type RefreshScheduleOutput struct{ *pulumi.OutputState }

func (RefreshScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RefreshSchedule)(nil)).Elem()
}

func (o RefreshScheduleOutput) ToRefreshScheduleOutput() RefreshScheduleOutput {
	return o
}

func (o RefreshScheduleOutput) ToRefreshScheduleOutputWithContext(ctx context.Context) RefreshScheduleOutput {
	return o
}

// <p>The Amazon Resource Name (ARN) of the data source.</p>
func (o RefreshScheduleOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *RefreshSchedule) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The AWS account ID of the account that you are creating a schedule in.
func (o RefreshScheduleOutput) AwsAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RefreshSchedule) pulumi.StringPtrOutput { return v.AwsAccountId }).(pulumi.StringPtrOutput)
}

// The ID of the dataset that you are creating a refresh schedule for.
func (o RefreshScheduleOutput) DataSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RefreshSchedule) pulumi.StringPtrOutput { return v.DataSetId }).(pulumi.StringPtrOutput)
}

// The refresh schedule of a dataset.
func (o RefreshScheduleOutput) Schedule() RefreshScheduleMapPtrOutput {
	return o.ApplyT(func(v *RefreshSchedule) RefreshScheduleMapPtrOutput { return v.Schedule }).(RefreshScheduleMapPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RefreshScheduleInput)(nil)).Elem(), &RefreshSchedule{})
	pulumi.RegisterOutputType(RefreshScheduleOutput{})
}
