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
    'GetTargetAccountConfigurationResult',
    'AwaitableGetTargetAccountConfigurationResult',
    'get_target_account_configuration',
    'get_target_account_configuration_output',
]

@pulumi.output_type
class GetTargetAccountConfigurationResult:
    def __init__(__self__, description=None, role_arn=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if role_arn and not isinstance(role_arn, str):
            raise TypeError("Expected argument 'role_arn' to be a str")
        pulumi.set(__self__, "role_arn", role_arn)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[str]:
        return pulumi.get(self, "role_arn")


class AwaitableGetTargetAccountConfigurationResult(GetTargetAccountConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTargetAccountConfigurationResult(
            description=self.description,
            role_arn=self.role_arn)


def get_target_account_configuration(account_id: Optional[str] = None,
                                     experiment_template_id: Optional[str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTargetAccountConfigurationResult:
    """
    Resource schema for AWS::FIS::TargetAccountConfiguration
    """
    __args__ = dict()
    __args__['accountId'] = account_id
    __args__['experimentTemplateId'] = experiment_template_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:fis:getTargetAccountConfiguration', __args__, opts=opts, typ=GetTargetAccountConfigurationResult).value

    return AwaitableGetTargetAccountConfigurationResult(
        description=pulumi.get(__ret__, 'description'),
        role_arn=pulumi.get(__ret__, 'role_arn'))


@_utilities.lift_output_func(get_target_account_configuration)
def get_target_account_configuration_output(account_id: Optional[pulumi.Input[str]] = None,
                                            experiment_template_id: Optional[pulumi.Input[str]] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTargetAccountConfigurationResult]:
    """
    Resource schema for AWS::FIS::TargetAccountConfiguration
    """
    ...