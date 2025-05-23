// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisGeospatialCircleSymbolStyleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The radius of the circle.
        /// </summary>
        [Input("circleRadius")]
        public Input<Inputs.AnalysisGeospatialCircleRadiusArgs>? CircleRadius { get; set; }

        /// <summary>
        /// The color and opacity values for the fill color.
        /// </summary>
        [Input("fillColor")]
        public Input<Inputs.AnalysisGeospatialColorArgs>? FillColor { get; set; }

        /// <summary>
        /// The color and opacity values for the stroke color.
        /// </summary>
        [Input("strokeColor")]
        public Input<Inputs.AnalysisGeospatialColorArgs>? StrokeColor { get; set; }

        /// <summary>
        /// The width of the stroke (border).
        /// </summary>
        [Input("strokeWidth")]
        public Input<Inputs.AnalysisGeospatialLineWidthArgs>? StrokeWidth { get; set; }

        public AnalysisGeospatialCircleSymbolStyleArgs()
        {
        }
        public static new AnalysisGeospatialCircleSymbolStyleArgs Empty => new AnalysisGeospatialCircleSymbolStyleArgs();
    }
}
