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
    'InvestigationGroupChatbotNotificationChannelArgs',
    'InvestigationGroupChatbotNotificationChannelArgsDict',
    'InvestigationGroupCrossAccountConfigurationArgs',
    'InvestigationGroupCrossAccountConfigurationArgsDict',
    'InvestigationGroupEncryptionConfigMapArgs',
    'InvestigationGroupEncryptionConfigMapArgsDict',
]

MYPY = False

if not MYPY:
    class InvestigationGroupChatbotNotificationChannelArgsDict(TypedDict):
        chat_configuration_arns: NotRequired[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]
        sns_topic_arn: NotRequired[pulumi.Input[builtins.str]]
elif False:
    InvestigationGroupChatbotNotificationChannelArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class InvestigationGroupChatbotNotificationChannelArgs:
    def __init__(__self__, *,
                 chat_configuration_arns: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 sns_topic_arn: Optional[pulumi.Input[builtins.str]] = None):
        if chat_configuration_arns is not None:
            pulumi.set(__self__, "chat_configuration_arns", chat_configuration_arns)
        if sns_topic_arn is not None:
            pulumi.set(__self__, "sns_topic_arn", sns_topic_arn)

    @property
    @pulumi.getter(name="chatConfigurationArns")
    def chat_configuration_arns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "chat_configuration_arns")

    @chat_configuration_arns.setter
    def chat_configuration_arns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "chat_configuration_arns", value)

    @property
    @pulumi.getter(name="snsTopicArn")
    def sns_topic_arn(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "sns_topic_arn")

    @sns_topic_arn.setter
    def sns_topic_arn(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "sns_topic_arn", value)


if not MYPY:
    class InvestigationGroupCrossAccountConfigurationArgsDict(TypedDict):
        source_role_arn: NotRequired[pulumi.Input[builtins.str]]
elif False:
    InvestigationGroupCrossAccountConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class InvestigationGroupCrossAccountConfigurationArgs:
    def __init__(__self__, *,
                 source_role_arn: Optional[pulumi.Input[builtins.str]] = None):
        if source_role_arn is not None:
            pulumi.set(__self__, "source_role_arn", source_role_arn)

    @property
    @pulumi.getter(name="sourceRoleArn")
    def source_role_arn(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "source_role_arn")

    @source_role_arn.setter
    def source_role_arn(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "source_role_arn", value)


if not MYPY:
    class InvestigationGroupEncryptionConfigMapArgsDict(TypedDict):
        encryption_configuration_type: NotRequired[pulumi.Input[builtins.str]]
        kms_key_id: NotRequired[pulumi.Input[builtins.str]]
elif False:
    InvestigationGroupEncryptionConfigMapArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class InvestigationGroupEncryptionConfigMapArgs:
    def __init__(__self__, *,
                 encryption_configuration_type: Optional[pulumi.Input[builtins.str]] = None,
                 kms_key_id: Optional[pulumi.Input[builtins.str]] = None):
        if encryption_configuration_type is not None:
            pulumi.set(__self__, "encryption_configuration_type", encryption_configuration_type)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)

    @property
    @pulumi.getter(name="encryptionConfigurationType")
    def encryption_configuration_type(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "encryption_configuration_type")

    @encryption_configuration_type.setter
    def encryption_configuration_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "encryption_configuration_type", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "kms_key_id", value)


