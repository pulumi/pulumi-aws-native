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
from .. import outputs as _root_outputs

__all__ = [
    'GetRegistryResult',
    'AwaitableGetRegistryResult',
    'get_registry',
    'get_registry_output',
]

@pulumi.output_type
class GetRegistryResult:
    def __init__(__self__, description=None, registry_arn=None, tags=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if registry_arn and not isinstance(registry_arn, str):
            raise TypeError("Expected argument 'registry_arn' to be a str")
        pulumi.set(__self__, "registry_arn", registry_arn)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        A description of the registry to be created.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="registryArn")
    def registry_arn(self) -> Optional[builtins.str]:
        """
        The ARN of the registry.
        """
        return pulumi.get(self, "registry_arn")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        Tags associated with the resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetRegistryResult(GetRegistryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegistryResult(
            description=self.description,
            registry_arn=self.registry_arn,
            tags=self.tags)


def get_registry(registry_arn: Optional[builtins.str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegistryResult:
    """
    Resource Type definition for AWS::EventSchemas::Registry


    :param builtins.str registry_arn: The ARN of the registry.
    """
    __args__ = dict()
    __args__['registryArn'] = registry_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:eventschemas:getRegistry', __args__, opts=opts, typ=GetRegistryResult).value

    return AwaitableGetRegistryResult(
        description=pulumi.get(__ret__, 'description'),
        registry_arn=pulumi.get(__ret__, 'registry_arn'),
        tags=pulumi.get(__ret__, 'tags'))
def get_registry_output(registry_arn: Optional[pulumi.Input[builtins.str]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRegistryResult]:
    """
    Resource Type definition for AWS::EventSchemas::Registry


    :param builtins.str registry_arn: The ARN of the registry.
    """
    __args__ = dict()
    __args__['registryArn'] = registry_arn
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:eventschemas:getRegistry', __args__, opts=opts, typ=GetRegistryResult)
    return __ret__.apply(lambda __response__: GetRegistryResult(
        description=pulumi.get(__response__, 'description'),
        registry_arn=pulumi.get(__response__, 'registry_arn'),
        tags=pulumi.get(__response__, 'tags')))
