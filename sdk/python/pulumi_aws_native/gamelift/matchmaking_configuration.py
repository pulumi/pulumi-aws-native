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

__all__ = ['MatchmakingConfigurationArgs', 'MatchmakingConfiguration']

@pulumi.input_type
class MatchmakingConfigurationArgs:
    def __init__(__self__, *,
                 acceptance_required: pulumi.Input[builtins.bool],
                 request_timeout_seconds: pulumi.Input[builtins.int],
                 rule_set_name: pulumi.Input[builtins.str],
                 acceptance_timeout_seconds: Optional[pulumi.Input[builtins.int]] = None,
                 additional_player_count: Optional[pulumi.Input[builtins.int]] = None,
                 backfill_mode: Optional[pulumi.Input['MatchmakingConfigurationBackfillMode']] = None,
                 creation_time: Optional[pulumi.Input[builtins.str]] = None,
                 custom_event_data: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 flex_match_mode: Optional[pulumi.Input['MatchmakingConfigurationFlexMatchMode']] = None,
                 game_properties: Optional[pulumi.Input[Sequence[pulumi.Input['MatchmakingConfigurationGamePropertyArgs']]]] = None,
                 game_session_data: Optional[pulumi.Input[builtins.str]] = None,
                 game_session_queue_arns: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 notification_target: Optional[pulumi.Input[builtins.str]] = None,
                 rule_set_arn: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a MatchmakingConfiguration resource.
        :param pulumi.Input[builtins.bool] acceptance_required: A flag that indicates whether a match that was created with this configuration must be accepted by the matched players
        :param pulumi.Input[builtins.int] request_timeout_seconds: The maximum duration, in seconds, that a matchmaking ticket can remain in process before timing out.
        :param pulumi.Input[builtins.str] rule_set_name: A unique identifier for the matchmaking rule set to use with this configuration.
        :param pulumi.Input[builtins.int] acceptance_timeout_seconds: The length of time (in seconds) to wait for players to accept a proposed match, if acceptance is required.
        :param pulumi.Input[builtins.int] additional_player_count: The number of player slots in a match to keep open for future players.
        :param pulumi.Input['MatchmakingConfigurationBackfillMode'] backfill_mode: The method used to backfill game sessions created with this matchmaking configuration.
        :param pulumi.Input[builtins.str] creation_time: A time stamp indicating when this data object was created.
        :param pulumi.Input[builtins.str] custom_event_data: Information to attach to all events related to the matchmaking configuration.
        :param pulumi.Input[builtins.str] description: A descriptive label that is associated with matchmaking configuration.
        :param pulumi.Input['MatchmakingConfigurationFlexMatchMode'] flex_match_mode: Indicates whether this matchmaking configuration is being used with Amazon GameLift hosting or as a standalone matchmaking solution.
        :param pulumi.Input[Sequence[pulumi.Input['MatchmakingConfigurationGamePropertyArgs']]] game_properties: A set of custom properties for a game session, formatted as key:value pairs.
        :param pulumi.Input[builtins.str] game_session_data: A set of custom game session properties, formatted as a single string value.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] game_session_queue_arns: The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift game session queue resource and uniquely identifies it.
        :param pulumi.Input[builtins.str] name: A unique identifier for the matchmaking configuration.
        :param pulumi.Input[builtins.str] notification_target: An SNS topic ARN that is set up to receive matchmaking notifications.
        :param pulumi.Input[builtins.str] rule_set_arn: The Amazon Resource Name (ARN) associated with the GameLift matchmaking rule set resource that this configuration uses.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: An array of key-value pairs to apply to this resource.
        """
        pulumi.set(__self__, "acceptance_required", acceptance_required)
        pulumi.set(__self__, "request_timeout_seconds", request_timeout_seconds)
        pulumi.set(__self__, "rule_set_name", rule_set_name)
        if acceptance_timeout_seconds is not None:
            pulumi.set(__self__, "acceptance_timeout_seconds", acceptance_timeout_seconds)
        if additional_player_count is not None:
            pulumi.set(__self__, "additional_player_count", additional_player_count)
        if backfill_mode is not None:
            pulumi.set(__self__, "backfill_mode", backfill_mode)
        if creation_time is not None:
            pulumi.set(__self__, "creation_time", creation_time)
        if custom_event_data is not None:
            pulumi.set(__self__, "custom_event_data", custom_event_data)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if flex_match_mode is not None:
            pulumi.set(__self__, "flex_match_mode", flex_match_mode)
        if game_properties is not None:
            pulumi.set(__self__, "game_properties", game_properties)
        if game_session_data is not None:
            pulumi.set(__self__, "game_session_data", game_session_data)
        if game_session_queue_arns is not None:
            pulumi.set(__self__, "game_session_queue_arns", game_session_queue_arns)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notification_target is not None:
            pulumi.set(__self__, "notification_target", notification_target)
        if rule_set_arn is not None:
            pulumi.set(__self__, "rule_set_arn", rule_set_arn)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="acceptanceRequired")
    def acceptance_required(self) -> pulumi.Input[builtins.bool]:
        """
        A flag that indicates whether a match that was created with this configuration must be accepted by the matched players
        """
        return pulumi.get(self, "acceptance_required")

    @acceptance_required.setter
    def acceptance_required(self, value: pulumi.Input[builtins.bool]):
        pulumi.set(self, "acceptance_required", value)

    @property
    @pulumi.getter(name="requestTimeoutSeconds")
    def request_timeout_seconds(self) -> pulumi.Input[builtins.int]:
        """
        The maximum duration, in seconds, that a matchmaking ticket can remain in process before timing out.
        """
        return pulumi.get(self, "request_timeout_seconds")

    @request_timeout_seconds.setter
    def request_timeout_seconds(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "request_timeout_seconds", value)

    @property
    @pulumi.getter(name="ruleSetName")
    def rule_set_name(self) -> pulumi.Input[builtins.str]:
        """
        A unique identifier for the matchmaking rule set to use with this configuration.
        """
        return pulumi.get(self, "rule_set_name")

    @rule_set_name.setter
    def rule_set_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "rule_set_name", value)

    @property
    @pulumi.getter(name="acceptanceTimeoutSeconds")
    def acceptance_timeout_seconds(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The length of time (in seconds) to wait for players to accept a proposed match, if acceptance is required.
        """
        return pulumi.get(self, "acceptance_timeout_seconds")

    @acceptance_timeout_seconds.setter
    def acceptance_timeout_seconds(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "acceptance_timeout_seconds", value)

    @property
    @pulumi.getter(name="additionalPlayerCount")
    def additional_player_count(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The number of player slots in a match to keep open for future players.
        """
        return pulumi.get(self, "additional_player_count")

    @additional_player_count.setter
    def additional_player_count(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "additional_player_count", value)

    @property
    @pulumi.getter(name="backfillMode")
    def backfill_mode(self) -> Optional[pulumi.Input['MatchmakingConfigurationBackfillMode']]:
        """
        The method used to backfill game sessions created with this matchmaking configuration.
        """
        return pulumi.get(self, "backfill_mode")

    @backfill_mode.setter
    def backfill_mode(self, value: Optional[pulumi.Input['MatchmakingConfigurationBackfillMode']]):
        pulumi.set(self, "backfill_mode", value)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A time stamp indicating when this data object was created.
        """
        return pulumi.get(self, "creation_time")

    @creation_time.setter
    def creation_time(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "creation_time", value)

    @property
    @pulumi.getter(name="customEventData")
    def custom_event_data(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Information to attach to all events related to the matchmaking configuration.
        """
        return pulumi.get(self, "custom_event_data")

    @custom_event_data.setter
    def custom_event_data(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "custom_event_data", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A descriptive label that is associated with matchmaking configuration.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="flexMatchMode")
    def flex_match_mode(self) -> Optional[pulumi.Input['MatchmakingConfigurationFlexMatchMode']]:
        """
        Indicates whether this matchmaking configuration is being used with Amazon GameLift hosting or as a standalone matchmaking solution.
        """
        return pulumi.get(self, "flex_match_mode")

    @flex_match_mode.setter
    def flex_match_mode(self, value: Optional[pulumi.Input['MatchmakingConfigurationFlexMatchMode']]):
        pulumi.set(self, "flex_match_mode", value)

    @property
    @pulumi.getter(name="gameProperties")
    def game_properties(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MatchmakingConfigurationGamePropertyArgs']]]]:
        """
        A set of custom properties for a game session, formatted as key:value pairs.
        """
        return pulumi.get(self, "game_properties")

    @game_properties.setter
    def game_properties(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MatchmakingConfigurationGamePropertyArgs']]]]):
        pulumi.set(self, "game_properties", value)

    @property
    @pulumi.getter(name="gameSessionData")
    def game_session_data(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A set of custom game session properties, formatted as a single string value.
        """
        return pulumi.get(self, "game_session_data")

    @game_session_data.setter
    def game_session_data(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "game_session_data", value)

    @property
    @pulumi.getter(name="gameSessionQueueArns")
    def game_session_queue_arns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift game session queue resource and uniquely identifies it.
        """
        return pulumi.get(self, "game_session_queue_arns")

    @game_session_queue_arns.setter
    def game_session_queue_arns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "game_session_queue_arns", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A unique identifier for the matchmaking configuration.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notificationTarget")
    def notification_target(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        An SNS topic ARN that is set up to receive matchmaking notifications.
        """
        return pulumi.get(self, "notification_target")

    @notification_target.setter
    def notification_target(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "notification_target", value)

    @property
    @pulumi.getter(name="ruleSetArn")
    def rule_set_arn(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Amazon Resource Name (ARN) associated with the GameLift matchmaking rule set resource that this configuration uses.
        """
        return pulumi.get(self, "rule_set_arn")

    @rule_set_arn.setter
    def rule_set_arn(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "rule_set_arn", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:gamelift:MatchmakingConfiguration")
class MatchmakingConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acceptance_required: Optional[pulumi.Input[builtins.bool]] = None,
                 acceptance_timeout_seconds: Optional[pulumi.Input[builtins.int]] = None,
                 additional_player_count: Optional[pulumi.Input[builtins.int]] = None,
                 backfill_mode: Optional[pulumi.Input['MatchmakingConfigurationBackfillMode']] = None,
                 creation_time: Optional[pulumi.Input[builtins.str]] = None,
                 custom_event_data: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 flex_match_mode: Optional[pulumi.Input['MatchmakingConfigurationFlexMatchMode']] = None,
                 game_properties: Optional[pulumi.Input[Sequence[pulumi.Input[Union['MatchmakingConfigurationGamePropertyArgs', 'MatchmakingConfigurationGamePropertyArgsDict']]]]] = None,
                 game_session_data: Optional[pulumi.Input[builtins.str]] = None,
                 game_session_queue_arns: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 notification_target: Optional[pulumi.Input[builtins.str]] = None,
                 request_timeout_seconds: Optional[pulumi.Input[builtins.int]] = None,
                 rule_set_arn: Optional[pulumi.Input[builtins.str]] = None,
                 rule_set_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        """
        The AWS::GameLift::MatchmakingConfiguration resource creates an Amazon GameLift (GameLift) matchmaking configuration.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] acceptance_required: A flag that indicates whether a match that was created with this configuration must be accepted by the matched players
        :param pulumi.Input[builtins.int] acceptance_timeout_seconds: The length of time (in seconds) to wait for players to accept a proposed match, if acceptance is required.
        :param pulumi.Input[builtins.int] additional_player_count: The number of player slots in a match to keep open for future players.
        :param pulumi.Input['MatchmakingConfigurationBackfillMode'] backfill_mode: The method used to backfill game sessions created with this matchmaking configuration.
        :param pulumi.Input[builtins.str] creation_time: A time stamp indicating when this data object was created.
        :param pulumi.Input[builtins.str] custom_event_data: Information to attach to all events related to the matchmaking configuration.
        :param pulumi.Input[builtins.str] description: A descriptive label that is associated with matchmaking configuration.
        :param pulumi.Input['MatchmakingConfigurationFlexMatchMode'] flex_match_mode: Indicates whether this matchmaking configuration is being used with Amazon GameLift hosting or as a standalone matchmaking solution.
        :param pulumi.Input[Sequence[pulumi.Input[Union['MatchmakingConfigurationGamePropertyArgs', 'MatchmakingConfigurationGamePropertyArgsDict']]]] game_properties: A set of custom properties for a game session, formatted as key:value pairs.
        :param pulumi.Input[builtins.str] game_session_data: A set of custom game session properties, formatted as a single string value.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] game_session_queue_arns: The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift game session queue resource and uniquely identifies it.
        :param pulumi.Input[builtins.str] name: A unique identifier for the matchmaking configuration.
        :param pulumi.Input[builtins.str] notification_target: An SNS topic ARN that is set up to receive matchmaking notifications.
        :param pulumi.Input[builtins.int] request_timeout_seconds: The maximum duration, in seconds, that a matchmaking ticket can remain in process before timing out.
        :param pulumi.Input[builtins.str] rule_set_arn: The Amazon Resource Name (ARN) associated with the GameLift matchmaking rule set resource that this configuration uses.
        :param pulumi.Input[builtins.str] rule_set_name: A unique identifier for the matchmaking rule set to use with this configuration.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] tags: An array of key-value pairs to apply to this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MatchmakingConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The AWS::GameLift::MatchmakingConfiguration resource creates an Amazon GameLift (GameLift) matchmaking configuration.

        :param str resource_name: The name of the resource.
        :param MatchmakingConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MatchmakingConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acceptance_required: Optional[pulumi.Input[builtins.bool]] = None,
                 acceptance_timeout_seconds: Optional[pulumi.Input[builtins.int]] = None,
                 additional_player_count: Optional[pulumi.Input[builtins.int]] = None,
                 backfill_mode: Optional[pulumi.Input['MatchmakingConfigurationBackfillMode']] = None,
                 creation_time: Optional[pulumi.Input[builtins.str]] = None,
                 custom_event_data: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 flex_match_mode: Optional[pulumi.Input['MatchmakingConfigurationFlexMatchMode']] = None,
                 game_properties: Optional[pulumi.Input[Sequence[pulumi.Input[Union['MatchmakingConfigurationGamePropertyArgs', 'MatchmakingConfigurationGamePropertyArgsDict']]]]] = None,
                 game_session_data: Optional[pulumi.Input[builtins.str]] = None,
                 game_session_queue_arns: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 notification_target: Optional[pulumi.Input[builtins.str]] = None,
                 request_timeout_seconds: Optional[pulumi.Input[builtins.int]] = None,
                 rule_set_arn: Optional[pulumi.Input[builtins.str]] = None,
                 rule_set_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MatchmakingConfigurationArgs.__new__(MatchmakingConfigurationArgs)

            if acceptance_required is None and not opts.urn:
                raise TypeError("Missing required property 'acceptance_required'")
            __props__.__dict__["acceptance_required"] = acceptance_required
            __props__.__dict__["acceptance_timeout_seconds"] = acceptance_timeout_seconds
            __props__.__dict__["additional_player_count"] = additional_player_count
            __props__.__dict__["backfill_mode"] = backfill_mode
            __props__.__dict__["creation_time"] = creation_time
            __props__.__dict__["custom_event_data"] = custom_event_data
            __props__.__dict__["description"] = description
            __props__.__dict__["flex_match_mode"] = flex_match_mode
            __props__.__dict__["game_properties"] = game_properties
            __props__.__dict__["game_session_data"] = game_session_data
            __props__.__dict__["game_session_queue_arns"] = game_session_queue_arns
            __props__.__dict__["name"] = name
            __props__.__dict__["notification_target"] = notification_target
            if request_timeout_seconds is None and not opts.urn:
                raise TypeError("Missing required property 'request_timeout_seconds'")
            __props__.__dict__["request_timeout_seconds"] = request_timeout_seconds
            __props__.__dict__["rule_set_arn"] = rule_set_arn
            if rule_set_name is None and not opts.urn:
                raise TypeError("Missing required property 'rule_set_name'")
            __props__.__dict__["rule_set_name"] = rule_set_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["name"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(MatchmakingConfiguration, __self__).__init__(
            'aws-native:gamelift:MatchmakingConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'MatchmakingConfiguration':
        """
        Get an existing MatchmakingConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = MatchmakingConfigurationArgs.__new__(MatchmakingConfigurationArgs)

        __props__.__dict__["acceptance_required"] = None
        __props__.__dict__["acceptance_timeout_seconds"] = None
        __props__.__dict__["additional_player_count"] = None
        __props__.__dict__["arn"] = None
        __props__.__dict__["backfill_mode"] = None
        __props__.__dict__["creation_time"] = None
        __props__.__dict__["custom_event_data"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["flex_match_mode"] = None
        __props__.__dict__["game_properties"] = None
        __props__.__dict__["game_session_data"] = None
        __props__.__dict__["game_session_queue_arns"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["notification_target"] = None
        __props__.__dict__["request_timeout_seconds"] = None
        __props__.__dict__["rule_set_arn"] = None
        __props__.__dict__["rule_set_name"] = None
        __props__.__dict__["tags"] = None
        return MatchmakingConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="acceptanceRequired")
    def acceptance_required(self) -> pulumi.Output[builtins.bool]:
        """
        A flag that indicates whether a match that was created with this configuration must be accepted by the matched players
        """
        return pulumi.get(self, "acceptance_required")

    @property
    @pulumi.getter(name="acceptanceTimeoutSeconds")
    def acceptance_timeout_seconds(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        The length of time (in seconds) to wait for players to accept a proposed match, if acceptance is required.
        """
        return pulumi.get(self, "acceptance_timeout_seconds")

    @property
    @pulumi.getter(name="additionalPlayerCount")
    def additional_player_count(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        The number of player slots in a match to keep open for future players.
        """
        return pulumi.get(self, "additional_player_count")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[builtins.str]:
        """
        The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift matchmaking configuration resource and uniquely identifies it.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="backfillMode")
    def backfill_mode(self) -> pulumi.Output[Optional['MatchmakingConfigurationBackfillMode']]:
        """
        The method used to backfill game sessions created with this matchmaking configuration.
        """
        return pulumi.get(self, "backfill_mode")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        A time stamp indicating when this data object was created.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="customEventData")
    def custom_event_data(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Information to attach to all events related to the matchmaking configuration.
        """
        return pulumi.get(self, "custom_event_data")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        A descriptive label that is associated with matchmaking configuration.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="flexMatchMode")
    def flex_match_mode(self) -> pulumi.Output[Optional['MatchmakingConfigurationFlexMatchMode']]:
        """
        Indicates whether this matchmaking configuration is being used with Amazon GameLift hosting or as a standalone matchmaking solution.
        """
        return pulumi.get(self, "flex_match_mode")

    @property
    @pulumi.getter(name="gameProperties")
    def game_properties(self) -> pulumi.Output[Optional[Sequence['outputs.MatchmakingConfigurationGameProperty']]]:
        """
        A set of custom properties for a game session, formatted as key:value pairs.
        """
        return pulumi.get(self, "game_properties")

    @property
    @pulumi.getter(name="gameSessionData")
    def game_session_data(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        A set of custom game session properties, formatted as a single string value.
        """
        return pulumi.get(self, "game_session_data")

    @property
    @pulumi.getter(name="gameSessionQueueArns")
    def game_session_queue_arns(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift game session queue resource and uniquely identifies it.
        """
        return pulumi.get(self, "game_session_queue_arns")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        A unique identifier for the matchmaking configuration.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationTarget")
    def notification_target(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        An SNS topic ARN that is set up to receive matchmaking notifications.
        """
        return pulumi.get(self, "notification_target")

    @property
    @pulumi.getter(name="requestTimeoutSeconds")
    def request_timeout_seconds(self) -> pulumi.Output[builtins.int]:
        """
        The maximum duration, in seconds, that a matchmaking ticket can remain in process before timing out.
        """
        return pulumi.get(self, "request_timeout_seconds")

    @property
    @pulumi.getter(name="ruleSetArn")
    def rule_set_arn(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The Amazon Resource Name (ARN) associated with the GameLift matchmaking rule set resource that this configuration uses.
        """
        return pulumi.get(self, "rule_set_arn")

    @property
    @pulumi.getter(name="ruleSetName")
    def rule_set_name(self) -> pulumi.Output[builtins.str]:
        """
        A unique identifier for the matchmaking rule set to use with this configuration.
        """
        return pulumi.get(self, "rule_set_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

