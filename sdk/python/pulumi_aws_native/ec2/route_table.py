# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['RouteTableArgs', 'RouteTable']

@pulumi.input_type
class RouteTableArgs:
    def __init__(__self__, *,
                 vpc_id: pulumi.Input[str],
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['RouteTableTagArgs']]]] = None):
        """
        The set of arguments for constructing a RouteTable resource.
        :param pulumi.Input[str] vpc_id: The ID of the VPC.
        :param pulumi.Input[Sequence[pulumi.Input['RouteTableTagArgs']]] tags: Any tags assigned to the route table.
        """
        pulumi.set(__self__, "vpc_id", vpc_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The ID of the VPC.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RouteTableTagArgs']]]]:
        """
        Any tags assigned to the route table.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RouteTableTagArgs']]]]):
        pulumi.set(self, "tags", value)


class RouteTable(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouteTableTagArgs']]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::EC2::RouteTable

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouteTableTagArgs']]]] tags: Any tags assigned to the route table.
        :param pulumi.Input[str] vpc_id: The ID of the VPC.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RouteTableArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::EC2::RouteTable

        :param str resource_name: The name of the resource.
        :param RouteTableArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RouteTableArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouteTableTagArgs']]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RouteTableArgs.__new__(RouteTableArgs)

            __props__.__dict__["tags"] = tags
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["route_table_id"] = None
        super(RouteTable, __self__).__init__(
            'aws-native:ec2:RouteTable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RouteTable':
        """
        Get an existing RouteTable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RouteTableArgs.__new__(RouteTableArgs)

        __props__.__dict__["route_table_id"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["vpc_id"] = None
        return RouteTable(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> pulumi.Output[str]:
        """
        The route table ID.
        """
        return pulumi.get(self, "route_table_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.RouteTableTag']]]:
        """
        Any tags assigned to the route table.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The ID of the VPC.
        """
        return pulumi.get(self, "vpc_id")
