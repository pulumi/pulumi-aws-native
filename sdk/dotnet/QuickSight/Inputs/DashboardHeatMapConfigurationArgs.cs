// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardHeatMapConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The color options (gradient color, point of divergence) in a heat map.
        /// </summary>
        [Input("colorScale")]
        public Input<Inputs.DashboardColorScaleArgs>? ColorScale { get; set; }

        /// <summary>
        /// The label options of the column that is displayed in a heat map.
        /// </summary>
        [Input("columnLabelOptions")]
        public Input<Inputs.DashboardChartAxisLabelOptionsArgs>? ColumnLabelOptions { get; set; }

        /// <summary>
        /// The options that determine if visual data labels are displayed.
        /// </summary>
        [Input("dataLabels")]
        public Input<Inputs.DashboardDataLabelOptionsArgs>? DataLabels { get; set; }

        /// <summary>
        /// The field wells of the visual.
        /// </summary>
        [Input("fieldWells")]
        public Input<Inputs.DashboardHeatMapFieldWellsArgs>? FieldWells { get; set; }

        /// <summary>
        /// The general visual interactions setup for a visual.
        /// </summary>
        [Input("interactions")]
        public Input<Inputs.DashboardVisualInteractionOptionsArgs>? Interactions { get; set; }

        /// <summary>
        /// The legend display setup of the visual.
        /// </summary>
        [Input("legend")]
        public Input<Inputs.DashboardLegendOptionsArgs>? Legend { get; set; }

        /// <summary>
        /// The label options of the row that is displayed in a `heat map` .
        /// </summary>
        [Input("rowLabelOptions")]
        public Input<Inputs.DashboardChartAxisLabelOptionsArgs>? RowLabelOptions { get; set; }

        /// <summary>
        /// The sort configuration of a heat map.
        /// </summary>
        [Input("sortConfiguration")]
        public Input<Inputs.DashboardHeatMapSortConfigurationArgs>? SortConfiguration { get; set; }

        /// <summary>
        /// The tooltip display setup of the visual.
        /// </summary>
        [Input("tooltip")]
        public Input<Inputs.DashboardTooltipOptionsArgs>? Tooltip { get; set; }

        public DashboardHeatMapConfigurationArgs()
        {
        }
        public static new DashboardHeatMapConfigurationArgs Empty => new DashboardHeatMapConfigurationArgs();
    }
}
