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
    'GetDedicatedIpPoolResult',
    'AwaitableGetDedicatedIpPoolResult',
    'get_dedicated_ip_pool',
    'get_dedicated_ip_pool_output',
]

@pulumi.output_type
class GetDedicatedIpPoolResult:
    def __init__(__self__, id=None, tags=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DedicatedIpPoolTags']]:
        return pulumi.get(self, "tags")


class AwaitableGetDedicatedIpPoolResult(GetDedicatedIpPoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDedicatedIpPoolResult(
            id=self.id,
            tags=self.tags)


def get_dedicated_ip_pool(id: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDedicatedIpPoolResult:
    """
    Resource Type definition for AWS::PinpointEmail::DedicatedIpPool
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:pinpointemail:getDedicatedIpPool', __args__, opts=opts, typ=GetDedicatedIpPoolResult).value

    return AwaitableGetDedicatedIpPoolResult(
        id=__ret__.id,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_dedicated_ip_pool)
def get_dedicated_ip_pool_output(id: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDedicatedIpPoolResult]:
    """
    Resource Type definition for AWS::PinpointEmail::DedicatedIpPool
    """
    ...