// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gamecast

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Type of the runtime environment used to run games. For initial launch it only includes wine staging branch but Motif
// will follow up with Proton support.
type ApplicationRuntimeEnvironmentType string

const (
	ApplicationRuntimeEnvironmentTypeWineStaging = ApplicationRuntimeEnvironmentType("WINE-STAGING")
)

func (ApplicationRuntimeEnvironmentType) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationRuntimeEnvironmentType)(nil)).Elem()
}

func (e ApplicationRuntimeEnvironmentType) ToApplicationRuntimeEnvironmentTypeOutput() ApplicationRuntimeEnvironmentTypeOutput {
	return pulumi.ToOutput(e).(ApplicationRuntimeEnvironmentTypeOutput)
}

func (e ApplicationRuntimeEnvironmentType) ToApplicationRuntimeEnvironmentTypeOutputWithContext(ctx context.Context) ApplicationRuntimeEnvironmentTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationRuntimeEnvironmentTypeOutput)
}

func (e ApplicationRuntimeEnvironmentType) ToApplicationRuntimeEnvironmentTypePtrOutput() ApplicationRuntimeEnvironmentTypePtrOutput {
	return e.ToApplicationRuntimeEnvironmentTypePtrOutputWithContext(context.Background())
}

func (e ApplicationRuntimeEnvironmentType) ToApplicationRuntimeEnvironmentTypePtrOutputWithContext(ctx context.Context) ApplicationRuntimeEnvironmentTypePtrOutput {
	return ApplicationRuntimeEnvironmentType(e).ToApplicationRuntimeEnvironmentTypeOutputWithContext(ctx).ToApplicationRuntimeEnvironmentTypePtrOutputWithContext(ctx)
}

func (e ApplicationRuntimeEnvironmentType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationRuntimeEnvironmentType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationRuntimeEnvironmentType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationRuntimeEnvironmentType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApplicationRuntimeEnvironmentTypeOutput struct{ *pulumi.OutputState }

func (ApplicationRuntimeEnvironmentTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationRuntimeEnvironmentType)(nil)).Elem()
}

func (o ApplicationRuntimeEnvironmentTypeOutput) ToApplicationRuntimeEnvironmentTypeOutput() ApplicationRuntimeEnvironmentTypeOutput {
	return o
}

func (o ApplicationRuntimeEnvironmentTypeOutput) ToApplicationRuntimeEnvironmentTypeOutputWithContext(ctx context.Context) ApplicationRuntimeEnvironmentTypeOutput {
	return o
}

func (o ApplicationRuntimeEnvironmentTypeOutput) ToApplicationRuntimeEnvironmentTypePtrOutput() ApplicationRuntimeEnvironmentTypePtrOutput {
	return o.ToApplicationRuntimeEnvironmentTypePtrOutputWithContext(context.Background())
}

func (o ApplicationRuntimeEnvironmentTypeOutput) ToApplicationRuntimeEnvironmentTypePtrOutputWithContext(ctx context.Context) ApplicationRuntimeEnvironmentTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationRuntimeEnvironmentType) *ApplicationRuntimeEnvironmentType {
		return &v
	}).(ApplicationRuntimeEnvironmentTypePtrOutput)
}

func (o ApplicationRuntimeEnvironmentTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationRuntimeEnvironmentTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationRuntimeEnvironmentType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationRuntimeEnvironmentTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationRuntimeEnvironmentTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationRuntimeEnvironmentType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationRuntimeEnvironmentTypePtrOutput struct{ *pulumi.OutputState }

func (ApplicationRuntimeEnvironmentTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationRuntimeEnvironmentType)(nil)).Elem()
}

func (o ApplicationRuntimeEnvironmentTypePtrOutput) ToApplicationRuntimeEnvironmentTypePtrOutput() ApplicationRuntimeEnvironmentTypePtrOutput {
	return o
}

func (o ApplicationRuntimeEnvironmentTypePtrOutput) ToApplicationRuntimeEnvironmentTypePtrOutputWithContext(ctx context.Context) ApplicationRuntimeEnvironmentTypePtrOutput {
	return o
}

func (o ApplicationRuntimeEnvironmentTypePtrOutput) Elem() ApplicationRuntimeEnvironmentTypeOutput {
	return o.ApplyT(func(v *ApplicationRuntimeEnvironmentType) ApplicationRuntimeEnvironmentType {
		if v != nil {
			return *v
		}
		var ret ApplicationRuntimeEnvironmentType
		return ret
	}).(ApplicationRuntimeEnvironmentTypeOutput)
}

func (o ApplicationRuntimeEnvironmentTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationRuntimeEnvironmentTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationRuntimeEnvironmentType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ApplicationRuntimeEnvironmentTypeInput is an input type that accepts ApplicationRuntimeEnvironmentTypeArgs and ApplicationRuntimeEnvironmentTypeOutput values.
// You can construct a concrete instance of `ApplicationRuntimeEnvironmentTypeInput` via:
//
//	ApplicationRuntimeEnvironmentTypeArgs{...}
type ApplicationRuntimeEnvironmentTypeInput interface {
	pulumi.Input

	ToApplicationRuntimeEnvironmentTypeOutput() ApplicationRuntimeEnvironmentTypeOutput
	ToApplicationRuntimeEnvironmentTypeOutputWithContext(context.Context) ApplicationRuntimeEnvironmentTypeOutput
}

var applicationRuntimeEnvironmentTypePtrType = reflect.TypeOf((**ApplicationRuntimeEnvironmentType)(nil)).Elem()

type ApplicationRuntimeEnvironmentTypePtrInput interface {
	pulumi.Input

	ToApplicationRuntimeEnvironmentTypePtrOutput() ApplicationRuntimeEnvironmentTypePtrOutput
	ToApplicationRuntimeEnvironmentTypePtrOutputWithContext(context.Context) ApplicationRuntimeEnvironmentTypePtrOutput
}

type applicationRuntimeEnvironmentTypePtr string

func ApplicationRuntimeEnvironmentTypePtr(v string) ApplicationRuntimeEnvironmentTypePtrInput {
	return (*applicationRuntimeEnvironmentTypePtr)(&v)
}

func (*applicationRuntimeEnvironmentTypePtr) ElementType() reflect.Type {
	return applicationRuntimeEnvironmentTypePtrType
}

func (in *applicationRuntimeEnvironmentTypePtr) ToApplicationRuntimeEnvironmentTypePtrOutput() ApplicationRuntimeEnvironmentTypePtrOutput {
	return pulumi.ToOutput(in).(ApplicationRuntimeEnvironmentTypePtrOutput)
}

func (in *applicationRuntimeEnvironmentTypePtr) ToApplicationRuntimeEnvironmentTypePtrOutputWithContext(ctx context.Context) ApplicationRuntimeEnvironmentTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationRuntimeEnvironmentTypePtrOutput)
}

// These are pre-packed Motif virtual instance type used to run customer games where mini spec maps to high-end
// integrated graphics and ultra maps to high-end gaming pc. These have different cost implications.
type StreamGroupStreamClass string

const (
	StreamGroupStreamClassMini  = StreamGroupStreamClass("MINI")
	StreamGroupStreamClassLow   = StreamGroupStreamClass("LOW")
	StreamGroupStreamClassMid   = StreamGroupStreamClass("MID")
	StreamGroupStreamClassHigh  = StreamGroupStreamClass("HIGH")
	StreamGroupStreamClassUltra = StreamGroupStreamClass("ULTRA")
)

func (StreamGroupStreamClass) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamGroupStreamClass)(nil)).Elem()
}

func (e StreamGroupStreamClass) ToStreamGroupStreamClassOutput() StreamGroupStreamClassOutput {
	return pulumi.ToOutput(e).(StreamGroupStreamClassOutput)
}

func (e StreamGroupStreamClass) ToStreamGroupStreamClassOutputWithContext(ctx context.Context) StreamGroupStreamClassOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StreamGroupStreamClassOutput)
}

func (e StreamGroupStreamClass) ToStreamGroupStreamClassPtrOutput() StreamGroupStreamClassPtrOutput {
	return e.ToStreamGroupStreamClassPtrOutputWithContext(context.Background())
}

func (e StreamGroupStreamClass) ToStreamGroupStreamClassPtrOutputWithContext(ctx context.Context) StreamGroupStreamClassPtrOutput {
	return StreamGroupStreamClass(e).ToStreamGroupStreamClassOutputWithContext(ctx).ToStreamGroupStreamClassPtrOutputWithContext(ctx)
}

func (e StreamGroupStreamClass) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StreamGroupStreamClass) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StreamGroupStreamClass) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StreamGroupStreamClass) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StreamGroupStreamClassOutput struct{ *pulumi.OutputState }

func (StreamGroupStreamClassOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamGroupStreamClass)(nil)).Elem()
}

func (o StreamGroupStreamClassOutput) ToStreamGroupStreamClassOutput() StreamGroupStreamClassOutput {
	return o
}

func (o StreamGroupStreamClassOutput) ToStreamGroupStreamClassOutputWithContext(ctx context.Context) StreamGroupStreamClassOutput {
	return o
}

func (o StreamGroupStreamClassOutput) ToStreamGroupStreamClassPtrOutput() StreamGroupStreamClassPtrOutput {
	return o.ToStreamGroupStreamClassPtrOutputWithContext(context.Background())
}

func (o StreamGroupStreamClassOutput) ToStreamGroupStreamClassPtrOutputWithContext(ctx context.Context) StreamGroupStreamClassPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamGroupStreamClass) *StreamGroupStreamClass {
		return &v
	}).(StreamGroupStreamClassPtrOutput)
}

func (o StreamGroupStreamClassOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StreamGroupStreamClassOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StreamGroupStreamClass) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StreamGroupStreamClassOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StreamGroupStreamClassOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StreamGroupStreamClass) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StreamGroupStreamClassPtrOutput struct{ *pulumi.OutputState }

func (StreamGroupStreamClassPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamGroupStreamClass)(nil)).Elem()
}

func (o StreamGroupStreamClassPtrOutput) ToStreamGroupStreamClassPtrOutput() StreamGroupStreamClassPtrOutput {
	return o
}

func (o StreamGroupStreamClassPtrOutput) ToStreamGroupStreamClassPtrOutputWithContext(ctx context.Context) StreamGroupStreamClassPtrOutput {
	return o
}

func (o StreamGroupStreamClassPtrOutput) Elem() StreamGroupStreamClassOutput {
	return o.ApplyT(func(v *StreamGroupStreamClass) StreamGroupStreamClass {
		if v != nil {
			return *v
		}
		var ret StreamGroupStreamClass
		return ret
	}).(StreamGroupStreamClassOutput)
}

func (o StreamGroupStreamClassPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StreamGroupStreamClassPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StreamGroupStreamClass) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// StreamGroupStreamClassInput is an input type that accepts StreamGroupStreamClassArgs and StreamGroupStreamClassOutput values.
// You can construct a concrete instance of `StreamGroupStreamClassInput` via:
//
//	StreamGroupStreamClassArgs{...}
type StreamGroupStreamClassInput interface {
	pulumi.Input

	ToStreamGroupStreamClassOutput() StreamGroupStreamClassOutput
	ToStreamGroupStreamClassOutputWithContext(context.Context) StreamGroupStreamClassOutput
}

var streamGroupStreamClassPtrType = reflect.TypeOf((**StreamGroupStreamClass)(nil)).Elem()

type StreamGroupStreamClassPtrInput interface {
	pulumi.Input

	ToStreamGroupStreamClassPtrOutput() StreamGroupStreamClassPtrOutput
	ToStreamGroupStreamClassPtrOutputWithContext(context.Context) StreamGroupStreamClassPtrOutput
}

type streamGroupStreamClassPtr string

func StreamGroupStreamClassPtr(v string) StreamGroupStreamClassPtrInput {
	return (*streamGroupStreamClassPtr)(&v)
}

func (*streamGroupStreamClassPtr) ElementType() reflect.Type {
	return streamGroupStreamClassPtrType
}

func (in *streamGroupStreamClassPtr) ToStreamGroupStreamClassPtrOutput() StreamGroupStreamClassPtrOutput {
	return pulumi.ToOutput(in).(StreamGroupStreamClassPtrOutput)
}

func (in *streamGroupStreamClassPtr) ToStreamGroupStreamClassPtrOutputWithContext(ctx context.Context) StreamGroupStreamClassPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StreamGroupStreamClassPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationRuntimeEnvironmentTypeInput)(nil)).Elem(), ApplicationRuntimeEnvironmentType("WINE-STAGING"))
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationRuntimeEnvironmentTypePtrInput)(nil)).Elem(), ApplicationRuntimeEnvironmentType("WINE-STAGING"))
	pulumi.RegisterInputType(reflect.TypeOf((*StreamGroupStreamClassInput)(nil)).Elem(), StreamGroupStreamClass("MINI"))
	pulumi.RegisterInputType(reflect.TypeOf((*StreamGroupStreamClassPtrInput)(nil)).Elem(), StreamGroupStreamClass("MINI"))
	pulumi.RegisterOutputType(ApplicationRuntimeEnvironmentTypeOutput{})
	pulumi.RegisterOutputType(ApplicationRuntimeEnvironmentTypePtrOutput{})
	pulumi.RegisterOutputType(StreamGroupStreamClassOutput{})
	pulumi.RegisterOutputType(StreamGroupStreamClassPtrOutput{})
}