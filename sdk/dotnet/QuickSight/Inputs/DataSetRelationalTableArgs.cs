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
    /// &lt;p&gt;A physical table type for relational data sources.&lt;/p&gt;
    /// </summary>
    public sealed class DataSetRelationalTableArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The catalog associated with a table.&lt;/p&gt;
        /// </summary>
        [Input("catalog")]
        public Input<string>? Catalog { get; set; }

        /// <summary>
        /// &lt;p&gt;The Amazon Resource Name (ARN) for the data source.&lt;/p&gt;
        /// </summary>
        [Input("dataSourceArn", required: true)]
        public Input<string> DataSourceArn { get; set; } = null!;

        [Input("inputColumns")]
        private InputList<Inputs.DataSetInputColumnArgs>? _inputColumns;

        /// <summary>
        /// &lt;p&gt;The column schema of the table.&lt;/p&gt;
        /// </summary>
        public InputList<Inputs.DataSetInputColumnArgs> InputColumns
        {
            get => _inputColumns ?? (_inputColumns = new InputList<Inputs.DataSetInputColumnArgs>());
            set => _inputColumns = value;
        }

        /// <summary>
        /// &lt;p&gt;The name of the relational table.&lt;/p&gt;
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// &lt;p&gt;The schema name. This name applies to certain relational database engines.&lt;/p&gt;
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        public DataSetRelationalTableArgs()
        {
        }
        public static new DataSetRelationalTableArgs Empty => new DataSetRelationalTableArgs();
    }
}
