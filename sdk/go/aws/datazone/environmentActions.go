// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datazone

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::DataZone::EnvironmentActions Resource Type
type EnvironmentActions struct {
	pulumi.CustomResourceState

	// The ID of the Amazon DataZone environment action.
	AwsId pulumi.StringOutput `pulumi:"awsId"`
	// The description of the Amazon DataZone environment action.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The identifier of the Amazon DataZone domain in which the environment is created.
	DomainId pulumi.StringOutput `pulumi:"domainId"`
	// The identifier of the Amazon DataZone domain in which the environment would be created.
	DomainIdentifier pulumi.StringPtrOutput `pulumi:"domainIdentifier"`
	// The identifier of the Amazon DataZone environment in which the action is taking place
	EnvironmentId pulumi.StringOutput `pulumi:"environmentId"`
	// The identifier of the Amazon DataZone environment in which the action is taking place
	EnvironmentIdentifier pulumi.StringPtrOutput `pulumi:"environmentIdentifier"`
	// The ID of the Amazon DataZone environment action.
	Identifier pulumi.StringPtrOutput `pulumi:"identifier"`
	// The name of the environment action.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parameters of the environment action.
	Parameters EnvironmentActionsAwsConsoleLinkParametersPtrOutput `pulumi:"parameters"`
}

// NewEnvironmentActions registers a new resource with the given unique name, arguments, and options.
func NewEnvironmentActions(ctx *pulumi.Context,
	name string, args *EnvironmentActionsArgs, opts ...pulumi.ResourceOption) (*EnvironmentActions, error) {
	if args == nil {
		args = &EnvironmentActionsArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"domainIdentifier",
		"environmentIdentifier",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EnvironmentActions
	err := ctx.RegisterResource("aws-native:datazone:EnvironmentActions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironmentActions gets an existing EnvironmentActions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironmentActions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentActionsState, opts ...pulumi.ResourceOption) (*EnvironmentActions, error) {
	var resource EnvironmentActions
	err := ctx.ReadResource("aws-native:datazone:EnvironmentActions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnvironmentActions resources.
type environmentActionsState struct {
}

type EnvironmentActionsState struct {
}

func (EnvironmentActionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentActionsState)(nil)).Elem()
}

type environmentActionsArgs struct {
	// The description of the Amazon DataZone environment action.
	Description *string `pulumi:"description"`
	// The identifier of the Amazon DataZone domain in which the environment would be created.
	DomainIdentifier *string `pulumi:"domainIdentifier"`
	// The identifier of the Amazon DataZone environment in which the action is taking place
	EnvironmentIdentifier *string `pulumi:"environmentIdentifier"`
	// The ID of the Amazon DataZone environment action.
	Identifier *string `pulumi:"identifier"`
	// The name of the environment action.
	Name *string `pulumi:"name"`
	// The parameters of the environment action.
	Parameters *EnvironmentActionsAwsConsoleLinkParameters `pulumi:"parameters"`
}

// The set of arguments for constructing a EnvironmentActions resource.
type EnvironmentActionsArgs struct {
	// The description of the Amazon DataZone environment action.
	Description pulumi.StringPtrInput
	// The identifier of the Amazon DataZone domain in which the environment would be created.
	DomainIdentifier pulumi.StringPtrInput
	// The identifier of the Amazon DataZone environment in which the action is taking place
	EnvironmentIdentifier pulumi.StringPtrInput
	// The ID of the Amazon DataZone environment action.
	Identifier pulumi.StringPtrInput
	// The name of the environment action.
	Name pulumi.StringPtrInput
	// The parameters of the environment action.
	Parameters EnvironmentActionsAwsConsoleLinkParametersPtrInput
}

func (EnvironmentActionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentActionsArgs)(nil)).Elem()
}

type EnvironmentActionsInput interface {
	pulumi.Input

	ToEnvironmentActionsOutput() EnvironmentActionsOutput
	ToEnvironmentActionsOutputWithContext(ctx context.Context) EnvironmentActionsOutput
}

func (*EnvironmentActions) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentActions)(nil)).Elem()
}

func (i *EnvironmentActions) ToEnvironmentActionsOutput() EnvironmentActionsOutput {
	return i.ToEnvironmentActionsOutputWithContext(context.Background())
}

func (i *EnvironmentActions) ToEnvironmentActionsOutputWithContext(ctx context.Context) EnvironmentActionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentActionsOutput)
}

type EnvironmentActionsOutput struct{ *pulumi.OutputState }

func (EnvironmentActionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentActions)(nil)).Elem()
}

func (o EnvironmentActionsOutput) ToEnvironmentActionsOutput() EnvironmentActionsOutput {
	return o
}

func (o EnvironmentActionsOutput) ToEnvironmentActionsOutputWithContext(ctx context.Context) EnvironmentActionsOutput {
	return o
}

// The ID of the Amazon DataZone environment action.
func (o EnvironmentActionsOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentActions) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// The description of the Amazon DataZone environment action.
func (o EnvironmentActionsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentActions) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The identifier of the Amazon DataZone domain in which the environment is created.
func (o EnvironmentActionsOutput) DomainId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentActions) pulumi.StringOutput { return v.DomainId }).(pulumi.StringOutput)
}

// The identifier of the Amazon DataZone domain in which the environment would be created.
func (o EnvironmentActionsOutput) DomainIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentActions) pulumi.StringPtrOutput { return v.DomainIdentifier }).(pulumi.StringPtrOutput)
}

// The identifier of the Amazon DataZone environment in which the action is taking place
func (o EnvironmentActionsOutput) EnvironmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentActions) pulumi.StringOutput { return v.EnvironmentId }).(pulumi.StringOutput)
}

// The identifier of the Amazon DataZone environment in which the action is taking place
func (o EnvironmentActionsOutput) EnvironmentIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentActions) pulumi.StringPtrOutput { return v.EnvironmentIdentifier }).(pulumi.StringPtrOutput)
}

// The ID of the Amazon DataZone environment action.
func (o EnvironmentActionsOutput) Identifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentActions) pulumi.StringPtrOutput { return v.Identifier }).(pulumi.StringPtrOutput)
}

// The name of the environment action.
func (o EnvironmentActionsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentActions) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The parameters of the environment action.
func (o EnvironmentActionsOutput) Parameters() EnvironmentActionsAwsConsoleLinkParametersPtrOutput {
	return o.ApplyT(func(v *EnvironmentActions) EnvironmentActionsAwsConsoleLinkParametersPtrOutput { return v.Parameters }).(EnvironmentActionsAwsConsoleLinkParametersPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentActionsInput)(nil)).Elem(), &EnvironmentActions{})
	pulumi.RegisterOutputType(EnvironmentActionsOutput{})
}