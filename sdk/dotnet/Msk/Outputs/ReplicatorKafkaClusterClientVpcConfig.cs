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
    /// Details of an Amazon VPC which has network connectivity to the Kafka cluster.
    /// </summary>
    [OutputType]
    public sealed class ReplicatorKafkaClusterClientVpcConfig
    {
        /// <summary>
        /// The AWS security groups to associate with the elastic network interfaces in order to specify what the replicator has access to. If a security group is not specified, the default security group associated with the VPC is used.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// The list of subnets to connect to in the virtual private cloud (VPC). AWS creates elastic network interfaces inside these subnets.
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;

        [OutputConstructor]
        private ReplicatorKafkaClusterClientVpcConfig(
            ImmutableArray<string> securityGroupIds,

            ImmutableArray<string> subnetIds)
        {
            SecurityGroupIds = securityGroupIds;
            SubnetIds = subnetIds;
        }
    }
}
