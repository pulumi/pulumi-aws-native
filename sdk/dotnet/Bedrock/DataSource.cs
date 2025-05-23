// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock
{
    /// <summary>
    /// Definition of AWS::Bedrock::DataSource Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:bedrock:DataSource")]
    public partial class DataSource : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The time at which the data source was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The data deletion policy for the data source.
        /// </summary>
        [Output("dataDeletionPolicy")]
        public Output<Pulumi.AwsNative.Bedrock.DataSourceDataDeletionPolicy?> DataDeletionPolicy { get; private set; } = null!;

        /// <summary>
        /// The connection configuration for the data source.
        /// </summary>
        [Output("dataSourceConfiguration")]
        public Output<Outputs.DataSourceConfiguration> DataSourceConfiguration { get; private set; } = null!;

        /// <summary>
        /// Identifier for a resource.
        /// </summary>
        [Output("dataSourceId")]
        public Output<string> DataSourceId { get; private set; } = null!;

        /// <summary>
        /// The status of the data source. The following statuses are possible:
        /// 
        /// - Available – The data source has been created and is ready for ingestion into the knowledge base.
        /// - Deleting – The data source is being deleted.
        /// </summary>
        [Output("dataSourceStatus")]
        public Output<Pulumi.AwsNative.Bedrock.DataSourceStatus> DataSourceStatus { get; private set; } = null!;

        /// <summary>
        /// Description of the Resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The details of the failure reasons related to the data source.
        /// </summary>
        [Output("failureReasons")]
        public Output<ImmutableArray<string>> FailureReasons { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the knowledge base to which to add the data source.
        /// </summary>
        [Output("knowledgeBaseId")]
        public Output<string> KnowledgeBaseId { get; private set; } = null!;

        /// <summary>
        /// The name of the data source.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Contains details about the configuration of the server-side encryption.
        /// </summary>
        [Output("serverSideEncryptionConfiguration")]
        public Output<Outputs.DataSourceServerSideEncryptionConfiguration?> ServerSideEncryptionConfiguration { get; private set; } = null!;

        /// <summary>
        /// The time at which the knowledge base was last updated.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// Contains details about how to ingest the documents in the data source.
        /// </summary>
        [Output("vectorIngestionConfiguration")]
        public Output<Outputs.DataSourceVectorIngestionConfiguration?> VectorIngestionConfiguration { get; private set; } = null!;


        /// <summary>
        /// Create a DataSource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataSource(string name, DataSourceArgs args, CustomResourceOptions? options = null)
            : base("aws-native:bedrock:DataSource", name, args ?? new DataSourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataSource(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:bedrock:DataSource", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "dataSourceConfiguration.type",
                    "knowledgeBaseId",
                    "vectorIngestionConfiguration.chunkingConfiguration",
                    "vectorIngestionConfiguration.parsingConfiguration",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DataSource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataSource Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DataSource(name, id, options);
        }
    }

    public sealed class DataSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The data deletion policy for the data source.
        /// </summary>
        [Input("dataDeletionPolicy")]
        public Input<Pulumi.AwsNative.Bedrock.DataSourceDataDeletionPolicy>? DataDeletionPolicy { get; set; }

        /// <summary>
        /// The connection configuration for the data source.
        /// </summary>
        [Input("dataSourceConfiguration", required: true)]
        public Input<Inputs.DataSourceConfigurationArgs> DataSourceConfiguration { get; set; } = null!;

        /// <summary>
        /// Description of the Resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The unique identifier of the knowledge base to which to add the data source.
        /// </summary>
        [Input("knowledgeBaseId", required: true)]
        public Input<string> KnowledgeBaseId { get; set; } = null!;

        /// <summary>
        /// The name of the data source.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Contains details about the configuration of the server-side encryption.
        /// </summary>
        [Input("serverSideEncryptionConfiguration")]
        public Input<Inputs.DataSourceServerSideEncryptionConfigurationArgs>? ServerSideEncryptionConfiguration { get; set; }

        /// <summary>
        /// Contains details about how to ingest the documents in the data source.
        /// </summary>
        [Input("vectorIngestionConfiguration")]
        public Input<Inputs.DataSourceVectorIngestionConfigurationArgs>? VectorIngestionConfiguration { get; set; }

        public DataSourceArgs()
        {
        }
        public static new DataSourceArgs Empty => new DataSourceArgs();
    }
}
