// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class AnalysisGeospatialPolygonSymbolStyle
    {
        /// <summary>
        /// The color and opacity values for the fill color.
        /// </summary>
        public readonly Outputs.AnalysisGeospatialColor? FillColor;
        /// <summary>
        /// The color and opacity values for the stroke color.
        /// </summary>
        public readonly Outputs.AnalysisGeospatialColor? StrokeColor;
        /// <summary>
        /// The width of the border stroke.
        /// </summary>
        public readonly Outputs.AnalysisGeospatialLineWidth? StrokeWidth;

        [OutputConstructor]
        private AnalysisGeospatialPolygonSymbolStyle(
            Outputs.AnalysisGeospatialColor? fillColor,

            Outputs.AnalysisGeospatialColor? strokeColor,

            Outputs.AnalysisGeospatialLineWidth? strokeWidth)
        {
            FillColor = fillColor;
            StrokeColor = strokeColor;
            StrokeWidth = strokeWidth;
        }
    }
}
