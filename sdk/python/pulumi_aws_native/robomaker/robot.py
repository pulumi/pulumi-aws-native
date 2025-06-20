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
from ._enums import *

__all__ = ['RobotArgs', 'Robot']

@pulumi.input_type
class RobotArgs:
    def __init__(__self__, *,
                 architecture: pulumi.Input['RobotArchitecture'],
                 greengrass_group_id: pulumi.Input[builtins.str],
                 fleet: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a Robot resource.
        :param pulumi.Input['RobotArchitecture'] architecture: The target architecture of the robot.
        :param pulumi.Input[builtins.str] greengrass_group_id: The Greengrass group id.
        :param pulumi.Input[builtins.str] fleet: The Amazon Resource Name (ARN) of the fleet.
        :param pulumi.Input[builtins.str] name: The name for the robot.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: A map that contains tag keys and tag values that are attached to the robot.
        """
        pulumi.set(__self__, "architecture", architecture)
        pulumi.set(__self__, "greengrass_group_id", greengrass_group_id)
        if fleet is not None:
            pulumi.set(__self__, "fleet", fleet)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def architecture(self) -> pulumi.Input['RobotArchitecture']:
        """
        The target architecture of the robot.
        """
        return pulumi.get(self, "architecture")

    @architecture.setter
    def architecture(self, value: pulumi.Input['RobotArchitecture']):
        pulumi.set(self, "architecture", value)

    @property
    @pulumi.getter(name="greengrassGroupId")
    def greengrass_group_id(self) -> pulumi.Input[builtins.str]:
        """
        The Greengrass group id.
        """
        return pulumi.get(self, "greengrass_group_id")

    @greengrass_group_id.setter
    def greengrass_group_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "greengrass_group_id", value)

    @property
    @pulumi.getter
    def fleet(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Amazon Resource Name (ARN) of the fleet.
        """
        return pulumi.get(self, "fleet")

    @fleet.setter
    def fleet(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "fleet", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name for the robot.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        A map that contains tag keys and tag values that are attached to the robot.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:robomaker:Robot")
class Robot(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 architecture: Optional[pulumi.Input['RobotArchitecture']] = None,
                 fleet: Optional[pulumi.Input[builtins.str]] = None,
                 greengrass_group_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        AWS::RoboMaker::Robot resource creates an AWS RoboMaker Robot.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['RobotArchitecture'] architecture: The target architecture of the robot.
        :param pulumi.Input[builtins.str] fleet: The Amazon Resource Name (ARN) of the fleet.
        :param pulumi.Input[builtins.str] greengrass_group_id: The Greengrass group id.
        :param pulumi.Input[builtins.str] name: The name for the robot.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: A map that contains tag keys and tag values that are attached to the robot.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RobotArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        AWS::RoboMaker::Robot resource creates an AWS RoboMaker Robot.

        :param str resource_name: The name of the resource.
        :param RobotArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RobotArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 architecture: Optional[pulumi.Input['RobotArchitecture']] = None,
                 fleet: Optional[pulumi.Input[builtins.str]] = None,
                 greengrass_group_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RobotArgs.__new__(RobotArgs)

            if architecture is None and not opts.urn:
                raise TypeError("Missing required property 'architecture'")
            __props__.__dict__["architecture"] = architecture
            __props__.__dict__["fleet"] = fleet
            if greengrass_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'greengrass_group_id'")
            __props__.__dict__["greengrass_group_id"] = greengrass_group_id
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["architecture", "fleet", "greengrassGroupId", "name"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Robot, __self__).__init__(
            'aws-native:robomaker:Robot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Robot':
        """
        Get an existing Robot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RobotArgs.__new__(RobotArgs)

        __props__.__dict__["architecture"] = None
        __props__.__dict__["arn"] = None
        __props__.__dict__["fleet"] = None
        __props__.__dict__["greengrass_group_id"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["tags"] = None
        return Robot(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def architecture(self) -> pulumi.Output['RobotArchitecture']:
        """
        The target architecture of the robot.
        """
        return pulumi.get(self, "architecture")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the robot.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def fleet(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The Amazon Resource Name (ARN) of the fleet.
        """
        return pulumi.get(self, "fleet")

    @property
    @pulumi.getter(name="greengrassGroupId")
    def greengrass_group_id(self) -> pulumi.Output[builtins.str]:
        """
        The Greengrass group id.
        """
        return pulumi.get(self, "greengrass_group_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The name for the robot.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        A map that contains tag keys and tag values that are attached to the robot.
        """
        return pulumi.get(self, "tags")

