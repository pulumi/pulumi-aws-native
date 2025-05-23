// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Efs.Outputs
{

    /// <summary>
    /// Specifies the directory on the Amazon EFS file system that the access point provides access to. The access point exposes the specified file system path as the root directory of your file system to applications using the access point. NFS clients using the access point can only access data in the access point's ``RootDirectory`` and its subdirectories.
    /// </summary>
    [OutputType]
    public sealed class AccessPointRootDirectory
    {
        /// <summary>
        /// (Optional) Specifies the POSIX IDs and permissions to apply to the access point's ``RootDirectory``. If the ``RootDirectory`` &gt; ``Path`` specified does not exist, EFS creates the root directory using the ``CreationInfo`` settings when a client connects to an access point. When specifying the ``CreationInfo``, you must provide values for all properties. 
        ///   If you do not provide ``CreationInfo`` and the specified ``RootDirectory`` &gt; ``Path`` does not exist, attempts to mount the file system using the access point will fail.
        /// </summary>
        public readonly Outputs.AccessPointCreationInfo? CreationInfo;
        /// <summary>
        /// Specifies the path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system. A path can have up to four subdirectories. If the specified path does not exist, you are required to provide the ``CreationInfo``.
        /// </summary>
        public readonly string? Path;

        [OutputConstructor]
        private AccessPointRootDirectory(
            Outputs.AccessPointCreationInfo? creationInfo,

            string? path)
        {
            CreationInfo = creationInfo;
            Path = path;
        }
    }
}
