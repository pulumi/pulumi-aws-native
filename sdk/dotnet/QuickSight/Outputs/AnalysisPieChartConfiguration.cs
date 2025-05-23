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
    public sealed class AnalysisPieChartConfiguration
    {
        /// <summary>
        /// The label options of the group/color that is displayed in a pie chart.
        /// </summary>
        public readonly Outputs.AnalysisChartAxisLabelOptions? CategoryLabelOptions;
        /// <summary>
        /// The contribution analysis (anomaly configuration) setup of the visual.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisContributionAnalysisDefault> ContributionAnalysisDefaults;
        /// <summary>
        /// The options that determine if visual data labels are displayed.
        /// </summary>
        public readonly Outputs.AnalysisDataLabelOptions? DataLabels;
        /// <summary>
        /// The options that determine the shape of the chart. This option determines whether the chart is a pie chart or a donut chart.
        /// </summary>
        public readonly Outputs.AnalysisDonutOptions? DonutOptions;
        /// <summary>
        /// The field wells of the visual.
        /// </summary>
        public readonly Outputs.AnalysisPieChartFieldWells? FieldWells;
        /// <summary>
        /// The general visual interactions setup for a visual.
        /// </summary>
        public readonly Outputs.AnalysisVisualInteractionOptions? Interactions;
        /// <summary>
        /// The legend display setup of the visual.
        /// </summary>
        public readonly Outputs.AnalysisLegendOptions? Legend;
        /// <summary>
        /// The small multiples setup for the visual.
        /// </summary>
        public readonly Outputs.AnalysisSmallMultiplesOptions? SmallMultiplesOptions;
        /// <summary>
        /// The sort configuration of a pie chart.
        /// </summary>
        public readonly Outputs.AnalysisPieChartSortConfiguration? SortConfiguration;
        /// <summary>
        /// The tooltip display setup of the visual.
        /// </summary>
        public readonly Outputs.AnalysisTooltipOptions? Tooltip;
        /// <summary>
        /// The label options for the value that is displayed in a pie chart.
        /// </summary>
        public readonly Outputs.AnalysisChartAxisLabelOptions? ValueLabelOptions;
        /// <summary>
        /// The palette (chart color) display setup of the visual.
        /// </summary>
        public readonly Outputs.AnalysisVisualPalette? VisualPalette;

        [OutputConstructor]
        private AnalysisPieChartConfiguration(
            Outputs.AnalysisChartAxisLabelOptions? categoryLabelOptions,

            ImmutableArray<Outputs.AnalysisContributionAnalysisDefault> contributionAnalysisDefaults,

            Outputs.AnalysisDataLabelOptions? dataLabels,

            Outputs.AnalysisDonutOptions? donutOptions,

            Outputs.AnalysisPieChartFieldWells? fieldWells,

            Outputs.AnalysisVisualInteractionOptions? interactions,

            Outputs.AnalysisLegendOptions? legend,

            Outputs.AnalysisSmallMultiplesOptions? smallMultiplesOptions,

            Outputs.AnalysisPieChartSortConfiguration? sortConfiguration,

            Outputs.AnalysisTooltipOptions? tooltip,

            Outputs.AnalysisChartAxisLabelOptions? valueLabelOptions,

            Outputs.AnalysisVisualPalette? visualPalette)
        {
            CategoryLabelOptions = categoryLabelOptions;
            ContributionAnalysisDefaults = contributionAnalysisDefaults;
            DataLabels = dataLabels;
            DonutOptions = donutOptions;
            FieldWells = fieldWells;
            Interactions = interactions;
            Legend = legend;
            SmallMultiplesOptions = smallMultiplesOptions;
            SortConfiguration = sortConfiguration;
            Tooltip = tooltip;
            ValueLabelOptions = valueLabelOptions;
            VisualPalette = visualPalette;
        }
    }
}
