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
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs

__all__ = ['NatGatewayArgs', 'NatGateway']

@pulumi.input_type
class NatGatewayArgs:
    def __init__(__self__, *,
                 subnet_id: pulumi.Input[builtins.str],
                 allocation_id: Optional[pulumi.Input[builtins.str]] = None,
                 connectivity_type: Optional[pulumi.Input[builtins.str]] = None,
                 max_drain_duration_seconds: Optional[pulumi.Input[builtins.int]] = None,
                 private_ip_address: Optional[pulumi.Input[builtins.str]] = None,
                 secondary_allocation_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 secondary_private_ip_address_count: Optional[pulumi.Input[builtins.int]] = None,
                 secondary_private_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a NatGateway resource.
        :param pulumi.Input[builtins.str] subnet_id: The ID of the subnet in which the NAT gateway is located.
        :param pulumi.Input[builtins.str] allocation_id: [Public NAT gateway only] The allocation ID of the Elastic IP address that's associated with the NAT gateway. This property is required for a public NAT gateway and cannot be specified with a private NAT gateway.
        :param pulumi.Input[builtins.str] connectivity_type: Indicates whether the NAT gateway supports public or private connectivity. The default is public connectivity.
        :param pulumi.Input[builtins.int] max_drain_duration_seconds: The maximum amount of time to wait (in seconds) before forcibly releasing the IP addresses if connections are still in progress. Default value is 350 seconds.
        :param pulumi.Input[builtins.str] private_ip_address: The private IPv4 address to assign to the NAT gateway. If you don't provide an address, a private IPv4 address will be automatically assigned.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] secondary_allocation_ids: Secondary EIP allocation IDs. For more information, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateway-working-with.html) in the *Amazon VPC User Guide*.
        :param pulumi.Input[builtins.int] secondary_private_ip_address_count: [Private NAT gateway only] The number of secondary private IPv4 addresses you want to assign to the NAT gateway. For more information about secondary addresses, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-creating) in the *Amazon Virtual Private Cloud User Guide*.
                ``SecondaryPrivateIpAddressCount`` and ``SecondaryPrivateIpAddresses`` cannot be set at the same time.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] secondary_private_ip_addresses: Secondary private IPv4 addresses. For more information about secondary addresses, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-creating) in the *Amazon Virtual Private Cloud User Guide*.
                ``SecondaryPrivateIpAddressCount`` and ``SecondaryPrivateIpAddresses`` cannot be set at the same time.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: The tags for the NAT gateway.
        """
        pulumi.set(__self__, "subnet_id", subnet_id)
        if allocation_id is not None:
            pulumi.set(__self__, "allocation_id", allocation_id)
        if connectivity_type is not None:
            pulumi.set(__self__, "connectivity_type", connectivity_type)
        if max_drain_duration_seconds is not None:
            pulumi.set(__self__, "max_drain_duration_seconds", max_drain_duration_seconds)
        if private_ip_address is not None:
            pulumi.set(__self__, "private_ip_address", private_ip_address)
        if secondary_allocation_ids is not None:
            pulumi.set(__self__, "secondary_allocation_ids", secondary_allocation_ids)
        if secondary_private_ip_address_count is not None:
            pulumi.set(__self__, "secondary_private_ip_address_count", secondary_private_ip_address_count)
        if secondary_private_ip_addresses is not None:
            pulumi.set(__self__, "secondary_private_ip_addresses", secondary_private_ip_addresses)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the subnet in which the NAT gateway is located.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="allocationId")
    def allocation_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        [Public NAT gateway only] The allocation ID of the Elastic IP address that's associated with the NAT gateway. This property is required for a public NAT gateway and cannot be specified with a private NAT gateway.
        """
        return pulumi.get(self, "allocation_id")

    @allocation_id.setter
    def allocation_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "allocation_id", value)

    @property
    @pulumi.getter(name="connectivityType")
    def connectivity_type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Indicates whether the NAT gateway supports public or private connectivity. The default is public connectivity.
        """
        return pulumi.get(self, "connectivity_type")

    @connectivity_type.setter
    def connectivity_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "connectivity_type", value)

    @property
    @pulumi.getter(name="maxDrainDurationSeconds")
    def max_drain_duration_seconds(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The maximum amount of time to wait (in seconds) before forcibly releasing the IP addresses if connections are still in progress. Default value is 350 seconds.
        """
        return pulumi.get(self, "max_drain_duration_seconds")

    @max_drain_duration_seconds.setter
    def max_drain_duration_seconds(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "max_drain_duration_seconds", value)

    @property
    @pulumi.getter(name="privateIpAddress")
    def private_ip_address(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The private IPv4 address to assign to the NAT gateway. If you don't provide an address, a private IPv4 address will be automatically assigned.
        """
        return pulumi.get(self, "private_ip_address")

    @private_ip_address.setter
    def private_ip_address(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "private_ip_address", value)

    @property
    @pulumi.getter(name="secondaryAllocationIds")
    def secondary_allocation_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Secondary EIP allocation IDs. For more information, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateway-working-with.html) in the *Amazon VPC User Guide*.
        """
        return pulumi.get(self, "secondary_allocation_ids")

    @secondary_allocation_ids.setter
    def secondary_allocation_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "secondary_allocation_ids", value)

    @property
    @pulumi.getter(name="secondaryPrivateIpAddressCount")
    def secondary_private_ip_address_count(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        [Private NAT gateway only] The number of secondary private IPv4 addresses you want to assign to the NAT gateway. For more information about secondary addresses, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-creating) in the *Amazon Virtual Private Cloud User Guide*.
         ``SecondaryPrivateIpAddressCount`` and ``SecondaryPrivateIpAddresses`` cannot be set at the same time.
        """
        return pulumi.get(self, "secondary_private_ip_address_count")

    @secondary_private_ip_address_count.setter
    def secondary_private_ip_address_count(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "secondary_private_ip_address_count", value)

    @property
    @pulumi.getter(name="secondaryPrivateIpAddresses")
    def secondary_private_ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Secondary private IPv4 addresses. For more information about secondary addresses, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-creating) in the *Amazon Virtual Private Cloud User Guide*.
         ``SecondaryPrivateIpAddressCount`` and ``SecondaryPrivateIpAddresses`` cannot be set at the same time.
        """
        return pulumi.get(self, "secondary_private_ip_addresses")

    @secondary_private_ip_addresses.setter
    def secondary_private_ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "secondary_private_ip_addresses", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        The tags for the NAT gateway.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:ec2:NatGateway")
class NatGateway(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allocation_id: Optional[pulumi.Input[builtins.str]] = None,
                 connectivity_type: Optional[pulumi.Input[builtins.str]] = None,
                 max_drain_duration_seconds: Optional[pulumi.Input[builtins.int]] = None,
                 private_ip_address: Optional[pulumi.Input[builtins.str]] = None,
                 secondary_allocation_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 secondary_private_ip_address_count: Optional[pulumi.Input[builtins.int]] = None,
                 secondary_private_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 subnet_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        """
        Specifies a network address translation (NAT) gateway in the specified subnet. You can create either a public NAT gateway or a private NAT gateway. The default is a public NAT gateway. If you create a public NAT gateway, you must specify an elastic IP address.
         With a NAT gateway, instances in a private subnet can connect to the internet, other AWS services, or an on-premises network using the IP address of the NAT gateway. For more information, see [NAT gateways](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html) in the *Amazon VPC User Guide*.
         If you add a default route (``AWS::EC2::Route`` resource) that points to a NAT gateway, specify the NAT gateway ID for the route's ``NatGatewayId`` property.
          When you associate an Elastic IP address or secondary Elastic IP address with a public NAT gateway, the network border group of the Elastic IP address must match the network border group of the Availability Zone (AZ) that the public NAT gateway is in. Otherwise, the NAT gateway fails to launch. You can see the network border group for the AZ by viewing the details of the subnet. Similarly, you can view the network border group for the Elastic IP address by viewing its details. For more information, see [Allocate an Elastic IP address](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-eips.html#allocate-eip) in the *Amazon VPC User Guide*.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] allocation_id: [Public NAT gateway only] The allocation ID of the Elastic IP address that's associated with the NAT gateway. This property is required for a public NAT gateway and cannot be specified with a private NAT gateway.
        :param pulumi.Input[builtins.str] connectivity_type: Indicates whether the NAT gateway supports public or private connectivity. The default is public connectivity.
        :param pulumi.Input[builtins.int] max_drain_duration_seconds: The maximum amount of time to wait (in seconds) before forcibly releasing the IP addresses if connections are still in progress. Default value is 350 seconds.
        :param pulumi.Input[builtins.str] private_ip_address: The private IPv4 address to assign to the NAT gateway. If you don't provide an address, a private IPv4 address will be automatically assigned.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] secondary_allocation_ids: Secondary EIP allocation IDs. For more information, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateway-working-with.html) in the *Amazon VPC User Guide*.
        :param pulumi.Input[builtins.int] secondary_private_ip_address_count: [Private NAT gateway only] The number of secondary private IPv4 addresses you want to assign to the NAT gateway. For more information about secondary addresses, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-creating) in the *Amazon Virtual Private Cloud User Guide*.
                ``SecondaryPrivateIpAddressCount`` and ``SecondaryPrivateIpAddresses`` cannot be set at the same time.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] secondary_private_ip_addresses: Secondary private IPv4 addresses. For more information about secondary addresses, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-creating) in the *Amazon Virtual Private Cloud User Guide*.
                ``SecondaryPrivateIpAddressCount`` and ``SecondaryPrivateIpAddresses`` cannot be set at the same time.
        :param pulumi.Input[builtins.str] subnet_id: The ID of the subnet in which the NAT gateway is located.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] tags: The tags for the NAT gateway.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NatGatewayArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Specifies a network address translation (NAT) gateway in the specified subnet. You can create either a public NAT gateway or a private NAT gateway. The default is a public NAT gateway. If you create a public NAT gateway, you must specify an elastic IP address.
         With a NAT gateway, instances in a private subnet can connect to the internet, other AWS services, or an on-premises network using the IP address of the NAT gateway. For more information, see [NAT gateways](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html) in the *Amazon VPC User Guide*.
         If you add a default route (``AWS::EC2::Route`` resource) that points to a NAT gateway, specify the NAT gateway ID for the route's ``NatGatewayId`` property.
          When you associate an Elastic IP address or secondary Elastic IP address with a public NAT gateway, the network border group of the Elastic IP address must match the network border group of the Availability Zone (AZ) that the public NAT gateway is in. Otherwise, the NAT gateway fails to launch. You can see the network border group for the AZ by viewing the details of the subnet. Similarly, you can view the network border group for the Elastic IP address by viewing its details. For more information, see [Allocate an Elastic IP address](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-eips.html#allocate-eip) in the *Amazon VPC User Guide*.

        :param str resource_name: The name of the resource.
        :param NatGatewayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NatGatewayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allocation_id: Optional[pulumi.Input[builtins.str]] = None,
                 connectivity_type: Optional[pulumi.Input[builtins.str]] = None,
                 max_drain_duration_seconds: Optional[pulumi.Input[builtins.int]] = None,
                 private_ip_address: Optional[pulumi.Input[builtins.str]] = None,
                 secondary_allocation_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 secondary_private_ip_address_count: Optional[pulumi.Input[builtins.int]] = None,
                 secondary_private_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 subnet_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NatGatewayArgs.__new__(NatGatewayArgs)

            __props__.__dict__["allocation_id"] = allocation_id
            __props__.__dict__["connectivity_type"] = connectivity_type
            __props__.__dict__["max_drain_duration_seconds"] = max_drain_duration_seconds
            __props__.__dict__["private_ip_address"] = private_ip_address
            __props__.__dict__["secondary_allocation_ids"] = secondary_allocation_ids
            __props__.__dict__["secondary_private_ip_address_count"] = secondary_private_ip_address_count
            __props__.__dict__["secondary_private_ip_addresses"] = secondary_private_ip_addresses
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["nat_gateway_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["allocationId", "connectivityType", "privateIpAddress", "subnetId"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(NatGateway, __self__).__init__(
            'aws-native:ec2:NatGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'NatGateway':
        """
        Get an existing NatGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = NatGatewayArgs.__new__(NatGatewayArgs)

        __props__.__dict__["allocation_id"] = None
        __props__.__dict__["connectivity_type"] = None
        __props__.__dict__["max_drain_duration_seconds"] = None
        __props__.__dict__["nat_gateway_id"] = None
        __props__.__dict__["private_ip_address"] = None
        __props__.__dict__["secondary_allocation_ids"] = None
        __props__.__dict__["secondary_private_ip_address_count"] = None
        __props__.__dict__["secondary_private_ip_addresses"] = None
        __props__.__dict__["subnet_id"] = None
        __props__.__dict__["tags"] = None
        return NatGateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allocationId")
    def allocation_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        [Public NAT gateway only] The allocation ID of the Elastic IP address that's associated with the NAT gateway. This property is required for a public NAT gateway and cannot be specified with a private NAT gateway.
        """
        return pulumi.get(self, "allocation_id")

    @property
    @pulumi.getter(name="connectivityType")
    def connectivity_type(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Indicates whether the NAT gateway supports public or private connectivity. The default is public connectivity.
        """
        return pulumi.get(self, "connectivity_type")

    @property
    @pulumi.getter(name="maxDrainDurationSeconds")
    def max_drain_duration_seconds(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        The maximum amount of time to wait (in seconds) before forcibly releasing the IP addresses if connections are still in progress. Default value is 350 seconds.
        """
        return pulumi.get(self, "max_drain_duration_seconds")

    @property
    @pulumi.getter(name="natGatewayId")
    def nat_gateway_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the NAT gateway.
        """
        return pulumi.get(self, "nat_gateway_id")

    @property
    @pulumi.getter(name="privateIpAddress")
    def private_ip_address(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The private IPv4 address to assign to the NAT gateway. If you don't provide an address, a private IPv4 address will be automatically assigned.
        """
        return pulumi.get(self, "private_ip_address")

    @property
    @pulumi.getter(name="secondaryAllocationIds")
    def secondary_allocation_ids(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        Secondary EIP allocation IDs. For more information, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateway-working-with.html) in the *Amazon VPC User Guide*.
        """
        return pulumi.get(self, "secondary_allocation_ids")

    @property
    @pulumi.getter(name="secondaryPrivateIpAddressCount")
    def secondary_private_ip_address_count(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        [Private NAT gateway only] The number of secondary private IPv4 addresses you want to assign to the NAT gateway. For more information about secondary addresses, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-creating) in the *Amazon Virtual Private Cloud User Guide*.
         ``SecondaryPrivateIpAddressCount`` and ``SecondaryPrivateIpAddresses`` cannot be set at the same time.
        """
        return pulumi.get(self, "secondary_private_ip_address_count")

    @property
    @pulumi.getter(name="secondaryPrivateIpAddresses")
    def secondary_private_ip_addresses(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        Secondary private IPv4 addresses. For more information about secondary addresses, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-creating) in the *Amazon Virtual Private Cloud User Guide*.
         ``SecondaryPrivateIpAddressCount`` and ``SecondaryPrivateIpAddresses`` cannot be set at the same time.
        """
        return pulumi.get(self, "secondary_private_ip_addresses")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the subnet in which the NAT gateway is located.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        The tags for the NAT gateway.
        """
        return pulumi.get(self, "tags")

