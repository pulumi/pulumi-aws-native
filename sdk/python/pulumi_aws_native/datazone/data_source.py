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

__all__ = ['DataSourceArgs', 'DataSource']

@pulumi.input_type
class DataSourceArgs:
    def __init__(__self__, *,
                 domain_identifier: pulumi.Input[str],
                 environment_identifier: pulumi.Input[str],
                 project_identifier: pulumi.Input[str],
                 type: pulumi.Input[str],
                 asset_forms_input: Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceFormInputArgs']]]] = None,
                 configuration: Optional[pulumi.Input[Union['DataSourceConfigurationInput0PropertiesArgs', 'DataSourceConfigurationInput1PropertiesArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_setting: Optional[pulumi.Input['DataSourceEnableSetting']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 publish_on_import: Optional[pulumi.Input[bool]] = None,
                 recommendation: Optional[pulumi.Input['DataSourceRecommendationConfigurationArgs']] = None,
                 schedule: Optional[pulumi.Input['DataSourceScheduleConfigurationArgs']] = None):
        """
        The set of arguments for constructing a DataSource resource.
        :param pulumi.Input[str] domain_identifier: The ID of the Amazon DataZone domain where the data source is created.
        :param pulumi.Input[str] environment_identifier: The unique identifier of the Amazon DataZone environment to which the data source publishes assets.
        :param pulumi.Input[str] project_identifier: The identifier of the Amazon DataZone project in which you want to add the data source.
        :param pulumi.Input[str] type: The type of the data source.
        :param pulumi.Input[Sequence[pulumi.Input['DataSourceFormInputArgs']]] asset_forms_input: The metadata forms that are to be attached to the assets that this data source works with.
        :param pulumi.Input[Union['DataSourceConfigurationInput0PropertiesArgs', 'DataSourceConfigurationInput1PropertiesArgs']] configuration: Configuration of the data source. It can be set to either glueRunConfiguration or redshiftRunConfiguration.
        :param pulumi.Input[str] description: The description of the data source.
        :param pulumi.Input['DataSourceEnableSetting'] enable_setting: Specifies whether the data source is enabled.
        :param pulumi.Input[str] name: The name of the data source.
        :param pulumi.Input[bool] publish_on_import: Specifies whether the assets that this data source creates in the inventory are to be also automatically published to the catalog.
        :param pulumi.Input['DataSourceRecommendationConfigurationArgs'] recommendation: Specifies whether the business name generation is to be enabled for this data source.
        :param pulumi.Input['DataSourceScheduleConfigurationArgs'] schedule: The schedule of the data source runs.
        """
        pulumi.set(__self__, "domain_identifier", domain_identifier)
        pulumi.set(__self__, "environment_identifier", environment_identifier)
        pulumi.set(__self__, "project_identifier", project_identifier)
        pulumi.set(__self__, "type", type)
        if asset_forms_input is not None:
            pulumi.set(__self__, "asset_forms_input", asset_forms_input)
        if configuration is not None:
            pulumi.set(__self__, "configuration", configuration)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_setting is not None:
            pulumi.set(__self__, "enable_setting", enable_setting)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if publish_on_import is not None:
            pulumi.set(__self__, "publish_on_import", publish_on_import)
        if recommendation is not None:
            pulumi.set(__self__, "recommendation", recommendation)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)

    @property
    @pulumi.getter(name="domainIdentifier")
    def domain_identifier(self) -> pulumi.Input[str]:
        """
        The ID of the Amazon DataZone domain where the data source is created.
        """
        return pulumi.get(self, "domain_identifier")

    @domain_identifier.setter
    def domain_identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain_identifier", value)

    @property
    @pulumi.getter(name="environmentIdentifier")
    def environment_identifier(self) -> pulumi.Input[str]:
        """
        The unique identifier of the Amazon DataZone environment to which the data source publishes assets.
        """
        return pulumi.get(self, "environment_identifier")

    @environment_identifier.setter
    def environment_identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "environment_identifier", value)

    @property
    @pulumi.getter(name="projectIdentifier")
    def project_identifier(self) -> pulumi.Input[str]:
        """
        The identifier of the Amazon DataZone project in which you want to add the data source.
        """
        return pulumi.get(self, "project_identifier")

    @project_identifier.setter
    def project_identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_identifier", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of the data source.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="assetFormsInput")
    def asset_forms_input(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceFormInputArgs']]]]:
        """
        The metadata forms that are to be attached to the assets that this data source works with.
        """
        return pulumi.get(self, "asset_forms_input")

    @asset_forms_input.setter
    def asset_forms_input(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceFormInputArgs']]]]):
        pulumi.set(self, "asset_forms_input", value)

    @property
    @pulumi.getter
    def configuration(self) -> Optional[pulumi.Input[Union['DataSourceConfigurationInput0PropertiesArgs', 'DataSourceConfigurationInput1PropertiesArgs']]]:
        """
        Configuration of the data source. It can be set to either glueRunConfiguration or redshiftRunConfiguration.
        """
        return pulumi.get(self, "configuration")

    @configuration.setter
    def configuration(self, value: Optional[pulumi.Input[Union['DataSourceConfigurationInput0PropertiesArgs', 'DataSourceConfigurationInput1PropertiesArgs']]]):
        pulumi.set(self, "configuration", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the data source.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enableSetting")
    def enable_setting(self) -> Optional[pulumi.Input['DataSourceEnableSetting']]:
        """
        Specifies whether the data source is enabled.
        """
        return pulumi.get(self, "enable_setting")

    @enable_setting.setter
    def enable_setting(self, value: Optional[pulumi.Input['DataSourceEnableSetting']]):
        pulumi.set(self, "enable_setting", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the data source.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="publishOnImport")
    def publish_on_import(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether the assets that this data source creates in the inventory are to be also automatically published to the catalog.
        """
        return pulumi.get(self, "publish_on_import")

    @publish_on_import.setter
    def publish_on_import(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "publish_on_import", value)

    @property
    @pulumi.getter
    def recommendation(self) -> Optional[pulumi.Input['DataSourceRecommendationConfigurationArgs']]:
        """
        Specifies whether the business name generation is to be enabled for this data source.
        """
        return pulumi.get(self, "recommendation")

    @recommendation.setter
    def recommendation(self, value: Optional[pulumi.Input['DataSourceRecommendationConfigurationArgs']]):
        pulumi.set(self, "recommendation", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input['DataSourceScheduleConfigurationArgs']]:
        """
        The schedule of the data source runs.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input['DataSourceScheduleConfigurationArgs']]):
        pulumi.set(self, "schedule", value)


class DataSource(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asset_forms_input: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceFormInputArgs']]]]] = None,
                 configuration: Optional[pulumi.Input[Union[pulumi.InputType['DataSourceConfigurationInput0PropertiesArgs'], pulumi.InputType['DataSourceConfigurationInput1PropertiesArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain_identifier: Optional[pulumi.Input[str]] = None,
                 enable_setting: Optional[pulumi.Input['DataSourceEnableSetting']] = None,
                 environment_identifier: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_identifier: Optional[pulumi.Input[str]] = None,
                 publish_on_import: Optional[pulumi.Input[bool]] = None,
                 recommendation: Optional[pulumi.Input[pulumi.InputType['DataSourceRecommendationConfigurationArgs']]] = None,
                 schedule: Optional[pulumi.Input[pulumi.InputType['DataSourceScheduleConfigurationArgs']]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Definition of AWS::DataZone::DataSource Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceFormInputArgs']]]] asset_forms_input: The metadata forms that are to be attached to the assets that this data source works with.
        :param pulumi.Input[Union[pulumi.InputType['DataSourceConfigurationInput0PropertiesArgs'], pulumi.InputType['DataSourceConfigurationInput1PropertiesArgs']]] configuration: Configuration of the data source. It can be set to either glueRunConfiguration or redshiftRunConfiguration.
        :param pulumi.Input[str] description: The description of the data source.
        :param pulumi.Input[str] domain_identifier: The ID of the Amazon DataZone domain where the data source is created.
        :param pulumi.Input['DataSourceEnableSetting'] enable_setting: Specifies whether the data source is enabled.
        :param pulumi.Input[str] environment_identifier: The unique identifier of the Amazon DataZone environment to which the data source publishes assets.
        :param pulumi.Input[str] name: The name of the data source.
        :param pulumi.Input[str] project_identifier: The identifier of the Amazon DataZone project in which you want to add the data source.
        :param pulumi.Input[bool] publish_on_import: Specifies whether the assets that this data source creates in the inventory are to be also automatically published to the catalog.
        :param pulumi.Input[pulumi.InputType['DataSourceRecommendationConfigurationArgs']] recommendation: Specifies whether the business name generation is to be enabled for this data source.
        :param pulumi.Input[pulumi.InputType['DataSourceScheduleConfigurationArgs']] schedule: The schedule of the data source runs.
        :param pulumi.Input[str] type: The type of the data source.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DataSourceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::DataZone::DataSource Resource Type

        :param str resource_name: The name of the resource.
        :param DataSourceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataSourceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asset_forms_input: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceFormInputArgs']]]]] = None,
                 configuration: Optional[pulumi.Input[Union[pulumi.InputType['DataSourceConfigurationInput0PropertiesArgs'], pulumi.InputType['DataSourceConfigurationInput1PropertiesArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain_identifier: Optional[pulumi.Input[str]] = None,
                 enable_setting: Optional[pulumi.Input['DataSourceEnableSetting']] = None,
                 environment_identifier: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_identifier: Optional[pulumi.Input[str]] = None,
                 publish_on_import: Optional[pulumi.Input[bool]] = None,
                 recommendation: Optional[pulumi.Input[pulumi.InputType['DataSourceRecommendationConfigurationArgs']]] = None,
                 schedule: Optional[pulumi.Input[pulumi.InputType['DataSourceScheduleConfigurationArgs']]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DataSourceArgs.__new__(DataSourceArgs)

            __props__.__dict__["asset_forms_input"] = asset_forms_input
            __props__.__dict__["configuration"] = configuration
            __props__.__dict__["description"] = description
            if domain_identifier is None and not opts.urn:
                raise TypeError("Missing required property 'domain_identifier'")
            __props__.__dict__["domain_identifier"] = domain_identifier
            __props__.__dict__["enable_setting"] = enable_setting
            if environment_identifier is None and not opts.urn:
                raise TypeError("Missing required property 'environment_identifier'")
            __props__.__dict__["environment_identifier"] = environment_identifier
            __props__.__dict__["name"] = name
            if project_identifier is None and not opts.urn:
                raise TypeError("Missing required property 'project_identifier'")
            __props__.__dict__["project_identifier"] = project_identifier
            __props__.__dict__["publish_on_import"] = publish_on_import
            __props__.__dict__["recommendation"] = recommendation
            __props__.__dict__["schedule"] = schedule
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["created_at"] = None
            __props__.__dict__["domain_id"] = None
            __props__.__dict__["environment_id"] = None
            __props__.__dict__["last_run_asset_count"] = None
            __props__.__dict__["last_run_at"] = None
            __props__.__dict__["last_run_status"] = None
            __props__.__dict__["project_id"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["updated_at"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["domain_identifier", "environment_identifier", "project_identifier", "type"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(DataSource, __self__).__init__(
            'aws-native:datazone:DataSource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DataSource':
        """
        Get an existing DataSource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DataSourceArgs.__new__(DataSourceArgs)

        __props__.__dict__["asset_forms_input"] = None
        __props__.__dict__["configuration"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["domain_id"] = None
        __props__.__dict__["domain_identifier"] = None
        __props__.__dict__["enable_setting"] = None
        __props__.__dict__["environment_id"] = None
        __props__.__dict__["environment_identifier"] = None
        __props__.__dict__["last_run_asset_count"] = None
        __props__.__dict__["last_run_at"] = None
        __props__.__dict__["last_run_status"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project_id"] = None
        __props__.__dict__["project_identifier"] = None
        __props__.__dict__["publish_on_import"] = None
        __props__.__dict__["recommendation"] = None
        __props__.__dict__["schedule"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["updated_at"] = None
        return DataSource(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="assetFormsInput")
    def asset_forms_input(self) -> pulumi.Output[Optional[Sequence['outputs.DataSourceFormInput']]]:
        """
        The metadata forms that are to be attached to the assets that this data source works with.
        """
        return pulumi.get(self, "asset_forms_input")

    @property
    @pulumi.getter
    def configuration(self) -> pulumi.Output[Optional[Any]]:
        """
        Configuration of the data source. It can be set to either glueRunConfiguration or redshiftRunConfiguration.
        """
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The timestamp of when the data source was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the data source.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> pulumi.Output[str]:
        """
        The ID of the Amazon DataZone domain where the data source is created.
        """
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter(name="domainIdentifier")
    def domain_identifier(self) -> pulumi.Output[str]:
        """
        The ID of the Amazon DataZone domain where the data source is created.
        """
        return pulumi.get(self, "domain_identifier")

    @property
    @pulumi.getter(name="enableSetting")
    def enable_setting(self) -> pulumi.Output[Optional['DataSourceEnableSetting']]:
        """
        Specifies whether the data source is enabled.
        """
        return pulumi.get(self, "enable_setting")

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Output[str]:
        """
        The unique identifier of the Amazon DataZone environment to which the data source publishes assets.
        """
        return pulumi.get(self, "environment_id")

    @property
    @pulumi.getter(name="environmentIdentifier")
    def environment_identifier(self) -> pulumi.Output[str]:
        """
        The unique identifier of the Amazon DataZone environment to which the data source publishes assets.
        """
        return pulumi.get(self, "environment_identifier")

    @property
    @pulumi.getter(name="lastRunAssetCount")
    def last_run_asset_count(self) -> pulumi.Output[float]:
        """
        The number of assets created by the data source during its last run.
        """
        return pulumi.get(self, "last_run_asset_count")

    @property
    @pulumi.getter(name="lastRunAt")
    def last_run_at(self) -> pulumi.Output[str]:
        """
        The timestamp that specifies when the data source was last run.
        """
        return pulumi.get(self, "last_run_at")

    @property
    @pulumi.getter(name="lastRunStatus")
    def last_run_status(self) -> pulumi.Output[str]:
        """
        The status of the last run of this data source.
        """
        return pulumi.get(self, "last_run_status")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the data source.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The ID of the Amazon DataZone project to which the data source is added.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="projectIdentifier")
    def project_identifier(self) -> pulumi.Output[str]:
        """
        The identifier of the Amazon DataZone project in which you want to add the data source.
        """
        return pulumi.get(self, "project_identifier")

    @property
    @pulumi.getter(name="publishOnImport")
    def publish_on_import(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether the assets that this data source creates in the inventory are to be also automatically published to the catalog.
        """
        return pulumi.get(self, "publish_on_import")

    @property
    @pulumi.getter
    def recommendation(self) -> pulumi.Output[Optional['outputs.DataSourceRecommendationConfiguration']]:
        """
        Specifies whether the business name generation is to be enabled for this data source.
        """
        return pulumi.get(self, "recommendation")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output[Optional['outputs.DataSourceScheduleConfiguration']]:
        """
        The schedule of the data source runs.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['DataSourceStatus']:
        """
        The status of the data source.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the data source.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The timestamp of when this data source was updated.
        """
        return pulumi.get(self, "updated_at")
