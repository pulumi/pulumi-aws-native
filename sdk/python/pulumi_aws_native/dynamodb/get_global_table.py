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
from ._enums import *

__all__ = [
    'GetGlobalTableResult',
    'AwaitableGetGlobalTableResult',
    'get_global_table',
    'get_global_table_output',
]

@pulumi.output_type
class GetGlobalTableResult:
    def __init__(__self__, arn=None, attribute_definitions=None, billing_mode=None, global_secondary_indexes=None, global_table_witnesses=None, multi_region_consistency=None, replicas=None, sse_specification=None, stream_arn=None, stream_specification=None, table_id=None, time_to_live_specification=None, warm_throughput=None, write_on_demand_throughput_settings=None, write_provisioned_throughput_settings=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if attribute_definitions and not isinstance(attribute_definitions, list):
            raise TypeError("Expected argument 'attribute_definitions' to be a list")
        pulumi.set(__self__, "attribute_definitions", attribute_definitions)
        if billing_mode and not isinstance(billing_mode, str):
            raise TypeError("Expected argument 'billing_mode' to be a str")
        pulumi.set(__self__, "billing_mode", billing_mode)
        if global_secondary_indexes and not isinstance(global_secondary_indexes, list):
            raise TypeError("Expected argument 'global_secondary_indexes' to be a list")
        pulumi.set(__self__, "global_secondary_indexes", global_secondary_indexes)
        if global_table_witnesses and not isinstance(global_table_witnesses, list):
            raise TypeError("Expected argument 'global_table_witnesses' to be a list")
        pulumi.set(__self__, "global_table_witnesses", global_table_witnesses)
        if multi_region_consistency and not isinstance(multi_region_consistency, str):
            raise TypeError("Expected argument 'multi_region_consistency' to be a str")
        pulumi.set(__self__, "multi_region_consistency", multi_region_consistency)
        if replicas and not isinstance(replicas, list):
            raise TypeError("Expected argument 'replicas' to be a list")
        pulumi.set(__self__, "replicas", replicas)
        if sse_specification and not isinstance(sse_specification, dict):
            raise TypeError("Expected argument 'sse_specification' to be a dict")
        pulumi.set(__self__, "sse_specification", sse_specification)
        if stream_arn and not isinstance(stream_arn, str):
            raise TypeError("Expected argument 'stream_arn' to be a str")
        pulumi.set(__self__, "stream_arn", stream_arn)
        if stream_specification and not isinstance(stream_specification, dict):
            raise TypeError("Expected argument 'stream_specification' to be a dict")
        pulumi.set(__self__, "stream_specification", stream_specification)
        if table_id and not isinstance(table_id, str):
            raise TypeError("Expected argument 'table_id' to be a str")
        pulumi.set(__self__, "table_id", table_id)
        if time_to_live_specification and not isinstance(time_to_live_specification, dict):
            raise TypeError("Expected argument 'time_to_live_specification' to be a dict")
        pulumi.set(__self__, "time_to_live_specification", time_to_live_specification)
        if warm_throughput and not isinstance(warm_throughput, dict):
            raise TypeError("Expected argument 'warm_throughput' to be a dict")
        pulumi.set(__self__, "warm_throughput", warm_throughput)
        if write_on_demand_throughput_settings and not isinstance(write_on_demand_throughput_settings, dict):
            raise TypeError("Expected argument 'write_on_demand_throughput_settings' to be a dict")
        pulumi.set(__self__, "write_on_demand_throughput_settings", write_on_demand_throughput_settings)
        if write_provisioned_throughput_settings and not isinstance(write_provisioned_throughput_settings, dict):
            raise TypeError("Expected argument 'write_provisioned_throughput_settings' to be a dict")
        pulumi.set(__self__, "write_provisioned_throughput_settings", write_provisioned_throughput_settings)

    @property
    @pulumi.getter
    def arn(self) -> Optional[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the DynamoDB table, such as `arn:aws:dynamodb:us-east-2:123456789012:table/myDynamoDBTable` . The ARN returned is that of the replica in the region the stack is deployed to.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="attributeDefinitions")
    def attribute_definitions(self) -> Optional[Sequence['outputs.GlobalTableAttributeDefinition']]:
        """
        A list of attributes that describe the key schema for the global table and indexes.
        """
        return pulumi.get(self, "attribute_definitions")

    @property
    @pulumi.getter(name="billingMode")
    def billing_mode(self) -> Optional[builtins.str]:
        """
        Specifies how you are charged for read and write throughput and how you manage capacity. Valid values are:

        - `PAY_PER_REQUEST`
        - `PROVISIONED`

        All replicas in your global table will have the same billing mode. If you use `PROVISIONED` billing mode, you must provide an auto scaling configuration via the `WriteProvisionedThroughputSettings` property. The default value of this property is `PROVISIONED` .
        """
        return pulumi.get(self, "billing_mode")

    @property
    @pulumi.getter(name="globalSecondaryIndexes")
    def global_secondary_indexes(self) -> Optional[Sequence['outputs.GlobalTableGlobalSecondaryIndex']]:
        """
        Global secondary indexes to be created on the global table. You can create up to 20 global secondary indexes. Each replica in your global table will have the same global secondary index settings. You can only create or delete one global secondary index in a single stack operation.

        Since the backfilling of an index could take a long time, CloudFormation does not wait for the index to become active. If a stack operation rolls back, CloudFormation might not delete an index that has been added. In that case, you will need to delete the index manually.
        """
        return pulumi.get(self, "global_secondary_indexes")

    @property
    @pulumi.getter(name="globalTableWitnesses")
    def global_table_witnesses(self) -> Optional[Sequence['outputs.GlobalTableWitness']]:
        """
        The list of witnesses of the MRSC global table. Only one witness Region can be configured per MRSC global table.
        """
        return pulumi.get(self, "global_table_witnesses")

    @property
    @pulumi.getter(name="multiRegionConsistency")
    def multi_region_consistency(self) -> Optional['GlobalTableMultiRegionConsistency']:
        """
        Specifies the consistency mode for a new global table.

        You can specify one of the following consistency modes:

        - `EVENTUAL` : Configures a new global table for multi-Region eventual consistency (MREC).
        - `STRONG` : Configures a new global table for multi-Region strong consistency (MRSC).

        If you don't specify this field, the global table consistency mode defaults to `EVENTUAL` . For more information about global tables consistency modes, see [Consistency modes](https://docs.aws.amazon.com/V2globaltables_HowItWorks.html#V2globaltables_HowItWorks.consistency-modes) in DynamoDB developer guide.
        """
        return pulumi.get(self, "multi_region_consistency")

    @property
    @pulumi.getter
    def replicas(self) -> Optional[Sequence['outputs.GlobalTableReplicaSpecification']]:
        """
        Specifies the list of replicas for your global table. The list must contain at least one element, the region where the stack defining the global table is deployed. For example, if you define your table in a stack deployed to us-east-1, you must have an entry in `Replicas` with the region us-east-1. You cannot remove the replica in the stack region.

        > Adding a replica might take a few minutes for an empty table, or up to several hours for large tables. If you want to add or remove a replica, we recommend submitting an `UpdateStack` operation containing only that change.
        > 
        > If you add or delete a replica during an update, we recommend that you don't update any other resources. If your stack fails to update and is rolled back while adding a new replica, you might need to manually delete the replica. 

        You can create a new global table with as many replicas as needed. You can add or remove replicas after table creation, but you can only add or remove a single replica in each update. For Multi-Region Strong Consistency (MRSC), you can add or remove up to 3 replicas, or 2 replicas plus a witness Region.
        """
        return pulumi.get(self, "replicas")

    @property
    @pulumi.getter(name="sseSpecification")
    def sse_specification(self) -> Optional['outputs.GlobalTableSseSpecification']:
        """
        Specifies the settings to enable server-side encryption. These settings will be applied to all replicas. If you plan to use customer-managed KMS keys, you must provide a key for each replica using the `ReplicaSpecification.ReplicaSSESpecification` property.
        """
        return pulumi.get(self, "sse_specification")

    @property
    @pulumi.getter(name="streamArn")
    def stream_arn(self) -> Optional[builtins.str]:
        """
        The ARN of the DynamoDB stream, such as `arn:aws:dynamodb:us-east-1:123456789012:table/testddbstack-myDynamoDBTable-012A1SL7SMP5Q/stream/2015-11-30T20:10:00.000` . The `StreamArn` returned is that of the replica in the region the stack is deployed to.

        > You must specify the `StreamSpecification` property to use this attribute.
        """
        return pulumi.get(self, "stream_arn")

    @property
    @pulumi.getter(name="streamSpecification")
    def stream_specification(self) -> Optional['outputs.GlobalTableStreamSpecification']:
        """
        Specifies the streams settings on your global table. You must provide a value for this property if your global table contains more than one replica. You can only change the streams settings if your global table has only one replica. For Multi-Region Strong Consistency (MRSC), you do not need to provide a value for this property and can change the settings at any time.
        """
        return pulumi.get(self, "stream_specification")

    @property
    @pulumi.getter(name="tableId")
    def table_id(self) -> Optional[builtins.str]:
        """
        Unique identifier for the table, such as `a123b456-01ab-23cd-123a-111222aaabbb` . The `TableId` returned is that of the replica in the region the stack is deployed to.
        """
        return pulumi.get(self, "table_id")

    @property
    @pulumi.getter(name="timeToLiveSpecification")
    def time_to_live_specification(self) -> Optional['outputs.GlobalTableTimeToLiveSpecification']:
        """
        Specifies the time to live (TTL) settings for the table. This setting will be applied to all replicas.
        """
        return pulumi.get(self, "time_to_live_specification")

    @property
    @pulumi.getter(name="warmThroughput")
    def warm_throughput(self) -> Optional['outputs.GlobalTableWarmThroughput']:
        """
        Provides visibility into the number of read and write operations your table or secondary index can instantaneously support. The settings can be modified using the `UpdateTable` operation to meet the throughput requirements of an upcoming peak event.
        """
        return pulumi.get(self, "warm_throughput")

    @property
    @pulumi.getter(name="writeOnDemandThroughputSettings")
    def write_on_demand_throughput_settings(self) -> Optional['outputs.GlobalTableWriteOnDemandThroughputSettings']:
        """
        Sets the write request settings for a global table or a global secondary index. You can only specify this setting if your resource uses the `PAY_PER_REQUEST` `BillingMode` .
        """
        return pulumi.get(self, "write_on_demand_throughput_settings")

    @property
    @pulumi.getter(name="writeProvisionedThroughputSettings")
    def write_provisioned_throughput_settings(self) -> Optional['outputs.GlobalTableWriteProvisionedThroughputSettings']:
        """
        Specifies an auto scaling policy for write capacity. This policy will be applied to all replicas. This setting must be specified if `BillingMode` is set to `PROVISIONED` .
        """
        return pulumi.get(self, "write_provisioned_throughput_settings")


class AwaitableGetGlobalTableResult(GetGlobalTableResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGlobalTableResult(
            arn=self.arn,
            attribute_definitions=self.attribute_definitions,
            billing_mode=self.billing_mode,
            global_secondary_indexes=self.global_secondary_indexes,
            global_table_witnesses=self.global_table_witnesses,
            multi_region_consistency=self.multi_region_consistency,
            replicas=self.replicas,
            sse_specification=self.sse_specification,
            stream_arn=self.stream_arn,
            stream_specification=self.stream_specification,
            table_id=self.table_id,
            time_to_live_specification=self.time_to_live_specification,
            warm_throughput=self.warm_throughput,
            write_on_demand_throughput_settings=self.write_on_demand_throughput_settings,
            write_provisioned_throughput_settings=self.write_provisioned_throughput_settings)


def get_global_table(table_name: Optional[builtins.str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGlobalTableResult:
    """
    Version: None. Resource Type definition for AWS::DynamoDB::GlobalTable


    :param builtins.str table_name: A name for the global table. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID as the table name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
           
           > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
    """
    __args__ = dict()
    __args__['tableName'] = table_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:dynamodb:getGlobalTable', __args__, opts=opts, typ=GetGlobalTableResult).value

    return AwaitableGetGlobalTableResult(
        arn=pulumi.get(__ret__, 'arn'),
        attribute_definitions=pulumi.get(__ret__, 'attribute_definitions'),
        billing_mode=pulumi.get(__ret__, 'billing_mode'),
        global_secondary_indexes=pulumi.get(__ret__, 'global_secondary_indexes'),
        global_table_witnesses=pulumi.get(__ret__, 'global_table_witnesses'),
        multi_region_consistency=pulumi.get(__ret__, 'multi_region_consistency'),
        replicas=pulumi.get(__ret__, 'replicas'),
        sse_specification=pulumi.get(__ret__, 'sse_specification'),
        stream_arn=pulumi.get(__ret__, 'stream_arn'),
        stream_specification=pulumi.get(__ret__, 'stream_specification'),
        table_id=pulumi.get(__ret__, 'table_id'),
        time_to_live_specification=pulumi.get(__ret__, 'time_to_live_specification'),
        warm_throughput=pulumi.get(__ret__, 'warm_throughput'),
        write_on_demand_throughput_settings=pulumi.get(__ret__, 'write_on_demand_throughput_settings'),
        write_provisioned_throughput_settings=pulumi.get(__ret__, 'write_provisioned_throughput_settings'))
def get_global_table_output(table_name: Optional[pulumi.Input[builtins.str]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetGlobalTableResult]:
    """
    Version: None. Resource Type definition for AWS::DynamoDB::GlobalTable


    :param builtins.str table_name: A name for the global table. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID as the table name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
           
           > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
    """
    __args__ = dict()
    __args__['tableName'] = table_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:dynamodb:getGlobalTable', __args__, opts=opts, typ=GetGlobalTableResult)
    return __ret__.apply(lambda __response__: GetGlobalTableResult(
        arn=pulumi.get(__response__, 'arn'),
        attribute_definitions=pulumi.get(__response__, 'attribute_definitions'),
        billing_mode=pulumi.get(__response__, 'billing_mode'),
        global_secondary_indexes=pulumi.get(__response__, 'global_secondary_indexes'),
        global_table_witnesses=pulumi.get(__response__, 'global_table_witnesses'),
        multi_region_consistency=pulumi.get(__response__, 'multi_region_consistency'),
        replicas=pulumi.get(__response__, 'replicas'),
        sse_specification=pulumi.get(__response__, 'sse_specification'),
        stream_arn=pulumi.get(__response__, 'stream_arn'),
        stream_specification=pulumi.get(__response__, 'stream_specification'),
        table_id=pulumi.get(__response__, 'table_id'),
        time_to_live_specification=pulumi.get(__response__, 'time_to_live_specification'),
        warm_throughput=pulumi.get(__response__, 'warm_throughput'),
        write_on_demand_throughput_settings=pulumi.get(__response__, 'write_on_demand_throughput_settings'),
        write_provisioned_throughput_settings=pulumi.get(__response__, 'write_provisioned_throughput_settings')))
