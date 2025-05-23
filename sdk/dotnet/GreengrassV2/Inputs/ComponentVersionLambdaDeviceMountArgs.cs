// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GreengrassV2.Inputs
{

    public sealed class ComponentVersionLambdaDeviceMountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether or not to add the component's system user as an owner of the device.
        /// 
        /// Default: `false`
        /// </summary>
        [Input("addGroupOwner")]
        public Input<bool>? AddGroupOwner { get; set; }

        /// <summary>
        /// The mount path for the device in the file system.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The permission to access the device: read/only ( `ro` ) or read/write ( `rw` ).
        /// 
        /// Default: `ro`
        /// </summary>
        [Input("permission")]
        public Input<Pulumi.AwsNative.GreengrassV2.ComponentVersionLambdaFilesystemPermission>? Permission { get; set; }

        public ComponentVersionLambdaDeviceMountArgs()
        {
        }
        public static new ComponentVersionLambdaDeviceMountArgs Empty => new ComponentVersionLambdaDeviceMountArgs();
    }
}
