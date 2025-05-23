// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Outputs
{

    /// <summary>
    /// Specifies encryption-related information for an Amazon S3 bucket that is a destination for replicated objects.
    ///   If you're specifying a customer managed KMS key, we recommend using a fully qualified KMS key ARN. If you use a KMS key alias instead, then KMS resolves the key within the requester’s account. This behavior can result in data that's encrypted with a KMS key that belongs to the requester, and not the bucket owner.
    /// </summary>
    [OutputType]
    public sealed class BucketEncryptionConfiguration
    {
        /// <summary>
        /// Specifies the ID (Key ARN or Alias ARN) of the customer managed AWS KMS key stored in AWS Key Management Service (KMS) for the destination bucket. Amazon S3 uses this key to encrypt replica objects. Amazon S3 only supports symmetric encryption KMS keys. For more information, see [Asymmetric keys in KMS](https://docs.aws.amazon.com//kms/latest/developerguide/symmetric-asymmetric.html) in the *Key Management Service Developer Guide*.
        /// </summary>
        public readonly string ReplicaKmsKeyId;

        [OutputConstructor]
        private BucketEncryptionConfiguration(string replicaKmsKeyId)
        {
            ReplicaKmsKeyId = replicaKmsKeyId;
        }
    }
}
