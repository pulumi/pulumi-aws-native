# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'NotificationChannelInsightSeverity',
    'NotificationChannelNotificationMessageType',
    'ResourceCollectionType',
]


class NotificationChannelInsightSeverity(str, Enum):
    """
    DevOps Guru Insight Severity Enum
    """
    LOW = "LOW"
    MEDIUM = "MEDIUM"
    HIGH = "HIGH"


class NotificationChannelNotificationMessageType(str, Enum):
    """
    DevOps Guru NotificationMessageType Enum
    """
    NEW_INSIGHT = "NEW_INSIGHT"
    CLOSED_INSIGHT = "CLOSED_INSIGHT"
    NEW_ASSOCIATION = "NEW_ASSOCIATION"
    SEVERITY_UPGRADED = "SEVERITY_UPGRADED"
    NEW_RECOMMENDATION = "NEW_RECOMMENDATION"


class ResourceCollectionType(str, Enum):
    """
    The type of ResourceCollection
    """
    AWS_CLOUD_FORMATION = "AWS_CLOUD_FORMATION"
    AWS_TAGS = "AWS_TAGS"
