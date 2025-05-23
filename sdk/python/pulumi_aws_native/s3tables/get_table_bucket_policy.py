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

__all__ = [
    'GetTableBucketPolicyResult',
    'AwaitableGetTableBucketPolicyResult',
    'get_table_bucket_policy',
    'get_table_bucket_policy_output',
]

@pulumi.output_type
class GetTableBucketPolicyResult:
    def __init__(__self__, resource_policy=None):
        if resource_policy and not isinstance(resource_policy, dict):
            raise TypeError("Expected argument 'resource_policy' to be a dict")
        pulumi.set(__self__, "resource_policy", resource_policy)

    @property
    @pulumi.getter(name="resourcePolicy")
    def resource_policy(self) -> Optional['outputs.TableBucketPolicyResourcePolicy']:
        """
        The bucket policy JSON for the table bucket.
        """
        return pulumi.get(self, "resource_policy")


class AwaitableGetTableBucketPolicyResult(GetTableBucketPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTableBucketPolicyResult(
            resource_policy=self.resource_policy)


def get_table_bucket_policy(table_bucket_arn: Optional[builtins.str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTableBucketPolicyResult:
    """
    Applies an IAM resource policy to a table bucket.


    :param builtins.str table_bucket_arn: The Amazon Resource Name (ARN) of the table bucket.
    """
    __args__ = dict()
    __args__['tableBucketArn'] = table_bucket_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:s3tables:getTableBucketPolicy', __args__, opts=opts, typ=GetTableBucketPolicyResult).value

    return AwaitableGetTableBucketPolicyResult(
        resource_policy=pulumi.get(__ret__, 'resource_policy'))
def get_table_bucket_policy_output(table_bucket_arn: Optional[pulumi.Input[builtins.str]] = None,
                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetTableBucketPolicyResult]:
    """
    Applies an IAM resource policy to a table bucket.


    :param builtins.str table_bucket_arn: The Amazon Resource Name (ARN) of the table bucket.
    """
    __args__ = dict()
    __args__['tableBucketArn'] = table_bucket_arn
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:s3tables:getTableBucketPolicy', __args__, opts=opts, typ=GetTableBucketPolicyResult)
    return __ret__.apply(lambda __response__: GetTableBucketPolicyResult(
        resource_policy=pulumi.get(__response__, 'resource_policy')))
