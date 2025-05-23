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
from .. import outputs as _root_outputs
from ._enums import *

__all__ = [
    'GetChannelNamespaceResult',
    'AwaitableGetChannelNamespaceResult',
    'get_channel_namespace',
    'get_channel_namespace_output',
]

@pulumi.output_type
class GetChannelNamespaceResult:
    def __init__(__self__, channel_namespace_arn=None, code_handlers=None, handler_configs=None, publish_auth_modes=None, subscribe_auth_modes=None, tags=None):
        if channel_namespace_arn and not isinstance(channel_namespace_arn, str):
            raise TypeError("Expected argument 'channel_namespace_arn' to be a str")
        pulumi.set(__self__, "channel_namespace_arn", channel_namespace_arn)
        if code_handlers and not isinstance(code_handlers, str):
            raise TypeError("Expected argument 'code_handlers' to be a str")
        pulumi.set(__self__, "code_handlers", code_handlers)
        if handler_configs and not isinstance(handler_configs, dict):
            raise TypeError("Expected argument 'handler_configs' to be a dict")
        pulumi.set(__self__, "handler_configs", handler_configs)
        if publish_auth_modes and not isinstance(publish_auth_modes, list):
            raise TypeError("Expected argument 'publish_auth_modes' to be a list")
        pulumi.set(__self__, "publish_auth_modes", publish_auth_modes)
        if subscribe_auth_modes and not isinstance(subscribe_auth_modes, list):
            raise TypeError("Expected argument 'subscribe_auth_modes' to be a list")
        pulumi.set(__self__, "subscribe_auth_modes", subscribe_auth_modes)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="channelNamespaceArn")
    def channel_namespace_arn(self) -> Optional[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the channel namespace.
        """
        return pulumi.get(self, "channel_namespace_arn")

    @property
    @pulumi.getter(name="codeHandlers")
    def code_handlers(self) -> Optional[builtins.str]:
        """
        The event handler functions that run custom business logic to process published events and subscribe requests.
        """
        return pulumi.get(self, "code_handlers")

    @property
    @pulumi.getter(name="handlerConfigs")
    def handler_configs(self) -> Optional['outputs.ChannelNamespaceHandlerConfigs']:
        """
        The configuration for the `OnPublish` and `OnSubscribe` handlers.
        """
        return pulumi.get(self, "handler_configs")

    @property
    @pulumi.getter(name="publishAuthModes")
    def publish_auth_modes(self) -> Optional[Sequence['outputs.ChannelNamespaceAuthMode']]:
        """
        List of AuthModes supported for Publish operations.
        """
        return pulumi.get(self, "publish_auth_modes")

    @property
    @pulumi.getter(name="subscribeAuthModes")
    def subscribe_auth_modes(self) -> Optional[Sequence['outputs.ChannelNamespaceAuthMode']]:
        """
        List of AuthModes supported for Subscribe operations.
        """
        return pulumi.get(self, "subscribe_auth_modes")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        A set of tags (key-value pairs) for this channel namespace.
        """
        return pulumi.get(self, "tags")


class AwaitableGetChannelNamespaceResult(GetChannelNamespaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetChannelNamespaceResult(
            channel_namespace_arn=self.channel_namespace_arn,
            code_handlers=self.code_handlers,
            handler_configs=self.handler_configs,
            publish_auth_modes=self.publish_auth_modes,
            subscribe_auth_modes=self.subscribe_auth_modes,
            tags=self.tags)


def get_channel_namespace(channel_namespace_arn: Optional[builtins.str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetChannelNamespaceResult:
    """
    Resource schema for AppSync ChannelNamespace


    :param builtins.str channel_namespace_arn: The Amazon Resource Name (ARN) of the channel namespace.
    """
    __args__ = dict()
    __args__['channelNamespaceArn'] = channel_namespace_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:appsync:getChannelNamespace', __args__, opts=opts, typ=GetChannelNamespaceResult).value

    return AwaitableGetChannelNamespaceResult(
        channel_namespace_arn=pulumi.get(__ret__, 'channel_namespace_arn'),
        code_handlers=pulumi.get(__ret__, 'code_handlers'),
        handler_configs=pulumi.get(__ret__, 'handler_configs'),
        publish_auth_modes=pulumi.get(__ret__, 'publish_auth_modes'),
        subscribe_auth_modes=pulumi.get(__ret__, 'subscribe_auth_modes'),
        tags=pulumi.get(__ret__, 'tags'))
def get_channel_namespace_output(channel_namespace_arn: Optional[pulumi.Input[builtins.str]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetChannelNamespaceResult]:
    """
    Resource schema for AppSync ChannelNamespace


    :param builtins.str channel_namespace_arn: The Amazon Resource Name (ARN) of the channel namespace.
    """
    __args__ = dict()
    __args__['channelNamespaceArn'] = channel_namespace_arn
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:appsync:getChannelNamespace', __args__, opts=opts, typ=GetChannelNamespaceResult)
    return __ret__.apply(lambda __response__: GetChannelNamespaceResult(
        channel_namespace_arn=pulumi.get(__response__, 'channel_namespace_arn'),
        code_handlers=pulumi.get(__response__, 'code_handlers'),
        handler_configs=pulumi.get(__response__, 'handler_configs'),
        publish_auth_modes=pulumi.get(__response__, 'publish_auth_modes'),
        subscribe_auth_modes=pulumi.get(__response__, 'subscribe_auth_modes'),
        tags=pulumi.get(__response__, 'tags')))
