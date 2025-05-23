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
    'GetVpnConnectionResult',
    'AwaitableGetVpnConnectionResult',
    'get_vpn_connection',
    'get_vpn_connection_output',
]

@pulumi.output_type
class GetVpnConnectionResult:
    def __init__(__self__, tags=None, vpn_connection_id=None):
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if vpn_connection_id and not isinstance(vpn_connection_id, str):
            raise TypeError("Expected argument 'vpn_connection_id' to be a str")
        pulumi.set(__self__, "vpn_connection_id", vpn_connection_id)

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        Any tags assigned to the VPN connection.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpnConnectionId")
    def vpn_connection_id(self) -> Optional[builtins.str]:
        """
        The ID of the VPN connection.
        """
        return pulumi.get(self, "vpn_connection_id")


class AwaitableGetVpnConnectionResult(GetVpnConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpnConnectionResult(
            tags=self.tags,
            vpn_connection_id=self.vpn_connection_id)


def get_vpn_connection(vpn_connection_id: Optional[builtins.str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpnConnectionResult:
    """
    Specifies a VPN connection between a virtual private gateway and a VPN customer gateway or a transit gateway and a VPN customer gateway.
     To specify a VPN connection between a transit gateway and customer gateway, use the ``TransitGatewayId`` and ``CustomerGatewayId`` properties.
     To specify a VPN connection between a virtual private gateway and customer gateway, use the ``VpnGatewayId`` and ``CustomerGatewayId`` properties.
     For more information, see [](https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the *User Guide*.


    :param builtins.str vpn_connection_id: The ID of the VPN connection.
    """
    __args__ = dict()
    __args__['vpnConnectionId'] = vpn_connection_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:ec2:getVpnConnection', __args__, opts=opts, typ=GetVpnConnectionResult).value

    return AwaitableGetVpnConnectionResult(
        tags=pulumi.get(__ret__, 'tags'),
        vpn_connection_id=pulumi.get(__ret__, 'vpn_connection_id'))
def get_vpn_connection_output(vpn_connection_id: Optional[pulumi.Input[builtins.str]] = None,
                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetVpnConnectionResult]:
    """
    Specifies a VPN connection between a virtual private gateway and a VPN customer gateway or a transit gateway and a VPN customer gateway.
     To specify a VPN connection between a transit gateway and customer gateway, use the ``TransitGatewayId`` and ``CustomerGatewayId`` properties.
     To specify a VPN connection between a virtual private gateway and customer gateway, use the ``VpnGatewayId`` and ``CustomerGatewayId`` properties.
     For more information, see [](https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the *User Guide*.


    :param builtins.str vpn_connection_id: The ID of the VPN connection.
    """
    __args__ = dict()
    __args__['vpnConnectionId'] = vpn_connection_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:ec2:getVpnConnection', __args__, opts=opts, typ=GetVpnConnectionResult)
    return __ret__.apply(lambda __response__: GetVpnConnectionResult(
        tags=pulumi.get(__response__, 'tags'),
        vpn_connection_id=pulumi.get(__response__, 'vpn_connection_id')))
