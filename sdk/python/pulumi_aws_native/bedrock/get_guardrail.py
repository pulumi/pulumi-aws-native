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
from . import outputs
from .. import outputs as _root_outputs
from ._enums import *

__all__ = [
    'GetGuardrailResult',
    'AwaitableGetGuardrailResult',
    'get_guardrail',
    'get_guardrail_output',
]

@pulumi.output_type
class GetGuardrailResult:
    def __init__(__self__, blocked_input_messaging=None, blocked_outputs_messaging=None, content_policy_config=None, contextual_grounding_policy_config=None, created_at=None, cross_region_config=None, description=None, failure_recommendations=None, guardrail_arn=None, guardrail_id=None, kms_key_arn=None, name=None, sensitive_information_policy_config=None, status=None, status_reasons=None, tags=None, topic_policy_config=None, updated_at=None, version=None, word_policy_config=None):
        if blocked_input_messaging and not isinstance(blocked_input_messaging, str):
            raise TypeError("Expected argument 'blocked_input_messaging' to be a str")
        pulumi.set(__self__, "blocked_input_messaging", blocked_input_messaging)
        if blocked_outputs_messaging and not isinstance(blocked_outputs_messaging, str):
            raise TypeError("Expected argument 'blocked_outputs_messaging' to be a str")
        pulumi.set(__self__, "blocked_outputs_messaging", blocked_outputs_messaging)
        if content_policy_config and not isinstance(content_policy_config, dict):
            raise TypeError("Expected argument 'content_policy_config' to be a dict")
        pulumi.set(__self__, "content_policy_config", content_policy_config)
        if contextual_grounding_policy_config and not isinstance(contextual_grounding_policy_config, dict):
            raise TypeError("Expected argument 'contextual_grounding_policy_config' to be a dict")
        pulumi.set(__self__, "contextual_grounding_policy_config", contextual_grounding_policy_config)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if cross_region_config and not isinstance(cross_region_config, dict):
            raise TypeError("Expected argument 'cross_region_config' to be a dict")
        pulumi.set(__self__, "cross_region_config", cross_region_config)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if failure_recommendations and not isinstance(failure_recommendations, list):
            raise TypeError("Expected argument 'failure_recommendations' to be a list")
        pulumi.set(__self__, "failure_recommendations", failure_recommendations)
        if guardrail_arn and not isinstance(guardrail_arn, str):
            raise TypeError("Expected argument 'guardrail_arn' to be a str")
        pulumi.set(__self__, "guardrail_arn", guardrail_arn)
        if guardrail_id and not isinstance(guardrail_id, str):
            raise TypeError("Expected argument 'guardrail_id' to be a str")
        pulumi.set(__self__, "guardrail_id", guardrail_id)
        if kms_key_arn and not isinstance(kms_key_arn, str):
            raise TypeError("Expected argument 'kms_key_arn' to be a str")
        pulumi.set(__self__, "kms_key_arn", kms_key_arn)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if sensitive_information_policy_config and not isinstance(sensitive_information_policy_config, dict):
            raise TypeError("Expected argument 'sensitive_information_policy_config' to be a dict")
        pulumi.set(__self__, "sensitive_information_policy_config", sensitive_information_policy_config)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if status_reasons and not isinstance(status_reasons, list):
            raise TypeError("Expected argument 'status_reasons' to be a list")
        pulumi.set(__self__, "status_reasons", status_reasons)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if topic_policy_config and not isinstance(topic_policy_config, dict):
            raise TypeError("Expected argument 'topic_policy_config' to be a dict")
        pulumi.set(__self__, "topic_policy_config", topic_policy_config)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)
        if word_policy_config and not isinstance(word_policy_config, dict):
            raise TypeError("Expected argument 'word_policy_config' to be a dict")
        pulumi.set(__self__, "word_policy_config", word_policy_config)

    @property
    @pulumi.getter(name="blockedInputMessaging")
    def blocked_input_messaging(self) -> Optional[builtins.str]:
        """
        Messaging for when violations are detected in text
        """
        return pulumi.get(self, "blocked_input_messaging")

    @property
    @pulumi.getter(name="blockedOutputsMessaging")
    def blocked_outputs_messaging(self) -> Optional[builtins.str]:
        """
        Messaging for when violations are detected in text
        """
        return pulumi.get(self, "blocked_outputs_messaging")

    @property
    @pulumi.getter(name="contentPolicyConfig")
    def content_policy_config(self) -> Optional['outputs.GuardrailContentPolicyConfig']:
        """
        The content filter policies to configure for the guardrail.
        """
        return pulumi.get(self, "content_policy_config")

    @property
    @pulumi.getter(name="contextualGroundingPolicyConfig")
    def contextual_grounding_policy_config(self) -> Optional['outputs.GuardrailContextualGroundingPolicyConfig']:
        return pulumi.get(self, "contextual_grounding_policy_config")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[builtins.str]:
        """
        Time Stamp
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="crossRegionConfig")
    def cross_region_config(self) -> Optional['outputs.GuardrailCrossRegionConfig']:
        return pulumi.get(self, "cross_region_config")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        Description of the guardrail or its version
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="failureRecommendations")
    def failure_recommendations(self) -> Optional[Sequence[builtins.str]]:
        """
        List of failure recommendations
        """
        return pulumi.get(self, "failure_recommendations")

    @property
    @pulumi.getter(name="guardrailArn")
    def guardrail_arn(self) -> Optional[builtins.str]:
        """
        Arn representation for the guardrail
        """
        return pulumi.get(self, "guardrail_arn")

    @property
    @pulumi.getter(name="guardrailId")
    def guardrail_id(self) -> Optional[builtins.str]:
        """
        Unique id for the guardrail
        """
        return pulumi.get(self, "guardrail_id")

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> Optional[builtins.str]:
        """
        The KMS key with which the guardrail was encrypted at rest
        """
        return pulumi.get(self, "kms_key_arn")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        """
        Name of the guardrail
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="sensitiveInformationPolicyConfig")
    def sensitive_information_policy_config(self) -> Optional['outputs.GuardrailSensitiveInformationPolicyConfig']:
        """
        The sensitive information policy to configure for the guardrail.
        """
        return pulumi.get(self, "sensitive_information_policy_config")

    @property
    @pulumi.getter
    def status(self) -> Optional['GuardrailStatus']:
        """
        The status of the guardrail.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusReasons")
    def status_reasons(self) -> Optional[Sequence[builtins.str]]:
        """
        List of status reasons
        """
        return pulumi.get(self, "status_reasons")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        List of Tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="topicPolicyConfig")
    def topic_policy_config(self) -> Optional['outputs.GuardrailTopicPolicyConfig']:
        """
        The topic policies to configure for the guardrail.
        """
        return pulumi.get(self, "topic_policy_config")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[builtins.str]:
        """
        Time Stamp
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def version(self) -> Optional[builtins.str]:
        """
        Guardrail version
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="wordPolicyConfig")
    def word_policy_config(self) -> Optional['outputs.GuardrailWordPolicyConfig']:
        """
        The word policy you configure for the guardrail.
        """
        return pulumi.get(self, "word_policy_config")


class AwaitableGetGuardrailResult(GetGuardrailResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGuardrailResult(
            blocked_input_messaging=self.blocked_input_messaging,
            blocked_outputs_messaging=self.blocked_outputs_messaging,
            content_policy_config=self.content_policy_config,
            contextual_grounding_policy_config=self.contextual_grounding_policy_config,
            created_at=self.created_at,
            cross_region_config=self.cross_region_config,
            description=self.description,
            failure_recommendations=self.failure_recommendations,
            guardrail_arn=self.guardrail_arn,
            guardrail_id=self.guardrail_id,
            kms_key_arn=self.kms_key_arn,
            name=self.name,
            sensitive_information_policy_config=self.sensitive_information_policy_config,
            status=self.status,
            status_reasons=self.status_reasons,
            tags=self.tags,
            topic_policy_config=self.topic_policy_config,
            updated_at=self.updated_at,
            version=self.version,
            word_policy_config=self.word_policy_config)


def get_guardrail(guardrail_arn: Optional[builtins.str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGuardrailResult:
    """
    Definition of AWS::Bedrock::Guardrail Resource Type


    :param builtins.str guardrail_arn: Arn representation for the guardrail
    """
    __args__ = dict()
    __args__['guardrailArn'] = guardrail_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:bedrock:getGuardrail', __args__, opts=opts, typ=GetGuardrailResult).value

    return AwaitableGetGuardrailResult(
        blocked_input_messaging=pulumi.get(__ret__, 'blocked_input_messaging'),
        blocked_outputs_messaging=pulumi.get(__ret__, 'blocked_outputs_messaging'),
        content_policy_config=pulumi.get(__ret__, 'content_policy_config'),
        contextual_grounding_policy_config=pulumi.get(__ret__, 'contextual_grounding_policy_config'),
        created_at=pulumi.get(__ret__, 'created_at'),
        cross_region_config=pulumi.get(__ret__, 'cross_region_config'),
        description=pulumi.get(__ret__, 'description'),
        failure_recommendations=pulumi.get(__ret__, 'failure_recommendations'),
        guardrail_arn=pulumi.get(__ret__, 'guardrail_arn'),
        guardrail_id=pulumi.get(__ret__, 'guardrail_id'),
        kms_key_arn=pulumi.get(__ret__, 'kms_key_arn'),
        name=pulumi.get(__ret__, 'name'),
        sensitive_information_policy_config=pulumi.get(__ret__, 'sensitive_information_policy_config'),
        status=pulumi.get(__ret__, 'status'),
        status_reasons=pulumi.get(__ret__, 'status_reasons'),
        tags=pulumi.get(__ret__, 'tags'),
        topic_policy_config=pulumi.get(__ret__, 'topic_policy_config'),
        updated_at=pulumi.get(__ret__, 'updated_at'),
        version=pulumi.get(__ret__, 'version'),
        word_policy_config=pulumi.get(__ret__, 'word_policy_config'))
def get_guardrail_output(guardrail_arn: Optional[pulumi.Input[builtins.str]] = None,
                         opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetGuardrailResult]:
    """
    Definition of AWS::Bedrock::Guardrail Resource Type


    :param builtins.str guardrail_arn: Arn representation for the guardrail
    """
    __args__ = dict()
    __args__['guardrailArn'] = guardrail_arn
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:bedrock:getGuardrail', __args__, opts=opts, typ=GetGuardrailResult)
    return __ret__.apply(lambda __response__: GetGuardrailResult(
        blocked_input_messaging=pulumi.get(__response__, 'blocked_input_messaging'),
        blocked_outputs_messaging=pulumi.get(__response__, 'blocked_outputs_messaging'),
        content_policy_config=pulumi.get(__response__, 'content_policy_config'),
        contextual_grounding_policy_config=pulumi.get(__response__, 'contextual_grounding_policy_config'),
        created_at=pulumi.get(__response__, 'created_at'),
        cross_region_config=pulumi.get(__response__, 'cross_region_config'),
        description=pulumi.get(__response__, 'description'),
        failure_recommendations=pulumi.get(__response__, 'failure_recommendations'),
        guardrail_arn=pulumi.get(__response__, 'guardrail_arn'),
        guardrail_id=pulumi.get(__response__, 'guardrail_id'),
        kms_key_arn=pulumi.get(__response__, 'kms_key_arn'),
        name=pulumi.get(__response__, 'name'),
        sensitive_information_policy_config=pulumi.get(__response__, 'sensitive_information_policy_config'),
        status=pulumi.get(__response__, 'status'),
        status_reasons=pulumi.get(__response__, 'status_reasons'),
        tags=pulumi.get(__response__, 'tags'),
        topic_policy_config=pulumi.get(__response__, 'topic_policy_config'),
        updated_at=pulumi.get(__response__, 'updated_at'),
        version=pulumi.get(__response__, 'version'),
        word_policy_config=pulumi.get(__response__, 'word_policy_config')))
