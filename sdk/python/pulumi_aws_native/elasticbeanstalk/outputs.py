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

__all__ = [
    'ApplicationMaxAgeRule',
    'ApplicationMaxCountRule',
    'ApplicationResourceLifecycleConfig',
    'ApplicationVersionLifecycleConfig',
    'ApplicationVersionSourceBundle',
    'ConfigurationTemplateConfigurationOptionSetting',
    'ConfigurationTemplateSourceConfiguration',
    'EnvironmentOptionSetting',
    'EnvironmentTier',
]

@pulumi.output_type
class ApplicationMaxAgeRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "deleteSourceFromS3":
            suggest = "delete_source_from_s3"
        elif key == "maxAgeInDays":
            suggest = "max_age_in_days"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationMaxAgeRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationMaxAgeRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationMaxAgeRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 delete_source_from_s3: Optional[builtins.bool] = None,
                 enabled: Optional[builtins.bool] = None,
                 max_age_in_days: Optional[builtins.int] = None):
        """
        :param builtins.bool delete_source_from_s3: Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.
        :param builtins.bool enabled: Specify true to apply the rule, or false to disable it.
        :param builtins.int max_age_in_days: Specify the number of days to retain an application versions.
        """
        if delete_source_from_s3 is not None:
            pulumi.set(__self__, "delete_source_from_s3", delete_source_from_s3)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if max_age_in_days is not None:
            pulumi.set(__self__, "max_age_in_days", max_age_in_days)

    @property
    @pulumi.getter(name="deleteSourceFromS3")
    def delete_source_from_s3(self) -> Optional[builtins.bool]:
        """
        Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.
        """
        return pulumi.get(self, "delete_source_from_s3")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[builtins.bool]:
        """
        Specify true to apply the rule, or false to disable it.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="maxAgeInDays")
    def max_age_in_days(self) -> Optional[builtins.int]:
        """
        Specify the number of days to retain an application versions.
        """
        return pulumi.get(self, "max_age_in_days")


@pulumi.output_type
class ApplicationMaxCountRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "deleteSourceFromS3":
            suggest = "delete_source_from_s3"
        elif key == "maxCount":
            suggest = "max_count"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationMaxCountRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationMaxCountRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationMaxCountRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 delete_source_from_s3: Optional[builtins.bool] = None,
                 enabled: Optional[builtins.bool] = None,
                 max_count: Optional[builtins.int] = None):
        """
        :param builtins.bool delete_source_from_s3: Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.
        :param builtins.bool enabled: Specify true to apply the rule, or false to disable it.
        :param builtins.int max_count: Specify the maximum number of application versions to retain.
        """
        if delete_source_from_s3 is not None:
            pulumi.set(__self__, "delete_source_from_s3", delete_source_from_s3)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if max_count is not None:
            pulumi.set(__self__, "max_count", max_count)

    @property
    @pulumi.getter(name="deleteSourceFromS3")
    def delete_source_from_s3(self) -> Optional[builtins.bool]:
        """
        Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.
        """
        return pulumi.get(self, "delete_source_from_s3")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[builtins.bool]:
        """
        Specify true to apply the rule, or false to disable it.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="maxCount")
    def max_count(self) -> Optional[builtins.int]:
        """
        Specify the maximum number of application versions to retain.
        """
        return pulumi.get(self, "max_count")


@pulumi.output_type
class ApplicationResourceLifecycleConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "serviceRole":
            suggest = "service_role"
        elif key == "versionLifecycleConfig":
            suggest = "version_lifecycle_config"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationResourceLifecycleConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationResourceLifecycleConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationResourceLifecycleConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 service_role: Optional[builtins.str] = None,
                 version_lifecycle_config: Optional['outputs.ApplicationVersionLifecycleConfig'] = None):
        """
        :param builtins.str service_role: The ARN of an IAM service role that Elastic Beanstalk has permission to assume. The ServiceRole property is required the first time that you provide a ResourceLifecycleConfig for the application. After you provide it once, Elastic Beanstalk persists the Service Role with the application, and you don't need to specify it again. You can, however, specify it in subsequent updates to change the Service Role to another value.
        :param 'ApplicationVersionLifecycleConfig' version_lifecycle_config: Defines lifecycle settings for application versions.
        """
        if service_role is not None:
            pulumi.set(__self__, "service_role", service_role)
        if version_lifecycle_config is not None:
            pulumi.set(__self__, "version_lifecycle_config", version_lifecycle_config)

    @property
    @pulumi.getter(name="serviceRole")
    def service_role(self) -> Optional[builtins.str]:
        """
        The ARN of an IAM service role that Elastic Beanstalk has permission to assume. The ServiceRole property is required the first time that you provide a ResourceLifecycleConfig for the application. After you provide it once, Elastic Beanstalk persists the Service Role with the application, and you don't need to specify it again. You can, however, specify it in subsequent updates to change the Service Role to another value.
        """
        return pulumi.get(self, "service_role")

    @property
    @pulumi.getter(name="versionLifecycleConfig")
    def version_lifecycle_config(self) -> Optional['outputs.ApplicationVersionLifecycleConfig']:
        """
        Defines lifecycle settings for application versions.
        """
        return pulumi.get(self, "version_lifecycle_config")


@pulumi.output_type
class ApplicationVersionLifecycleConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "maxAgeRule":
            suggest = "max_age_rule"
        elif key == "maxCountRule":
            suggest = "max_count_rule"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationVersionLifecycleConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationVersionLifecycleConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationVersionLifecycleConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 max_age_rule: Optional['outputs.ApplicationMaxAgeRule'] = None,
                 max_count_rule: Optional['outputs.ApplicationMaxCountRule'] = None):
        """
        :param 'ApplicationMaxAgeRule' max_age_rule: Specify a max age rule to restrict the length of time that application versions are retained for an application.
        :param 'ApplicationMaxCountRule' max_count_rule: Specify a max count rule to restrict the number of application versions that are retained for an application.
        """
        if max_age_rule is not None:
            pulumi.set(__self__, "max_age_rule", max_age_rule)
        if max_count_rule is not None:
            pulumi.set(__self__, "max_count_rule", max_count_rule)

    @property
    @pulumi.getter(name="maxAgeRule")
    def max_age_rule(self) -> Optional['outputs.ApplicationMaxAgeRule']:
        """
        Specify a max age rule to restrict the length of time that application versions are retained for an application.
        """
        return pulumi.get(self, "max_age_rule")

    @property
    @pulumi.getter(name="maxCountRule")
    def max_count_rule(self) -> Optional['outputs.ApplicationMaxCountRule']:
        """
        Specify a max count rule to restrict the number of application versions that are retained for an application.
        """
        return pulumi.get(self, "max_count_rule")


@pulumi.output_type
class ApplicationVersionSourceBundle(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "s3Bucket":
            suggest = "s3_bucket"
        elif key == "s3Key":
            suggest = "s3_key"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationVersionSourceBundle. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationVersionSourceBundle.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationVersionSourceBundle.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 s3_bucket: builtins.str,
                 s3_key: builtins.str):
        """
        :param builtins.str s3_bucket: The Amazon S3 bucket where the data is located.
        :param builtins.str s3_key: The Amazon S3 key where the data is located.
        """
        pulumi.set(__self__, "s3_bucket", s3_bucket)
        pulumi.set(__self__, "s3_key", s3_key)

    @property
    @pulumi.getter(name="s3Bucket")
    def s3_bucket(self) -> builtins.str:
        """
        The Amazon S3 bucket where the data is located.
        """
        return pulumi.get(self, "s3_bucket")

    @property
    @pulumi.getter(name="s3Key")
    def s3_key(self) -> builtins.str:
        """
        The Amazon S3 key where the data is located.
        """
        return pulumi.get(self, "s3_key")


@pulumi.output_type
class ConfigurationTemplateConfigurationOptionSetting(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "optionName":
            suggest = "option_name"
        elif key == "resourceName":
            suggest = "resource_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigurationTemplateConfigurationOptionSetting. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigurationTemplateConfigurationOptionSetting.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigurationTemplateConfigurationOptionSetting.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 namespace: builtins.str,
                 option_name: builtins.str,
                 resource_name: Optional[builtins.str] = None,
                 value: Optional[builtins.str] = None):
        """
        :param builtins.str namespace: A unique namespace that identifies the option's associated AWS resource.
        :param builtins.str option_name: The name of the configuration option.
        :param builtins.str resource_name: A unique resource name for the option setting. Use it for a time–based scaling configuration option. 
        :param builtins.str value: The current value for the configuration option.
        """
        pulumi.set(__self__, "namespace", namespace)
        pulumi.set(__self__, "option_name", option_name)
        if resource_name is not None:
            pulumi.set(__self__, "resource_name", resource_name)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def namespace(self) -> builtins.str:
        """
        A unique namespace that identifies the option's associated AWS resource.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="optionName")
    def option_name(self) -> builtins.str:
        """
        The name of the configuration option.
        """
        return pulumi.get(self, "option_name")

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> Optional[builtins.str]:
        """
        A unique resource name for the option setting. Use it for a time–based scaling configuration option. 
        """
        return pulumi.get(self, "resource_name")

    @property
    @pulumi.getter
    def value(self) -> Optional[builtins.str]:
        """
        The current value for the configuration option.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class ConfigurationTemplateSourceConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "applicationName":
            suggest = "application_name"
        elif key == "templateName":
            suggest = "template_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigurationTemplateSourceConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigurationTemplateSourceConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigurationTemplateSourceConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 application_name: builtins.str,
                 template_name: builtins.str):
        """
        :param builtins.str application_name: The name of the application associated with the configuration.
        :param builtins.str template_name: The name of the configuration template.
        """
        pulumi.set(__self__, "application_name", application_name)
        pulumi.set(__self__, "template_name", template_name)

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> builtins.str:
        """
        The name of the application associated with the configuration.
        """
        return pulumi.get(self, "application_name")

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> builtins.str:
        """
        The name of the configuration template.
        """
        return pulumi.get(self, "template_name")


@pulumi.output_type
class EnvironmentOptionSetting(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "optionName":
            suggest = "option_name"
        elif key == "resourceName":
            suggest = "resource_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in EnvironmentOptionSetting. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        EnvironmentOptionSetting.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        EnvironmentOptionSetting.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 namespace: builtins.str,
                 option_name: builtins.str,
                 resource_name: Optional[builtins.str] = None,
                 value: Optional[builtins.str] = None):
        """
        :param builtins.str namespace: A unique namespace that identifies the option's associated AWS resource.
        :param builtins.str option_name: The name of the configuration option.
        :param builtins.str resource_name: A unique resource name for the option setting. Use it for a time–based scaling configuration option.
        :param builtins.str value: The current value for the configuration option.
        """
        pulumi.set(__self__, "namespace", namespace)
        pulumi.set(__self__, "option_name", option_name)
        if resource_name is not None:
            pulumi.set(__self__, "resource_name", resource_name)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def namespace(self) -> builtins.str:
        """
        A unique namespace that identifies the option's associated AWS resource.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="optionName")
    def option_name(self) -> builtins.str:
        """
        The name of the configuration option.
        """
        return pulumi.get(self, "option_name")

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> Optional[builtins.str]:
        """
        A unique resource name for the option setting. Use it for a time–based scaling configuration option.
        """
        return pulumi.get(self, "resource_name")

    @property
    @pulumi.getter
    def value(self) -> Optional[builtins.str]:
        """
        The current value for the configuration option.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class EnvironmentTier(dict):
    def __init__(__self__, *,
                 name: Optional[builtins.str] = None,
                 type: Optional[builtins.str] = None,
                 version: Optional[builtins.str] = None):
        """
        :param builtins.str name: The name of this environment tier.
        :param builtins.str type: The type of this environment tier.
        :param builtins.str version: The version of this environment tier. When you don't set a value to it, Elastic Beanstalk uses the latest compatible worker tier version.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        """
        The name of this environment tier.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> Optional[builtins.str]:
        """
        The type of this environment tier.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> Optional[builtins.str]:
        """
        The version of this environment tier. When you don't set a value to it, Elastic Beanstalk uses the latest compatible worker tier version.
        """
        return pulumi.get(self, "version")


