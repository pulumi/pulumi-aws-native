# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import pulumi
from enum import Enum

__all__ = [
    'EventSubscriptionEventCategoriesItem',
    'EventSubscriptionSeverity',
    'EventSubscriptionSourceType',
    'EventSubscriptionStatus',
    'ScheduledActionState',
]


@pulumi.type_token("aws-native:redshift:EventSubscriptionEventCategoriesItem")
class EventSubscriptionEventCategoriesItem(builtins.str, Enum):
    CONFIGURATION = "configuration"
    MANAGEMENT = "management"
    MONITORING = "monitoring"
    SECURITY = "security"
    PENDING = "pending"


@pulumi.type_token("aws-native:redshift:EventSubscriptionSeverity")
class EventSubscriptionSeverity(builtins.str, Enum):
    """
    Specifies the Amazon Redshift event severity to be published by the event notification subscription.
    """
    ERROR = "ERROR"
    INFO = "INFO"


@pulumi.type_token("aws-native:redshift:EventSubscriptionSourceType")
class EventSubscriptionSourceType(builtins.str, Enum):
    """
    The type of source that will be generating the events.
    """
    CLUSTER = "cluster"
    CLUSTER_PARAMETER_GROUP = "cluster-parameter-group"
    CLUSTER_SECURITY_GROUP = "cluster-security-group"
    CLUSTER_SNAPSHOT = "cluster-snapshot"
    SCHEDULED_ACTION = "scheduled-action"


@pulumi.type_token("aws-native:redshift:EventSubscriptionStatus")
class EventSubscriptionStatus(builtins.str, Enum):
    """
    The status of the Amazon Redshift event notification subscription.
    """
    ACTIVE = "active"
    NO_PERMISSION = "no-permission"
    TOPIC_NOT_EXIST = "topic-not-exist"


@pulumi.type_token("aws-native:redshift:ScheduledActionState")
class ScheduledActionState(builtins.str, Enum):
    """
    The state of the scheduled action.
    """
    ACTIVE = "ACTIVE"
    DISABLED = "DISABLED"
