// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisTreeMapConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("colorLabelOptions")]
        public Input<Inputs.AnalysisChartAxisLabelOptionsArgs>? ColorLabelOptions { get; set; }

        [Input("colorScale")]
        public Input<Inputs.AnalysisColorScaleArgs>? ColorScale { get; set; }

        [Input("dataLabels")]
        public Input<Inputs.AnalysisDataLabelOptionsArgs>? DataLabels { get; set; }

        [Input("fieldWells")]
        public Input<Inputs.AnalysisTreeMapFieldWellsArgs>? FieldWells { get; set; }

        [Input("groupLabelOptions")]
        public Input<Inputs.AnalysisChartAxisLabelOptionsArgs>? GroupLabelOptions { get; set; }

        [Input("legend")]
        public Input<Inputs.AnalysisLegendOptionsArgs>? Legend { get; set; }

        [Input("sizeLabelOptions")]
        public Input<Inputs.AnalysisChartAxisLabelOptionsArgs>? SizeLabelOptions { get; set; }

        [Input("sortConfiguration")]
        public Input<Inputs.AnalysisTreeMapSortConfigurationArgs>? SortConfiguration { get; set; }

        [Input("tooltip")]
        public Input<Inputs.AnalysisTooltipOptionsArgs>? Tooltip { get; set; }

        public AnalysisTreeMapConfigurationArgs()
        {
        }
        public static new AnalysisTreeMapConfigurationArgs Empty => new AnalysisTreeMapConfigurationArgs();
    }
}