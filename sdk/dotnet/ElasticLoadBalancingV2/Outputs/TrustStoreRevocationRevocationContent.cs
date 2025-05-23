// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2.Outputs
{

    [OutputType]
    public sealed class TrustStoreRevocationRevocationContent
    {
        /// <summary>
        /// The type of revocation file.
        /// </summary>
        public readonly string? RevocationType;
        /// <summary>
        /// The Amazon S3 bucket for the revocation file.
        /// </summary>
        public readonly string? S3Bucket;
        /// <summary>
        /// The Amazon S3 path for the revocation file.
        /// </summary>
        public readonly string? S3Key;
        /// <summary>
        /// The Amazon S3 object version of the revocation file.
        /// </summary>
        public readonly string? S3ObjectVersion;

        [OutputConstructor]
        private TrustStoreRevocationRevocationContent(
            string? revocationType,

            string? s3Bucket,

            string? s3Key,

            string? s3ObjectVersion)
        {
            RevocationType = revocationType;
            S3Bucket = s3Bucket;
            S3Key = s3Key;
            S3ObjectVersion = s3ObjectVersion;
        }
    }
}
