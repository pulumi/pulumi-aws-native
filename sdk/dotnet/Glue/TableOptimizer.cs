// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue
{
    /// <summary>
    /// Resource Type definition for AWS::Glue::TableOptimizer
    /// </summary>
    [Obsolete(@"TableOptimizer is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:glue:TableOptimizer")]
    public partial class TableOptimizer : global::Pulumi.CustomResource
    {
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("catalogId")]
        public Output<string> CatalogId { get; private set; } = null!;

        [Output("databaseName")]
        public Output<string> DatabaseName { get; private set; } = null!;

        [Output("tableName")]
        public Output<string> TableName { get; private set; } = null!;

        [Output("tableOptimizerConfiguration")]
        public Output<Outputs.TableOptimizerConfiguration> TableOptimizerConfiguration { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a TableOptimizer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TableOptimizer(string name, TableOptimizerArgs args, CustomResourceOptions? options = null)
            : base("aws-native:glue:TableOptimizer", name, args ?? new TableOptimizerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TableOptimizer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:glue:TableOptimizer", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "catalogId",
                    "databaseName",
                    "tableName",
                    "type",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TableOptimizer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TableOptimizer Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TableOptimizer(name, id, options);
        }
    }

    public sealed class TableOptimizerArgs : global::Pulumi.ResourceArgs
    {
        [Input("catalogId", required: true)]
        public Input<string> CatalogId { get; set; } = null!;

        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        [Input("tableName", required: true)]
        public Input<string> TableName { get; set; } = null!;

        [Input("tableOptimizerConfiguration", required: true)]
        public Input<Inputs.TableOptimizerConfigurationArgs> TableOptimizerConfiguration { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public TableOptimizerArgs()
        {
        }
        public static new TableOptimizerArgs Empty => new TableOptimizerArgs();
    }
}