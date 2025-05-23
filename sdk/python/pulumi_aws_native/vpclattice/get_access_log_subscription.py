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
    'GetAccessLogSubscriptionResult',
    'AwaitableGetAccessLogSubscriptionResult',
    'get_access_log_subscription',
    'get_access_log_subscription_output',
]

@pulumi.output_type
class GetAccessLogSubscriptionResult:
    def __init__(__self__, arn=None, destination_arn=None, id=None, resource_arn=None, resource_id=None, service_network_log_type=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if destination_arn and not isinstance(destination_arn, str):
            raise TypeError("Expected argument 'destination_arn' to be a str")
        pulumi.set(__self__, "destination_arn", destination_arn)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if resource_arn and not isinstance(resource_arn, str):
            raise TypeError("Expected argument 'resource_arn' to be a str")
        pulumi.set(__self__, "resource_arn", resource_arn)
        if resource_id and not isinstance(resource_id, str):
            raise TypeError("Expected argument 'resource_id' to be a str")
        pulumi.set(__self__, "resource_id", resource_id)
        if service_network_log_type and not isinstance(service_network_log_type, str):
            raise TypeError("Expected argument 'service_network_log_type' to be a str")
        pulumi.set(__self__, "service_network_log_type", service_network_log_type)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the access log subscription.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="destinationArn")
    def destination_arn(self) -> Optional[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the destination. The supported destination types are CloudWatch Log groups, Kinesis Data Firehose delivery streams, and Amazon S3 buckets.
        """
        return pulumi.get(self, "destination_arn")

    @property
    @pulumi.getter
    def id(self) -> Optional[builtins.str]:
        """
        The ID of the access log subscription.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="resourceArn")
    def resource_arn(self) -> Optional[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the access log subscription.
        """
        return pulumi.get(self, "resource_arn")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[builtins.str]:
        """
        The ID of the service network or service.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="serviceNetworkLogType")
    def service_network_log_type(self) -> Optional['AccessLogSubscriptionServiceNetworkLogType']:
        """
        Log type of the service network.
        """
        return pulumi.get(self, "service_network_log_type")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        The tags for the access log subscription.
        """
        return pulumi.get(self, "tags")


class AwaitableGetAccessLogSubscriptionResult(GetAccessLogSubscriptionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccessLogSubscriptionResult(
            arn=self.arn,
            destination_arn=self.destination_arn,
            id=self.id,
            resource_arn=self.resource_arn,
            resource_id=self.resource_id,
            service_network_log_type=self.service_network_log_type,
            tags=self.tags)


def get_access_log_subscription(arn: Optional[builtins.str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccessLogSubscriptionResult:
    """
    Enables access logs to be sent to Amazon CloudWatch, Amazon S3, and Amazon Kinesis Data Firehose. The service network owner can use the access logs to audit the services in the network. The service network owner will only see access logs from clients and services that are associated with their service network. Access log entries represent traffic originated from VPCs associated with that network.


    :param builtins.str arn: The Amazon Resource Name (ARN) of the access log subscription.
    """
    __args__ = dict()
    __args__['arn'] = arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:vpclattice:getAccessLogSubscription', __args__, opts=opts, typ=GetAccessLogSubscriptionResult).value

    return AwaitableGetAccessLogSubscriptionResult(
        arn=pulumi.get(__ret__, 'arn'),
        destination_arn=pulumi.get(__ret__, 'destination_arn'),
        id=pulumi.get(__ret__, 'id'),
        resource_arn=pulumi.get(__ret__, 'resource_arn'),
        resource_id=pulumi.get(__ret__, 'resource_id'),
        service_network_log_type=pulumi.get(__ret__, 'service_network_log_type'),
        tags=pulumi.get(__ret__, 'tags'))
def get_access_log_subscription_output(arn: Optional[pulumi.Input[builtins.str]] = None,
                                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAccessLogSubscriptionResult]:
    """
    Enables access logs to be sent to Amazon CloudWatch, Amazon S3, and Amazon Kinesis Data Firehose. The service network owner can use the access logs to audit the services in the network. The service network owner will only see access logs from clients and services that are associated with their service network. Access log entries represent traffic originated from VPCs associated with that network.


    :param builtins.str arn: The Amazon Resource Name (ARN) of the access log subscription.
    """
    __args__ = dict()
    __args__['arn'] = arn
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:vpclattice:getAccessLogSubscription', __args__, opts=opts, typ=GetAccessLogSubscriptionResult)
    return __ret__.apply(lambda __response__: GetAccessLogSubscriptionResult(
        arn=pulumi.get(__response__, 'arn'),
        destination_arn=pulumi.get(__response__, 'destination_arn'),
        id=pulumi.get(__response__, 'id'),
        resource_arn=pulumi.get(__response__, 'resource_arn'),
        resource_id=pulumi.get(__response__, 'resource_id'),
        service_network_log_type=pulumi.get(__response__, 'service_network_log_type'),
        tags=pulumi.get(__response__, 'tags')))
