// *** WARNING: this file was generated by pulumi. ***
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
        /// Specify the KMSlong key ID for the Fargate ephemeral storage.
        ///  The key must be a single Region key.
        /// </summary>
        [Input("fargateEphemeralStorageKmsKeyId")]
        public Input<string>? FargateEphemeralStorageKmsKeyId { get; set; }

        /// <summary>
        /// Specify a KMSlong key ID to encrypt the managed storage.
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
