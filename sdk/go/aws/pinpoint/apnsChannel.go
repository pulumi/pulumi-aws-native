// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Pinpoint::APNSChannel
//
// Deprecated: ApnsChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ApnsChannel struct {
	pulumi.CustomResourceState

	ApplicationId               pulumi.StringOutput    `pulumi:"applicationId"`
	BundleId                    pulumi.StringPtrOutput `pulumi:"bundleId"`
	Certificate                 pulumi.StringPtrOutput `pulumi:"certificate"`
	DefaultAuthenticationMethod pulumi.StringPtrOutput `pulumi:"defaultAuthenticationMethod"`
	Enabled                     pulumi.BoolPtrOutput   `pulumi:"enabled"`
	PrivateKey                  pulumi.StringPtrOutput `pulumi:"privateKey"`
	TeamId                      pulumi.StringPtrOutput `pulumi:"teamId"`
	TokenKey                    pulumi.StringPtrOutput `pulumi:"tokenKey"`
	TokenKeyId                  pulumi.StringPtrOutput `pulumi:"tokenKeyId"`
}

// NewApnsChannel registers a new resource with the given unique name, arguments, and options.
func NewApnsChannel(ctx *pulumi.Context,
	name string, args *ApnsChannelArgs, opts ...pulumi.ResourceOption) (*ApnsChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApnsChannel
	err := ctx.RegisterResource("aws-native:pinpoint:ApnsChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApnsChannel gets an existing ApnsChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApnsChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApnsChannelState, opts ...pulumi.ResourceOption) (*ApnsChannel, error) {
	var resource ApnsChannel
	err := ctx.ReadResource("aws-native:pinpoint:ApnsChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApnsChannel resources.
type apnsChannelState struct {
}

type ApnsChannelState struct {
}

func (ApnsChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*apnsChannelState)(nil)).Elem()
}

type apnsChannelArgs struct {
	ApplicationId               string  `pulumi:"applicationId"`
	BundleId                    *string `pulumi:"bundleId"`
	Certificate                 *string `pulumi:"certificate"`
	DefaultAuthenticationMethod *string `pulumi:"defaultAuthenticationMethod"`
	Enabled                     *bool   `pulumi:"enabled"`
	PrivateKey                  *string `pulumi:"privateKey"`
	TeamId                      *string `pulumi:"teamId"`
	TokenKey                    *string `pulumi:"tokenKey"`
	TokenKeyId                  *string `pulumi:"tokenKeyId"`
}

// The set of arguments for constructing a ApnsChannel resource.
type ApnsChannelArgs struct {
	ApplicationId               pulumi.StringInput
	BundleId                    pulumi.StringPtrInput
	Certificate                 pulumi.StringPtrInput
	DefaultAuthenticationMethod pulumi.StringPtrInput
	Enabled                     pulumi.BoolPtrInput
	PrivateKey                  pulumi.StringPtrInput
	TeamId                      pulumi.StringPtrInput
	TokenKey                    pulumi.StringPtrInput
	TokenKeyId                  pulumi.StringPtrInput
}

func (ApnsChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apnsChannelArgs)(nil)).Elem()
}

type ApnsChannelInput interface {
	pulumi.Input

	ToApnsChannelOutput() ApnsChannelOutput
	ToApnsChannelOutputWithContext(ctx context.Context) ApnsChannelOutput
}

func (*ApnsChannel) ElementType() reflect.Type {
	return reflect.TypeOf((**ApnsChannel)(nil)).Elem()
}

func (i *ApnsChannel) ToApnsChannelOutput() ApnsChannelOutput {
	return i.ToApnsChannelOutputWithContext(context.Background())
}

func (i *ApnsChannel) ToApnsChannelOutputWithContext(ctx context.Context) ApnsChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsChannelOutput)
}

type ApnsChannelOutput struct{ *pulumi.OutputState }

func (ApnsChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApnsChannel)(nil)).Elem()
}

func (o ApnsChannelOutput) ToApnsChannelOutput() ApnsChannelOutput {
	return o
}

func (o ApnsChannelOutput) ToApnsChannelOutputWithContext(ctx context.Context) ApnsChannelOutput {
	return o
}

func (o ApnsChannelOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApnsChannel) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

func (o ApnsChannelOutput) BundleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsChannel) pulumi.StringPtrOutput { return v.BundleId }).(pulumi.StringPtrOutput)
}

func (o ApnsChannelOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsChannel) pulumi.StringPtrOutput { return v.Certificate }).(pulumi.StringPtrOutput)
}

func (o ApnsChannelOutput) DefaultAuthenticationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsChannel) pulumi.StringPtrOutput { return v.DefaultAuthenticationMethod }).(pulumi.StringPtrOutput)
}

func (o ApnsChannelOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApnsChannel) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ApnsChannelOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsChannel) pulumi.StringPtrOutput { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

func (o ApnsChannelOutput) TeamId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsChannel) pulumi.StringPtrOutput { return v.TeamId }).(pulumi.StringPtrOutput)
}

func (o ApnsChannelOutput) TokenKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsChannel) pulumi.StringPtrOutput { return v.TokenKey }).(pulumi.StringPtrOutput)
}

func (o ApnsChannelOutput) TokenKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsChannel) pulumi.StringPtrOutput { return v.TokenKeyId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApnsChannelInput)(nil)).Elem(), &ApnsChannel{})
	pulumi.RegisterOutputType(ApnsChannelOutput{})
}