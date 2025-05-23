// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisLineChartConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("contributionAnalysisDefaults")]
        private InputList<Inputs.AnalysisContributionAnalysisDefaultArgs>? _contributionAnalysisDefaults;

        /// <summary>
        /// The default configuration of a line chart's contribution analysis.
        /// </summary>
        public InputList<Inputs.AnalysisContributionAnalysisDefaultArgs> ContributionAnalysisDefaults
        {
            get => _contributionAnalysisDefaults ?? (_contributionAnalysisDefaults = new InputList<Inputs.AnalysisContributionAnalysisDefaultArgs>());
            set => _contributionAnalysisDefaults = value;
        }

        /// <summary>
        /// The data label configuration of a line chart.
        /// </summary>
        [Input("dataLabels")]
        public Input<Inputs.AnalysisDataLabelOptionsArgs>? DataLabels { get; set; }

        /// <summary>
        /// The options that determine the default presentation of all line series in `LineChartVisual` .
        /// </summary>
        [Input("defaultSeriesSettings")]
        public Input<Inputs.AnalysisLineChartDefaultSeriesSettingsArgs>? DefaultSeriesSettings { get; set; }

        /// <summary>
        /// The field well configuration of a line chart.
        /// </summary>
        [Input("fieldWells")]
        public Input<Inputs.AnalysisLineChartFieldWellsArgs>? FieldWells { get; set; }

        [Input("forecastConfigurations")]
        private InputList<Inputs.AnalysisForecastConfigurationArgs>? _forecastConfigurations;

        /// <summary>
        /// The forecast configuration of a line chart.
        /// </summary>
        public InputList<Inputs.AnalysisForecastConfigurationArgs> ForecastConfigurations
        {
            get => _forecastConfigurations ?? (_forecastConfigurations = new InputList<Inputs.AnalysisForecastConfigurationArgs>());
            set => _forecastConfigurations = value;
        }

        /// <summary>
        /// The general visual interactions setup for a visual.
        /// </summary>
        [Input("interactions")]
        public Input<Inputs.AnalysisVisualInteractionOptionsArgs>? Interactions { get; set; }

        /// <summary>
        /// The legend configuration of a line chart.
        /// </summary>
        [Input("legend")]
        public Input<Inputs.AnalysisLegendOptionsArgs>? Legend { get; set; }

        /// <summary>
        /// The series axis configuration of a line chart.
        /// </summary>
        [Input("primaryYAxisDisplayOptions")]
        public Input<Inputs.AnalysisLineSeriesAxisDisplayOptionsArgs>? PrimaryYAxisDisplayOptions { get; set; }

        /// <summary>
        /// The options that determine the presentation of the y-axis label.
        /// </summary>
        [Input("primaryYAxisLabelOptions")]
        public Input<Inputs.AnalysisChartAxisLabelOptionsArgs>? PrimaryYAxisLabelOptions { get; set; }

        [Input("referenceLines")]
        private InputList<Inputs.AnalysisReferenceLineArgs>? _referenceLines;

        /// <summary>
        /// The reference lines configuration of a line chart.
        /// </summary>
        public InputList<Inputs.AnalysisReferenceLineArgs> ReferenceLines
        {
            get => _referenceLines ?? (_referenceLines = new InputList<Inputs.AnalysisReferenceLineArgs>());
            set => _referenceLines = value;
        }

        /// <summary>
        /// The series axis configuration of a line chart.
        /// </summary>
        [Input("secondaryYAxisDisplayOptions")]
        public Input<Inputs.AnalysisLineSeriesAxisDisplayOptionsArgs>? SecondaryYAxisDisplayOptions { get; set; }

        /// <summary>
        /// The options that determine the presentation of the secondary y-axis label.
        /// </summary>
        [Input("secondaryYAxisLabelOptions")]
        public Input<Inputs.AnalysisChartAxisLabelOptionsArgs>? SecondaryYAxisLabelOptions { get; set; }

        [Input("series")]
        private InputList<Inputs.AnalysisSeriesItemArgs>? _series;

        /// <summary>
        /// The series item configuration of a line chart.
        /// </summary>
        public InputList<Inputs.AnalysisSeriesItemArgs> Series
        {
            get => _series ?? (_series = new InputList<Inputs.AnalysisSeriesItemArgs>());
            set => _series = value;
        }

        [Input("singleAxisOptions")]
        public Input<Inputs.AnalysisSingleAxisOptionsArgs>? SingleAxisOptions { get; set; }

        /// <summary>
        /// The small multiples setup for the visual.
        /// </summary>
        [Input("smallMultiplesOptions")]
        public Input<Inputs.AnalysisSmallMultiplesOptionsArgs>? SmallMultiplesOptions { get; set; }

        /// <summary>
        /// The sort configuration of a line chart.
        /// </summary>
        [Input("sortConfiguration")]
        public Input<Inputs.AnalysisLineChartSortConfigurationArgs>? SortConfiguration { get; set; }

        /// <summary>
        /// The tooltip configuration of a line chart.
        /// </summary>
        [Input("tooltip")]
        public Input<Inputs.AnalysisTooltipOptionsArgs>? Tooltip { get; set; }

        /// <summary>
        /// Determines the type of the line chart.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisLineChartType>? Type { get; set; }

        /// <summary>
        /// The visual palette configuration of a line chart.
        /// </summary>
        [Input("visualPalette")]
        public Input<Inputs.AnalysisVisualPaletteArgs>? VisualPalette { get; set; }

        /// <summary>
        /// The options that determine the presentation of the x-axis.
        /// </summary>
        [Input("xAxisDisplayOptions")]
        public Input<Inputs.AnalysisAxisDisplayOptionsArgs>? XAxisDisplayOptions { get; set; }

        /// <summary>
        /// The options that determine the presentation of the x-axis label.
        /// </summary>
        [Input("xAxisLabelOptions")]
        public Input<Inputs.AnalysisChartAxisLabelOptionsArgs>? XAxisLabelOptions { get; set; }

        public AnalysisLineChartConfigurationArgs()
        {
        }
        public static new AnalysisLineChartConfigurationArgs Empty => new AnalysisLineChartConfigurationArgs();
    }
}
