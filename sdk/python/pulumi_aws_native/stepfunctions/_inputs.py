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
from ._enums import *

__all__ = [
    'ActivityEncryptionConfigurationArgs',
    'ActivityEncryptionConfigurationArgsDict',
    'StateMachineAliasDeploymentPreferenceArgs',
    'StateMachineAliasDeploymentPreferenceArgsDict',
    'StateMachineAliasRoutingConfigurationVersionArgs',
    'StateMachineAliasRoutingConfigurationVersionArgsDict',
    'StateMachineCloudWatchLogsLogGroupArgs',
    'StateMachineCloudWatchLogsLogGroupArgsDict',
    'StateMachineDefinitionArgs',
    'StateMachineDefinitionArgsDict',
    'StateMachineEncryptionConfigurationArgs',
    'StateMachineEncryptionConfigurationArgsDict',
    'StateMachineLogDestinationArgs',
    'StateMachineLogDestinationArgsDict',
    'StateMachineLoggingConfigurationArgs',
    'StateMachineLoggingConfigurationArgsDict',
    'StateMachineS3LocationArgs',
    'StateMachineS3LocationArgsDict',
    'StateMachineTracingConfigurationArgs',
    'StateMachineTracingConfigurationArgsDict',
]

MYPY = False

if not MYPY:
    class ActivityEncryptionConfigurationArgsDict(TypedDict):
        type: pulumi.Input['ActivityEncryptionConfigurationType']
        """
        Encryption option for an activity.
        """
        kms_data_key_reuse_period_seconds: NotRequired[pulumi.Input[builtins.int]]
        """
        Maximum duration that Step Functions will reuse data keys. When the period expires, Step Functions will call `GenerateDataKey` . Only applies to customer managed keys.
        """
        kms_key_id: NotRequired[pulumi.Input[builtins.str]]
        """
        An alias, alias ARN, key ID, or key ARN of a symmetric encryption AWS KMS key to encrypt data. To specify a AWS KMS key in a different AWS account, you must use the key ARN or alias ARN.
        """
elif False:
    ActivityEncryptionConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ActivityEncryptionConfigurationArgs:
    def __init__(__self__, *,
                 type: pulumi.Input['ActivityEncryptionConfigurationType'],
                 kms_data_key_reuse_period_seconds: Optional[pulumi.Input[builtins.int]] = None,
                 kms_key_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        :param pulumi.Input['ActivityEncryptionConfigurationType'] type: Encryption option for an activity.
        :param pulumi.Input[builtins.int] kms_data_key_reuse_period_seconds: Maximum duration that Step Functions will reuse data keys. When the period expires, Step Functions will call `GenerateDataKey` . Only applies to customer managed keys.
        :param pulumi.Input[builtins.str] kms_key_id: An alias, alias ARN, key ID, or key ARN of a symmetric encryption AWS KMS key to encrypt data. To specify a AWS KMS key in a different AWS account, you must use the key ARN or alias ARN.
        """
        pulumi.set(__self__, "type", type)
        if kms_data_key_reuse_period_seconds is not None:
            pulumi.set(__self__, "kms_data_key_reuse_period_seconds", kms_data_key_reuse_period_seconds)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['ActivityEncryptionConfigurationType']:
        """
        Encryption option for an activity.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['ActivityEncryptionConfigurationType']):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="kmsDataKeyReusePeriodSeconds")
    def kms_data_key_reuse_period_seconds(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Maximum duration that Step Functions will reuse data keys. When the period expires, Step Functions will call `GenerateDataKey` . Only applies to customer managed keys.
        """
        return pulumi.get(self, "kms_data_key_reuse_period_seconds")

    @kms_data_key_reuse_period_seconds.setter
    def kms_data_key_reuse_period_seconds(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "kms_data_key_reuse_period_seconds", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        An alias, alias ARN, key ID, or key ARN of a symmetric encryption AWS KMS key to encrypt data. To specify a AWS KMS key in a different AWS account, you must use the key ARN or alias ARN.
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "kms_key_id", value)


if not MYPY:
    class StateMachineAliasDeploymentPreferenceArgsDict(TypedDict):
        """
        The settings to enable gradual state machine deployments.
        """
        state_machine_version_arn: pulumi.Input[builtins.str]
        """
        The Amazon Resource Name (ARN) of the [`AWS::StepFunctions::StateMachineVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachineversion.html) resource that will be the final version to which the alias points to when the traffic shifting is complete.

        While performing gradual deployments, you can only provide a single state machine version ARN. To explicitly set version weights in a CloudFormation template, use `RoutingConfiguration` instead.
        """
        type: pulumi.Input['StateMachineAliasDeploymentPreferenceType']
        """
        The type of deployment to perform.
        """
        alarms: NotRequired[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]
        """
        A list of CloudWatch alarm names that will be monitored during the deployment. The deployment will fail and rollback if any alarms go into ALARM state.
        """
        interval: NotRequired[pulumi.Input[builtins.int]]
        """
        The time in minutes between each traffic shifting increment.
        """
        percentage: NotRequired[pulumi.Input[builtins.int]]
        """
        The percentage of traffic to shift to the new version in each increment.
        """
elif False:
    StateMachineAliasDeploymentPreferenceArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class StateMachineAliasDeploymentPreferenceArgs:
    def __init__(__self__, *,
                 state_machine_version_arn: pulumi.Input[builtins.str],
                 type: pulumi.Input['StateMachineAliasDeploymentPreferenceType'],
                 alarms: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 interval: Optional[pulumi.Input[builtins.int]] = None,
                 percentage: Optional[pulumi.Input[builtins.int]] = None):
        """
        The settings to enable gradual state machine deployments.
        :param pulumi.Input[builtins.str] state_machine_version_arn: The Amazon Resource Name (ARN) of the [`AWS::StepFunctions::StateMachineVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachineversion.html) resource that will be the final version to which the alias points to when the traffic shifting is complete.
               
               While performing gradual deployments, you can only provide a single state machine version ARN. To explicitly set version weights in a CloudFormation template, use `RoutingConfiguration` instead.
        :param pulumi.Input['StateMachineAliasDeploymentPreferenceType'] type: The type of deployment to perform.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] alarms: A list of CloudWatch alarm names that will be monitored during the deployment. The deployment will fail and rollback if any alarms go into ALARM state.
        :param pulumi.Input[builtins.int] interval: The time in minutes between each traffic shifting increment.
        :param pulumi.Input[builtins.int] percentage: The percentage of traffic to shift to the new version in each increment.
        """
        pulumi.set(__self__, "state_machine_version_arn", state_machine_version_arn)
        pulumi.set(__self__, "type", type)
        if alarms is not None:
            pulumi.set(__self__, "alarms", alarms)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if percentage is not None:
            pulumi.set(__self__, "percentage", percentage)

    @property
    @pulumi.getter(name="stateMachineVersionArn")
    def state_machine_version_arn(self) -> pulumi.Input[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the [`AWS::StepFunctions::StateMachineVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachineversion.html) resource that will be the final version to which the alias points to when the traffic shifting is complete.

        While performing gradual deployments, you can only provide a single state machine version ARN. To explicitly set version weights in a CloudFormation template, use `RoutingConfiguration` instead.
        """
        return pulumi.get(self, "state_machine_version_arn")

    @state_machine_version_arn.setter
    def state_machine_version_arn(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "state_machine_version_arn", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['StateMachineAliasDeploymentPreferenceType']:
        """
        The type of deployment to perform.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['StateMachineAliasDeploymentPreferenceType']):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def alarms(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        A list of CloudWatch alarm names that will be monitored during the deployment. The deployment will fail and rollback if any alarms go into ALARM state.
        """
        return pulumi.get(self, "alarms")

    @alarms.setter
    def alarms(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "alarms", value)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The time in minutes between each traffic shifting increment.
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter
    def percentage(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The percentage of traffic to shift to the new version in each increment.
        """
        return pulumi.get(self, "percentage")

    @percentage.setter
    def percentage(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "percentage", value)


if not MYPY:
    class StateMachineAliasRoutingConfigurationVersionArgsDict(TypedDict):
        state_machine_version_arn: pulumi.Input[builtins.str]
        """
        The Amazon Resource Name (ARN) that identifies one or two state machine versions defined in the routing configuration.
        """
        weight: pulumi.Input[builtins.int]
        """
        The percentage of traffic you want to route to the state machine version. The sum of the weights in the routing configuration must be equal to 100.
        """
elif False:
    StateMachineAliasRoutingConfigurationVersionArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class StateMachineAliasRoutingConfigurationVersionArgs:
    def __init__(__self__, *,
                 state_machine_version_arn: pulumi.Input[builtins.str],
                 weight: pulumi.Input[builtins.int]):
        """
        :param pulumi.Input[builtins.str] state_machine_version_arn: The Amazon Resource Name (ARN) that identifies one or two state machine versions defined in the routing configuration.
        :param pulumi.Input[builtins.int] weight: The percentage of traffic you want to route to the state machine version. The sum of the weights in the routing configuration must be equal to 100.
        """
        pulumi.set(__self__, "state_machine_version_arn", state_machine_version_arn)
        pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter(name="stateMachineVersionArn")
    def state_machine_version_arn(self) -> pulumi.Input[builtins.str]:
        """
        The Amazon Resource Name (ARN) that identifies one or two state machine versions defined in the routing configuration.
        """
        return pulumi.get(self, "state_machine_version_arn")

    @state_machine_version_arn.setter
    def state_machine_version_arn(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "state_machine_version_arn", value)

    @property
    @pulumi.getter
    def weight(self) -> pulumi.Input[builtins.int]:
        """
        The percentage of traffic you want to route to the state machine version. The sum of the weights in the routing configuration must be equal to 100.
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "weight", value)


if not MYPY:
    class StateMachineCloudWatchLogsLogGroupArgsDict(TypedDict):
        log_group_arn: NotRequired[pulumi.Input[builtins.str]]
        """
        The ARN of the the CloudWatch log group to which you want your logs emitted to. The ARN must end with `:*`
        """
elif False:
    StateMachineCloudWatchLogsLogGroupArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class StateMachineCloudWatchLogsLogGroupArgs:
    def __init__(__self__, *,
                 log_group_arn: Optional[pulumi.Input[builtins.str]] = None):
        """
        :param pulumi.Input[builtins.str] log_group_arn: The ARN of the the CloudWatch log group to which you want your logs emitted to. The ARN must end with `:*`
        """
        if log_group_arn is not None:
            pulumi.set(__self__, "log_group_arn", log_group_arn)

    @property
    @pulumi.getter(name="logGroupArn")
    def log_group_arn(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ARN of the the CloudWatch log group to which you want your logs emitted to. The ARN must end with `:*`
        """
        return pulumi.get(self, "log_group_arn")

    @log_group_arn.setter
    def log_group_arn(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "log_group_arn", value)


if not MYPY:
    class StateMachineDefinitionArgsDict(TypedDict):
        pass
elif False:
    StateMachineDefinitionArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class StateMachineDefinitionArgs:
    def __init__(__self__):
        pass


if not MYPY:
    class StateMachineEncryptionConfigurationArgsDict(TypedDict):
        type: pulumi.Input['StateMachineEncryptionConfigurationType']
        """
        Encryption option for a state machine.
        """
        kms_data_key_reuse_period_seconds: NotRequired[pulumi.Input[builtins.int]]
        """
        Maximum duration that Step Functions will reuse data keys. When the period expires, Step Functions will call `GenerateDataKey` . Only applies to customer managed keys.
        """
        kms_key_id: NotRequired[pulumi.Input[builtins.str]]
        """
        An alias, alias ARN, key ID, or key ARN of a symmetric encryption AWS KMS key to encrypt data. To specify a AWS KMS key in a different AWS account, you must use the key ARN or alias ARN.
        """
elif False:
    StateMachineEncryptionConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class StateMachineEncryptionConfigurationArgs:
    def __init__(__self__, *,
                 type: pulumi.Input['StateMachineEncryptionConfigurationType'],
                 kms_data_key_reuse_period_seconds: Optional[pulumi.Input[builtins.int]] = None,
                 kms_key_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        :param pulumi.Input['StateMachineEncryptionConfigurationType'] type: Encryption option for a state machine.
        :param pulumi.Input[builtins.int] kms_data_key_reuse_period_seconds: Maximum duration that Step Functions will reuse data keys. When the period expires, Step Functions will call `GenerateDataKey` . Only applies to customer managed keys.
        :param pulumi.Input[builtins.str] kms_key_id: An alias, alias ARN, key ID, or key ARN of a symmetric encryption AWS KMS key to encrypt data. To specify a AWS KMS key in a different AWS account, you must use the key ARN or alias ARN.
        """
        pulumi.set(__self__, "type", type)
        if kms_data_key_reuse_period_seconds is not None:
            pulumi.set(__self__, "kms_data_key_reuse_period_seconds", kms_data_key_reuse_period_seconds)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['StateMachineEncryptionConfigurationType']:
        """
        Encryption option for a state machine.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['StateMachineEncryptionConfigurationType']):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="kmsDataKeyReusePeriodSeconds")
    def kms_data_key_reuse_period_seconds(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Maximum duration that Step Functions will reuse data keys. When the period expires, Step Functions will call `GenerateDataKey` . Only applies to customer managed keys.
        """
        return pulumi.get(self, "kms_data_key_reuse_period_seconds")

    @kms_data_key_reuse_period_seconds.setter
    def kms_data_key_reuse_period_seconds(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "kms_data_key_reuse_period_seconds", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        An alias, alias ARN, key ID, or key ARN of a symmetric encryption AWS KMS key to encrypt data. To specify a AWS KMS key in a different AWS account, you must use the key ARN or alias ARN.
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "kms_key_id", value)


if not MYPY:
    class StateMachineLogDestinationArgsDict(TypedDict):
        cloud_watch_logs_log_group: NotRequired[pulumi.Input['StateMachineCloudWatchLogsLogGroupArgsDict']]
        """
        An object describing a CloudWatch log group. For more information, see [AWS::Logs::LogGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html) in the AWS CloudFormation User Guide.
        """
elif False:
    StateMachineLogDestinationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class StateMachineLogDestinationArgs:
    def __init__(__self__, *,
                 cloud_watch_logs_log_group: Optional[pulumi.Input['StateMachineCloudWatchLogsLogGroupArgs']] = None):
        """
        :param pulumi.Input['StateMachineCloudWatchLogsLogGroupArgs'] cloud_watch_logs_log_group: An object describing a CloudWatch log group. For more information, see [AWS::Logs::LogGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html) in the AWS CloudFormation User Guide.
        """
        if cloud_watch_logs_log_group is not None:
            pulumi.set(__self__, "cloud_watch_logs_log_group", cloud_watch_logs_log_group)

    @property
    @pulumi.getter(name="cloudWatchLogsLogGroup")
    def cloud_watch_logs_log_group(self) -> Optional[pulumi.Input['StateMachineCloudWatchLogsLogGroupArgs']]:
        """
        An object describing a CloudWatch log group. For more information, see [AWS::Logs::LogGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html) in the AWS CloudFormation User Guide.
        """
        return pulumi.get(self, "cloud_watch_logs_log_group")

    @cloud_watch_logs_log_group.setter
    def cloud_watch_logs_log_group(self, value: Optional[pulumi.Input['StateMachineCloudWatchLogsLogGroupArgs']]):
        pulumi.set(self, "cloud_watch_logs_log_group", value)


if not MYPY:
    class StateMachineLoggingConfigurationArgsDict(TypedDict):
        destinations: NotRequired[pulumi.Input[Sequence[pulumi.Input['StateMachineLogDestinationArgsDict']]]]
        """
        An array of objects that describes where your execution history events will be logged. Limited to size 1. Required, if your log level is not set to `OFF` .
        """
        include_execution_data: NotRequired[pulumi.Input[builtins.bool]]
        """
        Determines whether execution data is included in your log. When set to `false` , data is excluded.
        """
        level: NotRequired[pulumi.Input['StateMachineLoggingConfigurationLevel']]
        """
        Defines which category of execution history events are logged.
        """
elif False:
    StateMachineLoggingConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class StateMachineLoggingConfigurationArgs:
    def __init__(__self__, *,
                 destinations: Optional[pulumi.Input[Sequence[pulumi.Input['StateMachineLogDestinationArgs']]]] = None,
                 include_execution_data: Optional[pulumi.Input[builtins.bool]] = None,
                 level: Optional[pulumi.Input['StateMachineLoggingConfigurationLevel']] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input['StateMachineLogDestinationArgs']]] destinations: An array of objects that describes where your execution history events will be logged. Limited to size 1. Required, if your log level is not set to `OFF` .
        :param pulumi.Input[builtins.bool] include_execution_data: Determines whether execution data is included in your log. When set to `false` , data is excluded.
        :param pulumi.Input['StateMachineLoggingConfigurationLevel'] level: Defines which category of execution history events are logged.
        """
        if destinations is not None:
            pulumi.set(__self__, "destinations", destinations)
        if include_execution_data is not None:
            pulumi.set(__self__, "include_execution_data", include_execution_data)
        if level is not None:
            pulumi.set(__self__, "level", level)

    @property
    @pulumi.getter
    def destinations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StateMachineLogDestinationArgs']]]]:
        """
        An array of objects that describes where your execution history events will be logged. Limited to size 1. Required, if your log level is not set to `OFF` .
        """
        return pulumi.get(self, "destinations")

    @destinations.setter
    def destinations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StateMachineLogDestinationArgs']]]]):
        pulumi.set(self, "destinations", value)

    @property
    @pulumi.getter(name="includeExecutionData")
    def include_execution_data(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Determines whether execution data is included in your log. When set to `false` , data is excluded.
        """
        return pulumi.get(self, "include_execution_data")

    @include_execution_data.setter
    def include_execution_data(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "include_execution_data", value)

    @property
    @pulumi.getter
    def level(self) -> Optional[pulumi.Input['StateMachineLoggingConfigurationLevel']]:
        """
        Defines which category of execution history events are logged.
        """
        return pulumi.get(self, "level")

    @level.setter
    def level(self, value: Optional[pulumi.Input['StateMachineLoggingConfigurationLevel']]):
        pulumi.set(self, "level", value)


if not MYPY:
    class StateMachineS3LocationArgsDict(TypedDict):
        bucket: pulumi.Input[builtins.str]
        """
        The name of the S3 bucket where the state machine definition JSON or YAML file is stored.
        """
        key: pulumi.Input[builtins.str]
        """
        The name of the state machine definition file (Amazon S3 object name).
        """
        version: NotRequired[pulumi.Input[builtins.str]]
        """
        For versioning-enabled buckets, a specific version of the state machine definition.
        """
elif False:
    StateMachineS3LocationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class StateMachineS3LocationArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[builtins.str],
                 key: pulumi.Input[builtins.str],
                 version: Optional[pulumi.Input[builtins.str]] = None):
        """
        :param pulumi.Input[builtins.str] bucket: The name of the S3 bucket where the state machine definition JSON or YAML file is stored.
        :param pulumi.Input[builtins.str] key: The name of the state machine definition file (Amazon S3 object name).
        :param pulumi.Input[builtins.str] version: For versioning-enabled buckets, a specific version of the state machine definition.
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "key", key)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[builtins.str]:
        """
        The name of the S3 bucket where the state machine definition JSON or YAML file is stored.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[builtins.str]:
        """
        The name of the state machine definition file (Amazon S3 object name).
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        For versioning-enabled buckets, a specific version of the state machine definition.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "version", value)


if not MYPY:
    class StateMachineTracingConfigurationArgsDict(TypedDict):
        enabled: NotRequired[pulumi.Input[builtins.bool]]
        """
        When set to `true` , X-Ray tracing is enabled.
        """
elif False:
    StateMachineTracingConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class StateMachineTracingConfigurationArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None):
        """
        :param pulumi.Input[builtins.bool] enabled: When set to `true` , X-Ray tracing is enabled.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        When set to `true` , X-Ray tracing is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)


