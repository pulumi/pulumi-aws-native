// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cassandra.Inputs
{

    public sealed class TableClusteringKeyColumnArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name and data type of this clustering key column.
        /// </summary>
        [Input("column", required: true)]
        public Input<Inputs.TableColumnArgs> Column { get; set; } = null!;

        /// <summary>
        /// The order in which this column's data is stored:
        /// 
        /// - `ASC` (default) - The column's data is stored in ascending order.
        /// - `DESC` - The column's data is stored in descending order.
        /// </summary>
        [Input("orderBy")]
        public Input<Pulumi.AwsNative.Cassandra.TableClusteringKeyColumnOrderBy>? OrderBy { get; set; }

        public TableClusteringKeyColumnArgs()
        {
        }
        public static new TableClusteringKeyColumnArgs Empty => new TableClusteringKeyColumnArgs();
    }
}
