// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Mitigation actions can be used to take actions to mitigate issues that were found in an Audit finding or Detect violation.
type MitigationAction struct {
	pulumi.CustomResourceState

	// A unique identifier for the mitigation action.
	ActionName          pulumi.StringPtrOutput             `pulumi:"actionName"`
	ActionParams        MitigationActionActionParamsOutput `pulumi:"actionParams"`
	MitigationActionArn pulumi.StringOutput                `pulumi:"mitigationActionArn"`
	MitigationActionId  pulumi.StringOutput                `pulumi:"mitigationActionId"`
	RoleArn             pulumi.StringOutput                `pulumi:"roleArn"`
	// An array of key-value pairs to apply to this resource.
	Tags MitigationActionTagArrayOutput `pulumi:"tags"`
}

// NewMitigationAction registers a new resource with the given unique name, arguments, and options.
func NewMitigationAction(ctx *pulumi.Context,
	name string, args *MitigationActionArgs, opts ...pulumi.ResourceOption) (*MitigationAction, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ActionParams == nil {
		return nil, errors.New("invalid value for required argument 'ActionParams'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource MitigationAction
	err := ctx.RegisterResource("aws-native:iot:MitigationAction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMitigationAction gets an existing MitigationAction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMitigationAction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MitigationActionState, opts ...pulumi.ResourceOption) (*MitigationAction, error) {
	var resource MitigationAction
	err := ctx.ReadResource("aws-native:iot:MitigationAction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MitigationAction resources.
type mitigationActionState struct {
}

type MitigationActionState struct {
}

func (MitigationActionState) ElementType() reflect.Type {
	return reflect.TypeOf((*mitigationActionState)(nil)).Elem()
}

type mitigationActionArgs struct {
	// A unique identifier for the mitigation action.
	ActionName   *string                      `pulumi:"actionName"`
	ActionParams MitigationActionActionParams `pulumi:"actionParams"`
	RoleArn      string                       `pulumi:"roleArn"`
	// An array of key-value pairs to apply to this resource.
	Tags []MitigationActionTag `pulumi:"tags"`
}

// The set of arguments for constructing a MitigationAction resource.
type MitigationActionArgs struct {
	// A unique identifier for the mitigation action.
	ActionName   pulumi.StringPtrInput
	ActionParams MitigationActionActionParamsInput
	RoleArn      pulumi.StringInput
	// An array of key-value pairs to apply to this resource.
	Tags MitigationActionTagArrayInput
}

func (MitigationActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mitigationActionArgs)(nil)).Elem()
}

type MitigationActionInput interface {
	pulumi.Input

	ToMitigationActionOutput() MitigationActionOutput
	ToMitigationActionOutputWithContext(ctx context.Context) MitigationActionOutput
}

func (*MitigationAction) ElementType() reflect.Type {
	return reflect.TypeOf((**MitigationAction)(nil)).Elem()
}

func (i *MitigationAction) ToMitigationActionOutput() MitigationActionOutput {
	return i.ToMitigationActionOutputWithContext(context.Background())
}

func (i *MitigationAction) ToMitigationActionOutputWithContext(ctx context.Context) MitigationActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MitigationActionOutput)
}

type MitigationActionOutput struct{ *pulumi.OutputState }

func (MitigationActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MitigationAction)(nil)).Elem()
}

func (o MitigationActionOutput) ToMitigationActionOutput() MitigationActionOutput {
	return o
}

func (o MitigationActionOutput) ToMitigationActionOutputWithContext(ctx context.Context) MitigationActionOutput {
	return o
}

// A unique identifier for the mitigation action.
func (o MitigationActionOutput) ActionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MitigationAction) pulumi.StringPtrOutput { return v.ActionName }).(pulumi.StringPtrOutput)
}

func (o MitigationActionOutput) ActionParams() MitigationActionActionParamsOutput {
	return o.ApplyT(func(v *MitigationAction) MitigationActionActionParamsOutput { return v.ActionParams }).(MitigationActionActionParamsOutput)
}

func (o MitigationActionOutput) MitigationActionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *MitigationAction) pulumi.StringOutput { return v.MitigationActionArn }).(pulumi.StringOutput)
}

func (o MitigationActionOutput) MitigationActionId() pulumi.StringOutput {
	return o.ApplyT(func(v *MitigationAction) pulumi.StringOutput { return v.MitigationActionId }).(pulumi.StringOutput)
}

func (o MitigationActionOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *MitigationAction) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

// An array of key-value pairs to apply to this resource.
func (o MitigationActionOutput) Tags() MitigationActionTagArrayOutput {
	return o.ApplyT(func(v *MitigationAction) MitigationActionTagArrayOutput { return v.Tags }).(MitigationActionTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MitigationActionInput)(nil)).Elem(), &MitigationAction{})
	pulumi.RegisterOutputType(MitigationActionOutput{})
}