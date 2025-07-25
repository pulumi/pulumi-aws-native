// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DynamoDb.Inputs
{

    /// <summary>
    /// Represents the properties of a global secondary index.
    /// </summary>
    public sealed class TableGlobalSecondaryIndexArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The settings used to enable or disable CloudWatch Contributor Insights for the specified global secondary index.
        /// </summary>
        [Input("contributorInsightsSpecification")]
        public Input<Inputs.TableContributorInsightsSpecificationArgs>? ContributorInsightsSpecification { get; set; }

        /// <summary>
        /// The name of the global secondary index. The name must be unique among all other indexes on this table.
        /// </summary>
        [Input("indexName", required: true)]
        public Input<string> IndexName { get; set; } = null!;

        [Input("keySchema", required: true)]
        private InputList<Inputs.TableKeySchemaArgs>? _keySchema;

        /// <summary>
        /// The complete key schema for a global secondary index, which consists of one or more pairs of attribute names and key types:
        ///   +  ``HASH`` - partition key
        ///   +  ``RANGE`` - sort key
        ///   
        ///   The partition key of an item is also known as its *hash attribute*. The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.
        ///  The sort key of an item is also known as its *range attribute*. The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.
        /// </summary>
        public InputList<Inputs.TableKeySchemaArgs> KeySchema
        {
            get => _keySchema ?? (_keySchema = new InputList<Inputs.TableKeySchemaArgs>());
            set => _keySchema = value;
        }

        /// <summary>
        /// The maximum number of read and write units for the specified global secondary index. If you use this parameter, you must specify ``MaxReadRequestUnits``, ``MaxWriteRequestUnits``, or both. You must use either ``OnDemandThroughput`` or ``ProvisionedThroughput`` based on your table's capacity mode.
        /// </summary>
        [Input("onDemandThroughput")]
        public Input<Inputs.TableOnDemandThroughputArgs>? OnDemandThroughput { get; set; }

        /// <summary>
        /// Represents attributes that are copied (projected) from the table into the global secondary index. These are in addition to the primary key attributes and index key attributes, which are automatically projected.
        /// </summary>
        [Input("projection", required: true)]
        public Input<Inputs.TableProjectionArgs> Projection { get; set; } = null!;

        /// <summary>
        /// Represents the provisioned throughput settings for the specified global secondary index. You must use either ``OnDemandThroughput`` or ``ProvisionedThroughput`` based on your table's capacity mode.
        ///  For current minimum and maximum provisioned throughput values, see [Service, Account, and Table Quotas](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the *Amazon DynamoDB Developer Guide*.
        /// </summary>
        [Input("provisionedThroughput")]
        public Input<Inputs.TableProvisionedThroughputArgs>? ProvisionedThroughput { get; set; }

        /// <summary>
        /// Represents the warm throughput value (in read units per second and write units per second) for the specified secondary index. If you use this parameter, you must specify ``ReadUnitsPerSecond``, ``WriteUnitsPerSecond``, or both.
        /// </summary>
        [Input("warmThroughput")]
        public Input<Inputs.TableWarmThroughputArgs>? WarmThroughput { get; set; }

        public TableGlobalSecondaryIndexArgs()
        {
        }
        public static new TableGlobalSecondaryIndexArgs Empty => new TableGlobalSecondaryIndexArgs();
    }
}
