// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaConnect.Inputs
{

    /// <summary>
    /// The network settings for a gateway.
    /// </summary>
    public sealed class GatewayNetworkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique IP address range to use for this network. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
        /// </summary>
        [Input("cidrBlock", required: true)]
        public Input<string> CidrBlock { get; set; } = null!;

        /// <summary>
        /// The name of the network. This name is used to reference the network and must be unique among networks in this gateway.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GatewayNetworkArgs()
        {
        }
        public static new GatewayNetworkArgs Empty => new GatewayNetworkArgs();
    }
}
