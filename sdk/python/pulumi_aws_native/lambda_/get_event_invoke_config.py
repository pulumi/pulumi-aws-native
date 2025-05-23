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
    'GetEventInvokeConfigResult',
    'AwaitableGetEventInvokeConfigResult',
    'get_event_invoke_config',
    'get_event_invoke_config_output',
]

@pulumi.output_type
class GetEventInvokeConfigResult:
    def __init__(__self__, destination_config=None, maximum_event_age_in_seconds=None, maximum_retry_attempts=None):
        if destination_config and not isinstance(destination_config, dict):
            raise TypeError("Expected argument 'destination_config' to be a dict")
        pulumi.set(__self__, "destination_config", destination_config)
        if maximum_event_age_in_seconds and not isinstance(maximum_event_age_in_seconds, int):
            raise TypeError("Expected argument 'maximum_event_age_in_seconds' to be a int")
        pulumi.set(__self__, "maximum_event_age_in_seconds", maximum_event_age_in_seconds)
        if maximum_retry_attempts and not isinstance(maximum_retry_attempts, int):
            raise TypeError("Expected argument 'maximum_retry_attempts' to be a int")
        pulumi.set(__self__, "maximum_retry_attempts", maximum_retry_attempts)

    @property
    @pulumi.getter(name="destinationConfig")
    def destination_config(self) -> Optional['outputs.EventInvokeConfigDestinationConfig']:
        """
        A destination for events after they have been sent to a function for processing.

        **Destinations** - *Function* - The Amazon Resource Name (ARN) of a Lambda function.
        - *Queue* - The ARN of a standard SQS queue.
        - *Bucket* - The ARN of an Amazon S3 bucket.
        - *Topic* - The ARN of a standard SNS topic.
        - *Event Bus* - The ARN of an Amazon EventBridge event bus.

        > S3 buckets are supported only for on-failure destinations. To retain records of successful invocations, use another destination type.
        """
        return pulumi.get(self, "destination_config")

    @property
    @pulumi.getter(name="maximumEventAgeInSeconds")
    def maximum_event_age_in_seconds(self) -> Optional[builtins.int]:
        """
        The maximum age of a request that Lambda sends to a function for processing.
        """
        return pulumi.get(self, "maximum_event_age_in_seconds")

    @property
    @pulumi.getter(name="maximumRetryAttempts")
    def maximum_retry_attempts(self) -> Optional[builtins.int]:
        """
        The maximum number of times to retry when the function returns an error.
        """
        return pulumi.get(self, "maximum_retry_attempts")


class AwaitableGetEventInvokeConfigResult(GetEventInvokeConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEventInvokeConfigResult(
            destination_config=self.destination_config,
            maximum_event_age_in_seconds=self.maximum_event_age_in_seconds,
            maximum_retry_attempts=self.maximum_retry_attempts)


def get_event_invoke_config(function_name: Optional[builtins.str] = None,
                            qualifier: Optional[builtins.str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEventInvokeConfigResult:
    """
    The AWS::Lambda::EventInvokeConfig resource configures options for asynchronous invocation on a version or an alias.


    :param builtins.str function_name: The name of the Lambda function.
    :param builtins.str qualifier: The identifier of a version or alias.
    """
    __args__ = dict()
    __args__['functionName'] = function_name
    __args__['qualifier'] = qualifier
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:lambda:getEventInvokeConfig', __args__, opts=opts, typ=GetEventInvokeConfigResult).value

    return AwaitableGetEventInvokeConfigResult(
        destination_config=pulumi.get(__ret__, 'destination_config'),
        maximum_event_age_in_seconds=pulumi.get(__ret__, 'maximum_event_age_in_seconds'),
        maximum_retry_attempts=pulumi.get(__ret__, 'maximum_retry_attempts'))
def get_event_invoke_config_output(function_name: Optional[pulumi.Input[builtins.str]] = None,
                                   qualifier: Optional[pulumi.Input[builtins.str]] = None,
                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetEventInvokeConfigResult]:
    """
    The AWS::Lambda::EventInvokeConfig resource configures options for asynchronous invocation on a version or an alias.


    :param builtins.str function_name: The name of the Lambda function.
    :param builtins.str qualifier: The identifier of a version or alias.
    """
    __args__ = dict()
    __args__['functionName'] = function_name
    __args__['qualifier'] = qualifier
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:lambda:getEventInvokeConfig', __args__, opts=opts, typ=GetEventInvokeConfigResult)
    return __ret__.apply(lambda __response__: GetEventInvokeConfigResult(
        destination_config=pulumi.get(__response__, 'destination_config'),
        maximum_event_age_in_seconds=pulumi.get(__response__, 'maximum_event_age_in_seconds'),
        maximum_retry_attempts=pulumi.get(__response__, 'maximum_retry_attempts')))
