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
    /// Details of a Kafka cluster for replication.
    /// </summary>
    [OutputType]
    public sealed class ReplicatorKafkaCluster
    {
        /// <summary>
        /// Details of an Amazon MSK cluster. Exactly one of AmazonMskCluster is required.
        /// </summary>
        public readonly Outputs.ReplicatorAmazonMskCluster AmazonMskCluster;
        /// <summary>
        /// Details of an Amazon VPC which has network connectivity to the Apache Kafka cluster.
        /// </summary>
        public readonly Outputs.ReplicatorKafkaClusterClientVpcConfig VpcConfig;

        [OutputConstructor]
        private ReplicatorKafkaCluster(
            Outputs.ReplicatorAmazonMskCluster amazonMskCluster,

            Outputs.ReplicatorKafkaClusterClientVpcConfig vpcConfig)
        {
            AmazonMskCluster = amazonMskCluster;
            VpcConfig = vpcConfig;
        }
    }
}
