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
    /// Configuration settings for an NFS or SMB protocol, currently only support NFS
    /// </summary>
    [OutputType]
    public sealed class LocationFSxOpenZfsProtocol
    {
        /// <summary>
        /// Represents the Network File System (NFS) protocol that DataSync uses to access your FSx for OpenZFS file system.
        /// </summary>
        public readonly Outputs.LocationFSxOpenZfsNfs? Nfs;

        [OutputConstructor]
        private LocationFSxOpenZfsProtocol(Outputs.LocationFSxOpenZfsNfs? nfs)
        {
            Nfs = nfs;
        }
    }
}
