// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package xray

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This schema provides construct and validation rules for AWS-XRay TransactionSearchConfig resource parameters.
func LookupTransactionSearchConfig(ctx *pulumi.Context, args *LookupTransactionSearchConfigArgs, opts ...pulumi.InvokeOption) (*LookupTransactionSearchConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTransactionSearchConfigResult
	err := ctx.Invoke("aws-native:xray:getTransactionSearchConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTransactionSearchConfigArgs struct {
	AccountId string `pulumi:"accountId"`
}

type LookupTransactionSearchConfigResult struct {
	AccountId          *string  `pulumi:"accountId"`
	IndexingPercentage *float64 `pulumi:"indexingPercentage"`
}

func LookupTransactionSearchConfigOutput(ctx *pulumi.Context, args LookupTransactionSearchConfigOutputArgs, opts ...pulumi.InvokeOption) LookupTransactionSearchConfigResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupTransactionSearchConfigResultOutput, error) {
			args := v.(LookupTransactionSearchConfigArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:xray:getTransactionSearchConfig", args, LookupTransactionSearchConfigResultOutput{}, options).(LookupTransactionSearchConfigResultOutput), nil
		}).(LookupTransactionSearchConfigResultOutput)
}

type LookupTransactionSearchConfigOutputArgs struct {
	AccountId pulumi.StringInput `pulumi:"accountId"`
}

func (LookupTransactionSearchConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTransactionSearchConfigArgs)(nil)).Elem()
}

type LookupTransactionSearchConfigResultOutput struct{ *pulumi.OutputState }

func (LookupTransactionSearchConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTransactionSearchConfigResult)(nil)).Elem()
}

func (o LookupTransactionSearchConfigResultOutput) ToLookupTransactionSearchConfigResultOutput() LookupTransactionSearchConfigResultOutput {
	return o
}

func (o LookupTransactionSearchConfigResultOutput) ToLookupTransactionSearchConfigResultOutputWithContext(ctx context.Context) LookupTransactionSearchConfigResultOutput {
	return o
}

func (o LookupTransactionSearchConfigResultOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTransactionSearchConfigResult) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

func (o LookupTransactionSearchConfigResultOutput) IndexingPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupTransactionSearchConfigResult) *float64 { return v.IndexingPercentage }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTransactionSearchConfigResultOutput{})
}
