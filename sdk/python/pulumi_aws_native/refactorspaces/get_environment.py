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
    'GetEnvironmentResult',
    'AwaitableGetEnvironmentResult',
    'get_environment',
    'get_environment_output',
]

@pulumi.output_type
class GetEnvironmentResult:
    def __init__(__self__, arn=None, environment_identifier=None, tags=None, transit_gateway_id=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if environment_identifier and not isinstance(environment_identifier, str):
            raise TypeError("Expected argument 'environment_identifier' to be a str")
        pulumi.set(__self__, "environment_identifier", environment_identifier)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if transit_gateway_id and not isinstance(transit_gateway_id, str):
            raise TypeError("Expected argument 'transit_gateway_id' to be a str")
        pulumi.set(__self__, "transit_gateway_id", transit_gateway_id)

    @property
    @pulumi.getter
    def arn(self) -> Optional[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the environment.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="environmentIdentifier")
    def environment_identifier(self) -> Optional[builtins.str]:
        """
        The unique identifier of the environment.
        """
        return pulumi.get(self, "environment_identifier")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="transitGatewayId")
    def transit_gateway_id(self) -> Optional[builtins.str]:
        """
        The ID of the AWS Transit Gateway set up by the environment.
        """
        return pulumi.get(self, "transit_gateway_id")


class AwaitableGetEnvironmentResult(GetEnvironmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEnvironmentResult(
            arn=self.arn,
            environment_identifier=self.environment_identifier,
            tags=self.tags,
            transit_gateway_id=self.transit_gateway_id)


def get_environment(environment_identifier: Optional[builtins.str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEnvironmentResult:
    """
    Definition of AWS::RefactorSpaces::Environment Resource Type


    :param builtins.str environment_identifier: The unique identifier of the environment.
    """
    __args__ = dict()
    __args__['environmentIdentifier'] = environment_identifier
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:refactorspaces:getEnvironment', __args__, opts=opts, typ=GetEnvironmentResult).value

    return AwaitableGetEnvironmentResult(
        arn=pulumi.get(__ret__, 'arn'),
        environment_identifier=pulumi.get(__ret__, 'environment_identifier'),
        tags=pulumi.get(__ret__, 'tags'),
        transit_gateway_id=pulumi.get(__ret__, 'transit_gateway_id'))
def get_environment_output(environment_identifier: Optional[pulumi.Input[builtins.str]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetEnvironmentResult]:
    """
    Definition of AWS::RefactorSpaces::Environment Resource Type


    :param builtins.str environment_identifier: The unique identifier of the environment.
    """
    __args__ = dict()
    __args__['environmentIdentifier'] = environment_identifier
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:refactorspaces:getEnvironment', __args__, opts=opts, typ=GetEnvironmentResult)
    return __ret__.apply(lambda __response__: GetEnvironmentResult(
        arn=pulumi.get(__response__, 'arn'),
        environment_identifier=pulumi.get(__response__, 'environment_identifier'),
        tags=pulumi.get(__response__, 'tags'),
        transit_gateway_id=pulumi.get(__response__, 'transit_gateway_id')))
