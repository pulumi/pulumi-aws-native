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
    'DataLakeEncryptionConfigurationArgs',
    'DataLakeEncryptionConfigurationArgsDict',
    'DataLakeExpirationArgs',
    'DataLakeExpirationArgsDict',
    'DataLakeLifecycleConfigurationArgs',
    'DataLakeLifecycleConfigurationArgsDict',
    'DataLakeReplicationConfigurationArgs',
    'DataLakeReplicationConfigurationArgsDict',
    'DataLakeTransitionsArgs',
    'DataLakeTransitionsArgsDict',
    'SubscriberAwsLogSourceArgs',
    'SubscriberAwsLogSourceArgsDict',
    'SubscriberCustomLogSourceArgs',
    'SubscriberCustomLogSourceArgsDict',
    'SubscriberIdentityPropertiesArgs',
    'SubscriberIdentityPropertiesArgsDict',
    'SubscriberNotificationHttpsNotificationConfigurationArgs',
    'SubscriberNotificationHttpsNotificationConfigurationArgsDict',
    'SubscriberNotificationNotificationConfigurationArgs',
    'SubscriberNotificationNotificationConfigurationArgsDict',
    'SubscriberNotificationSqsNotificationConfigurationArgs',
    'SubscriberNotificationSqsNotificationConfigurationArgsDict',
    'SubscriberSourceArgs',
    'SubscriberSourceArgsDict',
]

MYPY = False

if not MYPY:
    class DataLakeEncryptionConfigurationArgsDict(TypedDict):
        """
        Provides encryption details of Amazon Security Lake object.
        """
        kms_key_id: NotRequired[pulumi.Input[builtins.str]]
        """
        The id of KMS encryption key used by Amazon Security Lake to encrypt the Security Lake object.
        """
elif False:
    DataLakeEncryptionConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class DataLakeEncryptionConfigurationArgs:
    def __init__(__self__, *,
                 kms_key_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Provides encryption details of Amazon Security Lake object.
        :param pulumi.Input[builtins.str] kms_key_id: The id of KMS encryption key used by Amazon Security Lake to encrypt the Security Lake object.
        """
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The id of KMS encryption key used by Amazon Security Lake to encrypt the Security Lake object.
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "kms_key_id", value)


if not MYPY:
    class DataLakeExpirationArgsDict(TypedDict):
        """
        Provides data expiration details of Amazon Security Lake object.
        """
        days: NotRequired[pulumi.Input[builtins.int]]
        """
        The number of days before data expires in the Amazon Security Lake object.
        """
elif False:
    DataLakeExpirationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class DataLakeExpirationArgs:
    def __init__(__self__, *,
                 days: Optional[pulumi.Input[builtins.int]] = None):
        """
        Provides data expiration details of Amazon Security Lake object.
        :param pulumi.Input[builtins.int] days: The number of days before data expires in the Amazon Security Lake object.
        """
        if days is not None:
            pulumi.set(__self__, "days", days)

    @property
    @pulumi.getter
    def days(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The number of days before data expires in the Amazon Security Lake object.
        """
        return pulumi.get(self, "days")

    @days.setter
    def days(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "days", value)


if not MYPY:
    class DataLakeLifecycleConfigurationArgsDict(TypedDict):
        """
        Provides lifecycle details of Amazon Security Lake object.
        """
        expiration: NotRequired[pulumi.Input['DataLakeExpirationArgsDict']]
        """
        Provides data expiration details of the Amazon Security Lake object.
        """
        transitions: NotRequired[pulumi.Input[Sequence[pulumi.Input['DataLakeTransitionsArgsDict']]]]
        """
        Provides data storage transition details of Amazon Security Lake object.
        """
elif False:
    DataLakeLifecycleConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class DataLakeLifecycleConfigurationArgs:
    def __init__(__self__, *,
                 expiration: Optional[pulumi.Input['DataLakeExpirationArgs']] = None,
                 transitions: Optional[pulumi.Input[Sequence[pulumi.Input['DataLakeTransitionsArgs']]]] = None):
        """
        Provides lifecycle details of Amazon Security Lake object.
        :param pulumi.Input['DataLakeExpirationArgs'] expiration: Provides data expiration details of the Amazon Security Lake object.
        :param pulumi.Input[Sequence[pulumi.Input['DataLakeTransitionsArgs']]] transitions: Provides data storage transition details of Amazon Security Lake object.
        """
        if expiration is not None:
            pulumi.set(__self__, "expiration", expiration)
        if transitions is not None:
            pulumi.set(__self__, "transitions", transitions)

    @property
    @pulumi.getter
    def expiration(self) -> Optional[pulumi.Input['DataLakeExpirationArgs']]:
        """
        Provides data expiration details of the Amazon Security Lake object.
        """
        return pulumi.get(self, "expiration")

    @expiration.setter
    def expiration(self, value: Optional[pulumi.Input['DataLakeExpirationArgs']]):
        pulumi.set(self, "expiration", value)

    @property
    @pulumi.getter
    def transitions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataLakeTransitionsArgs']]]]:
        """
        Provides data storage transition details of Amazon Security Lake object.
        """
        return pulumi.get(self, "transitions")

    @transitions.setter
    def transitions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataLakeTransitionsArgs']]]]):
        pulumi.set(self, "transitions", value)


if not MYPY:
    class DataLakeReplicationConfigurationArgsDict(TypedDict):
        """
        Provides replication details of Amazon Security Lake object.
        """
        regions: NotRequired[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]
        """
        Specifies one or more centralized rollup Regions. The AWS Region specified in the region parameter of the `CreateDataLake` or `UpdateDataLake` operations contributes data to the rollup Region or Regions specified in this parameter.

        Replication enables automatic, asynchronous copying of objects across Amazon S3 buckets. S3 buckets that are configured for object replication can be owned by the same AWS account or by different accounts. You can replicate objects to a single destination bucket or to multiple destination buckets. The destination buckets can be in different Regions or within the same Region as the source bucket.
        """
        role_arn: NotRequired[pulumi.Input[builtins.str]]
        """
        Replication settings for the Amazon S3 buckets. This parameter uses the AWS Identity and Access Management (IAM) role you created that is managed by Security Lake, to ensure the replication setting is correct.
        """
elif False:
    DataLakeReplicationConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class DataLakeReplicationConfigurationArgs:
    def __init__(__self__, *,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 role_arn: Optional[pulumi.Input[builtins.str]] = None):
        """
        Provides replication details of Amazon Security Lake object.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] regions: Specifies one or more centralized rollup Regions. The AWS Region specified in the region parameter of the `CreateDataLake` or `UpdateDataLake` operations contributes data to the rollup Region or Regions specified in this parameter.
               
               Replication enables automatic, asynchronous copying of objects across Amazon S3 buckets. S3 buckets that are configured for object replication can be owned by the same AWS account or by different accounts. You can replicate objects to a single destination bucket or to multiple destination buckets. The destination buckets can be in different Regions or within the same Region as the source bucket.
        :param pulumi.Input[builtins.str] role_arn: Replication settings for the Amazon S3 buckets. This parameter uses the AWS Identity and Access Management (IAM) role you created that is managed by Security Lake, to ensure the replication setting is correct.
        """
        if regions is not None:
            pulumi.set(__self__, "regions", regions)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)

    @property
    @pulumi.getter
    def regions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Specifies one or more centralized rollup Regions. The AWS Region specified in the region parameter of the `CreateDataLake` or `UpdateDataLake` operations contributes data to the rollup Region or Regions specified in this parameter.

        Replication enables automatic, asynchronous copying of objects across Amazon S3 buckets. S3 buckets that are configured for object replication can be owned by the same AWS account or by different accounts. You can replicate objects to a single destination bucket or to multiple destination buckets. The destination buckets can be in different Regions or within the same Region as the source bucket.
        """
        return pulumi.get(self, "regions")

    @regions.setter
    def regions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "regions", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Replication settings for the Amazon S3 buckets. This parameter uses the AWS Identity and Access Management (IAM) role you created that is managed by Security Lake, to ensure the replication setting is correct.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "role_arn", value)


if not MYPY:
    class DataLakeTransitionsArgsDict(TypedDict):
        days: NotRequired[pulumi.Input[builtins.int]]
        """
        Number of days before data transitions to a different S3 Storage Class in the Amazon Security Lake object.
        """
        storage_class: NotRequired[pulumi.Input[builtins.str]]
        """
        The range of storage classes that you can choose from based on the data access, resiliency, and cost requirements of your workloads.
        """
elif False:
    DataLakeTransitionsArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class DataLakeTransitionsArgs:
    def __init__(__self__, *,
                 days: Optional[pulumi.Input[builtins.int]] = None,
                 storage_class: Optional[pulumi.Input[builtins.str]] = None):
        """
        :param pulumi.Input[builtins.int] days: Number of days before data transitions to a different S3 Storage Class in the Amazon Security Lake object.
        :param pulumi.Input[builtins.str] storage_class: The range of storage classes that you can choose from based on the data access, resiliency, and cost requirements of your workloads.
        """
        if days is not None:
            pulumi.set(__self__, "days", days)
        if storage_class is not None:
            pulumi.set(__self__, "storage_class", storage_class)

    @property
    @pulumi.getter
    def days(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Number of days before data transitions to a different S3 Storage Class in the Amazon Security Lake object.
        """
        return pulumi.get(self, "days")

    @days.setter
    def days(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "days", value)

    @property
    @pulumi.getter(name="storageClass")
    def storage_class(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The range of storage classes that you can choose from based on the data access, resiliency, and cost requirements of your workloads.
        """
        return pulumi.get(self, "storage_class")

    @storage_class.setter
    def storage_class(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "storage_class", value)


if not MYPY:
    class SubscriberAwsLogSourceArgsDict(TypedDict):
        """
        Amazon Security Lake supports log and event collection for natively supported AWS services.
        """
        source_name: NotRequired[pulumi.Input[builtins.str]]
        """
        The name for a AWS source. This must be a Regionally unique value.
        """
        source_version: NotRequired[pulumi.Input[builtins.str]]
        """
        The version for a AWS source. This must be a Regionally unique value.
        """
elif False:
    SubscriberAwsLogSourceArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class SubscriberAwsLogSourceArgs:
    def __init__(__self__, *,
                 source_name: Optional[pulumi.Input[builtins.str]] = None,
                 source_version: Optional[pulumi.Input[builtins.str]] = None):
        """
        Amazon Security Lake supports log and event collection for natively supported AWS services.
        :param pulumi.Input[builtins.str] source_name: The name for a AWS source. This must be a Regionally unique value.
        :param pulumi.Input[builtins.str] source_version: The version for a AWS source. This must be a Regionally unique value.
        """
        if source_name is not None:
            pulumi.set(__self__, "source_name", source_name)
        if source_version is not None:
            pulumi.set(__self__, "source_version", source_version)

    @property
    @pulumi.getter(name="sourceName")
    def source_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name for a AWS source. This must be a Regionally unique value.
        """
        return pulumi.get(self, "source_name")

    @source_name.setter
    def source_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "source_name", value)

    @property
    @pulumi.getter(name="sourceVersion")
    def source_version(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The version for a AWS source. This must be a Regionally unique value.
        """
        return pulumi.get(self, "source_version")

    @source_version.setter
    def source_version(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "source_version", value)


if not MYPY:
    class SubscriberCustomLogSourceArgsDict(TypedDict):
        source_name: NotRequired[pulumi.Input[builtins.str]]
        """
        The name for a third-party custom source. This must be a Regionally unique value.
        """
        source_version: NotRequired[pulumi.Input[builtins.str]]
        """
        The version for a third-party custom source. This must be a Regionally unique value.
        """
elif False:
    SubscriberCustomLogSourceArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class SubscriberCustomLogSourceArgs:
    def __init__(__self__, *,
                 source_name: Optional[pulumi.Input[builtins.str]] = None,
                 source_version: Optional[pulumi.Input[builtins.str]] = None):
        """
        :param pulumi.Input[builtins.str] source_name: The name for a third-party custom source. This must be a Regionally unique value.
        :param pulumi.Input[builtins.str] source_version: The version for a third-party custom source. This must be a Regionally unique value.
        """
        if source_name is not None:
            pulumi.set(__self__, "source_name", source_name)
        if source_version is not None:
            pulumi.set(__self__, "source_version", source_version)

    @property
    @pulumi.getter(name="sourceName")
    def source_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name for a third-party custom source. This must be a Regionally unique value.
        """
        return pulumi.get(self, "source_name")

    @source_name.setter
    def source_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "source_name", value)

    @property
    @pulumi.getter(name="sourceVersion")
    def source_version(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The version for a third-party custom source. This must be a Regionally unique value.
        """
        return pulumi.get(self, "source_version")

    @source_version.setter
    def source_version(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "source_version", value)


if not MYPY:
    class SubscriberIdentityPropertiesArgsDict(TypedDict):
        """
        The AWS identity used to access your data.
        """
        external_id: pulumi.Input[builtins.str]
        """
        The external ID used to establish trust relationship with the AWS identity.
        """
        principal: pulumi.Input[builtins.str]
        """
        The AWS identity principal.
        """
elif False:
    SubscriberIdentityPropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class SubscriberIdentityPropertiesArgs:
    def __init__(__self__, *,
                 external_id: pulumi.Input[builtins.str],
                 principal: pulumi.Input[builtins.str]):
        """
        The AWS identity used to access your data.
        :param pulumi.Input[builtins.str] external_id: The external ID used to establish trust relationship with the AWS identity.
        :param pulumi.Input[builtins.str] principal: The AWS identity principal.
        """
        pulumi.set(__self__, "external_id", external_id)
        pulumi.set(__self__, "principal", principal)

    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> pulumi.Input[builtins.str]:
        """
        The external ID used to establish trust relationship with the AWS identity.
        """
        return pulumi.get(self, "external_id")

    @external_id.setter
    def external_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "external_id", value)

    @property
    @pulumi.getter
    def principal(self) -> pulumi.Input[builtins.str]:
        """
        The AWS identity principal.
        """
        return pulumi.get(self, "principal")

    @principal.setter
    def principal(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "principal", value)


if not MYPY:
    class SubscriberNotificationHttpsNotificationConfigurationArgsDict(TypedDict):
        """
        The configuration for HTTPS subscriber notification.
        """
        endpoint: pulumi.Input[builtins.str]
        """
        The subscription endpoint in Security Lake.
        """
        target_role_arn: pulumi.Input[builtins.str]
        """
        The Amazon Resource Name (ARN) of the EventBridge API destinations IAM role that you created.
        """
        authorization_api_key_name: NotRequired[pulumi.Input[builtins.str]]
        """
        The key name for the notification subscription.
        """
        authorization_api_key_value: NotRequired[pulumi.Input[builtins.str]]
        """
        The key value for the notification subscription.
        """
        http_method: NotRequired[pulumi.Input['SubscriberNotificationHttpsNotificationConfigurationHttpMethod']]
        """
        The HTTPS method used for the notification subscription.
        """
elif False:
    SubscriberNotificationHttpsNotificationConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class SubscriberNotificationHttpsNotificationConfigurationArgs:
    def __init__(__self__, *,
                 endpoint: pulumi.Input[builtins.str],
                 target_role_arn: pulumi.Input[builtins.str],
                 authorization_api_key_name: Optional[pulumi.Input[builtins.str]] = None,
                 authorization_api_key_value: Optional[pulumi.Input[builtins.str]] = None,
                 http_method: Optional[pulumi.Input['SubscriberNotificationHttpsNotificationConfigurationHttpMethod']] = None):
        """
        The configuration for HTTPS subscriber notification.
        :param pulumi.Input[builtins.str] endpoint: The subscription endpoint in Security Lake.
        :param pulumi.Input[builtins.str] target_role_arn: The Amazon Resource Name (ARN) of the EventBridge API destinations IAM role that you created.
        :param pulumi.Input[builtins.str] authorization_api_key_name: The key name for the notification subscription.
        :param pulumi.Input[builtins.str] authorization_api_key_value: The key value for the notification subscription.
        :param pulumi.Input['SubscriberNotificationHttpsNotificationConfigurationHttpMethod'] http_method: The HTTPS method used for the notification subscription.
        """
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "target_role_arn", target_role_arn)
        if authorization_api_key_name is not None:
            pulumi.set(__self__, "authorization_api_key_name", authorization_api_key_name)
        if authorization_api_key_value is not None:
            pulumi.set(__self__, "authorization_api_key_value", authorization_api_key_value)
        if http_method is not None:
            pulumi.set(__self__, "http_method", http_method)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Input[builtins.str]:
        """
        The subscription endpoint in Security Lake.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter(name="targetRoleArn")
    def target_role_arn(self) -> pulumi.Input[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the EventBridge API destinations IAM role that you created.
        """
        return pulumi.get(self, "target_role_arn")

    @target_role_arn.setter
    def target_role_arn(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "target_role_arn", value)

    @property
    @pulumi.getter(name="authorizationApiKeyName")
    def authorization_api_key_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The key name for the notification subscription.
        """
        return pulumi.get(self, "authorization_api_key_name")

    @authorization_api_key_name.setter
    def authorization_api_key_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "authorization_api_key_name", value)

    @property
    @pulumi.getter(name="authorizationApiKeyValue")
    def authorization_api_key_value(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The key value for the notification subscription.
        """
        return pulumi.get(self, "authorization_api_key_value")

    @authorization_api_key_value.setter
    def authorization_api_key_value(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "authorization_api_key_value", value)

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> Optional[pulumi.Input['SubscriberNotificationHttpsNotificationConfigurationHttpMethod']]:
        """
        The HTTPS method used for the notification subscription.
        """
        return pulumi.get(self, "http_method")

    @http_method.setter
    def http_method(self, value: Optional[pulumi.Input['SubscriberNotificationHttpsNotificationConfigurationHttpMethod']]):
        pulumi.set(self, "http_method", value)


if not MYPY:
    class SubscriberNotificationNotificationConfigurationArgsDict(TypedDict):
        https_notification_configuration: NotRequired[pulumi.Input['SubscriberNotificationHttpsNotificationConfigurationArgsDict']]
        """
        The configurations used for HTTPS subscriber notification.
        """
        sqs_notification_configuration: NotRequired[pulumi.Input['SubscriberNotificationSqsNotificationConfigurationArgsDict']]
        """
        The configurations for SQS subscriber notification. The members of this structure are context-dependent.
        """
elif False:
    SubscriberNotificationNotificationConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class SubscriberNotificationNotificationConfigurationArgs:
    def __init__(__self__, *,
                 https_notification_configuration: Optional[pulumi.Input['SubscriberNotificationHttpsNotificationConfigurationArgs']] = None,
                 sqs_notification_configuration: Optional[pulumi.Input['SubscriberNotificationSqsNotificationConfigurationArgs']] = None):
        """
        :param pulumi.Input['SubscriberNotificationHttpsNotificationConfigurationArgs'] https_notification_configuration: The configurations used for HTTPS subscriber notification.
        :param pulumi.Input['SubscriberNotificationSqsNotificationConfigurationArgs'] sqs_notification_configuration: The configurations for SQS subscriber notification. The members of this structure are context-dependent.
        """
        if https_notification_configuration is not None:
            pulumi.set(__self__, "https_notification_configuration", https_notification_configuration)
        if sqs_notification_configuration is not None:
            pulumi.set(__self__, "sqs_notification_configuration", sqs_notification_configuration)

    @property
    @pulumi.getter(name="httpsNotificationConfiguration")
    def https_notification_configuration(self) -> Optional[pulumi.Input['SubscriberNotificationHttpsNotificationConfigurationArgs']]:
        """
        The configurations used for HTTPS subscriber notification.
        """
        return pulumi.get(self, "https_notification_configuration")

    @https_notification_configuration.setter
    def https_notification_configuration(self, value: Optional[pulumi.Input['SubscriberNotificationHttpsNotificationConfigurationArgs']]):
        pulumi.set(self, "https_notification_configuration", value)

    @property
    @pulumi.getter(name="sqsNotificationConfiguration")
    def sqs_notification_configuration(self) -> Optional[pulumi.Input['SubscriberNotificationSqsNotificationConfigurationArgs']]:
        """
        The configurations for SQS subscriber notification. The members of this structure are context-dependent.
        """
        return pulumi.get(self, "sqs_notification_configuration")

    @sqs_notification_configuration.setter
    def sqs_notification_configuration(self, value: Optional[pulumi.Input['SubscriberNotificationSqsNotificationConfigurationArgs']]):
        pulumi.set(self, "sqs_notification_configuration", value)


if not MYPY:
    class SubscriberNotificationSqsNotificationConfigurationArgsDict(TypedDict):
        """
        The configurations for SQS subscriber notification. The members of this structure are context-dependent.
        """
        pass
elif False:
    SubscriberNotificationSqsNotificationConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class SubscriberNotificationSqsNotificationConfigurationArgs:
    def __init__(__self__):
        """
        The configurations for SQS subscriber notification. The members of this structure are context-dependent.
        """
        pass


if not MYPY:
    class SubscriberSourceArgsDict(TypedDict):
        aws_log_source: NotRequired[pulumi.Input['SubscriberAwsLogSourceArgsDict']]
        custom_log_source: NotRequired[pulumi.Input['SubscriberCustomLogSourceArgsDict']]
elif False:
    SubscriberSourceArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class SubscriberSourceArgs:
    def __init__(__self__, *,
                 aws_log_source: Optional[pulumi.Input['SubscriberAwsLogSourceArgs']] = None,
                 custom_log_source: Optional[pulumi.Input['SubscriberCustomLogSourceArgs']] = None):
        if aws_log_source is not None:
            pulumi.set(__self__, "aws_log_source", aws_log_source)
        if custom_log_source is not None:
            pulumi.set(__self__, "custom_log_source", custom_log_source)

    @property
    @pulumi.getter(name="awsLogSource")
    def aws_log_source(self) -> Optional[pulumi.Input['SubscriberAwsLogSourceArgs']]:
        return pulumi.get(self, "aws_log_source")

    @aws_log_source.setter
    def aws_log_source(self, value: Optional[pulumi.Input['SubscriberAwsLogSourceArgs']]):
        pulumi.set(self, "aws_log_source", value)

    @property
    @pulumi.getter(name="customLogSource")
    def custom_log_source(self) -> Optional[pulumi.Input['SubscriberCustomLogSourceArgs']]:
        return pulumi.get(self, "custom_log_source")

    @custom_log_source.setter
    def custom_log_source(self, value: Optional[pulumi.Input['SubscriberCustomLogSourceArgs']]):
        pulumi.set(self, "custom_log_source", value)


