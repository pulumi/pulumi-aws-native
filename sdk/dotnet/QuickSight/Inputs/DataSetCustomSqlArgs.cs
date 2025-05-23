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
    /// &lt;p&gt;A physical table type built from the results of the custom SQL query.&lt;/p&gt;
    /// </summary>
    public sealed class DataSetCustomSqlArgs : global::Pulumi.ResourceArgs
    {
        [Input("columns")]
        private InputList<Inputs.DataSetInputColumnArgs>? _columns;

        /// <summary>
        /// &lt;p&gt;The column schema from the SQL query result set.&lt;/p&gt;
        /// </summary>
        public InputList<Inputs.DataSetInputColumnArgs> Columns
        {
            get => _columns ?? (_columns = new InputList<Inputs.DataSetInputColumnArgs>());
            set => _columns = value;
        }

        /// <summary>
        /// &lt;p&gt;The Amazon Resource Name (ARN) of the data source.&lt;/p&gt;
        /// </summary>
        [Input("dataSourceArn", required: true)]
        public Input<string> DataSourceArn { get; set; } = null!;

        /// <summary>
        /// &lt;p&gt;A display name for the SQL query result.&lt;/p&gt;
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// &lt;p&gt;The SQL query.&lt;/p&gt;
        /// </summary>
        [Input("sqlQuery", required: true)]
        public Input<string> SqlQuery { get; set; } = null!;

        public DataSetCustomSqlArgs()
        {
        }
        public static new DataSetCustomSqlArgs Empty => new DataSetCustomSqlArgs();
    }
}
