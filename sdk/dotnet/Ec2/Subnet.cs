// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    /// <summary>
    /// Specifies a subnet for the specified VPC.
    ///  For an IPv4 only subnet, specify an IPv4 CIDR block. If the VPC has an IPv6 CIDR block, you can create an IPv6 only subnet or a dual stack subnet instead. For an IPv6 only subnet, specify an IPv6 CIDR block. For a dual stack subnet, specify both an IPv4 CIDR block and an IPv6 CIDR block.
    ///  For more information, see [Subnets for your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/configure-subnets.html) in the *Amazon VPC User Guide*.
    /// </summary>
    [AwsNativeResourceType("aws-native:ec2:Subnet")]
    public partial class Subnet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates whether a network interface created in this subnet receives an IPv6 address. The default value is ``false``.
        ///  If you specify ``AssignIpv6AddressOnCreation``, you must also specify an IPv6 CIDR block.
        /// </summary>
        [Output("assignIpv6AddressOnCreation")]
        public Output<bool?> AssignIpv6AddressOnCreation { get; private set; } = null!;

        /// <summary>
        /// The Availability Zone of the subnet.
        ///  If you update this property, you must also update the ``CidrBlock`` property.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string?> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// The AZ ID of the subnet.
        /// </summary>
        [Output("availabilityZoneId")]
        public Output<string?> AvailabilityZoneId { get; private set; } = null!;

        [Output("blockPublicAccessStates")]
        public Output<Outputs.BlockPublicAccessStatesProperties> BlockPublicAccessStates { get; private set; } = null!;

        /// <summary>
        /// The IPv4 CIDR block assigned to the subnet.
        ///  If you update this property, we create a new subnet, and then delete the existing one.
        /// </summary>
        [Output("cidrBlock")]
        public Output<string?> CidrBlock { get; private set; } = null!;

        /// <summary>
        /// Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations.
        ///   You must first configure a NAT gateway in a public subnet (separate from the subnet containing the IPv6-only workloads). For example, the subnet containing the NAT gateway should have a ``0.0.0.0/0`` route pointing to the internet gateway. For more information, see [Configure DNS64 and NAT64](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateway-nat64-dns64.html#nat-gateway-nat64-dns64-walkthrough) in the *User Guide*.
        /// </summary>
        [Output("enableDns64")]
        public Output<bool?> EnableDns64 { get; private set; } = null!;

        /// <summary>
        /// Indicates the device position for local network interfaces in this subnet. For example, ``1`` indicates local network interfaces in this subnet are the secondary network interface (eth1).
        /// </summary>
        [Output("enableLniAtDeviceIndex")]
        public Output<int?> EnableLniAtDeviceIndex { get; private set; } = null!;

        /// <summary>
        /// An IPv4 IPAM pool ID for the subnet.
        /// </summary>
        [Output("ipv4IpamPoolId")]
        public Output<string?> Ipv4IpamPoolId { get; private set; } = null!;

        /// <summary>
        /// An IPv4 netmask length for the subnet.
        /// </summary>
        [Output("ipv4NetmaskLength")]
        public Output<int?> Ipv4NetmaskLength { get; private set; } = null!;

        /// <summary>
        /// The IPv6 CIDR block.
        ///  If you specify ``AssignIpv6AddressOnCreation``, you must also specify an IPv6 CIDR block.
        /// </summary>
        [Output("ipv6CidrBlock")]
        public Output<string?> Ipv6CidrBlock { get; private set; } = null!;

        /// <summary>
        /// The IPv6 CIDR blocks that are associated with the subnet.
        /// </summary>
        [Output("ipv6CidrBlocks")]
        public Output<ImmutableArray<string>> Ipv6CidrBlocks { get; private set; } = null!;

        /// <summary>
        /// An IPv6 IPAM pool ID for the subnet.
        /// </summary>
        [Output("ipv6IpamPoolId")]
        public Output<string?> Ipv6IpamPoolId { get; private set; } = null!;

        /// <summary>
        /// Indicates whether this is an IPv6 only subnet. For more information, see [Subnet basics](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html#subnet-basics) in the *User Guide*.
        /// </summary>
        [Output("ipv6Native")]
        public Output<bool?> Ipv6Native { get; private set; } = null!;

        /// <summary>
        /// An IPv6 netmask length for the subnet.
        /// </summary>
        [Output("ipv6NetmaskLength")]
        public Output<int?> Ipv6NetmaskLength { get; private set; } = null!;

        /// <summary>
        /// Indicates whether instances launched in this subnet receive a public IPv4 address. The default value is ``false``.
        ///  AWS charges for all public IPv4 addresses, including public IPv4 addresses associated with running instances and Elastic IP addresses. For more information, see the *Public IPv4 Address* tab on the [VPC pricing page](https://docs.aws.amazon.com/vpc/pricing/).
        /// </summary>
        [Output("mapPublicIpOnLaunch")]
        public Output<bool?> MapPublicIpOnLaunch { get; private set; } = null!;

        /// <summary>
        /// The ID of the network ACL that is associated with the subnet's VPC, such as `acl-5fb85d36` .
        /// </summary>
        [Output("networkAclAssociationId")]
        public Output<string> NetworkAclAssociationId { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Outpost.
        /// </summary>
        [Output("outpostArn")]
        public Output<string?> OutpostArn { get; private set; } = null!;

        /// <summary>
        /// The hostname type for EC2 instances launched into this subnet and how DNS A and AAAA record queries to the instances should be handled. For more information, see [Amazon EC2 instance hostname types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html) in the *User Guide*.
        ///  Available options:
        ///   +  EnableResourceNameDnsAAAARecord (true | false)
        ///   +  EnableResourceNameDnsARecord (true | false)
        ///   +  HostnameType (ip-name | resource-name)
        /// </summary>
        [Output("privateDnsNameOptionsOnLaunch")]
        public Output<Outputs.PrivateDnsNameOptionsOnLaunchProperties?> PrivateDnsNameOptionsOnLaunch { get; private set; } = null!;

        /// <summary>
        /// The ID of the subnet.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// Any tags assigned to the subnet.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPC the subnet is in.
        ///  If you update this property, you must also update the ``CidrBlock`` property.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a Subnet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Subnet(string name, SubnetArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ec2:Subnet", name, args ?? new SubnetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Subnet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:Subnet", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "availabilityZone",
                    "availabilityZoneId",
                    "cidrBlock",
                    "ipv4IpamPoolId",
                    "ipv4NetmaskLength",
                    "ipv6IpamPoolId",
                    "ipv6Native",
                    "ipv6NetmaskLength",
                    "outpostArn",
                    "vpcId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Subnet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Subnet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Subnet(name, id, options);
        }
    }

    public sealed class SubnetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether a network interface created in this subnet receives an IPv6 address. The default value is ``false``.
        ///  If you specify ``AssignIpv6AddressOnCreation``, you must also specify an IPv6 CIDR block.
        /// </summary>
        [Input("assignIpv6AddressOnCreation")]
        public Input<bool>? AssignIpv6AddressOnCreation { get; set; }

        /// <summary>
        /// The Availability Zone of the subnet.
        ///  If you update this property, you must also update the ``CidrBlock`` property.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// The AZ ID of the subnet.
        /// </summary>
        [Input("availabilityZoneId")]
        public Input<string>? AvailabilityZoneId { get; set; }

        /// <summary>
        /// The IPv4 CIDR block assigned to the subnet.
        ///  If you update this property, we create a new subnet, and then delete the existing one.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations.
        ///   You must first configure a NAT gateway in a public subnet (separate from the subnet containing the IPv6-only workloads). For example, the subnet containing the NAT gateway should have a ``0.0.0.0/0`` route pointing to the internet gateway. For more information, see [Configure DNS64 and NAT64](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateway-nat64-dns64.html#nat-gateway-nat64-dns64-walkthrough) in the *User Guide*.
        /// </summary>
        [Input("enableDns64")]
        public Input<bool>? EnableDns64 { get; set; }

        /// <summary>
        /// Indicates the device position for local network interfaces in this subnet. For example, ``1`` indicates local network interfaces in this subnet are the secondary network interface (eth1).
        /// </summary>
        [Input("enableLniAtDeviceIndex")]
        public Input<int>? EnableLniAtDeviceIndex { get; set; }

        /// <summary>
        /// An IPv4 IPAM pool ID for the subnet.
        /// </summary>
        [Input("ipv4IpamPoolId")]
        public Input<string>? Ipv4IpamPoolId { get; set; }

        /// <summary>
        /// An IPv4 netmask length for the subnet.
        /// </summary>
        [Input("ipv4NetmaskLength")]
        public Input<int>? Ipv4NetmaskLength { get; set; }

        /// <summary>
        /// The IPv6 CIDR block.
        ///  If you specify ``AssignIpv6AddressOnCreation``, you must also specify an IPv6 CIDR block.
        /// </summary>
        [Input("ipv6CidrBlock")]
        public Input<string>? Ipv6CidrBlock { get; set; }

        /// <summary>
        /// An IPv6 IPAM pool ID for the subnet.
        /// </summary>
        [Input("ipv6IpamPoolId")]
        public Input<string>? Ipv6IpamPoolId { get; set; }

        /// <summary>
        /// Indicates whether this is an IPv6 only subnet. For more information, see [Subnet basics](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html#subnet-basics) in the *User Guide*.
        /// </summary>
        [Input("ipv6Native")]
        public Input<bool>? Ipv6Native { get; set; }

        /// <summary>
        /// An IPv6 netmask length for the subnet.
        /// </summary>
        [Input("ipv6NetmaskLength")]
        public Input<int>? Ipv6NetmaskLength { get; set; }

        /// <summary>
        /// Indicates whether instances launched in this subnet receive a public IPv4 address. The default value is ``false``.
        ///  AWS charges for all public IPv4 addresses, including public IPv4 addresses associated with running instances and Elastic IP addresses. For more information, see the *Public IPv4 Address* tab on the [VPC pricing page](https://docs.aws.amazon.com/vpc/pricing/).
        /// </summary>
        [Input("mapPublicIpOnLaunch")]
        public Input<bool>? MapPublicIpOnLaunch { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Outpost.
        /// </summary>
        [Input("outpostArn")]
        public Input<string>? OutpostArn { get; set; }

        /// <summary>
        /// The hostname type for EC2 instances launched into this subnet and how DNS A and AAAA record queries to the instances should be handled. For more information, see [Amazon EC2 instance hostname types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html) in the *User Guide*.
        ///  Available options:
        ///   +  EnableResourceNameDnsAAAARecord (true | false)
        ///   +  EnableResourceNameDnsARecord (true | false)
        ///   +  HostnameType (ip-name | resource-name)
        /// </summary>
        [Input("privateDnsNameOptionsOnLaunch")]
        public Input<Inputs.PrivateDnsNameOptionsOnLaunchPropertiesArgs>? PrivateDnsNameOptionsOnLaunch { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// Any tags assigned to the subnet.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the VPC the subnet is in.
        ///  If you update this property, you must also update the ``CidrBlock`` property.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public SubnetArgs()
        {
        }
        public static new SubnetArgs Empty => new SubnetArgs();
    }
}
