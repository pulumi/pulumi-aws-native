// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateVisualArgs : global::Pulumi.ResourceArgs
    {
        [Input("barChartVisual")]
        public Input<Inputs.TemplateBarChartVisualArgs>? BarChartVisual { get; set; }

        [Input("boxPlotVisual")]
        public Input<Inputs.TemplateBoxPlotVisualArgs>? BoxPlotVisual { get; set; }

        [Input("comboChartVisual")]
        public Input<Inputs.TemplateComboChartVisualArgs>? ComboChartVisual { get; set; }

        [Input("customContentVisual")]
        public Input<Inputs.TemplateCustomContentVisualArgs>? CustomContentVisual { get; set; }

        [Input("emptyVisual")]
        public Input<Inputs.TemplateEmptyVisualArgs>? EmptyVisual { get; set; }

        [Input("filledMapVisual")]
        public Input<Inputs.TemplateFilledMapVisualArgs>? FilledMapVisual { get; set; }

        [Input("funnelChartVisual")]
        public Input<Inputs.TemplateFunnelChartVisualArgs>? FunnelChartVisual { get; set; }

        [Input("gaugeChartVisual")]
        public Input<Inputs.TemplateGaugeChartVisualArgs>? GaugeChartVisual { get; set; }

        [Input("geospatialMapVisual")]
        public Input<Inputs.TemplateGeospatialMapVisualArgs>? GeospatialMapVisual { get; set; }

        [Input("heatMapVisual")]
        public Input<Inputs.TemplateHeatMapVisualArgs>? HeatMapVisual { get; set; }

        [Input("histogramVisual")]
        public Input<Inputs.TemplateHistogramVisualArgs>? HistogramVisual { get; set; }

        [Input("insightVisual")]
        public Input<Inputs.TemplateInsightVisualArgs>? InsightVisual { get; set; }

        [Input("kPIVisual")]
        public Input<Inputs.TemplateKPIVisualArgs>? KPIVisual { get; set; }

        [Input("lineChartVisual")]
        public Input<Inputs.TemplateLineChartVisualArgs>? LineChartVisual { get; set; }

        [Input("pieChartVisual")]
        public Input<Inputs.TemplatePieChartVisualArgs>? PieChartVisual { get; set; }

        [Input("pivotTableVisual")]
        public Input<Inputs.TemplatePivotTableVisualArgs>? PivotTableVisual { get; set; }

        [Input("radarChartVisual")]
        public Input<Inputs.TemplateRadarChartVisualArgs>? RadarChartVisual { get; set; }

        [Input("sankeyDiagramVisual")]
        public Input<Inputs.TemplateSankeyDiagramVisualArgs>? SankeyDiagramVisual { get; set; }

        [Input("scatterPlotVisual")]
        public Input<Inputs.TemplateScatterPlotVisualArgs>? ScatterPlotVisual { get; set; }

        [Input("tableVisual")]
        public Input<Inputs.TemplateTableVisualArgs>? TableVisual { get; set; }

        [Input("treeMapVisual")]
        public Input<Inputs.TemplateTreeMapVisualArgs>? TreeMapVisual { get; set; }

        [Input("waterfallVisual")]
        public Input<Inputs.TemplateWaterfallVisualArgs>? WaterfallVisual { get; set; }

        [Input("wordCloudVisual")]
        public Input<Inputs.TemplateWordCloudVisualArgs>? WordCloudVisual { get; set; }

        public TemplateVisualArgs()
        {
        }
        public static new TemplateVisualArgs Empty => new TemplateVisualArgs();
    }
}