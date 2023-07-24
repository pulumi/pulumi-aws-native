// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateLegendOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// String based length that is composed of value and unit in px
        /// </summary>
        [Input("height")]
        public Input<string>? Height { get; set; }

        [Input("position")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateLegendPosition>? Position { get; set; }

        [Input("title")]
        public Input<Inputs.TemplateLabelOptionsArgs>? Title { get; set; }

        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateVisibility>? Visibility { get; set; }

        /// <summary>
        /// String based length that is composed of value and unit in px
        /// </summary>
        [Input("width")]
        public Input<string>? Width { get; set; }

        public TemplateLegendOptionsArgs()
        {
        }
        public static new TemplateLegendOptionsArgs Empty => new TemplateLegendOptionsArgs();
    }
}