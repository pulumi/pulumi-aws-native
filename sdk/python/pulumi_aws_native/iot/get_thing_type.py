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
    'GetThingTypeResult',
    'AwaitableGetThingTypeResult',
    'get_thing_type',
    'get_thing_type_output',
]

@pulumi.output_type
class GetThingTypeResult:
    def __init__(__self__, arn=None, deprecate_thing_type=None, id=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if deprecate_thing_type and not isinstance(deprecate_thing_type, bool):
            raise TypeError("Expected argument 'deprecate_thing_type' to be a bool")
        pulumi.set(__self__, "deprecate_thing_type", deprecate_thing_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="deprecateThingType")
    def deprecate_thing_type(self) -> Optional[bool]:
        return pulumi.get(self, "deprecate_thing_type")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.ThingTypeTag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetThingTypeResult(GetThingTypeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetThingTypeResult(
            arn=self.arn,
            deprecate_thing_type=self.deprecate_thing_type,
            id=self.id,
            tags=self.tags)


def get_thing_type(thing_type_name: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetThingTypeResult:
    """
    Resource Type definition for AWS::IoT::ThingType
    """
    __args__ = dict()
    __args__['thingTypeName'] = thing_type_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:iot:getThingType', __args__, opts=opts, typ=GetThingTypeResult).value

    return AwaitableGetThingTypeResult(
        arn=__ret__.arn,
        deprecate_thing_type=__ret__.deprecate_thing_type,
        id=__ret__.id,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_thing_type)
def get_thing_type_output(thing_type_name: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetThingTypeResult]:
    """
    Resource Type definition for AWS::IoT::ThingType
    """
    ...