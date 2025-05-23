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

__all__ = [
    'GetEntitlementResult',
    'AwaitableGetEntitlementResult',
    'get_entitlement',
    'get_entitlement_output',
]

@pulumi.output_type
class GetEntitlementResult:
    def __init__(__self__, app_visibility=None, attributes=None, created_time=None, description=None, last_modified_time=None):
        if app_visibility and not isinstance(app_visibility, str):
            raise TypeError("Expected argument 'app_visibility' to be a str")
        pulumi.set(__self__, "app_visibility", app_visibility)
        if attributes and not isinstance(attributes, list):
            raise TypeError("Expected argument 'attributes' to be a list")
        pulumi.set(__self__, "attributes", attributes)
        if created_time and not isinstance(created_time, str):
            raise TypeError("Expected argument 'created_time' to be a str")
        pulumi.set(__self__, "created_time", created_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if last_modified_time and not isinstance(last_modified_time, str):
            raise TypeError("Expected argument 'last_modified_time' to be a str")
        pulumi.set(__self__, "last_modified_time", last_modified_time)

    @property
    @pulumi.getter(name="appVisibility")
    def app_visibility(self) -> Optional[builtins.str]:
        """
        Specifies whether to entitle all apps or only selected apps.
        """
        return pulumi.get(self, "app_visibility")

    @property
    @pulumi.getter
    def attributes(self) -> Optional[Sequence['outputs.EntitlementAttribute']]:
        """
        The attributes of the entitlement.
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> Optional[builtins.str]:
        """
        The time when the entitlement was created.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        The description of the entitlement.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> Optional[builtins.str]:
        """
        The time when the entitlement was last modified.
        """
        return pulumi.get(self, "last_modified_time")


class AwaitableGetEntitlementResult(GetEntitlementResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEntitlementResult(
            app_visibility=self.app_visibility,
            attributes=self.attributes,
            created_time=self.created_time,
            description=self.description,
            last_modified_time=self.last_modified_time)


def get_entitlement(name: Optional[builtins.str] = None,
                    stack_name: Optional[builtins.str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEntitlementResult:
    """
    Resource Type definition for AWS::AppStream::Entitlement


    :param builtins.str name: The name of the entitlement.
    :param builtins.str stack_name: The name of the stack.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['stackName'] = stack_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:appstream:getEntitlement', __args__, opts=opts, typ=GetEntitlementResult).value

    return AwaitableGetEntitlementResult(
        app_visibility=pulumi.get(__ret__, 'app_visibility'),
        attributes=pulumi.get(__ret__, 'attributes'),
        created_time=pulumi.get(__ret__, 'created_time'),
        description=pulumi.get(__ret__, 'description'),
        last_modified_time=pulumi.get(__ret__, 'last_modified_time'))
def get_entitlement_output(name: Optional[pulumi.Input[builtins.str]] = None,
                           stack_name: Optional[pulumi.Input[builtins.str]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetEntitlementResult]:
    """
    Resource Type definition for AWS::AppStream::Entitlement


    :param builtins.str name: The name of the entitlement.
    :param builtins.str stack_name: The name of the stack.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['stackName'] = stack_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:appstream:getEntitlement', __args__, opts=opts, typ=GetEntitlementResult)
    return __ret__.apply(lambda __response__: GetEntitlementResult(
        app_visibility=pulumi.get(__response__, 'app_visibility'),
        attributes=pulumi.get(__response__, 'attributes'),
        created_time=pulumi.get(__response__, 'created_time'),
        description=pulumi.get(__response__, 'description'),
        last_modified_time=pulumi.get(__response__, 'last_modified_time')))
