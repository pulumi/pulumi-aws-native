// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package qldb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::QLDB::Stream.
func LookupStream(ctx *pulumi.Context, args *LookupStreamArgs, opts ...pulumi.InvokeOption) (*LookupStreamResult, error) {
	var rv LookupStreamResult
	err := ctx.Invoke("aws-native:qldb:getStream", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStreamArgs struct {
	Id         string `pulumi:"id"`
	LedgerName string `pulumi:"ledgerName"`
}

type LookupStreamResult struct {
	Arn *string `pulumi:"arn"`
	Id  *string `pulumi:"id"`
	// An array of key-value pairs to apply to this resource.
	Tags []StreamTag `pulumi:"tags"`
}

func LookupStreamOutput(ctx *pulumi.Context, args LookupStreamOutputArgs, opts ...pulumi.InvokeOption) LookupStreamResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStreamResult, error) {
			args := v.(LookupStreamArgs)
			r, err := LookupStream(ctx, &args, opts...)
			var s LookupStreamResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStreamResultOutput)
}

type LookupStreamOutputArgs struct {
	Id         pulumi.StringInput `pulumi:"id"`
	LedgerName pulumi.StringInput `pulumi:"ledgerName"`
}

func (LookupStreamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamArgs)(nil)).Elem()
}

type LookupStreamResultOutput struct{ *pulumi.OutputState }

func (LookupStreamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamResult)(nil)).Elem()
}

func (o LookupStreamResultOutput) ToLookupStreamResultOutput() LookupStreamResultOutput {
	return o
}

func (o LookupStreamResultOutput) ToLookupStreamResultOutputWithContext(ctx context.Context) LookupStreamResultOutput {
	return o
}

func (o LookupStreamResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupStreamResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupStreamResultOutput) Tags() StreamTagArrayOutput {
	return o.ApplyT(func(v LookupStreamResult) []StreamTag { return v.Tags }).(StreamTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStreamResultOutput{})
}