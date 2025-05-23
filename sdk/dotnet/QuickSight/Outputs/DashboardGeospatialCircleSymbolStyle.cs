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
    public sealed class DashboardGeospatialCircleSymbolStyle
    {
        /// <summary>
        /// The radius of the circle.
        /// </summary>
        public readonly Outputs.DashboardGeospatialCircleRadius? CircleRadius;
        /// <summary>
        /// The color and opacity values for the fill color.
        /// </summary>
        public readonly Outputs.DashboardGeospatialColor? FillColor;
        /// <summary>
        /// The color and opacity values for the stroke color.
        /// </summary>
        public readonly Outputs.DashboardGeospatialColor? StrokeColor;
        /// <summary>
        /// The width of the stroke (border).
        /// </summary>
        public readonly Outputs.DashboardGeospatialLineWidth? StrokeWidth;

        [OutputConstructor]
        private DashboardGeospatialCircleSymbolStyle(
            Outputs.DashboardGeospatialCircleRadius? circleRadius,

            Outputs.DashboardGeospatialColor? fillColor,

            Outputs.DashboardGeospatialColor? strokeColor,

            Outputs.DashboardGeospatialLineWidth? strokeWidth)
        {
            CircleRadius = circleRadius;
            FillColor = fillColor;
            StrokeColor = strokeColor;
            StrokeWidth = strokeWidth;
        }
    }
}
