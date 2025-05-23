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
    'GetModelExplainabilityJobDefinitionResult',
    'AwaitableGetModelExplainabilityJobDefinitionResult',
    'get_model_explainability_job_definition',
    'get_model_explainability_job_definition_output',
]

@pulumi.output_type
class GetModelExplainabilityJobDefinitionResult:
    def __init__(__self__, creation_time=None, job_definition_arn=None):
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if job_definition_arn and not isinstance(job_definition_arn, str):
            raise TypeError("Expected argument 'job_definition_arn' to be a str")
        pulumi.set(__self__, "job_definition_arn", job_definition_arn)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[builtins.str]:
        """
        The time at which the job definition was created.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="jobDefinitionArn")
    def job_definition_arn(self) -> Optional[builtins.str]:
        """
        The Amazon Resource Name (ARN) of job definition.
        """
        return pulumi.get(self, "job_definition_arn")


class AwaitableGetModelExplainabilityJobDefinitionResult(GetModelExplainabilityJobDefinitionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetModelExplainabilityJobDefinitionResult(
            creation_time=self.creation_time,
            job_definition_arn=self.job_definition_arn)


def get_model_explainability_job_definition(job_definition_arn: Optional[builtins.str] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetModelExplainabilityJobDefinitionResult:
    """
    Resource Type definition for AWS::SageMaker::ModelExplainabilityJobDefinition


    :param builtins.str job_definition_arn: The Amazon Resource Name (ARN) of job definition.
    """
    __args__ = dict()
    __args__['jobDefinitionArn'] = job_definition_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:sagemaker:getModelExplainabilityJobDefinition', __args__, opts=opts, typ=GetModelExplainabilityJobDefinitionResult).value

    return AwaitableGetModelExplainabilityJobDefinitionResult(
        creation_time=pulumi.get(__ret__, 'creation_time'),
        job_definition_arn=pulumi.get(__ret__, 'job_definition_arn'))
def get_model_explainability_job_definition_output(job_definition_arn: Optional[pulumi.Input[builtins.str]] = None,
                                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetModelExplainabilityJobDefinitionResult]:
    """
    Resource Type definition for AWS::SageMaker::ModelExplainabilityJobDefinition


    :param builtins.str job_definition_arn: The Amazon Resource Name (ARN) of job definition.
    """
    __args__ = dict()
    __args__['jobDefinitionArn'] = job_definition_arn
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:sagemaker:getModelExplainabilityJobDefinition', __args__, opts=opts, typ=GetModelExplainabilityJobDefinitionResult)
    return __ret__.apply(lambda __response__: GetModelExplainabilityJobDefinitionResult(
        creation_time=pulumi.get(__response__, 'creation_time'),
        job_definition_arn=pulumi.get(__response__, 'job_definition_arn')))
