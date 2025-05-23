// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Timestream
{
    public static class GetTable
    {
        /// <summary>
        /// The AWS::Timestream::Table resource creates a Timestream Table.
        /// </summary>
        public static Task<GetTableResult> InvokeAsync(GetTableArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTableResult>("aws-native:timestream:getTable", args ?? new GetTableArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::Timestream::Table resource creates a Timestream Table.
        /// </summary>
        public static Output<GetTableResult> Invoke(GetTableInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTableResult>("aws-native:timestream:getTable", args ?? new GetTableInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::Timestream::Table resource creates a Timestream Table.
        /// </summary>
        public static Output<GetTableResult> Invoke(GetTableInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetTableResult>("aws-native:timestream:getTable", args ?? new GetTableInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTableArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name for the database which the table to be created belongs to.
        /// </summary>
        [Input("databaseName", required: true)]
        public string DatabaseName { get; set; } = null!;

        /// <summary>
        /// The name for the table. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name.
        /// </summary>
        [Input("tableName", required: true)]
        public string TableName { get; set; } = null!;

        public GetTableArgs()
        {
        }
        public static new GetTableArgs Empty => new GetTableArgs();
    }

    public sealed class GetTableInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name for the database which the table to be created belongs to.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The name for the table. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name.
        /// </summary>
        [Input("tableName", required: true)]
        public Input<string> TableName { get; set; } = null!;

        public GetTableInvokeArgs()
        {
        }
        public static new GetTableInvokeArgs Empty => new GetTableInvokeArgs();
    }


    [OutputType]
    public sealed class GetTableResult
    {
        /// <summary>
        /// The `arn` of the table.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The properties that determine whether magnetic store writes are enabled.
        /// </summary>
        public readonly Outputs.MagneticStoreWritePropertiesProperties? MagneticStoreWriteProperties;
        /// <summary>
        /// The table name exposed as a read-only attribute.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The retention duration of the memory store and the magnetic store.
        /// </summary>
        public readonly Outputs.RetentionPropertiesProperties? RetentionProperties;
        /// <summary>
        /// A Schema specifies the expected data model of the table.
        /// </summary>
        public readonly Outputs.SchemaProperties? Schema;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetTableResult(
            string? arn,

            Outputs.MagneticStoreWritePropertiesProperties? magneticStoreWriteProperties,

            string? name,

            Outputs.RetentionPropertiesProperties? retentionProperties,

            Outputs.SchemaProperties? schema,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            Arn = arn;
            MagneticStoreWriteProperties = magneticStoreWriteProperties;
            Name = name;
            RetentionProperties = retentionProperties;
            Schema = schema;
            Tags = tags;
        }
    }
}
