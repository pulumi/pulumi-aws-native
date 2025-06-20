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
from . import outputs
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._inputs import *

__all__ = ['EncoderConfigurationArgs', 'EncoderConfiguration']

@pulumi.input_type
class EncoderConfigurationArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None,
                 video: Optional[pulumi.Input['VideoPropertiesArgs']] = None):
        """
        The set of arguments for constructing a EncoderConfiguration resource.
        :param pulumi.Input[builtins.str] name: Encoder configuration name.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: An array of key-value pairs to apply to this resource.
        :param pulumi.Input['VideoPropertiesArgs'] video: Video configuration. Default: video resolution 1280x720, bitrate 2500 kbps, 30 fps
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if video is not None:
            pulumi.set(__self__, "video", video)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Encoder configuration name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def video(self) -> Optional[pulumi.Input['VideoPropertiesArgs']]:
        """
        Video configuration. Default: video resolution 1280x720, bitrate 2500 kbps, 30 fps
        """
        return pulumi.get(self, "video")

    @video.setter
    def video(self, value: Optional[pulumi.Input['VideoPropertiesArgs']]):
        pulumi.set(self, "video", value)


@pulumi.type_token("aws-native:ivs:EncoderConfiguration")
class EncoderConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 video: Optional[pulumi.Input[Union['VideoPropertiesArgs', 'VideoPropertiesArgsDict']]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::IVS::EncoderConfiguration.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] name: Encoder configuration name.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] tags: An array of key-value pairs to apply to this resource.
        :param pulumi.Input[Union['VideoPropertiesArgs', 'VideoPropertiesArgsDict']] video: Video configuration. Default: video resolution 1280x720, bitrate 2500 kbps, 30 fps
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[EncoderConfigurationArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::IVS::EncoderConfiguration.

        :param str resource_name: The name of the resource.
        :param EncoderConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EncoderConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 video: Optional[pulumi.Input[Union['VideoPropertiesArgs', 'VideoPropertiesArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EncoderConfigurationArgs.__new__(EncoderConfigurationArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["video"] = video
            __props__.__dict__["arn"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["name", "video"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(EncoderConfiguration, __self__).__init__(
            'aws-native:ivs:EncoderConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'EncoderConfiguration':
        """
        Get an existing EncoderConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EncoderConfigurationArgs.__new__(EncoderConfigurationArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["video"] = None
        return EncoderConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[builtins.str]:
        """
        Encoder configuration identifier.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Encoder configuration name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def video(self) -> pulumi.Output[Optional['outputs.VideoProperties']]:
        """
        Video configuration. Default: video resolution 1280x720, bitrate 2500 kbps, 30 fps
        """
        return pulumi.get(self, "video")

