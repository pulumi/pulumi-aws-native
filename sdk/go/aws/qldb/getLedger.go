// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package qldb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::QLDB::Ledger
func LookupLedger(ctx *pulumi.Context, args *LookupLedgerArgs, opts ...pulumi.InvokeOption) (*LookupLedgerResult, error) {
	var rv LookupLedgerResult
	err := ctx.Invoke("aws-native:qldb:getLedger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLedgerArgs struct {
	Id string `pulumi:"id"`
}

type LookupLedgerResult struct {
	DeletionProtection *bool       `pulumi:"deletionProtection"`
	Id                 *string     `pulumi:"id"`
	KmsKey             *string     `pulumi:"kmsKey"`
	PermissionsMode    *string     `pulumi:"permissionsMode"`
	Tags               []LedgerTag `pulumi:"tags"`
}

func LookupLedgerOutput(ctx *pulumi.Context, args LookupLedgerOutputArgs, opts ...pulumi.InvokeOption) LookupLedgerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLedgerResult, error) {
			args := v.(LookupLedgerArgs)
			r, err := LookupLedger(ctx, &args, opts...)
			var s LookupLedgerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLedgerResultOutput)
}

type LookupLedgerOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupLedgerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLedgerArgs)(nil)).Elem()
}

type LookupLedgerResultOutput struct{ *pulumi.OutputState }

func (LookupLedgerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLedgerResult)(nil)).Elem()
}

func (o LookupLedgerResultOutput) ToLookupLedgerResultOutput() LookupLedgerResultOutput {
	return o
}

func (o LookupLedgerResultOutput) ToLookupLedgerResultOutputWithContext(ctx context.Context) LookupLedgerResultOutput {
	return o
}

func (o LookupLedgerResultOutput) DeletionProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLedgerResult) *bool { return v.DeletionProtection }).(pulumi.BoolPtrOutput)
}

func (o LookupLedgerResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLedgerResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupLedgerResultOutput) KmsKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLedgerResult) *string { return v.KmsKey }).(pulumi.StringPtrOutput)
}

func (o LookupLedgerResultOutput) PermissionsMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLedgerResult) *string { return v.PermissionsMode }).(pulumi.StringPtrOutput)
}

func (o LookupLedgerResultOutput) Tags() LedgerTagArrayOutput {
	return o.ApplyT(func(v LookupLedgerResult) []LedgerTag { return v.Tags }).(LedgerTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLedgerResultOutput{})
}