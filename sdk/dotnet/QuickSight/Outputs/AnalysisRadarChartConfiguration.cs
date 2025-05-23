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
    public sealed class AnalysisRadarChartConfiguration
    {
        /// <summary>
        /// Determines the visibility of the colors of alternatign bands in a radar chart.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisVisibility? AlternateBandColorsVisibility;
        /// <summary>
        /// The color of the even-numbered alternate bands of a radar chart.
        /// </summary>
        public readonly string? AlternateBandEvenColor;
        /// <summary>
        /// The color of the odd-numbered alternate bands of a radar chart.
        /// </summary>
        public readonly string? AlternateBandOddColor;
        /// <summary>
        /// The axis behavior options of a radar chart.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisRadarChartAxesRangeScale? AxesRangeScale;
        /// <summary>
        /// The base sreies settings of a radar chart.
        /// </summary>
        public readonly Outputs.AnalysisRadarChartSeriesSettings? BaseSeriesSettings;
        /// <summary>
        /// The category axis of a radar chart.
        /// </summary>
        public readonly Outputs.AnalysisAxisDisplayOptions? CategoryAxis;
        /// <summary>
        /// The category label options of a radar chart.
        /// </summary>
        public readonly Outputs.AnalysisChartAxisLabelOptions? CategoryLabelOptions;
        /// <summary>
        /// The color axis of a radar chart.
        /// </summary>
        public readonly Outputs.AnalysisAxisDisplayOptions? ColorAxis;
        /// <summary>
        /// The color label options of a radar chart.
        /// </summary>
        public readonly Outputs.AnalysisChartAxisLabelOptions? ColorLabelOptions;
        /// <summary>
        /// The field well configuration of a `RadarChartVisual` .
        /// </summary>
        public readonly Outputs.AnalysisRadarChartFieldWells? FieldWells;
        /// <summary>
        /// The general visual interactions setup for a visual.
        /// </summary>
        public readonly Outputs.AnalysisVisualInteractionOptions? Interactions;
        /// <summary>
        /// The legend display setup of the visual.
        /// </summary>
        public readonly Outputs.AnalysisLegendOptions? Legend;
        /// <summary>
        /// The shape of the radar chart.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisRadarChartShape? Shape;
        /// <summary>
        /// The sort configuration of a `RadarChartVisual` .
        /// </summary>
        public readonly Outputs.AnalysisRadarChartSortConfiguration? SortConfiguration;
        /// <summary>
        /// The start angle of a radar chart's axis.
        /// </summary>
        public readonly double? StartAngle;
        /// <summary>
        /// The palette (chart color) display setup of the visual.
        /// </summary>
        public readonly Outputs.AnalysisVisualPalette? VisualPalette;

        [OutputConstructor]
        private AnalysisRadarChartConfiguration(
            Pulumi.AwsNative.QuickSight.AnalysisVisibility? alternateBandColorsVisibility,

            string? alternateBandEvenColor,

            string? alternateBandOddColor,

            Pulumi.AwsNative.QuickSight.AnalysisRadarChartAxesRangeScale? axesRangeScale,

            Outputs.AnalysisRadarChartSeriesSettings? baseSeriesSettings,

            Outputs.AnalysisAxisDisplayOptions? categoryAxis,

            Outputs.AnalysisChartAxisLabelOptions? categoryLabelOptions,

            Outputs.AnalysisAxisDisplayOptions? colorAxis,

            Outputs.AnalysisChartAxisLabelOptions? colorLabelOptions,

            Outputs.AnalysisRadarChartFieldWells? fieldWells,

            Outputs.AnalysisVisualInteractionOptions? interactions,

            Outputs.AnalysisLegendOptions? legend,

            Pulumi.AwsNative.QuickSight.AnalysisRadarChartShape? shape,

            Outputs.AnalysisRadarChartSortConfiguration? sortConfiguration,

            double? startAngle,

            Outputs.AnalysisVisualPalette? visualPalette)
        {
            AlternateBandColorsVisibility = alternateBandColorsVisibility;
            AlternateBandEvenColor = alternateBandEvenColor;
            AlternateBandOddColor = alternateBandOddColor;
            AxesRangeScale = axesRangeScale;
            BaseSeriesSettings = baseSeriesSettings;
            CategoryAxis = categoryAxis;
            CategoryLabelOptions = categoryLabelOptions;
            ColorAxis = colorAxis;
            ColorLabelOptions = colorLabelOptions;
            FieldWells = fieldWells;
            Interactions = interactions;
            Legend = legend;
            Shape = shape;
            SortConfiguration = sortConfiguration;
            StartAngle = startAngle;
            VisualPalette = visualPalette;
        }
    }
}
