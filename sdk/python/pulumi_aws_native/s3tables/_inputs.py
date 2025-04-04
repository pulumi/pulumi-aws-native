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
from ._enums import *

__all__ = [
    'TableBucketPolicyResourcePolicyArgs',
    'TableBucketPolicyResourcePolicyArgsDict',
    'TableBucketUnreferencedFileRemovalArgs',
    'TableBucketUnreferencedFileRemovalArgsDict',
]

MYPY = False

if not MYPY:
    class TableBucketPolicyResourcePolicyArgsDict(TypedDict):
        """
        A policy document containing permissions to add to the specified table bucket. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM.
        """
        pass
elif False:
    TableBucketPolicyResourcePolicyArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class TableBucketPolicyResourcePolicyArgs:
    def __init__(__self__):
        """
        A policy document containing permissions to add to the specified table bucket. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM.
        """
        pass


if not MYPY:
    class TableBucketUnreferencedFileRemovalArgsDict(TypedDict):
        """
        Settings governing the Unreferenced File Removal maintenance action. Unreferenced file removal identifies and deletes all objects that are not referenced by any table snapshots.
        """
        noncurrent_days: NotRequired[pulumi.Input[int]]
        """
        S3 permanently deletes noncurrent objects after the number of days specified by the NoncurrentDays property.
        """
        status: NotRequired[pulumi.Input['TableBucketUnreferencedFileRemovalStatus']]
        """
        Indicates whether the Unreferenced File Removal maintenance action is enabled.
        """
        unreferenced_days: NotRequired[pulumi.Input[int]]
        """
        For any object not referenced by your table and older than the UnreferencedDays property, S3 creates a delete marker and marks the object version as noncurrent.
        """
elif False:
    TableBucketUnreferencedFileRemovalArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class TableBucketUnreferencedFileRemovalArgs:
    def __init__(__self__, *,
                 noncurrent_days: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input['TableBucketUnreferencedFileRemovalStatus']] = None,
                 unreferenced_days: Optional[pulumi.Input[int]] = None):
        """
        Settings governing the Unreferenced File Removal maintenance action. Unreferenced file removal identifies and deletes all objects that are not referenced by any table snapshots.
        :param pulumi.Input[int] noncurrent_days: S3 permanently deletes noncurrent objects after the number of days specified by the NoncurrentDays property.
        :param pulumi.Input['TableBucketUnreferencedFileRemovalStatus'] status: Indicates whether the Unreferenced File Removal maintenance action is enabled.
        :param pulumi.Input[int] unreferenced_days: For any object not referenced by your table and older than the UnreferencedDays property, S3 creates a delete marker and marks the object version as noncurrent.
        """
        if noncurrent_days is not None:
            pulumi.set(__self__, "noncurrent_days", noncurrent_days)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if unreferenced_days is not None:
            pulumi.set(__self__, "unreferenced_days", unreferenced_days)

    @property
    @pulumi.getter(name="noncurrentDays")
    def noncurrent_days(self) -> Optional[pulumi.Input[int]]:
        """
        S3 permanently deletes noncurrent objects after the number of days specified by the NoncurrentDays property.
        """
        return pulumi.get(self, "noncurrent_days")

    @noncurrent_days.setter
    def noncurrent_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "noncurrent_days", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['TableBucketUnreferencedFileRemovalStatus']]:
        """
        Indicates whether the Unreferenced File Removal maintenance action is enabled.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['TableBucketUnreferencedFileRemovalStatus']]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="unreferencedDays")
    def unreferenced_days(self) -> Optional[pulumi.Input[int]]:
        """
        For any object not referenced by your table and older than the UnreferencedDays property, S3 creates a delete marker and marks the object version as noncurrent.
        """
        return pulumi.get(self, "unreferenced_days")

    @unreferenced_days.setter
    def unreferenced_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "unreferenced_days", value)


