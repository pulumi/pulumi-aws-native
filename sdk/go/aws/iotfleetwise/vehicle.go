// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotfleetwise

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::IoTFleetWise::Vehicle Resource Type
//
// Deprecated: Vehicle is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Vehicle struct {
	pulumi.CustomResourceState

	Arn                  pulumi.StringOutput                 `pulumi:"arn"`
	AssociationBehavior  VehicleAssociationBehaviorPtrOutput `pulumi:"associationBehavior"`
	Attributes           VehicleattributesMapPtrOutput       `pulumi:"attributes"`
	CreationTime         pulumi.StringOutput                 `pulumi:"creationTime"`
	DecoderManifestArn   pulumi.StringOutput                 `pulumi:"decoderManifestArn"`
	LastModificationTime pulumi.StringOutput                 `pulumi:"lastModificationTime"`
	ModelManifestArn     pulumi.StringOutput                 `pulumi:"modelManifestArn"`
	Name                 pulumi.StringOutput                 `pulumi:"name"`
	Tags                 VehicleTagArrayOutput               `pulumi:"tags"`
}

// NewVehicle registers a new resource with the given unique name, arguments, and options.
func NewVehicle(ctx *pulumi.Context,
	name string, args *VehicleArgs, opts ...pulumi.ResourceOption) (*Vehicle, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DecoderManifestArn == nil {
		return nil, errors.New("invalid value for required argument 'DecoderManifestArn'")
	}
	if args.ModelManifestArn == nil {
		return nil, errors.New("invalid value for required argument 'ModelManifestArn'")
	}
	var resource Vehicle
	err := ctx.RegisterResource("aws-native:iotfleetwise:Vehicle", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVehicle gets an existing Vehicle resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVehicle(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VehicleState, opts ...pulumi.ResourceOption) (*Vehicle, error) {
	var resource Vehicle
	err := ctx.ReadResource("aws-native:iotfleetwise:Vehicle", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vehicle resources.
type vehicleState struct {
}

type VehicleState struct {
}

func (VehicleState) ElementType() reflect.Type {
	return reflect.TypeOf((*vehicleState)(nil)).Elem()
}

type vehicleArgs struct {
	AssociationBehavior *VehicleAssociationBehavior `pulumi:"associationBehavior"`
	Attributes          *VehicleattributesMap       `pulumi:"attributes"`
	DecoderManifestArn  string                      `pulumi:"decoderManifestArn"`
	ModelManifestArn    string                      `pulumi:"modelManifestArn"`
	Name                *string                     `pulumi:"name"`
	Tags                []VehicleTag                `pulumi:"tags"`
}

// The set of arguments for constructing a Vehicle resource.
type VehicleArgs struct {
	AssociationBehavior VehicleAssociationBehaviorPtrInput
	Attributes          VehicleattributesMapPtrInput
	DecoderManifestArn  pulumi.StringInput
	ModelManifestArn    pulumi.StringInput
	Name                pulumi.StringPtrInput
	Tags                VehicleTagArrayInput
}

func (VehicleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vehicleArgs)(nil)).Elem()
}

type VehicleInput interface {
	pulumi.Input

	ToVehicleOutput() VehicleOutput
	ToVehicleOutputWithContext(ctx context.Context) VehicleOutput
}

func (*Vehicle) ElementType() reflect.Type {
	return reflect.TypeOf((**Vehicle)(nil)).Elem()
}

func (i *Vehicle) ToVehicleOutput() VehicleOutput {
	return i.ToVehicleOutputWithContext(context.Background())
}

func (i *Vehicle) ToVehicleOutputWithContext(ctx context.Context) VehicleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VehicleOutput)
}

type VehicleOutput struct{ *pulumi.OutputState }

func (VehicleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vehicle)(nil)).Elem()
}

func (o VehicleOutput) ToVehicleOutput() VehicleOutput {
	return o
}

func (o VehicleOutput) ToVehicleOutputWithContext(ctx context.Context) VehicleOutput {
	return o
}

func (o VehicleOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Vehicle) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o VehicleOutput) AssociationBehavior() VehicleAssociationBehaviorPtrOutput {
	return o.ApplyT(func(v *Vehicle) VehicleAssociationBehaviorPtrOutput { return v.AssociationBehavior }).(VehicleAssociationBehaviorPtrOutput)
}

func (o VehicleOutput) Attributes() VehicleattributesMapPtrOutput {
	return o.ApplyT(func(v *Vehicle) VehicleattributesMapPtrOutput { return v.Attributes }).(VehicleattributesMapPtrOutput)
}

func (o VehicleOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Vehicle) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o VehicleOutput) DecoderManifestArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Vehicle) pulumi.StringOutput { return v.DecoderManifestArn }).(pulumi.StringOutput)
}

func (o VehicleOutput) LastModificationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Vehicle) pulumi.StringOutput { return v.LastModificationTime }).(pulumi.StringOutput)
}

func (o VehicleOutput) ModelManifestArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Vehicle) pulumi.StringOutput { return v.ModelManifestArn }).(pulumi.StringOutput)
}

func (o VehicleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Vehicle) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VehicleOutput) Tags() VehicleTagArrayOutput {
	return o.ApplyT(func(v *Vehicle) VehicleTagArrayOutput { return v.Tags }).(VehicleTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VehicleInput)(nil)).Elem(), &Vehicle{})
	pulumi.RegisterOutputType(VehicleOutput{})
}