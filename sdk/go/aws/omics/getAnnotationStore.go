// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package omics

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Omics::AnnotationStore Resource Type
func LookupAnnotationStore(ctx *pulumi.Context, args *LookupAnnotationStoreArgs, opts ...pulumi.InvokeOption) (*LookupAnnotationStoreResult, error) {
	var rv LookupAnnotationStoreResult
	err := ctx.Invoke("aws-native:omics:getAnnotationStore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAnnotationStoreArgs struct {
	Name string `pulumi:"name"`
}

type LookupAnnotationStoreResult struct {
	CreationTime   *string                     `pulumi:"creationTime"`
	Description    *string                     `pulumi:"description"`
	Id             *string                     `pulumi:"id"`
	Status         *AnnotationStoreStoreStatus `pulumi:"status"`
	StatusMessage  *string                     `pulumi:"statusMessage"`
	StoreArn       *string                     `pulumi:"storeArn"`
	StoreSizeBytes *float64                    `pulumi:"storeSizeBytes"`
	UpdateTime     *string                     `pulumi:"updateTime"`
}

func LookupAnnotationStoreOutput(ctx *pulumi.Context, args LookupAnnotationStoreOutputArgs, opts ...pulumi.InvokeOption) LookupAnnotationStoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAnnotationStoreResult, error) {
			args := v.(LookupAnnotationStoreArgs)
			r, err := LookupAnnotationStore(ctx, &args, opts...)
			var s LookupAnnotationStoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAnnotationStoreResultOutput)
}

type LookupAnnotationStoreOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupAnnotationStoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnnotationStoreArgs)(nil)).Elem()
}

type LookupAnnotationStoreResultOutput struct{ *pulumi.OutputState }

func (LookupAnnotationStoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnnotationStoreResult)(nil)).Elem()
}

func (o LookupAnnotationStoreResultOutput) ToLookupAnnotationStoreResultOutput() LookupAnnotationStoreResultOutput {
	return o
}

func (o LookupAnnotationStoreResultOutput) ToLookupAnnotationStoreResultOutputWithContext(ctx context.Context) LookupAnnotationStoreResultOutput {
	return o
}

func (o LookupAnnotationStoreResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnnotationStoreResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupAnnotationStoreResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnnotationStoreResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupAnnotationStoreResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnnotationStoreResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupAnnotationStoreResultOutput) Status() AnnotationStoreStoreStatusPtrOutput {
	return o.ApplyT(func(v LookupAnnotationStoreResult) *AnnotationStoreStoreStatus { return v.Status }).(AnnotationStoreStoreStatusPtrOutput)
}

func (o LookupAnnotationStoreResultOutput) StatusMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnnotationStoreResult) *string { return v.StatusMessage }).(pulumi.StringPtrOutput)
}

func (o LookupAnnotationStoreResultOutput) StoreArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnnotationStoreResult) *string { return v.StoreArn }).(pulumi.StringPtrOutput)
}

func (o LookupAnnotationStoreResultOutput) StoreSizeBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupAnnotationStoreResult) *float64 { return v.StoreSizeBytes }).(pulumi.Float64PtrOutput)
}

func (o LookupAnnotationStoreResultOutput) UpdateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnnotationStoreResult) *string { return v.UpdateTime }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAnnotationStoreResultOutput{})
}