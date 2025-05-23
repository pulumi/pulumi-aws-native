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
    'GetEgressOnlyInternetGatewayResult',
    'AwaitableGetEgressOnlyInternetGatewayResult',
    'get_egress_only_internet_gateway',
    'get_egress_only_internet_gateway_output',
]

@pulumi.output_type
class GetEgressOnlyInternetGatewayResult:
    def __init__(__self__, id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[builtins.str]:
        """
        Service Generated ID of the EgressOnlyInternetGateway
        """
        return pulumi.get(self, "id")


class AwaitableGetEgressOnlyInternetGatewayResult(GetEgressOnlyInternetGatewayResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEgressOnlyInternetGatewayResult(
            id=self.id)


def get_egress_only_internet_gateway(id: Optional[builtins.str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEgressOnlyInternetGatewayResult:
    """
    Resource Type definition for AWS::EC2::EgressOnlyInternetGateway


    :param builtins.str id: Service Generated ID of the EgressOnlyInternetGateway
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:ec2:getEgressOnlyInternetGateway', __args__, opts=opts, typ=GetEgressOnlyInternetGatewayResult).value

    return AwaitableGetEgressOnlyInternetGatewayResult(
        id=pulumi.get(__ret__, 'id'))
def get_egress_only_internet_gateway_output(id: Optional[pulumi.Input[builtins.str]] = None,
                                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetEgressOnlyInternetGatewayResult]:
    """
    Resource Type definition for AWS::EC2::EgressOnlyInternetGateway


    :param builtins.str id: Service Generated ID of the EgressOnlyInternetGateway
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:ec2:getEgressOnlyInternetGateway', __args__, opts=opts, typ=GetEgressOnlyInternetGatewayResult)
    return __ret__.apply(lambda __response__: GetEgressOnlyInternetGatewayResult(
        id=pulumi.get(__response__, 'id')))
