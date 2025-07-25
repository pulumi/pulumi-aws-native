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
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs

__all__ = ['StreamKeyArgs', 'StreamKey']

@pulumi.input_type
class StreamKeyArgs:
    def __init__(__self__, *,
                 channel_arn: pulumi.Input[builtins.str],
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a StreamKey resource.
        :param pulumi.Input[builtins.str] channel_arn: Channel ARN for the stream.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: A list of key-value pairs that contain metadata for the asset model.
        """
        pulumi.set(__self__, "channel_arn", channel_arn)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="channelArn")
    def channel_arn(self) -> pulumi.Input[builtins.str]:
        """
        Channel ARN for the stream.
        """
        return pulumi.get(self, "channel_arn")

    @channel_arn.setter
    def channel_arn(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "channel_arn", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        A list of key-value pairs that contain metadata for the asset model.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:ivs:StreamKey")
class StreamKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 channel_arn: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::IVS::StreamKey

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] channel_arn: Channel ARN for the stream.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] tags: A list of key-value pairs that contain metadata for the asset model.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StreamKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::IVS::StreamKey

        :param str resource_name: The name of the resource.
        :param StreamKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StreamKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 channel_arn: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StreamKeyArgs.__new__(StreamKeyArgs)

            if channel_arn is None and not opts.urn:
                raise TypeError("Missing required property 'channel_arn'")
            __props__.__dict__["channel_arn"] = channel_arn
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["value"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["channelArn"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(StreamKey, __self__).__init__(
            'aws-native:ivs:StreamKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'StreamKey':
        """
        Get an existing StreamKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = StreamKeyArgs.__new__(StreamKeyArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["channel_arn"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["value"] = None
        return StreamKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[builtins.str]:
        """
        Stream Key ARN is automatically generated on creation and assigned as the unique identifier.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="channelArn")
    def channel_arn(self) -> pulumi.Output[builtins.str]:
        """
        Channel ARN for the stream.
        """
        return pulumi.get(self, "channel_arn")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        A list of key-value pairs that contain metadata for the asset model.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[builtins.str]:
        """
        Stream-key value.
        """
        return pulumi.get(self, "value")

