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
from ._enums import *

__all__ = [
    'GetQueueEnvironmentResult',
    'AwaitableGetQueueEnvironmentResult',
    'get_queue_environment',
    'get_queue_environment_output',
]

@pulumi.output_type
class GetQueueEnvironmentResult:
    def __init__(__self__, name=None, priority=None, queue_environment_id=None, template=None, template_type=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if priority and not isinstance(priority, int):
            raise TypeError("Expected argument 'priority' to be a int")
        pulumi.set(__self__, "priority", priority)
        if queue_environment_id and not isinstance(queue_environment_id, str):
            raise TypeError("Expected argument 'queue_environment_id' to be a str")
        pulumi.set(__self__, "queue_environment_id", queue_environment_id)
        if template and not isinstance(template, str):
            raise TypeError("Expected argument 'template' to be a str")
        pulumi.set(__self__, "template", template)
        if template_type and not isinstance(template_type, str):
            raise TypeError("Expected argument 'template_type' to be a str")
        pulumi.set(__self__, "template_type", template_type)

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        """
        The name of the queue environment.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> Optional[builtins.int]:
        """
        The queue environment's priority.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="queueEnvironmentId")
    def queue_environment_id(self) -> Optional[builtins.str]:
        """
        The queue environment ID.
        """
        return pulumi.get(self, "queue_environment_id")

    @property
    @pulumi.getter
    def template(self) -> Optional[builtins.str]:
        """
        A JSON or YAML template that describes the processing environment for the queue.
        """
        return pulumi.get(self, "template")

    @property
    @pulumi.getter(name="templateType")
    def template_type(self) -> Optional['QueueEnvironmentEnvironmentTemplateType']:
        """
        Specifies whether the template for the queue environment is JSON or YAML.
        """
        return pulumi.get(self, "template_type")


class AwaitableGetQueueEnvironmentResult(GetQueueEnvironmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetQueueEnvironmentResult(
            name=self.name,
            priority=self.priority,
            queue_environment_id=self.queue_environment_id,
            template=self.template,
            template_type=self.template_type)


def get_queue_environment(farm_id: Optional[builtins.str] = None,
                          queue_environment_id: Optional[builtins.str] = None,
                          queue_id: Optional[builtins.str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetQueueEnvironmentResult:
    """
    Definition of AWS::Deadline::QueueEnvironment Resource Type


    :param builtins.str farm_id: The identifier assigned to the farm that contains the queue.
    :param builtins.str queue_environment_id: The queue environment ID.
    :param builtins.str queue_id: The unique identifier of the queue that contains the environment.
    """
    __args__ = dict()
    __args__['farmId'] = farm_id
    __args__['queueEnvironmentId'] = queue_environment_id
    __args__['queueId'] = queue_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:deadline:getQueueEnvironment', __args__, opts=opts, typ=GetQueueEnvironmentResult).value

    return AwaitableGetQueueEnvironmentResult(
        name=pulumi.get(__ret__, 'name'),
        priority=pulumi.get(__ret__, 'priority'),
        queue_environment_id=pulumi.get(__ret__, 'queue_environment_id'),
        template=pulumi.get(__ret__, 'template'),
        template_type=pulumi.get(__ret__, 'template_type'))
def get_queue_environment_output(farm_id: Optional[pulumi.Input[builtins.str]] = None,
                                 queue_environment_id: Optional[pulumi.Input[builtins.str]] = None,
                                 queue_id: Optional[pulumi.Input[builtins.str]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetQueueEnvironmentResult]:
    """
    Definition of AWS::Deadline::QueueEnvironment Resource Type


    :param builtins.str farm_id: The identifier assigned to the farm that contains the queue.
    :param builtins.str queue_environment_id: The queue environment ID.
    :param builtins.str queue_id: The unique identifier of the queue that contains the environment.
    """
    __args__ = dict()
    __args__['farmId'] = farm_id
    __args__['queueEnvironmentId'] = queue_environment_id
    __args__['queueId'] = queue_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:deadline:getQueueEnvironment', __args__, opts=opts, typ=GetQueueEnvironmentResult)
    return __ret__.apply(lambda __response__: GetQueueEnvironmentResult(
        name=pulumi.get(__response__, 'name'),
        priority=pulumi.get(__response__, 'priority'),
        queue_environment_id=pulumi.get(__response__, 'queue_environment_id'),
        template=pulumi.get(__response__, 'template'),
        template_type=pulumi.get(__response__, 'template_type')))
