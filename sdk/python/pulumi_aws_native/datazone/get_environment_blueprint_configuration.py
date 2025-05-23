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
    'GetEnvironmentBlueprintConfigurationResult',
    'AwaitableGetEnvironmentBlueprintConfigurationResult',
    'get_environment_blueprint_configuration',
    'get_environment_blueprint_configuration_output',
]

@pulumi.output_type
class GetEnvironmentBlueprintConfigurationResult:
    def __init__(__self__, created_at=None, domain_id=None, enabled_regions=None, environment_blueprint_id=None, manage_access_role_arn=None, provisioning_role_arn=None, regional_parameters=None, updated_at=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if domain_id and not isinstance(domain_id, str):
            raise TypeError("Expected argument 'domain_id' to be a str")
        pulumi.set(__self__, "domain_id", domain_id)
        if enabled_regions and not isinstance(enabled_regions, list):
            raise TypeError("Expected argument 'enabled_regions' to be a list")
        pulumi.set(__self__, "enabled_regions", enabled_regions)
        if environment_blueprint_id and not isinstance(environment_blueprint_id, str):
            raise TypeError("Expected argument 'environment_blueprint_id' to be a str")
        pulumi.set(__self__, "environment_blueprint_id", environment_blueprint_id)
        if manage_access_role_arn and not isinstance(manage_access_role_arn, str):
            raise TypeError("Expected argument 'manage_access_role_arn' to be a str")
        pulumi.set(__self__, "manage_access_role_arn", manage_access_role_arn)
        if provisioning_role_arn and not isinstance(provisioning_role_arn, str):
            raise TypeError("Expected argument 'provisioning_role_arn' to be a str")
        pulumi.set(__self__, "provisioning_role_arn", provisioning_role_arn)
        if regional_parameters and not isinstance(regional_parameters, list):
            raise TypeError("Expected argument 'regional_parameters' to be a list")
        pulumi.set(__self__, "regional_parameters", regional_parameters)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[builtins.str]:
        """
        The timestamp of when an environment blueprint was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> Optional[builtins.str]:
        """
        The identifier of the Amazon DataZone domain in which an environment blueprint exists.
        """
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter(name="enabledRegions")
    def enabled_regions(self) -> Optional[Sequence[builtins.str]]:
        """
        The enabled AWS Regions specified in a blueprint configuration.
        """
        return pulumi.get(self, "enabled_regions")

    @property
    @pulumi.getter(name="environmentBlueprintId")
    def environment_blueprint_id(self) -> Optional[builtins.str]:
        """
        The identifier of the environment blueprint. This identifier should be used when creating environment profiles.
        """
        return pulumi.get(self, "environment_blueprint_id")

    @property
    @pulumi.getter(name="manageAccessRoleArn")
    def manage_access_role_arn(self) -> Optional[builtins.str]:
        """
        The ARN of the manage access role.
        """
        return pulumi.get(self, "manage_access_role_arn")

    @property
    @pulumi.getter(name="provisioningRoleArn")
    def provisioning_role_arn(self) -> Optional[builtins.str]:
        """
        The ARN of the provisioning role.
        """
        return pulumi.get(self, "provisioning_role_arn")

    @property
    @pulumi.getter(name="regionalParameters")
    def regional_parameters(self) -> Optional[Sequence['outputs.EnvironmentBlueprintConfigurationRegionalParameter']]:
        """
        The regional parameters of the environment blueprint.
        """
        return pulumi.get(self, "regional_parameters")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[builtins.str]:
        """
        The timestamp of when the environment blueprint was updated.
        """
        return pulumi.get(self, "updated_at")


class AwaitableGetEnvironmentBlueprintConfigurationResult(GetEnvironmentBlueprintConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEnvironmentBlueprintConfigurationResult(
            created_at=self.created_at,
            domain_id=self.domain_id,
            enabled_regions=self.enabled_regions,
            environment_blueprint_id=self.environment_blueprint_id,
            manage_access_role_arn=self.manage_access_role_arn,
            provisioning_role_arn=self.provisioning_role_arn,
            regional_parameters=self.regional_parameters,
            updated_at=self.updated_at)


def get_environment_blueprint_configuration(domain_id: Optional[builtins.str] = None,
                                            environment_blueprint_id: Optional[builtins.str] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEnvironmentBlueprintConfigurationResult:
    """
    Definition of AWS::DataZone::EnvironmentBlueprintConfiguration Resource Type


    :param builtins.str domain_id: The identifier of the Amazon DataZone domain in which an environment blueprint exists.
    :param builtins.str environment_blueprint_id: The identifier of the environment blueprint. This identifier should be used when creating environment profiles.
    """
    __args__ = dict()
    __args__['domainId'] = domain_id
    __args__['environmentBlueprintId'] = environment_blueprint_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:datazone:getEnvironmentBlueprintConfiguration', __args__, opts=opts, typ=GetEnvironmentBlueprintConfigurationResult).value

    return AwaitableGetEnvironmentBlueprintConfigurationResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        domain_id=pulumi.get(__ret__, 'domain_id'),
        enabled_regions=pulumi.get(__ret__, 'enabled_regions'),
        environment_blueprint_id=pulumi.get(__ret__, 'environment_blueprint_id'),
        manage_access_role_arn=pulumi.get(__ret__, 'manage_access_role_arn'),
        provisioning_role_arn=pulumi.get(__ret__, 'provisioning_role_arn'),
        regional_parameters=pulumi.get(__ret__, 'regional_parameters'),
        updated_at=pulumi.get(__ret__, 'updated_at'))
def get_environment_blueprint_configuration_output(domain_id: Optional[pulumi.Input[builtins.str]] = None,
                                                   environment_blueprint_id: Optional[pulumi.Input[builtins.str]] = None,
                                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetEnvironmentBlueprintConfigurationResult]:
    """
    Definition of AWS::DataZone::EnvironmentBlueprintConfiguration Resource Type


    :param builtins.str domain_id: The identifier of the Amazon DataZone domain in which an environment blueprint exists.
    :param builtins.str environment_blueprint_id: The identifier of the environment blueprint. This identifier should be used when creating environment profiles.
    """
    __args__ = dict()
    __args__['domainId'] = domain_id
    __args__['environmentBlueprintId'] = environment_blueprint_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:datazone:getEnvironmentBlueprintConfiguration', __args__, opts=opts, typ=GetEnvironmentBlueprintConfigurationResult)
    return __ret__.apply(lambda __response__: GetEnvironmentBlueprintConfigurationResult(
        created_at=pulumi.get(__response__, 'created_at'),
        domain_id=pulumi.get(__response__, 'domain_id'),
        enabled_regions=pulumi.get(__response__, 'enabled_regions'),
        environment_blueprint_id=pulumi.get(__response__, 'environment_blueprint_id'),
        manage_access_role_arn=pulumi.get(__response__, 'manage_access_role_arn'),
        provisioning_role_arn=pulumi.get(__response__, 'provisioning_role_arn'),
        regional_parameters=pulumi.get(__response__, 'regional_parameters'),
        updated_at=pulumi.get(__response__, 'updated_at')))
