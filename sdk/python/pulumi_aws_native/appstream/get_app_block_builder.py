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
    'GetAppBlockBuilderResult',
    'AwaitableGetAppBlockBuilderResult',
    'get_app_block_builder',
    'get_app_block_builder_output',
]

@pulumi.output_type
class GetAppBlockBuilderResult:
    def __init__(__self__, access_endpoints=None, arn=None, created_time=None, description=None, display_name=None, enable_default_internet_access=None, iam_role_arn=None, instance_type=None, platform=None, vpc_config=None):
        if access_endpoints and not isinstance(access_endpoints, list):
            raise TypeError("Expected argument 'access_endpoints' to be a list")
        pulumi.set(__self__, "access_endpoints", access_endpoints)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if created_time and not isinstance(created_time, str):
            raise TypeError("Expected argument 'created_time' to be a str")
        pulumi.set(__self__, "created_time", created_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if enable_default_internet_access and not isinstance(enable_default_internet_access, bool):
            raise TypeError("Expected argument 'enable_default_internet_access' to be a bool")
        pulumi.set(__self__, "enable_default_internet_access", enable_default_internet_access)
        if iam_role_arn and not isinstance(iam_role_arn, str):
            raise TypeError("Expected argument 'iam_role_arn' to be a str")
        pulumi.set(__self__, "iam_role_arn", iam_role_arn)
        if instance_type and not isinstance(instance_type, str):
            raise TypeError("Expected argument 'instance_type' to be a str")
        pulumi.set(__self__, "instance_type", instance_type)
        if platform and not isinstance(platform, str):
            raise TypeError("Expected argument 'platform' to be a str")
        pulumi.set(__self__, "platform", platform)
        if vpc_config and not isinstance(vpc_config, dict):
            raise TypeError("Expected argument 'vpc_config' to be a dict")
        pulumi.set(__self__, "vpc_config", vpc_config)

    @property
    @pulumi.getter(name="accessEndpoints")
    def access_endpoints(self) -> Optional[Sequence['outputs.AppBlockBuilderAccessEndpoint']]:
        return pulumi.get(self, "access_endpoints")

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> Optional[str]:
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="enableDefaultInternetAccess")
    def enable_default_internet_access(self) -> Optional[bool]:
        return pulumi.get(self, "enable_default_internet_access")

    @property
    @pulumi.getter(name="iamRoleArn")
    def iam_role_arn(self) -> Optional[str]:
        return pulumi.get(self, "iam_role_arn")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[str]:
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter
    def platform(self) -> Optional[str]:
        return pulumi.get(self, "platform")

    @property
    @pulumi.getter(name="vpcConfig")
    def vpc_config(self) -> Optional['outputs.AppBlockBuilderVpcConfig']:
        return pulumi.get(self, "vpc_config")


class AwaitableGetAppBlockBuilderResult(GetAppBlockBuilderResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppBlockBuilderResult(
            access_endpoints=self.access_endpoints,
            arn=self.arn,
            created_time=self.created_time,
            description=self.description,
            display_name=self.display_name,
            enable_default_internet_access=self.enable_default_internet_access,
            iam_role_arn=self.iam_role_arn,
            instance_type=self.instance_type,
            platform=self.platform,
            vpc_config=self.vpc_config)


def get_app_block_builder(name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppBlockBuilderResult:
    """
    Resource Type definition for AWS::AppStream::AppBlockBuilder.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:appstream:getAppBlockBuilder', __args__, opts=opts, typ=GetAppBlockBuilderResult).value

    return AwaitableGetAppBlockBuilderResult(
        access_endpoints=pulumi.get(__ret__, 'access_endpoints'),
        arn=pulumi.get(__ret__, 'arn'),
        created_time=pulumi.get(__ret__, 'created_time'),
        description=pulumi.get(__ret__, 'description'),
        display_name=pulumi.get(__ret__, 'display_name'),
        enable_default_internet_access=pulumi.get(__ret__, 'enable_default_internet_access'),
        iam_role_arn=pulumi.get(__ret__, 'iam_role_arn'),
        instance_type=pulumi.get(__ret__, 'instance_type'),
        platform=pulumi.get(__ret__, 'platform'),
        vpc_config=pulumi.get(__ret__, 'vpc_config'))


@_utilities.lift_output_func(get_app_block_builder)
def get_app_block_builder_output(name: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppBlockBuilderResult]:
    """
    Resource Type definition for AWS::AppStream::AppBlockBuilder.
    """
    ...