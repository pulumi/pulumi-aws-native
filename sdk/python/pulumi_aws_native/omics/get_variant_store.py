# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'GetVariantStoreResult',
    'AwaitableGetVariantStoreResult',
    'get_variant_store',
    'get_variant_store_output',
]

@pulumi.output_type
class GetVariantStoreResult:
    def __init__(__self__, creation_time=None, description=None, id=None, status=None, status_message=None, store_arn=None, store_size_bytes=None, update_time=None):
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if status_message and not isinstance(status_message, str):
            raise TypeError("Expected argument 'status_message' to be a str")
        pulumi.set(__self__, "status_message", status_message)
        if store_arn and not isinstance(store_arn, str):
            raise TypeError("Expected argument 'store_arn' to be a str")
        pulumi.set(__self__, "store_arn", store_arn)
        if store_size_bytes and not isinstance(store_size_bytes, float):
            raise TypeError("Expected argument 'store_size_bytes' to be a float")
        pulumi.set(__self__, "store_size_bytes", store_size_bytes)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[str]:
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def status(self) -> Optional['VariantStoreStoreStatus']:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusMessage")
    def status_message(self) -> Optional[str]:
        return pulumi.get(self, "status_message")

    @property
    @pulumi.getter(name="storeArn")
    def store_arn(self) -> Optional[str]:
        return pulumi.get(self, "store_arn")

    @property
    @pulumi.getter(name="storeSizeBytes")
    def store_size_bytes(self) -> Optional[float]:
        return pulumi.get(self, "store_size_bytes")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[str]:
        return pulumi.get(self, "update_time")


class AwaitableGetVariantStoreResult(GetVariantStoreResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVariantStoreResult(
            creation_time=self.creation_time,
            description=self.description,
            id=self.id,
            status=self.status,
            status_message=self.status_message,
            store_arn=self.store_arn,
            store_size_bytes=self.store_size_bytes,
            update_time=self.update_time)


def get_variant_store(name: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVariantStoreResult:
    """
    Definition of AWS::Omics::VariantStore Resource Type
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:omics:getVariantStore', __args__, opts=opts, typ=GetVariantStoreResult).value

    return AwaitableGetVariantStoreResult(
        creation_time=__ret__.creation_time,
        description=__ret__.description,
        id=__ret__.id,
        status=__ret__.status,
        status_message=__ret__.status_message,
        store_arn=__ret__.store_arn,
        store_size_bytes=__ret__.store_size_bytes,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_variant_store)
def get_variant_store_output(name: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVariantStoreResult]:
    """
    Definition of AWS::Omics::VariantStore Resource Type
    """
    ...