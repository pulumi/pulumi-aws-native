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

__all__ = ['LoggingConfigurationInitArgs', 'LoggingConfiguration']

@pulumi.input_type
class LoggingConfigurationInitArgs:
    def __init__(__self__, *,
                 firewall_arn: pulumi.Input[builtins.str],
                 logging_configuration: pulumi.Input['LoggingConfigurationArgs'],
                 enable_monitoring_dashboard: Optional[pulumi.Input[builtins.bool]] = None,
                 firewall_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a LoggingConfiguration resource.
        :param pulumi.Input[builtins.str] firewall_arn: The Amazon Resource Name (ARN) of the `Firewall` that the logging configuration is associated with. You can't change the firewall specification after you create the logging configuration.
        :param pulumi.Input['LoggingConfigurationArgs'] logging_configuration: Defines how AWS Network Firewall performs logging for a `Firewall` .
        :param pulumi.Input[builtins.str] firewall_name: The name of the firewall that the logging configuration is associated with. You can't change the firewall specification after you create the logging configuration.
        """
        pulumi.set(__self__, "firewall_arn", firewall_arn)
        pulumi.set(__self__, "logging_configuration", logging_configuration)
        if enable_monitoring_dashboard is not None:
            pulumi.set(__self__, "enable_monitoring_dashboard", enable_monitoring_dashboard)
        if firewall_name is not None:
            pulumi.set(__self__, "firewall_name", firewall_name)

    @property
    @pulumi.getter(name="firewallArn")
    def firewall_arn(self) -> pulumi.Input[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the `Firewall` that the logging configuration is associated with. You can't change the firewall specification after you create the logging configuration.
        """
        return pulumi.get(self, "firewall_arn")

    @firewall_arn.setter
    def firewall_arn(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "firewall_arn", value)

    @property
    @pulumi.getter(name="loggingConfiguration")
    def logging_configuration(self) -> pulumi.Input['LoggingConfigurationArgs']:
        """
        Defines how AWS Network Firewall performs logging for a `Firewall` .
        """
        return pulumi.get(self, "logging_configuration")

    @logging_configuration.setter
    def logging_configuration(self, value: pulumi.Input['LoggingConfigurationArgs']):
        pulumi.set(self, "logging_configuration", value)

    @property
    @pulumi.getter(name="enableMonitoringDashboard")
    def enable_monitoring_dashboard(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "enable_monitoring_dashboard")

    @enable_monitoring_dashboard.setter
    def enable_monitoring_dashboard(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enable_monitoring_dashboard", value)

    @property
    @pulumi.getter(name="firewallName")
    def firewall_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the firewall that the logging configuration is associated with. You can't change the firewall specification after you create the logging configuration.
        """
        return pulumi.get(self, "firewall_name")

    @firewall_name.setter
    def firewall_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "firewall_name", value)


@pulumi.type_token("aws-native:networkfirewall:LoggingConfiguration")
class LoggingConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable_monitoring_dashboard: Optional[pulumi.Input[builtins.bool]] = None,
                 firewall_arn: Optional[pulumi.Input[builtins.str]] = None,
                 firewall_name: Optional[pulumi.Input[builtins.str]] = None,
                 logging_configuration: Optional[pulumi.Input[Union['LoggingConfigurationArgs', 'LoggingConfigurationArgsDict']]] = None,
                 __props__=None):
        """
        Resource type definition for AWS::NetworkFirewall::LoggingConfiguration

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] firewall_arn: The Amazon Resource Name (ARN) of the `Firewall` that the logging configuration is associated with. You can't change the firewall specification after you create the logging configuration.
        :param pulumi.Input[builtins.str] firewall_name: The name of the firewall that the logging configuration is associated with. You can't change the firewall specification after you create the logging configuration.
        :param pulumi.Input[Union['LoggingConfigurationArgs', 'LoggingConfigurationArgsDict']] logging_configuration: Defines how AWS Network Firewall performs logging for a `Firewall` .
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LoggingConfigurationInitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource type definition for AWS::NetworkFirewall::LoggingConfiguration

        :param str resource_name: The name of the resource.
        :param LoggingConfigurationInitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LoggingConfigurationInitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable_monitoring_dashboard: Optional[pulumi.Input[builtins.bool]] = None,
                 firewall_arn: Optional[pulumi.Input[builtins.str]] = None,
                 firewall_name: Optional[pulumi.Input[builtins.str]] = None,
                 logging_configuration: Optional[pulumi.Input[Union['LoggingConfigurationArgs', 'LoggingConfigurationArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LoggingConfigurationInitArgs.__new__(LoggingConfigurationInitArgs)

            __props__.__dict__["enable_monitoring_dashboard"] = enable_monitoring_dashboard
            if firewall_arn is None and not opts.urn:
                raise TypeError("Missing required property 'firewall_arn'")
            __props__.__dict__["firewall_arn"] = firewall_arn
            __props__.__dict__["firewall_name"] = firewall_name
            if logging_configuration is None and not opts.urn:
                raise TypeError("Missing required property 'logging_configuration'")
            __props__.__dict__["logging_configuration"] = logging_configuration
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["firewallArn", "firewallName"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(LoggingConfiguration, __self__).__init__(
            'aws-native:networkfirewall:LoggingConfiguration',
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

        __props__ = LoggingConfigurationInitArgs.__new__(LoggingConfigurationInitArgs)

        __props__.__dict__["enable_monitoring_dashboard"] = None
        __props__.__dict__["firewall_arn"] = None
        __props__.__dict__["firewall_name"] = None
        __props__.__dict__["logging_configuration"] = None
        return LoggingConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="enableMonitoringDashboard")
    def enable_monitoring_dashboard(self) -> pulumi.Output[Optional[builtins.bool]]:
        return pulumi.get(self, "enable_monitoring_dashboard")

    @property
    @pulumi.getter(name="firewallArn")
    def firewall_arn(self) -> pulumi.Output[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the `Firewall` that the logging configuration is associated with. You can't change the firewall specification after you create the logging configuration.
        """
        return pulumi.get(self, "firewall_arn")

    @property
    @pulumi.getter(name="firewallName")
    def firewall_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The name of the firewall that the logging configuration is associated with. You can't change the firewall specification after you create the logging configuration.
        """
        return pulumi.get(self, "firewall_name")

    @property
    @pulumi.getter(name="loggingConfiguration")
    def logging_configuration(self) -> pulumi.Output['outputs.LoggingConfiguration']:
        """
        Defines how AWS Network Firewall performs logging for a `Firewall` .
        """
        return pulumi.get(self, "logging_configuration")

