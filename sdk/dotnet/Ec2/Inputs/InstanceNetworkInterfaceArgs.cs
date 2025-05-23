// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    public sealed class InstanceNetworkInterfaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Not currently supported by AWS CloudFormation.
        /// </summary>
        [Input("associateCarrierIpAddress")]
        public Input<bool>? AssociateCarrierIpAddress { get; set; }

        /// <summary>
        /// Indicates whether to assign a public IPv4 address to an instance you launch in a VPC.
        /// </summary>
        [Input("associatePublicIpAddress")]
        public Input<bool>? AssociatePublicIpAddress { get; set; }

        /// <summary>
        /// If set to true, the interface is deleted when the instance is terminated.
        /// </summary>
        [Input("deleteOnTermination")]
        public Input<bool>? DeleteOnTermination { get; set; }

        /// <summary>
        /// The description of the network interface.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The position of the network interface in the attachment order. A primary network interface has a device index of 0.
        /// </summary>
        [Input("deviceIndex", required: true)]
        public Input<string> DeviceIndex { get; set; } = null!;

        /// <summary>
        /// Configures ENA Express for UDP network traffic.
        /// </summary>
        [Input("enaSrdSpecification")]
        public Input<Inputs.InstanceEnaSrdSpecificationArgs>? EnaSrdSpecification { get; set; }

        [Input("groupSet")]
        private InputList<string>? _groupSet;

        /// <summary>
        /// The IDs of the security groups for the network interface.
        /// </summary>
        public InputList<string> GroupSet
        {
            get => _groupSet ?? (_groupSet = new InputList<string>());
            set => _groupSet = value;
        }

        /// <summary>
        /// A number of IPv6 addresses to assign to the network interface.
        /// </summary>
        [Input("ipv6AddressCount")]
        public Input<int>? Ipv6AddressCount { get; set; }

        [Input("ipv6Addresses")]
        private InputList<Inputs.InstanceIpv6AddressArgs>? _ipv6Addresses;

        /// <summary>
        /// The IPv6 addresses associated with the network interface.
        /// </summary>
        public InputList<Inputs.InstanceIpv6AddressArgs> Ipv6Addresses
        {
            get => _ipv6Addresses ?? (_ipv6Addresses = new InputList<Inputs.InstanceIpv6AddressArgs>());
            set => _ipv6Addresses = value;
        }

        /// <summary>
        /// The ID of the network interface.
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        /// <summary>
        /// The private IPv4 address of the network interface.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        [Input("privateIpAddresses")]
        private InputList<Inputs.InstancePrivateIpAddressSpecificationArgs>? _privateIpAddresses;

        /// <summary>
        /// One or more private IPv4 addresses to assign to the network interface.
        /// </summary>
        public InputList<Inputs.InstancePrivateIpAddressSpecificationArgs> PrivateIpAddresses
        {
            get => _privateIpAddresses ?? (_privateIpAddresses = new InputList<Inputs.InstancePrivateIpAddressSpecificationArgs>());
            set => _privateIpAddresses = value;
        }

        /// <summary>
        /// The number of secondary private IPv4 addresses.
        /// </summary>
        [Input("secondaryPrivateIpAddressCount")]
        public Input<int>? SecondaryPrivateIpAddressCount { get; set; }

        /// <summary>
        /// The ID of the subnet.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public InstanceNetworkInterfaceArgs()
        {
        }
        public static new InstanceNetworkInterfaceArgs Empty => new InstanceNetworkInterfaceArgs();
    }
}
