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
    public sealed class AnalysisGeospatialHeatmapConfiguration
    {
        /// <summary>
        /// The color scale specification for the heatmap point style.
        /// </summary>
        public readonly Outputs.AnalysisGeospatialHeatmapColorScale? HeatmapColor;

        [OutputConstructor]
        private AnalysisGeospatialHeatmapConfiguration(Outputs.AnalysisGeospatialHeatmapColorScale? heatmapColor)
        {
            HeatmapColor = heatmapColor;
        }
    }
}
