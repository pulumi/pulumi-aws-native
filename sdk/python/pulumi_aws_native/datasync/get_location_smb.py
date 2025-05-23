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
from .. import outputs as _root_outputs
from ._enums import *

__all__ = [
    'GetLocationSmbResult',
    'AwaitableGetLocationSmbResult',
    'get_location_smb',
    'get_location_smb_output',
]

@pulumi.output_type
class GetLocationSmbResult:
    def __init__(__self__, agent_arns=None, authentication_type=None, dns_ip_addresses=None, domain=None, kerberos_principal=None, location_arn=None, location_uri=None, mount_options=None, tags=None, user=None):
        if agent_arns and not isinstance(agent_arns, list):
            raise TypeError("Expected argument 'agent_arns' to be a list")
        pulumi.set(__self__, "agent_arns", agent_arns)
        if authentication_type and not isinstance(authentication_type, str):
            raise TypeError("Expected argument 'authentication_type' to be a str")
        pulumi.set(__self__, "authentication_type", authentication_type)
        if dns_ip_addresses and not isinstance(dns_ip_addresses, list):
            raise TypeError("Expected argument 'dns_ip_addresses' to be a list")
        pulumi.set(__self__, "dns_ip_addresses", dns_ip_addresses)
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if kerberos_principal and not isinstance(kerberos_principal, str):
            raise TypeError("Expected argument 'kerberos_principal' to be a str")
        pulumi.set(__self__, "kerberos_principal", kerberos_principal)
        if location_arn and not isinstance(location_arn, str):
            raise TypeError("Expected argument 'location_arn' to be a str")
        pulumi.set(__self__, "location_arn", location_arn)
        if location_uri and not isinstance(location_uri, str):
            raise TypeError("Expected argument 'location_uri' to be a str")
        pulumi.set(__self__, "location_uri", location_uri)
        if mount_options and not isinstance(mount_options, dict):
            raise TypeError("Expected argument 'mount_options' to be a dict")
        pulumi.set(__self__, "mount_options", mount_options)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if user and not isinstance(user, str):
            raise TypeError("Expected argument 'user' to be a str")
        pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter(name="agentArns")
    def agent_arns(self) -> Optional[Sequence[builtins.str]]:
        """
        The Amazon Resource Names (ARNs) of agents to use for a Simple Message Block (SMB) location.
        """
        return pulumi.get(self, "agent_arns")

    @property
    @pulumi.getter(name="authenticationType")
    def authentication_type(self) -> Optional['LocationSmbAuthenticationType']:
        """
        The authentication mode used to determine identity of user.
        """
        return pulumi.get(self, "authentication_type")

    @property
    @pulumi.getter(name="dnsIpAddresses")
    def dns_ip_addresses(self) -> Optional[Sequence[builtins.str]]:
        """
        Specifies the IPv4 addresses for the DNS servers that your SMB file server belongs to. This parameter applies only if AuthenticationType is set to KERBEROS. If you have multiple domains in your environment, configuring this parameter makes sure that DataSync connects to the right SMB file server.
        """
        return pulumi.get(self, "dns_ip_addresses")

    @property
    @pulumi.getter
    def domain(self) -> Optional[builtins.str]:
        """
        The name of the Windows domain that the SMB server belongs to.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="kerberosPrincipal")
    def kerberos_principal(self) -> Optional[builtins.str]:
        """
        Specifies a service principal name (SPN), which is an identity in your Kerberos realm that has permission to access the files, folders, and file metadata in your SMB file server. SPNs are case sensitive and must include a prepended cifs/. For example, an SPN might look like cifs/kerberosuser@EXAMPLE.COM. Your task execution will fail if the SPN that you provide for this parameter doesn't match exactly what's in your keytab or krb5.conf files.
        """
        return pulumi.get(self, "kerberos_principal")

    @property
    @pulumi.getter(name="locationArn")
    def location_arn(self) -> Optional[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the SMB location that is created.
        """
        return pulumi.get(self, "location_arn")

    @property
    @pulumi.getter(name="locationUri")
    def location_uri(self) -> Optional[builtins.str]:
        """
        The URL of the SMB location that was described.
        """
        return pulumi.get(self, "location_uri")

    @property
    @pulumi.getter(name="mountOptions")
    def mount_options(self) -> Optional['outputs.LocationSmbMountOptions']:
        """
        Specifies the version of the SMB protocol that DataSync uses to access your SMB file server.
        """
        return pulumi.get(self, "mount_options")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def user(self) -> Optional[builtins.str]:
        """
        The user who can mount the share, has the permissions to access files and folders in the SMB share.
        """
        return pulumi.get(self, "user")


class AwaitableGetLocationSmbResult(GetLocationSmbResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLocationSmbResult(
            agent_arns=self.agent_arns,
            authentication_type=self.authentication_type,
            dns_ip_addresses=self.dns_ip_addresses,
            domain=self.domain,
            kerberos_principal=self.kerberos_principal,
            location_arn=self.location_arn,
            location_uri=self.location_uri,
            mount_options=self.mount_options,
            tags=self.tags,
            user=self.user)


def get_location_smb(location_arn: Optional[builtins.str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLocationSmbResult:
    """
    Resource schema for AWS::DataSync::LocationSMB.


    :param builtins.str location_arn: The Amazon Resource Name (ARN) of the SMB location that is created.
    """
    __args__ = dict()
    __args__['locationArn'] = location_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:datasync:getLocationSmb', __args__, opts=opts, typ=GetLocationSmbResult).value

    return AwaitableGetLocationSmbResult(
        agent_arns=pulumi.get(__ret__, 'agent_arns'),
        authentication_type=pulumi.get(__ret__, 'authentication_type'),
        dns_ip_addresses=pulumi.get(__ret__, 'dns_ip_addresses'),
        domain=pulumi.get(__ret__, 'domain'),
        kerberos_principal=pulumi.get(__ret__, 'kerberos_principal'),
        location_arn=pulumi.get(__ret__, 'location_arn'),
        location_uri=pulumi.get(__ret__, 'location_uri'),
        mount_options=pulumi.get(__ret__, 'mount_options'),
        tags=pulumi.get(__ret__, 'tags'),
        user=pulumi.get(__ret__, 'user'))
def get_location_smb_output(location_arn: Optional[pulumi.Input[builtins.str]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetLocationSmbResult]:
    """
    Resource schema for AWS::DataSync::LocationSMB.


    :param builtins.str location_arn: The Amazon Resource Name (ARN) of the SMB location that is created.
    """
    __args__ = dict()
    __args__['locationArn'] = location_arn
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:datasync:getLocationSmb', __args__, opts=opts, typ=GetLocationSmbResult)
    return __ret__.apply(lambda __response__: GetLocationSmbResult(
        agent_arns=pulumi.get(__response__, 'agent_arns'),
        authentication_type=pulumi.get(__response__, 'authentication_type'),
        dns_ip_addresses=pulumi.get(__response__, 'dns_ip_addresses'),
        domain=pulumi.get(__response__, 'domain'),
        kerberos_principal=pulumi.get(__response__, 'kerberos_principal'),
        location_arn=pulumi.get(__response__, 'location_arn'),
        location_uri=pulumi.get(__response__, 'location_uri'),
        mount_options=pulumi.get(__response__, 'mount_options'),
        tags=pulumi.get(__response__, 'tags'),
        user=pulumi.get(__response__, 'user')))
