// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateHeatMapConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("colorScale")]
        public Input<Inputs.TemplateColorScaleArgs>? ColorScale { get; set; }

        [Input("columnLabelOptions")]
        public Input<Inputs.TemplateChartAxisLabelOptionsArgs>? ColumnLabelOptions { get; set; }

        [Input("dataLabels")]
        public Input<Inputs.TemplateDataLabelOptionsArgs>? DataLabels { get; set; }

        [Input("fieldWells")]
        public Input<Inputs.TemplateHeatMapFieldWellsArgs>? FieldWells { get; set; }

        [Input("legend")]
        public Input<Inputs.TemplateLegendOptionsArgs>? Legend { get; set; }

        [Input("rowLabelOptions")]
        public Input<Inputs.TemplateChartAxisLabelOptionsArgs>? RowLabelOptions { get; set; }

        [Input("sortConfiguration")]
        public Input<Inputs.TemplateHeatMapSortConfigurationArgs>? SortConfiguration { get; set; }

        [Input("tooltip")]
        public Input<Inputs.TemplateTooltipOptionsArgs>? Tooltip { get; set; }

        public TemplateHeatMapConfigurationArgs()
        {
        }
        public static new TemplateHeatMapConfigurationArgs Empty => new TemplateHeatMapConfigurationArgs();
    }
}