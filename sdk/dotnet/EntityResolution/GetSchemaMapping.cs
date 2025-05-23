// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EntityResolution
{
    public static class GetSchemaMapping
    {
        /// <summary>
        /// SchemaMapping defined in AWS Entity Resolution service
        /// </summary>
        public static Task<GetSchemaMappingResult> InvokeAsync(GetSchemaMappingArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSchemaMappingResult>("aws-native:entityresolution:getSchemaMapping", args ?? new GetSchemaMappingArgs(), options.WithDefaults());

        /// <summary>
        /// SchemaMapping defined in AWS Entity Resolution service
        /// </summary>
        public static Output<GetSchemaMappingResult> Invoke(GetSchemaMappingInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSchemaMappingResult>("aws-native:entityresolution:getSchemaMapping", args ?? new GetSchemaMappingInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// SchemaMapping defined in AWS Entity Resolution service
        /// </summary>
        public static Output<GetSchemaMappingResult> Invoke(GetSchemaMappingInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSchemaMappingResult>("aws-native:entityresolution:getSchemaMapping", args ?? new GetSchemaMappingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSchemaMappingArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the SchemaMapping
        /// </summary>
        [Input("schemaName", required: true)]
        public string SchemaName { get; set; } = null!;

        public GetSchemaMappingArgs()
        {
        }
        public static new GetSchemaMappingArgs Empty => new GetSchemaMappingArgs();
    }

    public sealed class GetSchemaMappingInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the SchemaMapping
        /// </summary>
        [Input("schemaName", required: true)]
        public Input<string> SchemaName { get; set; } = null!;

        public GetSchemaMappingInvokeArgs()
        {
        }
        public static new GetSchemaMappingInvokeArgs Empty => new GetSchemaMappingInvokeArgs();
    }


    [OutputType]
    public sealed class GetSchemaMappingResult
    {
        public readonly string? CreatedAt;
        /// <summary>
        /// The description of the SchemaMapping
        /// </summary>
        public readonly string? Description;
        public readonly bool? HasWorkflows;
        /// <summary>
        /// The SchemaMapping attributes input
        /// </summary>
        public readonly ImmutableArray<Outputs.SchemaMappingSchemaInputAttribute> MappedInputFields;
        public readonly string? SchemaArn;
        /// <summary>
        /// The tags used to organize, track, or control access for this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        public readonly string? UpdatedAt;

        [OutputConstructor]
        private GetSchemaMappingResult(
            string? createdAt,

            string? description,

            bool? hasWorkflows,

            ImmutableArray<Outputs.SchemaMappingSchemaInputAttribute> mappedInputFields,

            string? schemaArn,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            string? updatedAt)
        {
            CreatedAt = createdAt;
            Description = description;
            HasWorkflows = hasWorkflows;
            MappedInputFields = mappedInputFields;
            SchemaArn = schemaArn;
            Tags = tags;
            UpdatedAt = updatedAt;
        }
    }
}
