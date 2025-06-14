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
from . import outputs
from .. import outputs as _root_outputs
from ._enums import *

__all__ = [
    'GetUserResult',
    'AwaitableGetUserResult',
    'get_user',
    'get_user_output',
]

@pulumi.output_type
class GetUserResult:
    def __init__(__self__, arn=None, home_directory=None, home_directory_mappings=None, home_directory_type=None, policy=None, posix_profile=None, role=None, ssh_public_keys=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if home_directory and not isinstance(home_directory, str):
            raise TypeError("Expected argument 'home_directory' to be a str")
        pulumi.set(__self__, "home_directory", home_directory)
        if home_directory_mappings and not isinstance(home_directory_mappings, list):
            raise TypeError("Expected argument 'home_directory_mappings' to be a list")
        pulumi.set(__self__, "home_directory_mappings", home_directory_mappings)
        if home_directory_type and not isinstance(home_directory_type, str):
            raise TypeError("Expected argument 'home_directory_type' to be a str")
        pulumi.set(__self__, "home_directory_type", home_directory_type)
        if policy and not isinstance(policy, str):
            raise TypeError("Expected argument 'policy' to be a str")
        pulumi.set(__self__, "policy", policy)
        if posix_profile and not isinstance(posix_profile, dict):
            raise TypeError("Expected argument 'posix_profile' to be a dict")
        pulumi.set(__self__, "posix_profile", posix_profile)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)
        if ssh_public_keys and not isinstance(ssh_public_keys, list):
            raise TypeError("Expected argument 'ssh_public_keys' to be a list")
        pulumi.set(__self__, "ssh_public_keys", ssh_public_keys)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[builtins.str]:
        """
        The Amazon Resource Name associated with the user, in the form `arn:aws:transfer:region: *account-id* :user/ *server-id* / *username*` .

        An example of a user ARN is: `arn:aws:transfer:us-east-1:123456789012:user/user1` .
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="homeDirectory")
    def home_directory(self) -> Optional[builtins.str]:
        """
        The landing directory (folder) for a user when they log in to the server using the client.

        A `HomeDirectory` example is `/bucket_name/home/mydirectory` .

        > You can use the `HomeDirectory` parameter for `HomeDirectoryType` when it is set to either `PATH` or `LOGICAL` .
        """
        return pulumi.get(self, "home_directory")

    @property
    @pulumi.getter(name="homeDirectoryMappings")
    def home_directory_mappings(self) -> Optional[Sequence['outputs.UserHomeDirectoryMapEntry']]:
        """
        Logical directory mappings that specify what Amazon S3 or Amazon EFS paths and keys should be visible to your user and how you want to make them visible. You must specify the `Entry` and `Target` pair, where `Entry` shows how the path is made visible and `Target` is the actual Amazon S3 or Amazon EFS path. If you only specify a target, it is displayed as is. You also must ensure that your AWS Identity and Access Management (IAM) role provides access to paths in `Target` . This value can be set only when `HomeDirectoryType` is set to *LOGICAL* .

        The following is an `Entry` and `Target` pair example.

        `[ { "Entry": "/directory1", "Target": "/bucket_name/home/mydirectory" } ]`

        In most cases, you can use this value instead of the session policy to lock your user down to the designated home directory (" `chroot` "). To do this, you can set `Entry` to `/` and set `Target` to the value the user should see for their home directory when they log in.

        The following is an `Entry` and `Target` pair example for `chroot` .

        `[ { "Entry": "/", "Target": "/bucket_name/home/mydirectory" } ]`
        """
        return pulumi.get(self, "home_directory_mappings")

    @property
    @pulumi.getter(name="homeDirectoryType")
    def home_directory_type(self) -> Optional['UserHomeDirectoryType']:
        """
        The type of landing directory (folder) that you want your users' home directory to be when they log in to the server. If you set it to `PATH` , the user will see the absolute Amazon S3 bucket or Amazon EFS path as is in their file transfer protocol clients. If you set it to `LOGICAL` , you need to provide mappings in the `HomeDirectoryMappings` for how you want to make Amazon S3 or Amazon EFS paths visible to your users.

        > If `HomeDirectoryType` is `LOGICAL` , you must provide mappings, using the `HomeDirectoryMappings` parameter. If, on the other hand, `HomeDirectoryType` is `PATH` , you provide an absolute path using the `HomeDirectory` parameter. You cannot have both `HomeDirectory` and `HomeDirectoryMappings` in your template.
        """
        return pulumi.get(self, "home_directory_type")

    @property
    @pulumi.getter
    def policy(self) -> Optional[builtins.str]:
        """
        A session policy for your user so you can use the same IAM role across multiple users. This policy restricts user access to portions of their Amazon S3 bucket. Variables that you can use inside this policy include `${Transfer:UserName}` , `${Transfer:HomeDirectory}` , and `${Transfer:HomeBucket}` .

        > For session policies, AWS Transfer Family stores the policy as a JSON blob, instead of the Amazon Resource Name (ARN) of the policy. You save the policy as a JSON blob and pass it in the `Policy` argument.
        > 
        > For an example of a session policy, see [Example session policy](https://docs.aws.amazon.com/transfer/latest/userguide/session-policy.html) .
        > 
        > For more information, see [AssumeRole](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html) in the *AWS Security Token Service API Reference* .
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter(name="posixProfile")
    def posix_profile(self) -> Optional['outputs.UserPosixProfile']:
        """
        Specifies the full POSIX identity, including user ID ( `Uid` ), group ID ( `Gid` ), and any secondary groups IDs ( `SecondaryGids` ), that controls your users' access to your Amazon Elastic File System (Amazon EFS) file systems. The POSIX permissions that are set on files and directories in your file system determine the level of access your users get when transferring files into and out of your Amazon EFS file systems.
        """
        return pulumi.get(self, "posix_profile")

    @property
    @pulumi.getter
    def role(self) -> Optional[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that controls your users' access to your Amazon S3 bucket or Amazon EFS file system. The policies attached to this role determine the level of access that you want to provide your users when transferring files into and out of your Amazon S3 bucket or Amazon EFS file system. The IAM role should also contain a trust relationship that allows the server to access your resources when servicing your users' transfer requests.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="sshPublicKeys")
    def ssh_public_keys(self) -> Optional[Sequence[builtins.str]]:
        """
        This represents the SSH User Public Keys for CloudFormation resource
        """
        return pulumi.get(self, "ssh_public_keys")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        Key-value pairs that can be used to group and search for users. Tags are metadata attached to users for any purpose.
        """
        return pulumi.get(self, "tags")


class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            arn=self.arn,
            home_directory=self.home_directory,
            home_directory_mappings=self.home_directory_mappings,
            home_directory_type=self.home_directory_type,
            policy=self.policy,
            posix_profile=self.posix_profile,
            role=self.role,
            ssh_public_keys=self.ssh_public_keys,
            tags=self.tags)


def get_user(arn: Optional[builtins.str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserResult:
    """
    Definition of AWS::Transfer::User Resource Type


    :param builtins.str arn: The Amazon Resource Name associated with the user, in the form `arn:aws:transfer:region: *account-id* :user/ *server-id* / *username*` .
           
           An example of a user ARN is: `arn:aws:transfer:us-east-1:123456789012:user/user1` .
    """
    __args__ = dict()
    __args__['arn'] = arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:transfer:getUser', __args__, opts=opts, typ=GetUserResult).value

    return AwaitableGetUserResult(
        arn=pulumi.get(__ret__, 'arn'),
        home_directory=pulumi.get(__ret__, 'home_directory'),
        home_directory_mappings=pulumi.get(__ret__, 'home_directory_mappings'),
        home_directory_type=pulumi.get(__ret__, 'home_directory_type'),
        policy=pulumi.get(__ret__, 'policy'),
        posix_profile=pulumi.get(__ret__, 'posix_profile'),
        role=pulumi.get(__ret__, 'role'),
        ssh_public_keys=pulumi.get(__ret__, 'ssh_public_keys'),
        tags=pulumi.get(__ret__, 'tags'))
def get_user_output(arn: Optional[pulumi.Input[builtins.str]] = None,
                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetUserResult]:
    """
    Definition of AWS::Transfer::User Resource Type


    :param builtins.str arn: The Amazon Resource Name associated with the user, in the form `arn:aws:transfer:region: *account-id* :user/ *server-id* / *username*` .
           
           An example of a user ARN is: `arn:aws:transfer:us-east-1:123456789012:user/user1` .
    """
    __args__ = dict()
    __args__['arn'] = arn
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:transfer:getUser', __args__, opts=opts, typ=GetUserResult)
    return __ret__.apply(lambda __response__: GetUserResult(
        arn=pulumi.get(__response__, 'arn'),
        home_directory=pulumi.get(__response__, 'home_directory'),
        home_directory_mappings=pulumi.get(__response__, 'home_directory_mappings'),
        home_directory_type=pulumi.get(__response__, 'home_directory_type'),
        policy=pulumi.get(__response__, 'policy'),
        posix_profile=pulumi.get(__response__, 'posix_profile'),
        role=pulumi.get(__response__, 'role'),
        ssh_public_keys=pulumi.get(__response__, 'ssh_public_keys'),
        tags=pulumi.get(__response__, 'tags')))
