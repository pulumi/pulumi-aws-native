// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package controltower

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = internal.GetEnvOrDefault

type LandingZoneTag struct {
	Key   *string `pulumi:"key"`
	Value *string `pulumi:"value"`
}

// LandingZoneTagInput is an input type that accepts LandingZoneTagArgs and LandingZoneTagOutput values.
// You can construct a concrete instance of `LandingZoneTagInput` via:
//
//	LandingZoneTagArgs{...}
type LandingZoneTagInput interface {
	pulumi.Input

	ToLandingZoneTagOutput() LandingZoneTagOutput
	ToLandingZoneTagOutputWithContext(context.Context) LandingZoneTagOutput
}

type LandingZoneTagArgs struct {
	Key   pulumi.StringPtrInput `pulumi:"key"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (LandingZoneTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LandingZoneTag)(nil)).Elem()
}

func (i LandingZoneTagArgs) ToLandingZoneTagOutput() LandingZoneTagOutput {
	return i.ToLandingZoneTagOutputWithContext(context.Background())
}

func (i LandingZoneTagArgs) ToLandingZoneTagOutputWithContext(ctx context.Context) LandingZoneTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LandingZoneTagOutput)
}

func (i LandingZoneTagArgs) ToOutput(ctx context.Context) pulumix.Output[LandingZoneTag] {
	return pulumix.Output[LandingZoneTag]{
		OutputState: i.ToLandingZoneTagOutputWithContext(ctx).OutputState,
	}
}

// LandingZoneTagArrayInput is an input type that accepts LandingZoneTagArray and LandingZoneTagArrayOutput values.
// You can construct a concrete instance of `LandingZoneTagArrayInput` via:
//
//	LandingZoneTagArray{ LandingZoneTagArgs{...} }
type LandingZoneTagArrayInput interface {
	pulumi.Input

	ToLandingZoneTagArrayOutput() LandingZoneTagArrayOutput
	ToLandingZoneTagArrayOutputWithContext(context.Context) LandingZoneTagArrayOutput
}

type LandingZoneTagArray []LandingZoneTagInput

func (LandingZoneTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LandingZoneTag)(nil)).Elem()
}

func (i LandingZoneTagArray) ToLandingZoneTagArrayOutput() LandingZoneTagArrayOutput {
	return i.ToLandingZoneTagArrayOutputWithContext(context.Background())
}

func (i LandingZoneTagArray) ToLandingZoneTagArrayOutputWithContext(ctx context.Context) LandingZoneTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LandingZoneTagArrayOutput)
}

func (i LandingZoneTagArray) ToOutput(ctx context.Context) pulumix.Output[[]LandingZoneTag] {
	return pulumix.Output[[]LandingZoneTag]{
		OutputState: i.ToLandingZoneTagArrayOutputWithContext(ctx).OutputState,
	}
}

type LandingZoneTagOutput struct{ *pulumi.OutputState }

func (LandingZoneTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LandingZoneTag)(nil)).Elem()
}

func (o LandingZoneTagOutput) ToLandingZoneTagOutput() LandingZoneTagOutput {
	return o
}

func (o LandingZoneTagOutput) ToLandingZoneTagOutputWithContext(ctx context.Context) LandingZoneTagOutput {
	return o
}

func (o LandingZoneTagOutput) ToOutput(ctx context.Context) pulumix.Output[LandingZoneTag] {
	return pulumix.Output[LandingZoneTag]{
		OutputState: o.OutputState,
	}
}

func (o LandingZoneTagOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LandingZoneTag) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o LandingZoneTagOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LandingZoneTag) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type LandingZoneTagArrayOutput struct{ *pulumi.OutputState }

func (LandingZoneTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LandingZoneTag)(nil)).Elem()
}

func (o LandingZoneTagArrayOutput) ToLandingZoneTagArrayOutput() LandingZoneTagArrayOutput {
	return o
}

func (o LandingZoneTagArrayOutput) ToLandingZoneTagArrayOutputWithContext(ctx context.Context) LandingZoneTagArrayOutput {
	return o
}

func (o LandingZoneTagArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]LandingZoneTag] {
	return pulumix.Output[[]LandingZoneTag]{
		OutputState: o.OutputState,
	}
}

func (o LandingZoneTagArrayOutput) Index(i pulumi.IntInput) LandingZoneTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LandingZoneTag {
		return vs[0].([]LandingZoneTag)[vs[1].(int)]
	}).(LandingZoneTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LandingZoneTagInput)(nil)).Elem(), LandingZoneTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LandingZoneTagArrayInput)(nil)).Elem(), LandingZoneTagArray{})
	pulumi.RegisterOutputType(LandingZoneTagOutput{})
	pulumi.RegisterOutputType(LandingZoneTagArrayOutput{})
}