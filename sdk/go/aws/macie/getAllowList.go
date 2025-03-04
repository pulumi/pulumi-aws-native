// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package macie

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Macie AllowList resource schema
func LookupAllowList(ctx *pulumi.Context, args *LookupAllowListArgs, opts ...pulumi.InvokeOption) (*LookupAllowListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAllowListResult
	err := ctx.Invoke("aws-native:macie:getAllowList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAllowListArgs struct {
	// AllowList ID.
	Id string `pulumi:"id"`
}

type LookupAllowListResult struct {
	// AllowList ARN.
	Arn *string `pulumi:"arn"`
	// AllowList criteria.
	Criteria *AllowListCriteria `pulumi:"criteria"`
	// Description of AllowList.
	Description *string `pulumi:"description"`
	// AllowList ID.
	Id *string `pulumi:"id"`
	// Name of AllowList.
	Name *string `pulumi:"name"`
	// AllowList status.
	Status *AllowListStatus `pulumi:"status"`
	// A collection of tags associated with a resource
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupAllowListOutput(ctx *pulumi.Context, args LookupAllowListOutputArgs, opts ...pulumi.InvokeOption) LookupAllowListResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAllowListResultOutput, error) {
			args := v.(LookupAllowListArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:macie:getAllowList", args, LookupAllowListResultOutput{}, options).(LookupAllowListResultOutput), nil
		}).(LookupAllowListResultOutput)
}

type LookupAllowListOutputArgs struct {
	// AllowList ID.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupAllowListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAllowListArgs)(nil)).Elem()
}

type LookupAllowListResultOutput struct{ *pulumi.OutputState }

func (LookupAllowListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAllowListResult)(nil)).Elem()
}

func (o LookupAllowListResultOutput) ToLookupAllowListResultOutput() LookupAllowListResultOutput {
	return o
}

func (o LookupAllowListResultOutput) ToLookupAllowListResultOutputWithContext(ctx context.Context) LookupAllowListResultOutput {
	return o
}

// AllowList ARN.
func (o LookupAllowListResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAllowListResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// AllowList criteria.
func (o LookupAllowListResultOutput) Criteria() AllowListCriteriaPtrOutput {
	return o.ApplyT(func(v LookupAllowListResult) *AllowListCriteria { return v.Criteria }).(AllowListCriteriaPtrOutput)
}

// Description of AllowList.
func (o LookupAllowListResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAllowListResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// AllowList ID.
func (o LookupAllowListResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAllowListResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Name of AllowList.
func (o LookupAllowListResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAllowListResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// AllowList status.
func (o LookupAllowListResultOutput) Status() AllowListStatusPtrOutput {
	return o.ApplyT(func(v LookupAllowListResult) *AllowListStatus { return v.Status }).(AllowListStatusPtrOutput)
}

// A collection of tags associated with a resource
func (o LookupAllowListResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupAllowListResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAllowListResultOutput{})
}
