// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class AnalysisPivotTotalOptions
    {
        /// <summary>
        /// The custom label string for the total cells.
        /// </summary>
        public readonly string? CustomLabel;
        /// <summary>
        /// The cell styling options for the total of header cells.
        /// </summary>
        public readonly Outputs.AnalysisTableCellStyle? MetricHeaderCellStyle;
        /// <summary>
        /// The placement (start, end) for the total cells.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisTableTotalsPlacement? Placement;
        /// <summary>
        /// The scroll status (pinned, scrolled) for the total cells.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisTableTotalsScrollStatus? ScrollStatus;
        /// <summary>
        /// The total aggregation options for each value field.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisTotalAggregationOption> TotalAggregationOptions;
        /// <summary>
        /// The cell styling options for the total cells.
        /// </summary>
        public readonly Outputs.AnalysisTableCellStyle? TotalCellStyle;
        /// <summary>
        /// The visibility configuration for the total cells.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisVisibility? TotalsVisibility;
        /// <summary>
        /// The cell styling options for the totals of value cells.
        /// </summary>
        public readonly Outputs.AnalysisTableCellStyle? ValueCellStyle;

        [OutputConstructor]
        private AnalysisPivotTotalOptions(
            string? customLabel,

            Outputs.AnalysisTableCellStyle? metricHeaderCellStyle,

            Pulumi.AwsNative.QuickSight.AnalysisTableTotalsPlacement? placement,

            Pulumi.AwsNative.QuickSight.AnalysisTableTotalsScrollStatus? scrollStatus,

            ImmutableArray<Outputs.AnalysisTotalAggregationOption> totalAggregationOptions,

            Outputs.AnalysisTableCellStyle? totalCellStyle,

            Pulumi.AwsNative.QuickSight.AnalysisVisibility? totalsVisibility,

            Outputs.AnalysisTableCellStyle? valueCellStyle)
        {
            CustomLabel = customLabel;
            MetricHeaderCellStyle = metricHeaderCellStyle;
            Placement = placement;
            ScrollStatus = scrollStatus;
            TotalAggregationOptions = totalAggregationOptions;
            TotalCellStyle = totalCellStyle;
            TotalsVisibility = totalsVisibility;
            ValueCellStyle = valueCellStyle;
        }
    }
}
