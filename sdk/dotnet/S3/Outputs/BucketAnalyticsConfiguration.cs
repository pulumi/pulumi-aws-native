// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Outputs
{

    /// <summary>
    /// Specifies the configuration and any analyses for the analytics filter of an Amazon S3 bucket.
    /// </summary>
    [OutputType]
    public sealed class BucketAnalyticsConfiguration
    {
        /// <summary>
        /// The ID that identifies the analytics configuration.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The prefix that an object must have to be included in the analytics results.
        /// </summary>
        public readonly string? Prefix;
        public readonly Outputs.BucketStorageClassAnalysis StorageClassAnalysis;
        public readonly ImmutableArray<Outputs.BucketTagFilter> TagFilters;

        [OutputConstructor]
        private BucketAnalyticsConfiguration(
            string id,

            string? prefix,

            Outputs.BucketStorageClassAnalysis storageClassAnalysis,

            ImmutableArray<Outputs.BucketTagFilter> tagFilters)
        {
            Id = id;
            Prefix = prefix;
            StorageClassAnalysis = storageClassAnalysis;
            TagFilters = tagFilters;
        }
    }
}