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
    /// The identifier for the S3 resource.
    /// </summary>
    [OutputType]
    public sealed class AgentS3Identifier
    {
        /// <summary>
        /// A bucket in S3.
        /// </summary>
        public readonly string? S3BucketName;
        /// <summary>
        /// A object key in S3.
        /// </summary>
        public readonly string? S3ObjectKey;

        [OutputConstructor]
        private AgentS3Identifier(
            string? s3BucketName,

            string? s3ObjectKey)
        {
            S3BucketName = s3BucketName;
            S3ObjectKey = s3ObjectKey;
        }
    }
}
