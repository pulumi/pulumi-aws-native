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
from ._enums import *

__all__ = [
    'GetQueueResult',
    'AwaitableGetQueueResult',
    'get_queue',
    'get_queue_output',
]

@pulumi.output_type
class GetQueueResult:
    def __init__(__self__, description=None, hours_of_operation_arn=None, instance_arn=None, max_contacts=None, name=None, outbound_caller_config=None, queue_arn=None, quick_connect_arns=None, status=None, tags=None, type=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if hours_of_operation_arn and not isinstance(hours_of_operation_arn, str):
            raise TypeError("Expected argument 'hours_of_operation_arn' to be a str")
        pulumi.set(__self__, "hours_of_operation_arn", hours_of_operation_arn)
        if instance_arn and not isinstance(instance_arn, str):
            raise TypeError("Expected argument 'instance_arn' to be a str")
        pulumi.set(__self__, "instance_arn", instance_arn)
        if max_contacts and not isinstance(max_contacts, int):
            raise TypeError("Expected argument 'max_contacts' to be a int")
        pulumi.set(__self__, "max_contacts", max_contacts)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if outbound_caller_config and not isinstance(outbound_caller_config, dict):
            raise TypeError("Expected argument 'outbound_caller_config' to be a dict")
        pulumi.set(__self__, "outbound_caller_config", outbound_caller_config)
        if queue_arn and not isinstance(queue_arn, str):
            raise TypeError("Expected argument 'queue_arn' to be a str")
        pulumi.set(__self__, "queue_arn", queue_arn)
        if quick_connect_arns and not isinstance(quick_connect_arns, list):
            raise TypeError("Expected argument 'quick_connect_arns' to be a list")
        pulumi.set(__self__, "quick_connect_arns", quick_connect_arns)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the queue.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="hoursOfOperationArn")
    def hours_of_operation_arn(self) -> Optional[str]:
        """
        The identifier for the hours of operation.
        """
        return pulumi.get(self, "hours_of_operation_arn")

    @property
    @pulumi.getter(name="instanceArn")
    def instance_arn(self) -> Optional[str]:
        """
        The identifier of the Amazon Connect instance.
        """
        return pulumi.get(self, "instance_arn")

    @property
    @pulumi.getter(name="maxContacts")
    def max_contacts(self) -> Optional[int]:
        """
        The maximum number of contacts that can be in the queue before it is considered full.
        """
        return pulumi.get(self, "max_contacts")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the queue.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outboundCallerConfig")
    def outbound_caller_config(self) -> Optional['outputs.QueueOutboundCallerConfig']:
        """
        The outbound caller ID name, number, and outbound whisper flow.
        """
        return pulumi.get(self, "outbound_caller_config")

    @property
    @pulumi.getter(name="queueArn")
    def queue_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) for the queue.
        """
        return pulumi.get(self, "queue_arn")

    @property
    @pulumi.getter(name="quickConnectArns")
    def quick_connect_arns(self) -> Optional[Sequence[str]]:
        """
        The quick connects available to agents who are working the queue.
        """
        return pulumi.get(self, "quick_connect_arns")

    @property
    @pulumi.getter
    def status(self) -> Optional['QueueStatus']:
        """
        The status of the queue.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.QueueTag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> Optional['QueueType']:
        """
        The type of queue.
        """
        return pulumi.get(self, "type")


class AwaitableGetQueueResult(GetQueueResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetQueueResult(
            description=self.description,
            hours_of_operation_arn=self.hours_of_operation_arn,
            instance_arn=self.instance_arn,
            max_contacts=self.max_contacts,
            name=self.name,
            outbound_caller_config=self.outbound_caller_config,
            queue_arn=self.queue_arn,
            quick_connect_arns=self.quick_connect_arns,
            status=self.status,
            tags=self.tags,
            type=self.type)


def get_queue(queue_arn: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetQueueResult:
    """
    Resource Type definition for AWS::Connect::Queue


    :param str queue_arn: The Amazon Resource Name (ARN) for the queue.
    """
    __args__ = dict()
    __args__['queueArn'] = queue_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:connect:getQueue', __args__, opts=opts, typ=GetQueueResult).value

    return AwaitableGetQueueResult(
        description=pulumi.get(__ret__, 'description'),
        hours_of_operation_arn=pulumi.get(__ret__, 'hours_of_operation_arn'),
        instance_arn=pulumi.get(__ret__, 'instance_arn'),
        max_contacts=pulumi.get(__ret__, 'max_contacts'),
        name=pulumi.get(__ret__, 'name'),
        outbound_caller_config=pulumi.get(__ret__, 'outbound_caller_config'),
        queue_arn=pulumi.get(__ret__, 'queue_arn'),
        quick_connect_arns=pulumi.get(__ret__, 'quick_connect_arns'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'))


@_utilities.lift_output_func(get_queue)
def get_queue_output(queue_arn: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetQueueResult]:
    """
    Resource Type definition for AWS::Connect::Queue


    :param str queue_arn: The Amazon Resource Name (ARN) for the queue.
    """
    ...