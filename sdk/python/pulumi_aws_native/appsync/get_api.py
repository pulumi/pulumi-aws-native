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
    'GetApiResult',
    'AwaitableGetApiResult',
    'get_api',
    'get_api_output',
]

@pulumi.output_type
class GetApiResult:
    def __init__(__self__, api_arn=None, api_id=None, dns=None, event_config=None, name=None, owner_contact=None, tags=None):
        if api_arn and not isinstance(api_arn, str):
            raise TypeError("Expected argument 'api_arn' to be a str")
        pulumi.set(__self__, "api_arn", api_arn)
        if api_id and not isinstance(api_id, str):
            raise TypeError("Expected argument 'api_id' to be a str")
        pulumi.set(__self__, "api_id", api_id)
        if dns and not isinstance(dns, dict):
            raise TypeError("Expected argument 'dns' to be a dict")
        pulumi.set(__self__, "dns", dns)
        if event_config and not isinstance(event_config, dict):
            raise TypeError("Expected argument 'event_config' to be a dict")
        pulumi.set(__self__, "event_config", event_config)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if owner_contact and not isinstance(owner_contact, str):
            raise TypeError("Expected argument 'owner_contact' to be a str")
        pulumi.set(__self__, "owner_contact", owner_contact)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="apiArn")
    def api_arn(self) -> Optional[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the AppSync Api
        """
        return pulumi.get(self, "api_arn")

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> Optional[builtins.str]:
        """
        The unique identifier for the AppSync Api generated by the service
        """
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter
    def dns(self) -> Optional['outputs.ApiDnsMap']:
        return pulumi.get(self, "dns")

    @property
    @pulumi.getter(name="eventConfig")
    def event_config(self) -> Optional['outputs.ApiEventConfig']:
        """
        Describes the authorization configuration for connections, message publishing, message subscriptions, and logging for an Event API.
        """
        return pulumi.get(self, "event_config")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        """
        The name of the `Api` .
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerContact")
    def owner_contact(self) -> Optional[builtins.str]:
        """
        The owner contact information for an API resource.

        This field accepts any string input with a length of 0 - 256 characters.
        """
        return pulumi.get(self, "owner_contact")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        A set of tags (key-value pairs) for this API.
        """
        return pulumi.get(self, "tags")


class AwaitableGetApiResult(GetApiResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApiResult(
            api_arn=self.api_arn,
            api_id=self.api_id,
            dns=self.dns,
            event_config=self.event_config,
            name=self.name,
            owner_contact=self.owner_contact,
            tags=self.tags)


def get_api(api_arn: Optional[builtins.str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApiResult:
    """
    Resource schema for AppSync Api


    :param builtins.str api_arn: The Amazon Resource Name (ARN) of the AppSync Api
    """
    __args__ = dict()
    __args__['apiArn'] = api_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:appsync:getApi', __args__, opts=opts, typ=GetApiResult).value

    return AwaitableGetApiResult(
        api_arn=pulumi.get(__ret__, 'api_arn'),
        api_id=pulumi.get(__ret__, 'api_id'),
        dns=pulumi.get(__ret__, 'dns'),
        event_config=pulumi.get(__ret__, 'event_config'),
        name=pulumi.get(__ret__, 'name'),
        owner_contact=pulumi.get(__ret__, 'owner_contact'),
        tags=pulumi.get(__ret__, 'tags'))
def get_api_output(api_arn: Optional[pulumi.Input[builtins.str]] = None,
                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetApiResult]:
    """
    Resource schema for AppSync Api


    :param builtins.str api_arn: The Amazon Resource Name (ARN) of the AppSync Api
    """
    __args__ = dict()
    __args__['apiArn'] = api_arn
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:appsync:getApi', __args__, opts=opts, typ=GetApiResult)
    return __ret__.apply(lambda __response__: GetApiResult(
        api_arn=pulumi.get(__response__, 'api_arn'),
        api_id=pulumi.get(__response__, 'api_id'),
        dns=pulumi.get(__response__, 'dns'),
        event_config=pulumi.get(__response__, 'event_config'),
        name=pulumi.get(__response__, 'name'),
        owner_contact=pulumi.get(__response__, 'owner_contact'),
        tags=pulumi.get(__response__, 'tags')))
