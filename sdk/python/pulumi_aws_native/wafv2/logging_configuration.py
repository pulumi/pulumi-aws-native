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
from ._enums import *
from ._inputs import *

__all__ = ['LoggingConfigurationArgs', 'LoggingConfiguration']

@pulumi.input_type
class LoggingConfigurationArgs:
    def __init__(__self__, *,
                 log_destination_configs: pulumi.Input[Sequence[pulumi.Input[builtins.str]]],
                 resource_arn: pulumi.Input[builtins.str],
                 logging_filter: Optional[pulumi.Input['LoggingFilterPropertiesArgs']] = None,
                 redacted_fields: Optional[pulumi.Input[Sequence[pulumi.Input['LoggingConfigurationFieldToMatchArgs']]]] = None):
        """
        The set of arguments for constructing a LoggingConfiguration resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] log_destination_configs: The Amazon Resource Names (ARNs) of the logging destinations that you want to associate with the web ACL.
        :param pulumi.Input[builtins.str] resource_arn: The Amazon Resource Name (ARN) of the web ACL that you want to associate with LogDestinationConfigs.
        :param pulumi.Input['LoggingFilterPropertiesArgs'] logging_filter: Filtering that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.
        :param pulumi.Input[Sequence[pulumi.Input['LoggingConfigurationFieldToMatchArgs']]] redacted_fields: The parts of the request that you want to keep out of the logs. For example, if you redact the HEADER field, the HEADER field in the firehose will be xxx.
        """
        pulumi.set(__self__, "log_destination_configs", log_destination_configs)
        pulumi.set(__self__, "resource_arn", resource_arn)
        if logging_filter is not None:
            pulumi.set(__self__, "logging_filter", logging_filter)
        if redacted_fields is not None:
            pulumi.set(__self__, "redacted_fields", redacted_fields)

    @property
    @pulumi.getter(name="logDestinationConfigs")
    def log_destination_configs(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        """
        The Amazon Resource Names (ARNs) of the logging destinations that you want to associate with the web ACL.
        """
        return pulumi.get(self, "log_destination_configs")

    @log_destination_configs.setter
    def log_destination_configs(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "log_destination_configs", value)

    @property
    @pulumi.getter(name="resourceArn")
    def resource_arn(self) -> pulumi.Input[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the web ACL that you want to associate with LogDestinationConfigs.
        """
        return pulumi.get(self, "resource_arn")

    @resource_arn.setter
    def resource_arn(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_arn", value)

    @property
    @pulumi.getter(name="loggingFilter")
    def logging_filter(self) -> Optional[pulumi.Input['LoggingFilterPropertiesArgs']]:
        """
        Filtering that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.
        """
        return pulumi.get(self, "logging_filter")

    @logging_filter.setter
    def logging_filter(self, value: Optional[pulumi.Input['LoggingFilterPropertiesArgs']]):
        pulumi.set(self, "logging_filter", value)

    @property
    @pulumi.getter(name="redactedFields")
    def redacted_fields(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LoggingConfigurationFieldToMatchArgs']]]]:
        """
        The parts of the request that you want to keep out of the logs. For example, if you redact the HEADER field, the HEADER field in the firehose will be xxx.
        """
        return pulumi.get(self, "redacted_fields")

    @redacted_fields.setter
    def redacted_fields(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LoggingConfigurationFieldToMatchArgs']]]]):
        pulumi.set(self, "redacted_fields", value)


@pulumi.type_token("aws-native:wafv2:LoggingConfiguration")
class LoggingConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 log_destination_configs: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 logging_filter: Optional[pulumi.Input[Union['LoggingFilterPropertiesArgs', 'LoggingFilterPropertiesArgsDict']]] = None,
                 redacted_fields: Optional[pulumi.Input[Sequence[pulumi.Input[Union['LoggingConfigurationFieldToMatchArgs', 'LoggingConfigurationFieldToMatchArgsDict']]]]] = None,
                 resource_arn: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        A WAFv2 Logging Configuration Resource Provider

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] log_destination_configs: The Amazon Resource Names (ARNs) of the logging destinations that you want to associate with the web ACL.
        :param pulumi.Input[Union['LoggingFilterPropertiesArgs', 'LoggingFilterPropertiesArgsDict']] logging_filter: Filtering that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.
        :param pulumi.Input[Sequence[pulumi.Input[Union['LoggingConfigurationFieldToMatchArgs', 'LoggingConfigurationFieldToMatchArgsDict']]]] redacted_fields: The parts of the request that you want to keep out of the logs. For example, if you redact the HEADER field, the HEADER field in the firehose will be xxx.
        :param pulumi.Input[builtins.str] resource_arn: The Amazon Resource Name (ARN) of the web ACL that you want to associate with LogDestinationConfigs.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LoggingConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A WAFv2 Logging Configuration Resource Provider

        :param str resource_name: The name of the resource.
        :param LoggingConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LoggingConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 log_destination_configs: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 logging_filter: Optional[pulumi.Input[Union['LoggingFilterPropertiesArgs', 'LoggingFilterPropertiesArgsDict']]] = None,
                 redacted_fields: Optional[pulumi.Input[Sequence[pulumi.Input[Union['LoggingConfigurationFieldToMatchArgs', 'LoggingConfigurationFieldToMatchArgsDict']]]]] = None,
                 resource_arn: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LoggingConfigurationArgs.__new__(LoggingConfigurationArgs)

            if log_destination_configs is None and not opts.urn:
                raise TypeError("Missing required property 'log_destination_configs'")
            __props__.__dict__["log_destination_configs"] = log_destination_configs
            __props__.__dict__["logging_filter"] = logging_filter
            __props__.__dict__["redacted_fields"] = redacted_fields
            if resource_arn is None and not opts.urn:
                raise TypeError("Missing required property 'resource_arn'")
            __props__.__dict__["resource_arn"] = resource_arn
            __props__.__dict__["managed_by_firewall_manager"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["resourceArn"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(LoggingConfiguration, __self__).__init__(
            'aws-native:wafv2:LoggingConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'LoggingConfiguration':
        """
        Get an existing LoggingConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = LoggingConfigurationArgs.__new__(LoggingConfigurationArgs)

        __props__.__dict__["log_destination_configs"] = None
        __props__.__dict__["logging_filter"] = None
        __props__.__dict__["managed_by_firewall_manager"] = None
        __props__.__dict__["redacted_fields"] = None
        __props__.__dict__["resource_arn"] = None
        return LoggingConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="logDestinationConfigs")
    def log_destination_configs(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        The Amazon Resource Names (ARNs) of the logging destinations that you want to associate with the web ACL.
        """
        return pulumi.get(self, "log_destination_configs")

    @property
    @pulumi.getter(name="loggingFilter")
    def logging_filter(self) -> pulumi.Output[Optional['outputs.LoggingFilterProperties']]:
        """
        Filtering that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.
        """
        return pulumi.get(self, "logging_filter")

    @property
    @pulumi.getter(name="managedByFirewallManager")
    def managed_by_firewall_manager(self) -> pulumi.Output[builtins.bool]:
        """
        Indicates whether the logging configuration was created by AWS Firewall Manager, as part of an AWS WAF policy configuration. If true, only Firewall Manager can modify or delete the configuration.
        """
        return pulumi.get(self, "managed_by_firewall_manager")

    @property
    @pulumi.getter(name="redactedFields")
    def redacted_fields(self) -> pulumi.Output[Optional[Sequence['outputs.LoggingConfigurationFieldToMatch']]]:
        """
        The parts of the request that you want to keep out of the logs. For example, if you redact the HEADER field, the HEADER field in the firehose will be xxx.
        """
        return pulumi.get(self, "redacted_fields")

    @property
    @pulumi.getter(name="resourceArn")
    def resource_arn(self) -> pulumi.Output[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the web ACL that you want to associate with LogDestinationConfigs.
        """
        return pulumi.get(self, "resource_arn")

