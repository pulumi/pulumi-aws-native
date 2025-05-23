// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datazone

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A owner can set up authorization permissions on their resources.
type Owner struct {
	pulumi.CustomResourceState

	// The ID of the domain in which you want to add the entity owner.
	DomainIdentifier pulumi.StringOutput `pulumi:"domainIdentifier"`
	// The ID of the entity to which you want to add an owner.
	EntityIdentifier pulumi.StringOutput `pulumi:"entityIdentifier"`
	// The type of an entity.
	EntityType OwnerEntityTypeOutput `pulumi:"entityType"`
	// The owner that you want to add to the entity.
	Owner OwnerPropertiesOutput `pulumi:"owner"`
}

// NewOwner registers a new resource with the given unique name, arguments, and options.
func NewOwner(ctx *pulumi.Context,
	name string, args *OwnerArgs, opts ...pulumi.ResourceOption) (*Owner, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'DomainIdentifier'")
	}
	if args.EntityIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'EntityIdentifier'")
	}
	if args.EntityType == nil {
		return nil, errors.New("invalid value for required argument 'EntityType'")
	}
	if args.Owner == nil {
		return nil, errors.New("invalid value for required argument 'Owner'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"domainIdentifier",
		"entityIdentifier",
		"entityType",
		"owner",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Owner
	err := ctx.RegisterResource("aws-native:datazone:Owner", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOwner gets an existing Owner resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOwner(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OwnerState, opts ...pulumi.ResourceOption) (*Owner, error) {
	var resource Owner
	err := ctx.ReadResource("aws-native:datazone:Owner", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Owner resources.
type ownerState struct {
}

type OwnerState struct {
}

func (OwnerState) ElementType() reflect.Type {
	return reflect.TypeOf((*ownerState)(nil)).Elem()
}

type ownerArgs struct {
	// The ID of the domain in which you want to add the entity owner.
	DomainIdentifier string `pulumi:"domainIdentifier"`
	// The ID of the entity to which you want to add an owner.
	EntityIdentifier string `pulumi:"entityIdentifier"`
	// The type of an entity.
	EntityType OwnerEntityType `pulumi:"entityType"`
	// The owner that you want to add to the entity.
	Owner OwnerProperties `pulumi:"owner"`
}

// The set of arguments for constructing a Owner resource.
type OwnerArgs struct {
	// The ID of the domain in which you want to add the entity owner.
	DomainIdentifier pulumi.StringInput
	// The ID of the entity to which you want to add an owner.
	EntityIdentifier pulumi.StringInput
	// The type of an entity.
	EntityType OwnerEntityTypeInput
	// The owner that you want to add to the entity.
	Owner OwnerPropertiesInput
}

func (OwnerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ownerArgs)(nil)).Elem()
}

type OwnerInput interface {
	pulumi.Input

	ToOwnerOutput() OwnerOutput
	ToOwnerOutputWithContext(ctx context.Context) OwnerOutput
}

func (*Owner) ElementType() reflect.Type {
	return reflect.TypeOf((**Owner)(nil)).Elem()
}

func (i *Owner) ToOwnerOutput() OwnerOutput {
	return i.ToOwnerOutputWithContext(context.Background())
}

func (i *Owner) ToOwnerOutputWithContext(ctx context.Context) OwnerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OwnerOutput)
}

type OwnerOutput struct{ *pulumi.OutputState }

func (OwnerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Owner)(nil)).Elem()
}

func (o OwnerOutput) ToOwnerOutput() OwnerOutput {
	return o
}

func (o OwnerOutput) ToOwnerOutputWithContext(ctx context.Context) OwnerOutput {
	return o
}

// The ID of the domain in which you want to add the entity owner.
func (o OwnerOutput) DomainIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Owner) pulumi.StringOutput { return v.DomainIdentifier }).(pulumi.StringOutput)
}

// The ID of the entity to which you want to add an owner.
func (o OwnerOutput) EntityIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Owner) pulumi.StringOutput { return v.EntityIdentifier }).(pulumi.StringOutput)
}

// The type of an entity.
func (o OwnerOutput) EntityType() OwnerEntityTypeOutput {
	return o.ApplyT(func(v *Owner) OwnerEntityTypeOutput { return v.EntityType }).(OwnerEntityTypeOutput)
}

// The owner that you want to add to the entity.
func (o OwnerOutput) Owner() OwnerPropertiesOutput {
	return o.ApplyT(func(v *Owner) OwnerPropertiesOutput { return v.Owner }).(OwnerPropertiesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OwnerInput)(nil)).Elem(), &Owner{})
	pulumi.RegisterOutputType(OwnerOutput{})
}
