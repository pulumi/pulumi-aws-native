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
    public sealed class DashboardFieldBasedTooltip
    {
        public readonly Pulumi.AwsNative.QuickSight.DashboardVisibility? AggregationVisibility;
        public readonly ImmutableArray<Outputs.DashboardTooltipItem> TooltipFields;
        public readonly Pulumi.AwsNative.QuickSight.DashboardTooltipTitleType? TooltipTitleType;

        [OutputConstructor]
        private DashboardFieldBasedTooltip(
            Pulumi.AwsNative.QuickSight.DashboardVisibility? aggregationVisibility,

            ImmutableArray<Outputs.DashboardTooltipItem> tooltipFields,

            Pulumi.AwsNative.QuickSight.DashboardTooltipTitleType? tooltipTitleType)
        {
            AggregationVisibility = aggregationVisibility;
            TooltipFields = tooltipFields;
            TooltipTitleType = tooltipTitleType;
        }
    }
}