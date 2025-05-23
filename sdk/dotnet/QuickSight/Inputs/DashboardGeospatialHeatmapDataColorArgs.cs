// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardGeospatialHeatmapDataColorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The hex color to be used in the heatmap point style.
        /// </summary>
        [Input("color", required: true)]
        public Input<string> Color { get; set; } = null!;

        public DashboardGeospatialHeatmapDataColorArgs()
        {
        }
        public static new DashboardGeospatialHeatmapDataColorArgs Empty => new DashboardGeospatialHeatmapDataColorArgs();
    }
}
