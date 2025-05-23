// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityhub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “AWS::SecurityHub::FindingAggregator“ resource enables cross-Region aggregation. When cross-Region aggregation is enabled, you can aggregate findings, finding updates, insights, control compliance statuses, and security scores from one or more linked Regions to a single aggregation Region. You can then view and manage all of this data from the aggregation Region. For more details about cross-Region aggregation, see [Cross-Region aggregation](https://docs.aws.amazon.com/securityhub/latest/userguide/finding-aggregation.html) in the *User Guide*
//
//	This resource must be created in the Region that you want to designate as your aggregation Region.
//	Cross-Region aggregation is also a prerequisite for using [central configuration](https://docs.aws.amazon.com/securityhub/latest/userguide/central-configuration-intro.html) in ASH.
func LookupFindingAggregator(ctx *pulumi.Context, args *LookupFindingAggregatorArgs, opts ...pulumi.InvokeOption) (*LookupFindingAggregatorResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFindingAggregatorResult
	err := ctx.Invoke("aws-native:securityhub:getFindingAggregator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFindingAggregatorArgs struct {
	// The ARN of the finding aggregator. You use the finding aggregator ARN to retrieve details for, update, and delete the finding aggregator.
	FindingAggregatorArn string `pulumi:"findingAggregatorArn"`
}

type LookupFindingAggregatorResult struct {
	// The home Region. Findings generated in linked Regions are replicated and sent to the home Region.
	FindingAggregationRegion *string `pulumi:"findingAggregationRegion"`
	// The ARN of the finding aggregator. You use the finding aggregator ARN to retrieve details for, update, and delete the finding aggregator.
	FindingAggregatorArn *string `pulumi:"findingAggregatorArn"`
	// Indicates whether to aggregate findings from all of the available Regions in the current partition. Also determines whether to automatically aggregate findings from new Regions as Security Hub supports them and you opt into them.
	//  The selected option also determines how to use the Regions provided in the Regions list.
	//  In CFN, the options for this property are as follows:
	//   +  ``ALL_REGIONS`` - Indicates to aggregate findings from all of the Regions where Security Hub is enabled. When you choose this option, Security Hub also automatically aggregates findings from new Regions as Security Hub supports them and you opt into them.
	//   +  ``ALL_REGIONS_EXCEPT_SPECIFIED`` - Indicates to aggregate findings from all of the Regions where Security Hub is enabled, except for the Regions listed in the ``Regions`` parameter. When you choose this option, Security Hub also automatically aggregates findings from new Regions as Security Hub supports them and you opt into them.
	//   +  ``SPECIFIED_REGIONS`` - Indicates to aggregate findings only from the Regions listed in the ``Regions`` parameter. Security Hub does not automatically aggregate findings from new Regions.
	RegionLinkingMode *FindingAggregatorRegionLinkingMode `pulumi:"regionLinkingMode"`
	// If ``RegionLinkingMode`` is ``ALL_REGIONS_EXCEPT_SPECIFIED``, then this is a space-separated list of Regions that do not aggregate findings to the aggregation Region.
	//  If ``RegionLinkingMode`` is ``SPECIFIED_REGIONS``, then this is a space-separated list of Regions that do aggregate findings to the aggregation Region.
	Regions []string `pulumi:"regions"`
}

func LookupFindingAggregatorOutput(ctx *pulumi.Context, args LookupFindingAggregatorOutputArgs, opts ...pulumi.InvokeOption) LookupFindingAggregatorResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupFindingAggregatorResultOutput, error) {
			args := v.(LookupFindingAggregatorArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:securityhub:getFindingAggregator", args, LookupFindingAggregatorResultOutput{}, options).(LookupFindingAggregatorResultOutput), nil
		}).(LookupFindingAggregatorResultOutput)
}

type LookupFindingAggregatorOutputArgs struct {
	// The ARN of the finding aggregator. You use the finding aggregator ARN to retrieve details for, update, and delete the finding aggregator.
	FindingAggregatorArn pulumi.StringInput `pulumi:"findingAggregatorArn"`
}

func (LookupFindingAggregatorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFindingAggregatorArgs)(nil)).Elem()
}

type LookupFindingAggregatorResultOutput struct{ *pulumi.OutputState }

func (LookupFindingAggregatorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFindingAggregatorResult)(nil)).Elem()
}

func (o LookupFindingAggregatorResultOutput) ToLookupFindingAggregatorResultOutput() LookupFindingAggregatorResultOutput {
	return o
}

func (o LookupFindingAggregatorResultOutput) ToLookupFindingAggregatorResultOutputWithContext(ctx context.Context) LookupFindingAggregatorResultOutput {
	return o
}

// The home Region. Findings generated in linked Regions are replicated and sent to the home Region.
func (o LookupFindingAggregatorResultOutput) FindingAggregationRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFindingAggregatorResult) *string { return v.FindingAggregationRegion }).(pulumi.StringPtrOutput)
}

// The ARN of the finding aggregator. You use the finding aggregator ARN to retrieve details for, update, and delete the finding aggregator.
func (o LookupFindingAggregatorResultOutput) FindingAggregatorArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFindingAggregatorResult) *string { return v.FindingAggregatorArn }).(pulumi.StringPtrOutput)
}

// Indicates whether to aggregate findings from all of the available Regions in the current partition. Also determines whether to automatically aggregate findings from new Regions as Security Hub supports them and you opt into them.
//
//	The selected option also determines how to use the Regions provided in the Regions list.
//	In CFN, the options for this property are as follows:
//	 +  ``ALL_REGIONS`` - Indicates to aggregate findings from all of the Regions where Security Hub is enabled. When you choose this option, Security Hub also automatically aggregates findings from new Regions as Security Hub supports them and you opt into them.
//	 +  ``ALL_REGIONS_EXCEPT_SPECIFIED`` - Indicates to aggregate findings from all of the Regions where Security Hub is enabled, except for the Regions listed in the ``Regions`` parameter. When you choose this option, Security Hub also automatically aggregates findings from new Regions as Security Hub supports them and you opt into them.
//	 +  ``SPECIFIED_REGIONS`` - Indicates to aggregate findings only from the Regions listed in the ``Regions`` parameter. Security Hub does not automatically aggregate findings from new Regions.
func (o LookupFindingAggregatorResultOutput) RegionLinkingMode() FindingAggregatorRegionLinkingModePtrOutput {
	return o.ApplyT(func(v LookupFindingAggregatorResult) *FindingAggregatorRegionLinkingMode { return v.RegionLinkingMode }).(FindingAggregatorRegionLinkingModePtrOutput)
}

// If “RegionLinkingMode“ is “ALL_REGIONS_EXCEPT_SPECIFIED“, then this is a space-separated list of Regions that do not aggregate findings to the aggregation Region.
//
//	If ``RegionLinkingMode`` is ``SPECIFIED_REGIONS``, then this is a space-separated list of Regions that do aggregate findings to the aggregation Region.
func (o LookupFindingAggregatorResultOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFindingAggregatorResult) []string { return v.Regions }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFindingAggregatorResultOutput{})
}
