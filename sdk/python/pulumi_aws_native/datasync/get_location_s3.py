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
from ._enums import *

__all__ = [
    'GetLocationS3Result',
    'AwaitableGetLocationS3Result',
    'get_location_s3',
    'get_location_s3_output',
]

@pulumi.output_type
class GetLocationS3Result:
    def __init__(__self__, location_arn=None, location_uri=None, s3_config=None, s3_storage_class=None, tags=None):
        if location_arn and not isinstance(location_arn, str):
            raise TypeError("Expected argument 'location_arn' to be a str")
        pulumi.set(__self__, "location_arn", location_arn)
        if location_uri and not isinstance(location_uri, str):
            raise TypeError("Expected argument 'location_uri' to be a str")
        pulumi.set(__self__, "location_uri", location_uri)
        if s3_config and not isinstance(s3_config, dict):
            raise TypeError("Expected argument 's3_config' to be a dict")
        pulumi.set(__self__, "s3_config", s3_config)
        if s3_storage_class and not isinstance(s3_storage_class, str):
            raise TypeError("Expected argument 's3_storage_class' to be a str")
        pulumi.set(__self__, "s3_storage_class", s3_storage_class)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="locationArn")
    def location_arn(self) -> Optional[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the Amazon S3 bucket location.
        """
        return pulumi.get(self, "location_arn")

    @property
    @pulumi.getter(name="locationUri")
    def location_uri(self) -> Optional[builtins.str]:
        """
        The URL of the S3 location that was described.
        """
        return pulumi.get(self, "location_uri")

    @property
    @pulumi.getter(name="s3Config")
    def s3_config(self) -> Optional['outputs.LocationS3s3Config']:
        """
        The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that is used to access an Amazon S3 bucket.

        For detailed information about using such a role, see [Creating a Location for Amazon S3](https://docs.aws.amazon.com/datasync/latest/userguide/working-with-locations.html#create-s3-location) in the *AWS DataSync User Guide* .
        """
        return pulumi.get(self, "s3_config")

    @property
    @pulumi.getter(name="s3StorageClass")
    def s3_storage_class(self) -> Optional['LocationS3S3StorageClass']:
        """
        The Amazon S3 storage class you want to store your files in when this location is used as a task destination.
        """
        return pulumi.get(self, "s3_storage_class")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetLocationS3Result(GetLocationS3Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLocationS3Result(
            location_arn=self.location_arn,
            location_uri=self.location_uri,
            s3_config=self.s3_config,
            s3_storage_class=self.s3_storage_class,
            tags=self.tags)


def get_location_s3(location_arn: Optional[builtins.str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLocationS3Result:
    """
    Resource schema for AWS::DataSync::LocationS3


    :param builtins.str location_arn: The Amazon Resource Name (ARN) of the Amazon S3 bucket location.
    """
    __args__ = dict()
    __args__['locationArn'] = location_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:datasync:getLocationS3', __args__, opts=opts, typ=GetLocationS3Result).value

    return AwaitableGetLocationS3Result(
        location_arn=pulumi.get(__ret__, 'location_arn'),
        location_uri=pulumi.get(__ret__, 'location_uri'),
        s3_config=pulumi.get(__ret__, 's3_config'),
        s3_storage_class=pulumi.get(__ret__, 's3_storage_class'),
        tags=pulumi.get(__ret__, 'tags'))
def get_location_s3_output(location_arn: Optional[pulumi.Input[builtins.str]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetLocationS3Result]:
    """
    Resource schema for AWS::DataSync::LocationS3


    :param builtins.str location_arn: The Amazon Resource Name (ARN) of the Amazon S3 bucket location.
    """
    __args__ = dict()
    __args__['locationArn'] = location_arn
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:datasync:getLocationS3', __args__, opts=opts, typ=GetLocationS3Result)
    return __ret__.apply(lambda __response__: GetLocationS3Result(
        location_arn=pulumi.get(__response__, 'location_arn'),
        location_uri=pulumi.get(__response__, 'location_uri'),
        s3_config=pulumi.get(__response__, 's3_config'),
        s3_storage_class=pulumi.get(__response__, 's3_storage_class'),
        tags=pulumi.get(__response__, 'tags')))
