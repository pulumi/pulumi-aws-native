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
from ._enums import *

__all__ = [
    'GetApplicationInstanceResult',
    'AwaitableGetApplicationInstanceResult',
    'get_application_instance',
    'get_application_instance_output',
]

@pulumi.output_type
class GetApplicationInstanceResult:
    def __init__(__self__, application_instance_id=None, arn=None, created_time=None, default_runtime_context_device_name=None, health_status=None, last_updated_time=None, status=None, status_description=None, tags=None):
        if application_instance_id and not isinstance(application_instance_id, str):
            raise TypeError("Expected argument 'application_instance_id' to be a str")
        pulumi.set(__self__, "application_instance_id", application_instance_id)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if created_time and not isinstance(created_time, int):
            raise TypeError("Expected argument 'created_time' to be a int")
        pulumi.set(__self__, "created_time", created_time)
        if default_runtime_context_device_name and not isinstance(default_runtime_context_device_name, str):
            raise TypeError("Expected argument 'default_runtime_context_device_name' to be a str")
        pulumi.set(__self__, "default_runtime_context_device_name", default_runtime_context_device_name)
        if health_status and not isinstance(health_status, str):
            raise TypeError("Expected argument 'health_status' to be a str")
        pulumi.set(__self__, "health_status", health_status)
        if last_updated_time and not isinstance(last_updated_time, int):
            raise TypeError("Expected argument 'last_updated_time' to be a int")
        pulumi.set(__self__, "last_updated_time", last_updated_time)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if status_description and not isinstance(status_description, str):
            raise TypeError("Expected argument 'status_description' to be a str")
        pulumi.set(__self__, "status_description", status_description)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="applicationInstanceId")
    def application_instance_id(self) -> Optional[builtins.str]:
        """
        The application instance's ID.
        """
        return pulumi.get(self, "application_instance_id")

    @property
    @pulumi.getter
    def arn(self) -> Optional[builtins.str]:
        """
        The application instance's ARN.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> Optional[builtins.int]:
        """
        The application instance's created time.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter(name="defaultRuntimeContextDeviceName")
    def default_runtime_context_device_name(self) -> Optional[builtins.str]:
        """
        The application instance's default runtime context device name.
        """
        return pulumi.get(self, "default_runtime_context_device_name")

    @property
    @pulumi.getter(name="healthStatus")
    def health_status(self) -> Optional['ApplicationInstanceHealthStatus']:
        """
        The application instance's health status.
        """
        return pulumi.get(self, "health_status")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> Optional[builtins.int]:
        """
        The application instance's last updated time.
        """
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter
    def status(self) -> Optional['ApplicationInstanceStatus']:
        """
        The application instance's status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusDescription")
    def status_description(self) -> Optional[builtins.str]:
        """
        The application instance's status description.
        """
        return pulumi.get(self, "status_description")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        Tags for the application instance.
        """
        return pulumi.get(self, "tags")


class AwaitableGetApplicationInstanceResult(GetApplicationInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationInstanceResult(
            application_instance_id=self.application_instance_id,
            arn=self.arn,
            created_time=self.created_time,
            default_runtime_context_device_name=self.default_runtime_context_device_name,
            health_status=self.health_status,
            last_updated_time=self.last_updated_time,
            status=self.status,
            status_description=self.status_description,
            tags=self.tags)


def get_application_instance(application_instance_id: Optional[builtins.str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationInstanceResult:
    """
    Creates an application instance and deploys it to a device.


    :param builtins.str application_instance_id: The application instance's ID.
    """
    __args__ = dict()
    __args__['applicationInstanceId'] = application_instance_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:panorama:getApplicationInstance', __args__, opts=opts, typ=GetApplicationInstanceResult).value

    return AwaitableGetApplicationInstanceResult(
        application_instance_id=pulumi.get(__ret__, 'application_instance_id'),
        arn=pulumi.get(__ret__, 'arn'),
        created_time=pulumi.get(__ret__, 'created_time'),
        default_runtime_context_device_name=pulumi.get(__ret__, 'default_runtime_context_device_name'),
        health_status=pulumi.get(__ret__, 'health_status'),
        last_updated_time=pulumi.get(__ret__, 'last_updated_time'),
        status=pulumi.get(__ret__, 'status'),
        status_description=pulumi.get(__ret__, 'status_description'),
        tags=pulumi.get(__ret__, 'tags'))
def get_application_instance_output(application_instance_id: Optional[pulumi.Input[builtins.str]] = None,
                                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetApplicationInstanceResult]:
    """
    Creates an application instance and deploys it to a device.


    :param builtins.str application_instance_id: The application instance's ID.
    """
    __args__ = dict()
    __args__['applicationInstanceId'] = application_instance_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:panorama:getApplicationInstance', __args__, opts=opts, typ=GetApplicationInstanceResult)
    return __ret__.apply(lambda __response__: GetApplicationInstanceResult(
        application_instance_id=pulumi.get(__response__, 'application_instance_id'),
        arn=pulumi.get(__response__, 'arn'),
        created_time=pulumi.get(__response__, 'created_time'),
        default_runtime_context_device_name=pulumi.get(__response__, 'default_runtime_context_device_name'),
        health_status=pulumi.get(__response__, 'health_status'),
        last_updated_time=pulumi.get(__response__, 'last_updated_time'),
        status=pulumi.get(__response__, 'status'),
        status_description=pulumi.get(__response__, 'status_description'),
        tags=pulumi.get(__response__, 'tags')))
