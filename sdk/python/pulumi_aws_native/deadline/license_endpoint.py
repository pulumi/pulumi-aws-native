# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = ['LicenseEndpointArgs', 'LicenseEndpoint']

@pulumi.input_type
class LicenseEndpointArgs:
    def __init__(__self__, *,
                 security_group_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 subnet_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 vpc_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a LicenseEndpoint resource.
        """
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "subnet_ids")

    @subnet_ids.setter
    def subnet_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "subnet_ids", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)


class LicenseEndpoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Definition of AWS::Deadline::LicenseEndpoint Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LicenseEndpointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::Deadline::LicenseEndpoint Resource Type

        :param str resource_name: The name of the resource.
        :param LicenseEndpointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LicenseEndpointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LicenseEndpointArgs.__new__(LicenseEndpointArgs)

            if security_group_ids is None and not opts.urn:
                raise TypeError("Missing required property 'security_group_ids'")
            __props__.__dict__["security_group_ids"] = security_group_ids
            if subnet_ids is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_ids'")
            __props__.__dict__["subnet_ids"] = subnet_ids
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["arn"] = None
            __props__.__dict__["dns_name"] = None
            __props__.__dict__["license_endpoint_id"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["status_message"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["securityGroupIds[*]", "subnetIds[*]", "vpcId"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(LicenseEndpoint, __self__).__init__(
            'aws-native:deadline:LicenseEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'LicenseEndpoint':
        """
        Get an existing LicenseEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = LicenseEndpointArgs.__new__(LicenseEndpointArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["dns_name"] = None
        __props__.__dict__["license_endpoint_id"] = None
        __props__.__dict__["security_group_ids"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["status_message"] = None
        __props__.__dict__["subnet_ids"] = None
        __props__.__dict__["vpc_id"] = None
        return LicenseEndpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "dns_name")

    @property
    @pulumi.getter(name="licenseEndpointId")
    def license_endpoint_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "license_endpoint_id")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['LicenseEndpointStatus']:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusMessage")
    def status_message(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status_message")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "vpc_id")
