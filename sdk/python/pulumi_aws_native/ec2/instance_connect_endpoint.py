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

__all__ = ['InstanceConnectEndpointArgs', 'InstanceConnectEndpoint']

@pulumi.input_type
class InstanceConnectEndpointArgs:
    def __init__(__self__, *,
                 subnet_id: pulumi.Input[str],
                 client_token: Optional[pulumi.Input[str]] = None,
                 preserve_client_ip: Optional[pulumi.Input[bool]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceConnectEndpointTagArgs']]]] = None):
        """
        The set of arguments for constructing a InstanceConnectEndpoint resource.
        :param pulumi.Input[str] subnet_id: The subnet id of the instance connect endpoint
        :param pulumi.Input[str] client_token: The client token of the instance connect endpoint.
        :param pulumi.Input[bool] preserve_client_ip: If true, the address of the loki client is preserved when connecting to the end resource
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: The security group IDs of the instance connect endpoint.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceConnectEndpointTagArgs']]] tags: The tags of the instance connect endpoint.
        """
        pulumi.set(__self__, "subnet_id", subnet_id)
        if client_token is not None:
            pulumi.set(__self__, "client_token", client_token)
        if preserve_client_ip is not None:
            pulumi.set(__self__, "preserve_client_ip", preserve_client_ip)
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        The subnet id of the instance connect endpoint
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="clientToken")
    def client_token(self) -> Optional[pulumi.Input[str]]:
        """
        The client token of the instance connect endpoint.
        """
        return pulumi.get(self, "client_token")

    @client_token.setter
    def client_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_token", value)

    @property
    @pulumi.getter(name="preserveClientIp")
    def preserve_client_ip(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, the address of the loki client is preserved when connecting to the end resource
        """
        return pulumi.get(self, "preserve_client_ip")

    @preserve_client_ip.setter
    def preserve_client_ip(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "preserve_client_ip", value)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The security group IDs of the instance connect endpoint.
        """
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceConnectEndpointTagArgs']]]]:
        """
        The tags of the instance connect endpoint.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceConnectEndpointTagArgs']]]]):
        pulumi.set(self, "tags", value)


class InstanceConnectEndpoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_token: Optional[pulumi.Input[str]] = None,
                 preserve_client_ip: Optional[pulumi.Input[bool]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceConnectEndpointTagArgs']]]]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::EC2::InstanceConnectEndpoint

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_token: The client token of the instance connect endpoint.
        :param pulumi.Input[bool] preserve_client_ip: If true, the address of the loki client is preserved when connecting to the end resource
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: The security group IDs of the instance connect endpoint.
        :param pulumi.Input[str] subnet_id: The subnet id of the instance connect endpoint
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceConnectEndpointTagArgs']]]] tags: The tags of the instance connect endpoint.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceConnectEndpointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::EC2::InstanceConnectEndpoint

        :param str resource_name: The name of the resource.
        :param InstanceConnectEndpointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceConnectEndpointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_token: Optional[pulumi.Input[str]] = None,
                 preserve_client_ip: Optional[pulumi.Input[bool]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceConnectEndpointTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceConnectEndpointArgs.__new__(InstanceConnectEndpointArgs)

            __props__.__dict__["client_token"] = client_token
            __props__.__dict__["preserve_client_ip"] = preserve_client_ip
            __props__.__dict__["security_group_ids"] = security_group_ids
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["tags"] = tags
        super(InstanceConnectEndpoint, __self__).__init__(
            'aws-native:ec2:InstanceConnectEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'InstanceConnectEndpoint':
        """
        Get an existing InstanceConnectEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = InstanceConnectEndpointArgs.__new__(InstanceConnectEndpointArgs)

        __props__.__dict__["client_token"] = None
        __props__.__dict__["preserve_client_ip"] = None
        __props__.__dict__["security_group_ids"] = None
        __props__.__dict__["subnet_id"] = None
        __props__.__dict__["tags"] = None
        return InstanceConnectEndpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientToken")
    def client_token(self) -> pulumi.Output[Optional[str]]:
        """
        The client token of the instance connect endpoint.
        """
        return pulumi.get(self, "client_token")

    @property
    @pulumi.getter(name="preserveClientIp")
    def preserve_client_ip(self) -> pulumi.Output[Optional[bool]]:
        """
        If true, the address of the loki client is preserved when connecting to the end resource
        """
        return pulumi.get(self, "preserve_client_ip")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The security group IDs of the instance connect endpoint.
        """
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        The subnet id of the instance connect endpoint
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.InstanceConnectEndpointTag']]]:
        """
        The tags of the instance connect endpoint.
        """
        return pulumi.get(self, "tags")
