# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from .. import outputs as _root_outputs
from ._enums import *

__all__ = [
    'GetPluginResult',
    'AwaitableGetPluginResult',
    'get_plugin',
    'get_plugin_output',
]

@pulumi.output_type
class GetPluginResult:
    def __init__(__self__, auth_configuration=None, build_status=None, created_at=None, custom_plugin_configuration=None, display_name=None, plugin_arn=None, plugin_id=None, server_url=None, state=None, tags=None, updated_at=None):
        if auth_configuration and not isinstance(auth_configuration, dict):
            raise TypeError("Expected argument 'auth_configuration' to be a dict")
        pulumi.set(__self__, "auth_configuration", auth_configuration)
        if build_status and not isinstance(build_status, str):
            raise TypeError("Expected argument 'build_status' to be a str")
        pulumi.set(__self__, "build_status", build_status)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if custom_plugin_configuration and not isinstance(custom_plugin_configuration, dict):
            raise TypeError("Expected argument 'custom_plugin_configuration' to be a dict")
        pulumi.set(__self__, "custom_plugin_configuration", custom_plugin_configuration)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if plugin_arn and not isinstance(plugin_arn, str):
            raise TypeError("Expected argument 'plugin_arn' to be a str")
        pulumi.set(__self__, "plugin_arn", plugin_arn)
        if plugin_id and not isinstance(plugin_id, str):
            raise TypeError("Expected argument 'plugin_id' to be a str")
        pulumi.set(__self__, "plugin_id", plugin_id)
        if server_url and not isinstance(server_url, str):
            raise TypeError("Expected argument 'server_url' to be a str")
        pulumi.set(__self__, "server_url", server_url)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="authConfiguration")
    def auth_configuration(self) -> Optional[Any]:
        return pulumi.get(self, "auth_configuration")

    @property
    @pulumi.getter(name="buildStatus")
    def build_status(self) -> Optional['PluginBuildStatus']:
        return pulumi.get(self, "build_status")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="customPluginConfiguration")
    def custom_plugin_configuration(self) -> Optional['outputs.PluginCustomPluginConfiguration']:
        return pulumi.get(self, "custom_plugin_configuration")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="pluginArn")
    def plugin_arn(self) -> Optional[str]:
        return pulumi.get(self, "plugin_arn")

    @property
    @pulumi.getter(name="pluginId")
    def plugin_id(self) -> Optional[str]:
        return pulumi.get(self, "plugin_id")

    @property
    @pulumi.getter(name="serverUrl")
    def server_url(self) -> Optional[str]:
        return pulumi.get(self, "server_url")

    @property
    @pulumi.getter
    def state(self) -> Optional['PluginState']:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[str]:
        return pulumi.get(self, "updated_at")


class AwaitableGetPluginResult(GetPluginResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPluginResult(
            auth_configuration=self.auth_configuration,
            build_status=self.build_status,
            created_at=self.created_at,
            custom_plugin_configuration=self.custom_plugin_configuration,
            display_name=self.display_name,
            plugin_arn=self.plugin_arn,
            plugin_id=self.plugin_id,
            server_url=self.server_url,
            state=self.state,
            tags=self.tags,
            updated_at=self.updated_at)


def get_plugin(application_id: Optional[str] = None,
               plugin_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPluginResult:
    """
    Definition of AWS::QBusiness::Plugin Resource Type
    """
    __args__ = dict()
    __args__['applicationId'] = application_id
    __args__['pluginId'] = plugin_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:qbusiness:getPlugin', __args__, opts=opts, typ=GetPluginResult).value

    return AwaitableGetPluginResult(
        auth_configuration=pulumi.get(__ret__, 'auth_configuration'),
        build_status=pulumi.get(__ret__, 'build_status'),
        created_at=pulumi.get(__ret__, 'created_at'),
        custom_plugin_configuration=pulumi.get(__ret__, 'custom_plugin_configuration'),
        display_name=pulumi.get(__ret__, 'display_name'),
        plugin_arn=pulumi.get(__ret__, 'plugin_arn'),
        plugin_id=pulumi.get(__ret__, 'plugin_id'),
        server_url=pulumi.get(__ret__, 'server_url'),
        state=pulumi.get(__ret__, 'state'),
        tags=pulumi.get(__ret__, 'tags'),
        updated_at=pulumi.get(__ret__, 'updated_at'))


@_utilities.lift_output_func(get_plugin)
def get_plugin_output(application_id: Optional[pulumi.Input[str]] = None,
                      plugin_id: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPluginResult]:
    """
    Definition of AWS::QBusiness::Plugin Resource Type
    """
    ...