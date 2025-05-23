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
    /// A container that describes additional filters for identifying the source objects that you want to replicate. You can choose to enable or disable the replication of these objects.
    /// </summary>
    [OutputType]
    public sealed class BucketSourceSelectionCriteria
    {
        /// <summary>
        /// A filter that you can specify for selection for modifications on replicas.
        /// </summary>
        public readonly Outputs.BucketReplicaModifications? ReplicaModifications;
        /// <summary>
        /// A container for filter information for the selection of Amazon S3 objects encrypted with AWS KMS.
        /// </summary>
        public readonly Outputs.BucketSseKmsEncryptedObjects? SseKmsEncryptedObjects;

        [OutputConstructor]
        private BucketSourceSelectionCriteria(
            Outputs.BucketReplicaModifications? replicaModifications,

            Outputs.BucketSseKmsEncryptedObjects? sseKmsEncryptedObjects)
        {
            ReplicaModifications = replicaModifications;
            SseKmsEncryptedObjects = sseKmsEncryptedObjects;
        }
    }
}
