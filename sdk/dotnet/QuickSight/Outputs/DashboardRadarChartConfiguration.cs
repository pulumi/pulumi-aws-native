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
    public sealed class DashboardRadarChartConfiguration
    {
        /// <summary>
        /// Determines the visibility of the colors of alternatign bands in a radar chart.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.DashboardVisibility? AlternateBandColorsVisibility;
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
        public readonly Pulumi.AwsNative.QuickSight.DashboardRadarChartAxesRangeScale? AxesRangeScale;
        /// <summary>
        /// The base sreies settings of a radar chart.
        /// </summary>
        public readonly Outputs.DashboardRadarChartSeriesSettings? BaseSeriesSettings;
        /// <summary>
        /// The category axis of a radar chart.
        /// </summary>
        public readonly Outputs.DashboardAxisDisplayOptions? CategoryAxis;
        /// <summary>
        /// The category label options of a radar chart.
        /// </summary>
        public readonly Outputs.DashboardChartAxisLabelOptions? CategoryLabelOptions;
        /// <summary>
        /// The color axis of a radar chart.
        /// </summary>
        public readonly Outputs.DashboardAxisDisplayOptions? ColorAxis;
        /// <summary>
        /// The color label options of a radar chart.
        /// </summary>
        public readonly Outputs.DashboardChartAxisLabelOptions? ColorLabelOptions;
        /// <summary>
        /// The field well configuration of a `RadarChartVisual` .
        /// </summary>
        public readonly Outputs.DashboardRadarChartFieldWells? FieldWells;
        /// <summary>
        /// The general visual interactions setup for a visual.
        /// </summary>
        public readonly Outputs.DashboardVisualInteractionOptions? Interactions;
        /// <summary>
        /// The legend display setup of the visual.
        /// </summary>
        public readonly Outputs.DashboardLegendOptions? Legend;
        /// <summary>
        /// The shape of the radar chart.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.DashboardRadarChartShape? Shape;
        /// <summary>
        /// The sort configuration of a `RadarChartVisual` .
        /// </summary>
        public readonly Outputs.DashboardRadarChartSortConfiguration? SortConfiguration;
        /// <summary>
        /// The start angle of a radar chart's axis.
        /// </summary>
        public readonly double? StartAngle;
        /// <summary>
        /// The palette (chart color) display setup of the visual.
        /// </summary>
        public readonly Outputs.DashboardVisualPalette? VisualPalette;

        [OutputConstructor]
        private DashboardRadarChartConfiguration(
            Pulumi.AwsNative.QuickSight.DashboardVisibility? alternateBandColorsVisibility,

            string? alternateBandEvenColor,

            string? alternateBandOddColor,

            Pulumi.AwsNative.QuickSight.DashboardRadarChartAxesRangeScale? axesRangeScale,

            Outputs.DashboardRadarChartSeriesSettings? baseSeriesSettings,

            Outputs.DashboardAxisDisplayOptions? categoryAxis,

            Outputs.DashboardChartAxisLabelOptions? categoryLabelOptions,

            Outputs.DashboardAxisDisplayOptions? colorAxis,

            Outputs.DashboardChartAxisLabelOptions? colorLabelOptions,

            Outputs.DashboardRadarChartFieldWells? fieldWells,

            Outputs.DashboardVisualInteractionOptions? interactions,

            Outputs.DashboardLegendOptions? legend,

            Pulumi.AwsNative.QuickSight.DashboardRadarChartShape? shape,

            Outputs.DashboardRadarChartSortConfiguration? sortConfiguration,

            double? startAngle,

            Outputs.DashboardVisualPalette? visualPalette)
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
