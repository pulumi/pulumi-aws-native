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
    'GetSubnetGroupResult',
    'AwaitableGetSubnetGroupResult',
    'get_subnet_group',
    'get_subnet_group_output',
]

@pulumi.output_type
class GetSubnetGroupResult:
    def __init__(__self__, a_rn=None, description=None, subnet_ids=None, tags=None):
        if a_rn and not isinstance(a_rn, str):
            raise TypeError("Expected argument 'a_rn' to be a str")
        pulumi.set(__self__, "a_rn", a_rn)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if subnet_ids and not isinstance(subnet_ids, list):
            raise TypeError("Expected argument 'subnet_ids' to be a list")
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="aRN")
    def a_rn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the subnet group.
        """
        return pulumi.get(self, "a_rn")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        An optional description of the subnet group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Optional[Sequence[str]]:
        """
        A list of VPC subnet IDs for the subnet group.
        """
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.SubnetGroupTag']]:
        """
        An array of key-value pairs to apply to this subnet group.
        """
        return pulumi.get(self, "tags")


class AwaitableGetSubnetGroupResult(GetSubnetGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSubnetGroupResult(
            a_rn=self.a_rn,
            description=self.description,
            subnet_ids=self.subnet_ids,
            tags=self.tags)


def get_subnet_group(subnet_group_name: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSubnetGroupResult:
    """
    The AWS::MemoryDB::SubnetGroup resource creates an Amazon MemoryDB Subnet Group.


    :param str subnet_group_name: The name of the subnet group. This value must be unique as it also serves as the subnet group identifier.
    """
    __args__ = dict()
    __args__['subnetGroupName'] = subnet_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:memorydb:getSubnetGroup', __args__, opts=opts, typ=GetSubnetGroupResult).value

    return AwaitableGetSubnetGroupResult(
        a_rn=__ret__.a_rn,
        description=__ret__.description,
        subnet_ids=__ret__.subnet_ids,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_subnet_group)
def get_subnet_group_output(subnet_group_name: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSubnetGroupResult]:
    """
    The AWS::MemoryDB::SubnetGroup resource creates an Amazon MemoryDB Subnet Group.


    :param str subnet_group_name: The name of the subnet group. This value must be unique as it also serves as the subnet group identifier.
    """
    ...