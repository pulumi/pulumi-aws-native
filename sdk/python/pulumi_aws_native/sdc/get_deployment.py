# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetDeploymentResult',
    'AwaitableGetDeploymentResult',
    'get_deployment',
    'get_deployment_output',
]

@pulumi.output_type
class GetDeploymentResult:
    def __init__(__self__, config_name=None, dimension=None, id=None, pipeline_id=None, s3_bucket=None, stage=None, target_region_override=None):
        if config_name and not isinstance(config_name, str):
            raise TypeError("Expected argument 'config_name' to be a str")
        pulumi.set(__self__, "config_name", config_name)
        if dimension and not isinstance(dimension, str):
            raise TypeError("Expected argument 'dimension' to be a str")
        pulumi.set(__self__, "dimension", dimension)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if pipeline_id and not isinstance(pipeline_id, str):
            raise TypeError("Expected argument 'pipeline_id' to be a str")
        pulumi.set(__self__, "pipeline_id", pipeline_id)
        if s3_bucket and not isinstance(s3_bucket, str):
            raise TypeError("Expected argument 's3_bucket' to be a str")
        pulumi.set(__self__, "s3_bucket", s3_bucket)
        if stage and not isinstance(stage, str):
            raise TypeError("Expected argument 'stage' to be a str")
        pulumi.set(__self__, "stage", stage)
        if target_region_override and not isinstance(target_region_override, str):
            raise TypeError("Expected argument 'target_region_override' to be a str")
        pulumi.set(__self__, "target_region_override", target_region_override)

    @property
    @pulumi.getter(name="configName")
    def config_name(self) -> Optional[str]:
        return pulumi.get(self, "config_name")

    @property
    @pulumi.getter
    def dimension(self) -> Optional[str]:
        return pulumi.get(self, "dimension")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="pipelineId")
    def pipeline_id(self) -> Optional[str]:
        return pulumi.get(self, "pipeline_id")

    @property
    @pulumi.getter(name="s3Bucket")
    def s3_bucket(self) -> Optional[str]:
        return pulumi.get(self, "s3_bucket")

    @property
    @pulumi.getter
    def stage(self) -> Optional[str]:
        return pulumi.get(self, "stage")

    @property
    @pulumi.getter(name="targetRegionOverride")
    def target_region_override(self) -> Optional[str]:
        return pulumi.get(self, "target_region_override")


class AwaitableGetDeploymentResult(GetDeploymentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeploymentResult(
            config_name=self.config_name,
            dimension=self.dimension,
            id=self.id,
            pipeline_id=self.pipeline_id,
            s3_bucket=self.s3_bucket,
            stage=self.stage,
            target_region_override=self.target_region_override)


def get_deployment(id: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeploymentResult:
    """
    Resource Type definition for AMZN::SDC::Deployment
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:sdc:getDeployment', __args__, opts=opts, typ=GetDeploymentResult).value

    return AwaitableGetDeploymentResult(
        config_name=pulumi.get(__ret__, 'config_name'),
        dimension=pulumi.get(__ret__, 'dimension'),
        id=pulumi.get(__ret__, 'id'),
        pipeline_id=pulumi.get(__ret__, 'pipeline_id'),
        s3_bucket=pulumi.get(__ret__, 's3_bucket'),
        stage=pulumi.get(__ret__, 'stage'),
        target_region_override=pulumi.get(__ret__, 'target_region_override'))


@_utilities.lift_output_func(get_deployment)
def get_deployment_output(id: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDeploymentResult]:
    """
    Resource Type definition for AMZN::SDC::Deployment
    """
    ...