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
    public sealed class DashboardStaticFile
    {
        /// <summary>
        /// The image static file.
        /// </summary>
        public readonly Outputs.DashboardImageStaticFile? ImageStaticFile;
        /// <summary>
        /// The spacial static file.
        /// </summary>
        public readonly Outputs.DashboardSpatialStaticFile? SpatialStaticFile;

        [OutputConstructor]
        private DashboardStaticFile(
            Outputs.DashboardImageStaticFile? imageStaticFile,

            Outputs.DashboardSpatialStaticFile? spatialStaticFile)
        {
            ImageStaticFile = imageStaticFile;
            SpatialStaticFile = spatialStaticFile;
        }
    }
}
