# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SecurityConfigurationArgs', 'SecurityConfiguration']

@pulumi.input_type
class SecurityConfigurationArgs:
    def __init__(__self__, *,
                 security_configuration: Any,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SecurityConfiguration resource.
        :param Any security_configuration: The security configuration details in JSON format.
        :param pulumi.Input[str] name: The name of the security configuration.
        """
        pulumi.set(__self__, "security_configuration", security_configuration)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="securityConfiguration")
    def security_configuration(self) -> Any:
        """
        The security configuration details in JSON format.
        """
        return pulumi.get(self, "security_configuration")

    @security_configuration.setter
    def security_configuration(self, value: Any):
        pulumi.set(self, "security_configuration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the security configuration.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class SecurityConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 security_configuration: Optional[Any] = None,
                 __props__=None):
        """
        Use a SecurityConfiguration resource to configure data encryption, Kerberos authentication, and Amazon S3 authorization for EMRFS.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the security configuration.
        :param Any security_configuration: The security configuration details in JSON format.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecurityConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use a SecurityConfiguration resource to configure data encryption, Kerberos authentication, and Amazon S3 authorization for EMRFS.

        :param str resource_name: The name of the resource.
        :param SecurityConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecurityConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 security_configuration: Optional[Any] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecurityConfigurationArgs.__new__(SecurityConfigurationArgs)

            __props__.__dict__["name"] = name
            if security_configuration is None and not opts.urn:
                raise TypeError("Missing required property 'security_configuration'")
            __props__.__dict__["security_configuration"] = security_configuration
        super(SecurityConfiguration, __self__).__init__(
            'aws-native:emr:SecurityConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SecurityConfiguration':
        """
        Get an existing SecurityConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SecurityConfigurationArgs.__new__(SecurityConfigurationArgs)

        __props__.__dict__["name"] = None
        __props__.__dict__["security_configuration"] = None
        return SecurityConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the security configuration.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="securityConfiguration")
    def security_configuration(self) -> pulumi.Output[Any]:
        """
        The security configuration details in JSON format.
        """
        return pulumi.get(self, "security_configuration")
