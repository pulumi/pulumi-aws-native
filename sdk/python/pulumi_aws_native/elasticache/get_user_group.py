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
    'GetUserGroupResult',
    'AwaitableGetUserGroupResult',
    'get_user_group',
    'get_user_group_output',
]

@pulumi.output_type
class GetUserGroupResult:
    def __init__(__self__, arn=None, status=None, user_ids=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if user_ids and not isinstance(user_ids, list):
            raise TypeError("Expected argument 'user_ids' to be a list")
        pulumi.set(__self__, "user_ids", user_ids)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the user account.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Indicates user group status. Can be "creating", "active", "modifying", "deleting".
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="userIds")
    def user_ids(self) -> Optional[Sequence[str]]:
        """
        List of users associated to this user group.
        """
        return pulumi.get(self, "user_ids")


class AwaitableGetUserGroupResult(GetUserGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserGroupResult(
            arn=self.arn,
            status=self.status,
            user_ids=self.user_ids)


def get_user_group(user_group_id: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserGroupResult:
    """
    Resource Type definition for AWS::ElastiCache::UserGroup


    :param str user_group_id: The ID of the user group.
    """
    __args__ = dict()
    __args__['userGroupId'] = user_group_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:elasticache:getUserGroup', __args__, opts=opts, typ=GetUserGroupResult).value

    return AwaitableGetUserGroupResult(
        arn=__ret__.arn,
        status=__ret__.status,
        user_ids=__ret__.user_ids)


@_utilities.lift_output_func(get_user_group)
def get_user_group_output(user_group_id: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserGroupResult]:
    """
    Resource Type definition for AWS::ElastiCache::UserGroup


    :param str user_group_id: The ID of the user group.
    """
    ...