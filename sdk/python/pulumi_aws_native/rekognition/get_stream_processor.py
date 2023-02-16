# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetStreamProcessorResult',
    'AwaitableGetStreamProcessorResult',
    'get_stream_processor',
    'get_stream_processor_output',
]

@pulumi.output_type
class GetStreamProcessorResult:
    def __init__(__self__, arn=None, status=None, status_message=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
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
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Current status of the stream processor.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusMessage")
    def status_message(self) -> Optional[str]:
        """
        Detailed status message about the stream processor.
        """
        return pulumi.get(self, "status_message")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.StreamProcessorTag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetStreamProcessorResult(GetStreamProcessorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStreamProcessorResult(
            arn=self.arn,
            status=self.status,
            status_message=self.status_message,
            tags=self.tags)


def get_stream_processor(name: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStreamProcessorResult:
    """
    The AWS::Rekognition::StreamProcessor type is used to create an Amazon Rekognition StreamProcessor that you can use to analyze streaming videos.


    :param str name: Name of the stream processor. It's an identifier you assign to the stream processor. You can use it to manage the stream processor.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:rekognition:getStreamProcessor', __args__, opts=opts, typ=GetStreamProcessorResult).value

    return AwaitableGetStreamProcessorResult(
        arn=__ret__.arn,
        status=__ret__.status,
        status_message=__ret__.status_message,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_stream_processor)
def get_stream_processor_output(name: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetStreamProcessorResult]:
    """
    The AWS::Rekognition::StreamProcessor type is used to create an Amazon Rekognition StreamProcessor that you can use to analyze streaming videos.


    :param str name: Name of the stream processor. It's an identifier you assign to the stream processor. You can use it to manage the stream processor.
    """
    ...