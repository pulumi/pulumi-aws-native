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
    public sealed class AnalysisHeatMapConfiguration
    {
        public readonly Outputs.AnalysisColorScale? ColorScale;
        public readonly Outputs.AnalysisChartAxisLabelOptions? ColumnLabelOptions;
        public readonly Outputs.AnalysisDataLabelOptions? DataLabels;
        public readonly Outputs.AnalysisHeatMapFieldWells? FieldWells;
        public readonly Outputs.AnalysisLegendOptions? Legend;
        public readonly Outputs.AnalysisChartAxisLabelOptions? RowLabelOptions;
        public readonly Outputs.AnalysisHeatMapSortConfiguration? SortConfiguration;
        public readonly Outputs.AnalysisTooltipOptions? Tooltip;

        [OutputConstructor]
        private AnalysisHeatMapConfiguration(
            Outputs.AnalysisColorScale? colorScale,

            Outputs.AnalysisChartAxisLabelOptions? columnLabelOptions,

            Outputs.AnalysisDataLabelOptions? dataLabels,

            Outputs.AnalysisHeatMapFieldWells? fieldWells,

            Outputs.AnalysisLegendOptions? legend,

            Outputs.AnalysisChartAxisLabelOptions? rowLabelOptions,

            Outputs.AnalysisHeatMapSortConfiguration? sortConfiguration,

            Outputs.AnalysisTooltipOptions? tooltip)
        {
            ColorScale = colorScale;
            ColumnLabelOptions = columnLabelOptions;
            DataLabels = dataLabels;
            FieldWells = fieldWells;
            Legend = legend;
            RowLabelOptions = rowLabelOptions;
            SortConfiguration = sortConfiguration;
            Tooltip = tooltip;
        }
    }
}