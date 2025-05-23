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
    public sealed class AnalysisVisual
    {
        /// <summary>
        /// A bar chart.
        /// 
        /// For more information, see [Using bar charts](https://docs.aws.amazon.com/quicksight/latest/user/bar-charts.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly Outputs.AnalysisBarChartVisual? BarChartVisual;
        /// <summary>
        /// A box plot.
        /// 
        /// For more information, see [Using box plots](https://docs.aws.amazon.com/quicksight/latest/user/box-plots.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly Outputs.AnalysisBoxPlotVisual? BoxPlotVisual;
        /// <summary>
        /// A combo chart.
        /// 
        /// For more information, see [Using combo charts](https://docs.aws.amazon.com/quicksight/latest/user/combo-charts.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly Outputs.AnalysisComboChartVisual? ComboChartVisual;
        /// <summary>
        /// A visual that contains custom content.
        /// 
        /// For more information, see [Using custom visual content](https://docs.aws.amazon.com/quicksight/latest/user/custom-visual-content.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly Outputs.AnalysisCustomContentVisual? CustomContentVisual;
        /// <summary>
        /// An empty visual.
        /// </summary>
        public readonly Outputs.AnalysisEmptyVisual? EmptyVisual;
        /// <summary>
        /// A filled map.
        /// 
        /// For more information, see [Creating filled maps](https://docs.aws.amazon.com/quicksight/latest/user/filled-maps.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly Outputs.AnalysisFilledMapVisual? FilledMapVisual;
        /// <summary>
        /// A funnel chart.
        /// 
        /// For more information, see [Using funnel charts](https://docs.aws.amazon.com/quicksight/latest/user/funnel-visual-content.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly Outputs.AnalysisFunnelChartVisual? FunnelChartVisual;
        /// <summary>
        /// A gauge chart.
        /// 
        /// For more information, see [Using gauge charts](https://docs.aws.amazon.com/quicksight/latest/user/gauge-chart.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly Outputs.AnalysisGaugeChartVisual? GaugeChartVisual;
        /// <summary>
        /// A geospatial map or a points on map visual.
        /// 
        /// For more information, see [Creating point maps](https://docs.aws.amazon.com/quicksight/latest/user/point-maps.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly Outputs.AnalysisGeospatialMapVisual? GeospatialMapVisual;
        /// <summary>
        /// A heat map.
        /// 
        /// For more information, see [Using heat maps](https://docs.aws.amazon.com/quicksight/latest/user/heat-map.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly Outputs.AnalysisHeatMapVisual? HeatMapVisual;
        /// <summary>
        /// A histogram.
        /// 
        /// For more information, see [Using histograms](https://docs.aws.amazon.com/quicksight/latest/user/histogram-charts.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly Outputs.AnalysisHistogramVisual? HistogramVisual;
        /// <summary>
        /// An insight visual.
        /// 
        /// For more information, see [Working with insights](https://docs.aws.amazon.com/quicksight/latest/user/computational-insights.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly Outputs.AnalysisInsightVisual? InsightVisual;
        /// <summary>
        /// A key performance indicator (KPI).
        /// 
        /// For more information, see [Using KPIs](https://docs.aws.amazon.com/quicksight/latest/user/kpi.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly Outputs.AnalysisKpiVisual? KpiVisual;
        /// <summary>
        /// The properties for a layer map visual
        /// </summary>
        public readonly Outputs.AnalysisLayerMapVisual? LayerMapVisual;
        /// <summary>
        /// A line chart.
        /// 
        /// For more information, see [Using line charts](https://docs.aws.amazon.com/quicksight/latest/user/line-charts.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly Outputs.AnalysisLineChartVisual? LineChartVisual;
        /// <summary>
        /// A pie or donut chart.
        /// 
        /// For more information, see [Using pie charts](https://docs.aws.amazon.com/quicksight/latest/user/pie-chart.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly Outputs.AnalysisPieChartVisual? PieChartVisual;
        /// <summary>
        /// A pivot table.
        /// 
        /// For more information, see [Using pivot tables](https://docs.aws.amazon.com/quicksight/latest/user/pivot-table.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly Outputs.AnalysisPivotTableVisual? PivotTableVisual;
        /// <summary>
        /// The custom plugin visual type.
        /// </summary>
        public readonly Outputs.AnalysisPluginVisual? PluginVisual;
        /// <summary>
        /// A radar chart visual.
        /// 
        /// For more information, see [Using radar charts](https://docs.aws.amazon.com/quicksight/latest/user/radar-chart.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly Outputs.AnalysisRadarChartVisual? RadarChartVisual;
        /// <summary>
        /// A sankey diagram.
        /// 
        /// For more information, see [Using Sankey diagrams](https://docs.aws.amazon.com/quicksight/latest/user/sankey-diagram.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly Outputs.AnalysisSankeyDiagramVisual? SankeyDiagramVisual;
        /// <summary>
        /// A scatter plot.
        /// 
        /// For more information, see [Using scatter plots](https://docs.aws.amazon.com/quicksight/latest/user/scatter-plot.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly Outputs.AnalysisScatterPlotVisual? ScatterPlotVisual;
        /// <summary>
        /// A table visual.
        /// 
        /// For more information, see [Using tables as visuals](https://docs.aws.amazon.com/quicksight/latest/user/tabular.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly Outputs.AnalysisTableVisual? TableVisual;
        /// <summary>
        /// A tree map.
        /// 
        /// For more information, see [Using tree maps](https://docs.aws.amazon.com/quicksight/latest/user/tree-map.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly Outputs.AnalysisTreeMapVisual? TreeMapVisual;
        /// <summary>
        /// A waterfall chart.
        /// 
        /// For more information, see [Using waterfall charts](https://docs.aws.amazon.com/quicksight/latest/user/waterfall-chart.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly Outputs.AnalysisWaterfallVisual? WaterfallVisual;
        /// <summary>
        /// A word cloud.
        /// 
        /// For more information, see [Using word clouds](https://docs.aws.amazon.com/quicksight/latest/user/word-cloud.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly Outputs.AnalysisWordCloudVisual? WordCloudVisual;

        [OutputConstructor]
        private AnalysisVisual(
            Outputs.AnalysisBarChartVisual? barChartVisual,

            Outputs.AnalysisBoxPlotVisual? boxPlotVisual,

            Outputs.AnalysisComboChartVisual? comboChartVisual,

            Outputs.AnalysisCustomContentVisual? customContentVisual,

            Outputs.AnalysisEmptyVisual? emptyVisual,

            Outputs.AnalysisFilledMapVisual? filledMapVisual,

            Outputs.AnalysisFunnelChartVisual? funnelChartVisual,

            Outputs.AnalysisGaugeChartVisual? gaugeChartVisual,

            Outputs.AnalysisGeospatialMapVisual? geospatialMapVisual,

            Outputs.AnalysisHeatMapVisual? heatMapVisual,

            Outputs.AnalysisHistogramVisual? histogramVisual,

            Outputs.AnalysisInsightVisual? insightVisual,

            Outputs.AnalysisKpiVisual? kpiVisual,

            Outputs.AnalysisLayerMapVisual? layerMapVisual,

            Outputs.AnalysisLineChartVisual? lineChartVisual,

            Outputs.AnalysisPieChartVisual? pieChartVisual,

            Outputs.AnalysisPivotTableVisual? pivotTableVisual,

            Outputs.AnalysisPluginVisual? pluginVisual,

            Outputs.AnalysisRadarChartVisual? radarChartVisual,

            Outputs.AnalysisSankeyDiagramVisual? sankeyDiagramVisual,

            Outputs.AnalysisScatterPlotVisual? scatterPlotVisual,

            Outputs.AnalysisTableVisual? tableVisual,

            Outputs.AnalysisTreeMapVisual? treeMapVisual,

            Outputs.AnalysisWaterfallVisual? waterfallVisual,

            Outputs.AnalysisWordCloudVisual? wordCloudVisual)
        {
            BarChartVisual = barChartVisual;
            BoxPlotVisual = boxPlotVisual;
            ComboChartVisual = comboChartVisual;
            CustomContentVisual = customContentVisual;
            EmptyVisual = emptyVisual;
            FilledMapVisual = filledMapVisual;
            FunnelChartVisual = funnelChartVisual;
            GaugeChartVisual = gaugeChartVisual;
            GeospatialMapVisual = geospatialMapVisual;
            HeatMapVisual = heatMapVisual;
            HistogramVisual = histogramVisual;
            InsightVisual = insightVisual;
            KpiVisual = kpiVisual;
            LayerMapVisual = layerMapVisual;
            LineChartVisual = lineChartVisual;
            PieChartVisual = pieChartVisual;
            PivotTableVisual = pivotTableVisual;
            PluginVisual = pluginVisual;
            RadarChartVisual = radarChartVisual;
            SankeyDiagramVisual = sankeyDiagramVisual;
            ScatterPlotVisual = scatterPlotVisual;
            TableVisual = tableVisual;
            TreeMapVisual = treeMapVisual;
            WaterfallVisual = waterfallVisual;
            WordCloudVisual = wordCloudVisual;
        }
    }
}
