# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = ['SlackChannelConfigurationArgs', 'SlackChannelConfiguration']

@pulumi.input_type
class SlackChannelConfigurationArgs:
    def __init__(__self__, *,
                 channel_id: pulumi.Input[str],
                 channel_role_arn: pulumi.Input[str],
                 notify_on_case_severity: pulumi.Input['SlackChannelConfigurationNotifyOnCaseSeverity'],
                 team_id: pulumi.Input[str],
                 channel_name: Optional[pulumi.Input[str]] = None,
                 notify_on_add_correspondence_to_case: Optional[pulumi.Input[bool]] = None,
                 notify_on_create_or_reopen_case: Optional[pulumi.Input[bool]] = None,
                 notify_on_resolve_case: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a SlackChannelConfiguration resource.
        :param pulumi.Input[str] channel_id: The channel ID in Slack, which identifies a channel within a workspace.
        :param pulumi.Input[str] channel_role_arn: The Amazon Resource Name (ARN) of an IAM role that grants the AWS Support App access to perform operations for AWS services.
        :param pulumi.Input['SlackChannelConfigurationNotifyOnCaseSeverity'] notify_on_case_severity: The severity level of a support case that a customer wants to get notified for.
        :param pulumi.Input[str] team_id: The team ID in Slack, which uniquely identifies a workspace.
        :param pulumi.Input[str] channel_name: The channel name in Slack.
        :param pulumi.Input[bool] notify_on_add_correspondence_to_case: Whether to notify when a correspondence is added to a case.
        :param pulumi.Input[bool] notify_on_create_or_reopen_case: Whether to notify when a case is created or reopened.
        :param pulumi.Input[bool] notify_on_resolve_case: Whether to notify when a case is resolved.
        """
        pulumi.set(__self__, "channel_id", channel_id)
        pulumi.set(__self__, "channel_role_arn", channel_role_arn)
        pulumi.set(__self__, "notify_on_case_severity", notify_on_case_severity)
        pulumi.set(__self__, "team_id", team_id)
        if channel_name is not None:
            pulumi.set(__self__, "channel_name", channel_name)
        if notify_on_add_correspondence_to_case is not None:
            pulumi.set(__self__, "notify_on_add_correspondence_to_case", notify_on_add_correspondence_to_case)
        if notify_on_create_or_reopen_case is not None:
            pulumi.set(__self__, "notify_on_create_or_reopen_case", notify_on_create_or_reopen_case)
        if notify_on_resolve_case is not None:
            pulumi.set(__self__, "notify_on_resolve_case", notify_on_resolve_case)

    @property
    @pulumi.getter(name="channelId")
    def channel_id(self) -> pulumi.Input[str]:
        """
        The channel ID in Slack, which identifies a channel within a workspace.
        """
        return pulumi.get(self, "channel_id")

    @channel_id.setter
    def channel_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "channel_id", value)

    @property
    @pulumi.getter(name="channelRoleArn")
    def channel_role_arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN) of an IAM role that grants the AWS Support App access to perform operations for AWS services.
        """
        return pulumi.get(self, "channel_role_arn")

    @channel_role_arn.setter
    def channel_role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "channel_role_arn", value)

    @property
    @pulumi.getter(name="notifyOnCaseSeverity")
    def notify_on_case_severity(self) -> pulumi.Input['SlackChannelConfigurationNotifyOnCaseSeverity']:
        """
        The severity level of a support case that a customer wants to get notified for.
        """
        return pulumi.get(self, "notify_on_case_severity")

    @notify_on_case_severity.setter
    def notify_on_case_severity(self, value: pulumi.Input['SlackChannelConfigurationNotifyOnCaseSeverity']):
        pulumi.set(self, "notify_on_case_severity", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> pulumi.Input[str]:
        """
        The team ID in Slack, which uniquely identifies a workspace.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "team_id", value)

    @property
    @pulumi.getter(name="channelName")
    def channel_name(self) -> Optional[pulumi.Input[str]]:
        """
        The channel name in Slack.
        """
        return pulumi.get(self, "channel_name")

    @channel_name.setter
    def channel_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "channel_name", value)

    @property
    @pulumi.getter(name="notifyOnAddCorrespondenceToCase")
    def notify_on_add_correspondence_to_case(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to notify when a correspondence is added to a case.
        """
        return pulumi.get(self, "notify_on_add_correspondence_to_case")

    @notify_on_add_correspondence_to_case.setter
    def notify_on_add_correspondence_to_case(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "notify_on_add_correspondence_to_case", value)

    @property
    @pulumi.getter(name="notifyOnCreateOrReopenCase")
    def notify_on_create_or_reopen_case(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to notify when a case is created or reopened.
        """
        return pulumi.get(self, "notify_on_create_or_reopen_case")

    @notify_on_create_or_reopen_case.setter
    def notify_on_create_or_reopen_case(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "notify_on_create_or_reopen_case", value)

    @property
    @pulumi.getter(name="notifyOnResolveCase")
    def notify_on_resolve_case(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to notify when a case is resolved.
        """
        return pulumi.get(self, "notify_on_resolve_case")

    @notify_on_resolve_case.setter
    def notify_on_resolve_case(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "notify_on_resolve_case", value)


class SlackChannelConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 channel_id: Optional[pulumi.Input[str]] = None,
                 channel_name: Optional[pulumi.Input[str]] = None,
                 channel_role_arn: Optional[pulumi.Input[str]] = None,
                 notify_on_add_correspondence_to_case: Optional[pulumi.Input[bool]] = None,
                 notify_on_case_severity: Optional[pulumi.Input['SlackChannelConfigurationNotifyOnCaseSeverity']] = None,
                 notify_on_create_or_reopen_case: Optional[pulumi.Input[bool]] = None,
                 notify_on_resolve_case: Optional[pulumi.Input[bool]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        An AWS Support App resource that creates, updates, lists and deletes Slack channel configurations.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] channel_id: The channel ID in Slack, which identifies a channel within a workspace.
        :param pulumi.Input[str] channel_name: The channel name in Slack.
        :param pulumi.Input[str] channel_role_arn: The Amazon Resource Name (ARN) of an IAM role that grants the AWS Support App access to perform operations for AWS services.
        :param pulumi.Input[bool] notify_on_add_correspondence_to_case: Whether to notify when a correspondence is added to a case.
        :param pulumi.Input['SlackChannelConfigurationNotifyOnCaseSeverity'] notify_on_case_severity: The severity level of a support case that a customer wants to get notified for.
        :param pulumi.Input[bool] notify_on_create_or_reopen_case: Whether to notify when a case is created or reopened.
        :param pulumi.Input[bool] notify_on_resolve_case: Whether to notify when a case is resolved.
        :param pulumi.Input[str] team_id: The team ID in Slack, which uniquely identifies a workspace.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SlackChannelConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An AWS Support App resource that creates, updates, lists and deletes Slack channel configurations.

        :param str resource_name: The name of the resource.
        :param SlackChannelConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SlackChannelConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 channel_id: Optional[pulumi.Input[str]] = None,
                 channel_name: Optional[pulumi.Input[str]] = None,
                 channel_role_arn: Optional[pulumi.Input[str]] = None,
                 notify_on_add_correspondence_to_case: Optional[pulumi.Input[bool]] = None,
                 notify_on_case_severity: Optional[pulumi.Input['SlackChannelConfigurationNotifyOnCaseSeverity']] = None,
                 notify_on_create_or_reopen_case: Optional[pulumi.Input[bool]] = None,
                 notify_on_resolve_case: Optional[pulumi.Input[bool]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SlackChannelConfigurationArgs.__new__(SlackChannelConfigurationArgs)

            if channel_id is None and not opts.urn:
                raise TypeError("Missing required property 'channel_id'")
            __props__.__dict__["channel_id"] = channel_id
            __props__.__dict__["channel_name"] = channel_name
            if channel_role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'channel_role_arn'")
            __props__.__dict__["channel_role_arn"] = channel_role_arn
            __props__.__dict__["notify_on_add_correspondence_to_case"] = notify_on_add_correspondence_to_case
            if notify_on_case_severity is None and not opts.urn:
                raise TypeError("Missing required property 'notify_on_case_severity'")
            __props__.__dict__["notify_on_case_severity"] = notify_on_case_severity
            __props__.__dict__["notify_on_create_or_reopen_case"] = notify_on_create_or_reopen_case
            __props__.__dict__["notify_on_resolve_case"] = notify_on_resolve_case
            if team_id is None and not opts.urn:
                raise TypeError("Missing required property 'team_id'")
            __props__.__dict__["team_id"] = team_id
        super(SlackChannelConfiguration, __self__).__init__(
            'aws-native:supportapp:SlackChannelConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SlackChannelConfiguration':
        """
        Get an existing SlackChannelConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SlackChannelConfigurationArgs.__new__(SlackChannelConfigurationArgs)

        __props__.__dict__["channel_id"] = None
        __props__.__dict__["channel_name"] = None
        __props__.__dict__["channel_role_arn"] = None
        __props__.__dict__["notify_on_add_correspondence_to_case"] = None
        __props__.__dict__["notify_on_case_severity"] = None
        __props__.__dict__["notify_on_create_or_reopen_case"] = None
        __props__.__dict__["notify_on_resolve_case"] = None
        __props__.__dict__["team_id"] = None
        return SlackChannelConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="channelId")
    def channel_id(self) -> pulumi.Output[str]:
        """
        The channel ID in Slack, which identifies a channel within a workspace.
        """
        return pulumi.get(self, "channel_id")

    @property
    @pulumi.getter(name="channelName")
    def channel_name(self) -> pulumi.Output[Optional[str]]:
        """
        The channel name in Slack.
        """
        return pulumi.get(self, "channel_name")

    @property
    @pulumi.getter(name="channelRoleArn")
    def channel_role_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of an IAM role that grants the AWS Support App access to perform operations for AWS services.
        """
        return pulumi.get(self, "channel_role_arn")

    @property
    @pulumi.getter(name="notifyOnAddCorrespondenceToCase")
    def notify_on_add_correspondence_to_case(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to notify when a correspondence is added to a case.
        """
        return pulumi.get(self, "notify_on_add_correspondence_to_case")

    @property
    @pulumi.getter(name="notifyOnCaseSeverity")
    def notify_on_case_severity(self) -> pulumi.Output['SlackChannelConfigurationNotifyOnCaseSeverity']:
        """
        The severity level of a support case that a customer wants to get notified for.
        """
        return pulumi.get(self, "notify_on_case_severity")

    @property
    @pulumi.getter(name="notifyOnCreateOrReopenCase")
    def notify_on_create_or_reopen_case(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to notify when a case is created or reopened.
        """
        return pulumi.get(self, "notify_on_create_or_reopen_case")

    @property
    @pulumi.getter(name="notifyOnResolveCase")
    def notify_on_resolve_case(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to notify when a case is resolved.
        """
        return pulumi.get(self, "notify_on_resolve_case")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> pulumi.Output[str]:
        """
        The team ID in Slack, which uniquely identifies a workspace.
        """
        return pulumi.get(self, "team_id")
