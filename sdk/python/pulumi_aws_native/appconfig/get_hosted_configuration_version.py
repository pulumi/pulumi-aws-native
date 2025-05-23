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
    'GetHostedConfigurationVersionResult',
    'AwaitableGetHostedConfigurationVersionResult',
    'get_hosted_configuration_version',
    'get_hosted_configuration_version_output',
]

@pulumi.output_type
class GetHostedConfigurationVersionResult:
    def __init__(__self__, version_number=None):
        if version_number and not isinstance(version_number, str):
            raise TypeError("Expected argument 'version_number' to be a str")
        pulumi.set(__self__, "version_number", version_number)

    @property
    @pulumi.getter(name="versionNumber")
    def version_number(self) -> Optional[builtins.str]:
        """
        Current version number of hosted configuration version.
        """
        return pulumi.get(self, "version_number")


class AwaitableGetHostedConfigurationVersionResult(GetHostedConfigurationVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHostedConfigurationVersionResult(
            version_number=self.version_number)


def get_hosted_configuration_version(application_id: Optional[builtins.str] = None,
                                     configuration_profile_id: Optional[builtins.str] = None,
                                     version_number: Optional[builtins.str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHostedConfigurationVersionResult:
    """
    Resource Type definition for AWS::AppConfig::HostedConfigurationVersion


    :param builtins.str application_id: The application ID.
    :param builtins.str configuration_profile_id: The configuration profile ID.
    :param builtins.str version_number: Current version number of hosted configuration version.
    """
    __args__ = dict()
    __args__['applicationId'] = application_id
    __args__['configurationProfileId'] = configuration_profile_id
    __args__['versionNumber'] = version_number
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:appconfig:getHostedConfigurationVersion', __args__, opts=opts, typ=GetHostedConfigurationVersionResult).value

    return AwaitableGetHostedConfigurationVersionResult(
        version_number=pulumi.get(__ret__, 'version_number'))
def get_hosted_configuration_version_output(application_id: Optional[pulumi.Input[builtins.str]] = None,
                                            configuration_profile_id: Optional[pulumi.Input[builtins.str]] = None,
                                            version_number: Optional[pulumi.Input[builtins.str]] = None,
                                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetHostedConfigurationVersionResult]:
    """
    Resource Type definition for AWS::AppConfig::HostedConfigurationVersion


    :param builtins.str application_id: The application ID.
    :param builtins.str configuration_profile_id: The configuration profile ID.
    :param builtins.str version_number: Current version number of hosted configuration version.
    """
    __args__ = dict()
    __args__['applicationId'] = application_id
    __args__['configurationProfileId'] = configuration_profile_id
    __args__['versionNumber'] = version_number
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:appconfig:getHostedConfigurationVersion', __args__, opts=opts, typ=GetHostedConfigurationVersionResult)
    return __ret__.apply(lambda __response__: GetHostedConfigurationVersionResult(
        version_number=pulumi.get(__response__, 'version_number')))
