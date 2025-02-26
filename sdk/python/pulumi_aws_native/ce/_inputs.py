# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
from ._enums import *

__all__ = [
    'AnomalyMonitorResourceTagArgs',
    'AnomalyMonitorResourceTagArgsDict',
    'AnomalySubscriptionResourceTagArgs',
    'AnomalySubscriptionResourceTagArgsDict',
    'AnomalySubscriptionSubscriberArgs',
    'AnomalySubscriptionSubscriberArgsDict',
]

MYPY = False

if not MYPY:
    class AnomalyMonitorResourceTagArgsDict(TypedDict):
        """
        A key-value pair to associate with a resource.
        """
        key: pulumi.Input[str]
        """
        The key name for the tag.
        """
        value: pulumi.Input[str]
        """
        The value for the tag.
        """
elif False:
    AnomalyMonitorResourceTagArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class AnomalyMonitorResourceTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A key-value pair to associate with a resource.
        :param pulumi.Input[str] key: The key name for the tag.
        :param pulumi.Input[str] value: The value for the tag.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The key name for the tag.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value for the tag.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


if not MYPY:
    class AnomalySubscriptionResourceTagArgsDict(TypedDict):
        """
        A key-value pair to associate with a resource.
        """
        key: pulumi.Input[str]
        """
        The key name for the tag.
        """
        value: pulumi.Input[str]
        """
        The value for the tag.
        """
elif False:
    AnomalySubscriptionResourceTagArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class AnomalySubscriptionResourceTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A key-value pair to associate with a resource.
        :param pulumi.Input[str] key: The key name for the tag.
        :param pulumi.Input[str] value: The value for the tag.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The key name for the tag.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value for the tag.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


if not MYPY:
    class AnomalySubscriptionSubscriberArgsDict(TypedDict):
        address: pulumi.Input[str]
        """
        The email address or SNS Topic Amazon Resource Name (ARN), depending on the `Type` .
        """
        type: pulumi.Input['AnomalySubscriptionSubscriberType']
        """
        The notification delivery channel.
        """
        status: NotRequired[pulumi.Input['AnomalySubscriptionSubscriberStatus']]
        """
        Indicates if the subscriber accepts the notifications.
        """
elif False:
    AnomalySubscriptionSubscriberArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class AnomalySubscriptionSubscriberArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str],
                 type: pulumi.Input['AnomalySubscriptionSubscriberType'],
                 status: Optional[pulumi.Input['AnomalySubscriptionSubscriberStatus']] = None):
        """
        :param pulumi.Input[str] address: The email address or SNS Topic Amazon Resource Name (ARN), depending on the `Type` .
        :param pulumi.Input['AnomalySubscriptionSubscriberType'] type: The notification delivery channel.
        :param pulumi.Input['AnomalySubscriptionSubscriberStatus'] status: Indicates if the subscriber accepts the notifications.
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "type", type)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        """
        The email address or SNS Topic Amazon Resource Name (ARN), depending on the `Type` .
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['AnomalySubscriptionSubscriberType']:
        """
        The notification delivery channel.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['AnomalySubscriptionSubscriberType']):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['AnomalySubscriptionSubscriberStatus']]:
        """
        Indicates if the subscriber accepts the notifications.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['AnomalySubscriptionSubscriberStatus']]):
        pulumi.set(self, "status", value)


