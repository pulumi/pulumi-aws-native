// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardGeospatialLineLayerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The visualization style for a line layer.
        /// </summary>
        [Input("style", required: true)]
        public Input<Inputs.DashboardGeospatialLineStyleArgs> Style { get; set; } = null!;

        public DashboardGeospatialLineLayerArgs()
        {
        }
        public static new DashboardGeospatialLineLayerArgs Empty => new DashboardGeospatialLineLayerArgs();
    }
}
