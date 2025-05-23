// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardFilledMapConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The field wells of the visual.
        /// </summary>
        [Input("fieldWells")]
        public Input<Inputs.DashboardFilledMapFieldWellsArgs>? FieldWells { get; set; }

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
        /// The map style options of the filled map visual.
        /// </summary>
        [Input("mapStyleOptions")]
        public Input<Inputs.DashboardGeospatialMapStyleOptionsArgs>? MapStyleOptions { get; set; }

        /// <summary>
        /// The sort configuration of a `FilledMapVisual` .
        /// </summary>
        [Input("sortConfiguration")]
        public Input<Inputs.DashboardFilledMapSortConfigurationArgs>? SortConfiguration { get; set; }

        /// <summary>
        /// The tooltip display setup of the visual.
        /// </summary>
        [Input("tooltip")]
        public Input<Inputs.DashboardTooltipOptionsArgs>? Tooltip { get; set; }

        /// <summary>
        /// The window options of the filled map visual.
        /// </summary>
        [Input("windowOptions")]
        public Input<Inputs.DashboardGeospatialWindowOptionsArgs>? WindowOptions { get; set; }

        public DashboardFilledMapConfigurationArgs()
        {
        }
        public static new DashboardFilledMapConfigurationArgs Empty => new DashboardFilledMapConfigurationArgs();
    }
}
