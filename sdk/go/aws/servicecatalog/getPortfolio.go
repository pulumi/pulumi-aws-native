// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ServiceCatalog::Portfolio
func LookupPortfolio(ctx *pulumi.Context, args *LookupPortfolioArgs, opts ...pulumi.InvokeOption) (*LookupPortfolioResult, error) {
	var rv LookupPortfolioResult
	err := ctx.Invoke("aws-native:servicecatalog:getPortfolio", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPortfolioArgs struct {
	Id string `pulumi:"id"`
}

type LookupPortfolioResult struct {
	AcceptLanguage *string        `pulumi:"acceptLanguage"`
	Description    *string        `pulumi:"description"`
	DisplayName    *string        `pulumi:"displayName"`
	Id             *string        `pulumi:"id"`
	PortfolioName  *string        `pulumi:"portfolioName"`
	ProviderName   *string        `pulumi:"providerName"`
	Tags           []PortfolioTag `pulumi:"tags"`
}

func LookupPortfolioOutput(ctx *pulumi.Context, args LookupPortfolioOutputArgs, opts ...pulumi.InvokeOption) LookupPortfolioResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPortfolioResult, error) {
			args := v.(LookupPortfolioArgs)
			r, err := LookupPortfolio(ctx, &args, opts...)
			var s LookupPortfolioResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPortfolioResultOutput)
}

type LookupPortfolioOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupPortfolioOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPortfolioArgs)(nil)).Elem()
}

type LookupPortfolioResultOutput struct{ *pulumi.OutputState }

func (LookupPortfolioResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPortfolioResult)(nil)).Elem()
}

func (o LookupPortfolioResultOutput) ToLookupPortfolioResultOutput() LookupPortfolioResultOutput {
	return o
}

func (o LookupPortfolioResultOutput) ToLookupPortfolioResultOutputWithContext(ctx context.Context) LookupPortfolioResultOutput {
	return o
}

func (o LookupPortfolioResultOutput) AcceptLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortfolioResult) *string { return v.AcceptLanguage }).(pulumi.StringPtrOutput)
}

func (o LookupPortfolioResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortfolioResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupPortfolioResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortfolioResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupPortfolioResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortfolioResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupPortfolioResultOutput) PortfolioName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortfolioResult) *string { return v.PortfolioName }).(pulumi.StringPtrOutput)
}

func (o LookupPortfolioResultOutput) ProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortfolioResult) *string { return v.ProviderName }).(pulumi.StringPtrOutput)
}

func (o LookupPortfolioResultOutput) Tags() PortfolioTagArrayOutput {
	return o.ApplyT(func(v LookupPortfolioResult) []PortfolioTag { return v.Tags }).(PortfolioTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPortfolioResultOutput{})
}