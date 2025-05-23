// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Outputs
{

    /// <summary>
    /// The configuration information to connect to Amazon S3 as your data source.
    /// </summary>
    [OutputType]
    public sealed class DataSourceS3DataSourceConfiguration
    {
        /// <summary>
        /// The ARN of the bucket that contains the data source.
        /// </summary>
        public readonly string BucketArn;
        /// <summary>
        /// The account ID for the owner of the S3 bucket.
        /// </summary>
        public readonly string? BucketOwnerAccountId;
        /// <summary>
        /// A list of S3 prefixes that define the object containing the data sources.
        /// </summary>
        public readonly ImmutableArray<string> InclusionPrefixes;

        [OutputConstructor]
        private DataSourceS3DataSourceConfiguration(
            string bucketArn,

            string? bucketOwnerAccountId,

            ImmutableArray<string> inclusionPrefixes)
        {
            BucketArn = bucketArn;
            BucketOwnerAccountId = bucketOwnerAccountId;
            InclusionPrefixes = inclusionPrefixes;
        }
    }
}
