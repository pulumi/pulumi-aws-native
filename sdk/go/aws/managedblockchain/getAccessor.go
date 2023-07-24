// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package managedblockchain

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::ManagedBlockchain::com.amazonaws.taiga.webservice.api#Accessor Resource Type
func LookupAccessor(ctx *pulumi.Context, args *LookupAccessorArgs, opts ...pulumi.InvokeOption) (*LookupAccessorResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAccessorResult
	err := ctx.Invoke("aws-native:managedblockchain:getAccessor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessorArgs struct {
	Id string `pulumi:"id"`
}

type LookupAccessorResult struct {
	Arn          *string         `pulumi:"arn"`
	BillingToken *string         `pulumi:"billingToken"`
	CreationDate *string         `pulumi:"creationDate"`
	Id           *string         `pulumi:"id"`
	Status       *AccessorStatus `pulumi:"status"`
}

func LookupAccessorOutput(ctx *pulumi.Context, args LookupAccessorOutputArgs, opts ...pulumi.InvokeOption) LookupAccessorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccessorResult, error) {
			args := v.(LookupAccessorArgs)
			r, err := LookupAccessor(ctx, &args, opts...)
			var s LookupAccessorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccessorResultOutput)
}

type LookupAccessorOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupAccessorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessorArgs)(nil)).Elem()
}

type LookupAccessorResultOutput struct{ *pulumi.OutputState }

func (LookupAccessorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessorResult)(nil)).Elem()
}

func (o LookupAccessorResultOutput) ToLookupAccessorResultOutput() LookupAccessorResultOutput {
	return o
}

func (o LookupAccessorResultOutput) ToLookupAccessorResultOutputWithContext(ctx context.Context) LookupAccessorResultOutput {
	return o
}

func (o LookupAccessorResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessorResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupAccessorResultOutput) BillingToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessorResult) *string { return v.BillingToken }).(pulumi.StringPtrOutput)
}

func (o LookupAccessorResultOutput) CreationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessorResult) *string { return v.CreationDate }).(pulumi.StringPtrOutput)
}

func (o LookupAccessorResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessorResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupAccessorResultOutput) Status() AccessorStatusPtrOutput {
	return o.ApplyT(func(v LookupAccessorResult) *AccessorStatus { return v.Status }).(AccessorStatusPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessorResultOutput{})
}