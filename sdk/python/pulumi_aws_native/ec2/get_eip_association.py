# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetEipAssociationResult',
    'AwaitableGetEipAssociationResult',
    'get_eip_association',
    'get_eip_association_output',
]

@pulumi.output_type
class GetEipAssociationResult:
    def __init__(__self__, allocation_id=None, eip=None, id=None, instance_id=None, network_interface_id=None, private_ip_address=None):
        if allocation_id and not isinstance(allocation_id, str):
            raise TypeError("Expected argument 'allocation_id' to be a str")
        pulumi.set(__self__, "allocation_id", allocation_id)
        if eip and not isinstance(eip, str):
            raise TypeError("Expected argument 'eip' to be a str")
        pulumi.set(__self__, "eip", eip)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if network_interface_id and not isinstance(network_interface_id, str):
            raise TypeError("Expected argument 'network_interface_id' to be a str")
        pulumi.set(__self__, "network_interface_id", network_interface_id)
        if private_ip_address and not isinstance(private_ip_address, str):
            raise TypeError("Expected argument 'private_ip_address' to be a str")
        pulumi.set(__self__, "private_ip_address", private_ip_address)

    @property
    @pulumi.getter(name="allocationId")
    def allocation_id(self) -> Optional[str]:
        return pulumi.get(self, "allocation_id")

    @property
    @pulumi.getter
    def eip(self) -> Optional[str]:
        return pulumi.get(self, "eip")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> Optional[str]:
        return pulumi.get(self, "network_interface_id")

    @property
    @pulumi.getter(name="privateIpAddress")
    def private_ip_address(self) -> Optional[str]:
        return pulumi.get(self, "private_ip_address")


class AwaitableGetEipAssociationResult(GetEipAssociationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEipAssociationResult(
            allocation_id=self.allocation_id,
            eip=self.eip,
            id=self.id,
            instance_id=self.instance_id,
            network_interface_id=self.network_interface_id,
            private_ip_address=self.private_ip_address)


def get_eip_association(id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEipAssociationResult:
    """
    Resource Type definition for AWS::EC2::EIPAssociation
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:ec2:getEipAssociation', __args__, opts=opts, typ=GetEipAssociationResult).value

    return AwaitableGetEipAssociationResult(
        allocation_id=pulumi.get(__ret__, 'allocation_id'),
        eip=pulumi.get(__ret__, 'eip'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        network_interface_id=pulumi.get(__ret__, 'network_interface_id'),
        private_ip_address=pulumi.get(__ret__, 'private_ip_address'))


@_utilities.lift_output_func(get_eip_association)
def get_eip_association_output(id: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEipAssociationResult]:
    """
    Resource Type definition for AWS::EC2::EIPAssociation
    """
    ...