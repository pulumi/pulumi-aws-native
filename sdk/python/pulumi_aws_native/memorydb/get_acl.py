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
    'GetAclResult',
    'AwaitableGetAclResult',
    'get_acl',
    'get_acl_output',
]

@pulumi.output_type
class GetAclResult:
    def __init__(__self__, arn=None, status=None, tags=None, user_names=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if user_names and not isinstance(user_names, list):
            raise TypeError("Expected argument 'user_names' to be a list")
        pulumi.set(__self__, "user_names", user_names)

    @property
    @pulumi.getter
    def arn(self) -> Optional[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the acl.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def status(self) -> Optional[builtins.str]:
        """
        Indicates acl status. Can be "creating", "active", "modifying", "deleting".
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        An array of key-value pairs to apply to this cluster.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="userNames")
    def user_names(self) -> Optional[Sequence[builtins.str]]:
        """
        List of users associated to this acl.
        """
        return pulumi.get(self, "user_names")


class AwaitableGetAclResult(GetAclResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAclResult(
            arn=self.arn,
            status=self.status,
            tags=self.tags,
            user_names=self.user_names)


def get_acl(acl_name: Optional[builtins.str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAclResult:
    """
    Resource Type definition for AWS::MemoryDB::ACL


    :param builtins.str acl_name: The name of the acl.
    """
    __args__ = dict()
    __args__['aclName'] = acl_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:memorydb:getAcl', __args__, opts=opts, typ=GetAclResult).value

    return AwaitableGetAclResult(
        arn=pulumi.get(__ret__, 'arn'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        user_names=pulumi.get(__ret__, 'user_names'))
def get_acl_output(acl_name: Optional[pulumi.Input[builtins.str]] = None,
                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAclResult]:
    """
    Resource Type definition for AWS::MemoryDB::ACL


    :param builtins.str acl_name: The name of the acl.
    """
    __args__ = dict()
    __args__['aclName'] = acl_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:memorydb:getAcl', __args__, opts=opts, typ=GetAclResult)
    return __ret__.apply(lambda __response__: GetAclResult(
        arn=pulumi.get(__response__, 'arn'),
        status=pulumi.get(__response__, 'status'),
        tags=pulumi.get(__response__, 'tags'),
        user_names=pulumi.get(__response__, 'user_names')))
