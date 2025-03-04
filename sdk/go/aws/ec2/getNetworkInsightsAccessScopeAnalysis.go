// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::EC2::NetworkInsightsAccessScopeAnalysis
func LookupNetworkInsightsAccessScopeAnalysis(ctx *pulumi.Context, args *LookupNetworkInsightsAccessScopeAnalysisArgs, opts ...pulumi.InvokeOption) (*LookupNetworkInsightsAccessScopeAnalysisResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkInsightsAccessScopeAnalysisResult
	err := ctx.Invoke("aws-native:ec2:getNetworkInsightsAccessScopeAnalysis", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkInsightsAccessScopeAnalysisArgs struct {
	// The ID of the Network Access Scope analysis.
	NetworkInsightsAccessScopeAnalysisId string `pulumi:"networkInsightsAccessScopeAnalysisId"`
}

type LookupNetworkInsightsAccessScopeAnalysisResult struct {
	// The number of network interfaces analyzed.
	AnalyzedEniCount *int `pulumi:"analyzedEniCount"`
	// The end date of the analysis.
	EndDate *string `pulumi:"endDate"`
	// Indicates whether there are findings (true | false | unknown).
	FindingsFound *NetworkInsightsAccessScopeAnalysisFindingsFound `pulumi:"findingsFound"`
	// The ARN of the Network Access Scope analysis.
	NetworkInsightsAccessScopeAnalysisArn *string `pulumi:"networkInsightsAccessScopeAnalysisArn"`
	// The ID of the Network Access Scope analysis.
	NetworkInsightsAccessScopeAnalysisId *string `pulumi:"networkInsightsAccessScopeAnalysisId"`
	// The start date of the analysis.
	StartDate *string `pulumi:"startDate"`
	// The status of the analysis (running | succeeded | failed).
	Status *NetworkInsightsAccessScopeAnalysisStatus `pulumi:"status"`
	// The status message.
	StatusMessage *string `pulumi:"statusMessage"`
	// The tags.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupNetworkInsightsAccessScopeAnalysisOutput(ctx *pulumi.Context, args LookupNetworkInsightsAccessScopeAnalysisOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkInsightsAccessScopeAnalysisResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNetworkInsightsAccessScopeAnalysisResultOutput, error) {
			args := v.(LookupNetworkInsightsAccessScopeAnalysisArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:ec2:getNetworkInsightsAccessScopeAnalysis", args, LookupNetworkInsightsAccessScopeAnalysisResultOutput{}, options).(LookupNetworkInsightsAccessScopeAnalysisResultOutput), nil
		}).(LookupNetworkInsightsAccessScopeAnalysisResultOutput)
}

type LookupNetworkInsightsAccessScopeAnalysisOutputArgs struct {
	// The ID of the Network Access Scope analysis.
	NetworkInsightsAccessScopeAnalysisId pulumi.StringInput `pulumi:"networkInsightsAccessScopeAnalysisId"`
}

func (LookupNetworkInsightsAccessScopeAnalysisOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkInsightsAccessScopeAnalysisArgs)(nil)).Elem()
}

type LookupNetworkInsightsAccessScopeAnalysisResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkInsightsAccessScopeAnalysisResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkInsightsAccessScopeAnalysisResult)(nil)).Elem()
}

func (o LookupNetworkInsightsAccessScopeAnalysisResultOutput) ToLookupNetworkInsightsAccessScopeAnalysisResultOutput() LookupNetworkInsightsAccessScopeAnalysisResultOutput {
	return o
}

func (o LookupNetworkInsightsAccessScopeAnalysisResultOutput) ToLookupNetworkInsightsAccessScopeAnalysisResultOutputWithContext(ctx context.Context) LookupNetworkInsightsAccessScopeAnalysisResultOutput {
	return o
}

// The number of network interfaces analyzed.
func (o LookupNetworkInsightsAccessScopeAnalysisResultOutput) AnalyzedEniCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupNetworkInsightsAccessScopeAnalysisResult) *int { return v.AnalyzedEniCount }).(pulumi.IntPtrOutput)
}

// The end date of the analysis.
func (o LookupNetworkInsightsAccessScopeAnalysisResultOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInsightsAccessScopeAnalysisResult) *string { return v.EndDate }).(pulumi.StringPtrOutput)
}

// Indicates whether there are findings (true | false | unknown).
func (o LookupNetworkInsightsAccessScopeAnalysisResultOutput) FindingsFound() NetworkInsightsAccessScopeAnalysisFindingsFoundPtrOutput {
	return o.ApplyT(func(v LookupNetworkInsightsAccessScopeAnalysisResult) *NetworkInsightsAccessScopeAnalysisFindingsFound {
		return v.FindingsFound
	}).(NetworkInsightsAccessScopeAnalysisFindingsFoundPtrOutput)
}

// The ARN of the Network Access Scope analysis.
func (o LookupNetworkInsightsAccessScopeAnalysisResultOutput) NetworkInsightsAccessScopeAnalysisArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInsightsAccessScopeAnalysisResult) *string {
		return v.NetworkInsightsAccessScopeAnalysisArn
	}).(pulumi.StringPtrOutput)
}

// The ID of the Network Access Scope analysis.
func (o LookupNetworkInsightsAccessScopeAnalysisResultOutput) NetworkInsightsAccessScopeAnalysisId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInsightsAccessScopeAnalysisResult) *string {
		return v.NetworkInsightsAccessScopeAnalysisId
	}).(pulumi.StringPtrOutput)
}

// The start date of the analysis.
func (o LookupNetworkInsightsAccessScopeAnalysisResultOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInsightsAccessScopeAnalysisResult) *string { return v.StartDate }).(pulumi.StringPtrOutput)
}

// The status of the analysis (running | succeeded | failed).
func (o LookupNetworkInsightsAccessScopeAnalysisResultOutput) Status() NetworkInsightsAccessScopeAnalysisStatusPtrOutput {
	return o.ApplyT(func(v LookupNetworkInsightsAccessScopeAnalysisResult) *NetworkInsightsAccessScopeAnalysisStatus {
		return v.Status
	}).(NetworkInsightsAccessScopeAnalysisStatusPtrOutput)
}

// The status message.
func (o LookupNetworkInsightsAccessScopeAnalysisResultOutput) StatusMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInsightsAccessScopeAnalysisResult) *string { return v.StatusMessage }).(pulumi.StringPtrOutput)
}

// The tags.
func (o LookupNetworkInsightsAccessScopeAnalysisResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupNetworkInsightsAccessScopeAnalysisResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkInsightsAccessScopeAnalysisResultOutput{})
}
