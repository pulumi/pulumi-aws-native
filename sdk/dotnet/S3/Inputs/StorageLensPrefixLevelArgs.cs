// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    /// <summary>
    /// Prefix-level metrics configurations.
    /// </summary>
    public sealed class StorageLensPrefixLevelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A property for the prefix-level storage metrics for Amazon S3 Storage Lens.
        /// </summary>
        [Input("storageMetrics", required: true)]
        public Input<Inputs.StorageLensPrefixLevelStorageMetricsArgs> StorageMetrics { get; set; } = null!;

        public StorageLensPrefixLevelArgs()
        {
        }
        public static new StorageLensPrefixLevelArgs Empty => new StorageLensPrefixLevelArgs();
    }
}
