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
    /// Represents *a single element* of a key schema. A key schema specifies the attributes that make up the primary key of a table, or the key attributes of an index.
    ///  A ``KeySchemaElement`` represents exactly one attribute of the primary key. For example, a simple primary key would be represented by one ``KeySchemaElement`` (for the partition key). A composite primary key would require one ``KeySchemaElement`` for the partition key, and another ``KeySchemaElement`` for the sort key.
    ///  A ``KeySchemaElement`` must be a scalar, top-level attribute (not a nested attribute). The data type must be one of String, Number, or Binary. The attribute cannot be nested within a List or a Map.
    /// </summary>
    public sealed class TableKeySchemaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of a key attribute.
        /// </summary>
        [Input("attributeName", required: true)]
        public Input<string> AttributeName { get; set; } = null!;

        /// <summary>
        /// The role that this key attribute will assume:
        ///   +  ``HASH`` - partition key
        ///   +  ``RANGE`` - sort key
        ///   
        ///   The partition key of an item is also known as its *hash attribute*. The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.
        ///  The sort key of an item is also known as its *range attribute*. The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.
        /// </summary>
        [Input("keyType", required: true)]
        public Input<string> KeyType { get; set; } = null!;

        public TableKeySchemaArgs()
        {
        }
        public static new TableKeySchemaArgs Empty => new TableKeySchemaArgs();
    }
}
