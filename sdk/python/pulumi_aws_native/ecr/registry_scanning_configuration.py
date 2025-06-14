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

__all__ = ['RegistryScanningConfigurationArgs', 'RegistryScanningConfiguration']

@pulumi.input_type
class RegistryScanningConfigurationArgs:
    def __init__(__self__, *,
                 rules: pulumi.Input[Sequence[pulumi.Input['RegistryScanningConfigurationScanningRuleArgs']]],
                 scan_type: pulumi.Input['RegistryScanningConfigurationScanType']):
        """
        The set of arguments for constructing a RegistryScanningConfiguration resource.
        :param pulumi.Input[Sequence[pulumi.Input['RegistryScanningConfigurationScanningRuleArgs']]] rules: The scanning rules associated with the registry.
        :param pulumi.Input['RegistryScanningConfigurationScanType'] scan_type: The type of scanning configured for the registry.
        """
        pulumi.set(__self__, "rules", rules)
        pulumi.set(__self__, "scan_type", scan_type)

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Input[Sequence[pulumi.Input['RegistryScanningConfigurationScanningRuleArgs']]]:
        """
        The scanning rules associated with the registry.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: pulumi.Input[Sequence[pulumi.Input['RegistryScanningConfigurationScanningRuleArgs']]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter(name="scanType")
    def scan_type(self) -> pulumi.Input['RegistryScanningConfigurationScanType']:
        """
        The type of scanning configured for the registry.
        """
        return pulumi.get(self, "scan_type")

    @scan_type.setter
    def scan_type(self, value: pulumi.Input['RegistryScanningConfigurationScanType']):
        pulumi.set(self, "scan_type", value)


@pulumi.type_token("aws-native:ecr:RegistryScanningConfiguration")
class RegistryScanningConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RegistryScanningConfigurationScanningRuleArgs', 'RegistryScanningConfigurationScanningRuleArgsDict']]]]] = None,
                 scan_type: Optional[pulumi.Input['RegistryScanningConfigurationScanType']] = None,
                 __props__=None):
        """
        The scanning configuration for a private registry.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['RegistryScanningConfigurationScanningRuleArgs', 'RegistryScanningConfigurationScanningRuleArgsDict']]]] rules: The scanning rules associated with the registry.
        :param pulumi.Input['RegistryScanningConfigurationScanType'] scan_type: The type of scanning configured for the registry.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RegistryScanningConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The scanning configuration for a private registry.

        :param str resource_name: The name of the resource.
        :param RegistryScanningConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RegistryScanningConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RegistryScanningConfigurationScanningRuleArgs', 'RegistryScanningConfigurationScanningRuleArgsDict']]]]] = None,
                 scan_type: Optional[pulumi.Input['RegistryScanningConfigurationScanType']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RegistryScanningConfigurationArgs.__new__(RegistryScanningConfigurationArgs)

            if rules is None and not opts.urn:
                raise TypeError("Missing required property 'rules'")
            __props__.__dict__["rules"] = rules
            if scan_type is None and not opts.urn:
                raise TypeError("Missing required property 'scan_type'")
            __props__.__dict__["scan_type"] = scan_type
            __props__.__dict__["registry_id"] = None
        super(RegistryScanningConfiguration, __self__).__init__(
            'aws-native:ecr:RegistryScanningConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RegistryScanningConfiguration':
        """
        Get an existing RegistryScanningConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RegistryScanningConfigurationArgs.__new__(RegistryScanningConfigurationArgs)

        __props__.__dict__["registry_id"] = None
        __props__.__dict__["rules"] = None
        __props__.__dict__["scan_type"] = None
        return RegistryScanningConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> pulumi.Output[builtins.str]:
        """
        The account ID of the destination registry.
        """
        return pulumi.get(self, "registry_id")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Sequence['outputs.RegistryScanningConfigurationScanningRule']]:
        """
        The scanning rules associated with the registry.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter(name="scanType")
    def scan_type(self) -> pulumi.Output['RegistryScanningConfigurationScanType']:
        """
        The type of scanning configured for the registry.
        """
        return pulumi.get(self, "scan_type")

