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
    'GetGeofenceCollectionResult',
    'AwaitableGetGeofenceCollectionResult',
    'get_geofence_collection',
    'get_geofence_collection_output',
]

@pulumi.output_type
class GetGeofenceCollectionResult:
    def __init__(__self__, arn=None, collection_arn=None, create_time=None, pricing_plan=None, pricing_plan_data_source=None, update_time=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if collection_arn and not isinstance(collection_arn, str):
            raise TypeError("Expected argument 'collection_arn' to be a str")
        pulumi.set(__self__, "collection_arn", collection_arn)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if pricing_plan and not isinstance(pricing_plan, str):
            raise TypeError("Expected argument 'pricing_plan' to be a str")
        pulumi.set(__self__, "pricing_plan", pricing_plan)
        if pricing_plan_data_source and not isinstance(pricing_plan_data_source, str):
            raise TypeError("Expected argument 'pricing_plan_data_source' to be a str")
        pulumi.set(__self__, "pricing_plan_data_source", pricing_plan_data_source)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="collectionArn")
    def collection_arn(self) -> Optional[str]:
        return pulumi.get(self, "collection_arn")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[str]:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="pricingPlan")
    def pricing_plan(self) -> Optional['GeofenceCollectionPricingPlan']:
        return pulumi.get(self, "pricing_plan")

    @property
    @pulumi.getter(name="pricingPlanDataSource")
    def pricing_plan_data_source(self) -> Optional[str]:
        return pulumi.get(self, "pricing_plan_data_source")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[str]:
        return pulumi.get(self, "update_time")


class AwaitableGetGeofenceCollectionResult(GetGeofenceCollectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGeofenceCollectionResult(
            arn=self.arn,
            collection_arn=self.collection_arn,
            create_time=self.create_time,
            pricing_plan=self.pricing_plan,
            pricing_plan_data_source=self.pricing_plan_data_source,
            update_time=self.update_time)


def get_geofence_collection(collection_name: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGeofenceCollectionResult:
    """
    Definition of AWS::Location::GeofenceCollection Resource Type
    """
    __args__ = dict()
    __args__['collectionName'] = collection_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:location:getGeofenceCollection', __args__, opts=opts, typ=GetGeofenceCollectionResult).value

    return AwaitableGetGeofenceCollectionResult(
        arn=__ret__.arn,
        collection_arn=__ret__.collection_arn,
        create_time=__ret__.create_time,
        pricing_plan=__ret__.pricing_plan,
        pricing_plan_data_source=__ret__.pricing_plan_data_source,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_geofence_collection)
def get_geofence_collection_output(collection_name: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGeofenceCollectionResult]:
    """
    Definition of AWS::Location::GeofenceCollection Resource Type
    """
    ...