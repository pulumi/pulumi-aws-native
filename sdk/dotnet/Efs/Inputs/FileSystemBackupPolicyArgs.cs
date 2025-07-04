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
    /// The backup policy turns automatic backups for the file system on or off.
    /// </summary>
    public sealed class FileSystemBackupPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set the backup policy status for the file system.
        ///   +  *ENABLED* - Turns automatic backups on for the file system. 
        ///   +  *DISABLED* - Turns automatic backups off for the file system.
        /// </summary>
        [Input("status", required: true)]
        public Input<Pulumi.AwsNative.Efs.FileSystemBackupPolicyStatus> Status { get; set; } = null!;

        public FileSystemBackupPolicyArgs()
        {
        }
        public static new FileSystemBackupPolicyArgs Empty => new FileSystemBackupPolicyArgs();
    }
}
