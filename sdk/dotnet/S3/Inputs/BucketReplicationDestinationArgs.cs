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
    /// A container for information about the replication destination and its configurations including enabling the S3 Replication Time Control (S3 RTC).
    /// </summary>
    public sealed class BucketReplicationDestinationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify this only in a cross-account scenario (where source and destination bucket owners are not the same), and you want to change replica ownership to the AWS-account that owns the destination bucket. If this is not specified in the replication configuration, the replicas are owned by same AWS-account that owns the source object.
        /// </summary>
        [Input("accessControlTranslation")]
        public Input<Inputs.BucketAccessControlTranslationArgs>? AccessControlTranslation { get; set; }

        /// <summary>
        /// Destination bucket owner account ID. In a cross-account scenario, if you direct Amazon S3 to change replica ownership to the AWS-account that owns the destination bucket by specifying the ``AccessControlTranslation`` property, this is the account ID of the destination bucket owner. For more information, see [Cross-Region Replication Additional Configuration: Change Replica Owner](https://docs.aws.amazon.com/AmazonS3/latest/dev/crr-change-owner.html) in the *Amazon S3 User Guide*.
        ///  If you specify the ``AccessControlTranslation`` property, the ``Account`` property is required.
        /// </summary>
        [Input("account")]
        public Input<string>? Account { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the bucket where you want Amazon S3 to store the results.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Specifies encryption-related information.
        /// </summary>
        [Input("encryptionConfiguration")]
        public Input<Inputs.BucketEncryptionConfigurationArgs>? EncryptionConfiguration { get; set; }

        /// <summary>
        /// A container specifying replication metrics-related settings enabling replication metrics and events.
        /// </summary>
        [Input("metrics")]
        public Input<Inputs.BucketMetricsArgs>? Metrics { get; set; }

        /// <summary>
        /// A container specifying S3 Replication Time Control (S3 RTC), including whether S3 RTC is enabled and the time when all objects and operations on objects must be replicated. Must be specified together with a ``Metrics`` block.
        /// </summary>
        [Input("replicationTime")]
        public Input<Inputs.BucketReplicationTimeArgs>? ReplicationTime { get; set; }

        /// <summary>
        /// The storage class to use when replicating objects, such as S3 Standard or reduced redundancy. By default, Amazon S3 uses the storage class of the source object to create the object replica. 
        ///  For valid values, see the ``StorageClass`` element of the [PUT Bucket replication](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTreplication.html) action in the *Amazon S3 API Reference*.
        /// </summary>
        [Input("storageClass")]
        public Input<Pulumi.AwsNative.S3.BucketReplicationDestinationStorageClass>? StorageClass { get; set; }

        public BucketReplicationDestinationArgs()
        {
        }
        public static new BucketReplicationDestinationArgs Empty => new BucketReplicationDestinationArgs();
    }
}
