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
    'GetVPNConnectionResult',
    'AwaitableGetVPNConnectionResult',
    'get_vpn_connection',
    'get_vpn_connection_output',
]

@pulumi.output_type
class GetVPNConnectionResult:
    def __init__(__self__, tags=None, vpn_connection_id=None):
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if vpn_connection_id and not isinstance(vpn_connection_id, str):
            raise TypeError("Expected argument 'vpn_connection_id' to be a str")
        pulumi.set(__self__, "vpn_connection_id", vpn_connection_id)

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.VPNConnectionTag']]:
        """
        Any tags assigned to the VPN connection.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpnConnectionId")
    def vpn_connection_id(self) -> Optional[str]:
        """
        The provider-assigned unique ID for this managed resource
        """
        return pulumi.get(self, "vpn_connection_id")


class AwaitableGetVPNConnectionResult(GetVPNConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVPNConnectionResult(
            tags=self.tags,
            vpn_connection_id=self.vpn_connection_id)


def get_vpn_connection(vpn_connection_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVPNConnectionResult:
    """
    Resource Type definition for AWS::EC2::VPNConnection


    :param str vpn_connection_id: The provider-assigned unique ID for this managed resource
    """
    __args__ = dict()
    __args__['vpnConnectionId'] = vpn_connection_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:ec2:getVPNConnection', __args__, opts=opts, typ=GetVPNConnectionResult).value

    return AwaitableGetVPNConnectionResult(
        tags=__ret__.tags,
        vpn_connection_id=__ret__.vpn_connection_id)


@_utilities.lift_output_func(get_vpn_connection)
def get_vpn_connection_output(vpn_connection_id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVPNConnectionResult]:
    """
    Resource Type definition for AWS::EC2::VPNConnection


    :param str vpn_connection_id: The provider-assigned unique ID for this managed resource
    """
    ...