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

__all__ = ['TlsInspectionConfigurationArgs', 'TlsInspectionConfiguration']

@pulumi.input_type
class TlsInspectionConfigurationArgs:
    def __init__(__self__, *,
                 tls_inspection_configuration: pulumi.Input['TlsInspectionConfigurationTlsInspectionConfigurationArgs'],
                 description: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['TlsInspectionConfigurationTagArgs']]]] = None,
                 tls_inspection_configuration_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TlsInspectionConfiguration resource.
        """
        pulumi.set(__self__, "tls_inspection_configuration", tls_inspection_configuration)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tls_inspection_configuration_name is not None:
            pulumi.set(__self__, "tls_inspection_configuration_name", tls_inspection_configuration_name)

    @property
    @pulumi.getter(name="tlsInspectionConfiguration")
    def tls_inspection_configuration(self) -> pulumi.Input['TlsInspectionConfigurationTlsInspectionConfigurationArgs']:
        return pulumi.get(self, "tls_inspection_configuration")

    @tls_inspection_configuration.setter
    def tls_inspection_configuration(self, value: pulumi.Input['TlsInspectionConfigurationTlsInspectionConfigurationArgs']):
        pulumi.set(self, "tls_inspection_configuration", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TlsInspectionConfigurationTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TlsInspectionConfigurationTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tlsInspectionConfigurationName")
    def tls_inspection_configuration_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tls_inspection_configuration_name")

    @tls_inspection_configuration_name.setter
    def tls_inspection_configuration_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_inspection_configuration_name", value)


class TlsInspectionConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TlsInspectionConfigurationTagArgs']]]]] = None,
                 tls_inspection_configuration: Optional[pulumi.Input[pulumi.InputType['TlsInspectionConfigurationTlsInspectionConfigurationArgs']]] = None,
                 tls_inspection_configuration_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource type definition for AWS::NetworkFirewall::TLSInspectionConfiguration

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TlsInspectionConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource type definition for AWS::NetworkFirewall::TLSInspectionConfiguration

        :param str resource_name: The name of the resource.
        :param TlsInspectionConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TlsInspectionConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TlsInspectionConfigurationTagArgs']]]]] = None,
                 tls_inspection_configuration: Optional[pulumi.Input[pulumi.InputType['TlsInspectionConfigurationTlsInspectionConfigurationArgs']]] = None,
                 tls_inspection_configuration_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TlsInspectionConfigurationArgs.__new__(TlsInspectionConfigurationArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["tags"] = tags
            if tls_inspection_configuration is None and not opts.urn:
                raise TypeError("Missing required property 'tls_inspection_configuration'")
            __props__.__dict__["tls_inspection_configuration"] = tls_inspection_configuration
            __props__.__dict__["tls_inspection_configuration_name"] = tls_inspection_configuration_name
            __props__.__dict__["tls_inspection_configuration_arn"] = None
            __props__.__dict__["tls_inspection_configuration_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["tls_inspection_configuration_name"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(TlsInspectionConfiguration, __self__).__init__(
            'aws-native:networkfirewall:TlsInspectionConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'TlsInspectionConfiguration':
        """
        Get an existing TlsInspectionConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TlsInspectionConfigurationArgs.__new__(TlsInspectionConfigurationArgs)

        __props__.__dict__["description"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["tls_inspection_configuration"] = None
        __props__.__dict__["tls_inspection_configuration_arn"] = None
        __props__.__dict__["tls_inspection_configuration_id"] = None
        __props__.__dict__["tls_inspection_configuration_name"] = None
        return TlsInspectionConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.TlsInspectionConfigurationTag']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tlsInspectionConfiguration")
    def tls_inspection_configuration(self) -> pulumi.Output['outputs.TlsInspectionConfigurationTlsInspectionConfiguration']:
        return pulumi.get(self, "tls_inspection_configuration")

    @property
    @pulumi.getter(name="tlsInspectionConfigurationArn")
    def tls_inspection_configuration_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "tls_inspection_configuration_arn")

    @property
    @pulumi.getter(name="tlsInspectionConfigurationId")
    def tls_inspection_configuration_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "tls_inspection_configuration_id")

    @property
    @pulumi.getter(name="tlsInspectionConfigurationName")
    def tls_inspection_configuration_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "tls_inspection_configuration_name")
