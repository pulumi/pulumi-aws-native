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
    public sealed class AnalysisKpiOptions
    {
        /// <summary>
        /// The comparison configuration of a KPI visual.
        /// </summary>
        public readonly Outputs.AnalysisComparisonConfiguration? Comparison;
        /// <summary>
        /// The options that determine the primary value display type.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisPrimaryValueDisplayType? PrimaryValueDisplayType;
        /// <summary>
        /// The options that determine the primary value font configuration.
        /// </summary>
        public readonly Outputs.AnalysisFontConfiguration? PrimaryValueFontConfiguration;
        /// <summary>
        /// The options that determine the presentation of the progress bar of a KPI visual.
        /// </summary>
        public readonly Outputs.AnalysisProgressBarOptions? ProgressBar;
        /// <summary>
        /// The options that determine the presentation of the secondary value of a KPI visual.
        /// </summary>
        public readonly Outputs.AnalysisSecondaryValueOptions? SecondaryValue;
        /// <summary>
        /// The options that determine the secondary value font configuration.
        /// </summary>
        public readonly Outputs.AnalysisFontConfiguration? SecondaryValueFontConfiguration;
        /// <summary>
        /// The options that determine the visibility, color, type, and tooltip visibility of the sparkline of a KPI visual.
        /// </summary>
        public readonly Outputs.AnalysisKpiSparklineOptions? Sparkline;
        /// <summary>
        /// The options that determine the presentation of trend arrows in a KPI visual.
        /// </summary>
        public readonly Outputs.AnalysisTrendArrowOptions? TrendArrows;
        /// <summary>
        /// The options that determine the layout a KPI visual.
        /// </summary>
        public readonly Outputs.AnalysisKpiVisualLayoutOptions? VisualLayoutOptions;

        [OutputConstructor]
        private AnalysisKpiOptions(
            Outputs.AnalysisComparisonConfiguration? comparison,

            Pulumi.AwsNative.QuickSight.AnalysisPrimaryValueDisplayType? primaryValueDisplayType,

            Outputs.AnalysisFontConfiguration? primaryValueFontConfiguration,

            Outputs.AnalysisProgressBarOptions? progressBar,

            Outputs.AnalysisSecondaryValueOptions? secondaryValue,

            Outputs.AnalysisFontConfiguration? secondaryValueFontConfiguration,

            Outputs.AnalysisKpiSparklineOptions? sparkline,

            Outputs.AnalysisTrendArrowOptions? trendArrows,

            Outputs.AnalysisKpiVisualLayoutOptions? visualLayoutOptions)
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
