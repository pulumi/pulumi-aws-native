// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Outputs
{

    /// <summary>
    /// Sends Verified Access logs to Amazon S3.
    /// </summary>
    [OutputType]
    public sealed class VerifiedAccessInstanceVerifiedAccessLogsS3Properties
    {
        /// <summary>
        /// The bucket name.
        /// </summary>
        public readonly string? BucketName;
        /// <summary>
        /// The ID of the AWS account that owns the Amazon S3 bucket.
        /// </summary>
        public readonly string? BucketOwner;
        /// <summary>
        /// Indicates whether logging is enabled.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The bucket prefix.
        /// </summary>
        public readonly string? Prefix;

        [OutputConstructor]
        private VerifiedAccessInstanceVerifiedAccessLogsS3Properties(
            string? bucketName,

            string? bucketOwner,

            bool? enabled,

            string? prefix)
        {
            BucketName = bucketName;
            BucketOwner = bucketOwner;
            Enabled = enabled;
            Prefix = prefix;
        }
    }
}
