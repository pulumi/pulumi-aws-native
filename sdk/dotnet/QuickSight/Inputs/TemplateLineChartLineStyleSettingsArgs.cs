// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateLineChartLineStyleSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Interpolation style for line series.
        /// 
        /// - `LINEAR` : Show as default, linear style.
        /// - `SMOOTH` : Show as a smooth curve.
        /// - `STEPPED` : Show steps in line.
        /// </summary>
        [Input("lineInterpolation")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateLineInterpolation>? LineInterpolation { get; set; }

        /// <summary>
        /// Line style for line series.
        /// 
        /// - `SOLID` : Show as a solid line.
        /// - `DOTTED` : Show as a dotted line.
        /// - `DASHED` : Show as a dashed line.
        /// </summary>
        [Input("lineStyle")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateLineChartLineStyle>? LineStyle { get; set; }

        /// <summary>
        /// Configuration option that determines whether to show the line for the series.
        /// </summary>
        [Input("lineVisibility")]
        public Input<object>? LineVisibility { get; set; }

        /// <summary>
        /// String based length that is composed of value and unit in px
        /// </summary>
        [Input("lineWidth")]
        public Input<string>? LineWidth { get; set; }

        public TemplateLineChartLineStyleSettingsArgs()
        {
        }
        public static new TemplateLineChartLineStyleSettingsArgs Empty => new TemplateLineChartLineStyleSettingsArgs();
    }
}
