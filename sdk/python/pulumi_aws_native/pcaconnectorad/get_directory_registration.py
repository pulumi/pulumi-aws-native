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
    'GetDirectoryRegistrationResult',
    'AwaitableGetDirectoryRegistrationResult',
    'get_directory_registration',
    'get_directory_registration_output',
]

@pulumi.output_type
class GetDirectoryRegistrationResult:
    def __init__(__self__, directory_registration_arn=None, tags=None):
        if directory_registration_arn and not isinstance(directory_registration_arn, str):
            raise TypeError("Expected argument 'directory_registration_arn' to be a str")
        pulumi.set(__self__, "directory_registration_arn", directory_registration_arn)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="directoryRegistrationArn")
    def directory_registration_arn(self) -> Optional[builtins.str]:
        """
        The Amazon Resource Name (ARN) that was returned when you called [CreateDirectoryRegistration](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration.html) .
        """
        return pulumi.get(self, "directory_registration_arn")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Metadata assigned to a directory registration consisting of a key-value pair.
        """
        return pulumi.get(self, "tags")


class AwaitableGetDirectoryRegistrationResult(GetDirectoryRegistrationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDirectoryRegistrationResult(
            directory_registration_arn=self.directory_registration_arn,
            tags=self.tags)


def get_directory_registration(directory_registration_arn: Optional[builtins.str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDirectoryRegistrationResult:
    """
    Definition of AWS::PCAConnectorAD::DirectoryRegistration Resource Type


    :param builtins.str directory_registration_arn: The Amazon Resource Name (ARN) that was returned when you called [CreateDirectoryRegistration](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration.html) .
    """
    __args__ = dict()
    __args__['directoryRegistrationArn'] = directory_registration_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:pcaconnectorad:getDirectoryRegistration', __args__, opts=opts, typ=GetDirectoryRegistrationResult).value

    return AwaitableGetDirectoryRegistrationResult(
        directory_registration_arn=pulumi.get(__ret__, 'directory_registration_arn'),
        tags=pulumi.get(__ret__, 'tags'))
def get_directory_registration_output(directory_registration_arn: Optional[pulumi.Input[builtins.str]] = None,
                                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDirectoryRegistrationResult]:
    """
    Definition of AWS::PCAConnectorAD::DirectoryRegistration Resource Type


    :param builtins.str directory_registration_arn: The Amazon Resource Name (ARN) that was returned when you called [CreateDirectoryRegistration](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration.html) .
    """
    __args__ = dict()
    __args__['directoryRegistrationArn'] = directory_registration_arn
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:pcaconnectorad:getDirectoryRegistration', __args__, opts=opts, typ=GetDirectoryRegistrationResult)
    return __ret__.apply(lambda __response__: GetDirectoryRegistrationResult(
        directory_registration_arn=pulumi.get(__response__, 'directory_registration_arn'),
        tags=pulumi.get(__response__, 'tags')))
