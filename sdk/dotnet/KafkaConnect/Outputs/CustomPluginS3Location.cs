// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KafkaConnect.Outputs
{

    /// <summary>
    /// The S3 bucket Amazon Resource Name (ARN), file key, and object version of the plugin file stored in Amazon S3.
    /// </summary>
    [OutputType]
    public sealed class CustomPluginS3Location
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of an S3 bucket.
        /// </summary>
        public readonly string BucketArn;
        /// <summary>
        /// The file key for an object in an S3 bucket.
        /// </summary>
        public readonly string FileKey;
        /// <summary>
        /// The version of an object in an S3 bucket.
        /// </summary>
        public readonly string? ObjectVersion;

        [OutputConstructor]
        private CustomPluginS3Location(
            string bucketArn,

            string fileKey,

            string? objectVersion)
        {
            BucketArn = bucketArn;
            FileKey = fileKey;
            ObjectVersion = objectVersion;
        }
    }
}
