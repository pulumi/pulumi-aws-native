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

__all__ = ['BillingGroupArgs', 'BillingGroup']

@pulumi.input_type
class BillingGroupArgs:
    def __init__(__self__, *,
                 billing_group_name: Optional[pulumi.Input[str]] = None,
                 billing_group_properties: Optional[pulumi.Input['BillingGroupPropertiesPropertiesArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['BillingGroupTagArgs']]]] = None):
        """
        The set of arguments for constructing a BillingGroup resource.
        :param pulumi.Input[Sequence[pulumi.Input['BillingGroupTagArgs']]] tags: An array of key-value pairs to apply to this resource.
        """
        if billing_group_name is not None:
            pulumi.set(__self__, "billing_group_name", billing_group_name)
        if billing_group_properties is not None:
            pulumi.set(__self__, "billing_group_properties", billing_group_properties)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="billingGroupName")
    def billing_group_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "billing_group_name")

    @billing_group_name.setter
    def billing_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "billing_group_name", value)

    @property
    @pulumi.getter(name="billingGroupProperties")
    def billing_group_properties(self) -> Optional[pulumi.Input['BillingGroupPropertiesPropertiesArgs']]:
        return pulumi.get(self, "billing_group_properties")

    @billing_group_properties.setter
    def billing_group_properties(self, value: Optional[pulumi.Input['BillingGroupPropertiesPropertiesArgs']]):
        pulumi.set(self, "billing_group_properties", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BillingGroupTagArgs']]]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BillingGroupTagArgs']]]]):
        pulumi.set(self, "tags", value)


class BillingGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 billing_group_name: Optional[pulumi.Input[str]] = None,
                 billing_group_properties: Optional[pulumi.Input[pulumi.InputType['BillingGroupPropertiesPropertiesArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BillingGroupTagArgs']]]]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::IoT::BillingGroup

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BillingGroupTagArgs']]]] tags: An array of key-value pairs to apply to this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[BillingGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::IoT::BillingGroup

        :param str resource_name: The name of the resource.
        :param BillingGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BillingGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 billing_group_name: Optional[pulumi.Input[str]] = None,
                 billing_group_properties: Optional[pulumi.Input[pulumi.InputType['BillingGroupPropertiesPropertiesArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BillingGroupTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BillingGroupArgs.__new__(BillingGroupArgs)

            __props__.__dict__["billing_group_name"] = billing_group_name
            __props__.__dict__["billing_group_properties"] = billing_group_properties
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
        super(BillingGroup, __self__).__init__(
            'aws-native:iot:BillingGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'BillingGroup':
        """
        Get an existing BillingGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = BillingGroupArgs.__new__(BillingGroupArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["billing_group_name"] = None
        __props__.__dict__["billing_group_properties"] = None
        __props__.__dict__["tags"] = None
        return BillingGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="billingGroupName")
    def billing_group_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "billing_group_name")

    @property
    @pulumi.getter(name="billingGroupProperties")
    def billing_group_properties(self) -> pulumi.Output[Optional['outputs.BillingGroupPropertiesProperties']]:
        return pulumi.get(self, "billing_group_properties")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.BillingGroupTag']]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")
