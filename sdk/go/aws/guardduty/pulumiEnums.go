// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type DetectorCfnFeatureConfigurationStatus string

const (
	DetectorCfnFeatureConfigurationStatusEnabled  = DetectorCfnFeatureConfigurationStatus("ENABLED")
	DetectorCfnFeatureConfigurationStatusDisabled = DetectorCfnFeatureConfigurationStatus("DISABLED")
)

func (DetectorCfnFeatureConfigurationStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*DetectorCfnFeatureConfigurationStatus)(nil)).Elem()
}

func (e DetectorCfnFeatureConfigurationStatus) ToDetectorCfnFeatureConfigurationStatusOutput() DetectorCfnFeatureConfigurationStatusOutput {
	return pulumi.ToOutput(e).(DetectorCfnFeatureConfigurationStatusOutput)
}

func (e DetectorCfnFeatureConfigurationStatus) ToDetectorCfnFeatureConfigurationStatusOutputWithContext(ctx context.Context) DetectorCfnFeatureConfigurationStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DetectorCfnFeatureConfigurationStatusOutput)
}

func (e DetectorCfnFeatureConfigurationStatus) ToDetectorCfnFeatureConfigurationStatusPtrOutput() DetectorCfnFeatureConfigurationStatusPtrOutput {
	return e.ToDetectorCfnFeatureConfigurationStatusPtrOutputWithContext(context.Background())
}

func (e DetectorCfnFeatureConfigurationStatus) ToDetectorCfnFeatureConfigurationStatusPtrOutputWithContext(ctx context.Context) DetectorCfnFeatureConfigurationStatusPtrOutput {
	return DetectorCfnFeatureConfigurationStatus(e).ToDetectorCfnFeatureConfigurationStatusOutputWithContext(ctx).ToDetectorCfnFeatureConfigurationStatusPtrOutputWithContext(ctx)
}

func (e DetectorCfnFeatureConfigurationStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DetectorCfnFeatureConfigurationStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DetectorCfnFeatureConfigurationStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DetectorCfnFeatureConfigurationStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DetectorCfnFeatureConfigurationStatusOutput struct{ *pulumi.OutputState }

func (DetectorCfnFeatureConfigurationStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DetectorCfnFeatureConfigurationStatus)(nil)).Elem()
}

func (o DetectorCfnFeatureConfigurationStatusOutput) ToDetectorCfnFeatureConfigurationStatusOutput() DetectorCfnFeatureConfigurationStatusOutput {
	return o
}

func (o DetectorCfnFeatureConfigurationStatusOutput) ToDetectorCfnFeatureConfigurationStatusOutputWithContext(ctx context.Context) DetectorCfnFeatureConfigurationStatusOutput {
	return o
}

func (o DetectorCfnFeatureConfigurationStatusOutput) ToDetectorCfnFeatureConfigurationStatusPtrOutput() DetectorCfnFeatureConfigurationStatusPtrOutput {
	return o.ToDetectorCfnFeatureConfigurationStatusPtrOutputWithContext(context.Background())
}

func (o DetectorCfnFeatureConfigurationStatusOutput) ToDetectorCfnFeatureConfigurationStatusPtrOutputWithContext(ctx context.Context) DetectorCfnFeatureConfigurationStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DetectorCfnFeatureConfigurationStatus) *DetectorCfnFeatureConfigurationStatus {
		return &v
	}).(DetectorCfnFeatureConfigurationStatusPtrOutput)
}

func (o DetectorCfnFeatureConfigurationStatusOutput) ToOutput(ctx context.Context) pulumix.Output[DetectorCfnFeatureConfigurationStatus] {
	return pulumix.Output[DetectorCfnFeatureConfigurationStatus]{
		OutputState: o.OutputState,
	}
}

func (o DetectorCfnFeatureConfigurationStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DetectorCfnFeatureConfigurationStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DetectorCfnFeatureConfigurationStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DetectorCfnFeatureConfigurationStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DetectorCfnFeatureConfigurationStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DetectorCfnFeatureConfigurationStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DetectorCfnFeatureConfigurationStatusPtrOutput struct{ *pulumi.OutputState }

func (DetectorCfnFeatureConfigurationStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DetectorCfnFeatureConfigurationStatus)(nil)).Elem()
}

func (o DetectorCfnFeatureConfigurationStatusPtrOutput) ToDetectorCfnFeatureConfigurationStatusPtrOutput() DetectorCfnFeatureConfigurationStatusPtrOutput {
	return o
}

func (o DetectorCfnFeatureConfigurationStatusPtrOutput) ToDetectorCfnFeatureConfigurationStatusPtrOutputWithContext(ctx context.Context) DetectorCfnFeatureConfigurationStatusPtrOutput {
	return o
}

func (o DetectorCfnFeatureConfigurationStatusPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*DetectorCfnFeatureConfigurationStatus] {
	return pulumix.Output[*DetectorCfnFeatureConfigurationStatus]{
		OutputState: o.OutputState,
	}
}

func (o DetectorCfnFeatureConfigurationStatusPtrOutput) Elem() DetectorCfnFeatureConfigurationStatusOutput {
	return o.ApplyT(func(v *DetectorCfnFeatureConfigurationStatus) DetectorCfnFeatureConfigurationStatus {
		if v != nil {
			return *v
		}
		var ret DetectorCfnFeatureConfigurationStatus
		return ret
	}).(DetectorCfnFeatureConfigurationStatusOutput)
}

func (o DetectorCfnFeatureConfigurationStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DetectorCfnFeatureConfigurationStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DetectorCfnFeatureConfigurationStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DetectorCfnFeatureConfigurationStatusInput is an input type that accepts DetectorCfnFeatureConfigurationStatusArgs and DetectorCfnFeatureConfigurationStatusOutput values.
// You can construct a concrete instance of `DetectorCfnFeatureConfigurationStatusInput` via:
//
//	DetectorCfnFeatureConfigurationStatusArgs{...}
type DetectorCfnFeatureConfigurationStatusInput interface {
	pulumi.Input

	ToDetectorCfnFeatureConfigurationStatusOutput() DetectorCfnFeatureConfigurationStatusOutput
	ToDetectorCfnFeatureConfigurationStatusOutputWithContext(context.Context) DetectorCfnFeatureConfigurationStatusOutput
}

var detectorCfnFeatureConfigurationStatusPtrType = reflect.TypeOf((**DetectorCfnFeatureConfigurationStatus)(nil)).Elem()

type DetectorCfnFeatureConfigurationStatusPtrInput interface {
	pulumi.Input

	ToDetectorCfnFeatureConfigurationStatusPtrOutput() DetectorCfnFeatureConfigurationStatusPtrOutput
	ToDetectorCfnFeatureConfigurationStatusPtrOutputWithContext(context.Context) DetectorCfnFeatureConfigurationStatusPtrOutput
}

type detectorCfnFeatureConfigurationStatusPtr string

func DetectorCfnFeatureConfigurationStatusPtr(v string) DetectorCfnFeatureConfigurationStatusPtrInput {
	return (*detectorCfnFeatureConfigurationStatusPtr)(&v)
}

func (*detectorCfnFeatureConfigurationStatusPtr) ElementType() reflect.Type {
	return detectorCfnFeatureConfigurationStatusPtrType
}

func (in *detectorCfnFeatureConfigurationStatusPtr) ToDetectorCfnFeatureConfigurationStatusPtrOutput() DetectorCfnFeatureConfigurationStatusPtrOutput {
	return pulumi.ToOutput(in).(DetectorCfnFeatureConfigurationStatusPtrOutput)
}

func (in *detectorCfnFeatureConfigurationStatusPtr) ToDetectorCfnFeatureConfigurationStatusPtrOutputWithContext(ctx context.Context) DetectorCfnFeatureConfigurationStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DetectorCfnFeatureConfigurationStatusPtrOutput)
}

func (in *detectorCfnFeatureConfigurationStatusPtr) ToOutput(ctx context.Context) pulumix.Output[*DetectorCfnFeatureConfigurationStatus] {
	return pulumix.Output[*DetectorCfnFeatureConfigurationStatus]{
		OutputState: in.ToDetectorCfnFeatureConfigurationStatusPtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DetectorCfnFeatureConfigurationStatusInput)(nil)).Elem(), DetectorCfnFeatureConfigurationStatus("ENABLED"))
	pulumi.RegisterInputType(reflect.TypeOf((*DetectorCfnFeatureConfigurationStatusPtrInput)(nil)).Elem(), DetectorCfnFeatureConfigurationStatus("ENABLED"))
	pulumi.RegisterOutputType(DetectorCfnFeatureConfigurationStatusOutput{})
	pulumi.RegisterOutputType(DetectorCfnFeatureConfigurationStatusPtrOutput{})
}