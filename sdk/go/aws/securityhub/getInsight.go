// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityhub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::SecurityHub::Insight resource represents the AWS Security Hub Insight in your account. An AWS Security Hub insight is a collection of related findings.
func LookupInsight(ctx *pulumi.Context, args *LookupInsightArgs, opts ...pulumi.InvokeOption) (*LookupInsightResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInsightResult
	err := ctx.Invoke("aws-native:securityhub:getInsight", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInsightArgs struct {
	// The ARN of a Security Hub insight
	InsightArn string `pulumi:"insightArn"`
}

type LookupInsightResult struct {
	// One or more attributes used to filter the findings included in the insight
	Filters *InsightAwsSecurityFindingFilters `pulumi:"filters"`
	// The grouping attribute for the insight's findings
	GroupByAttribute *string `pulumi:"groupByAttribute"`
	// The ARN of a Security Hub insight
	InsightArn *string `pulumi:"insightArn"`
	// The name of a Security Hub insight
	Name *string `pulumi:"name"`
}

func LookupInsightOutput(ctx *pulumi.Context, args LookupInsightOutputArgs, opts ...pulumi.InvokeOption) LookupInsightResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInsightResult, error) {
			args := v.(LookupInsightArgs)
			r, err := LookupInsight(ctx, &args, opts...)
			var s LookupInsightResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInsightResultOutput)
}

type LookupInsightOutputArgs struct {
	// The ARN of a Security Hub insight
	InsightArn pulumi.StringInput `pulumi:"insightArn"`
}

func (LookupInsightOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInsightArgs)(nil)).Elem()
}

type LookupInsightResultOutput struct{ *pulumi.OutputState }

func (LookupInsightResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInsightResult)(nil)).Elem()
}

func (o LookupInsightResultOutput) ToLookupInsightResultOutput() LookupInsightResultOutput {
	return o
}

func (o LookupInsightResultOutput) ToLookupInsightResultOutputWithContext(ctx context.Context) LookupInsightResultOutput {
	return o
}

// One or more attributes used to filter the findings included in the insight
func (o LookupInsightResultOutput) Filters() InsightAwsSecurityFindingFiltersPtrOutput {
	return o.ApplyT(func(v LookupInsightResult) *InsightAwsSecurityFindingFilters { return v.Filters }).(InsightAwsSecurityFindingFiltersPtrOutput)
}

// The grouping attribute for the insight's findings
func (o LookupInsightResultOutput) GroupByAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInsightResult) *string { return v.GroupByAttribute }).(pulumi.StringPtrOutput)
}

// The ARN of a Security Hub insight
func (o LookupInsightResultOutput) InsightArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInsightResult) *string { return v.InsightArn }).(pulumi.StringPtrOutput)
}

// The name of a Security Hub insight
func (o LookupInsightResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInsightResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInsightResultOutput{})
}