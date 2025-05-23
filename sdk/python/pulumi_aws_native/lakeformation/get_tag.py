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

__all__ = [
    'GetTagResult',
    'AwaitableGetTagResult',
    'get_tag',
    'get_tag_output',
]

@pulumi.output_type
class GetTagResult:
    def __init__(__self__, tag_values=None):
        if tag_values and not isinstance(tag_values, list):
            raise TypeError("Expected argument 'tag_values' to be a list")
        pulumi.set(__self__, "tag_values", tag_values)

    @property
    @pulumi.getter(name="tagValues")
    def tag_values(self) -> Optional[Sequence[builtins.str]]:
        """
        A list of possible values an attribute can take.
        """
        return pulumi.get(self, "tag_values")


class AwaitableGetTagResult(GetTagResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTagResult(
            tag_values=self.tag_values)


def get_tag(tag_key: Optional[builtins.str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTagResult:
    """
    A resource schema representing a Lake Formation Tag.


    :param builtins.str tag_key: The key-name for the LF-tag.
    """
    __args__ = dict()
    __args__['tagKey'] = tag_key
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:lakeformation:getTag', __args__, opts=opts, typ=GetTagResult).value

    return AwaitableGetTagResult(
        tag_values=pulumi.get(__ret__, 'tag_values'))
def get_tag_output(tag_key: Optional[pulumi.Input[builtins.str]] = None,
                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetTagResult]:
    """
    A resource schema representing a Lake Formation Tag.


    :param builtins.str tag_key: The key-name for the LF-tag.
    """
    __args__ = dict()
    __args__['tagKey'] = tag_key
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:lakeformation:getTag', __args__, opts=opts, typ=GetTagResult)
    return __ret__.apply(lambda __response__: GetTagResult(
        tag_values=pulumi.get(__response__, 'tag_values')))
