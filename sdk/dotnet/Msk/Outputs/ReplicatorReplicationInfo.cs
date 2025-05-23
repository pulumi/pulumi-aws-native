// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk.Outputs
{

    /// <summary>
    /// Specifies configuration for replication between a source and target Kafka cluster.
    /// </summary>
    [OutputType]
    public sealed class ReplicatorReplicationInfo
    {
        /// <summary>
        /// Configuration relating to consumer group replication.
        /// </summary>
        public readonly Outputs.ReplicatorConsumerGroupReplication ConsumerGroupReplication;
        /// <summary>
        /// Amazon Resource Name of the source Kafka cluster.
        /// </summary>
        public readonly string SourceKafkaClusterArn;
        /// <summary>
        /// The type of compression to use writing records to target Kafka cluster.
        /// </summary>
        public readonly Pulumi.AwsNative.Msk.ReplicatorReplicationInfoTargetCompressionType TargetCompressionType;
        /// <summary>
        /// Amazon Resource Name of the target Kafka cluster.
        /// </summary>
        public readonly string TargetKafkaClusterArn;
        /// <summary>
        /// Configuration relating to topic replication.
        /// </summary>
        public readonly Outputs.ReplicatorTopicReplication TopicReplication;

        [OutputConstructor]
        private ReplicatorReplicationInfo(
            Outputs.ReplicatorConsumerGroupReplication consumerGroupReplication,

            string sourceKafkaClusterArn,

            Pulumi.AwsNative.Msk.ReplicatorReplicationInfoTargetCompressionType targetCompressionType,

            string targetKafkaClusterArn,

            Outputs.ReplicatorTopicReplication topicReplication)
        {
            ConsumerGroupReplication = consumerGroupReplication;
            SourceKafkaClusterArn = sourceKafkaClusterArn;
            TargetCompressionType = targetCompressionType;
            TargetKafkaClusterArn = targetKafkaClusterArn;
            TopicReplication = topicReplication;
        }
    }
}
