// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardGeospatialHeatmapColorScaleArgs : global::Pulumi.ResourceArgs
    {
        [Input("colors")]
        private InputList<Inputs.DashboardGeospatialHeatmapDataColorArgs>? _colors;
        public InputList<Inputs.DashboardGeospatialHeatmapDataColorArgs> Colors
        {
            get => _colors ?? (_colors = new InputList<Inputs.DashboardGeospatialHeatmapDataColorArgs>());
            set => _colors = value;
        }

        public DashboardGeospatialHeatmapColorScaleArgs()
        {
        }
        public static new DashboardGeospatialHeatmapColorScaleArgs Empty => new DashboardGeospatialHeatmapColorScaleArgs();
    }
}