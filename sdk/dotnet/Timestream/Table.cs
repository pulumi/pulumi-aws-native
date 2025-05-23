// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Timestream
{
    /// <summary>
    /// The AWS::Timestream::Table resource creates a Timestream Table.
    /// </summary>
    [AwsNativeResourceType("aws-native:timestream:Table")]
    public partial class Table : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The `arn` of the table.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The name for the database which the table to be created belongs to.
        /// </summary>
        [Output("databaseName")]
        public Output<string> DatabaseName { get; private set; } = null!;

        /// <summary>
        /// The properties that determine whether magnetic store writes are enabled.
        /// </summary>
        [Output("magneticStoreWriteProperties")]
        public Output<Outputs.MagneticStoreWritePropertiesProperties?> MagneticStoreWriteProperties { get; private set; } = null!;

        /// <summary>
        /// The table name exposed as a read-only attribute.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The retention duration of the memory store and the magnetic store.
        /// </summary>
        [Output("retentionProperties")]
        public Output<Outputs.RetentionPropertiesProperties?> RetentionProperties { get; private set; } = null!;

        /// <summary>
        /// A Schema specifies the expected data model of the table.
        /// </summary>
        [Output("schema")]
        public Output<Outputs.SchemaProperties?> Schema { get; private set; } = null!;

        /// <summary>
        /// The name for the table. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name.
        /// </summary>
        [Output("tableName")]
        public Output<string?> TableName { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Table resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Table(string name, TableArgs args, CustomResourceOptions? options = null)
            : base("aws-native:timestream:Table", name, args ?? new TableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Table(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:timestream:Table", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "databaseName",
                    "tableName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Table resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Table Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Table(name, id, options);
        }
    }

    public sealed class TableArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name for the database which the table to be created belongs to.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The properties that determine whether magnetic store writes are enabled.
        /// </summary>
        [Input("magneticStoreWriteProperties")]
        public Input<Inputs.MagneticStoreWritePropertiesPropertiesArgs>? MagneticStoreWriteProperties { get; set; }

        /// <summary>
        /// The retention duration of the memory store and the magnetic store.
        /// </summary>
        [Input("retentionProperties")]
        public Input<Inputs.RetentionPropertiesPropertiesArgs>? RetentionProperties { get; set; }

        /// <summary>
        /// A Schema specifies the expected data model of the table.
        /// </summary>
        [Input("schema")]
        public Input<Inputs.SchemaPropertiesArgs>? Schema { get; set; }

        /// <summary>
        /// The name for the table. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name.
        /// </summary>
        [Input("tableName")]
        public Input<string>? TableName { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public TableArgs()
        {
        }
        public static new TableArgs Empty => new TableArgs();
    }
}
