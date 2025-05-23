// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Eks.Inputs
{

    /// <summary>
    /// Configuration fields for specifying on-premises node and pod CIDRs that are external to the VPC passed during cluster creation.
    /// </summary>
    public sealed class ClusterRemoteNetworkConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("remoteNodeNetworks", required: true)]
        private InputList<Inputs.ClusterRemoteNodeNetworkArgs>? _remoteNodeNetworks;

        /// <summary>
        /// Network configuration of nodes run on-premises with EKS Hybrid Nodes.
        /// </summary>
        public InputList<Inputs.ClusterRemoteNodeNetworkArgs> RemoteNodeNetworks
        {
            get => _remoteNodeNetworks ?? (_remoteNodeNetworks = new InputList<Inputs.ClusterRemoteNodeNetworkArgs>());
            set => _remoteNodeNetworks = value;
        }

        [Input("remotePodNetworks")]
        private InputList<Inputs.ClusterRemotePodNetworkArgs>? _remotePodNetworks;

        /// <summary>
        /// Network configuration of pods run on-premises with EKS Hybrid Nodes.
        /// </summary>
        public InputList<Inputs.ClusterRemotePodNetworkArgs> RemotePodNetworks
        {
            get => _remotePodNetworks ?? (_remotePodNetworks = new InputList<Inputs.ClusterRemotePodNetworkArgs>());
            set => _remotePodNetworks = value;
        }

        public ClusterRemoteNetworkConfigArgs()
        {
        }
        public static new ClusterRemoteNetworkConfigArgs Empty => new ClusterRemoteNetworkConfigArgs();
    }
}
