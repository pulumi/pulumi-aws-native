// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes an association between a local gateway route table and a VPC.
type LocalGatewayRouteTableVPCAssociation struct {
	pulumi.CustomResourceState

	// The ID of the local gateway.
	LocalGatewayId pulumi.StringOutput `pulumi:"localGatewayId"`
	// The ID of the local gateway route table.
	LocalGatewayRouteTableId pulumi.StringOutput `pulumi:"localGatewayRouteTableId"`
	// The ID of the association.
	LocalGatewayRouteTableVpcAssociationId pulumi.StringOutput `pulumi:"localGatewayRouteTableVpcAssociationId"`
	// The state of the association.
	State pulumi.StringOutput `pulumi:"state"`
	// The tags for the association.
	Tags LocalGatewayRouteTableVPCAssociationTagArrayOutput `pulumi:"tags"`
	// The ID of the VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewLocalGatewayRouteTableVPCAssociation registers a new resource with the given unique name, arguments, and options.
func NewLocalGatewayRouteTableVPCAssociation(ctx *pulumi.Context,
	name string, args *LocalGatewayRouteTableVPCAssociationArgs, opts ...pulumi.ResourceOption) (*LocalGatewayRouteTableVPCAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LocalGatewayRouteTableId == nil {
		return nil, errors.New("invalid value for required argument 'LocalGatewayRouteTableId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource LocalGatewayRouteTableVPCAssociation
	err := ctx.RegisterResource("aws-native:ec2:LocalGatewayRouteTableVPCAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalGatewayRouteTableVPCAssociation gets an existing LocalGatewayRouteTableVPCAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalGatewayRouteTableVPCAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalGatewayRouteTableVPCAssociationState, opts ...pulumi.ResourceOption) (*LocalGatewayRouteTableVPCAssociation, error) {
	var resource LocalGatewayRouteTableVPCAssociation
	err := ctx.ReadResource("aws-native:ec2:LocalGatewayRouteTableVPCAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalGatewayRouteTableVPCAssociation resources.
type localGatewayRouteTableVPCAssociationState struct {
}

type LocalGatewayRouteTableVPCAssociationState struct {
}

func (LocalGatewayRouteTableVPCAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*localGatewayRouteTableVPCAssociationState)(nil)).Elem()
}

type localGatewayRouteTableVPCAssociationArgs struct {
	// The ID of the local gateway route table.
	LocalGatewayRouteTableId string `pulumi:"localGatewayRouteTableId"`
	// The tags for the association.
	Tags []LocalGatewayRouteTableVPCAssociationTag `pulumi:"tags"`
	// The ID of the VPC.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a LocalGatewayRouteTableVPCAssociation resource.
type LocalGatewayRouteTableVPCAssociationArgs struct {
	// The ID of the local gateway route table.
	LocalGatewayRouteTableId pulumi.StringInput
	// The tags for the association.
	Tags LocalGatewayRouteTableVPCAssociationTagArrayInput
	// The ID of the VPC.
	VpcId pulumi.StringInput
}

func (LocalGatewayRouteTableVPCAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localGatewayRouteTableVPCAssociationArgs)(nil)).Elem()
}

type LocalGatewayRouteTableVPCAssociationInput interface {
	pulumi.Input

	ToLocalGatewayRouteTableVPCAssociationOutput() LocalGatewayRouteTableVPCAssociationOutput
	ToLocalGatewayRouteTableVPCAssociationOutputWithContext(ctx context.Context) LocalGatewayRouteTableVPCAssociationOutput
}

func (*LocalGatewayRouteTableVPCAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalGatewayRouteTableVPCAssociation)(nil)).Elem()
}

func (i *LocalGatewayRouteTableVPCAssociation) ToLocalGatewayRouteTableVPCAssociationOutput() LocalGatewayRouteTableVPCAssociationOutput {
	return i.ToLocalGatewayRouteTableVPCAssociationOutputWithContext(context.Background())
}

func (i *LocalGatewayRouteTableVPCAssociation) ToLocalGatewayRouteTableVPCAssociationOutputWithContext(ctx context.Context) LocalGatewayRouteTableVPCAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalGatewayRouteTableVPCAssociationOutput)
}

type LocalGatewayRouteTableVPCAssociationOutput struct{ *pulumi.OutputState }

func (LocalGatewayRouteTableVPCAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalGatewayRouteTableVPCAssociation)(nil)).Elem()
}

func (o LocalGatewayRouteTableVPCAssociationOutput) ToLocalGatewayRouteTableVPCAssociationOutput() LocalGatewayRouteTableVPCAssociationOutput {
	return o
}

func (o LocalGatewayRouteTableVPCAssociationOutput) ToLocalGatewayRouteTableVPCAssociationOutputWithContext(ctx context.Context) LocalGatewayRouteTableVPCAssociationOutput {
	return o
}

// The ID of the local gateway.
func (o LocalGatewayRouteTableVPCAssociationOutput) LocalGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalGatewayRouteTableVPCAssociation) pulumi.StringOutput { return v.LocalGatewayId }).(pulumi.StringOutput)
}

// The ID of the local gateway route table.
func (o LocalGatewayRouteTableVPCAssociationOutput) LocalGatewayRouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalGatewayRouteTableVPCAssociation) pulumi.StringOutput { return v.LocalGatewayRouteTableId }).(pulumi.StringOutput)
}

// The ID of the association.
func (o LocalGatewayRouteTableVPCAssociationOutput) LocalGatewayRouteTableVpcAssociationId() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalGatewayRouteTableVPCAssociation) pulumi.StringOutput {
		return v.LocalGatewayRouteTableVpcAssociationId
	}).(pulumi.StringOutput)
}

// The state of the association.
func (o LocalGatewayRouteTableVPCAssociationOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalGatewayRouteTableVPCAssociation) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The tags for the association.
func (o LocalGatewayRouteTableVPCAssociationOutput) Tags() LocalGatewayRouteTableVPCAssociationTagArrayOutput {
	return o.ApplyT(func(v *LocalGatewayRouteTableVPCAssociation) LocalGatewayRouteTableVPCAssociationTagArrayOutput {
		return v.Tags
	}).(LocalGatewayRouteTableVPCAssociationTagArrayOutput)
}

// The ID of the VPC.
func (o LocalGatewayRouteTableVPCAssociationOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalGatewayRouteTableVPCAssociation) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalGatewayRouteTableVPCAssociationInput)(nil)).Elem(), &LocalGatewayRouteTableVPCAssociation{})
	pulumi.RegisterOutputType(LocalGatewayRouteTableVPCAssociationOutput{})
}