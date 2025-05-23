// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue
{
    public static class GetSchemaVersion
    {
        /// <summary>
        /// This resource represents an individual schema version of a schema defined in Glue Schema Registry.
        /// </summary>
        public static Task<GetSchemaVersionResult> InvokeAsync(GetSchemaVersionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSchemaVersionResult>("aws-native:glue:getSchemaVersion", args ?? new GetSchemaVersionArgs(), options.WithDefaults());

        /// <summary>
        /// This resource represents an individual schema version of a schema defined in Glue Schema Registry.
        /// </summary>
        public static Output<GetSchemaVersionResult> Invoke(GetSchemaVersionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSchemaVersionResult>("aws-native:glue:getSchemaVersion", args ?? new GetSchemaVersionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// This resource represents an individual schema version of a schema defined in Glue Schema Registry.
        /// </summary>
        public static Output<GetSchemaVersionResult> Invoke(GetSchemaVersionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSchemaVersionResult>("aws-native:glue:getSchemaVersion", args ?? new GetSchemaVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSchemaVersionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Represents the version ID associated with the schema version.
        /// </summary>
        [Input("versionId", required: true)]
        public string VersionId { get; set; } = null!;

        public GetSchemaVersionArgs()
        {
        }
        public static new GetSchemaVersionArgs Empty => new GetSchemaVersionArgs();
    }

    public sealed class GetSchemaVersionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Represents the version ID associated with the schema version.
        /// </summary>
        [Input("versionId", required: true)]
        public Input<string> VersionId { get; set; } = null!;

        public GetSchemaVersionInvokeArgs()
        {
        }
        public static new GetSchemaVersionInvokeArgs Empty => new GetSchemaVersionInvokeArgs();
    }


    [OutputType]
    public sealed class GetSchemaVersionResult
    {
        /// <summary>
        /// Represents the version ID associated with the schema version.
        /// </summary>
        public readonly string? VersionId;

        [OutputConstructor]
        private GetSchemaVersionResult(string? versionId)
        {
            VersionId = versionId;
        }
    }
}
