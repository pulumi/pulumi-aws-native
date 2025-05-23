// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateKpiSparklineOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The color of the sparkline.
        /// </summary>
        [Input("color")]
        public Input<string>? Color { get; set; }

        /// <summary>
        /// The tooltip visibility of the sparkline.
        /// </summary>
        [Input("tooltipVisibility")]
        public Input<object>? TooltipVisibility { get; set; }

        /// <summary>
        /// The type of the sparkline.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.QuickSight.TemplateKpiSparklineType> Type { get; set; } = null!;

        /// <summary>
        /// The visibility of the sparkline.
        /// </summary>
        [Input("visibility")]
        public Input<object>? Visibility { get; set; }

        public TemplateKpiSparklineOptionsArgs()
        {
        }
        public static new TemplateKpiSparklineOptionsArgs Empty => new TemplateKpiSparklineOptionsArgs();
    }
}
