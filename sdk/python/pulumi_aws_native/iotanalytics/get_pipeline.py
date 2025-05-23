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
    'GetPipelineResult',
    'AwaitableGetPipelineResult',
    'get_pipeline',
    'get_pipeline_output',
]

@pulumi.output_type
class GetPipelineResult:
    def __init__(__self__, id=None, pipeline_activities=None, tags=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if pipeline_activities and not isinstance(pipeline_activities, list):
            raise TypeError("Expected argument 'pipeline_activities' to be a list")
        pulumi.set(__self__, "pipeline_activities", pipeline_activities)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def id(self) -> Optional[builtins.str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="pipelineActivities")
    def pipeline_activities(self) -> Optional[Sequence['outputs.PipelineActivity']]:
        """
        A list of "PipelineActivity" objects. Activities perform transformations on your messages, such as removing, renaming or adding message attributes; filtering messages based on attribute values; invoking your Lambda functions on messages for advanced processing; or performing mathematical transformations to normalize device data.

        The list can be 2-25 *PipelineActivity* objects and must contain both a `channel` and a `datastore` activity. Each entry in the list must contain only one activity, for example:

        `pipelineActivities = [ { "channel": { ... } }, { "lambda": { ... } }, ... ]`
        """
        return pulumi.get(self, "pipeline_activities")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        Metadata which can be used to manage the pipeline.

        For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
        """
        return pulumi.get(self, "tags")


class AwaitableGetPipelineResult(GetPipelineResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPipelineResult(
            id=self.id,
            pipeline_activities=self.pipeline_activities,
            tags=self.tags)


def get_pipeline(pipeline_name: Optional[builtins.str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPipelineResult:
    """
    Resource Type definition for AWS::IoTAnalytics::Pipeline


    :param builtins.str pipeline_name: The name of the pipeline.
    """
    __args__ = dict()
    __args__['pipelineName'] = pipeline_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:iotanalytics:getPipeline', __args__, opts=opts, typ=GetPipelineResult).value

    return AwaitableGetPipelineResult(
        id=pulumi.get(__ret__, 'id'),
        pipeline_activities=pulumi.get(__ret__, 'pipeline_activities'),
        tags=pulumi.get(__ret__, 'tags'))
def get_pipeline_output(pipeline_name: Optional[pulumi.Input[builtins.str]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetPipelineResult]:
    """
    Resource Type definition for AWS::IoTAnalytics::Pipeline


    :param builtins.str pipeline_name: The name of the pipeline.
    """
    __args__ = dict()
    __args__['pipelineName'] = pipeline_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:iotanalytics:getPipeline', __args__, opts=opts, typ=GetPipelineResult)
    return __ret__.apply(lambda __response__: GetPipelineResult(
        id=pulumi.get(__response__, 'id'),
        pipeline_activities=pulumi.get(__response__, 'pipeline_activities'),
        tags=pulumi.get(__response__, 'tags')))
