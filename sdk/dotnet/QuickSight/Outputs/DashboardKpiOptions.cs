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
    public sealed class DashboardKpiOptions
    {
        public readonly Outputs.DashboardComparisonConfiguration? Comparison;
        public readonly Pulumi.AwsNative.QuickSight.DashboardPrimaryValueDisplayType? PrimaryValueDisplayType;
        public readonly Outputs.DashboardFontConfiguration? PrimaryValueFontConfiguration;
        public readonly Outputs.DashboardProgressBarOptions? ProgressBar;
        public readonly Outputs.DashboardSecondaryValueOptions? SecondaryValue;
        public readonly Outputs.DashboardFontConfiguration? SecondaryValueFontConfiguration;
        public readonly Outputs.DashboardKpiSparklineOptions? Sparkline;
        public readonly Outputs.DashboardTrendArrowOptions? TrendArrows;
        public readonly Outputs.DashboardKpiVisualLayoutOptions? VisualLayoutOptions;

        [OutputConstructor]
        private DashboardKpiOptions(
            Outputs.DashboardComparisonConfiguration? comparison,

            Pulumi.AwsNative.QuickSight.DashboardPrimaryValueDisplayType? primaryValueDisplayType,

            Outputs.DashboardFontConfiguration? primaryValueFontConfiguration,

            Outputs.DashboardProgressBarOptions? progressBar,

            Outputs.DashboardSecondaryValueOptions? secondaryValue,

            Outputs.DashboardFontConfiguration? secondaryValueFontConfiguration,

            Outputs.DashboardKpiSparklineOptions? sparkline,

            Outputs.DashboardTrendArrowOptions? trendArrows,

            Outputs.DashboardKpiVisualLayoutOptions? visualLayoutOptions)
        {
            Comparison = comparison;
            PrimaryValueDisplayType = primaryValueDisplayType;
            PrimaryValueFontConfiguration = primaryValueFontConfiguration;
            ProgressBar = progressBar;
            SecondaryValue = secondaryValue;
            SecondaryValueFontConfiguration = secondaryValueFontConfiguration;
            Sparkline = sparkline;
            TrendArrows = trendArrows;
            VisualLayoutOptions = visualLayoutOptions;
        }
    }
}