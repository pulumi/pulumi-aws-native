# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetPrincipalPermissionsResult',
    'AwaitableGetPrincipalPermissionsResult',
    'get_principal_permissions',
    'get_principal_permissions_output',
]

@pulumi.output_type
class GetPrincipalPermissionsResult:
    def __init__(__self__, principal_identifier=None, resource_identifier=None):
        if principal_identifier and not isinstance(principal_identifier, str):
            raise TypeError("Expected argument 'principal_identifier' to be a str")
        pulumi.set(__self__, "principal_identifier", principal_identifier)
        if resource_identifier and not isinstance(resource_identifier, str):
            raise TypeError("Expected argument 'resource_identifier' to be a str")
        pulumi.set(__self__, "resource_identifier", resource_identifier)

    @property
    @pulumi.getter(name="principalIdentifier")
    def principal_identifier(self) -> Optional[str]:
        return pulumi.get(self, "principal_identifier")

    @property
    @pulumi.getter(name="resourceIdentifier")
    def resource_identifier(self) -> Optional[str]:
        return pulumi.get(self, "resource_identifier")


class AwaitableGetPrincipalPermissionsResult(GetPrincipalPermissionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPrincipalPermissionsResult(
            principal_identifier=self.principal_identifier,
            resource_identifier=self.resource_identifier)


def get_principal_permissions(principal_identifier: Optional[str] = None,
                              resource_identifier: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPrincipalPermissionsResult:
    """
    A resource schema representing a Lake Formation Permission.
    """
    __args__ = dict()
    __args__['principalIdentifier'] = principal_identifier
    __args__['resourceIdentifier'] = resource_identifier
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:lakeformation:getPrincipalPermissions', __args__, opts=opts, typ=GetPrincipalPermissionsResult).value

    return AwaitableGetPrincipalPermissionsResult(
        principal_identifier=__ret__.principal_identifier,
        resource_identifier=__ret__.resource_identifier)


@_utilities.lift_output_func(get_principal_permissions)
def get_principal_permissions_output(principal_identifier: Optional[pulumi.Input[str]] = None,
                                     resource_identifier: Optional[pulumi.Input[str]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPrincipalPermissionsResult]:
    """
    A resource schema representing a Lake Formation Permission.
    """
    ...