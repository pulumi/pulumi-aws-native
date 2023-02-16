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

__all__ = ['SignalingChannelArgs', 'SignalingChannel']

@pulumi.input_type
class SignalingChannelArgs:
    def __init__(__self__, *,
                 message_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['SignalingChannelTagArgs']]]] = None,
                 type: Optional[pulumi.Input['SignalingChannelType']] = None):
        """
        The set of arguments for constructing a SignalingChannel resource.
        :param pulumi.Input[int] message_ttl_seconds: The period of time a signaling channel retains undelivered messages before they are discarded.
        :param pulumi.Input[str] name: The name of the Kinesis Video Signaling Channel.
        :param pulumi.Input[Sequence[pulumi.Input['SignalingChannelTagArgs']]] tags: An array of key-value pairs to apply to this resource.
        :param pulumi.Input['SignalingChannelType'] type: The type of the Kinesis Video Signaling Channel to create. Currently, SINGLE_MASTER is the only supported channel type.
        """
        if message_ttl_seconds is not None:
            pulumi.set(__self__, "message_ttl_seconds", message_ttl_seconds)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="messageTtlSeconds")
    def message_ttl_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The period of time a signaling channel retains undelivered messages before they are discarded.
        """
        return pulumi.get(self, "message_ttl_seconds")

    @message_ttl_seconds.setter
    def message_ttl_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "message_ttl_seconds", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Kinesis Video Signaling Channel.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SignalingChannelTagArgs']]]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SignalingChannelTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input['SignalingChannelType']]:
        """
        The type of the Kinesis Video Signaling Channel to create. Currently, SINGLE_MASTER is the only supported channel type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input['SignalingChannelType']]):
        pulumi.set(self, "type", value)


class SignalingChannel(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 message_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SignalingChannelTagArgs']]]]] = None,
                 type: Optional[pulumi.Input['SignalingChannelType']] = None,
                 __props__=None):
        """
        Resource Type Definition for AWS::KinesisVideo::SignalingChannel

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] message_ttl_seconds: The period of time a signaling channel retains undelivered messages before they are discarded.
        :param pulumi.Input[str] name: The name of the Kinesis Video Signaling Channel.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SignalingChannelTagArgs']]]] tags: An array of key-value pairs to apply to this resource.
        :param pulumi.Input['SignalingChannelType'] type: The type of the Kinesis Video Signaling Channel to create. Currently, SINGLE_MASTER is the only supported channel type.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SignalingChannelArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type Definition for AWS::KinesisVideo::SignalingChannel

        :param str resource_name: The name of the resource.
        :param SignalingChannelArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SignalingChannelArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 message_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SignalingChannelTagArgs']]]]] = None,
                 type: Optional[pulumi.Input['SignalingChannelType']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SignalingChannelArgs.__new__(SignalingChannelArgs)

            __props__.__dict__["message_ttl_seconds"] = message_ttl_seconds
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["type"] = type
            __props__.__dict__["arn"] = None
        super(SignalingChannel, __self__).__init__(
            'aws-native:kinesisvideo:SignalingChannel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SignalingChannel':
        """
        Get an existing SignalingChannel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SignalingChannelArgs.__new__(SignalingChannelArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["message_ttl_seconds"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        return SignalingChannel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the Kinesis Video Signaling Channel.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="messageTtlSeconds")
    def message_ttl_seconds(self) -> pulumi.Output[Optional[int]]:
        """
        The period of time a signaling channel retains undelivered messages before they are discarded.
        """
        return pulumi.get(self, "message_ttl_seconds")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the Kinesis Video Signaling Channel.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.SignalingChannelTag']]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional['SignalingChannelType']]:
        """
        The type of the Kinesis Video Signaling Channel to create. Currently, SINGLE_MASTER is the only supported channel type.
        """
        return pulumi.get(self, "type")
