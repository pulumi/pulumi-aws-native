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
    'GetStorageConfigurationResult',
    'AwaitableGetStorageConfigurationResult',
    'get_storage_configuration',
    'get_storage_configuration_output',
]

@pulumi.output_type
class GetStorageConfigurationResult:
    def __init__(__self__, arn=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[builtins.str]:
        """
        Storage Configuration ARN is automatically generated on creation and assigned as the unique identifier.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        A list of key-value pairs that contain metadata for the asset model.
        """
        return pulumi.get(self, "tags")


class AwaitableGetStorageConfigurationResult(GetStorageConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStorageConfigurationResult(
            arn=self.arn,
            tags=self.tags)


def get_storage_configuration(arn: Optional[builtins.str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStorageConfigurationResult:
    """
    Resource Type definition for AWS::IVS::StorageConfiguration


    :param builtins.str arn: Storage Configuration ARN is automatically generated on creation and assigned as the unique identifier.
    """
    __args__ = dict()
    __args__['arn'] = arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:ivs:getStorageConfiguration', __args__, opts=opts, typ=GetStorageConfigurationResult).value

    return AwaitableGetStorageConfigurationResult(
        arn=pulumi.get(__ret__, 'arn'),
        tags=pulumi.get(__ret__, 'tags'))
def get_storage_configuration_output(arn: Optional[pulumi.Input[builtins.str]] = None,
                                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetStorageConfigurationResult]:
    """
    Resource Type definition for AWS::IVS::StorageConfiguration


    :param builtins.str arn: Storage Configuration ARN is automatically generated on creation and assigned as the unique identifier.
    """
    __args__ = dict()
    __args__['arn'] = arn
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:ivs:getStorageConfiguration', __args__, opts=opts, typ=GetStorageConfigurationResult)
    return __ret__.apply(lambda __response__: GetStorageConfigurationResult(
        arn=pulumi.get(__response__, 'arn'),
        tags=pulumi.get(__response__, 'tags')))
