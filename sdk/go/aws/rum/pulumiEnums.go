// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rum

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppMonitorTelemetry string

const (
	AppMonitorTelemetryErrors      = AppMonitorTelemetry("errors")
	AppMonitorTelemetryPerformance = AppMonitorTelemetry("performance")
	AppMonitorTelemetryHttp        = AppMonitorTelemetry("http")
)

func (AppMonitorTelemetry) ElementType() reflect.Type {
	return reflect.TypeOf((*AppMonitorTelemetry)(nil)).Elem()
}

func (e AppMonitorTelemetry) ToAppMonitorTelemetryOutput() AppMonitorTelemetryOutput {
	return pulumi.ToOutput(e).(AppMonitorTelemetryOutput)
}

func (e AppMonitorTelemetry) ToAppMonitorTelemetryOutputWithContext(ctx context.Context) AppMonitorTelemetryOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AppMonitorTelemetryOutput)
}

func (e AppMonitorTelemetry) ToAppMonitorTelemetryPtrOutput() AppMonitorTelemetryPtrOutput {
	return e.ToAppMonitorTelemetryPtrOutputWithContext(context.Background())
}

func (e AppMonitorTelemetry) ToAppMonitorTelemetryPtrOutputWithContext(ctx context.Context) AppMonitorTelemetryPtrOutput {
	return AppMonitorTelemetry(e).ToAppMonitorTelemetryOutputWithContext(ctx).ToAppMonitorTelemetryPtrOutputWithContext(ctx)
}

func (e AppMonitorTelemetry) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AppMonitorTelemetry) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AppMonitorTelemetry) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AppMonitorTelemetry) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AppMonitorTelemetryOutput struct{ *pulumi.OutputState }

func (AppMonitorTelemetryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppMonitorTelemetry)(nil)).Elem()
}

func (o AppMonitorTelemetryOutput) ToAppMonitorTelemetryOutput() AppMonitorTelemetryOutput {
	return o
}

func (o AppMonitorTelemetryOutput) ToAppMonitorTelemetryOutputWithContext(ctx context.Context) AppMonitorTelemetryOutput {
	return o
}

func (o AppMonitorTelemetryOutput) ToAppMonitorTelemetryPtrOutput() AppMonitorTelemetryPtrOutput {
	return o.ToAppMonitorTelemetryPtrOutputWithContext(context.Background())
}

func (o AppMonitorTelemetryOutput) ToAppMonitorTelemetryPtrOutputWithContext(ctx context.Context) AppMonitorTelemetryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppMonitorTelemetry) *AppMonitorTelemetry {
		return &v
	}).(AppMonitorTelemetryPtrOutput)
}

func (o AppMonitorTelemetryOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AppMonitorTelemetryOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AppMonitorTelemetry) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AppMonitorTelemetryOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AppMonitorTelemetryOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AppMonitorTelemetry) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AppMonitorTelemetryPtrOutput struct{ *pulumi.OutputState }

func (AppMonitorTelemetryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppMonitorTelemetry)(nil)).Elem()
}

func (o AppMonitorTelemetryPtrOutput) ToAppMonitorTelemetryPtrOutput() AppMonitorTelemetryPtrOutput {
	return o
}

func (o AppMonitorTelemetryPtrOutput) ToAppMonitorTelemetryPtrOutputWithContext(ctx context.Context) AppMonitorTelemetryPtrOutput {
	return o
}

func (o AppMonitorTelemetryPtrOutput) Elem() AppMonitorTelemetryOutput {
	return o.ApplyT(func(v *AppMonitorTelemetry) AppMonitorTelemetry {
		if v != nil {
			return *v
		}
		var ret AppMonitorTelemetry
		return ret
	}).(AppMonitorTelemetryOutput)
}

func (o AppMonitorTelemetryPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AppMonitorTelemetryPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AppMonitorTelemetry) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AppMonitorTelemetryInput is an input type that accepts AppMonitorTelemetryArgs and AppMonitorTelemetryOutput values.
// You can construct a concrete instance of `AppMonitorTelemetryInput` via:
//
//	AppMonitorTelemetryArgs{...}
type AppMonitorTelemetryInput interface {
	pulumi.Input

	ToAppMonitorTelemetryOutput() AppMonitorTelemetryOutput
	ToAppMonitorTelemetryOutputWithContext(context.Context) AppMonitorTelemetryOutput
}

var appMonitorTelemetryPtrType = reflect.TypeOf((**AppMonitorTelemetry)(nil)).Elem()

type AppMonitorTelemetryPtrInput interface {
	pulumi.Input

	ToAppMonitorTelemetryPtrOutput() AppMonitorTelemetryPtrOutput
	ToAppMonitorTelemetryPtrOutputWithContext(context.Context) AppMonitorTelemetryPtrOutput
}

type appMonitorTelemetryPtr string

func AppMonitorTelemetryPtr(v string) AppMonitorTelemetryPtrInput {
	return (*appMonitorTelemetryPtr)(&v)
}

func (*appMonitorTelemetryPtr) ElementType() reflect.Type {
	return appMonitorTelemetryPtrType
}

func (in *appMonitorTelemetryPtr) ToAppMonitorTelemetryPtrOutput() AppMonitorTelemetryPtrOutput {
	return pulumi.ToOutput(in).(AppMonitorTelemetryPtrOutput)
}

func (in *appMonitorTelemetryPtr) ToAppMonitorTelemetryPtrOutputWithContext(ctx context.Context) AppMonitorTelemetryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AppMonitorTelemetryPtrOutput)
}

// AppMonitorTelemetryArrayInput is an input type that accepts AppMonitorTelemetryArray and AppMonitorTelemetryArrayOutput values.
// You can construct a concrete instance of `AppMonitorTelemetryArrayInput` via:
//
//	AppMonitorTelemetryArray{ AppMonitorTelemetryArgs{...} }
type AppMonitorTelemetryArrayInput interface {
	pulumi.Input

	ToAppMonitorTelemetryArrayOutput() AppMonitorTelemetryArrayOutput
	ToAppMonitorTelemetryArrayOutputWithContext(context.Context) AppMonitorTelemetryArrayOutput
}

type AppMonitorTelemetryArray []AppMonitorTelemetry

func (AppMonitorTelemetryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppMonitorTelemetry)(nil)).Elem()
}

func (i AppMonitorTelemetryArray) ToAppMonitorTelemetryArrayOutput() AppMonitorTelemetryArrayOutput {
	return i.ToAppMonitorTelemetryArrayOutputWithContext(context.Background())
}

func (i AppMonitorTelemetryArray) ToAppMonitorTelemetryArrayOutputWithContext(ctx context.Context) AppMonitorTelemetryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppMonitorTelemetryArrayOutput)
}

type AppMonitorTelemetryArrayOutput struct{ *pulumi.OutputState }

func (AppMonitorTelemetryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppMonitorTelemetry)(nil)).Elem()
}

func (o AppMonitorTelemetryArrayOutput) ToAppMonitorTelemetryArrayOutput() AppMonitorTelemetryArrayOutput {
	return o
}

func (o AppMonitorTelemetryArrayOutput) ToAppMonitorTelemetryArrayOutputWithContext(ctx context.Context) AppMonitorTelemetryArrayOutput {
	return o
}

func (o AppMonitorTelemetryArrayOutput) Index(i pulumi.IntInput) AppMonitorTelemetryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppMonitorTelemetry {
		return vs[0].([]AppMonitorTelemetry)[vs[1].(int)]
	}).(AppMonitorTelemetryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppMonitorTelemetryInput)(nil)).Elem(), AppMonitorTelemetry("errors"))
	pulumi.RegisterInputType(reflect.TypeOf((*AppMonitorTelemetryPtrInput)(nil)).Elem(), AppMonitorTelemetry("errors"))
	pulumi.RegisterInputType(reflect.TypeOf((*AppMonitorTelemetryArrayInput)(nil)).Elem(), AppMonitorTelemetryArray{})
	pulumi.RegisterOutputType(AppMonitorTelemetryOutput{})
	pulumi.RegisterOutputType(AppMonitorTelemetryPtrOutput{})
	pulumi.RegisterOutputType(AppMonitorTelemetryArrayOutput{})
}