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
    'NotificationRuleTarget',
]

@pulumi.output_type
class NotificationRuleTarget(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "targetAddress":
            suggest = "target_address"
        elif key == "targetType":
            suggest = "target_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NotificationRuleTarget. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NotificationRuleTarget.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NotificationRuleTarget.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 target_address: str,
                 target_type: str):
        """
        :param str target_address: The Amazon Resource Name (ARN) of the AWS Chatbot topic or AWS Chatbot client.
        :param str target_type: The target type. Can be an Amazon Simple Notification Service topic or AWS Chatbot client.
               
               - Amazon Simple Notification Service topics are specified as `SNS` .
               - AWS Chatbot clients are specified as `AWSChatbotSlack` .
               - AWS Chatbot clients for Microsoft Teams are specified as `AWSChatbotMicrosoftTeams` .
        """
        pulumi.set(__self__, "target_address", target_address)
        pulumi.set(__self__, "target_type", target_type)

    @property
    @pulumi.getter(name="targetAddress")
    def target_address(self) -> str:
        """
        The Amazon Resource Name (ARN) of the AWS Chatbot topic or AWS Chatbot client.
        """
        return pulumi.get(self, "target_address")

    @property
    @pulumi.getter(name="targetType")
    def target_type(self) -> str:
        """
        The target type. Can be an Amazon Simple Notification Service topic or AWS Chatbot client.

        - Amazon Simple Notification Service topics are specified as `SNS` .
        - AWS Chatbot clients are specified as `AWSChatbotSlack` .
        - AWS Chatbot clients for Microsoft Teams are specified as `AWSChatbotMicrosoftTeams` .
        """
        return pulumi.get(self, "target_type")


