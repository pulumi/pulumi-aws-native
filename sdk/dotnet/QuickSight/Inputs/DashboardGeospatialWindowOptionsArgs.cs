// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardGeospatialWindowOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("bounds")]
        public Input<Inputs.DashboardGeospatialCoordinateBoundsArgs>? Bounds { get; set; }

        [Input("mapZoomMode")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardMapZoomMode>? MapZoomMode { get; set; }

        public DashboardGeospatialWindowOptionsArgs()
        {
        }
        public static new DashboardGeospatialWindowOptionsArgs Empty => new DashboardGeospatialWindowOptionsArgs();
    }
}