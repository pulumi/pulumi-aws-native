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
    'GetCustomDataIdentifierResult',
    'AwaitableGetCustomDataIdentifierResult',
    'get_custom_data_identifier',
    'get_custom_data_identifier_output',
]

@pulumi.output_type
class GetCustomDataIdentifierResult:
    def __init__(__self__, arn=None, id=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[builtins.str]:
        """
        Custom data identifier ARN.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def id(self) -> Optional[builtins.str]:
        """
        Custom data identifier ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        A collection of tags associated with a resource
        """
        return pulumi.get(self, "tags")


class AwaitableGetCustomDataIdentifierResult(GetCustomDataIdentifierResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCustomDataIdentifierResult(
            arn=self.arn,
            id=self.id,
            tags=self.tags)


def get_custom_data_identifier(id: Optional[builtins.str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCustomDataIdentifierResult:
    """
    Macie CustomDataIdentifier resource schema


    :param builtins.str id: Custom data identifier ID.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:macie:getCustomDataIdentifier', __args__, opts=opts, typ=GetCustomDataIdentifierResult).value

    return AwaitableGetCustomDataIdentifierResult(
        arn=pulumi.get(__ret__, 'arn'),
        id=pulumi.get(__ret__, 'id'),
        tags=pulumi.get(__ret__, 'tags'))
def get_custom_data_identifier_output(id: Optional[pulumi.Input[builtins.str]] = None,
                                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetCustomDataIdentifierResult]:
    """
    Macie CustomDataIdentifier resource schema


    :param builtins.str id: Custom data identifier ID.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:macie:getCustomDataIdentifier', __args__, opts=opts, typ=GetCustomDataIdentifierResult)
    return __ret__.apply(lambda __response__: GetCustomDataIdentifierResult(
        arn=pulumi.get(__response__, 'arn'),
        id=pulumi.get(__response__, 'id'),
        tags=pulumi.get(__response__, 'tags')))
