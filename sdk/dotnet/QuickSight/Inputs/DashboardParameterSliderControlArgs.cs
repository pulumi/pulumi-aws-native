// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardParameterSliderControlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The display options of a control.
        /// </summary>
        [Input("displayOptions")]
        public Input<Inputs.DashboardSliderControlDisplayOptionsArgs>? DisplayOptions { get; set; }

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
        /// The ID of the `ParameterSliderControl` .
        /// </summary>
        [Input("parameterControlId", required: true)]
        public Input<string> ParameterControlId { get; set; } = null!;

        /// <summary>
        /// The source parameter name of the `ParameterSliderControl` .
        /// </summary>
        [Input("sourceParameterName", required: true)]
        public Input<string> SourceParameterName { get; set; } = null!;

        /// <summary>
        /// The number of increments that the slider bar is divided into.
        /// </summary>
        [Input("stepSize", required: true)]
        public Input<double> StepSize { get; set; } = null!;

        /// <summary>
        /// The title of the `ParameterSliderControl` .
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public DashboardParameterSliderControlArgs()
        {
        }
        public static new DashboardParameterSliderControlArgs Empty => new DashboardParameterSliderControlArgs();
    }
}
