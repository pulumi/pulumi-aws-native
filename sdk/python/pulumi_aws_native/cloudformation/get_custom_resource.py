# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetCustomResourceResult',
    'AwaitableGetCustomResourceResult',
    'get_custom_resource',
    'get_custom_resource_output',
]

@pulumi.output_type
class GetCustomResourceResult:
    def __init__(__self__, id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")


class AwaitableGetCustomResourceResult(GetCustomResourceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCustomResourceResult(
            id=self.id)


def get_custom_resource(id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCustomResourceResult:
    """
    Resource Type definition for AWS::CloudFormation::CustomResource
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:cloudformation:getCustomResource', __args__, opts=opts, typ=GetCustomResourceResult).value

    return AwaitableGetCustomResourceResult(
        id=__ret__.id)


@_utilities.lift_output_func(get_custom_resource)
def get_custom_resource_output(id: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCustomResourceResult]:
    """
    Resource Type definition for AWS::CloudFormation::CustomResource
    """
    ...