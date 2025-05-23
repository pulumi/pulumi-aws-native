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

__all__ = [
    'GetSecurityKeyResult',
    'AwaitableGetSecurityKeyResult',
    'get_security_key',
    'get_security_key_output',
]

@pulumi.output_type
class GetSecurityKeyResult:
    def __init__(__self__, association_id=None):
        if association_id and not isinstance(association_id, str):
            raise TypeError("Expected argument 'association_id' to be a str")
        pulumi.set(__self__, "association_id", association_id)

    @property
    @pulumi.getter(name="associationId")
    def association_id(self) -> Optional[builtins.str]:
        """
        An `AssociationId` is automatically generated when a storage config is associated with an instance.
        """
        return pulumi.get(self, "association_id")


class AwaitableGetSecurityKeyResult(GetSecurityKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecurityKeyResult(
            association_id=self.association_id)


def get_security_key(association_id: Optional[builtins.str] = None,
                     instance_id: Optional[builtins.str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecurityKeyResult:
    """
    Resource Type definition for AWS::Connect::SecurityKey


    :param builtins.str association_id: An `AssociationId` is automatically generated when a storage config is associated with an instance.
    :param builtins.str instance_id: The Amazon Resource Name (ARN) of the instance.
           
           *Minimum* : `1`
           
           *Maximum* : `100`
    """
    __args__ = dict()
    __args__['associationId'] = association_id
    __args__['instanceId'] = instance_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:connect:getSecurityKey', __args__, opts=opts, typ=GetSecurityKeyResult).value

    return AwaitableGetSecurityKeyResult(
        association_id=pulumi.get(__ret__, 'association_id'))
def get_security_key_output(association_id: Optional[pulumi.Input[builtins.str]] = None,
                            instance_id: Optional[pulumi.Input[builtins.str]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetSecurityKeyResult]:
    """
    Resource Type definition for AWS::Connect::SecurityKey


    :param builtins.str association_id: An `AssociationId` is automatically generated when a storage config is associated with an instance.
    :param builtins.str instance_id: The Amazon Resource Name (ARN) of the instance.
           
           *Minimum* : `1`
           
           *Maximum* : `100`
    """
    __args__ = dict()
    __args__['associationId'] = association_id
    __args__['instanceId'] = instance_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:connect:getSecurityKey', __args__, opts=opts, typ=GetSecurityKeyResult)
    return __ret__.apply(lambda __response__: GetSecurityKeyResult(
        association_id=pulumi.get(__response__, 'association_id')))
