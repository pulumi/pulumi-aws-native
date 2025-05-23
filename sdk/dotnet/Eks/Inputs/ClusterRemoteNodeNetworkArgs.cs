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
    /// Network configuration of nodes run on-premises with EKS Hybrid Nodes.
    /// </summary>
    public sealed class ClusterRemoteNodeNetworkArgs : global::Pulumi.ResourceArgs
    {
        [Input("cidrs", required: true)]
        private InputList<string>? _cidrs;

        /// <summary>
        /// Specifies the list of remote node CIDRs.
        /// </summary>
        public InputList<string> Cidrs
        {
            get => _cidrs ?? (_cidrs = new InputList<string>());
            set => _cidrs = value;
        }

        public ClusterRemoteNodeNetworkArgs()
        {
        }
        public static new ClusterRemoteNodeNetworkArgs Empty => new ClusterRemoteNodeNetworkArgs();
    }
}
