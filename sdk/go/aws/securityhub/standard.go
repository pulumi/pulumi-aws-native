// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::SecurityHub::Standard resource represents the implementation of an individual AWS Security Hub Standard in your account. It requires you have SecurityHub enabled before you can enable the Standard.
type Standard struct {
	pulumi.CustomResourceState

	// StandardsControls to disable from this Standard.
	DisabledStandardsControls StandardsControlArrayOutput `pulumi:"disabledStandardsControls"`
	// The ARN of the Standard being enabled
	StandardsArn pulumi.StringOutput `pulumi:"standardsArn"`
	// The ARN of the StandardsSubscription for the account ID, region, and Standard.
	StandardsSubscriptionArn pulumi.StringOutput `pulumi:"standardsSubscriptionArn"`
}

// NewStandard registers a new resource with the given unique name, arguments, and options.
func NewStandard(ctx *pulumi.Context,
	name string, args *StandardArgs, opts ...pulumi.ResourceOption) (*Standard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StandardsArn == nil {
		return nil, errors.New("invalid value for required argument 'StandardsArn'")
	}
	var resource Standard
	err := ctx.RegisterResource("aws-native:securityhub:Standard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStandard gets an existing Standard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStandard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StandardState, opts ...pulumi.ResourceOption) (*Standard, error) {
	var resource Standard
	err := ctx.ReadResource("aws-native:securityhub:Standard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Standard resources.
type standardState struct {
}

type StandardState struct {
}

func (StandardState) ElementType() reflect.Type {
	return reflect.TypeOf((*standardState)(nil)).Elem()
}

type standardArgs struct {
	// StandardsControls to disable from this Standard.
	DisabledStandardsControls []StandardsControl `pulumi:"disabledStandardsControls"`
	// The ARN of the Standard being enabled
	StandardsArn string `pulumi:"standardsArn"`
}

// The set of arguments for constructing a Standard resource.
type StandardArgs struct {
	// StandardsControls to disable from this Standard.
	DisabledStandardsControls StandardsControlArrayInput
	// The ARN of the Standard being enabled
	StandardsArn pulumi.StringInput
}

func (StandardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*standardArgs)(nil)).Elem()
}

type StandardInput interface {
	pulumi.Input

	ToStandardOutput() StandardOutput
	ToStandardOutputWithContext(ctx context.Context) StandardOutput
}

func (*Standard) ElementType() reflect.Type {
	return reflect.TypeOf((**Standard)(nil)).Elem()
}

func (i *Standard) ToStandardOutput() StandardOutput {
	return i.ToStandardOutputWithContext(context.Background())
}

func (i *Standard) ToStandardOutputWithContext(ctx context.Context) StandardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StandardOutput)
}

type StandardOutput struct{ *pulumi.OutputState }

func (StandardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Standard)(nil)).Elem()
}

func (o StandardOutput) ToStandardOutput() StandardOutput {
	return o
}

func (o StandardOutput) ToStandardOutputWithContext(ctx context.Context) StandardOutput {
	return o
}

// StandardsControls to disable from this Standard.
func (o StandardOutput) DisabledStandardsControls() StandardsControlArrayOutput {
	return o.ApplyT(func(v *Standard) StandardsControlArrayOutput { return v.DisabledStandardsControls }).(StandardsControlArrayOutput)
}

// The ARN of the Standard being enabled
func (o StandardOutput) StandardsArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Standard) pulumi.StringOutput { return v.StandardsArn }).(pulumi.StringOutput)
}

// The ARN of the StandardsSubscription for the account ID, region, and Standard.
func (o StandardOutput) StandardsSubscriptionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Standard) pulumi.StringOutput { return v.StandardsSubscriptionArn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StandardInput)(nil)).Elem(), &Standard{})
	pulumi.RegisterOutputType(StandardOutput{})
}