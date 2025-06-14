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

__all__ = ['StreamArgs', 'Stream']

@pulumi.input_type
class StreamArgs:
    def __init__(__self__, *,
                 desired_shard_level_metrics: Optional[pulumi.Input[Sequence[pulumi.Input['StreamEnhancedMetric']]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 retention_period_hours: Optional[pulumi.Input[builtins.int]] = None,
                 shard_count: Optional[pulumi.Input[builtins.int]] = None,
                 stream_encryption: Optional[pulumi.Input['StreamEncryptionArgs']] = None,
                 stream_mode_details: Optional[pulumi.Input['StreamModeDetailsArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a Stream resource.
        :param pulumi.Input[Sequence[pulumi.Input['StreamEnhancedMetric']]] desired_shard_level_metrics: The final list of shard-level metrics
        :param pulumi.Input[builtins.str] name: The name of the Kinesis stream.
        :param pulumi.Input[builtins.int] retention_period_hours: The number of hours for the data records that are stored in shards to remain accessible.
        :param pulumi.Input[builtins.int] shard_count: The number of shards that the stream uses. Required when StreamMode = PROVISIONED is passed.
        :param pulumi.Input['StreamEncryptionArgs'] stream_encryption: When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream.
        :param pulumi.Input['StreamModeDetailsArgs'] stream_mode_details: The mode in which the stream is running.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: An arbitrary set of tags (key–value pairs) to associate with the Kinesis stream.
        """
        if desired_shard_level_metrics is not None:
            pulumi.set(__self__, "desired_shard_level_metrics", desired_shard_level_metrics)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if retention_period_hours is not None:
            pulumi.set(__self__, "retention_period_hours", retention_period_hours)
        if shard_count is not None:
            pulumi.set(__self__, "shard_count", shard_count)
        if stream_encryption is not None:
            pulumi.set(__self__, "stream_encryption", stream_encryption)
        if stream_mode_details is not None:
            pulumi.set(__self__, "stream_mode_details", stream_mode_details)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="desiredShardLevelMetrics")
    def desired_shard_level_metrics(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StreamEnhancedMetric']]]]:
        """
        The final list of shard-level metrics
        """
        return pulumi.get(self, "desired_shard_level_metrics")

    @desired_shard_level_metrics.setter
    def desired_shard_level_metrics(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StreamEnhancedMetric']]]]):
        pulumi.set(self, "desired_shard_level_metrics", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the Kinesis stream.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="retentionPeriodHours")
    def retention_period_hours(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The number of hours for the data records that are stored in shards to remain accessible.
        """
        return pulumi.get(self, "retention_period_hours")

    @retention_period_hours.setter
    def retention_period_hours(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "retention_period_hours", value)

    @property
    @pulumi.getter(name="shardCount")
    def shard_count(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The number of shards that the stream uses. Required when StreamMode = PROVISIONED is passed.
        """
        return pulumi.get(self, "shard_count")

    @shard_count.setter
    def shard_count(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "shard_count", value)

    @property
    @pulumi.getter(name="streamEncryption")
    def stream_encryption(self) -> Optional[pulumi.Input['StreamEncryptionArgs']]:
        """
        When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream.
        """
        return pulumi.get(self, "stream_encryption")

    @stream_encryption.setter
    def stream_encryption(self, value: Optional[pulumi.Input['StreamEncryptionArgs']]):
        pulumi.set(self, "stream_encryption", value)

    @property
    @pulumi.getter(name="streamModeDetails")
    def stream_mode_details(self) -> Optional[pulumi.Input['StreamModeDetailsArgs']]:
        """
        The mode in which the stream is running.
        """
        return pulumi.get(self, "stream_mode_details")

    @stream_mode_details.setter
    def stream_mode_details(self, value: Optional[pulumi.Input['StreamModeDetailsArgs']]):
        pulumi.set(self, "stream_mode_details", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        An arbitrary set of tags (key–value pairs) to associate with the Kinesis stream.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:kinesis:Stream")
class Stream(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 desired_shard_level_metrics: Optional[pulumi.Input[Sequence[pulumi.Input['StreamEnhancedMetric']]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 retention_period_hours: Optional[pulumi.Input[builtins.int]] = None,
                 shard_count: Optional[pulumi.Input[builtins.int]] = None,
                 stream_encryption: Optional[pulumi.Input[Union['StreamEncryptionArgs', 'StreamEncryptionArgsDict']]] = None,
                 stream_mode_details: Optional[pulumi.Input[Union['StreamModeDetailsArgs', 'StreamModeDetailsArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Kinesis::Stream

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input['StreamEnhancedMetric']]] desired_shard_level_metrics: The final list of shard-level metrics
        :param pulumi.Input[builtins.str] name: The name of the Kinesis stream.
        :param pulumi.Input[builtins.int] retention_period_hours: The number of hours for the data records that are stored in shards to remain accessible.
        :param pulumi.Input[builtins.int] shard_count: The number of shards that the stream uses. Required when StreamMode = PROVISIONED is passed.
        :param pulumi.Input[Union['StreamEncryptionArgs', 'StreamEncryptionArgsDict']] stream_encryption: When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream.
        :param pulumi.Input[Union['StreamModeDetailsArgs', 'StreamModeDetailsArgsDict']] stream_mode_details: The mode in which the stream is running.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] tags: An arbitrary set of tags (key–value pairs) to associate with the Kinesis stream.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[StreamArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Kinesis::Stream

        :param str resource_name: The name of the resource.
        :param StreamArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StreamArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 desired_shard_level_metrics: Optional[pulumi.Input[Sequence[pulumi.Input['StreamEnhancedMetric']]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 retention_period_hours: Optional[pulumi.Input[builtins.int]] = None,
                 shard_count: Optional[pulumi.Input[builtins.int]] = None,
                 stream_encryption: Optional[pulumi.Input[Union['StreamEncryptionArgs', 'StreamEncryptionArgsDict']]] = None,
                 stream_mode_details: Optional[pulumi.Input[Union['StreamModeDetailsArgs', 'StreamModeDetailsArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StreamArgs.__new__(StreamArgs)

            __props__.__dict__["desired_shard_level_metrics"] = desired_shard_level_metrics
            __props__.__dict__["name"] = name
            __props__.__dict__["retention_period_hours"] = retention_period_hours
            __props__.__dict__["shard_count"] = shard_count
            __props__.__dict__["stream_encryption"] = stream_encryption
            __props__.__dict__["stream_mode_details"] = stream_mode_details
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["name"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Stream, __self__).__init__(
            'aws-native:kinesis:Stream',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Stream':
        """
        Get an existing Stream resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = StreamArgs.__new__(StreamArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["desired_shard_level_metrics"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["retention_period_hours"] = None
        __props__.__dict__["shard_count"] = None
        __props__.__dict__["stream_encryption"] = None
        __props__.__dict__["stream_mode_details"] = None
        __props__.__dict__["tags"] = None
        return Stream(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[builtins.str]:
        """
        The Amazon resource name (ARN) of the Kinesis stream
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="desiredShardLevelMetrics")
    def desired_shard_level_metrics(self) -> pulumi.Output[Optional[Sequence['StreamEnhancedMetric']]]:
        """
        The final list of shard-level metrics
        """
        return pulumi.get(self, "desired_shard_level_metrics")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The name of the Kinesis stream.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="retentionPeriodHours")
    def retention_period_hours(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        The number of hours for the data records that are stored in shards to remain accessible.
        """
        return pulumi.get(self, "retention_period_hours")

    @property
    @pulumi.getter(name="shardCount")
    def shard_count(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        The number of shards that the stream uses. Required when StreamMode = PROVISIONED is passed.
        """
        return pulumi.get(self, "shard_count")

    @property
    @pulumi.getter(name="streamEncryption")
    def stream_encryption(self) -> pulumi.Output[Optional['outputs.StreamEncryption']]:
        """
        When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream.
        """
        return pulumi.get(self, "stream_encryption")

    @property
    @pulumi.getter(name="streamModeDetails")
    def stream_mode_details(self) -> pulumi.Output[Optional['outputs.StreamModeDetails']]:
        """
        The mode in which the stream is running.
        """
        return pulumi.get(self, "stream_mode_details")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        An arbitrary set of tags (key–value pairs) to associate with the Kinesis stream.
        """
        return pulumi.get(self, "tags")

