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
    'GetKeyspaceResult',
    'AwaitableGetKeyspaceResult',
    'get_keyspace',
    'get_keyspace_output',
]

@pulumi.output_type
class GetKeyspaceResult:
    def __init__(__self__, client_side_timestamps_enabled=None, replication_specification=None, tags=None):
        if client_side_timestamps_enabled and not isinstance(client_side_timestamps_enabled, bool):
            raise TypeError("Expected argument 'client_side_timestamps_enabled' to be a bool")
        pulumi.set(__self__, "client_side_timestamps_enabled", client_side_timestamps_enabled)
        if replication_specification and not isinstance(replication_specification, dict):
            raise TypeError("Expected argument 'replication_specification' to be a dict")
        pulumi.set(__self__, "replication_specification", replication_specification)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="clientSideTimestampsEnabled")
    def client_side_timestamps_enabled(self) -> Optional[builtins.bool]:
        """
        Indicates whether client-side timestamps are enabled (true) or disabled (false) for all tables in the keyspace. To add a Region to a single-Region keyspace with at least one table, the value must be set to true. After you enabled client-side timestamps for a table, you can’t disable it again.
        """
        return pulumi.get(self, "client_side_timestamps_enabled")

    @property
    @pulumi.getter(name="replicationSpecification")
    def replication_specification(self) -> Optional['outputs.KeyspaceReplicationSpecification']:
        """
        Specifies the `ReplicationStrategy` of a keyspace. The options are:

        - `SINGLE_REGION` for a single Region keyspace (optional) or
        - `MULTI_REGION` for a multi-Region keyspace

        If no `ReplicationStrategy` is provided, the default is `SINGLE_REGION` . If you choose `MULTI_REGION` , you must also provide a `RegionList` with the AWS Regions that the keyspace is replicated in.
        """
        return pulumi.get(self, "replication_specification")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        An array of key-value pairs to apply to this resource.

        For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
        """
        return pulumi.get(self, "tags")


class AwaitableGetKeyspaceResult(GetKeyspaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKeyspaceResult(
            client_side_timestamps_enabled=self.client_side_timestamps_enabled,
            replication_specification=self.replication_specification,
            tags=self.tags)


def get_keyspace(keyspace_name: Optional[builtins.str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKeyspaceResult:
    """
    Resource schema for AWS::Cassandra::Keyspace


    :param builtins.str keyspace_name: Name for Cassandra keyspace
    """
    __args__ = dict()
    __args__['keyspaceName'] = keyspace_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:cassandra:getKeyspace', __args__, opts=opts, typ=GetKeyspaceResult).value

    return AwaitableGetKeyspaceResult(
        client_side_timestamps_enabled=pulumi.get(__ret__, 'client_side_timestamps_enabled'),
        replication_specification=pulumi.get(__ret__, 'replication_specification'),
        tags=pulumi.get(__ret__, 'tags'))
def get_keyspace_output(keyspace_name: Optional[pulumi.Input[builtins.str]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetKeyspaceResult]:
    """
    Resource schema for AWS::Cassandra::Keyspace


    :param builtins.str keyspace_name: Name for Cassandra keyspace
    """
    __args__ = dict()
    __args__['keyspaceName'] = keyspace_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:cassandra:getKeyspace', __args__, opts=opts, typ=GetKeyspaceResult)
    return __ret__.apply(lambda __response__: GetKeyspaceResult(
        client_side_timestamps_enabled=pulumi.get(__response__, 'client_side_timestamps_enabled'),
        replication_specification=pulumi.get(__response__, 'replication_specification'),
        tags=pulumi.get(__response__, 'tags')))
