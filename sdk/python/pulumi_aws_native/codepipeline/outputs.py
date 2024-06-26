# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'CustomActionTypeArtifactDetails',
    'CustomActionTypeConfigurationProperties',
    'CustomActionTypeSettings',
]

@pulumi.output_type
class CustomActionTypeArtifactDetails(dict):
    """
    Returns information about the details of an artifact.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "maximumCount":
            suggest = "maximum_count"
        elif key == "minimumCount":
            suggest = "minimum_count"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CustomActionTypeArtifactDetails. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CustomActionTypeArtifactDetails.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CustomActionTypeArtifactDetails.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 maximum_count: int,
                 minimum_count: int):
        """
        Returns information about the details of an artifact.
        :param int maximum_count: The maximum number of artifacts allowed for the action type.
        :param int minimum_count: The minimum number of artifacts allowed for the action type.
        """
        pulumi.set(__self__, "maximum_count", maximum_count)
        pulumi.set(__self__, "minimum_count", minimum_count)

    @property
    @pulumi.getter(name="maximumCount")
    def maximum_count(self) -> int:
        """
        The maximum number of artifacts allowed for the action type.
        """
        return pulumi.get(self, "maximum_count")

    @property
    @pulumi.getter(name="minimumCount")
    def minimum_count(self) -> int:
        """
        The minimum number of artifacts allowed for the action type.
        """
        return pulumi.get(self, "minimum_count")


@pulumi.output_type
class CustomActionTypeConfigurationProperties(dict):
    """
    The configuration properties for the custom action.
    """
    def __init__(__self__, *,
                 key: bool,
                 name: str,
                 required: bool,
                 secret: bool,
                 description: Optional[str] = None,
                 queryable: Optional[bool] = None,
                 type: Optional[str] = None):
        """
        The configuration properties for the custom action.
        :param bool key: Whether the configuration property is a key.
        :param str name: The name of the action configuration property.
        :param bool required: Whether the configuration property is a required value.
        :param bool secret: Whether the configuration property is secret. Secrets are hidden from all calls except for GetJobDetails, GetThirdPartyJobDetails, PollForJobs, and PollForThirdPartyJobs.
        :param str description: The description of the action configuration property that is displayed to users. 
        :param bool queryable: Indicates that the property is used with PollForJobs. When creating a custom action, an action can have up to one queryable property. If it has one, that property must be both required and not secret.If you create a pipeline with a custom action type, and that custom action contains a queryable property, the value for that configuration property is subject to other restrictions. The value must be less than or equal to twenty (20) characters. The value can contain only alphanumeric characters, underscores, and hyphens. 
        :param str type: The type of the configuration property.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "required", required)
        pulumi.set(__self__, "secret", secret)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if queryable is not None:
            pulumi.set(__self__, "queryable", queryable)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def key(self) -> bool:
        """
        Whether the configuration property is a key.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the action configuration property.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def required(self) -> bool:
        """
        Whether the configuration property is a required value.
        """
        return pulumi.get(self, "required")

    @property
    @pulumi.getter
    def secret(self) -> bool:
        """
        Whether the configuration property is secret. Secrets are hidden from all calls except for GetJobDetails, GetThirdPartyJobDetails, PollForJobs, and PollForThirdPartyJobs.
        """
        return pulumi.get(self, "secret")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the action configuration property that is displayed to users. 
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def queryable(self) -> Optional[bool]:
        """
        Indicates that the property is used with PollForJobs. When creating a custom action, an action can have up to one queryable property. If it has one, that property must be both required and not secret.If you create a pipeline with a custom action type, and that custom action contains a queryable property, the value for that configuration property is subject to other restrictions. The value must be less than or equal to twenty (20) characters. The value can contain only alphanumeric characters, underscores, and hyphens. 
        """
        return pulumi.get(self, "queryable")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of the configuration property.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class CustomActionTypeSettings(dict):
    """
    Settings is a property of the AWS::CodePipeline::CustomActionType resource that provides URLs that users can access to view information about the CodePipeline custom action. 
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "entityUrlTemplate":
            suggest = "entity_url_template"
        elif key == "executionUrlTemplate":
            suggest = "execution_url_template"
        elif key == "revisionUrlTemplate":
            suggest = "revision_url_template"
        elif key == "thirdPartyConfigurationUrl":
            suggest = "third_party_configuration_url"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CustomActionTypeSettings. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CustomActionTypeSettings.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CustomActionTypeSettings.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 entity_url_template: Optional[str] = None,
                 execution_url_template: Optional[str] = None,
                 revision_url_template: Optional[str] = None,
                 third_party_configuration_url: Optional[str] = None):
        """
        Settings is a property of the AWS::CodePipeline::CustomActionType resource that provides URLs that users can access to view information about the CodePipeline custom action. 
        :param str entity_url_template: The URL returned to the AWS CodePipeline console that provides a deep link to the resources of the external system, such as the configuration page for an AWS CodeDeploy deployment group. This link is provided as part of the action display in the pipeline. 
        :param str execution_url_template: The URL returned to the AWS CodePipeline console that contains a link to the top-level landing page for the external system, such as the console page for AWS CodeDeploy. This link is shown on the pipeline view page in the AWS CodePipeline console and provides a link to the execution entity of the external action. 
        :param str revision_url_template: The URL returned to the AWS CodePipeline console that contains a link to the page where customers can update or change the configuration of the external action. 
        :param str third_party_configuration_url: The URL of a sign-up page where users can sign up for an external service and perform initial configuration of the action provided by that service.
        """
        if entity_url_template is not None:
            pulumi.set(__self__, "entity_url_template", entity_url_template)
        if execution_url_template is not None:
            pulumi.set(__self__, "execution_url_template", execution_url_template)
        if revision_url_template is not None:
            pulumi.set(__self__, "revision_url_template", revision_url_template)
        if third_party_configuration_url is not None:
            pulumi.set(__self__, "third_party_configuration_url", third_party_configuration_url)

    @property
    @pulumi.getter(name="entityUrlTemplate")
    def entity_url_template(self) -> Optional[str]:
        """
        The URL returned to the AWS CodePipeline console that provides a deep link to the resources of the external system, such as the configuration page for an AWS CodeDeploy deployment group. This link is provided as part of the action display in the pipeline. 
        """
        return pulumi.get(self, "entity_url_template")

    @property
    @pulumi.getter(name="executionUrlTemplate")
    def execution_url_template(self) -> Optional[str]:
        """
        The URL returned to the AWS CodePipeline console that contains a link to the top-level landing page for the external system, such as the console page for AWS CodeDeploy. This link is shown on the pipeline view page in the AWS CodePipeline console and provides a link to the execution entity of the external action. 
        """
        return pulumi.get(self, "execution_url_template")

    @property
    @pulumi.getter(name="revisionUrlTemplate")
    def revision_url_template(self) -> Optional[str]:
        """
        The URL returned to the AWS CodePipeline console that contains a link to the page where customers can update or change the configuration of the external action. 
        """
        return pulumi.get(self, "revision_url_template")

    @property
    @pulumi.getter(name="thirdPartyConfigurationUrl")
    def third_party_configuration_url(self) -> Optional[str]:
        """
        The URL of a sign-up page where users can sign up for an external service and perform initial configuration of the action provided by that service.
        """
        return pulumi.get(self, "third_party_configuration_url")


