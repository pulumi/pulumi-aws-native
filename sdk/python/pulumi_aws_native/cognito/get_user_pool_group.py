# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
    'GetUserPoolGroupResult',
    'AwaitableGetUserPoolGroupResult',
    'get_user_pool_group',
    'get_user_pool_group_output',
]

@pulumi.output_type
class GetUserPoolGroupResult:
    def __init__(__self__, description=None, precedence=None, role_arn=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if precedence and not isinstance(precedence, int):
            raise TypeError("Expected argument 'precedence' to be a int")
        pulumi.set(__self__, "precedence", precedence)
        if role_arn and not isinstance(role_arn, str):
            raise TypeError("Expected argument 'role_arn' to be a str")
        pulumi.set(__self__, "role_arn", role_arn)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A string containing the description of the group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def precedence(self) -> Optional[int]:
        """
        A non-negative integer value that specifies the precedence of this group relative to the other groups that a user can belong to in the user pool. Zero is the highest precedence value. Groups with lower `Precedence` values take precedence over groups with higher or null `Precedence` values. If a user belongs to two or more groups, it is the group with the lowest precedence value whose role ARN is given in the user's tokens for the `cognito:roles` and `cognito:preferred_role` claims.

        Two groups can have the same `Precedence` value. If this happens, neither group takes precedence over the other. If two groups with the same `Precedence` have the same role ARN, that role is used in the `cognito:preferred_role` claim in tokens for users in each group. If the two groups have different role ARNs, the `cognito:preferred_role` claim isn't set in users' tokens.

        The default `Precedence` value is null. The maximum `Precedence` value is `2^31-1` .
        """
        return pulumi.get(self, "precedence")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[str]:
        """
        The role Amazon Resource Name (ARN) for the group.
        """
        return pulumi.get(self, "role_arn")


class AwaitableGetUserPoolGroupResult(GetUserPoolGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserPoolGroupResult(
            description=self.description,
            precedence=self.precedence,
            role_arn=self.role_arn)


def get_user_pool_group(group_name: Optional[str] = None,
                        user_pool_id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserPoolGroupResult:
    """
    Resource Type definition for AWS::Cognito::UserPoolGroup


    :param str group_name: The name of the group. Must be unique.
    :param str user_pool_id: The user pool ID for the user pool.
    """
    __args__ = dict()
    __args__['groupName'] = group_name
    __args__['userPoolId'] = user_pool_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:cognito:getUserPoolGroup', __args__, opts=opts, typ=GetUserPoolGroupResult).value

    return AwaitableGetUserPoolGroupResult(
        description=pulumi.get(__ret__, 'description'),
        precedence=pulumi.get(__ret__, 'precedence'),
        role_arn=pulumi.get(__ret__, 'role_arn'))
def get_user_pool_group_output(group_name: Optional[pulumi.Input[str]] = None,
                               user_pool_id: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserPoolGroupResult]:
    """
    Resource Type definition for AWS::Cognito::UserPoolGroup


    :param str group_name: The name of the group. Must be unique.
    :param str user_pool_id: The user pool ID for the user pool.
    """
    __args__ = dict()
    __args__['groupName'] = group_name
    __args__['userPoolId'] = user_pool_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:cognito:getUserPoolGroup', __args__, opts=opts, typ=GetUserPoolGroupResult)
    return __ret__.apply(lambda __response__: GetUserPoolGroupResult(
        description=pulumi.get(__response__, 'description'),
        precedence=pulumi.get(__response__, 'precedence'),
        role_arn=pulumi.get(__response__, 'role_arn')))
