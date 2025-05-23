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
    public sealed class AnalysisSubtotalOptions
    {
        /// <summary>
        /// The custom label string for the subtotal cells.
        /// </summary>
        public readonly string? CustomLabel;
        /// <summary>
        /// The field level (all, custom, last) for the subtotal cells.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisPivotTableSubtotalLevel? FieldLevel;
        /// <summary>
        /// The optional configuration of subtotal cells.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisPivotTableFieldSubtotalOptions> FieldLevelOptions;
        /// <summary>
        /// The cell styling options for the subtotals of header cells.
        /// </summary>
        public readonly Outputs.AnalysisTableCellStyle? MetricHeaderCellStyle;
        /// <summary>
        /// The style targets options for subtotals.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisTableStyleTarget> StyleTargets;
        /// <summary>
        /// The cell styling options for the subtotal cells.
        /// </summary>
        public readonly Outputs.AnalysisTableCellStyle? TotalCellStyle;
        /// <summary>
        /// The visibility configuration for the subtotal cells.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisVisibility? TotalsVisibility;
        /// <summary>
        /// The cell styling options for the subtotals of value cells.
        /// </summary>
        public readonly Outputs.AnalysisTableCellStyle? ValueCellStyle;

        [OutputConstructor]
        private AnalysisSubtotalOptions(
            string? customLabel,

            Pulumi.AwsNative.QuickSight.AnalysisPivotTableSubtotalLevel? fieldLevel,

            ImmutableArray<Outputs.AnalysisPivotTableFieldSubtotalOptions> fieldLevelOptions,

            Outputs.AnalysisTableCellStyle? metricHeaderCellStyle,

            ImmutableArray<Outputs.AnalysisTableStyleTarget> styleTargets,

            Outputs.AnalysisTableCellStyle? totalCellStyle,

            Pulumi.AwsNative.QuickSight.AnalysisVisibility? totalsVisibility,

            Outputs.AnalysisTableCellStyle? valueCellStyle)
        {
            CustomLabel = customLabel;
            FieldLevel = fieldLevel;
            FieldLevelOptions = fieldLevelOptions;
            MetricHeaderCellStyle = metricHeaderCellStyle;
            StyleTargets = styleTargets;
            TotalCellStyle = totalCellStyle;
            TotalsVisibility = totalsVisibility;
            ValueCellStyle = valueCellStyle;
        }
    }
}
