// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaConnect.Outputs
{

    /// <summary>
    /// The source of the bridge. A network source originates at your premises.
    /// </summary>
    [OutputType]
    public sealed class BridgeNetworkSource
    {
        /// <summary>
        /// The network source multicast IP.
        /// </summary>
        public readonly string MulticastIp;
        /// <summary>
        /// The name of the network source.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The network source's gateway network name.
        /// </summary>
        public readonly string NetworkName;
        /// <summary>
        /// The network source port.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The network source protocol.
        /// </summary>
        public readonly Pulumi.AwsNative.MediaConnect.BridgeProtocolEnum Protocol;

        [OutputConstructor]
        private BridgeNetworkSource(
            string multicastIp,

            string name,

            string networkName,

            int port,

            Pulumi.AwsNative.MediaConnect.BridgeProtocolEnum protocol)
        {
            MulticastIp = multicastIp;
            Name = name;
            NetworkName = networkName;
            Port = port;
            Protocol = protocol;
        }
    }
}