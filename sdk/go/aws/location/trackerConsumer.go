// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package location

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Location::TrackerConsumer Resource Type
type TrackerConsumer struct {
	pulumi.CustomResourceState

	ConsumerArn pulumi.StringOutput `pulumi:"consumerArn"`
	TrackerName pulumi.StringOutput `pulumi:"trackerName"`
}

// NewTrackerConsumer registers a new resource with the given unique name, arguments, and options.
func NewTrackerConsumer(ctx *pulumi.Context,
	name string, args *TrackerConsumerArgs, opts ...pulumi.ResourceOption) (*TrackerConsumer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsumerArn == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerArn'")
	}
	if args.TrackerName == nil {
		return nil, errors.New("invalid value for required argument 'TrackerName'")
	}
	var resource TrackerConsumer
	err := ctx.RegisterResource("aws-native:location:TrackerConsumer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrackerConsumer gets an existing TrackerConsumer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrackerConsumer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrackerConsumerState, opts ...pulumi.ResourceOption) (*TrackerConsumer, error) {
	var resource TrackerConsumer
	err := ctx.ReadResource("aws-native:location:TrackerConsumer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrackerConsumer resources.
type trackerConsumerState struct {
}

type TrackerConsumerState struct {
}

func (TrackerConsumerState) ElementType() reflect.Type {
	return reflect.TypeOf((*trackerConsumerState)(nil)).Elem()
}

type trackerConsumerArgs struct {
	ConsumerArn string `pulumi:"consumerArn"`
	TrackerName string `pulumi:"trackerName"`
}

// The set of arguments for constructing a TrackerConsumer resource.
type TrackerConsumerArgs struct {
	ConsumerArn pulumi.StringInput
	TrackerName pulumi.StringInput
}

func (TrackerConsumerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trackerConsumerArgs)(nil)).Elem()
}

type TrackerConsumerInput interface {
	pulumi.Input

	ToTrackerConsumerOutput() TrackerConsumerOutput
	ToTrackerConsumerOutputWithContext(ctx context.Context) TrackerConsumerOutput
}

func (*TrackerConsumer) ElementType() reflect.Type {
	return reflect.TypeOf((**TrackerConsumer)(nil)).Elem()
}

func (i *TrackerConsumer) ToTrackerConsumerOutput() TrackerConsumerOutput {
	return i.ToTrackerConsumerOutputWithContext(context.Background())
}

func (i *TrackerConsumer) ToTrackerConsumerOutputWithContext(ctx context.Context) TrackerConsumerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackerConsumerOutput)
}

type TrackerConsumerOutput struct{ *pulumi.OutputState }

func (TrackerConsumerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrackerConsumer)(nil)).Elem()
}

func (o TrackerConsumerOutput) ToTrackerConsumerOutput() TrackerConsumerOutput {
	return o
}

func (o TrackerConsumerOutput) ToTrackerConsumerOutputWithContext(ctx context.Context) TrackerConsumerOutput {
	return o
}

func (o TrackerConsumerOutput) ConsumerArn() pulumi.StringOutput {
	return o.ApplyT(func(v *TrackerConsumer) pulumi.StringOutput { return v.ConsumerArn }).(pulumi.StringOutput)
}

func (o TrackerConsumerOutput) TrackerName() pulumi.StringOutput {
	return o.ApplyT(func(v *TrackerConsumer) pulumi.StringOutput { return v.TrackerName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrackerConsumerInput)(nil)).Elem(), &TrackerConsumer{})
	pulumi.RegisterOutputType(TrackerConsumerOutput{})
}