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

__all__ = [
    'GetDiscovererResult',
    'AwaitableGetDiscovererResult',
    'get_discoverer',
    'get_discoverer_output',
]

@pulumi.output_type
class GetDiscovererResult:
    def __init__(__self__, cross_account=None, description=None, discoverer_arn=None, discoverer_id=None, tags=None):
        if cross_account and not isinstance(cross_account, bool):
            raise TypeError("Expected argument 'cross_account' to be a bool")
        pulumi.set(__self__, "cross_account", cross_account)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if discoverer_arn and not isinstance(discoverer_arn, str):
            raise TypeError("Expected argument 'discoverer_arn' to be a str")
        pulumi.set(__self__, "discoverer_arn", discoverer_arn)
        if discoverer_id and not isinstance(discoverer_id, str):
            raise TypeError("Expected argument 'discoverer_id' to be a str")
        pulumi.set(__self__, "discoverer_id", discoverer_id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="crossAccount")
    def cross_account(self) -> Optional[bool]:
        return pulumi.get(self, "cross_account")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="discovererArn")
    def discoverer_arn(self) -> Optional[str]:
        return pulumi.get(self, "discoverer_arn")

    @property
    @pulumi.getter(name="discovererId")
    def discoverer_id(self) -> Optional[str]:
        return pulumi.get(self, "discoverer_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DiscovererTagsEntry']]:
        return pulumi.get(self, "tags")


class AwaitableGetDiscovererResult(GetDiscovererResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDiscovererResult(
            cross_account=self.cross_account,
            description=self.description,
            discoverer_arn=self.discoverer_arn,
            discoverer_id=self.discoverer_id,
            tags=self.tags)


def get_discoverer(discoverer_id: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDiscovererResult:
    """
    Resource Type definition for AWS::EventSchemas::Discoverer
    """
    __args__ = dict()
    __args__['discovererId'] = discoverer_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:eventschemas:getDiscoverer', __args__, opts=opts, typ=GetDiscovererResult).value

    return AwaitableGetDiscovererResult(
        cross_account=__ret__.cross_account,
        description=__ret__.description,
        discoverer_arn=__ret__.discoverer_arn,
        discoverer_id=__ret__.discoverer_id,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_discoverer)
def get_discoverer_output(discoverer_id: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDiscovererResult]:
    """
    Resource Type definition for AWS::EventSchemas::Discoverer
    """
    ...