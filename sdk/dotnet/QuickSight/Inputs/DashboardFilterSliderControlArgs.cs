// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardFilterSliderControlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The display options of a control.
        /// </summary>
        [Input("displayOptions")]
        public Input<Inputs.DashboardSliderControlDisplayOptionsArgs>? DisplayOptions { get; set; }

        /// <summary>
        /// The ID of the `FilterSliderControl` .
        /// </summary>
        [Input("filterControlId", required: true)]
        public Input<string> FilterControlId { get; set; } = null!;

        /// <summary>
        /// The larger value that is displayed at the right of the slider.
        /// </summary>
        [Input("maximumValue", required: true)]
        public Input<double> MaximumValue { get; set; } = null!;

        /// <summary>
        /// The smaller value that is displayed at the left of the slider.
        /// </summary>
        [Input("minimumValue", required: true)]
        public Input<double> MinimumValue { get; set; } = null!;

        /// <summary>
        /// The source filter ID of the `FilterSliderControl` .
        /// </summary>
        [Input("sourceFilterId", required: true)]
        public Input<string> SourceFilterId { get; set; } = null!;

        /// <summary>
        /// The number of increments that the slider bar is divided into.
        /// </summary>
        [Input("stepSize", required: true)]
        public Input<double> StepSize { get; set; } = null!;

        /// <summary>
        /// The title of the `FilterSliderControl` .
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        /// <summary>
        /// The type of the `FilterSliderControl` . Choose one of the following options:
        /// 
        /// - `SINGLE_POINT` : Filter against(equals) a single data point.
        /// - `RANGE` : Filter data that is in a specified range.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardSheetControlSliderType>? Type { get; set; }

        public DashboardFilterSliderControlArgs()
        {
        }
        public static new DashboardFilterSliderControlArgs Empty => new DashboardFilterSliderControlArgs();
    }
}
