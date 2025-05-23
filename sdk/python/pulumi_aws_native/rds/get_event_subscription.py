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

__all__ = [
    'GetEventSubscriptionResult',
    'AwaitableGetEventSubscriptionResult',
    'get_event_subscription',
    'get_event_subscription_output',
]

@pulumi.output_type
class GetEventSubscriptionResult:
    def __init__(__self__, enabled=None, event_categories=None, source_ids=None, source_type=None, tags=None):
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if event_categories and not isinstance(event_categories, list):
            raise TypeError("Expected argument 'event_categories' to be a list")
        pulumi.set(__self__, "event_categories", event_categories)
        if source_ids and not isinstance(source_ids, list):
            raise TypeError("Expected argument 'source_ids' to be a list")
        pulumi.set(__self__, "source_ids", source_ids)
        if source_type and not isinstance(source_type, str):
            raise TypeError("Expected argument 'source_type' to be a str")
        pulumi.set(__self__, "source_type", source_type)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[builtins.bool]:
        """
        Specifies whether to activate the subscription. If the event notification subscription isn't activated, the subscription is created but not active.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="eventCategories")
    def event_categories(self) -> Optional[Sequence[builtins.str]]:
        """
        A list of event categories for a particular source type (``SourceType``) that you want to subscribe to. You can see a list of the categories for a given source type in the "Amazon RDS event categories and event messages" section of the [Amazon RDS User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.Messages.html) or the [Amazon Aurora User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Events.Messages.html). You can also see this list by using the ``DescribeEventCategories`` operation.
        """
        return pulumi.get(self, "event_categories")

    @property
    @pulumi.getter(name="sourceIds")
    def source_ids(self) -> Optional[Sequence[builtins.str]]:
        """
        The list of identifiers of the event sources for which events are returned. If not specified, then all sources are included in the response. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens. It can't end with a hyphen or contain two consecutive hyphens.
         Constraints:
          +  If ``SourceIds`` are supplied, ``SourceType`` must also be provided.
          +  If the source type is a DB instance, a ``DBInstanceIdentifier`` value must be supplied.
          +  If the source type is a DB cluster, a ``DBClusterIdentifier`` value must be supplied.
          +  If the source type is a DB parameter group, a ``DBParameterGroupName`` value must be supplied.
          +  If the source type is a DB security group, a ``DBSecurityGroupName`` value must be supplied.
          +  If the source type is a DB snapshot, a ``DBSnapshotIdentifier`` value must be supplied.
          +  If the source type is a DB cluster snapshot, a ``DBClusterSnapshotIdentifier`` value must be supplied.
          +  If the source type is an RDS Proxy, a ``DBProxyName`` value must be supplied.
        """
        return pulumi.get(self, "source_ids")

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> Optional[builtins.str]:
        """
        The type of source that is generating the events. For example, if you want to be notified of events generated by a DB instance, you set this parameter to ``db-instance``. For RDS Proxy events, specify ``db-proxy``. If this value isn't specified, all events are returned.
         Valid Values:``db-instance | db-cluster | db-parameter-group | db-security-group | db-snapshot | db-cluster-snapshot | db-proxy | zero-etl | custom-engine-version | blue-green-deployment``
        """
        return pulumi.get(self, "source_type")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        An optional array of key-value pairs to apply to this subscription.
        """
        return pulumi.get(self, "tags")


class AwaitableGetEventSubscriptionResult(GetEventSubscriptionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEventSubscriptionResult(
            enabled=self.enabled,
            event_categories=self.event_categories,
            source_ids=self.source_ids,
            source_type=self.source_type,
            tags=self.tags)


def get_event_subscription(subscription_name: Optional[builtins.str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEventSubscriptionResult:
    """
    The ``AWS::RDS::EventSubscription`` resource allows you to receive notifications for Amazon Relational Database Service events through the Amazon Simple Notification Service (Amazon SNS). For more information, see [Using Amazon RDS Event Notification](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.html) in the *Amazon RDS User Guide*.


    :param builtins.str subscription_name: The name of the subscription.
            Constraints: The name must be less than 255 characters.
    """
    __args__ = dict()
    __args__['subscriptionName'] = subscription_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:rds:getEventSubscription', __args__, opts=opts, typ=GetEventSubscriptionResult).value

    return AwaitableGetEventSubscriptionResult(
        enabled=pulumi.get(__ret__, 'enabled'),
        event_categories=pulumi.get(__ret__, 'event_categories'),
        source_ids=pulumi.get(__ret__, 'source_ids'),
        source_type=pulumi.get(__ret__, 'source_type'),
        tags=pulumi.get(__ret__, 'tags'))
def get_event_subscription_output(subscription_name: Optional[pulumi.Input[builtins.str]] = None,
                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetEventSubscriptionResult]:
    """
    The ``AWS::RDS::EventSubscription`` resource allows you to receive notifications for Amazon Relational Database Service events through the Amazon Simple Notification Service (Amazon SNS). For more information, see [Using Amazon RDS Event Notification](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.html) in the *Amazon RDS User Guide*.


    :param builtins.str subscription_name: The name of the subscription.
            Constraints: The name must be less than 255 characters.
    """
    __args__ = dict()
    __args__['subscriptionName'] = subscription_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:rds:getEventSubscription', __args__, opts=opts, typ=GetEventSubscriptionResult)
    return __ret__.apply(lambda __response__: GetEventSubscriptionResult(
        enabled=pulumi.get(__response__, 'enabled'),
        event_categories=pulumi.get(__response__, 'event_categories'),
        source_ids=pulumi.get(__response__, 'source_ids'),
        source_type=pulumi.get(__response__, 'source_type'),
        tags=pulumi.get(__response__, 'tags')))
