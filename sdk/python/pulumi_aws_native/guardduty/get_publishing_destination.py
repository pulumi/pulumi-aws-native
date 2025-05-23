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

__all__ = [
    'GetPublishingDestinationResult',
    'AwaitableGetPublishingDestinationResult',
    'get_publishing_destination',
    'get_publishing_destination_output',
]

@pulumi.output_type
class GetPublishingDestinationResult:
    def __init__(__self__, destination_properties=None, destination_type=None, id=None, publishing_failure_start_timestamp=None, status=None, tags=None):
        if destination_properties and not isinstance(destination_properties, dict):
            raise TypeError("Expected argument 'destination_properties' to be a dict")
        pulumi.set(__self__, "destination_properties", destination_properties)
        if destination_type and not isinstance(destination_type, str):
            raise TypeError("Expected argument 'destination_type' to be a str")
        pulumi.set(__self__, "destination_type", destination_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if publishing_failure_start_timestamp and not isinstance(publishing_failure_start_timestamp, str):
            raise TypeError("Expected argument 'publishing_failure_start_timestamp' to be a str")
        pulumi.set(__self__, "publishing_failure_start_timestamp", publishing_failure_start_timestamp)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="destinationProperties")
    def destination_properties(self) -> Optional['outputs.PublishingDestinationCfnDestinationProperties']:
        """
        Contains the Amazon Resource Name (ARN) of the resource to publish to, such as an S3 bucket, and the ARN of the KMS key to use to encrypt published findings.
        """
        return pulumi.get(self, "destination_properties")

    @property
    @pulumi.getter(name="destinationType")
    def destination_type(self) -> Optional[builtins.str]:
        """
        The type of resource for the publishing destination. Currently only Amazon S3 buckets are supported.
        """
        return pulumi.get(self, "destination_type")

    @property
    @pulumi.getter
    def id(self) -> Optional[builtins.str]:
        """
        The ID of the publishing destination.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="publishingFailureStartTimestamp")
    def publishing_failure_start_timestamp(self) -> Optional[builtins.str]:
        """
        The time, in epoch millisecond format, at which GuardDuty was first unable to publish findings to the destination.
        """
        return pulumi.get(self, "publishing_failure_start_timestamp")

    @property
    @pulumi.getter
    def status(self) -> Optional[builtins.str]:
        """
        The status of the publishing destination.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        Describes a tag.
        """
        return pulumi.get(self, "tags")


class AwaitableGetPublishingDestinationResult(GetPublishingDestinationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPublishingDestinationResult(
            destination_properties=self.destination_properties,
            destination_type=self.destination_type,
            id=self.id,
            publishing_failure_start_timestamp=self.publishing_failure_start_timestamp,
            status=self.status,
            tags=self.tags)


def get_publishing_destination(detector_id: Optional[builtins.str] = None,
                               id: Optional[builtins.str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPublishingDestinationResult:
    """
    Resource Type definition for AWS::GuardDuty::PublishingDestination.


    :param builtins.str detector_id: The ID of the GuardDuty detector associated with the publishing destination.
    :param builtins.str id: The ID of the publishing destination.
    """
    __args__ = dict()
    __args__['detectorId'] = detector_id
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:guardduty:getPublishingDestination', __args__, opts=opts, typ=GetPublishingDestinationResult).value

    return AwaitableGetPublishingDestinationResult(
        destination_properties=pulumi.get(__ret__, 'destination_properties'),
        destination_type=pulumi.get(__ret__, 'destination_type'),
        id=pulumi.get(__ret__, 'id'),
        publishing_failure_start_timestamp=pulumi.get(__ret__, 'publishing_failure_start_timestamp'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'))
def get_publishing_destination_output(detector_id: Optional[pulumi.Input[builtins.str]] = None,
                                      id: Optional[pulumi.Input[builtins.str]] = None,
                                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetPublishingDestinationResult]:
    """
    Resource Type definition for AWS::GuardDuty::PublishingDestination.


    :param builtins.str detector_id: The ID of the GuardDuty detector associated with the publishing destination.
    :param builtins.str id: The ID of the publishing destination.
    """
    __args__ = dict()
    __args__['detectorId'] = detector_id
    __args__['id'] = id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:guardduty:getPublishingDestination', __args__, opts=opts, typ=GetPublishingDestinationResult)
    return __ret__.apply(lambda __response__: GetPublishingDestinationResult(
        destination_properties=pulumi.get(__response__, 'destination_properties'),
        destination_type=pulumi.get(__response__, 'destination_type'),
        id=pulumi.get(__response__, 'id'),
        publishing_failure_start_timestamp=pulumi.get(__response__, 'publishing_failure_start_timestamp'),
        status=pulumi.get(__response__, 'status'),
        tags=pulumi.get(__response__, 'tags')))
