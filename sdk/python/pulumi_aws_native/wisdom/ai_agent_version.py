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

__all__ = ['AiAgentVersionArgs', 'AiAgentVersion']

@pulumi.input_type
class AiAgentVersionArgs:
    def __init__(__self__, *,
                 ai_agent_id: pulumi.Input[builtins.str],
                 assistant_id: pulumi.Input[builtins.str],
                 modified_time_seconds: Optional[pulumi.Input[builtins.float]] = None):
        """
        The set of arguments for constructing a AiAgentVersion resource.
        :param pulumi.Input[builtins.str] ai_agent_id: The identifier of the AI Agent.
        :param pulumi.Input[builtins.float] modified_time_seconds: The time the AI Agent version was last modified in seconds.
        """
        pulumi.set(__self__, "ai_agent_id", ai_agent_id)
        pulumi.set(__self__, "assistant_id", assistant_id)
        if modified_time_seconds is not None:
            pulumi.set(__self__, "modified_time_seconds", modified_time_seconds)

    @property
    @pulumi.getter(name="aiAgentId")
    def ai_agent_id(self) -> pulumi.Input[builtins.str]:
        """
        The identifier of the AI Agent.
        """
        return pulumi.get(self, "ai_agent_id")

    @ai_agent_id.setter
    def ai_agent_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "ai_agent_id", value)

    @property
    @pulumi.getter(name="assistantId")
    def assistant_id(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "assistant_id")

    @assistant_id.setter
    def assistant_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "assistant_id", value)

    @property
    @pulumi.getter(name="modifiedTimeSeconds")
    def modified_time_seconds(self) -> Optional[pulumi.Input[builtins.float]]:
        """
        The time the AI Agent version was last modified in seconds.
        """
        return pulumi.get(self, "modified_time_seconds")

    @modified_time_seconds.setter
    def modified_time_seconds(self, value: Optional[pulumi.Input[builtins.float]]):
        pulumi.set(self, "modified_time_seconds", value)


@pulumi.type_token("aws-native:wisdom:AiAgentVersion")
class AiAgentVersion(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ai_agent_id: Optional[pulumi.Input[builtins.str]] = None,
                 assistant_id: Optional[pulumi.Input[builtins.str]] = None,
                 modified_time_seconds: Optional[pulumi.Input[builtins.float]] = None,
                 __props__=None):
        """
        Definition of AWS::Wisdom::AIAgentVersion Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] ai_agent_id: The identifier of the AI Agent.
        :param pulumi.Input[builtins.float] modified_time_seconds: The time the AI Agent version was last modified in seconds.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AiAgentVersionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::Wisdom::AIAgentVersion Resource Type

        :param str resource_name: The name of the resource.
        :param AiAgentVersionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AiAgentVersionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ai_agent_id: Optional[pulumi.Input[builtins.str]] = None,
                 assistant_id: Optional[pulumi.Input[builtins.str]] = None,
                 modified_time_seconds: Optional[pulumi.Input[builtins.float]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AiAgentVersionArgs.__new__(AiAgentVersionArgs)

            if ai_agent_id is None and not opts.urn:
                raise TypeError("Missing required property 'ai_agent_id'")
            __props__.__dict__["ai_agent_id"] = ai_agent_id
            if assistant_id is None and not opts.urn:
                raise TypeError("Missing required property 'assistant_id'")
            __props__.__dict__["assistant_id"] = assistant_id
            __props__.__dict__["modified_time_seconds"] = modified_time_seconds
            __props__.__dict__["ai_agent_arn"] = None
            __props__.__dict__["ai_agent_version_id"] = None
            __props__.__dict__["assistant_arn"] = None
            __props__.__dict__["version_number"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["aiAgentId", "assistantId", "modifiedTimeSeconds"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(AiAgentVersion, __self__).__init__(
            'aws-native:wisdom:AiAgentVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AiAgentVersion':
        """
        Get an existing AiAgentVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AiAgentVersionArgs.__new__(AiAgentVersionArgs)

        __props__.__dict__["ai_agent_arn"] = None
        __props__.__dict__["ai_agent_id"] = None
        __props__.__dict__["ai_agent_version_id"] = None
        __props__.__dict__["assistant_arn"] = None
        __props__.__dict__["assistant_id"] = None
        __props__.__dict__["modified_time_seconds"] = None
        __props__.__dict__["version_number"] = None
        return AiAgentVersion(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="aiAgentArn")
    def ai_agent_arn(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "ai_agent_arn")

    @property
    @pulumi.getter(name="aiAgentId")
    def ai_agent_id(self) -> pulumi.Output[builtins.str]:
        """
        The identifier of the AI Agent.
        """
        return pulumi.get(self, "ai_agent_id")

    @property
    @pulumi.getter(name="aiAgentVersionId")
    def ai_agent_version_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "ai_agent_version_id")

    @property
    @pulumi.getter(name="assistantArn")
    def assistant_arn(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "assistant_arn")

    @property
    @pulumi.getter(name="assistantId")
    def assistant_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "assistant_id")

    @property
    @pulumi.getter(name="modifiedTimeSeconds")
    def modified_time_seconds(self) -> pulumi.Output[Optional[builtins.float]]:
        """
        The time the AI Agent version was last modified in seconds.
        """
        return pulumi.get(self, "modified_time_seconds")

    @property
    @pulumi.getter(name="versionNumber")
    def version_number(self) -> pulumi.Output[builtins.float]:
        """
        The version number for this AI Agent version.
        """
        return pulumi.get(self, "version_number")

