# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['LocalGatewayRouteArgs', 'LocalGatewayRoute']

@pulumi.input_type
class LocalGatewayRouteArgs:
    def __init__(__self__, *,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 local_gateway_route_table_id: Optional[pulumi.Input[str]] = None,
                 local_gateway_virtual_interface_group_id: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LocalGatewayRoute resource.
        :param pulumi.Input[str] destination_cidr_block: The CIDR block used for destination matches.
        :param pulumi.Input[str] local_gateway_route_table_id: The ID of the local gateway route table.
        :param pulumi.Input[str] local_gateway_virtual_interface_group_id: The ID of the virtual interface group.
        :param pulumi.Input[str] network_interface_id: The ID of the network interface.
        """
        if destination_cidr_block is not None:
            pulumi.set(__self__, "destination_cidr_block", destination_cidr_block)
        if local_gateway_route_table_id is not None:
            pulumi.set(__self__, "local_gateway_route_table_id", local_gateway_route_table_id)
        if local_gateway_virtual_interface_group_id is not None:
            pulumi.set(__self__, "local_gateway_virtual_interface_group_id", local_gateway_virtual_interface_group_id)
        if network_interface_id is not None:
            pulumi.set(__self__, "network_interface_id", network_interface_id)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The CIDR block used for destination matches.
        """
        return pulumi.get(self, "destination_cidr_block")

    @destination_cidr_block.setter
    def destination_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_cidr_block", value)

    @property
    @pulumi.getter(name="localGatewayRouteTableId")
    def local_gateway_route_table_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the local gateway route table.
        """
        return pulumi.get(self, "local_gateway_route_table_id")

    @local_gateway_route_table_id.setter
    def local_gateway_route_table_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_gateway_route_table_id", value)

    @property
    @pulumi.getter(name="localGatewayVirtualInterfaceGroupId")
    def local_gateway_virtual_interface_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the virtual interface group.
        """
        return pulumi.get(self, "local_gateway_virtual_interface_group_id")

    @local_gateway_virtual_interface_group_id.setter
    def local_gateway_virtual_interface_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_gateway_virtual_interface_group_id", value)

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the network interface.
        """
        return pulumi.get(self, "network_interface_id")

    @network_interface_id.setter
    def network_interface_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_interface_id", value)


class LocalGatewayRoute(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 local_gateway_route_table_id: Optional[pulumi.Input[str]] = None,
                 local_gateway_virtual_interface_group_id: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Describes a route for a local gateway route table.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_cidr_block: The CIDR block used for destination matches.
        :param pulumi.Input[str] local_gateway_route_table_id: The ID of the local gateway route table.
        :param pulumi.Input[str] local_gateway_virtual_interface_group_id: The ID of the virtual interface group.
        :param pulumi.Input[str] network_interface_id: The ID of the network interface.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[LocalGatewayRouteArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Describes a route for a local gateway route table.

        :param str resource_name: The name of the resource.
        :param LocalGatewayRouteArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LocalGatewayRouteArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 local_gateway_route_table_id: Optional[pulumi.Input[str]] = None,
                 local_gateway_virtual_interface_group_id: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LocalGatewayRouteArgs.__new__(LocalGatewayRouteArgs)

            __props__.__dict__["destination_cidr_block"] = destination_cidr_block
            __props__.__dict__["local_gateway_route_table_id"] = local_gateway_route_table_id
            __props__.__dict__["local_gateway_virtual_interface_group_id"] = local_gateway_virtual_interface_group_id
            __props__.__dict__["network_interface_id"] = network_interface_id
            __props__.__dict__["state"] = None
            __props__.__dict__["type"] = None
        super(LocalGatewayRoute, __self__).__init__(
            'aws-native:ec2:LocalGatewayRoute',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'LocalGatewayRoute':
        """
        Get an existing LocalGatewayRoute resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = LocalGatewayRouteArgs.__new__(LocalGatewayRouteArgs)

        __props__.__dict__["destination_cidr_block"] = None
        __props__.__dict__["local_gateway_route_table_id"] = None
        __props__.__dict__["local_gateway_virtual_interface_group_id"] = None
        __props__.__dict__["network_interface_id"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["type"] = None
        return LocalGatewayRoute(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> pulumi.Output[Optional[str]]:
        """
        The CIDR block used for destination matches.
        """
        return pulumi.get(self, "destination_cidr_block")

    @property
    @pulumi.getter(name="localGatewayRouteTableId")
    def local_gateway_route_table_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the local gateway route table.
        """
        return pulumi.get(self, "local_gateway_route_table_id")

    @property
    @pulumi.getter(name="localGatewayVirtualInterfaceGroupId")
    def local_gateway_virtual_interface_group_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the virtual interface group.
        """
        return pulumi.get(self, "local_gateway_virtual_interface_group_id")

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the network interface.
        """
        return pulumi.get(self, "network_interface_id")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The state of the route.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The route type.
        """
        return pulumi.get(self, "type")
