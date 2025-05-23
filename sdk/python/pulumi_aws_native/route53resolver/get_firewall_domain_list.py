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
from .. import outputs as _root_outputs
from ._enums import *

__all__ = [
    'GetFirewallDomainListResult',
    'AwaitableGetFirewallDomainListResult',
    'get_firewall_domain_list',
    'get_firewall_domain_list_output',
]

@pulumi.output_type
class GetFirewallDomainListResult:
    def __init__(__self__, arn=None, creation_time=None, creator_request_id=None, domain_count=None, id=None, managed_owner_name=None, modification_time=None, status=None, status_message=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if creator_request_id and not isinstance(creator_request_id, str):
            raise TypeError("Expected argument 'creator_request_id' to be a str")
        pulumi.set(__self__, "creator_request_id", creator_request_id)
        if domain_count and not isinstance(domain_count, int):
            raise TypeError("Expected argument 'domain_count' to be a int")
        pulumi.set(__self__, "domain_count", domain_count)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if managed_owner_name and not isinstance(managed_owner_name, str):
            raise TypeError("Expected argument 'managed_owner_name' to be a str")
        pulumi.set(__self__, "managed_owner_name", managed_owner_name)
        if modification_time and not isinstance(modification_time, str):
            raise TypeError("Expected argument 'modification_time' to be a str")
        pulumi.set(__self__, "modification_time", modification_time)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if status_message and not isinstance(status_message, str):
            raise TypeError("Expected argument 'status_message' to be a str")
        pulumi.set(__self__, "status_message", status_message)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[builtins.str]:
        """
        Arn
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[builtins.str]:
        """
        Rfc3339TimeString
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="creatorRequestId")
    def creator_request_id(self) -> Optional[builtins.str]:
        """
        The id of the creator request.
        """
        return pulumi.get(self, "creator_request_id")

    @property
    @pulumi.getter(name="domainCount")
    def domain_count(self) -> Optional[builtins.int]:
        """
        Count
        """
        return pulumi.get(self, "domain_count")

    @property
    @pulumi.getter
    def id(self) -> Optional[builtins.str]:
        """
        ResourceId
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="managedOwnerName")
    def managed_owner_name(self) -> Optional[builtins.str]:
        """
        ServicePrincipal
        """
        return pulumi.get(self, "managed_owner_name")

    @property
    @pulumi.getter(name="modificationTime")
    def modification_time(self) -> Optional[builtins.str]:
        """
        Rfc3339TimeString
        """
        return pulumi.get(self, "modification_time")

    @property
    @pulumi.getter
    def status(self) -> Optional['FirewallDomainListStatus']:
        """
        ResolverFirewallDomainList, possible values are COMPLETE, DELETING, UPDATING, COMPLETE_IMPORT_FAILED, IMPORTING, and INACTIVE_OWNER_ACCOUNT_CLOSED.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusMessage")
    def status_message(self) -> Optional[builtins.str]:
        """
        FirewallDomainListAssociationStatus
        """
        return pulumi.get(self, "status_message")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        Tags
        """
        return pulumi.get(self, "tags")


class AwaitableGetFirewallDomainListResult(GetFirewallDomainListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallDomainListResult(
            arn=self.arn,
            creation_time=self.creation_time,
            creator_request_id=self.creator_request_id,
            domain_count=self.domain_count,
            id=self.id,
            managed_owner_name=self.managed_owner_name,
            modification_time=self.modification_time,
            status=self.status,
            status_message=self.status_message,
            tags=self.tags)


def get_firewall_domain_list(id: Optional[builtins.str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallDomainListResult:
    """
    Resource schema for AWS::Route53Resolver::FirewallDomainList.


    :param builtins.str id: ResourceId
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:route53resolver:getFirewallDomainList', __args__, opts=opts, typ=GetFirewallDomainListResult).value

    return AwaitableGetFirewallDomainListResult(
        arn=pulumi.get(__ret__, 'arn'),
        creation_time=pulumi.get(__ret__, 'creation_time'),
        creator_request_id=pulumi.get(__ret__, 'creator_request_id'),
        domain_count=pulumi.get(__ret__, 'domain_count'),
        id=pulumi.get(__ret__, 'id'),
        managed_owner_name=pulumi.get(__ret__, 'managed_owner_name'),
        modification_time=pulumi.get(__ret__, 'modification_time'),
        status=pulumi.get(__ret__, 'status'),
        status_message=pulumi.get(__ret__, 'status_message'),
        tags=pulumi.get(__ret__, 'tags'))
def get_firewall_domain_list_output(id: Optional[pulumi.Input[builtins.str]] = None,
                                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetFirewallDomainListResult]:
    """
    Resource schema for AWS::Route53Resolver::FirewallDomainList.


    :param builtins.str id: ResourceId
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:route53resolver:getFirewallDomainList', __args__, opts=opts, typ=GetFirewallDomainListResult)
    return __ret__.apply(lambda __response__: GetFirewallDomainListResult(
        arn=pulumi.get(__response__, 'arn'),
        creation_time=pulumi.get(__response__, 'creation_time'),
        creator_request_id=pulumi.get(__response__, 'creator_request_id'),
        domain_count=pulumi.get(__response__, 'domain_count'),
        id=pulumi.get(__response__, 'id'),
        managed_owner_name=pulumi.get(__response__, 'managed_owner_name'),
        modification_time=pulumi.get(__response__, 'modification_time'),
        status=pulumi.get(__response__, 'status'),
        status_message=pulumi.get(__response__, 'status_message'),
        tags=pulumi.get(__response__, 'tags')))
