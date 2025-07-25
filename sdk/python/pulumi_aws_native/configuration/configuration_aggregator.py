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
from ._inputs import *

__all__ = ['ConfigurationAggregatorArgs', 'ConfigurationAggregator']

@pulumi.input_type
class ConfigurationAggregatorArgs:
    def __init__(__self__, *,
                 account_aggregation_sources: Optional[pulumi.Input[Sequence[pulumi.Input['ConfigurationAggregatorAccountAggregationSourceArgs']]]] = None,
                 configuration_aggregator_name: Optional[pulumi.Input[builtins.str]] = None,
                 organization_aggregation_source: Optional[pulumi.Input['ConfigurationAggregatorOrganizationAggregationSourceArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a ConfigurationAggregator resource.
        :param pulumi.Input[Sequence[pulumi.Input['ConfigurationAggregatorAccountAggregationSourceArgs']]] account_aggregation_sources: Provides a list of source accounts and regions to be aggregated.
        :param pulumi.Input[builtins.str] configuration_aggregator_name: The name of the aggregator.
        :param pulumi.Input['ConfigurationAggregatorOrganizationAggregationSourceArgs'] organization_aggregation_source: Provides an organization and list of regions to be aggregated.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: The tags for the configuration aggregator.
        """
        if account_aggregation_sources is not None:
            pulumi.set(__self__, "account_aggregation_sources", account_aggregation_sources)
        if configuration_aggregator_name is not None:
            pulumi.set(__self__, "configuration_aggregator_name", configuration_aggregator_name)
        if organization_aggregation_source is not None:
            pulumi.set(__self__, "organization_aggregation_source", organization_aggregation_source)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="accountAggregationSources")
    def account_aggregation_sources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ConfigurationAggregatorAccountAggregationSourceArgs']]]]:
        """
        Provides a list of source accounts and regions to be aggregated.
        """
        return pulumi.get(self, "account_aggregation_sources")

    @account_aggregation_sources.setter
    def account_aggregation_sources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ConfigurationAggregatorAccountAggregationSourceArgs']]]]):
        pulumi.set(self, "account_aggregation_sources", value)

    @property
    @pulumi.getter(name="configurationAggregatorName")
    def configuration_aggregator_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the aggregator.
        """
        return pulumi.get(self, "configuration_aggregator_name")

    @configuration_aggregator_name.setter
    def configuration_aggregator_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "configuration_aggregator_name", value)

    @property
    @pulumi.getter(name="organizationAggregationSource")
    def organization_aggregation_source(self) -> Optional[pulumi.Input['ConfigurationAggregatorOrganizationAggregationSourceArgs']]:
        """
        Provides an organization and list of regions to be aggregated.
        """
        return pulumi.get(self, "organization_aggregation_source")

    @organization_aggregation_source.setter
    def organization_aggregation_source(self, value: Optional[pulumi.Input['ConfigurationAggregatorOrganizationAggregationSourceArgs']]):
        pulumi.set(self, "organization_aggregation_source", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        The tags for the configuration aggregator.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:configuration:ConfigurationAggregator")
class ConfigurationAggregator(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_aggregation_sources: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ConfigurationAggregatorAccountAggregationSourceArgs', 'ConfigurationAggregatorAccountAggregationSourceArgsDict']]]]] = None,
                 configuration_aggregator_name: Optional[pulumi.Input[builtins.str]] = None,
                 organization_aggregation_source: Optional[pulumi.Input[Union['ConfigurationAggregatorOrganizationAggregationSourceArgs', 'ConfigurationAggregatorOrganizationAggregationSourceArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Config::ConfigurationAggregator

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ConfigurationAggregatorAccountAggregationSourceArgs', 'ConfigurationAggregatorAccountAggregationSourceArgsDict']]]] account_aggregation_sources: Provides a list of source accounts and regions to be aggregated.
        :param pulumi.Input[builtins.str] configuration_aggregator_name: The name of the aggregator.
        :param pulumi.Input[Union['ConfigurationAggregatorOrganizationAggregationSourceArgs', 'ConfigurationAggregatorOrganizationAggregationSourceArgsDict']] organization_aggregation_source: Provides an organization and list of regions to be aggregated.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] tags: The tags for the configuration aggregator.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ConfigurationAggregatorArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Config::ConfigurationAggregator

        :param str resource_name: The name of the resource.
        :param ConfigurationAggregatorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConfigurationAggregatorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_aggregation_sources: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ConfigurationAggregatorAccountAggregationSourceArgs', 'ConfigurationAggregatorAccountAggregationSourceArgsDict']]]]] = None,
                 configuration_aggregator_name: Optional[pulumi.Input[builtins.str]] = None,
                 organization_aggregation_source: Optional[pulumi.Input[Union['ConfigurationAggregatorOrganizationAggregationSourceArgs', 'ConfigurationAggregatorOrganizationAggregationSourceArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConfigurationAggregatorArgs.__new__(ConfigurationAggregatorArgs)

            __props__.__dict__["account_aggregation_sources"] = account_aggregation_sources
            __props__.__dict__["configuration_aggregator_name"] = configuration_aggregator_name
            __props__.__dict__["organization_aggregation_source"] = organization_aggregation_source
            __props__.__dict__["tags"] = tags
            __props__.__dict__["configuration_aggregator_arn"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["configurationAggregatorName"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(ConfigurationAggregator, __self__).__init__(
            'aws-native:configuration:ConfigurationAggregator',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ConfigurationAggregator':
        """
        Get an existing ConfigurationAggregator resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ConfigurationAggregatorArgs.__new__(ConfigurationAggregatorArgs)

        __props__.__dict__["account_aggregation_sources"] = None
        __props__.__dict__["configuration_aggregator_arn"] = None
        __props__.__dict__["configuration_aggregator_name"] = None
        __props__.__dict__["organization_aggregation_source"] = None
        __props__.__dict__["tags"] = None
        return ConfigurationAggregator(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountAggregationSources")
    def account_aggregation_sources(self) -> pulumi.Output[Optional[Sequence['outputs.ConfigurationAggregatorAccountAggregationSource']]]:
        """
        Provides a list of source accounts and regions to be aggregated.
        """
        return pulumi.get(self, "account_aggregation_sources")

    @property
    @pulumi.getter(name="configurationAggregatorArn")
    def configuration_aggregator_arn(self) -> pulumi.Output[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the aggregator.
        """
        return pulumi.get(self, "configuration_aggregator_arn")

    @property
    @pulumi.getter(name="configurationAggregatorName")
    def configuration_aggregator_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The name of the aggregator.
        """
        return pulumi.get(self, "configuration_aggregator_name")

    @property
    @pulumi.getter(name="organizationAggregationSource")
    def organization_aggregation_source(self) -> pulumi.Output[Optional['outputs.ConfigurationAggregatorOrganizationAggregationSource']]:
        """
        Provides an organization and list of regions to be aggregated.
        """
        return pulumi.get(self, "organization_aggregation_source")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        The tags for the configuration aggregator.
        """
        return pulumi.get(self, "tags")

