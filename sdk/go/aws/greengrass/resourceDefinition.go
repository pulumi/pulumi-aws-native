// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package greengrass

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Greengrass::ResourceDefinition
//
// Deprecated: ResourceDefinition is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ResourceDefinition struct {
	pulumi.CustomResourceState

	Arn              pulumi.StringOutput                    `pulumi:"arn"`
	InitialVersion   ResourceDefinitionVersionTypePtrOutput `pulumi:"initialVersion"`
	LatestVersionArn pulumi.StringOutput                    `pulumi:"latestVersionArn"`
	Name             pulumi.StringOutput                    `pulumi:"name"`
	Tags             pulumi.AnyOutput                       `pulumi:"tags"`
}

// NewResourceDefinition registers a new resource with the given unique name, arguments, and options.
func NewResourceDefinition(ctx *pulumi.Context,
	name string, args *ResourceDefinitionArgs, opts ...pulumi.ResourceOption) (*ResourceDefinition, error) {
	if args == nil {
		args = &ResourceDefinitionArgs{}
	}

	var resource ResourceDefinition
	err := ctx.RegisterResource("aws-native:greengrass:ResourceDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceDefinition gets an existing ResourceDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceDefinitionState, opts ...pulumi.ResourceOption) (*ResourceDefinition, error) {
	var resource ResourceDefinition
	err := ctx.ReadResource("aws-native:greengrass:ResourceDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceDefinition resources.
type resourceDefinitionState struct {
}

type ResourceDefinitionState struct {
}

func (ResourceDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceDefinitionState)(nil)).Elem()
}

type resourceDefinitionArgs struct {
	InitialVersion *ResourceDefinitionVersionType `pulumi:"initialVersion"`
	Name           *string                        `pulumi:"name"`
	Tags           interface{}                    `pulumi:"tags"`
}

// The set of arguments for constructing a ResourceDefinition resource.
type ResourceDefinitionArgs struct {
	InitialVersion ResourceDefinitionVersionTypePtrInput
	Name           pulumi.StringPtrInput
	Tags           pulumi.Input
}

func (ResourceDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceDefinitionArgs)(nil)).Elem()
}

type ResourceDefinitionInput interface {
	pulumi.Input

	ToResourceDefinitionOutput() ResourceDefinitionOutput
	ToResourceDefinitionOutputWithContext(ctx context.Context) ResourceDefinitionOutput
}

func (*ResourceDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceDefinition)(nil)).Elem()
}

func (i *ResourceDefinition) ToResourceDefinitionOutput() ResourceDefinitionOutput {
	return i.ToResourceDefinitionOutputWithContext(context.Background())
}

func (i *ResourceDefinition) ToResourceDefinitionOutputWithContext(ctx context.Context) ResourceDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceDefinitionOutput)
}

type ResourceDefinitionOutput struct{ *pulumi.OutputState }

func (ResourceDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceDefinition)(nil)).Elem()
}

func (o ResourceDefinitionOutput) ToResourceDefinitionOutput() ResourceDefinitionOutput {
	return o
}

func (o ResourceDefinitionOutput) ToResourceDefinitionOutputWithContext(ctx context.Context) ResourceDefinitionOutput {
	return o
}

func (o ResourceDefinitionOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceDefinition) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o ResourceDefinitionOutput) InitialVersion() ResourceDefinitionVersionTypePtrOutput {
	return o.ApplyT(func(v *ResourceDefinition) ResourceDefinitionVersionTypePtrOutput { return v.InitialVersion }).(ResourceDefinitionVersionTypePtrOutput)
}

func (o ResourceDefinitionOutput) LatestVersionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceDefinition) pulumi.StringOutput { return v.LatestVersionArn }).(pulumi.StringOutput)
}

func (o ResourceDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceDefinition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceDefinitionOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *ResourceDefinition) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceDefinitionInput)(nil)).Elem(), &ResourceDefinition{})
	pulumi.RegisterOutputType(ResourceDefinitionOutput{})
}