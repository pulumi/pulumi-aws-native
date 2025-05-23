// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    public sealed class FeatureGroupDataCatalogConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Glue table catalog.
        /// </summary>
        [Input("catalog", required: true)]
        public Input<string> Catalog { get; set; } = null!;

        /// <summary>
        /// The name of the Glue table database.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// The name of the Glue table.
        /// </summary>
        [Input("tableName", required: true)]
        public Input<string> TableName { get; set; } = null!;

        public FeatureGroupDataCatalogConfigArgs()
        {
        }
        public static new FeatureGroupDataCatalogConfigArgs Empty => new FeatureGroupDataCatalogConfigArgs();
    }
}
