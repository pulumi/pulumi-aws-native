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
    'SuiteDefinitionConfigurationPropertiesArgs',
    'SuiteDefinitionDeviceUnderTestArgs',
    'SuiteDefinitionTagArgs',
]

@pulumi.input_type
class SuiteDefinitionConfigurationPropertiesArgs:
    def __init__(__self__, *,
                 device_permission_role_arn: pulumi.Input[str],
                 root_group: pulumi.Input[str],
                 devices: Optional[pulumi.Input[Sequence[pulumi.Input['SuiteDefinitionDeviceUnderTestArgs']]]] = None,
                 intended_for_qualification: Optional[pulumi.Input[bool]] = None,
                 suite_definition_name: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "device_permission_role_arn", device_permission_role_arn)
        pulumi.set(__self__, "root_group", root_group)
        if devices is not None:
            pulumi.set(__self__, "devices", devices)
        if intended_for_qualification is not None:
            pulumi.set(__self__, "intended_for_qualification", intended_for_qualification)
        if suite_definition_name is not None:
            pulumi.set(__self__, "suite_definition_name", suite_definition_name)

    @property
    @pulumi.getter(name="devicePermissionRoleArn")
    def device_permission_role_arn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "device_permission_role_arn")

    @device_permission_role_arn.setter
    def device_permission_role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "device_permission_role_arn", value)

    @property
    @pulumi.getter(name="rootGroup")
    def root_group(self) -> pulumi.Input[str]:
        return pulumi.get(self, "root_group")

    @root_group.setter
    def root_group(self, value: pulumi.Input[str]):
        pulumi.set(self, "root_group", value)

    @property
    @pulumi.getter
    def devices(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SuiteDefinitionDeviceUnderTestArgs']]]]:
        return pulumi.get(self, "devices")

    @devices.setter
    def devices(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SuiteDefinitionDeviceUnderTestArgs']]]]):
        pulumi.set(self, "devices", value)

    @property
    @pulumi.getter(name="intendedForQualification")
    def intended_for_qualification(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "intended_for_qualification")

    @intended_for_qualification.setter
    def intended_for_qualification(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "intended_for_qualification", value)

    @property
    @pulumi.getter(name="suiteDefinitionName")
    def suite_definition_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "suite_definition_name")

    @suite_definition_name.setter
    def suite_definition_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "suite_definition_name", value)


@pulumi.input_type
class SuiteDefinitionDeviceUnderTestArgs:
    def __init__(__self__, *,
                 certificate_arn: Optional[pulumi.Input[str]] = None,
                 thing_arn: Optional[pulumi.Input[str]] = None):
        if certificate_arn is not None:
            pulumi.set(__self__, "certificate_arn", certificate_arn)
        if thing_arn is not None:
            pulumi.set(__self__, "thing_arn", thing_arn)

    @property
    @pulumi.getter(name="certificateArn")
    def certificate_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "certificate_arn")

    @certificate_arn.setter
    def certificate_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_arn", value)

    @property
    @pulumi.getter(name="thingArn")
    def thing_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "thing_arn")

    @thing_arn.setter
    def thing_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "thing_arn", value)


@pulumi.input_type
class SuiteDefinitionTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A key-value pair to associate with a resource.
        :param pulumi.Input[str] key: The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        :param pulumi.Input[str] value: The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

