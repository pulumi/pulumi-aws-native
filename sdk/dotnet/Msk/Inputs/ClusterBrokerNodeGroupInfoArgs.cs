// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk.Inputs
{

    public sealed class ClusterBrokerNodeGroupInfoArgs : global::Pulumi.ResourceArgs
    {
        [Input("brokerAzDistribution")]
        public Input<string>? BrokerAzDistribution { get; set; }

        [Input("clientSubnets", required: true)]
        private InputList<string>? _clientSubnets;
        public InputList<string> ClientSubnets
        {
            get => _clientSubnets ?? (_clientSubnets = new InputList<string>());
            set => _clientSubnets = value;
        }

        [Input("connectivityInfo")]
        public Input<Inputs.ClusterConnectivityInfoArgs>? ConnectivityInfo { get; set; }

        /// <summary>
        /// The type of Amazon EC2 instances to use for brokers. The following instance types are allowed: kafka.m5.large, kafka.m5.xlarge, kafka.m5.2xlarge, kafka.m5.4xlarge, kafka.m5.8xlarge, kafka.m5.12xlarge, kafka.m5.16xlarge, kafka.m5.24xlarge, and kafka.t3.small.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("storageInfo")]
        public Input<Inputs.ClusterStorageInfoArgs>? StorageInfo { get; set; }

        public ClusterBrokerNodeGroupInfoArgs()
        {
        }
        public static new ClusterBrokerNodeGroupInfoArgs Empty => new ClusterBrokerNodeGroupInfoArgs();
    }
}
