// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardGridLayoutCanvasSizeOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("screenCanvasSizeOptions")]
        public Input<Inputs.DashboardGridLayoutScreenCanvasSizeOptionsArgs>? ScreenCanvasSizeOptions { get; set; }

        public DashboardGridLayoutCanvasSizeOptionsArgs()
        {
        }
        public static new DashboardGridLayoutCanvasSizeOptionsArgs Empty => new DashboardGridLayoutCanvasSizeOptionsArgs();
    }
}