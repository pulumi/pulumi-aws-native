// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardGeospatialPointStyleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The circle symbol style for a point layer.
        /// </summary>
        [Input("circleSymbolStyle")]
        public Input<Inputs.DashboardGeospatialCircleSymbolStyleArgs>? CircleSymbolStyle { get; set; }

        public DashboardGeospatialPointStyleArgs()
        {
        }
        public static new DashboardGeospatialPointStyleArgs Empty => new DashboardGeospatialPointStyleArgs();
    }
}
