# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetApplicationResult',
    'AwaitableGetApplicationResult',
    'get_application',
    'get_application_output',
]

@pulumi.output_type
class GetApplicationResult:
    def __init__(__self__, application_arn=None, application_creation_date=None, application_description=None, application_id=None, application_last_update_date=None, application_name=None, application_state=None, application_url=None, error_message=None, role_arn=None, sso_client_id=None, tags=None):
        if application_arn and not isinstance(application_arn, str):
            raise TypeError("Expected argument 'application_arn' to be a str")
        pulumi.set(__self__, "application_arn", application_arn)
        if application_creation_date and not isinstance(application_creation_date, int):
            raise TypeError("Expected argument 'application_creation_date' to be a int")
        pulumi.set(__self__, "application_creation_date", application_creation_date)
        if application_description and not isinstance(application_description, str):
            raise TypeError("Expected argument 'application_description' to be a str")
        pulumi.set(__self__, "application_description", application_description)
        if application_id and not isinstance(application_id, str):
            raise TypeError("Expected argument 'application_id' to be a str")
        pulumi.set(__self__, "application_id", application_id)
        if application_last_update_date and not isinstance(application_last_update_date, int):
            raise TypeError("Expected argument 'application_last_update_date' to be a int")
        pulumi.set(__self__, "application_last_update_date", application_last_update_date)
        if application_name and not isinstance(application_name, str):
            raise TypeError("Expected argument 'application_name' to be a str")
        pulumi.set(__self__, "application_name", application_name)
        if application_state and not isinstance(application_state, str):
            raise TypeError("Expected argument 'application_state' to be a str")
        pulumi.set(__self__, "application_state", application_state)
        if application_url and not isinstance(application_url, str):
            raise TypeError("Expected argument 'application_url' to be a str")
        pulumi.set(__self__, "application_url", application_url)
        if error_message and not isinstance(error_message, str):
            raise TypeError("Expected argument 'error_message' to be a str")
        pulumi.set(__self__, "error_message", error_message)
        if role_arn and not isinstance(role_arn, str):
            raise TypeError("Expected argument 'role_arn' to be a str")
        pulumi.set(__self__, "role_arn", role_arn)
        if sso_client_id and not isinstance(sso_client_id, str):
            raise TypeError("Expected argument 'sso_client_id' to be a str")
        pulumi.set(__self__, "sso_client_id", sso_client_id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="applicationArn")
    def application_arn(self) -> Optional[str]:
        """
        The ARN of the application.
        """
        return pulumi.get(self, "application_arn")

    @property
    @pulumi.getter(name="applicationCreationDate")
    def application_creation_date(self) -> Optional[int]:
        """
        When the Application was created
        """
        return pulumi.get(self, "application_creation_date")

    @property
    @pulumi.getter(name="applicationDescription")
    def application_description(self) -> Optional[str]:
        """
        Application Description, should be between 1 and 2048 characters.
        """
        return pulumi.get(self, "application_description")

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[str]:
        """
        The ID of the application.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="applicationLastUpdateDate")
    def application_last_update_date(self) -> Optional[int]:
        """
        When the Application was last updated
        """
        return pulumi.get(self, "application_last_update_date")

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> Optional[str]:
        """
        Application Name, should be between 1 and 256 characters.
        """
        return pulumi.get(self, "application_name")

    @property
    @pulumi.getter(name="applicationState")
    def application_state(self) -> Optional[str]:
        """
        The current state of the application.
        """
        return pulumi.get(self, "application_state")

    @property
    @pulumi.getter(name="applicationUrl")
    def application_url(self) -> Optional[str]:
        """
        The URL of the application.
        """
        return pulumi.get(self, "application_url")

    @property
    @pulumi.getter(name="errorMessage")
    def error_message(self) -> Optional[str]:
        """
        A message indicating why Create or Delete Application failed.
        """
        return pulumi.get(self, "error_message")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[str]:
        """
        The ARN of the role that the web application assumes when it interacts with AWS IoT Core. For more info on configuring this attribute, see https://docs.aws.amazon.com/iot/latest/apireference/API_iotfleethub_CreateApplication.html#API_iotfleethub_CreateApplication_RequestSyntax
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="ssoClientId")
    def sso_client_id(self) -> Optional[str]:
        """
        The AWS SSO application generated client ID (used with AWS SSO APIs).
        """
        return pulumi.get(self, "sso_client_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.ApplicationTag']]:
        """
        A list of key-value pairs that contain metadata for the application.
        """
        return pulumi.get(self, "tags")


class AwaitableGetApplicationResult(GetApplicationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationResult(
            application_arn=self.application_arn,
            application_creation_date=self.application_creation_date,
            application_description=self.application_description,
            application_id=self.application_id,
            application_last_update_date=self.application_last_update_date,
            application_name=self.application_name,
            application_state=self.application_state,
            application_url=self.application_url,
            error_message=self.error_message,
            role_arn=self.role_arn,
            sso_client_id=self.sso_client_id,
            tags=self.tags)


def get_application(application_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationResult:
    """
    Resource schema for AWS::IoTFleetHub::Application


    :param str application_id: The ID of the application.
    """
    __args__ = dict()
    __args__['applicationId'] = application_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:iotfleethub:getApplication', __args__, opts=opts, typ=GetApplicationResult).value

    return AwaitableGetApplicationResult(
        application_arn=__ret__.application_arn,
        application_creation_date=__ret__.application_creation_date,
        application_description=__ret__.application_description,
        application_id=__ret__.application_id,
        application_last_update_date=__ret__.application_last_update_date,
        application_name=__ret__.application_name,
        application_state=__ret__.application_state,
        application_url=__ret__.application_url,
        error_message=__ret__.error_message,
        role_arn=__ret__.role_arn,
        sso_client_id=__ret__.sso_client_id,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_application)
def get_application_output(application_id: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApplicationResult]:
    """
    Resource schema for AWS::IoTFleetHub::Application


    :param str application_id: The ID of the application.
    """
    ...