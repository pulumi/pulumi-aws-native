// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Inputs
{

    /// <summary>
    /// The managed storage configuration for the cluster.
    /// </summary>
    public sealed class ClusterManagedStorageConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify the KMSlong key ID for Fargate ephemeral storage.
        ///  When you specify a ``fargateEphemeralStorageKmsKeyId``, AWS Fargate uses the key to encrypt data at rest in ephemeral storage. For more information about Fargate ephemeral storage encryption, see [Customer managed keys for Fargate ephemeral storage for Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/fargate-storage-encryption.html) in the *Amazon Elastic Container Service Developer Guide*.
        ///  The key must be a single Region key.
        /// </summary>
        [Input("fargateEphemeralStorageKmsKeyId")]
        public Input<string>? FargateEphemeralStorageKmsKeyId { get; set; }

        /// <summary>
        /// Specify a KMSlong key ID to encrypt Amazon ECS managed storage.
        ///   When you specify a ``kmsKeyId``, Amazon ECS uses the key to encrypt data volumes managed by Amazon ECS that are attached to tasks in the cluster. The following data volumes are managed by Amazon ECS: Amazon EBS. For more information about encryption of Amazon EBS volumes attached to Amazon ECS tasks, see [Encrypt data stored in Amazon EBS volumes for Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ebs-kms-encryption.html) in the *Amazon Elastic Container Service Developer Guide*.
        ///  The key must be a single Region key.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        public ClusterManagedStorageConfigurationArgs()
        {
        }
        public static new ClusterManagedStorageConfigurationArgs Empty => new ClusterManagedStorageConfigurationArgs();
    }
}
