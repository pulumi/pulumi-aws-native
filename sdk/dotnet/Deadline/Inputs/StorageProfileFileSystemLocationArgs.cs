// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Deadline.Inputs
{

    public sealed class StorageProfileFileSystemLocationArgs : global::Pulumi.ResourceArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.Deadline.StorageProfileFileSystemLocationType> Type { get; set; } = null!;

        public StorageProfileFileSystemLocationArgs()
        {
        }
        public static new StorageProfileFileSystemLocationArgs Empty => new StorageProfileFileSystemLocationArgs();
    }
}