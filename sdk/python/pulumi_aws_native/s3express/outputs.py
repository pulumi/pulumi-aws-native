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
    'DirectoryBucketBucketEncryption',
    'DirectoryBucketServerSideEncryptionByDefault',
    'DirectoryBucketServerSideEncryptionRule',
]

@pulumi.output_type
class DirectoryBucketBucketEncryption(dict):
    """
    Specifies default encryption for a bucket using server-side encryption with Amazon S3 managed keys (SSE-S3) or AWS KMS keys (SSE-KMS).
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "serverSideEncryptionConfiguration":
            suggest = "server_side_encryption_configuration"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DirectoryBucketBucketEncryption. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DirectoryBucketBucketEncryption.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DirectoryBucketBucketEncryption.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 server_side_encryption_configuration: Sequence['outputs.DirectoryBucketServerSideEncryptionRule']):
        """
        Specifies default encryption for a bucket using server-side encryption with Amazon S3 managed keys (SSE-S3) or AWS KMS keys (SSE-KMS).
        :param Sequence['DirectoryBucketServerSideEncryptionRule'] server_side_encryption_configuration: Specifies the default server-side-encryption configuration.
        """
        pulumi.set(__self__, "server_side_encryption_configuration", server_side_encryption_configuration)

    @property
    @pulumi.getter(name="serverSideEncryptionConfiguration")
    def server_side_encryption_configuration(self) -> Sequence['outputs.DirectoryBucketServerSideEncryptionRule']:
        """
        Specifies the default server-side-encryption configuration.
        """
        return pulumi.get(self, "server_side_encryption_configuration")


@pulumi.output_type
class DirectoryBucketServerSideEncryptionByDefault(dict):
    """
    Specifies the default server-side encryption to apply to new objects in the bucket. If a PUT Object request doesn't specify any server-side encryption, this default encryption will be applied.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "sseAlgorithm":
            suggest = "sse_algorithm"
        elif key == "kmsMasterKeyId":
            suggest = "kms_master_key_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DirectoryBucketServerSideEncryptionByDefault. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DirectoryBucketServerSideEncryptionByDefault.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DirectoryBucketServerSideEncryptionByDefault.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 sse_algorithm: 'DirectoryBucketServerSideEncryptionByDefaultSseAlgorithm',
                 kms_master_key_id: Optional[str] = None):
        """
        Specifies the default server-side encryption to apply to new objects in the bucket. If a PUT Object request doesn't specify any server-side encryption, this default encryption will be applied.
        :param 'DirectoryBucketServerSideEncryptionByDefaultSseAlgorithm' sse_algorithm: Server-side encryption algorithm to use for the default encryption.
               
               > For directory buckets, there are only two supported values for server-side encryption: `AES256` and `aws:kms` .
        :param str kms_master_key_id: AWS Key Management Service (KMS) customer managed key ID to use for the default encryption. This parameter is allowed only if SSEAlgorithm is set to aws:kms. You can specify this parameter with the key ID or the Amazon Resource Name (ARN) of the KMS key
        """
        pulumi.set(__self__, "sse_algorithm", sse_algorithm)
        if kms_master_key_id is not None:
            pulumi.set(__self__, "kms_master_key_id", kms_master_key_id)

    @property
    @pulumi.getter(name="sseAlgorithm")
    def sse_algorithm(self) -> 'DirectoryBucketServerSideEncryptionByDefaultSseAlgorithm':
        """
        Server-side encryption algorithm to use for the default encryption.

        > For directory buckets, there are only two supported values for server-side encryption: `AES256` and `aws:kms` .
        """
        return pulumi.get(self, "sse_algorithm")

    @property
    @pulumi.getter(name="kmsMasterKeyId")
    def kms_master_key_id(self) -> Optional[str]:
        """
        AWS Key Management Service (KMS) customer managed key ID to use for the default encryption. This parameter is allowed only if SSEAlgorithm is set to aws:kms. You can specify this parameter with the key ID or the Amazon Resource Name (ARN) of the KMS key
        """
        return pulumi.get(self, "kms_master_key_id")


@pulumi.output_type
class DirectoryBucketServerSideEncryptionRule(dict):
    """
    Specifies the default server-side encryption configuration.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "bucketKeyEnabled":
            suggest = "bucket_key_enabled"
        elif key == "serverSideEncryptionByDefault":
            suggest = "server_side_encryption_by_default"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DirectoryBucketServerSideEncryptionRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DirectoryBucketServerSideEncryptionRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DirectoryBucketServerSideEncryptionRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 bucket_key_enabled: Optional[bool] = None,
                 server_side_encryption_by_default: Optional['outputs.DirectoryBucketServerSideEncryptionByDefault'] = None):
        """
        Specifies the default server-side encryption configuration.
        :param bool bucket_key_enabled: Specifies whether Amazon S3 should use an S3 Bucket Key with server-side encryption using KMS (SSE-KMS) for new objects in the bucket. Existing objects are not affected. Amazon S3 Express One Zone uses an S3 Bucket Key with SSE-KMS and S3 Bucket Key cannot be disabled. It's only allowed to set the BucketKeyEnabled element to true.
        :param 'DirectoryBucketServerSideEncryptionByDefault' server_side_encryption_by_default: Specifies the default server-side encryption to apply to new objects in the bucket. If a PUT Object request doesn't specify any server-side encryption, this default encryption will be applied.
        """
        if bucket_key_enabled is not None:
            pulumi.set(__self__, "bucket_key_enabled", bucket_key_enabled)
        if server_side_encryption_by_default is not None:
            pulumi.set(__self__, "server_side_encryption_by_default", server_side_encryption_by_default)

    @property
    @pulumi.getter(name="bucketKeyEnabled")
    def bucket_key_enabled(self) -> Optional[bool]:
        """
        Specifies whether Amazon S3 should use an S3 Bucket Key with server-side encryption using KMS (SSE-KMS) for new objects in the bucket. Existing objects are not affected. Amazon S3 Express One Zone uses an S3 Bucket Key with SSE-KMS and S3 Bucket Key cannot be disabled. It's only allowed to set the BucketKeyEnabled element to true.
        """
        return pulumi.get(self, "bucket_key_enabled")

    @property
    @pulumi.getter(name="serverSideEncryptionByDefault")
    def server_side_encryption_by_default(self) -> Optional['outputs.DirectoryBucketServerSideEncryptionByDefault']:
        """
        Specifies the default server-side encryption to apply to new objects in the bucket. If a PUT Object request doesn't specify any server-side encryption, this default encryption will be applied.
        """
        return pulumi.get(self, "server_side_encryption_by_default")

