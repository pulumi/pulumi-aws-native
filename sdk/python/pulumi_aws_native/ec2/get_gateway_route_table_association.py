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
    'GetGatewayRouteTableAssociationResult',
    'AwaitableGetGatewayRouteTableAssociationResult',
    'get_gateway_route_table_association',
    'get_gateway_route_table_association_output',
]

@pulumi.output_type
class GetGatewayRouteTableAssociationResult:
    def __init__(__self__, association_id=None, route_table_id=None):
        if association_id and not isinstance(association_id, str):
            raise TypeError("Expected argument 'association_id' to be a str")
        pulumi.set(__self__, "association_id", association_id)
        if route_table_id and not isinstance(route_table_id, str):
            raise TypeError("Expected argument 'route_table_id' to be a str")
        pulumi.set(__self__, "route_table_id", route_table_id)

    @property
    @pulumi.getter(name="associationId")
    def association_id(self) -> Optional[builtins.str]:
        """
        The route table association ID.
        """
        return pulumi.get(self, "association_id")

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> Optional[builtins.str]:
        """
        The ID of the route table.
        """
        return pulumi.get(self, "route_table_id")


class AwaitableGetGatewayRouteTableAssociationResult(GetGatewayRouteTableAssociationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGatewayRouteTableAssociationResult(
            association_id=self.association_id,
            route_table_id=self.route_table_id)


def get_gateway_route_table_association(gateway_id: Optional[builtins.str] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGatewayRouteTableAssociationResult:
    """
    Associates a gateway with a route table. The gateway and route table must be in the same VPC. This association causes the incoming traffic to the gateway to be routed according to the routes in the route table.


    :param builtins.str gateway_id: The ID of the gateway.
    """
    __args__ = dict()
    __args__['gatewayId'] = gateway_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:ec2:getGatewayRouteTableAssociation', __args__, opts=opts, typ=GetGatewayRouteTableAssociationResult).value

    return AwaitableGetGatewayRouteTableAssociationResult(
        association_id=pulumi.get(__ret__, 'association_id'),
        route_table_id=pulumi.get(__ret__, 'route_table_id'))
def get_gateway_route_table_association_output(gateway_id: Optional[pulumi.Input[builtins.str]] = None,
                                               opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetGatewayRouteTableAssociationResult]:
    """
    Associates a gateway with a route table. The gateway and route table must be in the same VPC. This association causes the incoming traffic to the gateway to be routed according to the routes in the route table.


    :param builtins.str gateway_id: The ID of the gateway.
    """
    __args__ = dict()
    __args__['gatewayId'] = gateway_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:ec2:getGatewayRouteTableAssociation', __args__, opts=opts, typ=GetGatewayRouteTableAssociationResult)
    return __ret__.apply(lambda __response__: GetGatewayRouteTableAssociationResult(
        association_id=pulumi.get(__response__, 'association_id'),
        route_table_id=pulumi.get(__response__, 'route_table_id')))
