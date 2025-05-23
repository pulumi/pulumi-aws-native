// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardGeospatialStaticFileSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the static file.
        /// </summary>
        [Input("staticFileId", required: true)]
        public Input<string> StaticFileId { get; set; } = null!;

        public DashboardGeospatialStaticFileSourceArgs()
        {
        }
        public static new DashboardGeospatialStaticFileSourceArgs Empty => new DashboardGeospatialStaticFileSourceArgs();
    }
}
