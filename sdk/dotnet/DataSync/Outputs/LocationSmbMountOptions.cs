// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataSync.Outputs
{

    /// <summary>
    /// The mount options used by DataSync to access the SMB server.
    /// </summary>
    [OutputType]
    public sealed class LocationSmbMountOptions
    {
        /// <summary>
        /// The specific SMB version that you want DataSync to use to mount your SMB share.
        /// </summary>
        public readonly Pulumi.AwsNative.DataSync.LocationSmbMountOptionsVersion? Version;

        [OutputConstructor]
        private LocationSmbMountOptions(Pulumi.AwsNative.DataSync.LocationSmbMountOptionsVersion? version)
        {
            Version = version;
        }
    }
}
