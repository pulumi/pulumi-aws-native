// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaConnect.Outputs
{

    /// <summary>
    /// The output of the bridge. A network output is delivered to your premises.
    /// </summary>
    [OutputType]
    public sealed class BridgeOutputResourceBridgeNetworkOutput
    {
        /// <summary>
        /// The network output IP Address.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// The network output's gateway network name.
        /// </summary>
        public readonly string NetworkName;
        /// <summary>
        /// The network output port.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The network output protocol.
        /// </summary>
        public readonly Pulumi.AwsNative.MediaConnect.BridgeOutputResourceBridgeNetworkOutputProtocol Protocol;
        /// <summary>
        /// The network output TTL.
        /// </summary>
        public readonly int Ttl;

        [OutputConstructor]
        private BridgeOutputResourceBridgeNetworkOutput(
            string ipAddress,

            string networkName,

            int port,

            Pulumi.AwsNative.MediaConnect.BridgeOutputResourceBridgeNetworkOutputProtocol protocol,

            int ttl)
        {
            IpAddress = ipAddress;
            NetworkName = networkName;
            Port = port;
            Protocol = protocol;
            Ttl = ttl;
        }
    }
}
