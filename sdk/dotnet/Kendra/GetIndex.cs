// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra
{
    public static class GetIndex
    {
        /// <summary>
        /// A Kendra index
        /// </summary>
        public static Task<GetIndexResult> InvokeAsync(GetIndexArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIndexResult>("aws-native:kendra:getIndex", args ?? new GetIndexArgs(), options.WithDefaults());

        /// <summary>
        /// A Kendra index
        /// </summary>
        public static Output<GetIndexResult> Invoke(GetIndexInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIndexResult>("aws-native:kendra:getIndex", args ?? new GetIndexInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIndexArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetIndexArgs()
        {
        }
        public static new GetIndexArgs Empty => new GetIndexArgs();
    }

    public sealed class GetIndexInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetIndexInvokeArgs()
        {
        }
        public static new GetIndexInvokeArgs Empty => new GetIndexInvokeArgs();
    }


    [OutputType]
    public sealed class GetIndexResult
    {
        public readonly string? Arn;
        /// <summary>
        /// Capacity units
        /// </summary>
        public readonly Outputs.IndexCapacityUnitsConfiguration? CapacityUnits;
        /// <summary>
        /// A description for the index
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Document metadata configurations
        /// </summary>
        public readonly ImmutableArray<Outputs.IndexDocumentMetadataConfiguration> DocumentMetadataConfigurations;
        public readonly string? Id;
        public readonly string? Name;
        public readonly string? RoleArn;
        /// <summary>
        /// Tags for labeling the index
        /// </summary>
        public readonly ImmutableArray<Outputs.IndexTag> Tags;
        public readonly Pulumi.AwsNative.Kendra.IndexUserContextPolicy? UserContextPolicy;
        public readonly ImmutableArray<Outputs.IndexUserTokenConfiguration> UserTokenConfigurations;

        [OutputConstructor]
        private GetIndexResult(
            string? arn,

            Outputs.IndexCapacityUnitsConfiguration? capacityUnits,

            string? description,

            ImmutableArray<Outputs.IndexDocumentMetadataConfiguration> documentMetadataConfigurations,

            string? id,

            string? name,

            string? roleArn,

            ImmutableArray<Outputs.IndexTag> tags,

            Pulumi.AwsNative.Kendra.IndexUserContextPolicy? userContextPolicy,

            ImmutableArray<Outputs.IndexUserTokenConfiguration> userTokenConfigurations)
        {
            Arn = arn;
            CapacityUnits = capacityUnits;
            Description = description;
            DocumentMetadataConfigurations = documentMetadataConfigurations;
            Id = id;
            Name = name;
            RoleArn = roleArn;
            Tags = tags;
            UserContextPolicy = userContextPolicy;
            UserTokenConfigurations = userTokenConfigurations;
        }
    }
}