// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    [OutputType]
    public sealed class UserProfileCustomFileSystemConfig
    {
        /// <summary>
        /// The settings for a custom Amazon EFS file system.
        /// </summary>
        public readonly Outputs.UserProfileEfsFileSystemConfig? EfsFileSystemConfig;

        [OutputConstructor]
        private UserProfileCustomFileSystemConfig(Outputs.UserProfileEfsFileSystemConfig? efsFileSystemConfig)
        {
            EfsFileSystemConfig = efsFileSystemConfig;
        }
    }
}
