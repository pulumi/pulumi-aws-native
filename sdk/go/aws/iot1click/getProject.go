// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot1click

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IoT1Click::Project
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	var rv LookupProjectResult
	err := ctx.Invoke("aws-native:iot1click:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProjectArgs struct {
	Id string `pulumi:"id"`
}

type LookupProjectResult struct {
	Arn               *string                   `pulumi:"arn"`
	Description       *string                   `pulumi:"description"`
	Id                *string                   `pulumi:"id"`
	PlacementTemplate *ProjectPlacementTemplate `pulumi:"placementTemplate"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectResult, error) {
			args := v.(LookupProjectArgs)
			r, err := LookupProject(ctx, &args, opts...)
			var s LookupProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectResultOutput)
}

type LookupProjectOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupProjectResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupProjectResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupProjectResultOutput) PlacementTemplate() ProjectPlacementTemplatePtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *ProjectPlacementTemplate { return v.PlacementTemplate }).(ProjectPlacementTemplatePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}