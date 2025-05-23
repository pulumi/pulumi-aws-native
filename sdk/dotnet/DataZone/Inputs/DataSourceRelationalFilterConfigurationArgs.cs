// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone.Inputs
{

    /// <summary>
    /// The relational filter configuration for the data source.
    /// </summary>
    public sealed class DataSourceRelationalFilterConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The database name specified in the relational filter configuration for the data source.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        [Input("filterExpressions")]
        private InputList<Inputs.DataSourceFilterExpressionArgs>? _filterExpressions;

        /// <summary>
        /// The filter expressions specified in the relational filter configuration for the data source.
        /// </summary>
        public InputList<Inputs.DataSourceFilterExpressionArgs> FilterExpressions
        {
            get => _filterExpressions ?? (_filterExpressions = new InputList<Inputs.DataSourceFilterExpressionArgs>());
            set => _filterExpressions = value;
        }

        /// <summary>
        /// The schema name specified in the relational filter configuration for the data source.
        /// </summary>
        [Input("schemaName")]
        public Input<string>? SchemaName { get; set; }

        public DataSourceRelationalFilterConfigurationArgs()
        {
        }
        public static new DataSourceRelationalFilterConfigurationArgs Empty => new DataSourceRelationalFilterConfigurationArgs();
    }
}
