// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package controltower

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Definition of AWS::ControlTower::LandingZone Resource Type
type LandingZone struct {
	pulumi.CustomResourceState

	Arn                    pulumi.StringOutput          `pulumi:"arn"`
	DriftStatus            LandingZoneDriftStatusOutput `pulumi:"driftStatus"`
	LandingZoneIdentifier  pulumi.StringOutput          `pulumi:"landingZoneIdentifier"`
	LatestAvailableVersion pulumi.StringOutput          `pulumi:"latestAvailableVersion"`
	Manifest               pulumi.AnyOutput             `pulumi:"manifest"`
	Status                 LandingZoneStatusOutput      `pulumi:"status"`
	Tags                   LandingZoneTagArrayOutput    `pulumi:"tags"`
	Version                pulumi.StringOutput          `pulumi:"version"`
}

// NewLandingZone registers a new resource with the given unique name, arguments, and options.
func NewLandingZone(ctx *pulumi.Context,
	name string, args *LandingZoneArgs, opts ...pulumi.ResourceOption) (*LandingZone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Manifest == nil {
		return nil, errors.New("invalid value for required argument 'Manifest'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LandingZone
	err := ctx.RegisterResource("aws-native:controltower:LandingZone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLandingZone gets an existing LandingZone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLandingZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LandingZoneState, opts ...pulumi.ResourceOption) (*LandingZone, error) {
	var resource LandingZone
	err := ctx.ReadResource("aws-native:controltower:LandingZone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LandingZone resources.
type landingZoneState struct {
}

type LandingZoneState struct {
}

func (LandingZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*landingZoneState)(nil)).Elem()
}

type landingZoneArgs struct {
	Manifest interface{}      `pulumi:"manifest"`
	Tags     []LandingZoneTag `pulumi:"tags"`
	Version  string           `pulumi:"version"`
}

// The set of arguments for constructing a LandingZone resource.
type LandingZoneArgs struct {
	Manifest pulumi.Input
	Tags     LandingZoneTagArrayInput
	Version  pulumi.StringInput
}

func (LandingZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*landingZoneArgs)(nil)).Elem()
}

type LandingZoneInput interface {
	pulumi.Input

	ToLandingZoneOutput() LandingZoneOutput
	ToLandingZoneOutputWithContext(ctx context.Context) LandingZoneOutput
}

func (*LandingZone) ElementType() reflect.Type {
	return reflect.TypeOf((**LandingZone)(nil)).Elem()
}

func (i *LandingZone) ToLandingZoneOutput() LandingZoneOutput {
	return i.ToLandingZoneOutputWithContext(context.Background())
}

func (i *LandingZone) ToLandingZoneOutputWithContext(ctx context.Context) LandingZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LandingZoneOutput)
}

func (i *LandingZone) ToOutput(ctx context.Context) pulumix.Output[*LandingZone] {
	return pulumix.Output[*LandingZone]{
		OutputState: i.ToLandingZoneOutputWithContext(ctx).OutputState,
	}
}

type LandingZoneOutput struct{ *pulumi.OutputState }

func (LandingZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LandingZone)(nil)).Elem()
}

func (o LandingZoneOutput) ToLandingZoneOutput() LandingZoneOutput {
	return o
}

func (o LandingZoneOutput) ToLandingZoneOutputWithContext(ctx context.Context) LandingZoneOutput {
	return o
}

func (o LandingZoneOutput) ToOutput(ctx context.Context) pulumix.Output[*LandingZone] {
	return pulumix.Output[*LandingZone]{
		OutputState: o.OutputState,
	}
}

func (o LandingZoneOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *LandingZone) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o LandingZoneOutput) DriftStatus() LandingZoneDriftStatusOutput {
	return o.ApplyT(func(v *LandingZone) LandingZoneDriftStatusOutput { return v.DriftStatus }).(LandingZoneDriftStatusOutput)
}

func (o LandingZoneOutput) LandingZoneIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *LandingZone) pulumi.StringOutput { return v.LandingZoneIdentifier }).(pulumi.StringOutput)
}

func (o LandingZoneOutput) LatestAvailableVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *LandingZone) pulumi.StringOutput { return v.LatestAvailableVersion }).(pulumi.StringOutput)
}

func (o LandingZoneOutput) Manifest() pulumi.AnyOutput {
	return o.ApplyT(func(v *LandingZone) pulumi.AnyOutput { return v.Manifest }).(pulumi.AnyOutput)
}

func (o LandingZoneOutput) Status() LandingZoneStatusOutput {
	return o.ApplyT(func(v *LandingZone) LandingZoneStatusOutput { return v.Status }).(LandingZoneStatusOutput)
}

func (o LandingZoneOutput) Tags() LandingZoneTagArrayOutput {
	return o.ApplyT(func(v *LandingZone) LandingZoneTagArrayOutput { return v.Tags }).(LandingZoneTagArrayOutput)
}

func (o LandingZoneOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *LandingZone) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LandingZoneInput)(nil)).Elem(), &LandingZone{})
	pulumi.RegisterOutputType(LandingZoneOutput{})
}