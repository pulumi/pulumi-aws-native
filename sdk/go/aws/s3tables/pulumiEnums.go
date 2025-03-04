// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3tables

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Indicates whether the Unreferenced File Removal maintenance action is enabled.
type TableBucketUnreferencedFileRemovalStatus string

const (
	TableBucketUnreferencedFileRemovalStatusEnabled  = TableBucketUnreferencedFileRemovalStatus("Enabled")
	TableBucketUnreferencedFileRemovalStatusDisabled = TableBucketUnreferencedFileRemovalStatus("Disabled")
)

func (TableBucketUnreferencedFileRemovalStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*TableBucketUnreferencedFileRemovalStatus)(nil)).Elem()
}

func (e TableBucketUnreferencedFileRemovalStatus) ToTableBucketUnreferencedFileRemovalStatusOutput() TableBucketUnreferencedFileRemovalStatusOutput {
	return pulumi.ToOutput(e).(TableBucketUnreferencedFileRemovalStatusOutput)
}

func (e TableBucketUnreferencedFileRemovalStatus) ToTableBucketUnreferencedFileRemovalStatusOutputWithContext(ctx context.Context) TableBucketUnreferencedFileRemovalStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TableBucketUnreferencedFileRemovalStatusOutput)
}

func (e TableBucketUnreferencedFileRemovalStatus) ToTableBucketUnreferencedFileRemovalStatusPtrOutput() TableBucketUnreferencedFileRemovalStatusPtrOutput {
	return e.ToTableBucketUnreferencedFileRemovalStatusPtrOutputWithContext(context.Background())
}

func (e TableBucketUnreferencedFileRemovalStatus) ToTableBucketUnreferencedFileRemovalStatusPtrOutputWithContext(ctx context.Context) TableBucketUnreferencedFileRemovalStatusPtrOutput {
	return TableBucketUnreferencedFileRemovalStatus(e).ToTableBucketUnreferencedFileRemovalStatusOutputWithContext(ctx).ToTableBucketUnreferencedFileRemovalStatusPtrOutputWithContext(ctx)
}

func (e TableBucketUnreferencedFileRemovalStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TableBucketUnreferencedFileRemovalStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TableBucketUnreferencedFileRemovalStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TableBucketUnreferencedFileRemovalStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TableBucketUnreferencedFileRemovalStatusOutput struct{ *pulumi.OutputState }

func (TableBucketUnreferencedFileRemovalStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableBucketUnreferencedFileRemovalStatus)(nil)).Elem()
}

func (o TableBucketUnreferencedFileRemovalStatusOutput) ToTableBucketUnreferencedFileRemovalStatusOutput() TableBucketUnreferencedFileRemovalStatusOutput {
	return o
}

func (o TableBucketUnreferencedFileRemovalStatusOutput) ToTableBucketUnreferencedFileRemovalStatusOutputWithContext(ctx context.Context) TableBucketUnreferencedFileRemovalStatusOutput {
	return o
}

func (o TableBucketUnreferencedFileRemovalStatusOutput) ToTableBucketUnreferencedFileRemovalStatusPtrOutput() TableBucketUnreferencedFileRemovalStatusPtrOutput {
	return o.ToTableBucketUnreferencedFileRemovalStatusPtrOutputWithContext(context.Background())
}

func (o TableBucketUnreferencedFileRemovalStatusOutput) ToTableBucketUnreferencedFileRemovalStatusPtrOutputWithContext(ctx context.Context) TableBucketUnreferencedFileRemovalStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TableBucketUnreferencedFileRemovalStatus) *TableBucketUnreferencedFileRemovalStatus {
		return &v
	}).(TableBucketUnreferencedFileRemovalStatusPtrOutput)
}

func (o TableBucketUnreferencedFileRemovalStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TableBucketUnreferencedFileRemovalStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TableBucketUnreferencedFileRemovalStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TableBucketUnreferencedFileRemovalStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TableBucketUnreferencedFileRemovalStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TableBucketUnreferencedFileRemovalStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TableBucketUnreferencedFileRemovalStatusPtrOutput struct{ *pulumi.OutputState }

func (TableBucketUnreferencedFileRemovalStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableBucketUnreferencedFileRemovalStatus)(nil)).Elem()
}

func (o TableBucketUnreferencedFileRemovalStatusPtrOutput) ToTableBucketUnreferencedFileRemovalStatusPtrOutput() TableBucketUnreferencedFileRemovalStatusPtrOutput {
	return o
}

func (o TableBucketUnreferencedFileRemovalStatusPtrOutput) ToTableBucketUnreferencedFileRemovalStatusPtrOutputWithContext(ctx context.Context) TableBucketUnreferencedFileRemovalStatusPtrOutput {
	return o
}

func (o TableBucketUnreferencedFileRemovalStatusPtrOutput) Elem() TableBucketUnreferencedFileRemovalStatusOutput {
	return o.ApplyT(func(v *TableBucketUnreferencedFileRemovalStatus) TableBucketUnreferencedFileRemovalStatus {
		if v != nil {
			return *v
		}
		var ret TableBucketUnreferencedFileRemovalStatus
		return ret
	}).(TableBucketUnreferencedFileRemovalStatusOutput)
}

func (o TableBucketUnreferencedFileRemovalStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TableBucketUnreferencedFileRemovalStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TableBucketUnreferencedFileRemovalStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// TableBucketUnreferencedFileRemovalStatusInput is an input type that accepts values of the TableBucketUnreferencedFileRemovalStatus enum
// A concrete instance of `TableBucketUnreferencedFileRemovalStatusInput` can be one of the following:
//
//	TableBucketUnreferencedFileRemovalStatusEnabled
//	TableBucketUnreferencedFileRemovalStatusDisabled
type TableBucketUnreferencedFileRemovalStatusInput interface {
	pulumi.Input

	ToTableBucketUnreferencedFileRemovalStatusOutput() TableBucketUnreferencedFileRemovalStatusOutput
	ToTableBucketUnreferencedFileRemovalStatusOutputWithContext(context.Context) TableBucketUnreferencedFileRemovalStatusOutput
}

var tableBucketUnreferencedFileRemovalStatusPtrType = reflect.TypeOf((**TableBucketUnreferencedFileRemovalStatus)(nil)).Elem()

type TableBucketUnreferencedFileRemovalStatusPtrInput interface {
	pulumi.Input

	ToTableBucketUnreferencedFileRemovalStatusPtrOutput() TableBucketUnreferencedFileRemovalStatusPtrOutput
	ToTableBucketUnreferencedFileRemovalStatusPtrOutputWithContext(context.Context) TableBucketUnreferencedFileRemovalStatusPtrOutput
}

type tableBucketUnreferencedFileRemovalStatusPtr string

func TableBucketUnreferencedFileRemovalStatusPtr(v string) TableBucketUnreferencedFileRemovalStatusPtrInput {
	return (*tableBucketUnreferencedFileRemovalStatusPtr)(&v)
}

func (*tableBucketUnreferencedFileRemovalStatusPtr) ElementType() reflect.Type {
	return tableBucketUnreferencedFileRemovalStatusPtrType
}

func (in *tableBucketUnreferencedFileRemovalStatusPtr) ToTableBucketUnreferencedFileRemovalStatusPtrOutput() TableBucketUnreferencedFileRemovalStatusPtrOutput {
	return pulumi.ToOutput(in).(TableBucketUnreferencedFileRemovalStatusPtrOutput)
}

func (in *tableBucketUnreferencedFileRemovalStatusPtr) ToTableBucketUnreferencedFileRemovalStatusPtrOutputWithContext(ctx context.Context) TableBucketUnreferencedFileRemovalStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TableBucketUnreferencedFileRemovalStatusPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TableBucketUnreferencedFileRemovalStatusInput)(nil)).Elem(), TableBucketUnreferencedFileRemovalStatus("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*TableBucketUnreferencedFileRemovalStatusPtrInput)(nil)).Elem(), TableBucketUnreferencedFileRemovalStatus("Enabled"))
	pulumi.RegisterOutputType(TableBucketUnreferencedFileRemovalStatusOutput{})
	pulumi.RegisterOutputType(TableBucketUnreferencedFileRemovalStatusPtrOutput{})
}
