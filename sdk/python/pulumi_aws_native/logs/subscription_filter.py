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

__all__ = ['SubscriptionFilterArgs', 'SubscriptionFilter']

@pulumi.input_type
class SubscriptionFilterArgs:
    def __init__(__self__, *,
                 destination_arn: pulumi.Input[str],
                 filter_pattern: pulumi.Input[str],
                 log_group_name: pulumi.Input[str],
                 distribution: Optional[pulumi.Input['SubscriptionFilterDistribution']] = None,
                 filter_name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SubscriptionFilter resource.
        :param pulumi.Input[str] destination_arn: The Amazon Resource Name (ARN) of the destination.
        :param pulumi.Input[str] filter_pattern: The filtering expressions that restrict what gets delivered to the destination AWS resource.
        :param pulumi.Input[str] log_group_name: Existing log group that you want to associate with this filter.
        :param pulumi.Input['SubscriptionFilterDistribution'] distribution: The method used to distribute log data to the destination. By default, log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream.
        :param pulumi.Input[str] filter_name: The name of the filter generated by resource.
        :param pulumi.Input[str] role_arn: The ARN of an IAM role that grants CloudWatch Logs permissions to deliver ingested log events to the destination stream. You don't need to provide the ARN when you are working with a logical destination for cross-account delivery.
        """
        pulumi.set(__self__, "destination_arn", destination_arn)
        pulumi.set(__self__, "filter_pattern", filter_pattern)
        pulumi.set(__self__, "log_group_name", log_group_name)
        if distribution is not None:
            pulumi.set(__self__, "distribution", distribution)
        if filter_name is not None:
            pulumi.set(__self__, "filter_name", filter_name)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)

    @property
    @pulumi.getter(name="destinationArn")
    def destination_arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN) of the destination.
        """
        return pulumi.get(self, "destination_arn")

    @destination_arn.setter
    def destination_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_arn", value)

    @property
    @pulumi.getter(name="filterPattern")
    def filter_pattern(self) -> pulumi.Input[str]:
        """
        The filtering expressions that restrict what gets delivered to the destination AWS resource.
        """
        return pulumi.get(self, "filter_pattern")

    @filter_pattern.setter
    def filter_pattern(self, value: pulumi.Input[str]):
        pulumi.set(self, "filter_pattern", value)

    @property
    @pulumi.getter(name="logGroupName")
    def log_group_name(self) -> pulumi.Input[str]:
        """
        Existing log group that you want to associate with this filter.
        """
        return pulumi.get(self, "log_group_name")

    @log_group_name.setter
    def log_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "log_group_name", value)

    @property
    @pulumi.getter
    def distribution(self) -> Optional[pulumi.Input['SubscriptionFilterDistribution']]:
        """
        The method used to distribute log data to the destination. By default, log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream.
        """
        return pulumi.get(self, "distribution")

    @distribution.setter
    def distribution(self, value: Optional[pulumi.Input['SubscriptionFilterDistribution']]):
        pulumi.set(self, "distribution", value)

    @property
    @pulumi.getter(name="filterName")
    def filter_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the filter generated by resource.
        """
        return pulumi.get(self, "filter_name")

    @filter_name.setter
    def filter_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter_name", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of an IAM role that grants CloudWatch Logs permissions to deliver ingested log events to the destination stream. You don't need to provide the ARN when you are working with a logical destination for cross-account delivery.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)


class SubscriptionFilter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_arn: Optional[pulumi.Input[str]] = None,
                 distribution: Optional[pulumi.Input['SubscriptionFilterDistribution']] = None,
                 filter_name: Optional[pulumi.Input[str]] = None,
                 filter_pattern: Optional[pulumi.Input[str]] = None,
                 log_group_name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        An example resource schema demonstrating some basic constructs and validation rules.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_arn: The Amazon Resource Name (ARN) of the destination.
        :param pulumi.Input['SubscriptionFilterDistribution'] distribution: The method used to distribute log data to the destination. By default, log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream.
        :param pulumi.Input[str] filter_name: The name of the filter generated by resource.
        :param pulumi.Input[str] filter_pattern: The filtering expressions that restrict what gets delivered to the destination AWS resource.
        :param pulumi.Input[str] log_group_name: Existing log group that you want to associate with this filter.
        :param pulumi.Input[str] role_arn: The ARN of an IAM role that grants CloudWatch Logs permissions to deliver ingested log events to the destination stream. You don't need to provide the ARN when you are working with a logical destination for cross-account delivery.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SubscriptionFilterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An example resource schema demonstrating some basic constructs and validation rules.

        :param str resource_name: The name of the resource.
        :param SubscriptionFilterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SubscriptionFilterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_arn: Optional[pulumi.Input[str]] = None,
                 distribution: Optional[pulumi.Input['SubscriptionFilterDistribution']] = None,
                 filter_name: Optional[pulumi.Input[str]] = None,
                 filter_pattern: Optional[pulumi.Input[str]] = None,
                 log_group_name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SubscriptionFilterArgs.__new__(SubscriptionFilterArgs)

            if destination_arn is None and not opts.urn:
                raise TypeError("Missing required property 'destination_arn'")
            __props__.__dict__["destination_arn"] = destination_arn
            __props__.__dict__["distribution"] = distribution
            __props__.__dict__["filter_name"] = filter_name
            if filter_pattern is None and not opts.urn:
                raise TypeError("Missing required property 'filter_pattern'")
            __props__.__dict__["filter_pattern"] = filter_pattern
            if log_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'log_group_name'")
            __props__.__dict__["log_group_name"] = log_group_name
            __props__.__dict__["role_arn"] = role_arn
        super(SubscriptionFilter, __self__).__init__(
            'aws-native:logs:SubscriptionFilter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SubscriptionFilter':
        """
        Get an existing SubscriptionFilter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SubscriptionFilterArgs.__new__(SubscriptionFilterArgs)

        __props__.__dict__["destination_arn"] = None
        __props__.__dict__["distribution"] = None
        __props__.__dict__["filter_name"] = None
        __props__.__dict__["filter_pattern"] = None
        __props__.__dict__["log_group_name"] = None
        __props__.__dict__["role_arn"] = None
        return SubscriptionFilter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="destinationArn")
    def destination_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the destination.
        """
        return pulumi.get(self, "destination_arn")

    @property
    @pulumi.getter
    def distribution(self) -> pulumi.Output[Optional['SubscriptionFilterDistribution']]:
        """
        The method used to distribute log data to the destination. By default, log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream.
        """
        return pulumi.get(self, "distribution")

    @property
    @pulumi.getter(name="filterName")
    def filter_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the filter generated by resource.
        """
        return pulumi.get(self, "filter_name")

    @property
    @pulumi.getter(name="filterPattern")
    def filter_pattern(self) -> pulumi.Output[str]:
        """
        The filtering expressions that restrict what gets delivered to the destination AWS resource.
        """
        return pulumi.get(self, "filter_pattern")

    @property
    @pulumi.getter(name="logGroupName")
    def log_group_name(self) -> pulumi.Output[str]:
        """
        Existing log group that you want to associate with this filter.
        """
        return pulumi.get(self, "log_group_name")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The ARN of an IAM role that grants CloudWatch Logs permissions to deliver ingested log events to the destination stream. You don't need to provide the ARN when you are working with a logical destination for cross-account delivery.
        """
        return pulumi.get(self, "role_arn")
