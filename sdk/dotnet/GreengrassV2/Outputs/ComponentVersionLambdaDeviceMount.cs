// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GreengrassV2.Outputs
{

    [OutputType]
    public sealed class ComponentVersionLambdaDeviceMount
    {
        /// <summary>
        /// Whether or not to add the component's system user as an owner of the device.
        /// 
        /// Default: `false`
        /// </summary>
        public readonly bool? AddGroupOwner;
        /// <summary>
        /// The mount path for the device in the file system.
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// The permission to access the device: read/only ( `ro` ) or read/write ( `rw` ).
        /// 
        /// Default: `ro`
        /// </summary>
        public readonly Pulumi.AwsNative.GreengrassV2.ComponentVersionLambdaFilesystemPermission? Permission;

        [OutputConstructor]
        private ComponentVersionLambdaDeviceMount(
            bool? addGroupOwner,

            string? path,

            Pulumi.AwsNative.GreengrassV2.ComponentVersionLambdaFilesystemPermission? permission)
        {
            AddGroupOwner = addGroupOwner;
            Path = path;
            Permission = permission;
        }
    }
}
