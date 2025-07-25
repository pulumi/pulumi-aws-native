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

__all__ = ['TrailArgs', 'Trail']

@pulumi.input_type
class TrailArgs:
    def __init__(__self__, *,
                 is_logging: pulumi.Input[builtins.bool],
                 s3_bucket_name: pulumi.Input[builtins.str],
                 advanced_event_selectors: Optional[pulumi.Input[Sequence[pulumi.Input['TrailAdvancedEventSelectorArgs']]]] = None,
                 cloud_watch_logs_log_group_arn: Optional[pulumi.Input[builtins.str]] = None,
                 cloud_watch_logs_role_arn: Optional[pulumi.Input[builtins.str]] = None,
                 enable_log_file_validation: Optional[pulumi.Input[builtins.bool]] = None,
                 event_selectors: Optional[pulumi.Input[Sequence[pulumi.Input['TrailEventSelectorArgs']]]] = None,
                 include_global_service_events: Optional[pulumi.Input[builtins.bool]] = None,
                 insight_selectors: Optional[pulumi.Input[Sequence[pulumi.Input['TrailInsightSelectorArgs']]]] = None,
                 is_multi_region_trail: Optional[pulumi.Input[builtins.bool]] = None,
                 is_organization_trail: Optional[pulumi.Input[builtins.bool]] = None,
                 kms_key_id: Optional[pulumi.Input[builtins.str]] = None,
                 s3_key_prefix: Optional[pulumi.Input[builtins.str]] = None,
                 sns_topic_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None,
                 trail_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Trail resource.
        :param pulumi.Input[builtins.bool] is_logging: Whether the CloudTrail is currently logging AWS API calls.
        :param pulumi.Input[builtins.str] s3_bucket_name: Specifies the name of the Amazon S3 bucket designated for publishing log files. See Amazon S3 Bucket Naming Requirements.
        :param pulumi.Input[Sequence[pulumi.Input['TrailAdvancedEventSelectorArgs']]] advanced_event_selectors: The advanced event selectors that were used to select events for the data store.
        :param pulumi.Input[builtins.str] cloud_watch_logs_log_group_arn: Specifies a log group name using an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs will be delivered. Not required unless you specify CloudWatchLogsRoleArn.
        :param pulumi.Input[builtins.str] cloud_watch_logs_role_arn: Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's log group.
        :param pulumi.Input[builtins.bool] enable_log_file_validation: Specifies whether log file validation is enabled. The default is false.
        :param pulumi.Input[Sequence[pulumi.Input['TrailEventSelectorArgs']]] event_selectors: Use event selectors to further specify the management and data event settings for your trail. By default, trails created without specific event selectors will be configured to log all read and write management events, and no data events. When an event occurs in your account, CloudTrail evaluates the event selector for all trails. For each trail, if the event matches any event selector, the trail processes and logs the event. If the event doesn't match any event selector, the trail doesn't log the event. You can configure up to five event selectors for a trail.
        :param pulumi.Input[builtins.bool] include_global_service_events: Specifies whether the trail is publishing events from global services such as IAM to the log files.
        :param pulumi.Input[Sequence[pulumi.Input['TrailInsightSelectorArgs']]] insight_selectors: Lets you enable Insights event logging by specifying the Insights selectors that you want to enable on an existing trail.
        :param pulumi.Input[builtins.bool] is_multi_region_trail: Specifies whether the trail applies only to the current region or to all regions. The default is false. If the trail exists only in the current region and this value is set to true, shadow trails (replications of the trail) will be created in the other regions. If the trail exists in all regions and this value is set to false, the trail will remain in the region where it was created, and its shadow trails in other regions will be deleted. As a best practice, consider using trails that log events in all regions.
        :param pulumi.Input[builtins.bool] is_organization_trail: Specifies whether the trail is created for all accounts in an organization in AWS Organizations, or only for the current AWS account. The default is false, and cannot be true unless the call is made on behalf of an AWS account that is the master account for an organization in AWS Organizations.
        :param pulumi.Input[builtins.str] kms_key_id: Specifies the KMS key ID to use to encrypt the logs delivered by CloudTrail. The value can be an alias name prefixed by 'alias/', a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
        :param pulumi.Input[builtins.str] s3_key_prefix: Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery. For more information, see Finding Your CloudTrail Log Files. The maximum length is 200 characters.
        :param pulumi.Input[builtins.str] sns_topic_name: Specifies the name of the Amazon SNS topic defined for notification of log file delivery. The maximum length is 256 characters.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: A custom set of tags (key-value pairs) for this trail.
        :param pulumi.Input[builtins.str] trail_name: Specifies the name of the trail. The name must meet the following requirements:
               
               - Contain only ASCII letters (a-z, A-Z), numbers (0-9), periods (.), underscores (_), or dashes (-)
               - Start with a letter or number, and end with a letter or number
               - Be between 3 and 128 characters
               - Have no adjacent periods, underscores or dashes. Names like `my-_namespace` and `my--namespace` are not valid.
               - Not be in IP address format (for example, 192.168.5.4)
        """
        pulumi.set(__self__, "is_logging", is_logging)
        pulumi.set(__self__, "s3_bucket_name", s3_bucket_name)
        if advanced_event_selectors is not None:
            pulumi.set(__self__, "advanced_event_selectors", advanced_event_selectors)
        if cloud_watch_logs_log_group_arn is not None:
            pulumi.set(__self__, "cloud_watch_logs_log_group_arn", cloud_watch_logs_log_group_arn)
        if cloud_watch_logs_role_arn is not None:
            pulumi.set(__self__, "cloud_watch_logs_role_arn", cloud_watch_logs_role_arn)
        if enable_log_file_validation is not None:
            pulumi.set(__self__, "enable_log_file_validation", enable_log_file_validation)
        if event_selectors is not None:
            pulumi.set(__self__, "event_selectors", event_selectors)
        if include_global_service_events is not None:
            pulumi.set(__self__, "include_global_service_events", include_global_service_events)
        if insight_selectors is not None:
            pulumi.set(__self__, "insight_selectors", insight_selectors)
        if is_multi_region_trail is not None:
            pulumi.set(__self__, "is_multi_region_trail", is_multi_region_trail)
        if is_organization_trail is not None:
            pulumi.set(__self__, "is_organization_trail", is_organization_trail)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if s3_key_prefix is not None:
            pulumi.set(__self__, "s3_key_prefix", s3_key_prefix)
        if sns_topic_name is not None:
            pulumi.set(__self__, "sns_topic_name", sns_topic_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if trail_name is not None:
            pulumi.set(__self__, "trail_name", trail_name)

    @property
    @pulumi.getter(name="isLogging")
    def is_logging(self) -> pulumi.Input[builtins.bool]:
        """
        Whether the CloudTrail is currently logging AWS API calls.
        """
        return pulumi.get(self, "is_logging")

    @is_logging.setter
    def is_logging(self, value: pulumi.Input[builtins.bool]):
        pulumi.set(self, "is_logging", value)

    @property
    @pulumi.getter(name="s3BucketName")
    def s3_bucket_name(self) -> pulumi.Input[builtins.str]:
        """
        Specifies the name of the Amazon S3 bucket designated for publishing log files. See Amazon S3 Bucket Naming Requirements.
        """
        return pulumi.get(self, "s3_bucket_name")

    @s3_bucket_name.setter
    def s3_bucket_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "s3_bucket_name", value)

    @property
    @pulumi.getter(name="advancedEventSelectors")
    def advanced_event_selectors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TrailAdvancedEventSelectorArgs']]]]:
        """
        The advanced event selectors that were used to select events for the data store.
        """
        return pulumi.get(self, "advanced_event_selectors")

    @advanced_event_selectors.setter
    def advanced_event_selectors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TrailAdvancedEventSelectorArgs']]]]):
        pulumi.set(self, "advanced_event_selectors", value)

    @property
    @pulumi.getter(name="cloudWatchLogsLogGroupArn")
    def cloud_watch_logs_log_group_arn(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies a log group name using an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs will be delivered. Not required unless you specify CloudWatchLogsRoleArn.
        """
        return pulumi.get(self, "cloud_watch_logs_log_group_arn")

    @cloud_watch_logs_log_group_arn.setter
    def cloud_watch_logs_log_group_arn(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "cloud_watch_logs_log_group_arn", value)

    @property
    @pulumi.getter(name="cloudWatchLogsRoleArn")
    def cloud_watch_logs_role_arn(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's log group.
        """
        return pulumi.get(self, "cloud_watch_logs_role_arn")

    @cloud_watch_logs_role_arn.setter
    def cloud_watch_logs_role_arn(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "cloud_watch_logs_role_arn", value)

    @property
    @pulumi.getter(name="enableLogFileValidation")
    def enable_log_file_validation(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Specifies whether log file validation is enabled. The default is false.
        """
        return pulumi.get(self, "enable_log_file_validation")

    @enable_log_file_validation.setter
    def enable_log_file_validation(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enable_log_file_validation", value)

    @property
    @pulumi.getter(name="eventSelectors")
    def event_selectors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TrailEventSelectorArgs']]]]:
        """
        Use event selectors to further specify the management and data event settings for your trail. By default, trails created without specific event selectors will be configured to log all read and write management events, and no data events. When an event occurs in your account, CloudTrail evaluates the event selector for all trails. For each trail, if the event matches any event selector, the trail processes and logs the event. If the event doesn't match any event selector, the trail doesn't log the event. You can configure up to five event selectors for a trail.
        """
        return pulumi.get(self, "event_selectors")

    @event_selectors.setter
    def event_selectors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TrailEventSelectorArgs']]]]):
        pulumi.set(self, "event_selectors", value)

    @property
    @pulumi.getter(name="includeGlobalServiceEvents")
    def include_global_service_events(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Specifies whether the trail is publishing events from global services such as IAM to the log files.
        """
        return pulumi.get(self, "include_global_service_events")

    @include_global_service_events.setter
    def include_global_service_events(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "include_global_service_events", value)

    @property
    @pulumi.getter(name="insightSelectors")
    def insight_selectors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TrailInsightSelectorArgs']]]]:
        """
        Lets you enable Insights event logging by specifying the Insights selectors that you want to enable on an existing trail.
        """
        return pulumi.get(self, "insight_selectors")

    @insight_selectors.setter
    def insight_selectors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TrailInsightSelectorArgs']]]]):
        pulumi.set(self, "insight_selectors", value)

    @property
    @pulumi.getter(name="isMultiRegionTrail")
    def is_multi_region_trail(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Specifies whether the trail applies only to the current region or to all regions. The default is false. If the trail exists only in the current region and this value is set to true, shadow trails (replications of the trail) will be created in the other regions. If the trail exists in all regions and this value is set to false, the trail will remain in the region where it was created, and its shadow trails in other regions will be deleted. As a best practice, consider using trails that log events in all regions.
        """
        return pulumi.get(self, "is_multi_region_trail")

    @is_multi_region_trail.setter
    def is_multi_region_trail(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "is_multi_region_trail", value)

    @property
    @pulumi.getter(name="isOrganizationTrail")
    def is_organization_trail(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Specifies whether the trail is created for all accounts in an organization in AWS Organizations, or only for the current AWS account. The default is false, and cannot be true unless the call is made on behalf of an AWS account that is the master account for an organization in AWS Organizations.
        """
        return pulumi.get(self, "is_organization_trail")

    @is_organization_trail.setter
    def is_organization_trail(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "is_organization_trail", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the KMS key ID to use to encrypt the logs delivered by CloudTrail. The value can be an alias name prefixed by 'alias/', a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter(name="s3KeyPrefix")
    def s3_key_prefix(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery. For more information, see Finding Your CloudTrail Log Files. The maximum length is 200 characters.
        """
        return pulumi.get(self, "s3_key_prefix")

    @s3_key_prefix.setter
    def s3_key_prefix(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "s3_key_prefix", value)

    @property
    @pulumi.getter(name="snsTopicName")
    def sns_topic_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the name of the Amazon SNS topic defined for notification of log file delivery. The maximum length is 256 characters.
        """
        return pulumi.get(self, "sns_topic_name")

    @sns_topic_name.setter
    def sns_topic_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "sns_topic_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        A custom set of tags (key-value pairs) for this trail.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="trailName")
    def trail_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the name of the trail. The name must meet the following requirements:

        - Contain only ASCII letters (a-z, A-Z), numbers (0-9), periods (.), underscores (_), or dashes (-)
        - Start with a letter or number, and end with a letter or number
        - Be between 3 and 128 characters
        - Have no adjacent periods, underscores or dashes. Names like `my-_namespace` and `my--namespace` are not valid.
        - Not be in IP address format (for example, 192.168.5.4)
        """
        return pulumi.get(self, "trail_name")

    @trail_name.setter
    def trail_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "trail_name", value)


@pulumi.type_token("aws-native:cloudtrail:Trail")
class Trail(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 advanced_event_selectors: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TrailAdvancedEventSelectorArgs', 'TrailAdvancedEventSelectorArgsDict']]]]] = None,
                 cloud_watch_logs_log_group_arn: Optional[pulumi.Input[builtins.str]] = None,
                 cloud_watch_logs_role_arn: Optional[pulumi.Input[builtins.str]] = None,
                 enable_log_file_validation: Optional[pulumi.Input[builtins.bool]] = None,
                 event_selectors: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TrailEventSelectorArgs', 'TrailEventSelectorArgsDict']]]]] = None,
                 include_global_service_events: Optional[pulumi.Input[builtins.bool]] = None,
                 insight_selectors: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TrailInsightSelectorArgs', 'TrailInsightSelectorArgsDict']]]]] = None,
                 is_logging: Optional[pulumi.Input[builtins.bool]] = None,
                 is_multi_region_trail: Optional[pulumi.Input[builtins.bool]] = None,
                 is_organization_trail: Optional[pulumi.Input[builtins.bool]] = None,
                 kms_key_id: Optional[pulumi.Input[builtins.str]] = None,
                 s3_bucket_name: Optional[pulumi.Input[builtins.str]] = None,
                 s3_key_prefix: Optional[pulumi.Input[builtins.str]] = None,
                 sns_topic_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 trail_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Creates a trail that specifies the settings for delivery of log data to an Amazon S3 bucket. A maximum of five trails can exist in a region, irrespective of the region in which they were created.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['TrailAdvancedEventSelectorArgs', 'TrailAdvancedEventSelectorArgsDict']]]] advanced_event_selectors: The advanced event selectors that were used to select events for the data store.
        :param pulumi.Input[builtins.str] cloud_watch_logs_log_group_arn: Specifies a log group name using an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs will be delivered. Not required unless you specify CloudWatchLogsRoleArn.
        :param pulumi.Input[builtins.str] cloud_watch_logs_role_arn: Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's log group.
        :param pulumi.Input[builtins.bool] enable_log_file_validation: Specifies whether log file validation is enabled. The default is false.
        :param pulumi.Input[Sequence[pulumi.Input[Union['TrailEventSelectorArgs', 'TrailEventSelectorArgsDict']]]] event_selectors: Use event selectors to further specify the management and data event settings for your trail. By default, trails created without specific event selectors will be configured to log all read and write management events, and no data events. When an event occurs in your account, CloudTrail evaluates the event selector for all trails. For each trail, if the event matches any event selector, the trail processes and logs the event. If the event doesn't match any event selector, the trail doesn't log the event. You can configure up to five event selectors for a trail.
        :param pulumi.Input[builtins.bool] include_global_service_events: Specifies whether the trail is publishing events from global services such as IAM to the log files.
        :param pulumi.Input[Sequence[pulumi.Input[Union['TrailInsightSelectorArgs', 'TrailInsightSelectorArgsDict']]]] insight_selectors: Lets you enable Insights event logging by specifying the Insights selectors that you want to enable on an existing trail.
        :param pulumi.Input[builtins.bool] is_logging: Whether the CloudTrail is currently logging AWS API calls.
        :param pulumi.Input[builtins.bool] is_multi_region_trail: Specifies whether the trail applies only to the current region or to all regions. The default is false. If the trail exists only in the current region and this value is set to true, shadow trails (replications of the trail) will be created in the other regions. If the trail exists in all regions and this value is set to false, the trail will remain in the region where it was created, and its shadow trails in other regions will be deleted. As a best practice, consider using trails that log events in all regions.
        :param pulumi.Input[builtins.bool] is_organization_trail: Specifies whether the trail is created for all accounts in an organization in AWS Organizations, or only for the current AWS account. The default is false, and cannot be true unless the call is made on behalf of an AWS account that is the master account for an organization in AWS Organizations.
        :param pulumi.Input[builtins.str] kms_key_id: Specifies the KMS key ID to use to encrypt the logs delivered by CloudTrail. The value can be an alias name prefixed by 'alias/', a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
        :param pulumi.Input[builtins.str] s3_bucket_name: Specifies the name of the Amazon S3 bucket designated for publishing log files. See Amazon S3 Bucket Naming Requirements.
        :param pulumi.Input[builtins.str] s3_key_prefix: Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery. For more information, see Finding Your CloudTrail Log Files. The maximum length is 200 characters.
        :param pulumi.Input[builtins.str] sns_topic_name: Specifies the name of the Amazon SNS topic defined for notification of log file delivery. The maximum length is 256 characters.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] tags: A custom set of tags (key-value pairs) for this trail.
        :param pulumi.Input[builtins.str] trail_name: Specifies the name of the trail. The name must meet the following requirements:
               
               - Contain only ASCII letters (a-z, A-Z), numbers (0-9), periods (.), underscores (_), or dashes (-)
               - Start with a letter or number, and end with a letter or number
               - Be between 3 and 128 characters
               - Have no adjacent periods, underscores or dashes. Names like `my-_namespace` and `my--namespace` are not valid.
               - Not be in IP address format (for example, 192.168.5.4)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TrailArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a trail that specifies the settings for delivery of log data to an Amazon S3 bucket. A maximum of five trails can exist in a region, irrespective of the region in which they were created.

        :param str resource_name: The name of the resource.
        :param TrailArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TrailArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 advanced_event_selectors: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TrailAdvancedEventSelectorArgs', 'TrailAdvancedEventSelectorArgsDict']]]]] = None,
                 cloud_watch_logs_log_group_arn: Optional[pulumi.Input[builtins.str]] = None,
                 cloud_watch_logs_role_arn: Optional[pulumi.Input[builtins.str]] = None,
                 enable_log_file_validation: Optional[pulumi.Input[builtins.bool]] = None,
                 event_selectors: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TrailEventSelectorArgs', 'TrailEventSelectorArgsDict']]]]] = None,
                 include_global_service_events: Optional[pulumi.Input[builtins.bool]] = None,
                 insight_selectors: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TrailInsightSelectorArgs', 'TrailInsightSelectorArgsDict']]]]] = None,
                 is_logging: Optional[pulumi.Input[builtins.bool]] = None,
                 is_multi_region_trail: Optional[pulumi.Input[builtins.bool]] = None,
                 is_organization_trail: Optional[pulumi.Input[builtins.bool]] = None,
                 kms_key_id: Optional[pulumi.Input[builtins.str]] = None,
                 s3_bucket_name: Optional[pulumi.Input[builtins.str]] = None,
                 s3_key_prefix: Optional[pulumi.Input[builtins.str]] = None,
                 sns_topic_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 trail_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TrailArgs.__new__(TrailArgs)

            __props__.__dict__["advanced_event_selectors"] = advanced_event_selectors
            __props__.__dict__["cloud_watch_logs_log_group_arn"] = cloud_watch_logs_log_group_arn
            __props__.__dict__["cloud_watch_logs_role_arn"] = cloud_watch_logs_role_arn
            __props__.__dict__["enable_log_file_validation"] = enable_log_file_validation
            __props__.__dict__["event_selectors"] = event_selectors
            __props__.__dict__["include_global_service_events"] = include_global_service_events
            __props__.__dict__["insight_selectors"] = insight_selectors
            if is_logging is None and not opts.urn:
                raise TypeError("Missing required property 'is_logging'")
            __props__.__dict__["is_logging"] = is_logging
            __props__.__dict__["is_multi_region_trail"] = is_multi_region_trail
            __props__.__dict__["is_organization_trail"] = is_organization_trail
            __props__.__dict__["kms_key_id"] = kms_key_id
            if s3_bucket_name is None and not opts.urn:
                raise TypeError("Missing required property 's3_bucket_name'")
            __props__.__dict__["s3_bucket_name"] = s3_bucket_name
            __props__.__dict__["s3_key_prefix"] = s3_key_prefix
            __props__.__dict__["sns_topic_name"] = sns_topic_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["trail_name"] = trail_name
            __props__.__dict__["arn"] = None
            __props__.__dict__["sns_topic_arn"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["trailName"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Trail, __self__).__init__(
            'aws-native:cloudtrail:Trail',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Trail':
        """
        Get an existing Trail resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TrailArgs.__new__(TrailArgs)

        __props__.__dict__["advanced_event_selectors"] = None
        __props__.__dict__["arn"] = None
        __props__.__dict__["cloud_watch_logs_log_group_arn"] = None
        __props__.__dict__["cloud_watch_logs_role_arn"] = None
        __props__.__dict__["enable_log_file_validation"] = None
        __props__.__dict__["event_selectors"] = None
        __props__.__dict__["include_global_service_events"] = None
        __props__.__dict__["insight_selectors"] = None
        __props__.__dict__["is_logging"] = None
        __props__.__dict__["is_multi_region_trail"] = None
        __props__.__dict__["is_organization_trail"] = None
        __props__.__dict__["kms_key_id"] = None
        __props__.__dict__["s3_bucket_name"] = None
        __props__.__dict__["s3_key_prefix"] = None
        __props__.__dict__["sns_topic_arn"] = None
        __props__.__dict__["sns_topic_name"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["trail_name"] = None
        return Trail(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="advancedEventSelectors")
    def advanced_event_selectors(self) -> pulumi.Output[Optional[Sequence['outputs.TrailAdvancedEventSelector']]]:
        """
        The advanced event selectors that were used to select events for the data store.
        """
        return pulumi.get(self, "advanced_event_selectors")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[builtins.str]:
        """
        `Ref` returns the ARN of the CloudTrail trail, such as `arn:aws:cloudtrail:us-east-2:123456789012:trail/myCloudTrail` .
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="cloudWatchLogsLogGroupArn")
    def cloud_watch_logs_log_group_arn(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Specifies a log group name using an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs will be delivered. Not required unless you specify CloudWatchLogsRoleArn.
        """
        return pulumi.get(self, "cloud_watch_logs_log_group_arn")

    @property
    @pulumi.getter(name="cloudWatchLogsRoleArn")
    def cloud_watch_logs_role_arn(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's log group.
        """
        return pulumi.get(self, "cloud_watch_logs_role_arn")

    @property
    @pulumi.getter(name="enableLogFileValidation")
    def enable_log_file_validation(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Specifies whether log file validation is enabled. The default is false.
        """
        return pulumi.get(self, "enable_log_file_validation")

    @property
    @pulumi.getter(name="eventSelectors")
    def event_selectors(self) -> pulumi.Output[Optional[Sequence['outputs.TrailEventSelector']]]:
        """
        Use event selectors to further specify the management and data event settings for your trail. By default, trails created without specific event selectors will be configured to log all read and write management events, and no data events. When an event occurs in your account, CloudTrail evaluates the event selector for all trails. For each trail, if the event matches any event selector, the trail processes and logs the event. If the event doesn't match any event selector, the trail doesn't log the event. You can configure up to five event selectors for a trail.
        """
        return pulumi.get(self, "event_selectors")

    @property
    @pulumi.getter(name="includeGlobalServiceEvents")
    def include_global_service_events(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Specifies whether the trail is publishing events from global services such as IAM to the log files.
        """
        return pulumi.get(self, "include_global_service_events")

    @property
    @pulumi.getter(name="insightSelectors")
    def insight_selectors(self) -> pulumi.Output[Optional[Sequence['outputs.TrailInsightSelector']]]:
        """
        Lets you enable Insights event logging by specifying the Insights selectors that you want to enable on an existing trail.
        """
        return pulumi.get(self, "insight_selectors")

    @property
    @pulumi.getter(name="isLogging")
    def is_logging(self) -> pulumi.Output[builtins.bool]:
        """
        Whether the CloudTrail is currently logging AWS API calls.
        """
        return pulumi.get(self, "is_logging")

    @property
    @pulumi.getter(name="isMultiRegionTrail")
    def is_multi_region_trail(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Specifies whether the trail applies only to the current region or to all regions. The default is false. If the trail exists only in the current region and this value is set to true, shadow trails (replications of the trail) will be created in the other regions. If the trail exists in all regions and this value is set to false, the trail will remain in the region where it was created, and its shadow trails in other regions will be deleted. As a best practice, consider using trails that log events in all regions.
        """
        return pulumi.get(self, "is_multi_region_trail")

    @property
    @pulumi.getter(name="isOrganizationTrail")
    def is_organization_trail(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Specifies whether the trail is created for all accounts in an organization in AWS Organizations, or only for the current AWS account. The default is false, and cannot be true unless the call is made on behalf of an AWS account that is the master account for an organization in AWS Organizations.
        """
        return pulumi.get(self, "is_organization_trail")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Specifies the KMS key ID to use to encrypt the logs delivered by CloudTrail. The value can be an alias name prefixed by 'alias/', a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
        """
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter(name="s3BucketName")
    def s3_bucket_name(self) -> pulumi.Output[builtins.str]:
        """
        Specifies the name of the Amazon S3 bucket designated for publishing log files. See Amazon S3 Bucket Naming Requirements.
        """
        return pulumi.get(self, "s3_bucket_name")

    @property
    @pulumi.getter(name="s3KeyPrefix")
    def s3_key_prefix(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery. For more information, see Finding Your CloudTrail Log Files. The maximum length is 200 characters.
        """
        return pulumi.get(self, "s3_key_prefix")

    @property
    @pulumi.getter(name="snsTopicArn")
    def sns_topic_arn(self) -> pulumi.Output[builtins.str]:
        """
        `Ref` returns the ARN of the Amazon SNS topic that's associated with the CloudTrail trail, such as `arn:aws:sns:us-east-2:123456789012:mySNSTopic` .
        """
        return pulumi.get(self, "sns_topic_arn")

    @property
    @pulumi.getter(name="snsTopicName")
    def sns_topic_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Specifies the name of the Amazon SNS topic defined for notification of log file delivery. The maximum length is 256 characters.
        """
        return pulumi.get(self, "sns_topic_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        A custom set of tags (key-value pairs) for this trail.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="trailName")
    def trail_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Specifies the name of the trail. The name must meet the following requirements:

        - Contain only ASCII letters (a-z, A-Z), numbers (0-9), periods (.), underscores (_), or dashes (-)
        - Start with a letter or number, and end with a letter or number
        - Be between 3 and 128 characters
        - Have no adjacent periods, underscores or dashes. Names like `my-_namespace` and `my--namespace` are not valid.
        - Not be in IP address format (for example, 192.168.5.4)
        """
        return pulumi.get(self, "trail_name")

