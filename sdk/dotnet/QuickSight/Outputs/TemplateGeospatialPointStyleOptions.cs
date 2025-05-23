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
    public sealed class TemplateGeospatialPointStyleOptions
    {
        /// <summary>
        /// The cluster marker configuration of the geospatial point style.
        /// </summary>
        public readonly Outputs.TemplateClusterMarkerConfiguration? ClusterMarkerConfiguration;
        /// <summary>
        /// The heatmap configuration of the geospatial point style.
        /// </summary>
        public readonly Outputs.TemplateGeospatialHeatmapConfiguration? HeatmapConfiguration;
        /// <summary>
        /// The selected point styles (point, cluster) of the geospatial map.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.TemplateGeospatialSelectedPointStyle? SelectedPointStyle;

        [OutputConstructor]
        private TemplateGeospatialPointStyleOptions(
            Outputs.TemplateClusterMarkerConfiguration? clusterMarkerConfiguration,

            Outputs.TemplateGeospatialHeatmapConfiguration? heatmapConfiguration,

            Pulumi.AwsNative.QuickSight.TemplateGeospatialSelectedPointStyle? selectedPointStyle)
        {
            ClusterMarkerConfiguration = clusterMarkerConfiguration;
            HeatmapConfiguration = heatmapConfiguration;
            SelectedPointStyle = selectedPointStyle;
        }
    }
}
