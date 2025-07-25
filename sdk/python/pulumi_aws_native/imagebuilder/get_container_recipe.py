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
    'GetContainerRecipeResult',
    'AwaitableGetContainerRecipeResult',
    'get_container_recipe',
    'get_container_recipe_output',
]

@pulumi.output_type
class GetContainerRecipeResult:
    def __init__(__self__, arn=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the container recipe.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Tags that are attached to the container recipe.
        """
        return pulumi.get(self, "tags")


class AwaitableGetContainerRecipeResult(GetContainerRecipeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetContainerRecipeResult(
            arn=self.arn,
            tags=self.tags)


def get_container_recipe(arn: Optional[builtins.str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetContainerRecipeResult:
    """
    Resource schema for AWS::ImageBuilder::ContainerRecipe


    :param builtins.str arn: The Amazon Resource Name (ARN) of the container recipe.
    """
    __args__ = dict()
    __args__['arn'] = arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:imagebuilder:getContainerRecipe', __args__, opts=opts, typ=GetContainerRecipeResult).value

    return AwaitableGetContainerRecipeResult(
        arn=pulumi.get(__ret__, 'arn'),
        tags=pulumi.get(__ret__, 'tags'))
def get_container_recipe_output(arn: Optional[pulumi.Input[builtins.str]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetContainerRecipeResult]:
    """
    Resource schema for AWS::ImageBuilder::ContainerRecipe


    :param builtins.str arn: The Amazon Resource Name (ARN) of the container recipe.
    """
    __args__ = dict()
    __args__['arn'] = arn
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:imagebuilder:getContainerRecipe', __args__, opts=opts, typ=GetContainerRecipeResult)
    return __ret__.apply(lambda __response__: GetContainerRecipeResult(
        arn=pulumi.get(__response__, 'arn'),
        tags=pulumi.get(__response__, 'tags')))
