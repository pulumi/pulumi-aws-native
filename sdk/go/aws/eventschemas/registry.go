// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventschemas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EventSchemas::Registry
//
// Deprecated: Registry is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Registry struct {
	pulumi.CustomResourceState

	Description  pulumi.StringPtrOutput       `pulumi:"description"`
	RegistryArn  pulumi.StringOutput          `pulumi:"registryArn"`
	RegistryName pulumi.StringPtrOutput       `pulumi:"registryName"`
	Tags         RegistryTagsEntryArrayOutput `pulumi:"tags"`
}

// NewRegistry registers a new resource with the given unique name, arguments, and options.
func NewRegistry(ctx *pulumi.Context,
	name string, args *RegistryArgs, opts ...pulumi.ResourceOption) (*Registry, error) {
	if args == nil {
		args = &RegistryArgs{}
	}

	var resource Registry
	err := ctx.RegisterResource("aws-native:eventschemas:Registry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistry gets an existing Registry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryState, opts ...pulumi.ResourceOption) (*Registry, error) {
	var resource Registry
	err := ctx.ReadResource("aws-native:eventschemas:Registry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Registry resources.
type registryState struct {
}

type RegistryState struct {
}

func (RegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryState)(nil)).Elem()
}

type registryArgs struct {
	Description  *string             `pulumi:"description"`
	RegistryName *string             `pulumi:"registryName"`
	Tags         []RegistryTagsEntry `pulumi:"tags"`
}

// The set of arguments for constructing a Registry resource.
type RegistryArgs struct {
	Description  pulumi.StringPtrInput
	RegistryName pulumi.StringPtrInput
	Tags         RegistryTagsEntryArrayInput
}

func (RegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryArgs)(nil)).Elem()
}

type RegistryInput interface {
	pulumi.Input

	ToRegistryOutput() RegistryOutput
	ToRegistryOutputWithContext(ctx context.Context) RegistryOutput
}

func (*Registry) ElementType() reflect.Type {
	return reflect.TypeOf((**Registry)(nil)).Elem()
}

func (i *Registry) ToRegistryOutput() RegistryOutput {
	return i.ToRegistryOutputWithContext(context.Background())
}

func (i *Registry) ToRegistryOutputWithContext(ctx context.Context) RegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryOutput)
}

type RegistryOutput struct{ *pulumi.OutputState }

func (RegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Registry)(nil)).Elem()
}

func (o RegistryOutput) ToRegistryOutput() RegistryOutput {
	return o
}

func (o RegistryOutput) ToRegistryOutputWithContext(ctx context.Context) RegistryOutput {
	return o
}

func (o RegistryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RegistryOutput) RegistryArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringOutput { return v.RegistryArn }).(pulumi.StringOutput)
}

func (o RegistryOutput) RegistryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringPtrOutput { return v.RegistryName }).(pulumi.StringPtrOutput)
}

func (o RegistryOutput) Tags() RegistryTagsEntryArrayOutput {
	return o.ApplyT(func(v *Registry) RegistryTagsEntryArrayOutput { return v.Tags }).(RegistryTagsEntryArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryInput)(nil)).Elem(), &Registry{})
	pulumi.RegisterOutputType(RegistryOutput{})
}