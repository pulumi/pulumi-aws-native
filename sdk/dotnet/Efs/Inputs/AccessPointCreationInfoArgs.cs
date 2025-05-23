// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Efs.Inputs
{

    /// <summary>
    /// Required if the ``RootDirectory`` &gt; ``Path`` specified does not exist. Specifies the POSIX IDs and permissions to apply to the access point's ``RootDirectory`` &gt; ``Path``. If the access point root directory does not exist, EFS creates it with these settings when a client connects to the access point. When specifying ``CreationInfo``, you must include values for all properties. 
    ///  Amazon EFS creates a root directory only if you have provided the CreationInfo: OwnUid, OwnGID, and permissions for the directory. If you do not provide this information, Amazon EFS does not create the root directory. If the root directory does not exist, attempts to mount using the access point will fail.
    ///   If you do not provide ``CreationInfo`` and the specified ``RootDirectory`` does not exist, attempts to mount the file system using the access point will fail.
    /// </summary>
    public sealed class AccessPointCreationInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the POSIX group ID to apply to the ``RootDirectory``. Accepts values from 0 to 2^32 (4294967295).
        /// </summary>
        [Input("ownerGid", required: true)]
        public Input<string> OwnerGid { get; set; } = null!;

        /// <summary>
        /// Specifies the POSIX user ID to apply to the ``RootDirectory``. Accepts values from 0 to 2^32 (4294967295).
        /// </summary>
        [Input("ownerUid", required: true)]
        public Input<string> OwnerUid { get; set; } = null!;

        /// <summary>
        /// Specifies the POSIX permissions to apply to the ``RootDirectory``, in the format of an octal number representing the file's mode bits.
        /// </summary>
        [Input("permissions", required: true)]
        public Input<string> Permissions { get; set; } = null!;

        public AccessPointCreationInfoArgs()
        {
        }
        public static new AccessPointCreationInfoArgs Empty => new AccessPointCreationInfoArgs();
    }
}
