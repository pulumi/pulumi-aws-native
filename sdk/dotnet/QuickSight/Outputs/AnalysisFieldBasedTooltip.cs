// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class AnalysisFieldBasedTooltip
    {
        public readonly Pulumi.AwsNative.QuickSight.AnalysisVisibility? AggregationVisibility;
        public readonly ImmutableArray<Outputs.AnalysisTooltipItem> TooltipFields;
        public readonly Pulumi.AwsNative.QuickSight.AnalysisTooltipTitleType? TooltipTitleType;

        [OutputConstructor]
        private AnalysisFieldBasedTooltip(
            Pulumi.AwsNative.QuickSight.AnalysisVisibility? aggregationVisibility,

            ImmutableArray<Outputs.AnalysisTooltipItem> tooltipFields,

            Pulumi.AwsNative.QuickSight.AnalysisTooltipTitleType? tooltipTitleType)
        {
            AggregationVisibility = aggregationVisibility;
            TooltipFields = tooltipFields;
            TooltipTitleType = tooltipTitleType;
        }
    }
}