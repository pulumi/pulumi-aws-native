# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['ApplicationArgs', 'Application']

@pulumi.input_type
class ApplicationArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[str],
                 auto_configuration_enabled: Optional[pulumi.Input[bool]] = None,
                 c_we_monitor_enabled: Optional[pulumi.Input[bool]] = None,
                 component_monitoring_settings: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationComponentMonitoringSettingArgs']]]] = None,
                 custom_components: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationCustomComponentArgs']]]] = None,
                 grouping_type: Optional[pulumi.Input['ApplicationGroupingType']] = None,
                 log_pattern_sets: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationLogPatternSetArgs']]]] = None,
                 ops_center_enabled: Optional[pulumi.Input[bool]] = None,
                 ops_item_sns_topic_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationTagArgs']]]] = None):
        """
        The set of arguments for constructing a Application resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[bool] auto_configuration_enabled: If set to true, application will be configured with recommended monitoring configuration.
        :param pulumi.Input[bool] c_we_monitor_enabled: Indicates whether Application Insights can listen to CloudWatch events for the application resources.
        :param pulumi.Input[Sequence[pulumi.Input['ApplicationComponentMonitoringSettingArgs']]] component_monitoring_settings: The monitoring settings of the components.
        :param pulumi.Input[Sequence[pulumi.Input['ApplicationCustomComponentArgs']]] custom_components: The custom grouped components.
        :param pulumi.Input['ApplicationGroupingType'] grouping_type: The grouping type of the application
        :param pulumi.Input[Sequence[pulumi.Input['ApplicationLogPatternSetArgs']]] log_pattern_sets: The log pattern sets.
        :param pulumi.Input[bool] ops_center_enabled: When set to true, creates opsItems for any problems detected on an application.
        :param pulumi.Input[str] ops_item_sns_topic_arn: The SNS topic provided to Application Insights that is associated to the created opsItem.
        :param pulumi.Input[Sequence[pulumi.Input['ApplicationTagArgs']]] tags: The tags of Application Insights application.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if auto_configuration_enabled is not None:
            pulumi.set(__self__, "auto_configuration_enabled", auto_configuration_enabled)
        if c_we_monitor_enabled is not None:
            pulumi.set(__self__, "c_we_monitor_enabled", c_we_monitor_enabled)
        if component_monitoring_settings is not None:
            pulumi.set(__self__, "component_monitoring_settings", component_monitoring_settings)
        if custom_components is not None:
            pulumi.set(__self__, "custom_components", custom_components)
        if grouping_type is not None:
            pulumi.set(__self__, "grouping_type", grouping_type)
        if log_pattern_sets is not None:
            pulumi.set(__self__, "log_pattern_sets", log_pattern_sets)
        if ops_center_enabled is not None:
            pulumi.set(__self__, "ops_center_enabled", ops_center_enabled)
        if ops_item_sns_topic_arn is not None:
            pulumi.set(__self__, "ops_item_sns_topic_arn", ops_item_sns_topic_arn)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="autoConfigurationEnabled")
    def auto_configuration_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to true, application will be configured with recommended monitoring configuration.
        """
        return pulumi.get(self, "auto_configuration_enabled")

    @auto_configuration_enabled.setter
    def auto_configuration_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_configuration_enabled", value)

    @property
    @pulumi.getter(name="cWEMonitorEnabled")
    def c_we_monitor_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether Application Insights can listen to CloudWatch events for the application resources.
        """
        return pulumi.get(self, "c_we_monitor_enabled")

    @c_we_monitor_enabled.setter
    def c_we_monitor_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "c_we_monitor_enabled", value)

    @property
    @pulumi.getter(name="componentMonitoringSettings")
    def component_monitoring_settings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationComponentMonitoringSettingArgs']]]]:
        """
        The monitoring settings of the components.
        """
        return pulumi.get(self, "component_monitoring_settings")

    @component_monitoring_settings.setter
    def component_monitoring_settings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationComponentMonitoringSettingArgs']]]]):
        pulumi.set(self, "component_monitoring_settings", value)

    @property
    @pulumi.getter(name="customComponents")
    def custom_components(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationCustomComponentArgs']]]]:
        """
        The custom grouped components.
        """
        return pulumi.get(self, "custom_components")

    @custom_components.setter
    def custom_components(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationCustomComponentArgs']]]]):
        pulumi.set(self, "custom_components", value)

    @property
    @pulumi.getter(name="groupingType")
    def grouping_type(self) -> Optional[pulumi.Input['ApplicationGroupingType']]:
        """
        The grouping type of the application
        """
        return pulumi.get(self, "grouping_type")

    @grouping_type.setter
    def grouping_type(self, value: Optional[pulumi.Input['ApplicationGroupingType']]):
        pulumi.set(self, "grouping_type", value)

    @property
    @pulumi.getter(name="logPatternSets")
    def log_pattern_sets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationLogPatternSetArgs']]]]:
        """
        The log pattern sets.
        """
        return pulumi.get(self, "log_pattern_sets")

    @log_pattern_sets.setter
    def log_pattern_sets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationLogPatternSetArgs']]]]):
        pulumi.set(self, "log_pattern_sets", value)

    @property
    @pulumi.getter(name="opsCenterEnabled")
    def ops_center_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        When set to true, creates opsItems for any problems detected on an application.
        """
        return pulumi.get(self, "ops_center_enabled")

    @ops_center_enabled.setter
    def ops_center_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ops_center_enabled", value)

    @property
    @pulumi.getter(name="opsItemSNSTopicArn")
    def ops_item_sns_topic_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The SNS topic provided to Application Insights that is associated to the created opsItem.
        """
        return pulumi.get(self, "ops_item_sns_topic_arn")

    @ops_item_sns_topic_arn.setter
    def ops_item_sns_topic_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ops_item_sns_topic_arn", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationTagArgs']]]]:
        """
        The tags of Application Insights application.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationTagArgs']]]]):
        pulumi.set(self, "tags", value)


class Application(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_configuration_enabled: Optional[pulumi.Input[bool]] = None,
                 c_we_monitor_enabled: Optional[pulumi.Input[bool]] = None,
                 component_monitoring_settings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationComponentMonitoringSettingArgs']]]]] = None,
                 custom_components: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationCustomComponentArgs']]]]] = None,
                 grouping_type: Optional[pulumi.Input['ApplicationGroupingType']] = None,
                 log_pattern_sets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationLogPatternSetArgs']]]]] = None,
                 ops_center_enabled: Optional[pulumi.Input[bool]] = None,
                 ops_item_sns_topic_arn: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationTagArgs']]]]] = None,
                 __props__=None):
        """
        Resource schema for AWS::ApplicationInsights::Application

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_configuration_enabled: If set to true, application will be configured with recommended monitoring configuration.
        :param pulumi.Input[bool] c_we_monitor_enabled: Indicates whether Application Insights can listen to CloudWatch events for the application resources.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationComponentMonitoringSettingArgs']]]] component_monitoring_settings: The monitoring settings of the components.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationCustomComponentArgs']]]] custom_components: The custom grouped components.
        :param pulumi.Input['ApplicationGroupingType'] grouping_type: The grouping type of the application
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationLogPatternSetArgs']]]] log_pattern_sets: The log pattern sets.
        :param pulumi.Input[bool] ops_center_enabled: When set to true, creates opsItems for any problems detected on an application.
        :param pulumi.Input[str] ops_item_sns_topic_arn: The SNS topic provided to Application Insights that is associated to the created opsItem.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationTagArgs']]]] tags: The tags of Application Insights application.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource schema for AWS::ApplicationInsights::Application

        :param str resource_name: The name of the resource.
        :param ApplicationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_configuration_enabled: Optional[pulumi.Input[bool]] = None,
                 c_we_monitor_enabled: Optional[pulumi.Input[bool]] = None,
                 component_monitoring_settings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationComponentMonitoringSettingArgs']]]]] = None,
                 custom_components: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationCustomComponentArgs']]]]] = None,
                 grouping_type: Optional[pulumi.Input['ApplicationGroupingType']] = None,
                 log_pattern_sets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationLogPatternSetArgs']]]]] = None,
                 ops_center_enabled: Optional[pulumi.Input[bool]] = None,
                 ops_item_sns_topic_arn: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationArgs.__new__(ApplicationArgs)

            __props__.__dict__["auto_configuration_enabled"] = auto_configuration_enabled
            __props__.__dict__["c_we_monitor_enabled"] = c_we_monitor_enabled
            __props__.__dict__["component_monitoring_settings"] = component_monitoring_settings
            __props__.__dict__["custom_components"] = custom_components
            __props__.__dict__["grouping_type"] = grouping_type
            __props__.__dict__["log_pattern_sets"] = log_pattern_sets
            __props__.__dict__["ops_center_enabled"] = ops_center_enabled
            __props__.__dict__["ops_item_sns_topic_arn"] = ops_item_sns_topic_arn
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["application_arn"] = None
        super(Application, __self__).__init__(
            'aws-native:applicationinsights:Application',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Application':
        """
        Get an existing Application resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ApplicationArgs.__new__(ApplicationArgs)

        __props__.__dict__["application_arn"] = None
        __props__.__dict__["auto_configuration_enabled"] = None
        __props__.__dict__["c_we_monitor_enabled"] = None
        __props__.__dict__["component_monitoring_settings"] = None
        __props__.__dict__["custom_components"] = None
        __props__.__dict__["grouping_type"] = None
        __props__.__dict__["log_pattern_sets"] = None
        __props__.__dict__["ops_center_enabled"] = None
        __props__.__dict__["ops_item_sns_topic_arn"] = None
        __props__.__dict__["resource_group_name"] = None
        __props__.__dict__["tags"] = None
        return Application(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationARN")
    def application_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the ApplicationInsights application.
        """
        return pulumi.get(self, "application_arn")

    @property
    @pulumi.getter(name="autoConfigurationEnabled")
    def auto_configuration_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        If set to true, application will be configured with recommended monitoring configuration.
        """
        return pulumi.get(self, "auto_configuration_enabled")

    @property
    @pulumi.getter(name="cWEMonitorEnabled")
    def c_we_monitor_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether Application Insights can listen to CloudWatch events for the application resources.
        """
        return pulumi.get(self, "c_we_monitor_enabled")

    @property
    @pulumi.getter(name="componentMonitoringSettings")
    def component_monitoring_settings(self) -> pulumi.Output[Optional[Sequence['outputs.ApplicationComponentMonitoringSetting']]]:
        """
        The monitoring settings of the components.
        """
        return pulumi.get(self, "component_monitoring_settings")

    @property
    @pulumi.getter(name="customComponents")
    def custom_components(self) -> pulumi.Output[Optional[Sequence['outputs.ApplicationCustomComponent']]]:
        """
        The custom grouped components.
        """
        return pulumi.get(self, "custom_components")

    @property
    @pulumi.getter(name="groupingType")
    def grouping_type(self) -> pulumi.Output[Optional['ApplicationGroupingType']]:
        """
        The grouping type of the application
        """
        return pulumi.get(self, "grouping_type")

    @property
    @pulumi.getter(name="logPatternSets")
    def log_pattern_sets(self) -> pulumi.Output[Optional[Sequence['outputs.ApplicationLogPatternSet']]]:
        """
        The log pattern sets.
        """
        return pulumi.get(self, "log_pattern_sets")

    @property
    @pulumi.getter(name="opsCenterEnabled")
    def ops_center_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        When set to true, creates opsItems for any problems detected on an application.
        """
        return pulumi.get(self, "ops_center_enabled")

    @property
    @pulumi.getter(name="opsItemSNSTopicArn")
    def ops_item_sns_topic_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The SNS topic provided to Application Insights that is associated to the created opsItem.
        """
        return pulumi.get(self, "ops_item_sns_topic_arn")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.ApplicationTag']]]:
        """
        The tags of Application Insights application.
        """
        return pulumi.get(self, "tags")
