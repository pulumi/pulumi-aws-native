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

__all__ = ['RouteServerAssociationArgs', 'RouteServerAssociation']

@pulumi.input_type
class RouteServerAssociationArgs:
    def __init__(__self__, *,
                 route_server_id: pulumi.Input[builtins.str],
                 vpc_id: pulumi.Input[builtins.str]):
        """
        The set of arguments for constructing a RouteServerAssociation resource.
        :param pulumi.Input[builtins.str] route_server_id: Route Server ID
        :param pulumi.Input[builtins.str] vpc_id: VPC ID
        """
        pulumi.set(__self__, "route_server_id", route_server_id)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="routeServerId")
    def route_server_id(self) -> pulumi.Input[builtins.str]:
        """
        Route Server ID
        """
        return pulumi.get(self, "route_server_id")

    @route_server_id.setter
    def route_server_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "route_server_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[builtins.str]:
        """
        VPC ID
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "vpc_id", value)


@pulumi.type_token("aws-native:ec2:RouteServerAssociation")
class RouteServerAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 route_server_id: Optional[pulumi.Input[builtins.str]] = None,
                 vpc_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        VPC Route Server Association

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] route_server_id: Route Server ID
        :param pulumi.Input[builtins.str] vpc_id: VPC ID
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RouteServerAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        VPC Route Server Association

        :param str resource_name: The name of the resource.
        :param RouteServerAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RouteServerAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 route_server_id: Optional[pulumi.Input[builtins.str]] = None,
                 vpc_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RouteServerAssociationArgs.__new__(RouteServerAssociationArgs)

            if route_server_id is None and not opts.urn:
                raise TypeError("Missing required property 'route_server_id'")
            __props__.__dict__["route_server_id"] = route_server_id
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["routeServerId", "vpcId"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(RouteServerAssociation, __self__).__init__(
            'aws-native:ec2:RouteServerAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RouteServerAssociation':
        """
        Get an existing RouteServerAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RouteServerAssociationArgs.__new__(RouteServerAssociationArgs)

        __props__.__dict__["route_server_id"] = None
        __props__.__dict__["vpc_id"] = None
        return RouteServerAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="routeServerId")
    def route_server_id(self) -> pulumi.Output[builtins.str]:
        """
        Route Server ID
        """
        return pulumi.get(self, "route_server_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[builtins.str]:
        """
        VPC ID
        """
        return pulumi.get(self, "vpc_id")

