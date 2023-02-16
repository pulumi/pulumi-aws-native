// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotwireless

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create and manage wireless gateways, including LoRa gateways.
type WirelessGateway struct {
	pulumi.CustomResourceState

	// Arn for Wireless Gateway. Returned upon successful create.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of Wireless Gateway.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The date and time when the most recent uplink was received.
	LastUplinkReceivedAt pulumi.StringPtrOutput `pulumi:"lastUplinkReceivedAt"`
	// The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Gateway.
	LoRaWAN WirelessGatewayLoRaWANGatewayOutput `pulumi:"loRaWAN"`
	// Name of Wireless Gateway.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// A list of key-value pairs that contain metadata for the gateway.
	Tags WirelessGatewayTagArrayOutput `pulumi:"tags"`
	// Thing Arn. Passed into Update to associate a Thing with the Wireless Gateway.
	ThingArn pulumi.StringPtrOutput `pulumi:"thingArn"`
	// Thing Name. If there is a Thing created, this can be returned with a Get call.
	ThingName pulumi.StringPtrOutput `pulumi:"thingName"`
}

// NewWirelessGateway registers a new resource with the given unique name, arguments, and options.
func NewWirelessGateway(ctx *pulumi.Context,
	name string, args *WirelessGatewayArgs, opts ...pulumi.ResourceOption) (*WirelessGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoRaWAN == nil {
		return nil, errors.New("invalid value for required argument 'LoRaWAN'")
	}
	var resource WirelessGateway
	err := ctx.RegisterResource("aws-native:iotwireless:WirelessGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessGateway gets an existing WirelessGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessGatewayState, opts ...pulumi.ResourceOption) (*WirelessGateway, error) {
	var resource WirelessGateway
	err := ctx.ReadResource("aws-native:iotwireless:WirelessGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessGateway resources.
type wirelessGatewayState struct {
}

type WirelessGatewayState struct {
}

func (WirelessGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessGatewayState)(nil)).Elem()
}

type wirelessGatewayArgs struct {
	// Description of Wireless Gateway.
	Description *string `pulumi:"description"`
	// The date and time when the most recent uplink was received.
	LastUplinkReceivedAt *string `pulumi:"lastUplinkReceivedAt"`
	// The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Gateway.
	LoRaWAN WirelessGatewayLoRaWANGateway `pulumi:"loRaWAN"`
	// Name of Wireless Gateway.
	Name *string `pulumi:"name"`
	// A list of key-value pairs that contain metadata for the gateway.
	Tags []WirelessGatewayTag `pulumi:"tags"`
	// Thing Arn. Passed into Update to associate a Thing with the Wireless Gateway.
	ThingArn *string `pulumi:"thingArn"`
	// Thing Name. If there is a Thing created, this can be returned with a Get call.
	ThingName *string `pulumi:"thingName"`
}

// The set of arguments for constructing a WirelessGateway resource.
type WirelessGatewayArgs struct {
	// Description of Wireless Gateway.
	Description pulumi.StringPtrInput
	// The date and time when the most recent uplink was received.
	LastUplinkReceivedAt pulumi.StringPtrInput
	// The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Gateway.
	LoRaWAN WirelessGatewayLoRaWANGatewayInput
	// Name of Wireless Gateway.
	Name pulumi.StringPtrInput
	// A list of key-value pairs that contain metadata for the gateway.
	Tags WirelessGatewayTagArrayInput
	// Thing Arn. Passed into Update to associate a Thing with the Wireless Gateway.
	ThingArn pulumi.StringPtrInput
	// Thing Name. If there is a Thing created, this can be returned with a Get call.
	ThingName pulumi.StringPtrInput
}

func (WirelessGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessGatewayArgs)(nil)).Elem()
}

type WirelessGatewayInput interface {
	pulumi.Input

	ToWirelessGatewayOutput() WirelessGatewayOutput
	ToWirelessGatewayOutputWithContext(ctx context.Context) WirelessGatewayOutput
}

func (*WirelessGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessGateway)(nil)).Elem()
}

func (i *WirelessGateway) ToWirelessGatewayOutput() WirelessGatewayOutput {
	return i.ToWirelessGatewayOutputWithContext(context.Background())
}

func (i *WirelessGateway) ToWirelessGatewayOutputWithContext(ctx context.Context) WirelessGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessGatewayOutput)
}

type WirelessGatewayOutput struct{ *pulumi.OutputState }

func (WirelessGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessGateway)(nil)).Elem()
}

func (o WirelessGatewayOutput) ToWirelessGatewayOutput() WirelessGatewayOutput {
	return o
}

func (o WirelessGatewayOutput) ToWirelessGatewayOutputWithContext(ctx context.Context) WirelessGatewayOutput {
	return o
}

// Arn for Wireless Gateway. Returned upon successful create.
func (o WirelessGatewayOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessGateway) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Description of Wireless Gateway.
func (o WirelessGatewayOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessGateway) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The date and time when the most recent uplink was received.
func (o WirelessGatewayOutput) LastUplinkReceivedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessGateway) pulumi.StringPtrOutput { return v.LastUplinkReceivedAt }).(pulumi.StringPtrOutput)
}

// The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Gateway.
func (o WirelessGatewayOutput) LoRaWAN() WirelessGatewayLoRaWANGatewayOutput {
	return o.ApplyT(func(v *WirelessGateway) WirelessGatewayLoRaWANGatewayOutput { return v.LoRaWAN }).(WirelessGatewayLoRaWANGatewayOutput)
}

// Name of Wireless Gateway.
func (o WirelessGatewayOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessGateway) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// A list of key-value pairs that contain metadata for the gateway.
func (o WirelessGatewayOutput) Tags() WirelessGatewayTagArrayOutput {
	return o.ApplyT(func(v *WirelessGateway) WirelessGatewayTagArrayOutput { return v.Tags }).(WirelessGatewayTagArrayOutput)
}

// Thing Arn. Passed into Update to associate a Thing with the Wireless Gateway.
func (o WirelessGatewayOutput) ThingArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessGateway) pulumi.StringPtrOutput { return v.ThingArn }).(pulumi.StringPtrOutput)
}

// Thing Name. If there is a Thing created, this can be returned with a Get call.
func (o WirelessGatewayOutput) ThingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessGateway) pulumi.StringPtrOutput { return v.ThingName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessGatewayInput)(nil)).Elem(), &WirelessGateway{})
	pulumi.RegisterOutputType(WirelessGatewayOutput{})
}