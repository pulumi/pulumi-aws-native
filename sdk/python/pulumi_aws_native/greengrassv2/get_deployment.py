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
    'GetDeploymentResult',
    'AwaitableGetDeploymentResult',
    'get_deployment',
    'get_deployment_output',
]

@pulumi.output_type
class GetDeploymentResult:
    def __init__(__self__, deployment_id=None, tags=None):
        if deployment_id and not isinstance(deployment_id, str):
            raise TypeError("Expected argument 'deployment_id' to be a str")
        pulumi.set(__self__, "deployment_id", deployment_id)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="deploymentId")
    def deployment_id(self) -> Optional[str]:
        return pulumi.get(self, "deployment_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Any]:
        return pulumi.get(self, "tags")


class AwaitableGetDeploymentResult(GetDeploymentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeploymentResult(
            deployment_id=self.deployment_id,
            tags=self.tags)


def get_deployment(deployment_id: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeploymentResult:
    """
    Resource for Greengrass V2 deployment.
    """
    __args__ = dict()
    __args__['deploymentId'] = deployment_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:greengrassv2:getDeployment', __args__, opts=opts, typ=GetDeploymentResult).value

    return AwaitableGetDeploymentResult(
        deployment_id=__ret__.deployment_id,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_deployment)
def get_deployment_output(deployment_id: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDeploymentResult]:
    """
    Resource for Greengrass V2 deployment.
    """
    ...