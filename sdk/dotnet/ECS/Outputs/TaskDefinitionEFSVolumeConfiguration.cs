// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ECS.Outputs
{

    [OutputType]
    public sealed class TaskDefinitionEFSVolumeConfiguration
    {
        public readonly Outputs.TaskDefinitionAuthorizationConfig? AuthorizationConfig;
        public readonly string FilesystemId;
        public readonly string? RootDirectory;
        public readonly Pulumi.AwsNative.ECS.TaskDefinitionEFSVolumeConfigurationTransitEncryption? TransitEncryption;
        public readonly int? TransitEncryptionPort;

        [OutputConstructor]
        private TaskDefinitionEFSVolumeConfiguration(
            Outputs.TaskDefinitionAuthorizationConfig? authorizationConfig,

            string filesystemId,

            string? rootDirectory,

            Pulumi.AwsNative.ECS.TaskDefinitionEFSVolumeConfigurationTransitEncryption? transitEncryption,

            int? transitEncryptionPort)
        {
            AuthorizationConfig = authorizationConfig;
            FilesystemId = filesystemId;
            RootDirectory = rootDirectory;
            TransitEncryption = transitEncryption;
            TransitEncryptionPort = transitEncryptionPort;
        }
    }
}