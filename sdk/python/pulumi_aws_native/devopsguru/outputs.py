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
from ._enums import *

__all__ = [
    'NotificationChannelConfig',
    'NotificationChannelNotificationFilterConfig',
    'NotificationChannelSnsChannelConfig',
    'ResourceCollectionCloudFormationCollectionFilter',
    'ResourceCollectionFilter',
    'ResourceCollectionTagCollection',
]

@pulumi.output_type
class NotificationChannelConfig(dict):
    """
    Information about notification channels you have configured with DevOps Guru.
    """
    def __init__(__self__, *,
                 filters: Optional['outputs.NotificationChannelNotificationFilterConfig'] = None,
                 sns: Optional['outputs.NotificationChannelSnsChannelConfig'] = None):
        """
        Information about notification channels you have configured with DevOps Guru.
        """
        if filters is not None:
            pulumi.set(__self__, "filters", filters)
        if sns is not None:
            pulumi.set(__self__, "sns", sns)

    @property
    @pulumi.getter
    def filters(self) -> Optional['outputs.NotificationChannelNotificationFilterConfig']:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def sns(self) -> Optional['outputs.NotificationChannelSnsChannelConfig']:
        return pulumi.get(self, "sns")


@pulumi.output_type
class NotificationChannelNotificationFilterConfig(dict):
    """
    Information about filters of a notification channel configured in DevOpsGuru to filter for insights.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "messageTypes":
            suggest = "message_types"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NotificationChannelNotificationFilterConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NotificationChannelNotificationFilterConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NotificationChannelNotificationFilterConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 message_types: Optional[Sequence['NotificationChannelNotificationMessageType']] = None,
                 severities: Optional[Sequence['NotificationChannelInsightSeverity']] = None):
        """
        Information about filters of a notification channel configured in DevOpsGuru to filter for insights.
        """
        if message_types is not None:
            pulumi.set(__self__, "message_types", message_types)
        if severities is not None:
            pulumi.set(__self__, "severities", severities)

    @property
    @pulumi.getter(name="messageTypes")
    def message_types(self) -> Optional[Sequence['NotificationChannelNotificationMessageType']]:
        return pulumi.get(self, "message_types")

    @property
    @pulumi.getter
    def severities(self) -> Optional[Sequence['NotificationChannelInsightSeverity']]:
        return pulumi.get(self, "severities")


@pulumi.output_type
class NotificationChannelSnsChannelConfig(dict):
    """
    Information about a notification channel configured in DevOps Guru to send notifications when insights are created.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "topicArn":
            suggest = "topic_arn"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NotificationChannelSnsChannelConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NotificationChannelSnsChannelConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NotificationChannelSnsChannelConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 topic_arn: Optional[str] = None):
        """
        Information about a notification channel configured in DevOps Guru to send notifications when insights are created.
        """
        if topic_arn is not None:
            pulumi.set(__self__, "topic_arn", topic_arn)

    @property
    @pulumi.getter(name="topicArn")
    def topic_arn(self) -> Optional[str]:
        return pulumi.get(self, "topic_arn")


@pulumi.output_type
class ResourceCollectionCloudFormationCollectionFilter(dict):
    """
    CloudFormation resource for DevOps Guru to monitor
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "stackNames":
            suggest = "stack_names"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ResourceCollectionCloudFormationCollectionFilter. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ResourceCollectionCloudFormationCollectionFilter.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ResourceCollectionCloudFormationCollectionFilter.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 stack_names: Optional[Sequence[str]] = None):
        """
        CloudFormation resource for DevOps Guru to monitor
        :param Sequence[str] stack_names: An array of CloudFormation stack names.
        """
        if stack_names is not None:
            pulumi.set(__self__, "stack_names", stack_names)

    @property
    @pulumi.getter(name="stackNames")
    def stack_names(self) -> Optional[Sequence[str]]:
        """
        An array of CloudFormation stack names.
        """
        return pulumi.get(self, "stack_names")


@pulumi.output_type
class ResourceCollectionFilter(dict):
    """
    Information about a filter used to specify which AWS resources are analyzed for anomalous behavior by DevOps Guru.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "cloudFormation":
            suggest = "cloud_formation"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ResourceCollectionFilter. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ResourceCollectionFilter.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ResourceCollectionFilter.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cloud_formation: Optional['outputs.ResourceCollectionCloudFormationCollectionFilter'] = None,
                 tags: Optional[Sequence['outputs.ResourceCollectionTagCollection']] = None):
        """
        Information about a filter used to specify which AWS resources are analyzed for anomalous behavior by DevOps Guru.
        """
        if cloud_formation is not None:
            pulumi.set(__self__, "cloud_formation", cloud_formation)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="cloudFormation")
    def cloud_formation(self) -> Optional['outputs.ResourceCollectionCloudFormationCollectionFilter']:
        return pulumi.get(self, "cloud_formation")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.ResourceCollectionTagCollection']]:
        return pulumi.get(self, "tags")


@pulumi.output_type
class ResourceCollectionTagCollection(dict):
    """
    Tagged resource for DevOps Guru to monitor
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "appBoundaryKey":
            suggest = "app_boundary_key"
        elif key == "tagValues":
            suggest = "tag_values"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ResourceCollectionTagCollection. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ResourceCollectionTagCollection.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ResourceCollectionTagCollection.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 app_boundary_key: Optional[str] = None,
                 tag_values: Optional[Sequence[str]] = None):
        """
        Tagged resource for DevOps Guru to monitor
        :param str app_boundary_key: A Tag key for DevOps Guru app boundary.
        :param Sequence[str] tag_values: Tag values of DevOps Guru app boundary.
        """
        if app_boundary_key is not None:
            pulumi.set(__self__, "app_boundary_key", app_boundary_key)
        if tag_values is not None:
            pulumi.set(__self__, "tag_values", tag_values)

    @property
    @pulumi.getter(name="appBoundaryKey")
    def app_boundary_key(self) -> Optional[str]:
        """
        A Tag key for DevOps Guru app boundary.
        """
        return pulumi.get(self, "app_boundary_key")

    @property
    @pulumi.getter(name="tagValues")
    def tag_values(self) -> Optional[Sequence[str]]:
        """
        Tag values of DevOps Guru app boundary.
        """
        return pulumi.get(self, "tag_values")

