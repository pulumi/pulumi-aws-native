// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Outputs
{

    [OutputType]
    public sealed class TaskDefinitionFSxWindowsFileServerVolumeConfiguration
    {
        public readonly Outputs.TaskDefinitionFSxAuthorizationConfig? AuthorizationConfig;
        public readonly string FileSystemId;
        public readonly string RootDirectory;

        [OutputConstructor]
        private TaskDefinitionFSxWindowsFileServerVolumeConfiguration(
            Outputs.TaskDefinitionFSxAuthorizationConfig? authorizationConfig,

            string fileSystemId,

            string rootDirectory)
        {
            AuthorizationConfig = authorizationConfig;
            FileSystemId = fileSystemId;
            RootDirectory = rootDirectory;
        }
    }
}