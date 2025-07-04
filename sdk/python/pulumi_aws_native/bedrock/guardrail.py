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
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._enums import *
from ._inputs import *

__all__ = ['GuardrailArgs', 'Guardrail']

@pulumi.input_type
class GuardrailArgs:
    def __init__(__self__, *,
                 blocked_input_messaging: pulumi.Input[builtins.str],
                 blocked_outputs_messaging: pulumi.Input[builtins.str],
                 content_policy_config: Optional[pulumi.Input['GuardrailContentPolicyConfigArgs']] = None,
                 contextual_grounding_policy_config: Optional[pulumi.Input['GuardrailContextualGroundingPolicyConfigArgs']] = None,
                 cross_region_config: Optional[pulumi.Input['GuardrailCrossRegionConfigArgs']] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 kms_key_arn: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 sensitive_information_policy_config: Optional[pulumi.Input['GuardrailSensitiveInformationPolicyConfigArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None,
                 topic_policy_config: Optional[pulumi.Input['GuardrailTopicPolicyConfigArgs']] = None,
                 word_policy_config: Optional[pulumi.Input['GuardrailWordPolicyConfigArgs']] = None):
        """
        The set of arguments for constructing a Guardrail resource.
        :param pulumi.Input[builtins.str] blocked_input_messaging: Messaging for when violations are detected in text
        :param pulumi.Input[builtins.str] blocked_outputs_messaging: Messaging for when violations are detected in text
        :param pulumi.Input['GuardrailContentPolicyConfigArgs'] content_policy_config: The content filter policies to configure for the guardrail.
        :param pulumi.Input[builtins.str] description: Description of the guardrail or its version
        :param pulumi.Input[builtins.str] kms_key_arn: The KMS key with which the guardrail was encrypted at rest
        :param pulumi.Input[builtins.str] name: Name of the guardrail
        :param pulumi.Input['GuardrailSensitiveInformationPolicyConfigArgs'] sensitive_information_policy_config: The sensitive information policy to configure for the guardrail.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: List of Tags
        :param pulumi.Input['GuardrailTopicPolicyConfigArgs'] topic_policy_config: The topic policies to configure for the guardrail.
        :param pulumi.Input['GuardrailWordPolicyConfigArgs'] word_policy_config: The word policy you configure for the guardrail.
        """
        pulumi.set(__self__, "blocked_input_messaging", blocked_input_messaging)
        pulumi.set(__self__, "blocked_outputs_messaging", blocked_outputs_messaging)
        if content_policy_config is not None:
            pulumi.set(__self__, "content_policy_config", content_policy_config)
        if contextual_grounding_policy_config is not None:
            pulumi.set(__self__, "contextual_grounding_policy_config", contextual_grounding_policy_config)
        if cross_region_config is not None:
            pulumi.set(__self__, "cross_region_config", cross_region_config)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if kms_key_arn is not None:
            pulumi.set(__self__, "kms_key_arn", kms_key_arn)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if sensitive_information_policy_config is not None:
            pulumi.set(__self__, "sensitive_information_policy_config", sensitive_information_policy_config)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if topic_policy_config is not None:
            pulumi.set(__self__, "topic_policy_config", topic_policy_config)
        if word_policy_config is not None:
            pulumi.set(__self__, "word_policy_config", word_policy_config)

    @property
    @pulumi.getter(name="blockedInputMessaging")
    def blocked_input_messaging(self) -> pulumi.Input[builtins.str]:
        """
        Messaging for when violations are detected in text
        """
        return pulumi.get(self, "blocked_input_messaging")

    @blocked_input_messaging.setter
    def blocked_input_messaging(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "blocked_input_messaging", value)

    @property
    @pulumi.getter(name="blockedOutputsMessaging")
    def blocked_outputs_messaging(self) -> pulumi.Input[builtins.str]:
        """
        Messaging for when violations are detected in text
        """
        return pulumi.get(self, "blocked_outputs_messaging")

    @blocked_outputs_messaging.setter
    def blocked_outputs_messaging(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "blocked_outputs_messaging", value)

    @property
    @pulumi.getter(name="contentPolicyConfig")
    def content_policy_config(self) -> Optional[pulumi.Input['GuardrailContentPolicyConfigArgs']]:
        """
        The content filter policies to configure for the guardrail.
        """
        return pulumi.get(self, "content_policy_config")

    @content_policy_config.setter
    def content_policy_config(self, value: Optional[pulumi.Input['GuardrailContentPolicyConfigArgs']]):
        pulumi.set(self, "content_policy_config", value)

    @property
    @pulumi.getter(name="contextualGroundingPolicyConfig")
    def contextual_grounding_policy_config(self) -> Optional[pulumi.Input['GuardrailContextualGroundingPolicyConfigArgs']]:
        return pulumi.get(self, "contextual_grounding_policy_config")

    @contextual_grounding_policy_config.setter
    def contextual_grounding_policy_config(self, value: Optional[pulumi.Input['GuardrailContextualGroundingPolicyConfigArgs']]):
        pulumi.set(self, "contextual_grounding_policy_config", value)

    @property
    @pulumi.getter(name="crossRegionConfig")
    def cross_region_config(self) -> Optional[pulumi.Input['GuardrailCrossRegionConfigArgs']]:
        return pulumi.get(self, "cross_region_config")

    @cross_region_config.setter
    def cross_region_config(self, value: Optional[pulumi.Input['GuardrailCrossRegionConfigArgs']]):
        pulumi.set(self, "cross_region_config", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Description of the guardrail or its version
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The KMS key with which the guardrail was encrypted at rest
        """
        return pulumi.get(self, "kms_key_arn")

    @kms_key_arn.setter
    def kms_key_arn(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "kms_key_arn", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the guardrail
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="sensitiveInformationPolicyConfig")
    def sensitive_information_policy_config(self) -> Optional[pulumi.Input['GuardrailSensitiveInformationPolicyConfigArgs']]:
        """
        The sensitive information policy to configure for the guardrail.
        """
        return pulumi.get(self, "sensitive_information_policy_config")

    @sensitive_information_policy_config.setter
    def sensitive_information_policy_config(self, value: Optional[pulumi.Input['GuardrailSensitiveInformationPolicyConfigArgs']]):
        pulumi.set(self, "sensitive_information_policy_config", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        List of Tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="topicPolicyConfig")
    def topic_policy_config(self) -> Optional[pulumi.Input['GuardrailTopicPolicyConfigArgs']]:
        """
        The topic policies to configure for the guardrail.
        """
        return pulumi.get(self, "topic_policy_config")

    @topic_policy_config.setter
    def topic_policy_config(self, value: Optional[pulumi.Input['GuardrailTopicPolicyConfigArgs']]):
        pulumi.set(self, "topic_policy_config", value)

    @property
    @pulumi.getter(name="wordPolicyConfig")
    def word_policy_config(self) -> Optional[pulumi.Input['GuardrailWordPolicyConfigArgs']]:
        """
        The word policy you configure for the guardrail.
        """
        return pulumi.get(self, "word_policy_config")

    @word_policy_config.setter
    def word_policy_config(self, value: Optional[pulumi.Input['GuardrailWordPolicyConfigArgs']]):
        pulumi.set(self, "word_policy_config", value)


@pulumi.type_token("aws-native:bedrock:Guardrail")
class Guardrail(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blocked_input_messaging: Optional[pulumi.Input[builtins.str]] = None,
                 blocked_outputs_messaging: Optional[pulumi.Input[builtins.str]] = None,
                 content_policy_config: Optional[pulumi.Input[Union['GuardrailContentPolicyConfigArgs', 'GuardrailContentPolicyConfigArgsDict']]] = None,
                 contextual_grounding_policy_config: Optional[pulumi.Input[Union['GuardrailContextualGroundingPolicyConfigArgs', 'GuardrailContextualGroundingPolicyConfigArgsDict']]] = None,
                 cross_region_config: Optional[pulumi.Input[Union['GuardrailCrossRegionConfigArgs', 'GuardrailCrossRegionConfigArgsDict']]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 kms_key_arn: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 sensitive_information_policy_config: Optional[pulumi.Input[Union['GuardrailSensitiveInformationPolicyConfigArgs', 'GuardrailSensitiveInformationPolicyConfigArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 topic_policy_config: Optional[pulumi.Input[Union['GuardrailTopicPolicyConfigArgs', 'GuardrailTopicPolicyConfigArgsDict']]] = None,
                 word_policy_config: Optional[pulumi.Input[Union['GuardrailWordPolicyConfigArgs', 'GuardrailWordPolicyConfigArgsDict']]] = None,
                 __props__=None):
        """
        Definition of AWS::Bedrock::Guardrail Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] blocked_input_messaging: Messaging for when violations are detected in text
        :param pulumi.Input[builtins.str] blocked_outputs_messaging: Messaging for when violations are detected in text
        :param pulumi.Input[Union['GuardrailContentPolicyConfigArgs', 'GuardrailContentPolicyConfigArgsDict']] content_policy_config: The content filter policies to configure for the guardrail.
        :param pulumi.Input[builtins.str] description: Description of the guardrail or its version
        :param pulumi.Input[builtins.str] kms_key_arn: The KMS key with which the guardrail was encrypted at rest
        :param pulumi.Input[builtins.str] name: Name of the guardrail
        :param pulumi.Input[Union['GuardrailSensitiveInformationPolicyConfigArgs', 'GuardrailSensitiveInformationPolicyConfigArgsDict']] sensitive_information_policy_config: The sensitive information policy to configure for the guardrail.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] tags: List of Tags
        :param pulumi.Input[Union['GuardrailTopicPolicyConfigArgs', 'GuardrailTopicPolicyConfigArgsDict']] topic_policy_config: The topic policies to configure for the guardrail.
        :param pulumi.Input[Union['GuardrailWordPolicyConfigArgs', 'GuardrailWordPolicyConfigArgsDict']] word_policy_config: The word policy you configure for the guardrail.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GuardrailArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::Bedrock::Guardrail Resource Type

        :param str resource_name: The name of the resource.
        :param GuardrailArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GuardrailArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blocked_input_messaging: Optional[pulumi.Input[builtins.str]] = None,
                 blocked_outputs_messaging: Optional[pulumi.Input[builtins.str]] = None,
                 content_policy_config: Optional[pulumi.Input[Union['GuardrailContentPolicyConfigArgs', 'GuardrailContentPolicyConfigArgsDict']]] = None,
                 contextual_grounding_policy_config: Optional[pulumi.Input[Union['GuardrailContextualGroundingPolicyConfigArgs', 'GuardrailContextualGroundingPolicyConfigArgsDict']]] = None,
                 cross_region_config: Optional[pulumi.Input[Union['GuardrailCrossRegionConfigArgs', 'GuardrailCrossRegionConfigArgsDict']]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 kms_key_arn: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 sensitive_information_policy_config: Optional[pulumi.Input[Union['GuardrailSensitiveInformationPolicyConfigArgs', 'GuardrailSensitiveInformationPolicyConfigArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 topic_policy_config: Optional[pulumi.Input[Union['GuardrailTopicPolicyConfigArgs', 'GuardrailTopicPolicyConfigArgsDict']]] = None,
                 word_policy_config: Optional[pulumi.Input[Union['GuardrailWordPolicyConfigArgs', 'GuardrailWordPolicyConfigArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GuardrailArgs.__new__(GuardrailArgs)

            if blocked_input_messaging is None and not opts.urn:
                raise TypeError("Missing required property 'blocked_input_messaging'")
            __props__.__dict__["blocked_input_messaging"] = blocked_input_messaging
            if blocked_outputs_messaging is None and not opts.urn:
                raise TypeError("Missing required property 'blocked_outputs_messaging'")
            __props__.__dict__["blocked_outputs_messaging"] = blocked_outputs_messaging
            __props__.__dict__["content_policy_config"] = content_policy_config
            __props__.__dict__["contextual_grounding_policy_config"] = contextual_grounding_policy_config
            __props__.__dict__["cross_region_config"] = cross_region_config
            __props__.__dict__["description"] = description
            __props__.__dict__["kms_key_arn"] = kms_key_arn
            __props__.__dict__["name"] = name
            __props__.__dict__["sensitive_information_policy_config"] = sensitive_information_policy_config
            __props__.__dict__["tags"] = tags
            __props__.__dict__["topic_policy_config"] = topic_policy_config
            __props__.__dict__["word_policy_config"] = word_policy_config
            __props__.__dict__["created_at"] = None
            __props__.__dict__["failure_recommendations"] = None
            __props__.__dict__["guardrail_arn"] = None
            __props__.__dict__["guardrail_id"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["status_reasons"] = None
            __props__.__dict__["updated_at"] = None
            __props__.__dict__["version"] = None
        super(Guardrail, __self__).__init__(
            'aws-native:bedrock:Guardrail',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Guardrail':
        """
        Get an existing Guardrail resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = GuardrailArgs.__new__(GuardrailArgs)

        __props__.__dict__["blocked_input_messaging"] = None
        __props__.__dict__["blocked_outputs_messaging"] = None
        __props__.__dict__["content_policy_config"] = None
        __props__.__dict__["contextual_grounding_policy_config"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["cross_region_config"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["failure_recommendations"] = None
        __props__.__dict__["guardrail_arn"] = None
        __props__.__dict__["guardrail_id"] = None
        __props__.__dict__["kms_key_arn"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["sensitive_information_policy_config"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["status_reasons"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["topic_policy_config"] = None
        __props__.__dict__["updated_at"] = None
        __props__.__dict__["version"] = None
        __props__.__dict__["word_policy_config"] = None
        return Guardrail(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="blockedInputMessaging")
    def blocked_input_messaging(self) -> pulumi.Output[builtins.str]:
        """
        Messaging for when violations are detected in text
        """
        return pulumi.get(self, "blocked_input_messaging")

    @property
    @pulumi.getter(name="blockedOutputsMessaging")
    def blocked_outputs_messaging(self) -> pulumi.Output[builtins.str]:
        """
        Messaging for when violations are detected in text
        """
        return pulumi.get(self, "blocked_outputs_messaging")

    @property
    @pulumi.getter(name="contentPolicyConfig")
    def content_policy_config(self) -> pulumi.Output[Optional['outputs.GuardrailContentPolicyConfig']]:
        """
        The content filter policies to configure for the guardrail.
        """
        return pulumi.get(self, "content_policy_config")

    @property
    @pulumi.getter(name="contextualGroundingPolicyConfig")
    def contextual_grounding_policy_config(self) -> pulumi.Output[Optional['outputs.GuardrailContextualGroundingPolicyConfig']]:
        return pulumi.get(self, "contextual_grounding_policy_config")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        """
        Time Stamp
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="crossRegionConfig")
    def cross_region_config(self) -> pulumi.Output[Optional['outputs.GuardrailCrossRegionConfig']]:
        return pulumi.get(self, "cross_region_config")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Description of the guardrail or its version
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="failureRecommendations")
    def failure_recommendations(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        List of failure recommendations
        """
        return pulumi.get(self, "failure_recommendations")

    @property
    @pulumi.getter(name="guardrailArn")
    def guardrail_arn(self) -> pulumi.Output[builtins.str]:
        """
        Arn representation for the guardrail
        """
        return pulumi.get(self, "guardrail_arn")

    @property
    @pulumi.getter(name="guardrailId")
    def guardrail_id(self) -> pulumi.Output[builtins.str]:
        """
        Unique id for the guardrail
        """
        return pulumi.get(self, "guardrail_id")

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The KMS key with which the guardrail was encrypted at rest
        """
        return pulumi.get(self, "kms_key_arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the guardrail
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="sensitiveInformationPolicyConfig")
    def sensitive_information_policy_config(self) -> pulumi.Output[Optional['outputs.GuardrailSensitiveInformationPolicyConfig']]:
        """
        The sensitive information policy to configure for the guardrail.
        """
        return pulumi.get(self, "sensitive_information_policy_config")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['GuardrailStatus']:
        """
        The status of the guardrail.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusReasons")
    def status_reasons(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        List of status reasons
        """
        return pulumi.get(self, "status_reasons")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        List of Tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="topicPolicyConfig")
    def topic_policy_config(self) -> pulumi.Output[Optional['outputs.GuardrailTopicPolicyConfig']]:
        """
        The topic policies to configure for the guardrail.
        """
        return pulumi.get(self, "topic_policy_config")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[builtins.str]:
        """
        Time Stamp
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[builtins.str]:
        """
        Guardrail version
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="wordPolicyConfig")
    def word_policy_config(self) -> pulumi.Output[Optional['outputs.GuardrailWordPolicyConfig']]:
        """
        The word policy you configure for the guardrail.
        """
        return pulumi.get(self, "word_policy_config")

