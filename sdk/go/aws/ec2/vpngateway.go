// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Schema for EC2 VPN Gateway
type VPNGateway struct {
	pulumi.CustomResourceState

	// The private Autonomous System Number (ASN) for the Amazon side of a BGP session.
	AmazonSideAsn pulumi.IntPtrOutput `pulumi:"amazonSideAsn"`
	// Any tags assigned to the virtual private gateway.
	Tags VPNGatewayTagArrayOutput `pulumi:"tags"`
	// The type of VPN connection the virtual private gateway supports.
	Type pulumi.StringOutput `pulumi:"type"`
	// VPN Gateway ID generated by service
	VPNGatewayId pulumi.StringOutput `pulumi:"vPNGatewayId"`
}

// NewVPNGateway registers a new resource with the given unique name, arguments, and options.
func NewVPNGateway(ctx *pulumi.Context,
	name string, args *VPNGatewayArgs, opts ...pulumi.ResourceOption) (*VPNGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource VPNGateway
	err := ctx.RegisterResource("aws-native:ec2:VPNGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVPNGateway gets an existing VPNGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVPNGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VPNGatewayState, opts ...pulumi.ResourceOption) (*VPNGateway, error) {
	var resource VPNGateway
	err := ctx.ReadResource("aws-native:ec2:VPNGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VPNGateway resources.
type vpngatewayState struct {
}

type VPNGatewayState struct {
}

func (VPNGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpngatewayState)(nil)).Elem()
}

type vpngatewayArgs struct {
	// The private Autonomous System Number (ASN) for the Amazon side of a BGP session.
	AmazonSideAsn *int `pulumi:"amazonSideAsn"`
	// Any tags assigned to the virtual private gateway.
	Tags []VPNGatewayTag `pulumi:"tags"`
	// The type of VPN connection the virtual private gateway supports.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a VPNGateway resource.
type VPNGatewayArgs struct {
	// The private Autonomous System Number (ASN) for the Amazon side of a BGP session.
	AmazonSideAsn pulumi.IntPtrInput
	// Any tags assigned to the virtual private gateway.
	Tags VPNGatewayTagArrayInput
	// The type of VPN connection the virtual private gateway supports.
	Type pulumi.StringInput
}

func (VPNGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpngatewayArgs)(nil)).Elem()
}

type VPNGatewayInput interface {
	pulumi.Input

	ToVPNGatewayOutput() VPNGatewayOutput
	ToVPNGatewayOutputWithContext(ctx context.Context) VPNGatewayOutput
}

func (*VPNGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**VPNGateway)(nil)).Elem()
}

func (i *VPNGateway) ToVPNGatewayOutput() VPNGatewayOutput {
	return i.ToVPNGatewayOutputWithContext(context.Background())
}

func (i *VPNGateway) ToVPNGatewayOutputWithContext(ctx context.Context) VPNGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VPNGatewayOutput)
}

type VPNGatewayOutput struct{ *pulumi.OutputState }

func (VPNGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VPNGateway)(nil)).Elem()
}

func (o VPNGatewayOutput) ToVPNGatewayOutput() VPNGatewayOutput {
	return o
}

func (o VPNGatewayOutput) ToVPNGatewayOutputWithContext(ctx context.Context) VPNGatewayOutput {
	return o
}

// The private Autonomous System Number (ASN) for the Amazon side of a BGP session.
func (o VPNGatewayOutput) AmazonSideAsn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VPNGateway) pulumi.IntPtrOutput { return v.AmazonSideAsn }).(pulumi.IntPtrOutput)
}

// Any tags assigned to the virtual private gateway.
func (o VPNGatewayOutput) Tags() VPNGatewayTagArrayOutput {
	return o.ApplyT(func(v *VPNGateway) VPNGatewayTagArrayOutput { return v.Tags }).(VPNGatewayTagArrayOutput)
}

// The type of VPN connection the virtual private gateway supports.
func (o VPNGatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VPNGateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// VPN Gateway ID generated by service
func (o VPNGatewayOutput) VPNGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *VPNGateway) pulumi.StringOutput { return v.VPNGatewayId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VPNGatewayInput)(nil)).Elem(), &VPNGateway{})
	pulumi.RegisterOutputType(VPNGatewayOutput{})
}