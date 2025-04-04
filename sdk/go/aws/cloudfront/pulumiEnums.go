// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The type of traffic configuration.
type ContinuousDeploymentPolicyConfigType string

const (
	ContinuousDeploymentPolicyConfigTypeSingleWeight = ContinuousDeploymentPolicyConfigType("SingleWeight")
	ContinuousDeploymentPolicyConfigTypeSingleHeader = ContinuousDeploymentPolicyConfigType("SingleHeader")
)

func (ContinuousDeploymentPolicyConfigType) ElementType() reflect.Type {
	return reflect.TypeOf((*ContinuousDeploymentPolicyConfigType)(nil)).Elem()
}

func (e ContinuousDeploymentPolicyConfigType) ToContinuousDeploymentPolicyConfigTypeOutput() ContinuousDeploymentPolicyConfigTypeOutput {
	return pulumi.ToOutput(e).(ContinuousDeploymentPolicyConfigTypeOutput)
}

func (e ContinuousDeploymentPolicyConfigType) ToContinuousDeploymentPolicyConfigTypeOutputWithContext(ctx context.Context) ContinuousDeploymentPolicyConfigTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContinuousDeploymentPolicyConfigTypeOutput)
}

func (e ContinuousDeploymentPolicyConfigType) ToContinuousDeploymentPolicyConfigTypePtrOutput() ContinuousDeploymentPolicyConfigTypePtrOutput {
	return e.ToContinuousDeploymentPolicyConfigTypePtrOutputWithContext(context.Background())
}

func (e ContinuousDeploymentPolicyConfigType) ToContinuousDeploymentPolicyConfigTypePtrOutputWithContext(ctx context.Context) ContinuousDeploymentPolicyConfigTypePtrOutput {
	return ContinuousDeploymentPolicyConfigType(e).ToContinuousDeploymentPolicyConfigTypeOutputWithContext(ctx).ToContinuousDeploymentPolicyConfigTypePtrOutputWithContext(ctx)
}

func (e ContinuousDeploymentPolicyConfigType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContinuousDeploymentPolicyConfigType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContinuousDeploymentPolicyConfigType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContinuousDeploymentPolicyConfigType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContinuousDeploymentPolicyConfigTypeOutput struct{ *pulumi.OutputState }

func (ContinuousDeploymentPolicyConfigTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContinuousDeploymentPolicyConfigType)(nil)).Elem()
}

func (o ContinuousDeploymentPolicyConfigTypeOutput) ToContinuousDeploymentPolicyConfigTypeOutput() ContinuousDeploymentPolicyConfigTypeOutput {
	return o
}

func (o ContinuousDeploymentPolicyConfigTypeOutput) ToContinuousDeploymentPolicyConfigTypeOutputWithContext(ctx context.Context) ContinuousDeploymentPolicyConfigTypeOutput {
	return o
}

func (o ContinuousDeploymentPolicyConfigTypeOutput) ToContinuousDeploymentPolicyConfigTypePtrOutput() ContinuousDeploymentPolicyConfigTypePtrOutput {
	return o.ToContinuousDeploymentPolicyConfigTypePtrOutputWithContext(context.Background())
}

func (o ContinuousDeploymentPolicyConfigTypeOutput) ToContinuousDeploymentPolicyConfigTypePtrOutputWithContext(ctx context.Context) ContinuousDeploymentPolicyConfigTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContinuousDeploymentPolicyConfigType) *ContinuousDeploymentPolicyConfigType {
		return &v
	}).(ContinuousDeploymentPolicyConfigTypePtrOutput)
}

func (o ContinuousDeploymentPolicyConfigTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContinuousDeploymentPolicyConfigTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContinuousDeploymentPolicyConfigType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContinuousDeploymentPolicyConfigTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContinuousDeploymentPolicyConfigTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContinuousDeploymentPolicyConfigType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContinuousDeploymentPolicyConfigTypePtrOutput struct{ *pulumi.OutputState }

func (ContinuousDeploymentPolicyConfigTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContinuousDeploymentPolicyConfigType)(nil)).Elem()
}

func (o ContinuousDeploymentPolicyConfigTypePtrOutput) ToContinuousDeploymentPolicyConfigTypePtrOutput() ContinuousDeploymentPolicyConfigTypePtrOutput {
	return o
}

func (o ContinuousDeploymentPolicyConfigTypePtrOutput) ToContinuousDeploymentPolicyConfigTypePtrOutputWithContext(ctx context.Context) ContinuousDeploymentPolicyConfigTypePtrOutput {
	return o
}

func (o ContinuousDeploymentPolicyConfigTypePtrOutput) Elem() ContinuousDeploymentPolicyConfigTypeOutput {
	return o.ApplyT(func(v *ContinuousDeploymentPolicyConfigType) ContinuousDeploymentPolicyConfigType {
		if v != nil {
			return *v
		}
		var ret ContinuousDeploymentPolicyConfigType
		return ret
	}).(ContinuousDeploymentPolicyConfigTypeOutput)
}

func (o ContinuousDeploymentPolicyConfigTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContinuousDeploymentPolicyConfigTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContinuousDeploymentPolicyConfigType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ContinuousDeploymentPolicyConfigTypeInput is an input type that accepts values of the ContinuousDeploymentPolicyConfigType enum
// A concrete instance of `ContinuousDeploymentPolicyConfigTypeInput` can be one of the following:
//
//	ContinuousDeploymentPolicyConfigTypeSingleWeight
//	ContinuousDeploymentPolicyConfigTypeSingleHeader
type ContinuousDeploymentPolicyConfigTypeInput interface {
	pulumi.Input

	ToContinuousDeploymentPolicyConfigTypeOutput() ContinuousDeploymentPolicyConfigTypeOutput
	ToContinuousDeploymentPolicyConfigTypeOutputWithContext(context.Context) ContinuousDeploymentPolicyConfigTypeOutput
}

var continuousDeploymentPolicyConfigTypePtrType = reflect.TypeOf((**ContinuousDeploymentPolicyConfigType)(nil)).Elem()

type ContinuousDeploymentPolicyConfigTypePtrInput interface {
	pulumi.Input

	ToContinuousDeploymentPolicyConfigTypePtrOutput() ContinuousDeploymentPolicyConfigTypePtrOutput
	ToContinuousDeploymentPolicyConfigTypePtrOutputWithContext(context.Context) ContinuousDeploymentPolicyConfigTypePtrOutput
}

type continuousDeploymentPolicyConfigTypePtr string

func ContinuousDeploymentPolicyConfigTypePtr(v string) ContinuousDeploymentPolicyConfigTypePtrInput {
	return (*continuousDeploymentPolicyConfigTypePtr)(&v)
}

func (*continuousDeploymentPolicyConfigTypePtr) ElementType() reflect.Type {
	return continuousDeploymentPolicyConfigTypePtrType
}

func (in *continuousDeploymentPolicyConfigTypePtr) ToContinuousDeploymentPolicyConfigTypePtrOutput() ContinuousDeploymentPolicyConfigTypePtrOutput {
	return pulumi.ToOutput(in).(ContinuousDeploymentPolicyConfigTypePtrOutput)
}

func (in *continuousDeploymentPolicyConfigTypePtr) ToContinuousDeploymentPolicyConfigTypePtrOutputWithContext(ctx context.Context) ContinuousDeploymentPolicyConfigTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContinuousDeploymentPolicyConfigTypePtrOutput)
}

// The type of traffic configuration.
type ContinuousDeploymentPolicyTrafficConfigType string

const (
	ContinuousDeploymentPolicyTrafficConfigTypeSingleWeight = ContinuousDeploymentPolicyTrafficConfigType("SingleWeight")
	ContinuousDeploymentPolicyTrafficConfigTypeSingleHeader = ContinuousDeploymentPolicyTrafficConfigType("SingleHeader")
)

func (ContinuousDeploymentPolicyTrafficConfigType) ElementType() reflect.Type {
	return reflect.TypeOf((*ContinuousDeploymentPolicyTrafficConfigType)(nil)).Elem()
}

func (e ContinuousDeploymentPolicyTrafficConfigType) ToContinuousDeploymentPolicyTrafficConfigTypeOutput() ContinuousDeploymentPolicyTrafficConfigTypeOutput {
	return pulumi.ToOutput(e).(ContinuousDeploymentPolicyTrafficConfigTypeOutput)
}

func (e ContinuousDeploymentPolicyTrafficConfigType) ToContinuousDeploymentPolicyTrafficConfigTypeOutputWithContext(ctx context.Context) ContinuousDeploymentPolicyTrafficConfigTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContinuousDeploymentPolicyTrafficConfigTypeOutput)
}

func (e ContinuousDeploymentPolicyTrafficConfigType) ToContinuousDeploymentPolicyTrafficConfigTypePtrOutput() ContinuousDeploymentPolicyTrafficConfigTypePtrOutput {
	return e.ToContinuousDeploymentPolicyTrafficConfigTypePtrOutputWithContext(context.Background())
}

func (e ContinuousDeploymentPolicyTrafficConfigType) ToContinuousDeploymentPolicyTrafficConfigTypePtrOutputWithContext(ctx context.Context) ContinuousDeploymentPolicyTrafficConfigTypePtrOutput {
	return ContinuousDeploymentPolicyTrafficConfigType(e).ToContinuousDeploymentPolicyTrafficConfigTypeOutputWithContext(ctx).ToContinuousDeploymentPolicyTrafficConfigTypePtrOutputWithContext(ctx)
}

func (e ContinuousDeploymentPolicyTrafficConfigType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContinuousDeploymentPolicyTrafficConfigType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContinuousDeploymentPolicyTrafficConfigType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContinuousDeploymentPolicyTrafficConfigType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContinuousDeploymentPolicyTrafficConfigTypeOutput struct{ *pulumi.OutputState }

func (ContinuousDeploymentPolicyTrafficConfigTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContinuousDeploymentPolicyTrafficConfigType)(nil)).Elem()
}

func (o ContinuousDeploymentPolicyTrafficConfigTypeOutput) ToContinuousDeploymentPolicyTrafficConfigTypeOutput() ContinuousDeploymentPolicyTrafficConfigTypeOutput {
	return o
}

func (o ContinuousDeploymentPolicyTrafficConfigTypeOutput) ToContinuousDeploymentPolicyTrafficConfigTypeOutputWithContext(ctx context.Context) ContinuousDeploymentPolicyTrafficConfigTypeOutput {
	return o
}

func (o ContinuousDeploymentPolicyTrafficConfigTypeOutput) ToContinuousDeploymentPolicyTrafficConfigTypePtrOutput() ContinuousDeploymentPolicyTrafficConfigTypePtrOutput {
	return o.ToContinuousDeploymentPolicyTrafficConfigTypePtrOutputWithContext(context.Background())
}

func (o ContinuousDeploymentPolicyTrafficConfigTypeOutput) ToContinuousDeploymentPolicyTrafficConfigTypePtrOutputWithContext(ctx context.Context) ContinuousDeploymentPolicyTrafficConfigTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContinuousDeploymentPolicyTrafficConfigType) *ContinuousDeploymentPolicyTrafficConfigType {
		return &v
	}).(ContinuousDeploymentPolicyTrafficConfigTypePtrOutput)
}

func (o ContinuousDeploymentPolicyTrafficConfigTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContinuousDeploymentPolicyTrafficConfigTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContinuousDeploymentPolicyTrafficConfigType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContinuousDeploymentPolicyTrafficConfigTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContinuousDeploymentPolicyTrafficConfigTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContinuousDeploymentPolicyTrafficConfigType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContinuousDeploymentPolicyTrafficConfigTypePtrOutput struct{ *pulumi.OutputState }

func (ContinuousDeploymentPolicyTrafficConfigTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContinuousDeploymentPolicyTrafficConfigType)(nil)).Elem()
}

func (o ContinuousDeploymentPolicyTrafficConfigTypePtrOutput) ToContinuousDeploymentPolicyTrafficConfigTypePtrOutput() ContinuousDeploymentPolicyTrafficConfigTypePtrOutput {
	return o
}

func (o ContinuousDeploymentPolicyTrafficConfigTypePtrOutput) ToContinuousDeploymentPolicyTrafficConfigTypePtrOutputWithContext(ctx context.Context) ContinuousDeploymentPolicyTrafficConfigTypePtrOutput {
	return o
}

func (o ContinuousDeploymentPolicyTrafficConfigTypePtrOutput) Elem() ContinuousDeploymentPolicyTrafficConfigTypeOutput {
	return o.ApplyT(func(v *ContinuousDeploymentPolicyTrafficConfigType) ContinuousDeploymentPolicyTrafficConfigType {
		if v != nil {
			return *v
		}
		var ret ContinuousDeploymentPolicyTrafficConfigType
		return ret
	}).(ContinuousDeploymentPolicyTrafficConfigTypeOutput)
}

func (o ContinuousDeploymentPolicyTrafficConfigTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContinuousDeploymentPolicyTrafficConfigTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContinuousDeploymentPolicyTrafficConfigType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ContinuousDeploymentPolicyTrafficConfigTypeInput is an input type that accepts values of the ContinuousDeploymentPolicyTrafficConfigType enum
// A concrete instance of `ContinuousDeploymentPolicyTrafficConfigTypeInput` can be one of the following:
//
//	ContinuousDeploymentPolicyTrafficConfigTypeSingleWeight
//	ContinuousDeploymentPolicyTrafficConfigTypeSingleHeader
type ContinuousDeploymentPolicyTrafficConfigTypeInput interface {
	pulumi.Input

	ToContinuousDeploymentPolicyTrafficConfigTypeOutput() ContinuousDeploymentPolicyTrafficConfigTypeOutput
	ToContinuousDeploymentPolicyTrafficConfigTypeOutputWithContext(context.Context) ContinuousDeploymentPolicyTrafficConfigTypeOutput
}

var continuousDeploymentPolicyTrafficConfigTypePtrType = reflect.TypeOf((**ContinuousDeploymentPolicyTrafficConfigType)(nil)).Elem()

type ContinuousDeploymentPolicyTrafficConfigTypePtrInput interface {
	pulumi.Input

	ToContinuousDeploymentPolicyTrafficConfigTypePtrOutput() ContinuousDeploymentPolicyTrafficConfigTypePtrOutput
	ToContinuousDeploymentPolicyTrafficConfigTypePtrOutputWithContext(context.Context) ContinuousDeploymentPolicyTrafficConfigTypePtrOutput
}

type continuousDeploymentPolicyTrafficConfigTypePtr string

func ContinuousDeploymentPolicyTrafficConfigTypePtr(v string) ContinuousDeploymentPolicyTrafficConfigTypePtrInput {
	return (*continuousDeploymentPolicyTrafficConfigTypePtr)(&v)
}

func (*continuousDeploymentPolicyTrafficConfigTypePtr) ElementType() reflect.Type {
	return continuousDeploymentPolicyTrafficConfigTypePtrType
}

func (in *continuousDeploymentPolicyTrafficConfigTypePtr) ToContinuousDeploymentPolicyTrafficConfigTypePtrOutput() ContinuousDeploymentPolicyTrafficConfigTypePtrOutput {
	return pulumi.ToOutput(in).(ContinuousDeploymentPolicyTrafficConfigTypePtrOutput)
}

func (in *continuousDeploymentPolicyTrafficConfigTypePtr) ToContinuousDeploymentPolicyTrafficConfigTypePtrOutputWithContext(ctx context.Context) ContinuousDeploymentPolicyTrafficConfigTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContinuousDeploymentPolicyTrafficConfigTypePtrOutput)
}

type DistributionOriginGroupSelectionCriteria string

const (
	DistributionOriginGroupSelectionCriteriaDefault           = DistributionOriginGroupSelectionCriteria("default")
	DistributionOriginGroupSelectionCriteriaMediaQualityBased = DistributionOriginGroupSelectionCriteria("media-quality-based")
)

func (DistributionOriginGroupSelectionCriteria) ElementType() reflect.Type {
	return reflect.TypeOf((*DistributionOriginGroupSelectionCriteria)(nil)).Elem()
}

func (e DistributionOriginGroupSelectionCriteria) ToDistributionOriginGroupSelectionCriteriaOutput() DistributionOriginGroupSelectionCriteriaOutput {
	return pulumi.ToOutput(e).(DistributionOriginGroupSelectionCriteriaOutput)
}

func (e DistributionOriginGroupSelectionCriteria) ToDistributionOriginGroupSelectionCriteriaOutputWithContext(ctx context.Context) DistributionOriginGroupSelectionCriteriaOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DistributionOriginGroupSelectionCriteriaOutput)
}

func (e DistributionOriginGroupSelectionCriteria) ToDistributionOriginGroupSelectionCriteriaPtrOutput() DistributionOriginGroupSelectionCriteriaPtrOutput {
	return e.ToDistributionOriginGroupSelectionCriteriaPtrOutputWithContext(context.Background())
}

func (e DistributionOriginGroupSelectionCriteria) ToDistributionOriginGroupSelectionCriteriaPtrOutputWithContext(ctx context.Context) DistributionOriginGroupSelectionCriteriaPtrOutput {
	return DistributionOriginGroupSelectionCriteria(e).ToDistributionOriginGroupSelectionCriteriaOutputWithContext(ctx).ToDistributionOriginGroupSelectionCriteriaPtrOutputWithContext(ctx)
}

func (e DistributionOriginGroupSelectionCriteria) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DistributionOriginGroupSelectionCriteria) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DistributionOriginGroupSelectionCriteria) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DistributionOriginGroupSelectionCriteria) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DistributionOriginGroupSelectionCriteriaOutput struct{ *pulumi.OutputState }

func (DistributionOriginGroupSelectionCriteriaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DistributionOriginGroupSelectionCriteria)(nil)).Elem()
}

func (o DistributionOriginGroupSelectionCriteriaOutput) ToDistributionOriginGroupSelectionCriteriaOutput() DistributionOriginGroupSelectionCriteriaOutput {
	return o
}

func (o DistributionOriginGroupSelectionCriteriaOutput) ToDistributionOriginGroupSelectionCriteriaOutputWithContext(ctx context.Context) DistributionOriginGroupSelectionCriteriaOutput {
	return o
}

func (o DistributionOriginGroupSelectionCriteriaOutput) ToDistributionOriginGroupSelectionCriteriaPtrOutput() DistributionOriginGroupSelectionCriteriaPtrOutput {
	return o.ToDistributionOriginGroupSelectionCriteriaPtrOutputWithContext(context.Background())
}

func (o DistributionOriginGroupSelectionCriteriaOutput) ToDistributionOriginGroupSelectionCriteriaPtrOutputWithContext(ctx context.Context) DistributionOriginGroupSelectionCriteriaPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DistributionOriginGroupSelectionCriteria) *DistributionOriginGroupSelectionCriteria {
		return &v
	}).(DistributionOriginGroupSelectionCriteriaPtrOutput)
}

func (o DistributionOriginGroupSelectionCriteriaOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DistributionOriginGroupSelectionCriteriaOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DistributionOriginGroupSelectionCriteria) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DistributionOriginGroupSelectionCriteriaOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DistributionOriginGroupSelectionCriteriaOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DistributionOriginGroupSelectionCriteria) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DistributionOriginGroupSelectionCriteriaPtrOutput struct{ *pulumi.OutputState }

func (DistributionOriginGroupSelectionCriteriaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DistributionOriginGroupSelectionCriteria)(nil)).Elem()
}

func (o DistributionOriginGroupSelectionCriteriaPtrOutput) ToDistributionOriginGroupSelectionCriteriaPtrOutput() DistributionOriginGroupSelectionCriteriaPtrOutput {
	return o
}

func (o DistributionOriginGroupSelectionCriteriaPtrOutput) ToDistributionOriginGroupSelectionCriteriaPtrOutputWithContext(ctx context.Context) DistributionOriginGroupSelectionCriteriaPtrOutput {
	return o
}

func (o DistributionOriginGroupSelectionCriteriaPtrOutput) Elem() DistributionOriginGroupSelectionCriteriaOutput {
	return o.ApplyT(func(v *DistributionOriginGroupSelectionCriteria) DistributionOriginGroupSelectionCriteria {
		if v != nil {
			return *v
		}
		var ret DistributionOriginGroupSelectionCriteria
		return ret
	}).(DistributionOriginGroupSelectionCriteriaOutput)
}

func (o DistributionOriginGroupSelectionCriteriaPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DistributionOriginGroupSelectionCriteriaPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DistributionOriginGroupSelectionCriteria) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DistributionOriginGroupSelectionCriteriaInput is an input type that accepts values of the DistributionOriginGroupSelectionCriteria enum
// A concrete instance of `DistributionOriginGroupSelectionCriteriaInput` can be one of the following:
//
//	DistributionOriginGroupSelectionCriteriaDefault
//	DistributionOriginGroupSelectionCriteriaMediaQualityBased
type DistributionOriginGroupSelectionCriteriaInput interface {
	pulumi.Input

	ToDistributionOriginGroupSelectionCriteriaOutput() DistributionOriginGroupSelectionCriteriaOutput
	ToDistributionOriginGroupSelectionCriteriaOutputWithContext(context.Context) DistributionOriginGroupSelectionCriteriaOutput
}

var distributionOriginGroupSelectionCriteriaPtrType = reflect.TypeOf((**DistributionOriginGroupSelectionCriteria)(nil)).Elem()

type DistributionOriginGroupSelectionCriteriaPtrInput interface {
	pulumi.Input

	ToDistributionOriginGroupSelectionCriteriaPtrOutput() DistributionOriginGroupSelectionCriteriaPtrOutput
	ToDistributionOriginGroupSelectionCriteriaPtrOutputWithContext(context.Context) DistributionOriginGroupSelectionCriteriaPtrOutput
}

type distributionOriginGroupSelectionCriteriaPtr string

func DistributionOriginGroupSelectionCriteriaPtr(v string) DistributionOriginGroupSelectionCriteriaPtrInput {
	return (*distributionOriginGroupSelectionCriteriaPtr)(&v)
}

func (*distributionOriginGroupSelectionCriteriaPtr) ElementType() reflect.Type {
	return distributionOriginGroupSelectionCriteriaPtrType
}

func (in *distributionOriginGroupSelectionCriteriaPtr) ToDistributionOriginGroupSelectionCriteriaPtrOutput() DistributionOriginGroupSelectionCriteriaPtrOutput {
	return pulumi.ToOutput(in).(DistributionOriginGroupSelectionCriteriaPtrOutput)
}

func (in *distributionOriginGroupSelectionCriteriaPtr) ToDistributionOriginGroupSelectionCriteriaPtrOutputWithContext(ctx context.Context) DistributionOriginGroupSelectionCriteriaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DistributionOriginGroupSelectionCriteriaPtrOutput)
}

// A flag that indicates whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
type MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus string

const (
	MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusEnabled  = MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus("Enabled")
	MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusDisabled = MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus("Disabled")
)

func (MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus)(nil)).Elem()
}

func (e MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus) ToMonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput() MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput {
	return pulumi.ToOutput(e).(MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput)
}

func (e MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus) ToMonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutputWithContext(ctx context.Context) MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput)
}

func (e MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus) ToMonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput() MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput {
	return e.ToMonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutputWithContext(context.Background())
}

func (e MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus) ToMonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutputWithContext(ctx context.Context) MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput {
	return MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus(e).ToMonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutputWithContext(ctx).ToMonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutputWithContext(ctx)
}

func (e MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput struct{ *pulumi.OutputState }

func (MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus)(nil)).Elem()
}

func (o MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput) ToMonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput() MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput {
	return o
}

func (o MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput) ToMonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutputWithContext(ctx context.Context) MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput {
	return o
}

func (o MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput) ToMonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput() MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput {
	return o.ToMonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutputWithContext(context.Background())
}

func (o MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput) ToMonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutputWithContext(ctx context.Context) MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus) *MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus {
		return &v
	}).(MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput)
}

func (o MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput struct{ *pulumi.OutputState }

func (MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus)(nil)).Elem()
}

func (o MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput) ToMonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput() MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput {
	return o
}

func (o MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput) ToMonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutputWithContext(ctx context.Context) MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput {
	return o
}

func (o MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput) Elem() MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput {
	return o.ApplyT(func(v *MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus) MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus {
		if v != nil {
			return *v
		}
		var ret MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus
		return ret
	}).(MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput)
}

func (o MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusInput is an input type that accepts values of the MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus enum
// A concrete instance of `MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusInput` can be one of the following:
//
//	MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusEnabled
//	MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusDisabled
type MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusInput interface {
	pulumi.Input

	ToMonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput() MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput
	ToMonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutputWithContext(context.Context) MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput
}

var monitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrType = reflect.TypeOf((**MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus)(nil)).Elem()

type MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrInput interface {
	pulumi.Input

	ToMonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput() MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput
	ToMonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutputWithContext(context.Context) MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput
}

type monitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtr string

func MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtr(v string) MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrInput {
	return (*monitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtr)(&v)
}

func (*monitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtr) ElementType() reflect.Type {
	return monitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrType
}

func (in *monitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtr) ToMonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput() MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput {
	return pulumi.ToOutput(in).(MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput)
}

func (in *monitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtr) ToMonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutputWithContext(ctx context.Context) MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContinuousDeploymentPolicyConfigTypeInput)(nil)).Elem(), ContinuousDeploymentPolicyConfigType("SingleWeight"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContinuousDeploymentPolicyConfigTypePtrInput)(nil)).Elem(), ContinuousDeploymentPolicyConfigType("SingleWeight"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContinuousDeploymentPolicyTrafficConfigTypeInput)(nil)).Elem(), ContinuousDeploymentPolicyTrafficConfigType("SingleWeight"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContinuousDeploymentPolicyTrafficConfigTypePtrInput)(nil)).Elem(), ContinuousDeploymentPolicyTrafficConfigType("SingleWeight"))
	pulumi.RegisterInputType(reflect.TypeOf((*DistributionOriginGroupSelectionCriteriaInput)(nil)).Elem(), DistributionOriginGroupSelectionCriteria("default"))
	pulumi.RegisterInputType(reflect.TypeOf((*DistributionOriginGroupSelectionCriteriaPtrInput)(nil)).Elem(), DistributionOriginGroupSelectionCriteria("default"))
	pulumi.RegisterInputType(reflect.TypeOf((*MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusInput)(nil)).Elem(), MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrInput)(nil)).Elem(), MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatus("Enabled"))
	pulumi.RegisterOutputType(ContinuousDeploymentPolicyConfigTypeOutput{})
	pulumi.RegisterOutputType(ContinuousDeploymentPolicyConfigTypePtrOutput{})
	pulumi.RegisterOutputType(ContinuousDeploymentPolicyTrafficConfigTypeOutput{})
	pulumi.RegisterOutputType(ContinuousDeploymentPolicyTrafficConfigTypePtrOutput{})
	pulumi.RegisterOutputType(DistributionOriginGroupSelectionCriteriaOutput{})
	pulumi.RegisterOutputType(DistributionOriginGroupSelectionCriteriaPtrOutput{})
	pulumi.RegisterOutputType(MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusOutput{})
	pulumi.RegisterOutputType(MonitoringSubscriptionRealtimeMetricsSubscriptionConfigRealtimeMetricsSubscriptionStatusPtrOutput{})
}
