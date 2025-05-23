// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    /// <summary>
    /// ENA Express uses AWS Scalable Reliable Datagram (SRD) technology to increase the maximum bandwidth used per stream and minimize tail latency of network traffic between EC2 instances. With ENA Express, you can communicate between two EC2 instances in the same subnet within the same account, or in different accounts. Both sending and receiving instances must have ENA Express enabled.
    ///  To improve the reliability of network packet delivery, ENA Express reorders network packets on the receiving end by default. However, some UDP-based applications are designed to handle network packets that are out of order to reduce the overhead for packet delivery at the network layer. When ENA Express is enabled, you can specify whether UDP network traffic uses it.
    /// </summary>
    public sealed class NetworkInterfaceAttachmentEnaSrdSpecificationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether ENA Express is enabled for the network interface.
        /// </summary>
        [Input("enaSrdEnabled")]
        public Input<bool>? EnaSrdEnabled { get; set; }

        /// <summary>
        /// Configures ENA Express for UDP network traffic.
        /// </summary>
        [Input("enaSrdUdpSpecification")]
        public Input<Inputs.NetworkInterfaceAttachmentEnaSrdSpecificationEnaSrdUdpSpecificationPropertiesArgs>? EnaSrdUdpSpecification { get; set; }

        public NetworkInterfaceAttachmentEnaSrdSpecificationArgs()
        {
        }
        public static new NetworkInterfaceAttachmentEnaSrdSpecificationArgs Empty => new NetworkInterfaceAttachmentEnaSrdSpecificationArgs();
    }
}
