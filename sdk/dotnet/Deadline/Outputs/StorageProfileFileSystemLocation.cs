// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Deadline.Outputs
{

    [OutputType]
    public sealed class StorageProfileFileSystemLocation
    {
        public readonly string Name;
        public readonly string Path;
        public readonly Pulumi.AwsNative.Deadline.StorageProfileFileSystemLocationType Type;

        [OutputConstructor]
        private StorageProfileFileSystemLocation(
            string name,

            string path,

            Pulumi.AwsNative.Deadline.StorageProfileFileSystemLocationType type)
        {
            Name = name;
            Path = path;
            Type = type;
        }
    }
}