// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package personalize

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The domain of a Domain dataset group.
type DatasetGroupDomain string

const (
	DatasetGroupDomainEcommerce     = DatasetGroupDomain("ECOMMERCE")
	DatasetGroupDomainVideoOnDemand = DatasetGroupDomain("VIDEO_ON_DEMAND")
)

func (DatasetGroupDomain) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetGroupDomain)(nil)).Elem()
}

func (e DatasetGroupDomain) ToDatasetGroupDomainOutput() DatasetGroupDomainOutput {
	return pulumi.ToOutput(e).(DatasetGroupDomainOutput)
}

func (e DatasetGroupDomain) ToDatasetGroupDomainOutputWithContext(ctx context.Context) DatasetGroupDomainOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DatasetGroupDomainOutput)
}

func (e DatasetGroupDomain) ToDatasetGroupDomainPtrOutput() DatasetGroupDomainPtrOutput {
	return e.ToDatasetGroupDomainPtrOutputWithContext(context.Background())
}

func (e DatasetGroupDomain) ToDatasetGroupDomainPtrOutputWithContext(ctx context.Context) DatasetGroupDomainPtrOutput {
	return DatasetGroupDomain(e).ToDatasetGroupDomainOutputWithContext(ctx).ToDatasetGroupDomainPtrOutputWithContext(ctx)
}

func (e DatasetGroupDomain) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatasetGroupDomain) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatasetGroupDomain) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DatasetGroupDomain) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DatasetGroupDomainOutput struct{ *pulumi.OutputState }

func (DatasetGroupDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetGroupDomain)(nil)).Elem()
}

func (o DatasetGroupDomainOutput) ToDatasetGroupDomainOutput() DatasetGroupDomainOutput {
	return o
}

func (o DatasetGroupDomainOutput) ToDatasetGroupDomainOutputWithContext(ctx context.Context) DatasetGroupDomainOutput {
	return o
}

func (o DatasetGroupDomainOutput) ToDatasetGroupDomainPtrOutput() DatasetGroupDomainPtrOutput {
	return o.ToDatasetGroupDomainPtrOutputWithContext(context.Background())
}

func (o DatasetGroupDomainOutput) ToDatasetGroupDomainPtrOutputWithContext(ctx context.Context) DatasetGroupDomainPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatasetGroupDomain) *DatasetGroupDomain {
		return &v
	}).(DatasetGroupDomainPtrOutput)
}

func (o DatasetGroupDomainOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DatasetGroupDomainOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatasetGroupDomain) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DatasetGroupDomainOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatasetGroupDomainOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatasetGroupDomain) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DatasetGroupDomainPtrOutput struct{ *pulumi.OutputState }

func (DatasetGroupDomainPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetGroupDomain)(nil)).Elem()
}

func (o DatasetGroupDomainPtrOutput) ToDatasetGroupDomainPtrOutput() DatasetGroupDomainPtrOutput {
	return o
}

func (o DatasetGroupDomainPtrOutput) ToDatasetGroupDomainPtrOutputWithContext(ctx context.Context) DatasetGroupDomainPtrOutput {
	return o
}

func (o DatasetGroupDomainPtrOutput) Elem() DatasetGroupDomainOutput {
	return o.ApplyT(func(v *DatasetGroupDomain) DatasetGroupDomain {
		if v != nil {
			return *v
		}
		var ret DatasetGroupDomain
		return ret
	}).(DatasetGroupDomainOutput)
}

func (o DatasetGroupDomainPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatasetGroupDomainPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DatasetGroupDomain) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DatasetGroupDomainInput is an input type that accepts DatasetGroupDomainArgs and DatasetGroupDomainOutput values.
// You can construct a concrete instance of `DatasetGroupDomainInput` via:
//
//	DatasetGroupDomainArgs{...}
type DatasetGroupDomainInput interface {
	pulumi.Input

	ToDatasetGroupDomainOutput() DatasetGroupDomainOutput
	ToDatasetGroupDomainOutputWithContext(context.Context) DatasetGroupDomainOutput
}

var datasetGroupDomainPtrType = reflect.TypeOf((**DatasetGroupDomain)(nil)).Elem()

type DatasetGroupDomainPtrInput interface {
	pulumi.Input

	ToDatasetGroupDomainPtrOutput() DatasetGroupDomainPtrOutput
	ToDatasetGroupDomainPtrOutputWithContext(context.Context) DatasetGroupDomainPtrOutput
}

type datasetGroupDomainPtr string

func DatasetGroupDomainPtr(v string) DatasetGroupDomainPtrInput {
	return (*datasetGroupDomainPtr)(&v)
}

func (*datasetGroupDomainPtr) ElementType() reflect.Type {
	return datasetGroupDomainPtrType
}

func (in *datasetGroupDomainPtr) ToDatasetGroupDomainPtrOutput() DatasetGroupDomainPtrOutput {
	return pulumi.ToOutput(in).(DatasetGroupDomainPtrOutput)
}

func (in *datasetGroupDomainPtr) ToDatasetGroupDomainPtrOutputWithContext(ctx context.Context) DatasetGroupDomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DatasetGroupDomainPtrOutput)
}

// The type of dataset
type DatasetType string

const (
	DatasetTypeInteractions = DatasetType("Interactions")
	DatasetTypeItems        = DatasetType("Items")
	DatasetTypeUsers        = DatasetType("Users")
)

func (DatasetType) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetType)(nil)).Elem()
}

func (e DatasetType) ToDatasetTypeOutput() DatasetTypeOutput {
	return pulumi.ToOutput(e).(DatasetTypeOutput)
}

func (e DatasetType) ToDatasetTypeOutputWithContext(ctx context.Context) DatasetTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DatasetTypeOutput)
}

func (e DatasetType) ToDatasetTypePtrOutput() DatasetTypePtrOutput {
	return e.ToDatasetTypePtrOutputWithContext(context.Background())
}

func (e DatasetType) ToDatasetTypePtrOutputWithContext(ctx context.Context) DatasetTypePtrOutput {
	return DatasetType(e).ToDatasetTypeOutputWithContext(ctx).ToDatasetTypePtrOutputWithContext(ctx)
}

func (e DatasetType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatasetType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatasetType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DatasetType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DatasetTypeOutput struct{ *pulumi.OutputState }

func (DatasetTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetType)(nil)).Elem()
}

func (o DatasetTypeOutput) ToDatasetTypeOutput() DatasetTypeOutput {
	return o
}

func (o DatasetTypeOutput) ToDatasetTypeOutputWithContext(ctx context.Context) DatasetTypeOutput {
	return o
}

func (o DatasetTypeOutput) ToDatasetTypePtrOutput() DatasetTypePtrOutput {
	return o.ToDatasetTypePtrOutputWithContext(context.Background())
}

func (o DatasetTypeOutput) ToDatasetTypePtrOutputWithContext(ctx context.Context) DatasetTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatasetType) *DatasetType {
		return &v
	}).(DatasetTypePtrOutput)
}

func (o DatasetTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DatasetTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatasetType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DatasetTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatasetTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatasetType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DatasetTypePtrOutput struct{ *pulumi.OutputState }

func (DatasetTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetType)(nil)).Elem()
}

func (o DatasetTypePtrOutput) ToDatasetTypePtrOutput() DatasetTypePtrOutput {
	return o
}

func (o DatasetTypePtrOutput) ToDatasetTypePtrOutputWithContext(ctx context.Context) DatasetTypePtrOutput {
	return o
}

func (o DatasetTypePtrOutput) Elem() DatasetTypeOutput {
	return o.ApplyT(func(v *DatasetType) DatasetType {
		if v != nil {
			return *v
		}
		var ret DatasetType
		return ret
	}).(DatasetTypeOutput)
}

func (o DatasetTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatasetTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DatasetType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DatasetTypeInput is an input type that accepts DatasetTypeArgs and DatasetTypeOutput values.
// You can construct a concrete instance of `DatasetTypeInput` via:
//
//	DatasetTypeArgs{...}
type DatasetTypeInput interface {
	pulumi.Input

	ToDatasetTypeOutput() DatasetTypeOutput
	ToDatasetTypeOutputWithContext(context.Context) DatasetTypeOutput
}

var datasetTypePtrType = reflect.TypeOf((**DatasetType)(nil)).Elem()

type DatasetTypePtrInput interface {
	pulumi.Input

	ToDatasetTypePtrOutput() DatasetTypePtrOutput
	ToDatasetTypePtrOutputWithContext(context.Context) DatasetTypePtrOutput
}

type datasetTypePtr string

func DatasetTypePtr(v string) DatasetTypePtrInput {
	return (*datasetTypePtr)(&v)
}

func (*datasetTypePtr) ElementType() reflect.Type {
	return datasetTypePtrType
}

func (in *datasetTypePtr) ToDatasetTypePtrOutput() DatasetTypePtrOutput {
	return pulumi.ToOutput(in).(DatasetTypePtrOutput)
}

func (in *datasetTypePtr) ToDatasetTypePtrOutputWithContext(ctx context.Context) DatasetTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DatasetTypePtrOutput)
}

// The domain of a Domain dataset group.
type SchemaDomain string

const (
	SchemaDomainEcommerce     = SchemaDomain("ECOMMERCE")
	SchemaDomainVideoOnDemand = SchemaDomain("VIDEO_ON_DEMAND")
)

func (SchemaDomain) ElementType() reflect.Type {
	return reflect.TypeOf((*SchemaDomain)(nil)).Elem()
}

func (e SchemaDomain) ToSchemaDomainOutput() SchemaDomainOutput {
	return pulumi.ToOutput(e).(SchemaDomainOutput)
}

func (e SchemaDomain) ToSchemaDomainOutputWithContext(ctx context.Context) SchemaDomainOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SchemaDomainOutput)
}

func (e SchemaDomain) ToSchemaDomainPtrOutput() SchemaDomainPtrOutput {
	return e.ToSchemaDomainPtrOutputWithContext(context.Background())
}

func (e SchemaDomain) ToSchemaDomainPtrOutputWithContext(ctx context.Context) SchemaDomainPtrOutput {
	return SchemaDomain(e).ToSchemaDomainOutputWithContext(ctx).ToSchemaDomainPtrOutputWithContext(ctx)
}

func (e SchemaDomain) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SchemaDomain) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SchemaDomain) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SchemaDomain) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SchemaDomainOutput struct{ *pulumi.OutputState }

func (SchemaDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SchemaDomain)(nil)).Elem()
}

func (o SchemaDomainOutput) ToSchemaDomainOutput() SchemaDomainOutput {
	return o
}

func (o SchemaDomainOutput) ToSchemaDomainOutputWithContext(ctx context.Context) SchemaDomainOutput {
	return o
}

func (o SchemaDomainOutput) ToSchemaDomainPtrOutput() SchemaDomainPtrOutput {
	return o.ToSchemaDomainPtrOutputWithContext(context.Background())
}

func (o SchemaDomainOutput) ToSchemaDomainPtrOutputWithContext(ctx context.Context) SchemaDomainPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SchemaDomain) *SchemaDomain {
		return &v
	}).(SchemaDomainPtrOutput)
}

func (o SchemaDomainOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SchemaDomainOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SchemaDomain) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SchemaDomainOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SchemaDomainOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SchemaDomain) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SchemaDomainPtrOutput struct{ *pulumi.OutputState }

func (SchemaDomainPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SchemaDomain)(nil)).Elem()
}

func (o SchemaDomainPtrOutput) ToSchemaDomainPtrOutput() SchemaDomainPtrOutput {
	return o
}

func (o SchemaDomainPtrOutput) ToSchemaDomainPtrOutputWithContext(ctx context.Context) SchemaDomainPtrOutput {
	return o
}

func (o SchemaDomainPtrOutput) Elem() SchemaDomainOutput {
	return o.ApplyT(func(v *SchemaDomain) SchemaDomain {
		if v != nil {
			return *v
		}
		var ret SchemaDomain
		return ret
	}).(SchemaDomainOutput)
}

func (o SchemaDomainPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SchemaDomainPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SchemaDomain) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SchemaDomainInput is an input type that accepts SchemaDomainArgs and SchemaDomainOutput values.
// You can construct a concrete instance of `SchemaDomainInput` via:
//
//	SchemaDomainArgs{...}
type SchemaDomainInput interface {
	pulumi.Input

	ToSchemaDomainOutput() SchemaDomainOutput
	ToSchemaDomainOutputWithContext(context.Context) SchemaDomainOutput
}

var schemaDomainPtrType = reflect.TypeOf((**SchemaDomain)(nil)).Elem()

type SchemaDomainPtrInput interface {
	pulumi.Input

	ToSchemaDomainPtrOutput() SchemaDomainPtrOutput
	ToSchemaDomainPtrOutputWithContext(context.Context) SchemaDomainPtrOutput
}

type schemaDomainPtr string

func SchemaDomainPtr(v string) SchemaDomainPtrInput {
	return (*schemaDomainPtr)(&v)
}

func (*schemaDomainPtr) ElementType() reflect.Type {
	return schemaDomainPtrType
}

func (in *schemaDomainPtr) ToSchemaDomainPtrOutput() SchemaDomainPtrOutput {
	return pulumi.ToOutput(in).(SchemaDomainPtrOutput)
}

func (in *schemaDomainPtr) ToSchemaDomainPtrOutputWithContext(ctx context.Context) SchemaDomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SchemaDomainPtrOutput)
}

// The type of the metric. Valid values are Maximize and Minimize.
type SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType string

const (
	SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeMaximize = SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType("Maximize")
	SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeMinimize = SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType("Minimize")
)

func (SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType) ElementType() reflect.Type {
	return reflect.TypeOf((*SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType)(nil)).Elem()
}

func (e SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType) ToSolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput() SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput {
	return pulumi.ToOutput(e).(SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput)
}

func (e SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType) ToSolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutputWithContext(ctx context.Context) SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput)
}

func (e SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType) ToSolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput() SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput {
	return e.ToSolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutputWithContext(context.Background())
}

func (e SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType) ToSolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutputWithContext(ctx context.Context) SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput {
	return SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType(e).ToSolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutputWithContext(ctx).ToSolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutputWithContext(ctx)
}

func (e SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput struct{ *pulumi.OutputState }

func (SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType)(nil)).Elem()
}

func (o SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput) ToSolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput() SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput {
	return o
}

func (o SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput) ToSolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutputWithContext(ctx context.Context) SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput {
	return o
}

func (o SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput) ToSolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput() SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput {
	return o.ToSolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutputWithContext(context.Background())
}

func (o SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput) ToSolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutputWithContext(ctx context.Context) SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType) *SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType {
		return &v
	}).(SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput)
}

func (o SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput struct{ *pulumi.OutputState }

func (SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType)(nil)).Elem()
}

func (o SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput) ToSolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput() SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput {
	return o
}

func (o SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput) ToSolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutputWithContext(ctx context.Context) SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput {
	return o
}

func (o SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput) Elem() SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput {
	return o.ApplyT(func(v *SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType) SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType {
		if v != nil {
			return *v
		}
		var ret SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType
		return ret
	}).(SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput)
}

func (o SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeInput is an input type that accepts SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeArgs and SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput values.
// You can construct a concrete instance of `SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeInput` via:
//
//	SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeArgs{...}
type SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeInput interface {
	pulumi.Input

	ToSolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput() SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput
	ToSolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutputWithContext(context.Context) SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput
}

var solutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrType = reflect.TypeOf((**SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType)(nil)).Elem()

type SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrInput interface {
	pulumi.Input

	ToSolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput() SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput
	ToSolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutputWithContext(context.Context) SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput
}

type solutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtr string

func SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtr(v string) SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrInput {
	return (*solutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtr)(&v)
}

func (*solutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtr) ElementType() reflect.Type {
	return solutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrType
}

func (in *solutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtr) ToSolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput() SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput {
	return pulumi.ToOutput(in).(SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput)
}

func (in *solutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtr) ToSolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutputWithContext(ctx context.Context) SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetGroupDomainInput)(nil)).Elem(), DatasetGroupDomain("ECOMMERCE"))
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetGroupDomainPtrInput)(nil)).Elem(), DatasetGroupDomain("ECOMMERCE"))
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetTypeInput)(nil)).Elem(), DatasetType("Interactions"))
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetTypePtrInput)(nil)).Elem(), DatasetType("Interactions"))
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaDomainInput)(nil)).Elem(), SchemaDomain("ECOMMERCE"))
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaDomainPtrInput)(nil)).Elem(), SchemaDomain("ECOMMERCE"))
	pulumi.RegisterInputType(reflect.TypeOf((*SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeInput)(nil)).Elem(), SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType("Maximize"))
	pulumi.RegisterInputType(reflect.TypeOf((*SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrInput)(nil)).Elem(), SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType("Maximize"))
	pulumi.RegisterOutputType(DatasetGroupDomainOutput{})
	pulumi.RegisterOutputType(DatasetGroupDomainPtrOutput{})
	pulumi.RegisterOutputType(DatasetTypeOutput{})
	pulumi.RegisterOutputType(DatasetTypePtrOutput{})
	pulumi.RegisterOutputType(SchemaDomainOutput{})
	pulumi.RegisterOutputType(SchemaDomainPtrOutput{})
	pulumi.RegisterOutputType(SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypeOutput{})
	pulumi.RegisterOutputType(SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesTypePtrOutput{})
}