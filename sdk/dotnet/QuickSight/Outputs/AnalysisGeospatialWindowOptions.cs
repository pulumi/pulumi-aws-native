// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class AnalysisGeospatialWindowOptions
    {
        public readonly Outputs.AnalysisGeospatialCoordinateBounds? Bounds;
        public readonly Pulumi.AwsNative.QuickSight.AnalysisMapZoomMode? MapZoomMode;

        [OutputConstructor]
        private AnalysisGeospatialWindowOptions(
            Outputs.AnalysisGeospatialCoordinateBounds? bounds,

            Pulumi.AwsNative.QuickSight.AnalysisMapZoomMode? mapZoomMode)
        {
            Bounds = bounds;
            MapZoomMode = mapZoomMode;
        }
    }
}