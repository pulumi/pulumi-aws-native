# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
from ._enums import *

__all__ = [
    'GetDirectoryBucketResult',
    'AwaitableGetDirectoryBucketResult',
    'get_directory_bucket',
    'get_directory_bucket_output',
]

@pulumi.output_type
class GetDirectoryBucketResult:
    def __init__(__self__, arn=None, availability_zone_name=None, bucket_encryption=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if availability_zone_name and not isinstance(availability_zone_name, str):
            raise TypeError("Expected argument 'availability_zone_name' to be a str")
        pulumi.set(__self__, "availability_zone_name", availability_zone_name)
        if bucket_encryption and not isinstance(bucket_encryption, dict):
            raise TypeError("Expected argument 'bucket_encryption' to be a dict")
        pulumi.set(__self__, "bucket_encryption", bucket_encryption)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        Returns the Amazon Resource Name (ARN) of the specified bucket.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="availabilityZoneName")
    def availability_zone_name(self) -> Optional[str]:
        """
        Returns the code for the Availability Zone where the directory bucket was created.
        """
        return pulumi.get(self, "availability_zone_name")

    @property
    @pulumi.getter(name="bucketEncryption")
    def bucket_encryption(self) -> Optional['outputs.DirectoryBucketBucketEncryption']:
        """
        Specifies default encryption for a bucket using server-side encryption with Amazon S3 managed keys (SSE-S3) or AWS KMS keys (SSE-KMS). For information about default encryption for directory buckets, see [Setting and monitoring default encryption for directory buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-bucket-encryption.html) in the *Amazon S3 User Guide* .
        """
        return pulumi.get(self, "bucket_encryption")


class AwaitableGetDirectoryBucketResult(GetDirectoryBucketResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDirectoryBucketResult(
            arn=self.arn,
            availability_zone_name=self.availability_zone_name,
            bucket_encryption=self.bucket_encryption)


def get_directory_bucket(bucket_name: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDirectoryBucketResult:
    """
    Resource Type definition for AWS::S3Express::DirectoryBucket.


    :param str bucket_name: Specifies a name for the bucket. The bucket name must contain only lowercase letters, numbers, and hyphens (-). A directory bucket name must be unique in the chosen Availability Zone. The bucket name must also follow the format 'bucket_base_name--az_id--x-s3' (for example, 'DOC-EXAMPLE-BUCKET--usw2-az1--x-s3'). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
    """
    __args__ = dict()
    __args__['bucketName'] = bucket_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:s3express:getDirectoryBucket', __args__, opts=opts, typ=GetDirectoryBucketResult).value

    return AwaitableGetDirectoryBucketResult(
        arn=pulumi.get(__ret__, 'arn'),
        availability_zone_name=pulumi.get(__ret__, 'availability_zone_name'),
        bucket_encryption=pulumi.get(__ret__, 'bucket_encryption'))
def get_directory_bucket_output(bucket_name: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDirectoryBucketResult]:
    """
    Resource Type definition for AWS::S3Express::DirectoryBucket.


    :param str bucket_name: Specifies a name for the bucket. The bucket name must contain only lowercase letters, numbers, and hyphens (-). A directory bucket name must be unique in the chosen Availability Zone. The bucket name must also follow the format 'bucket_base_name--az_id--x-s3' (for example, 'DOC-EXAMPLE-BUCKET--usw2-az1--x-s3'). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
    """
    __args__ = dict()
    __args__['bucketName'] = bucket_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:s3express:getDirectoryBucket', __args__, opts=opts, typ=GetDirectoryBucketResult)
    return __ret__.apply(lambda __response__: GetDirectoryBucketResult(
        arn=pulumi.get(__response__, 'arn'),
        availability_zone_name=pulumi.get(__response__, 'availability_zone_name'),
        bucket_encryption=pulumi.get(__response__, 'bucket_encryption')))
