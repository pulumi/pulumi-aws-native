# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs
from .. import outputs as _root_outputs

__all__ = [
    'GetNetworkInterfaceResult',
    'AwaitableGetNetworkInterfaceResult',
    'get_network_interface',
    'get_network_interface_output',
]

@pulumi.output_type
class GetNetworkInterfaceResult:
    def __init__(__self__, connection_tracking_specification=None, description=None, enable_primary_ipv6=None, group_set=None, id=None, ipv4_prefix_count=None, ipv4_prefixes=None, ipv6_address_count=None, ipv6_addresses=None, ipv6_prefix_count=None, ipv6_prefixes=None, primary_ipv6_address=None, primary_private_ip_address=None, private_ip_addresses=None, secondary_private_ip_address_count=None, secondary_private_ip_addresses=None, source_dest_check=None, tags=None, vpc_id=None):
        if connection_tracking_specification and not isinstance(connection_tracking_specification, dict):
            raise TypeError("Expected argument 'connection_tracking_specification' to be a dict")
        pulumi.set(__self__, "connection_tracking_specification", connection_tracking_specification)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enable_primary_ipv6 and not isinstance(enable_primary_ipv6, bool):
            raise TypeError("Expected argument 'enable_primary_ipv6' to be a bool")
        pulumi.set(__self__, "enable_primary_ipv6", enable_primary_ipv6)
        if group_set and not isinstance(group_set, list):
            raise TypeError("Expected argument 'group_set' to be a list")
        pulumi.set(__self__, "group_set", group_set)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ipv4_prefix_count and not isinstance(ipv4_prefix_count, int):
            raise TypeError("Expected argument 'ipv4_prefix_count' to be a int")
        pulumi.set(__self__, "ipv4_prefix_count", ipv4_prefix_count)
        if ipv4_prefixes and not isinstance(ipv4_prefixes, list):
            raise TypeError("Expected argument 'ipv4_prefixes' to be a list")
        pulumi.set(__self__, "ipv4_prefixes", ipv4_prefixes)
        if ipv6_address_count and not isinstance(ipv6_address_count, int):
            raise TypeError("Expected argument 'ipv6_address_count' to be a int")
        pulumi.set(__self__, "ipv6_address_count", ipv6_address_count)
        if ipv6_addresses and not isinstance(ipv6_addresses, list):
            raise TypeError("Expected argument 'ipv6_addresses' to be a list")
        pulumi.set(__self__, "ipv6_addresses", ipv6_addresses)
        if ipv6_prefix_count and not isinstance(ipv6_prefix_count, int):
            raise TypeError("Expected argument 'ipv6_prefix_count' to be a int")
        pulumi.set(__self__, "ipv6_prefix_count", ipv6_prefix_count)
        if ipv6_prefixes and not isinstance(ipv6_prefixes, list):
            raise TypeError("Expected argument 'ipv6_prefixes' to be a list")
        pulumi.set(__self__, "ipv6_prefixes", ipv6_prefixes)
        if primary_ipv6_address and not isinstance(primary_ipv6_address, str):
            raise TypeError("Expected argument 'primary_ipv6_address' to be a str")
        pulumi.set(__self__, "primary_ipv6_address", primary_ipv6_address)
        if primary_private_ip_address and not isinstance(primary_private_ip_address, str):
            raise TypeError("Expected argument 'primary_private_ip_address' to be a str")
        pulumi.set(__self__, "primary_private_ip_address", primary_private_ip_address)
        if private_ip_addresses and not isinstance(private_ip_addresses, list):
            raise TypeError("Expected argument 'private_ip_addresses' to be a list")
        pulumi.set(__self__, "private_ip_addresses", private_ip_addresses)
        if secondary_private_ip_address_count and not isinstance(secondary_private_ip_address_count, int):
            raise TypeError("Expected argument 'secondary_private_ip_address_count' to be a int")
        pulumi.set(__self__, "secondary_private_ip_address_count", secondary_private_ip_address_count)
        if secondary_private_ip_addresses and not isinstance(secondary_private_ip_addresses, list):
            raise TypeError("Expected argument 'secondary_private_ip_addresses' to be a list")
        pulumi.set(__self__, "secondary_private_ip_addresses", secondary_private_ip_addresses)
        if source_dest_check and not isinstance(source_dest_check, bool):
            raise TypeError("Expected argument 'source_dest_check' to be a bool")
        pulumi.set(__self__, "source_dest_check", source_dest_check)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="connectionTrackingSpecification")
    def connection_tracking_specification(self) -> Optional['outputs.NetworkInterfaceConnectionTrackingSpecification']:
        """
        A connection tracking specification for the network interface.
        """
        return pulumi.get(self, "connection_tracking_specification")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        A description for the network interface.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enablePrimaryIpv6")
    def enable_primary_ipv6(self) -> Optional[builtins.bool]:
        """
        If you have instances or ENIs that rely on the IPv6 address not changing, to avoid disrupting traffic to instances or ENIs, you can enable a primary IPv6 address. Enable this option to automatically assign an IPv6 associated with the ENI attached to your instance to be the primary IPv6 address. When you enable an IPv6 address to be a primary IPv6, you cannot disable it. Traffic will be routed to the primary IPv6 address until the instance is terminated or the ENI is detached. If you have multiple IPv6 addresses associated with an ENI and you enable a primary IPv6 address, the first IPv6 address associated with the ENI becomes the primary IPv6 address.
        """
        return pulumi.get(self, "enable_primary_ipv6")

    @property
    @pulumi.getter(name="groupSet")
    def group_set(self) -> Optional[Sequence[builtins.str]]:
        """
        A list of security group IDs associated with this network interface.
        """
        return pulumi.get(self, "group_set")

    @property
    @pulumi.getter
    def id(self) -> Optional[builtins.str]:
        """
        Network interface id.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipv4PrefixCount")
    def ipv4_prefix_count(self) -> Optional[builtins.int]:
        """
        The number of IPv4 prefixes to assign to a network interface. When you specify a number of IPv4 prefixes, Amazon EC2 selects these prefixes from your existing subnet CIDR reservations, if available, or from free spaces in the subnet. By default, these will be /28 prefixes. You can't specify a count of IPv4 prefixes if you've specified one of the following: specific IPv4 prefixes, specific private IPv4 addresses, or a count of private IPv4 addresses.
        """
        return pulumi.get(self, "ipv4_prefix_count")

    @property
    @pulumi.getter(name="ipv4Prefixes")
    def ipv4_prefixes(self) -> Optional[Sequence['outputs.NetworkInterfaceIpv4PrefixSpecification']]:
        """
        Assigns a list of IPv4 prefixes to the network interface. If you want EC2 to automatically assign IPv4 prefixes, use the Ipv4PrefixCount property and do not specify this property. Presently, only /28 prefixes are supported. You can't specify IPv4 prefixes if you've specified one of the following: a count of IPv4 prefixes, specific private IPv4 addresses, or a count of private IPv4 addresses.
        """
        return pulumi.get(self, "ipv4_prefixes")

    @property
    @pulumi.getter(name="ipv6AddressCount")
    def ipv6_address_count(self) -> Optional[builtins.int]:
        """
        The number of IPv6 addresses to assign to a network interface. Amazon EC2 automatically selects the IPv6 addresses from the subnet range. To specify specific IPv6 addresses, use the Ipv6Addresses property and don't specify this property.
        """
        return pulumi.get(self, "ipv6_address_count")

    @property
    @pulumi.getter(name="ipv6Addresses")
    def ipv6_addresses(self) -> Optional[Sequence['outputs.NetworkInterfaceInstanceIpv6Address']]:
        """
        One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet to associate with the network interface. If you're specifying a number of IPv6 addresses, use the Ipv6AddressCount property and don't specify this property.
        """
        return pulumi.get(self, "ipv6_addresses")

    @property
    @pulumi.getter(name="ipv6PrefixCount")
    def ipv6_prefix_count(self) -> Optional[builtins.int]:
        """
        The number of IPv6 prefixes to assign to a network interface. When you specify a number of IPv6 prefixes, Amazon EC2 selects these prefixes from your existing subnet CIDR reservations, if available, or from free spaces in the subnet. By default, these will be /80 prefixes. You can't specify a count of IPv6 prefixes if you've specified one of the following: specific IPv6 prefixes, specific IPv6 addresses, or a count of IPv6 addresses.
        """
        return pulumi.get(self, "ipv6_prefix_count")

    @property
    @pulumi.getter(name="ipv6Prefixes")
    def ipv6_prefixes(self) -> Optional[Sequence['outputs.NetworkInterfaceIpv6PrefixSpecification']]:
        """
        Assigns a list of IPv6 prefixes to the network interface. If you want EC2 to automatically assign IPv6 prefixes, use the Ipv6PrefixCount property and do not specify this property. Presently, only /80 prefixes are supported. You can't specify IPv6 prefixes if you've specified one of the following: a count of IPv6 prefixes, specific IPv6 addresses, or a count of IPv6 addresses.
        """
        return pulumi.get(self, "ipv6_prefixes")

    @property
    @pulumi.getter(name="primaryIpv6Address")
    def primary_ipv6_address(self) -> Optional[builtins.str]:
        """
        The primary IPv6 address
        """
        return pulumi.get(self, "primary_ipv6_address")

    @property
    @pulumi.getter(name="primaryPrivateIpAddress")
    def primary_private_ip_address(self) -> Optional[builtins.str]:
        """
        Returns the primary private IP address of the network interface.
        """
        return pulumi.get(self, "primary_private_ip_address")

    @property
    @pulumi.getter(name="privateIpAddresses")
    def private_ip_addresses(self) -> Optional[Sequence['outputs.NetworkInterfacePrivateIpAddressSpecification']]:
        """
        Assigns a list of private IP addresses to the network interface. You can specify a primary private IP address by setting the value of the Primary property to true in the PrivateIpAddressSpecification property. If you want EC2 to automatically assign private IP addresses, use the SecondaryPrivateIpAddressCount property and do not specify this property.
        """
        return pulumi.get(self, "private_ip_addresses")

    @property
    @pulumi.getter(name="secondaryPrivateIpAddressCount")
    def secondary_private_ip_address_count(self) -> Optional[builtins.int]:
        """
        The number of secondary private IPv4 addresses to assign to a network interface. When you specify a number of secondary IPv4 addresses, Amazon EC2 selects these IP addresses within the subnet's IPv4 CIDR range. You can't specify this option and specify more than one private IP address using privateIpAddresses
        """
        return pulumi.get(self, "secondary_private_ip_address_count")

    @property
    @pulumi.getter(name="secondaryPrivateIpAddresses")
    def secondary_private_ip_addresses(self) -> Optional[Sequence[builtins.str]]:
        """
        Returns the secondary private IP addresses of the network interface.
        """
        return pulumi.get(self, "secondary_private_ip_addresses")

    @property
    @pulumi.getter(name="sourceDestCheck")
    def source_dest_check(self) -> Optional[builtins.bool]:
        """
        Indicates whether traffic to or from the instance is validated.
        """
        return pulumi.get(self, "source_dest_check")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        An arbitrary set of tags (key-value pairs) for this network interface.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[builtins.str]:
        """
        The ID of the VPC
        """
        return pulumi.get(self, "vpc_id")


class AwaitableGetNetworkInterfaceResult(GetNetworkInterfaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkInterfaceResult(
            connection_tracking_specification=self.connection_tracking_specification,
            description=self.description,
            enable_primary_ipv6=self.enable_primary_ipv6,
            group_set=self.group_set,
            id=self.id,
            ipv4_prefix_count=self.ipv4_prefix_count,
            ipv4_prefixes=self.ipv4_prefixes,
            ipv6_address_count=self.ipv6_address_count,
            ipv6_addresses=self.ipv6_addresses,
            ipv6_prefix_count=self.ipv6_prefix_count,
            ipv6_prefixes=self.ipv6_prefixes,
            primary_ipv6_address=self.primary_ipv6_address,
            primary_private_ip_address=self.primary_private_ip_address,
            private_ip_addresses=self.private_ip_addresses,
            secondary_private_ip_address_count=self.secondary_private_ip_address_count,
            secondary_private_ip_addresses=self.secondary_private_ip_addresses,
            source_dest_check=self.source_dest_check,
            tags=self.tags,
            vpc_id=self.vpc_id)


def get_network_interface(id: Optional[builtins.str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkInterfaceResult:
    """
    The AWS::EC2::NetworkInterface resource creates network interface


    :param builtins.str id: Network interface id.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:ec2:getNetworkInterface', __args__, opts=opts, typ=GetNetworkInterfaceResult).value

    return AwaitableGetNetworkInterfaceResult(
        connection_tracking_specification=pulumi.get(__ret__, 'connection_tracking_specification'),
        description=pulumi.get(__ret__, 'description'),
        enable_primary_ipv6=pulumi.get(__ret__, 'enable_primary_ipv6'),
        group_set=pulumi.get(__ret__, 'group_set'),
        id=pulumi.get(__ret__, 'id'),
        ipv4_prefix_count=pulumi.get(__ret__, 'ipv4_prefix_count'),
        ipv4_prefixes=pulumi.get(__ret__, 'ipv4_prefixes'),
        ipv6_address_count=pulumi.get(__ret__, 'ipv6_address_count'),
        ipv6_addresses=pulumi.get(__ret__, 'ipv6_addresses'),
        ipv6_prefix_count=pulumi.get(__ret__, 'ipv6_prefix_count'),
        ipv6_prefixes=pulumi.get(__ret__, 'ipv6_prefixes'),
        primary_ipv6_address=pulumi.get(__ret__, 'primary_ipv6_address'),
        primary_private_ip_address=pulumi.get(__ret__, 'primary_private_ip_address'),
        private_ip_addresses=pulumi.get(__ret__, 'private_ip_addresses'),
        secondary_private_ip_address_count=pulumi.get(__ret__, 'secondary_private_ip_address_count'),
        secondary_private_ip_addresses=pulumi.get(__ret__, 'secondary_private_ip_addresses'),
        source_dest_check=pulumi.get(__ret__, 'source_dest_check'),
        tags=pulumi.get(__ret__, 'tags'),
        vpc_id=pulumi.get(__ret__, 'vpc_id'))
def get_network_interface_output(id: Optional[pulumi.Input[builtins.str]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetNetworkInterfaceResult]:
    """
    The AWS::EC2::NetworkInterface resource creates network interface


    :param builtins.str id: Network interface id.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:ec2:getNetworkInterface', __args__, opts=opts, typ=GetNetworkInterfaceResult)
    return __ret__.apply(lambda __response__: GetNetworkInterfaceResult(
        connection_tracking_specification=pulumi.get(__response__, 'connection_tracking_specification'),
        description=pulumi.get(__response__, 'description'),
        enable_primary_ipv6=pulumi.get(__response__, 'enable_primary_ipv6'),
        group_set=pulumi.get(__response__, 'group_set'),
        id=pulumi.get(__response__, 'id'),
        ipv4_prefix_count=pulumi.get(__response__, 'ipv4_prefix_count'),
        ipv4_prefixes=pulumi.get(__response__, 'ipv4_prefixes'),
        ipv6_address_count=pulumi.get(__response__, 'ipv6_address_count'),
        ipv6_addresses=pulumi.get(__response__, 'ipv6_addresses'),
        ipv6_prefix_count=pulumi.get(__response__, 'ipv6_prefix_count'),
        ipv6_prefixes=pulumi.get(__response__, 'ipv6_prefixes'),
        primary_ipv6_address=pulumi.get(__response__, 'primary_ipv6_address'),
        primary_private_ip_address=pulumi.get(__response__, 'primary_private_ip_address'),
        private_ip_addresses=pulumi.get(__response__, 'private_ip_addresses'),
        secondary_private_ip_address_count=pulumi.get(__response__, 'secondary_private_ip_address_count'),
        secondary_private_ip_addresses=pulumi.get(__response__, 'secondary_private_ip_addresses'),
        source_dest_check=pulumi.get(__response__, 'source_dest_check'),
        tags=pulumi.get(__response__, 'tags'),
        vpc_id=pulumi.get(__response__, 'vpc_id')))
