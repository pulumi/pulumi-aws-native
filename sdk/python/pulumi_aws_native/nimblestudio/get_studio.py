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

__all__ = [
    'GetStudioResult',
    'AwaitableGetStudioResult',
    'get_studio',
    'get_studio_output',
]

@pulumi.output_type
class GetStudioResult:
    def __init__(__self__, admin_role_arn=None, display_name=None, home_region=None, sso_client_id=None, studio_encryption_configuration=None, studio_id=None, studio_url=None, user_role_arn=None):
        if admin_role_arn and not isinstance(admin_role_arn, str):
            raise TypeError("Expected argument 'admin_role_arn' to be a str")
        pulumi.set(__self__, "admin_role_arn", admin_role_arn)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if home_region and not isinstance(home_region, str):
            raise TypeError("Expected argument 'home_region' to be a str")
        pulumi.set(__self__, "home_region", home_region)
        if sso_client_id and not isinstance(sso_client_id, str):
            raise TypeError("Expected argument 'sso_client_id' to be a str")
        pulumi.set(__self__, "sso_client_id", sso_client_id)
        if studio_encryption_configuration and not isinstance(studio_encryption_configuration, dict):
            raise TypeError("Expected argument 'studio_encryption_configuration' to be a dict")
        pulumi.set(__self__, "studio_encryption_configuration", studio_encryption_configuration)
        if studio_id and not isinstance(studio_id, str):
            raise TypeError("Expected argument 'studio_id' to be a str")
        pulumi.set(__self__, "studio_id", studio_id)
        if studio_url and not isinstance(studio_url, str):
            raise TypeError("Expected argument 'studio_url' to be a str")
        pulumi.set(__self__, "studio_url", studio_url)
        if user_role_arn and not isinstance(user_role_arn, str):
            raise TypeError("Expected argument 'user_role_arn' to be a str")
        pulumi.set(__self__, "user_role_arn", user_role_arn)

    @property
    @pulumi.getter(name="adminRoleArn")
    def admin_role_arn(self) -> Optional[builtins.str]:
        return pulumi.get(self, "admin_role_arn")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[builtins.str]:
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="homeRegion")
    def home_region(self) -> Optional[builtins.str]:
        return pulumi.get(self, "home_region")

    @property
    @pulumi.getter(name="ssoClientId")
    def sso_client_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "sso_client_id")

    @property
    @pulumi.getter(name="studioEncryptionConfiguration")
    def studio_encryption_configuration(self) -> Optional['outputs.StudioEncryptionConfiguration']:
        return pulumi.get(self, "studio_encryption_configuration")

    @property
    @pulumi.getter(name="studioId")
    def studio_id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "studio_id")

    @property
    @pulumi.getter(name="studioUrl")
    def studio_url(self) -> Optional[builtins.str]:
        return pulumi.get(self, "studio_url")

    @property
    @pulumi.getter(name="userRoleArn")
    def user_role_arn(self) -> Optional[builtins.str]:
        return pulumi.get(self, "user_role_arn")


class AwaitableGetStudioResult(GetStudioResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStudioResult(
            admin_role_arn=self.admin_role_arn,
            display_name=self.display_name,
            home_region=self.home_region,
            sso_client_id=self.sso_client_id,
            studio_encryption_configuration=self.studio_encryption_configuration,
            studio_id=self.studio_id,
            studio_url=self.studio_url,
            user_role_arn=self.user_role_arn)


def get_studio(studio_id: Optional[builtins.str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStudioResult:
    """
    Resource Type definition for AWS::NimbleStudio::Studio
    """
    __args__ = dict()
    __args__['studioId'] = studio_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:nimblestudio:getStudio', __args__, opts=opts, typ=GetStudioResult).value

    return AwaitableGetStudioResult(
        admin_role_arn=pulumi.get(__ret__, 'admin_role_arn'),
        display_name=pulumi.get(__ret__, 'display_name'),
        home_region=pulumi.get(__ret__, 'home_region'),
        sso_client_id=pulumi.get(__ret__, 'sso_client_id'),
        studio_encryption_configuration=pulumi.get(__ret__, 'studio_encryption_configuration'),
        studio_id=pulumi.get(__ret__, 'studio_id'),
        studio_url=pulumi.get(__ret__, 'studio_url'),
        user_role_arn=pulumi.get(__ret__, 'user_role_arn'))
def get_studio_output(studio_id: Optional[pulumi.Input[builtins.str]] = None,
                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetStudioResult]:
    """
    Resource Type definition for AWS::NimbleStudio::Studio
    """
    __args__ = dict()
    __args__['studioId'] = studio_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:nimblestudio:getStudio', __args__, opts=opts, typ=GetStudioResult)
    return __ret__.apply(lambda __response__: GetStudioResult(
        admin_role_arn=pulumi.get(__response__, 'admin_role_arn'),
        display_name=pulumi.get(__response__, 'display_name'),
        home_region=pulumi.get(__response__, 'home_region'),
        sso_client_id=pulumi.get(__response__, 'sso_client_id'),
        studio_encryption_configuration=pulumi.get(__response__, 'studio_encryption_configuration'),
        studio_id=pulumi.get(__response__, 'studio_id'),
        studio_url=pulumi.get(__response__, 'studio_url'),
        user_role_arn=pulumi.get(__response__, 'user_role_arn')))
