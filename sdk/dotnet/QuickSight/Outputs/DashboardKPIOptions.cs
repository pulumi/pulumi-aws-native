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
    public sealed class DashboardKPIOptions
    {
        public readonly Outputs.DashboardComparisonConfiguration? Comparison;
        public readonly Pulumi.AwsNative.QuickSight.DashboardPrimaryValueDisplayType? PrimaryValueDisplayType;
        public readonly Outputs.DashboardFontConfiguration? PrimaryValueFontConfiguration;
        public readonly Outputs.DashboardProgressBarOptions? ProgressBar;
        public readonly Outputs.DashboardSecondaryValueOptions? SecondaryValue;
        public readonly Outputs.DashboardFontConfiguration? SecondaryValueFontConfiguration;
        public readonly Outputs.DashboardTrendArrowOptions? TrendArrows;

        [OutputConstructor]
        private DashboardKPIOptions(
            Outputs.DashboardComparisonConfiguration? comparison,

            Pulumi.AwsNative.QuickSight.DashboardPrimaryValueDisplayType? primaryValueDisplayType,

            Outputs.DashboardFontConfiguration? primaryValueFontConfiguration,

            Outputs.DashboardProgressBarOptions? progressBar,

            Outputs.DashboardSecondaryValueOptions? secondaryValue,

            Outputs.DashboardFontConfiguration? secondaryValueFontConfiguration,

            Outputs.DashboardTrendArrowOptions? trendArrows)
        {
            Comparison = comparison;
            PrimaryValueDisplayType = primaryValueDisplayType;
            PrimaryValueFontConfiguration = primaryValueFontConfiguration;
            ProgressBar = progressBar;
            SecondaryValue = secondaryValue;
            SecondaryValueFontConfiguration = secondaryValueFontConfiguration;
            TrendArrows = trendArrows;
        }
    }
}