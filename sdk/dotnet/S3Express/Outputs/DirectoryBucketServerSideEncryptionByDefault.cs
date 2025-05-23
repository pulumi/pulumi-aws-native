// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3Express.Outputs
{

    /// <summary>
    /// Specifies the default server-side encryption to apply to new objects in the bucket. If a PUT Object request doesn't specify any server-side encryption, this default encryption will be applied.
    /// </summary>
    [OutputType]
    public sealed class DirectoryBucketServerSideEncryptionByDefault
    {
        /// <summary>
        /// AWS Key Management Service (KMS) customer managed key ID to use for the default encryption. This parameter is allowed only if SSEAlgorithm is set to aws:kms. You can specify this parameter with the key ID or the Amazon Resource Name (ARN) of the KMS key
        /// </summary>
        public readonly string? KmsMasterKeyId;
        /// <summary>
        /// Server-side encryption algorithm to use for the default encryption.
        /// 
        /// &gt; For directory buckets, there are only two supported values for server-side encryption: `AES256` and `aws:kms` .
        /// </summary>
        public readonly Pulumi.AwsNative.S3Express.DirectoryBucketServerSideEncryptionByDefaultSseAlgorithm SseAlgorithm;

        [OutputConstructor]
        private DirectoryBucketServerSideEncryptionByDefault(
            string? kmsMasterKeyId,

            Pulumi.AwsNative.S3Express.DirectoryBucketServerSideEncryptionByDefaultSseAlgorithm sseAlgorithm)
        {
            KmsMasterKeyId = kmsMasterKeyId;
            SseAlgorithm = sseAlgorithm;
        }
    }
}
