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
    'SuiteDefinitionConfigurationPropertiesArgs',
    'SuiteDefinitionConfigurationPropertiesArgsDict',
    'SuiteDefinitionDeviceUnderTestArgs',
    'SuiteDefinitionDeviceUnderTestArgsDict',
]

MYPY = False

if not MYPY:
    class SuiteDefinitionConfigurationPropertiesArgsDict(TypedDict):
        """
        The configuration of the Suite Definition. Listed below are the required elements of the `SuiteDefinitionConfiguration` .

        - ***devicePermissionRoleArn*** - The device permission arn.

        This is a required element.

        *Type:* String
        - ***devices*** - The list of configured devices under test. For more information on devices under test, see [DeviceUnderTest](https://docs.aws.amazon.com/iot/latest/apireference/API_iotdeviceadvisor_DeviceUnderTest.html)

        Not a required element.

        *Type:* List of devices under test
        - ***intendedForQualification*** - The tests intended for qualification in a suite.

        Not a required element.

        *Type:* Boolean
        - ***rootGroup*** - The test suite root group. For more information on creating and using root groups see the [Device Advisor workflow](https://docs.aws.amazon.com/iot/latest/developerguide/device-advisor-workflow.html) .

        This is a required element.

        *Type:* String
        - ***suiteDefinitionName*** - The Suite Definition Configuration name.

        This is a required element.

        *Type:* String
        """
        device_permission_role_arn: pulumi.Input[builtins.str]
        """
        Gets the device permission ARN. This is a required parameter.
        """
        root_group: pulumi.Input[builtins.str]
        """
        Gets the test suite root group. This is a required parameter. For updating or creating the latest qualification suite, if `intendedForQualification` is set to true, `rootGroup` can be an empty string. If `intendedForQualification` is false, `rootGroup` cannot be an empty string. If `rootGroup` is empty, and `intendedForQualification` is set to true, all the qualification tests are included, and the configuration is default.

        For a qualification suite, the minimum length is 0, and the maximum is 2048. For a non-qualification suite, the minimum length is 1, and the maximum is 2048.
        """
        devices: NotRequired[pulumi.Input[Sequence[pulumi.Input['SuiteDefinitionDeviceUnderTestArgsDict']]]]
        """
        Gets the devices configured.
        """
        intended_for_qualification: NotRequired[pulumi.Input[builtins.bool]]
        """
        Gets the tests intended for qualification in a suite.
        """
        suite_definition_name: NotRequired[pulumi.Input[builtins.str]]
        """
        Gets the suite definition name. This is a required parameter.
        """
elif False:
    SuiteDefinitionConfigurationPropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class SuiteDefinitionConfigurationPropertiesArgs:
    def __init__(__self__, *,
                 device_permission_role_arn: pulumi.Input[builtins.str],
                 root_group: pulumi.Input[builtins.str],
                 devices: Optional[pulumi.Input[Sequence[pulumi.Input['SuiteDefinitionDeviceUnderTestArgs']]]] = None,
                 intended_for_qualification: Optional[pulumi.Input[builtins.bool]] = None,
                 suite_definition_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The configuration of the Suite Definition. Listed below are the required elements of the `SuiteDefinitionConfiguration` .

        - ***devicePermissionRoleArn*** - The device permission arn.

        This is a required element.

        *Type:* String
        - ***devices*** - The list of configured devices under test. For more information on devices under test, see [DeviceUnderTest](https://docs.aws.amazon.com/iot/latest/apireference/API_iotdeviceadvisor_DeviceUnderTest.html)

        Not a required element.

        *Type:* List of devices under test
        - ***intendedForQualification*** - The tests intended for qualification in a suite.

        Not a required element.

        *Type:* Boolean
        - ***rootGroup*** - The test suite root group. For more information on creating and using root groups see the [Device Advisor workflow](https://docs.aws.amazon.com/iot/latest/developerguide/device-advisor-workflow.html) .

        This is a required element.

        *Type:* String
        - ***suiteDefinitionName*** - The Suite Definition Configuration name.

        This is a required element.

        *Type:* String
        :param pulumi.Input[builtins.str] device_permission_role_arn: Gets the device permission ARN. This is a required parameter.
        :param pulumi.Input[builtins.str] root_group: Gets the test suite root group. This is a required parameter. For updating or creating the latest qualification suite, if `intendedForQualification` is set to true, `rootGroup` can be an empty string. If `intendedForQualification` is false, `rootGroup` cannot be an empty string. If `rootGroup` is empty, and `intendedForQualification` is set to true, all the qualification tests are included, and the configuration is default.
               
               For a qualification suite, the minimum length is 0, and the maximum is 2048. For a non-qualification suite, the minimum length is 1, and the maximum is 2048.
        :param pulumi.Input[Sequence[pulumi.Input['SuiteDefinitionDeviceUnderTestArgs']]] devices: Gets the devices configured.
        :param pulumi.Input[builtins.bool] intended_for_qualification: Gets the tests intended for qualification in a suite.
        :param pulumi.Input[builtins.str] suite_definition_name: Gets the suite definition name. This is a required parameter.
        """
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
    def device_permission_role_arn(self) -> pulumi.Input[builtins.str]:
        """
        Gets the device permission ARN. This is a required parameter.
        """
        return pulumi.get(self, "device_permission_role_arn")

    @device_permission_role_arn.setter
    def device_permission_role_arn(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "device_permission_role_arn", value)

    @property
    @pulumi.getter(name="rootGroup")
    def root_group(self) -> pulumi.Input[builtins.str]:
        """
        Gets the test suite root group. This is a required parameter. For updating or creating the latest qualification suite, if `intendedForQualification` is set to true, `rootGroup` can be an empty string. If `intendedForQualification` is false, `rootGroup` cannot be an empty string. If `rootGroup` is empty, and `intendedForQualification` is set to true, all the qualification tests are included, and the configuration is default.

        For a qualification suite, the minimum length is 0, and the maximum is 2048. For a non-qualification suite, the minimum length is 1, and the maximum is 2048.
        """
        return pulumi.get(self, "root_group")

    @root_group.setter
    def root_group(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "root_group", value)

    @property
    @pulumi.getter
    def devices(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SuiteDefinitionDeviceUnderTestArgs']]]]:
        """
        Gets the devices configured.
        """
        return pulumi.get(self, "devices")

    @devices.setter
    def devices(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SuiteDefinitionDeviceUnderTestArgs']]]]):
        pulumi.set(self, "devices", value)

    @property
    @pulumi.getter(name="intendedForQualification")
    def intended_for_qualification(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Gets the tests intended for qualification in a suite.
        """
        return pulumi.get(self, "intended_for_qualification")

    @intended_for_qualification.setter
    def intended_for_qualification(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "intended_for_qualification", value)

    @property
    @pulumi.getter(name="suiteDefinitionName")
    def suite_definition_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Gets the suite definition name. This is a required parameter.
        """
        return pulumi.get(self, "suite_definition_name")

    @suite_definition_name.setter
    def suite_definition_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "suite_definition_name", value)


if not MYPY:
    class SuiteDefinitionDeviceUnderTestArgsDict(TypedDict):
        certificate_arn: NotRequired[pulumi.Input[builtins.str]]
        thing_arn: NotRequired[pulumi.Input[builtins.str]]
elif False:
    SuiteDefinitionDeviceUnderTestArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class SuiteDefinitionDeviceUnderTestArgs:
    def __init__(__self__, *,
                 certificate_arn: Optional[pulumi.Input[builtins.str]] = None,
                 thing_arn: Optional[pulumi.Input[builtins.str]] = None):
        if certificate_arn is not None:
            pulumi.set(__self__, "certificate_arn", certificate_arn)
        if thing_arn is not None:
            pulumi.set(__self__, "thing_arn", thing_arn)

    @property
    @pulumi.getter(name="certificateArn")
    def certificate_arn(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "certificate_arn")

    @certificate_arn.setter
    def certificate_arn(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "certificate_arn", value)

    @property
    @pulumi.getter(name="thingArn")
    def thing_arn(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "thing_arn")

    @thing_arn.setter
    def thing_arn(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "thing_arn", value)


