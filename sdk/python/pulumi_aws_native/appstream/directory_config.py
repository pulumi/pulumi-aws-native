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
from ._inputs import *

__all__ = ['DirectoryConfigArgs', 'DirectoryConfig']

@pulumi.input_type
class DirectoryConfigArgs:
    def __init__(__self__, *,
                 directory_name: pulumi.Input[builtins.str],
                 organizational_unit_distinguished_names: pulumi.Input[Sequence[pulumi.Input[builtins.str]]],
                 service_account_credentials: pulumi.Input['DirectoryConfigServiceAccountCredentialsArgs'],
                 certificate_based_auth_properties: Optional[pulumi.Input['DirectoryConfigCertificateBasedAuthPropertiesArgs']] = None):
        """
        The set of arguments for constructing a DirectoryConfig resource.
        :param pulumi.Input[builtins.str] directory_name: The fully qualified name of the directory (for example, corp.example.com).
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] organizational_unit_distinguished_names: The distinguished names of the organizational units for computer accounts.
        :param pulumi.Input['DirectoryConfigServiceAccountCredentialsArgs'] service_account_credentials: The credentials for the service account used by the streaming instance to connect to the directory. Do not use this parameter directly. Use `ServiceAccountCredentials` as an input parameter with `noEcho` as shown in the [Parameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/parameters-section-structure.html) . For best practices information, see [Do Not Embed Credentials in Your Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/best-practices.html#creds) .
        :param pulumi.Input['DirectoryConfigCertificateBasedAuthPropertiesArgs'] certificate_based_auth_properties: The certificate-based authentication properties used to authenticate SAML 2.0 Identity Provider (IdP) user identities to Active Directory domain-joined streaming instances.
        """
        pulumi.set(__self__, "directory_name", directory_name)
        pulumi.set(__self__, "organizational_unit_distinguished_names", organizational_unit_distinguished_names)
        pulumi.set(__self__, "service_account_credentials", service_account_credentials)
        if certificate_based_auth_properties is not None:
            pulumi.set(__self__, "certificate_based_auth_properties", certificate_based_auth_properties)

    @property
    @pulumi.getter(name="directoryName")
    def directory_name(self) -> pulumi.Input[builtins.str]:
        """
        The fully qualified name of the directory (for example, corp.example.com).
        """
        return pulumi.get(self, "directory_name")

    @directory_name.setter
    def directory_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "directory_name", value)

    @property
    @pulumi.getter(name="organizationalUnitDistinguishedNames")
    def organizational_unit_distinguished_names(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        """
        The distinguished names of the organizational units for computer accounts.
        """
        return pulumi.get(self, "organizational_unit_distinguished_names")

    @organizational_unit_distinguished_names.setter
    def organizational_unit_distinguished_names(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "organizational_unit_distinguished_names", value)

    @property
    @pulumi.getter(name="serviceAccountCredentials")
    def service_account_credentials(self) -> pulumi.Input['DirectoryConfigServiceAccountCredentialsArgs']:
        """
        The credentials for the service account used by the streaming instance to connect to the directory. Do not use this parameter directly. Use `ServiceAccountCredentials` as an input parameter with `noEcho` as shown in the [Parameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/parameters-section-structure.html) . For best practices information, see [Do Not Embed Credentials in Your Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/best-practices.html#creds) .
        """
        return pulumi.get(self, "service_account_credentials")

    @service_account_credentials.setter
    def service_account_credentials(self, value: pulumi.Input['DirectoryConfigServiceAccountCredentialsArgs']):
        pulumi.set(self, "service_account_credentials", value)

    @property
    @pulumi.getter(name="certificateBasedAuthProperties")
    def certificate_based_auth_properties(self) -> Optional[pulumi.Input['DirectoryConfigCertificateBasedAuthPropertiesArgs']]:
        """
        The certificate-based authentication properties used to authenticate SAML 2.0 Identity Provider (IdP) user identities to Active Directory domain-joined streaming instances.
        """
        return pulumi.get(self, "certificate_based_auth_properties")

    @certificate_based_auth_properties.setter
    def certificate_based_auth_properties(self, value: Optional[pulumi.Input['DirectoryConfigCertificateBasedAuthPropertiesArgs']]):
        pulumi.set(self, "certificate_based_auth_properties", value)


@pulumi.type_token("aws-native:appstream:DirectoryConfig")
class DirectoryConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_based_auth_properties: Optional[pulumi.Input[Union['DirectoryConfigCertificateBasedAuthPropertiesArgs', 'DirectoryConfigCertificateBasedAuthPropertiesArgsDict']]] = None,
                 directory_name: Optional[pulumi.Input[builtins.str]] = None,
                 organizational_unit_distinguished_names: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 service_account_credentials: Optional[pulumi.Input[Union['DirectoryConfigServiceAccountCredentialsArgs', 'DirectoryConfigServiceAccountCredentialsArgsDict']]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::AppStream::DirectoryConfig

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['DirectoryConfigCertificateBasedAuthPropertiesArgs', 'DirectoryConfigCertificateBasedAuthPropertiesArgsDict']] certificate_based_auth_properties: The certificate-based authentication properties used to authenticate SAML 2.0 Identity Provider (IdP) user identities to Active Directory domain-joined streaming instances.
        :param pulumi.Input[builtins.str] directory_name: The fully qualified name of the directory (for example, corp.example.com).
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] organizational_unit_distinguished_names: The distinguished names of the organizational units for computer accounts.
        :param pulumi.Input[Union['DirectoryConfigServiceAccountCredentialsArgs', 'DirectoryConfigServiceAccountCredentialsArgsDict']] service_account_credentials: The credentials for the service account used by the streaming instance to connect to the directory. Do not use this parameter directly. Use `ServiceAccountCredentials` as an input parameter with `noEcho` as shown in the [Parameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/parameters-section-structure.html) . For best practices information, see [Do Not Embed Credentials in Your Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/best-practices.html#creds) .
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DirectoryConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::AppStream::DirectoryConfig

        :param str resource_name: The name of the resource.
        :param DirectoryConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DirectoryConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_based_auth_properties: Optional[pulumi.Input[Union['DirectoryConfigCertificateBasedAuthPropertiesArgs', 'DirectoryConfigCertificateBasedAuthPropertiesArgsDict']]] = None,
                 directory_name: Optional[pulumi.Input[builtins.str]] = None,
                 organizational_unit_distinguished_names: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 service_account_credentials: Optional[pulumi.Input[Union['DirectoryConfigServiceAccountCredentialsArgs', 'DirectoryConfigServiceAccountCredentialsArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DirectoryConfigArgs.__new__(DirectoryConfigArgs)

            __props__.__dict__["certificate_based_auth_properties"] = certificate_based_auth_properties
            if directory_name is None and not opts.urn:
                raise TypeError("Missing required property 'directory_name'")
            __props__.__dict__["directory_name"] = directory_name
            if organizational_unit_distinguished_names is None and not opts.urn:
                raise TypeError("Missing required property 'organizational_unit_distinguished_names'")
            __props__.__dict__["organizational_unit_distinguished_names"] = organizational_unit_distinguished_names
            if service_account_credentials is None and not opts.urn:
                raise TypeError("Missing required property 'service_account_credentials'")
            __props__.__dict__["service_account_credentials"] = service_account_credentials
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["directoryName"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(DirectoryConfig, __self__).__init__(
            'aws-native:appstream:DirectoryConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DirectoryConfig':
        """
        Get an existing DirectoryConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DirectoryConfigArgs.__new__(DirectoryConfigArgs)

        __props__.__dict__["certificate_based_auth_properties"] = None
        __props__.__dict__["directory_name"] = None
        __props__.__dict__["organizational_unit_distinguished_names"] = None
        __props__.__dict__["service_account_credentials"] = None
        return DirectoryConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="certificateBasedAuthProperties")
    def certificate_based_auth_properties(self) -> pulumi.Output[Optional['outputs.DirectoryConfigCertificateBasedAuthProperties']]:
        """
        The certificate-based authentication properties used to authenticate SAML 2.0 Identity Provider (IdP) user identities to Active Directory domain-joined streaming instances.
        """
        return pulumi.get(self, "certificate_based_auth_properties")

    @property
    @pulumi.getter(name="directoryName")
    def directory_name(self) -> pulumi.Output[builtins.str]:
        """
        The fully qualified name of the directory (for example, corp.example.com).
        """
        return pulumi.get(self, "directory_name")

    @property
    @pulumi.getter(name="organizationalUnitDistinguishedNames")
    def organizational_unit_distinguished_names(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        The distinguished names of the organizational units for computer accounts.
        """
        return pulumi.get(self, "organizational_unit_distinguished_names")

    @property
    @pulumi.getter(name="serviceAccountCredentials")
    def service_account_credentials(self) -> pulumi.Output['outputs.DirectoryConfigServiceAccountCredentials']:
        """
        The credentials for the service account used by the streaming instance to connect to the directory. Do not use this parameter directly. Use `ServiceAccountCredentials` as an input parameter with `noEcho` as shown in the [Parameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/parameters-section-structure.html) . For best practices information, see [Do Not Embed Credentials in Your Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/best-practices.html#creds) .
        """
        return pulumi.get(self, "service_account_credentials")

