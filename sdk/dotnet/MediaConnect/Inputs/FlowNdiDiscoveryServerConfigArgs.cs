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
    /// Specifies the configuration settings for individual NDI discovery servers. A maximum of 3 servers is allowed.
    /// </summary>
    public sealed class FlowNdiDiscoveryServerConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique network address of the NDI discovery server.
        /// </summary>
        [Input("discoveryServerAddress", required: true)]
        public Input<string> DiscoveryServerAddress { get; set; } = null!;

        /// <summary>
        /// The port for the NDI discovery server. Defaults to 5959 if a custom port isn't specified.
        /// </summary>
        [Input("discoveryServerPort")]
        public Input<int>? DiscoveryServerPort { get; set; }

        /// <summary>
        /// The identifier for the Virtual Private Cloud (VPC) network interface used by the flow.
        /// </summary>
        [Input("vpcInterfaceAdapter", required: true)]
        public Input<string> VpcInterfaceAdapter { get; set; } = null!;

        public FlowNdiDiscoveryServerConfigArgs()
        {
        }
        public static new FlowNdiDiscoveryServerConfigArgs Empty => new FlowNdiDiscoveryServerConfigArgs();
    }
}
