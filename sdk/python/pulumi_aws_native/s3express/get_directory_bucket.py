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
    'GetDirectoryBucketResult',
    'AwaitableGetDirectoryBucketResult',
    'get_directory_bucket',
    'get_directory_bucket_output',
]

@pulumi.output_type
class GetDirectoryBucketResult:
    def __init__(__self__, arn=None, availability_zone_name=None, bucket_encryption=None, lifecycle_configuration=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if availability_zone_name and not isinstance(availability_zone_name, str):
            raise TypeError("Expected argument 'availability_zone_name' to be a str")
        pulumi.set(__self__, "availability_zone_name", availability_zone_name)
        if bucket_encryption and not isinstance(bucket_encryption, dict):
            raise TypeError("Expected argument 'bucket_encryption' to be a dict")
        pulumi.set(__self__, "bucket_encryption", bucket_encryption)
        if lifecycle_configuration and not isinstance(lifecycle_configuration, dict):
            raise TypeError("Expected argument 'lifecycle_configuration' to be a dict")
        pulumi.set(__self__, "lifecycle_configuration", lifecycle_configuration)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[builtins.str]:
        """
        Returns the Amazon Resource Name (ARN) of the specified bucket.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="availabilityZoneName")
    def availability_zone_name(self) -> Optional[builtins.str]:
        """
        Returns the code for the Availability Zone or Local Zone where the directory bucket was created. An example for the code of an Availability Zone is 'us-east-1f'.
        """
        return pulumi.get(self, "availability_zone_name")

    @property
    @pulumi.getter(name="bucketEncryption")
    def bucket_encryption(self) -> Optional['outputs.DirectoryBucketBucketEncryption']:
        """
        Specifies default encryption for a bucket using server-side encryption with Amazon S3 managed keys (SSE-S3) or AWS KMS keys (SSE-KMS). For information about default encryption for directory buckets, see [Setting and monitoring default encryption for directory buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-bucket-encryption.html) in the *Amazon S3 User Guide* .
        """
        return pulumi.get(self, "bucket_encryption")

    @property
    @pulumi.getter(name="lifecycleConfiguration")
    def lifecycle_configuration(self) -> Optional['outputs.DirectoryBucketLifecycleConfiguration']:
        """
        Lifecycle rules that define how Amazon S3 Express manages objects during their lifetime.
        """
        return pulumi.get(self, "lifecycle_configuration")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        return pulumi.get(self, "tags")


class AwaitableGetDirectoryBucketResult(GetDirectoryBucketResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDirectoryBucketResult(
            arn=self.arn,
            availability_zone_name=self.availability_zone_name,
            bucket_encryption=self.bucket_encryption,
            lifecycle_configuration=self.lifecycle_configuration,
            tags=self.tags)


def get_directory_bucket(bucket_name: Optional[builtins.str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDirectoryBucketResult:
    """
    Resource Type definition for AWS::S3Express::DirectoryBucket.


    :param builtins.str bucket_name: Specifies a name for the bucket. The bucket name must contain only lowercase letters, numbers, and hyphens (-). A directory bucket name must be unique in the chosen Availability Zone or Local Zone. The bucket name must also follow the format 'bucket_base_name--zone_id--x-s3'. The zone_id can be the ID of an Availability Zone or a Local Zone. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
    """
    __args__ = dict()
    __args__['bucketName'] = bucket_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:s3express:getDirectoryBucket', __args__, opts=opts, typ=GetDirectoryBucketResult).value

    return AwaitableGetDirectoryBucketResult(
        arn=pulumi.get(__ret__, 'arn'),
        availability_zone_name=pulumi.get(__ret__, 'availability_zone_name'),
        bucket_encryption=pulumi.get(__ret__, 'bucket_encryption'),
        lifecycle_configuration=pulumi.get(__ret__, 'lifecycle_configuration'),
        tags=pulumi.get(__ret__, 'tags'))
def get_directory_bucket_output(bucket_name: Optional[pulumi.Input[builtins.str]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDirectoryBucketResult]:
    """
    Resource Type definition for AWS::S3Express::DirectoryBucket.


    :param builtins.str bucket_name: Specifies a name for the bucket. The bucket name must contain only lowercase letters, numbers, and hyphens (-). A directory bucket name must be unique in the chosen Availability Zone or Local Zone. The bucket name must also follow the format 'bucket_base_name--zone_id--x-s3'. The zone_id can be the ID of an Availability Zone or a Local Zone. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
    """
    __args__ = dict()
    __args__['bucketName'] = bucket_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:s3express:getDirectoryBucket', __args__, opts=opts, typ=GetDirectoryBucketResult)
    return __ret__.apply(lambda __response__: GetDirectoryBucketResult(
        arn=pulumi.get(__response__, 'arn'),
        availability_zone_name=pulumi.get(__response__, 'availability_zone_name'),
        bucket_encryption=pulumi.get(__response__, 'bucket_encryption'),
        lifecycle_configuration=pulumi.get(__response__, 'lifecycle_configuration'),
        tags=pulumi.get(__response__, 'tags')))
