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
from ._enums import *
from ._inputs import *

__all__ = ['SpotFleetArgs', 'SpotFleet']

@pulumi.input_type
class SpotFleetArgs:
    def __init__(__self__, *,
                 spot_fleet_request_config_data: pulumi.Input['SpotFleetRequestConfigDataArgs']):
        """
        The set of arguments for constructing a SpotFleet resource.
        """
        pulumi.set(__self__, "spot_fleet_request_config_data", spot_fleet_request_config_data)

    @property
    @pulumi.getter(name="spotFleetRequestConfigData")
    def spot_fleet_request_config_data(self) -> pulumi.Input['SpotFleetRequestConfigDataArgs']:
        return pulumi.get(self, "spot_fleet_request_config_data")

    @spot_fleet_request_config_data.setter
    def spot_fleet_request_config_data(self, value: pulumi.Input['SpotFleetRequestConfigDataArgs']):
        pulumi.set(self, "spot_fleet_request_config_data", value)


class SpotFleet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 spot_fleet_request_config_data: Optional[pulumi.Input[pulumi.InputType['SpotFleetRequestConfigDataArgs']]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::EC2::SpotFleet

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SpotFleetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::EC2::SpotFleet

        :param str resource_name: The name of the resource.
        :param SpotFleetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SpotFleetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 spot_fleet_request_config_data: Optional[pulumi.Input[pulumi.InputType['SpotFleetRequestConfigDataArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SpotFleetArgs.__new__(SpotFleetArgs)

            if spot_fleet_request_config_data is None and not opts.urn:
                raise TypeError("Missing required property 'spot_fleet_request_config_data'")
            __props__.__dict__["spot_fleet_request_config_data"] = spot_fleet_request_config_data
        super(SpotFleet, __self__).__init__(
            'aws-native:ec2:SpotFleet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SpotFleet':
        """
        Get an existing SpotFleet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SpotFleetArgs.__new__(SpotFleetArgs)

        __props__.__dict__["spot_fleet_request_config_data"] = None
        return SpotFleet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="spotFleetRequestConfigData")
    def spot_fleet_request_config_data(self) -> pulumi.Output['outputs.SpotFleetRequestConfigData']:
        return pulumi.get(self, "spot_fleet_request_config_data")
