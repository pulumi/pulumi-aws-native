// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateAxisDisplayOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("axisLineVisibility")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateVisibility>? AxisLineVisibility { get; set; }

        /// <summary>
        /// String based length that is composed of value and unit in px
        /// </summary>
        [Input("axisOffset")]
        public Input<string>? AxisOffset { get; set; }

        [Input("dataOptions")]
        public Input<Inputs.TemplateAxisDataOptionsArgs>? DataOptions { get; set; }

        [Input("gridLineVisibility")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateVisibility>? GridLineVisibility { get; set; }

        [Input("scrollbarOptions")]
        public Input<Inputs.TemplateScrollBarOptionsArgs>? ScrollbarOptions { get; set; }

        [Input("tickLabelOptions")]
        public Input<Inputs.TemplateAxisTickLabelOptionsArgs>? TickLabelOptions { get; set; }

        public TemplateAxisDisplayOptionsArgs()
        {
        }
        public static new TemplateAxisDisplayOptionsArgs Empty => new TemplateAxisDisplayOptionsArgs();
    }
}