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
from ._enums import *
from ._inputs import *

__all__ = ['ChannelNamespaceArgs', 'ChannelNamespace']

@pulumi.input_type
class ChannelNamespaceArgs:
    def __init__(__self__, *,
                 api_id: pulumi.Input[builtins.str],
                 code_handlers: Optional[pulumi.Input[builtins.str]] = None,
                 code_s3_location: Optional[pulumi.Input[builtins.str]] = None,
                 handler_configs: Optional[pulumi.Input['ChannelNamespaceHandlerConfigsArgs']] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 publish_auth_modes: Optional[pulumi.Input[Sequence[pulumi.Input['ChannelNamespaceAuthModeArgs']]]] = None,
                 subscribe_auth_modes: Optional[pulumi.Input[Sequence[pulumi.Input['ChannelNamespaceAuthModeArgs']]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a ChannelNamespace resource.
        :param pulumi.Input[builtins.str] api_id: AppSync Api Id that this Channel Namespace belongs to.
        :param pulumi.Input[builtins.str] code_handlers: The event handler functions that run custom business logic to process published events and subscribe requests.
        :param pulumi.Input[builtins.str] code_s3_location: The Amazon S3 endpoint where the code is located.
        :param pulumi.Input['ChannelNamespaceHandlerConfigsArgs'] handler_configs: The configuration for the `OnPublish` and `OnSubscribe` handlers.
        :param pulumi.Input[builtins.str] name: The name of the channel namespace. This name must be unique within the `Api` .
        :param pulumi.Input[Sequence[pulumi.Input['ChannelNamespaceAuthModeArgs']]] publish_auth_modes: List of AuthModes supported for Publish operations.
        :param pulumi.Input[Sequence[pulumi.Input['ChannelNamespaceAuthModeArgs']]] subscribe_auth_modes: List of AuthModes supported for Subscribe operations.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: A set of tags (key-value pairs) for this channel namespace.
        """
        pulumi.set(__self__, "api_id", api_id)
        if code_handlers is not None:
            pulumi.set(__self__, "code_handlers", code_handlers)
        if code_s3_location is not None:
            pulumi.set(__self__, "code_s3_location", code_s3_location)
        if handler_configs is not None:
            pulumi.set(__self__, "handler_configs", handler_configs)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if publish_auth_modes is not None:
            pulumi.set(__self__, "publish_auth_modes", publish_auth_modes)
        if subscribe_auth_modes is not None:
            pulumi.set(__self__, "subscribe_auth_modes", subscribe_auth_modes)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Input[builtins.str]:
        """
        AppSync Api Id that this Channel Namespace belongs to.
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter(name="codeHandlers")
    def code_handlers(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The event handler functions that run custom business logic to process published events and subscribe requests.
        """
        return pulumi.get(self, "code_handlers")

    @code_handlers.setter
    def code_handlers(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "code_handlers", value)

    @property
    @pulumi.getter(name="codeS3Location")
    def code_s3_location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Amazon S3 endpoint where the code is located.
        """
        return pulumi.get(self, "code_s3_location")

    @code_s3_location.setter
    def code_s3_location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "code_s3_location", value)

    @property
    @pulumi.getter(name="handlerConfigs")
    def handler_configs(self) -> Optional[pulumi.Input['ChannelNamespaceHandlerConfigsArgs']]:
        """
        The configuration for the `OnPublish` and `OnSubscribe` handlers.
        """
        return pulumi.get(self, "handler_configs")

    @handler_configs.setter
    def handler_configs(self, value: Optional[pulumi.Input['ChannelNamespaceHandlerConfigsArgs']]):
        pulumi.set(self, "handler_configs", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the channel namespace. This name must be unique within the `Api` .
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="publishAuthModes")
    def publish_auth_modes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ChannelNamespaceAuthModeArgs']]]]:
        """
        List of AuthModes supported for Publish operations.
        """
        return pulumi.get(self, "publish_auth_modes")

    @publish_auth_modes.setter
    def publish_auth_modes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ChannelNamespaceAuthModeArgs']]]]):
        pulumi.set(self, "publish_auth_modes", value)

    @property
    @pulumi.getter(name="subscribeAuthModes")
    def subscribe_auth_modes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ChannelNamespaceAuthModeArgs']]]]:
        """
        List of AuthModes supported for Subscribe operations.
        """
        return pulumi.get(self, "subscribe_auth_modes")

    @subscribe_auth_modes.setter
    def subscribe_auth_modes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ChannelNamespaceAuthModeArgs']]]]):
        pulumi.set(self, "subscribe_auth_modes", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        A set of tags (key-value pairs) for this channel namespace.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:appsync:ChannelNamespace")
class ChannelNamespace(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[builtins.str]] = None,
                 code_handlers: Optional[pulumi.Input[builtins.str]] = None,
                 code_s3_location: Optional[pulumi.Input[builtins.str]] = None,
                 handler_configs: Optional[pulumi.Input[Union['ChannelNamespaceHandlerConfigsArgs', 'ChannelNamespaceHandlerConfigsArgsDict']]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 publish_auth_modes: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ChannelNamespaceAuthModeArgs', 'ChannelNamespaceAuthModeArgsDict']]]]] = None,
                 subscribe_auth_modes: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ChannelNamespaceAuthModeArgs', 'ChannelNamespaceAuthModeArgsDict']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        """
        Resource schema for AppSync ChannelNamespace

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] api_id: AppSync Api Id that this Channel Namespace belongs to.
        :param pulumi.Input[builtins.str] code_handlers: The event handler functions that run custom business logic to process published events and subscribe requests.
        :param pulumi.Input[builtins.str] code_s3_location: The Amazon S3 endpoint where the code is located.
        :param pulumi.Input[Union['ChannelNamespaceHandlerConfigsArgs', 'ChannelNamespaceHandlerConfigsArgsDict']] handler_configs: The configuration for the `OnPublish` and `OnSubscribe` handlers.
        :param pulumi.Input[builtins.str] name: The name of the channel namespace. This name must be unique within the `Api` .
        :param pulumi.Input[Sequence[pulumi.Input[Union['ChannelNamespaceAuthModeArgs', 'ChannelNamespaceAuthModeArgsDict']]]] publish_auth_modes: List of AuthModes supported for Publish operations.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ChannelNamespaceAuthModeArgs', 'ChannelNamespaceAuthModeArgsDict']]]] subscribe_auth_modes: List of AuthModes supported for Subscribe operations.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] tags: A set of tags (key-value pairs) for this channel namespace.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ChannelNamespaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource schema for AppSync ChannelNamespace

        :param str resource_name: The name of the resource.
        :param ChannelNamespaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ChannelNamespaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[builtins.str]] = None,
                 code_handlers: Optional[pulumi.Input[builtins.str]] = None,
                 code_s3_location: Optional[pulumi.Input[builtins.str]] = None,
                 handler_configs: Optional[pulumi.Input[Union['ChannelNamespaceHandlerConfigsArgs', 'ChannelNamespaceHandlerConfigsArgsDict']]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 publish_auth_modes: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ChannelNamespaceAuthModeArgs', 'ChannelNamespaceAuthModeArgsDict']]]]] = None,
                 subscribe_auth_modes: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ChannelNamespaceAuthModeArgs', 'ChannelNamespaceAuthModeArgsDict']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ChannelNamespaceArgs.__new__(ChannelNamespaceArgs)

            if api_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_id'")
            __props__.__dict__["api_id"] = api_id
            __props__.__dict__["code_handlers"] = code_handlers
            __props__.__dict__["code_s3_location"] = code_s3_location
            __props__.__dict__["handler_configs"] = handler_configs
            __props__.__dict__["name"] = name
            __props__.__dict__["publish_auth_modes"] = publish_auth_modes
            __props__.__dict__["subscribe_auth_modes"] = subscribe_auth_modes
            __props__.__dict__["tags"] = tags
            __props__.__dict__["channel_namespace_arn"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["apiId", "name"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(ChannelNamespace, __self__).__init__(
            'aws-native:appsync:ChannelNamespace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ChannelNamespace':
        """
        Get an existing ChannelNamespace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ChannelNamespaceArgs.__new__(ChannelNamespaceArgs)

        __props__.__dict__["api_id"] = None
        __props__.__dict__["channel_namespace_arn"] = None
        __props__.__dict__["code_handlers"] = None
        __props__.__dict__["code_s3_location"] = None
        __props__.__dict__["handler_configs"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["publish_auth_modes"] = None
        __props__.__dict__["subscribe_auth_modes"] = None
        __props__.__dict__["tags"] = None
        return ChannelNamespace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Output[builtins.str]:
        """
        AppSync Api Id that this Channel Namespace belongs to.
        """
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter(name="channelNamespaceArn")
    def channel_namespace_arn(self) -> pulumi.Output[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the channel namespace.
        """
        return pulumi.get(self, "channel_namespace_arn")

    @property
    @pulumi.getter(name="codeHandlers")
    def code_handlers(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The event handler functions that run custom business logic to process published events and subscribe requests.
        """
        return pulumi.get(self, "code_handlers")

    @property
    @pulumi.getter(name="codeS3Location")
    def code_s3_location(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The Amazon S3 endpoint where the code is located.
        """
        return pulumi.get(self, "code_s3_location")

    @property
    @pulumi.getter(name="handlerConfigs")
    def handler_configs(self) -> pulumi.Output[Optional['outputs.ChannelNamespaceHandlerConfigs']]:
        """
        The configuration for the `OnPublish` and `OnSubscribe` handlers.
        """
        return pulumi.get(self, "handler_configs")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the channel namespace. This name must be unique within the `Api` .
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="publishAuthModes")
    def publish_auth_modes(self) -> pulumi.Output[Optional[Sequence['outputs.ChannelNamespaceAuthMode']]]:
        """
        List of AuthModes supported for Publish operations.
        """
        return pulumi.get(self, "publish_auth_modes")

    @property
    @pulumi.getter(name="subscribeAuthModes")
    def subscribe_auth_modes(self) -> pulumi.Output[Optional[Sequence['outputs.ChannelNamespaceAuthMode']]]:
        """
        List of AuthModes supported for Subscribe operations.
        """
        return pulumi.get(self, "subscribe_auth_modes")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        A set of tags (key-value pairs) for this channel namespace.
        """
        return pulumi.get(self, "tags")

