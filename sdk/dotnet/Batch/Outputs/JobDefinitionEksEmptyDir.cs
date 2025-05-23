// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Outputs
{

    [OutputType]
    public sealed class JobDefinitionEksEmptyDir
    {
        /// <summary>
        /// The medium to store the volume. The default value is an empty string, which uses the storage of the node.
        /// 
        /// - **""** - *(Default)* Use the disk storage of the node.
        /// - **"Memory"** - Use the `tmpfs` volume that's backed by the RAM of the node. Contents of the volume are lost when the node reboots, and any storage on the volume counts against the container's memory limit.
        /// </summary>
        public readonly string? Medium;
        /// <summary>
        /// The maximum size of the volume. By default, there's no maximum size defined.
        /// </summary>
        public readonly string? SizeLimit;

        [OutputConstructor]
        private JobDefinitionEksEmptyDir(
            string? medium,

            string? sizeLimit)
        {
            Medium = medium;
            SizeLimit = sizeLimit;
        }
    }
}
