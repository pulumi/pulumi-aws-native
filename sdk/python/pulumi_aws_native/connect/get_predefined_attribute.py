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
    'GetPredefinedAttributeResult',
    'AwaitableGetPredefinedAttributeResult',
    'get_predefined_attribute',
    'get_predefined_attribute_output',
]

@pulumi.output_type
class GetPredefinedAttributeResult:
    def __init__(__self__, values=None):
        if values and not isinstance(values, dict):
            raise TypeError("Expected argument 'values' to be a dict")
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def values(self) -> Optional['outputs.ValuesProperties']:
        """
        The values of a predefined attribute.
        """
        return pulumi.get(self, "values")


class AwaitableGetPredefinedAttributeResult(GetPredefinedAttributeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPredefinedAttributeResult(
            values=self.values)


def get_predefined_attribute(instance_arn: Optional[str] = None,
                             name: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPredefinedAttributeResult:
    """
    Resource Type definition for AWS::Connect::PredefinedAttribute


    :param str instance_arn: The identifier of the Amazon Connect instance.
    :param str name: The name of the predefined attribute.
    """
    __args__ = dict()
    __args__['instanceArn'] = instance_arn
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:connect:getPredefinedAttribute', __args__, opts=opts, typ=GetPredefinedAttributeResult).value

    return AwaitableGetPredefinedAttributeResult(
        values=pulumi.get(__ret__, 'values'))


@_utilities.lift_output_func(get_predefined_attribute)
def get_predefined_attribute_output(instance_arn: Optional[pulumi.Input[str]] = None,
                                    name: Optional[pulumi.Input[str]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPredefinedAttributeResult]:
    """
    Resource Type definition for AWS::Connect::PredefinedAttribute


    :param str instance_arn: The identifier of the Amazon Connect instance.
    :param str name: The name of the predefined attribute.
    """
    ...