// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    /// <summary>
    /// &lt;p&gt;A transform operation that casts a column to a different type.&lt;/p&gt;
    /// </summary>
    public sealed class DataSetCastColumnTypeOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;Column name.&lt;/p&gt;
        /// </summary>
        [Input("columnName", required: true)]
        public Input<string> ColumnName { get; set; } = null!;

        /// <summary>
        /// &lt;p&gt;When casting a column from string to datetime type, you can supply a string in a
        ///             format supported by Amazon QuickSight to denote the source data format.&lt;/p&gt;
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// New column data type.
        /// </summary>
        [Input("newColumnType", required: true)]
        public Input<Pulumi.AwsNative.QuickSight.DataSetColumnDataType> NewColumnType { get; set; } = null!;

        /// <summary>
        /// The sub data type of the new column. Sub types are only available for decimal columns that are part of a SPICE dataset.
        /// </summary>
        [Input("subType")]
        public Input<Pulumi.AwsNative.QuickSight.DataSetColumnDataSubType>? SubType { get; set; }

        public DataSetCastColumnTypeOperationArgs()
        {
        }
        public static new DataSetCastColumnTypeOperationArgs Empty => new DataSetCastColumnTypeOperationArgs();
    }
}
