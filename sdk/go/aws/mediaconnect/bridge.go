// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediaconnect

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::MediaConnect::Bridge
//
// Deprecated: Bridge is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Bridge struct {
	pulumi.CustomResourceState

	// The Amazon Resource Number (ARN) of the bridge.
	BridgeArn            pulumi.StringOutput                 `pulumi:"bridgeArn"`
	BridgeState          BridgeStateEnumOutput               `pulumi:"bridgeState"`
	EgressGatewayBridge  BridgeEgressGatewayBridgePtrOutput  `pulumi:"egressGatewayBridge"`
	IngressGatewayBridge BridgeIngressGatewayBridgePtrOutput `pulumi:"ingressGatewayBridge"`
	// The name of the bridge.
	Name pulumi.StringOutput `pulumi:"name"`
	// The outputs on this bridge.
	Outputs BridgeOutputResourceArrayOutput `pulumi:"outputs"`
	// The placement Amazon Resource Number (ARN) of the bridge.
	PlacementArn         pulumi.StringOutput           `pulumi:"placementArn"`
	SourceFailoverConfig BridgeFailoverConfigPtrOutput `pulumi:"sourceFailoverConfig"`
	// The sources on this bridge.
	Sources BridgeSourceTypeArrayOutput `pulumi:"sources"`
}

// NewBridge registers a new resource with the given unique name, arguments, and options.
func NewBridge(ctx *pulumi.Context,
	name string, args *BridgeArgs, opts ...pulumi.ResourceOption) (*Bridge, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PlacementArn == nil {
		return nil, errors.New("invalid value for required argument 'PlacementArn'")
	}
	if args.Sources == nil {
		return nil, errors.New("invalid value for required argument 'Sources'")
	}
	var resource Bridge
	err := ctx.RegisterResource("aws-native:mediaconnect:Bridge", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBridge gets an existing Bridge resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBridge(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BridgeState, opts ...pulumi.ResourceOption) (*Bridge, error) {
	var resource Bridge
	err := ctx.ReadResource("aws-native:mediaconnect:Bridge", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bridge resources.
type bridgeState struct {
}

type BridgeState struct {
}

func (BridgeState) ElementType() reflect.Type {
	return reflect.TypeOf((*bridgeState)(nil)).Elem()
}

type bridgeArgs struct {
	EgressGatewayBridge  *BridgeEgressGatewayBridge  `pulumi:"egressGatewayBridge"`
	IngressGatewayBridge *BridgeIngressGatewayBridge `pulumi:"ingressGatewayBridge"`
	// The name of the bridge.
	Name *string `pulumi:"name"`
	// The outputs on this bridge.
	Outputs []BridgeOutputResource `pulumi:"outputs"`
	// The placement Amazon Resource Number (ARN) of the bridge.
	PlacementArn         string                `pulumi:"placementArn"`
	SourceFailoverConfig *BridgeFailoverConfig `pulumi:"sourceFailoverConfig"`
	// The sources on this bridge.
	Sources []BridgeSourceType `pulumi:"sources"`
}

// The set of arguments for constructing a Bridge resource.
type BridgeArgs struct {
	EgressGatewayBridge  BridgeEgressGatewayBridgePtrInput
	IngressGatewayBridge BridgeIngressGatewayBridgePtrInput
	// The name of the bridge.
	Name pulumi.StringPtrInput
	// The outputs on this bridge.
	Outputs BridgeOutputResourceArrayInput
	// The placement Amazon Resource Number (ARN) of the bridge.
	PlacementArn         pulumi.StringInput
	SourceFailoverConfig BridgeFailoverConfigPtrInput
	// The sources on this bridge.
	Sources BridgeSourceTypeArrayInput
}

func (BridgeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bridgeArgs)(nil)).Elem()
}

type BridgeInput interface {
	pulumi.Input

	ToBridgeOutput() BridgeOutput
	ToBridgeOutputWithContext(ctx context.Context) BridgeOutput
}

func (*Bridge) ElementType() reflect.Type {
	return reflect.TypeOf((**Bridge)(nil)).Elem()
}

func (i *Bridge) ToBridgeOutput() BridgeOutput {
	return i.ToBridgeOutputWithContext(context.Background())
}

func (i *Bridge) ToBridgeOutputWithContext(ctx context.Context) BridgeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BridgeOutput)
}

type BridgeOutput struct{ *pulumi.OutputState }

func (BridgeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bridge)(nil)).Elem()
}

func (o BridgeOutput) ToBridgeOutput() BridgeOutput {
	return o
}

func (o BridgeOutput) ToBridgeOutputWithContext(ctx context.Context) BridgeOutput {
	return o
}

// The Amazon Resource Number (ARN) of the bridge.
func (o BridgeOutput) BridgeArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Bridge) pulumi.StringOutput { return v.BridgeArn }).(pulumi.StringOutput)
}

func (o BridgeOutput) BridgeState() BridgeStateEnumOutput {
	return o.ApplyT(func(v *Bridge) BridgeStateEnumOutput { return v.BridgeState }).(BridgeStateEnumOutput)
}

func (o BridgeOutput) EgressGatewayBridge() BridgeEgressGatewayBridgePtrOutput {
	return o.ApplyT(func(v *Bridge) BridgeEgressGatewayBridgePtrOutput { return v.EgressGatewayBridge }).(BridgeEgressGatewayBridgePtrOutput)
}

func (o BridgeOutput) IngressGatewayBridge() BridgeIngressGatewayBridgePtrOutput {
	return o.ApplyT(func(v *Bridge) BridgeIngressGatewayBridgePtrOutput { return v.IngressGatewayBridge }).(BridgeIngressGatewayBridgePtrOutput)
}

// The name of the bridge.
func (o BridgeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Bridge) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The outputs on this bridge.
func (o BridgeOutput) Outputs() BridgeOutputResourceArrayOutput {
	return o.ApplyT(func(v *Bridge) BridgeOutputResourceArrayOutput { return v.Outputs }).(BridgeOutputResourceArrayOutput)
}

// The placement Amazon Resource Number (ARN) of the bridge.
func (o BridgeOutput) PlacementArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Bridge) pulumi.StringOutput { return v.PlacementArn }).(pulumi.StringOutput)
}

func (o BridgeOutput) SourceFailoverConfig() BridgeFailoverConfigPtrOutput {
	return o.ApplyT(func(v *Bridge) BridgeFailoverConfigPtrOutput { return v.SourceFailoverConfig }).(BridgeFailoverConfigPtrOutput)
}

// The sources on this bridge.
func (o BridgeOutput) Sources() BridgeSourceTypeArrayOutput {
	return o.ApplyT(func(v *Bridge) BridgeSourceTypeArrayOutput { return v.Sources }).(BridgeSourceTypeArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BridgeInput)(nil)).Elem(), &Bridge{})
	pulumi.RegisterOutputType(BridgeOutput{})
}