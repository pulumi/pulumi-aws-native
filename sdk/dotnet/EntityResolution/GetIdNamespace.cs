// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EntityResolution
{
    public static class GetIdNamespace
    {
        /// <summary>
        /// IdNamespace defined in AWS Entity Resolution service
        /// </summary>
        public static Task<GetIdNamespaceResult> InvokeAsync(GetIdNamespaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIdNamespaceResult>("aws-native:entityresolution:getIdNamespace", args ?? new GetIdNamespaceArgs(), options.WithDefaults());

        /// <summary>
        /// IdNamespace defined in AWS Entity Resolution service
        /// </summary>
        public static Output<GetIdNamespaceResult> Invoke(GetIdNamespaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIdNamespaceResult>("aws-native:entityresolution:getIdNamespace", args ?? new GetIdNamespaceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// IdNamespace defined in AWS Entity Resolution service
        /// </summary>
        public static Output<GetIdNamespaceResult> Invoke(GetIdNamespaceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIdNamespaceResult>("aws-native:entityresolution:getIdNamespace", args ?? new GetIdNamespaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIdNamespaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the ID namespace.
        /// </summary>
        [Input("idNamespaceName", required: true)]
        public string IdNamespaceName { get; set; } = null!;

        public GetIdNamespaceArgs()
        {
        }
        public static new GetIdNamespaceArgs Empty => new GetIdNamespaceArgs();
    }

    public sealed class GetIdNamespaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the ID namespace.
        /// </summary>
        [Input("idNamespaceName", required: true)]
        public Input<string> IdNamespaceName { get; set; } = null!;

        public GetIdNamespaceInvokeArgs()
        {
        }
        public static new GetIdNamespaceInvokeArgs Empty => new GetIdNamespaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetIdNamespaceResult
    {
        /// <summary>
        /// The date and time when the IdNamespace was created
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// The description of the ID namespace.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Determines the properties of `IdMappingWorflow` where this `IdNamespace` can be used as a `Source` or a `Target` .
        /// </summary>
        public readonly ImmutableArray<Outputs.IdNamespaceIdMappingWorkflowProperties> IdMappingWorkflowProperties;
        /// <summary>
        /// The arn associated with the IdNamespace
        /// </summary>
        public readonly string? IdNamespaceArn;
        /// <summary>
        /// A list of `InputSource` objects, which have the fields `InputSourceARN` and `SchemaName` .
        /// </summary>
        public readonly ImmutableArray<Outputs.IdNamespaceInputSource> InputSourceConfig;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role. AWS Entity Resolution assumes this role to access the resources defined in this `IdNamespace` on your behalf as part of the workflow run.
        /// </summary>
        public readonly string? RoleArn;
        /// <summary>
        /// The tags used to organize, track, or control access for this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// The type of ID namespace. There are two types: `SOURCE` and `TARGET` .
        /// 
        /// The `SOURCE` contains configurations for `sourceId` data that will be processed in an ID mapping workflow.
        /// 
        /// The `TARGET` contains a configuration of `targetId` which all `sourceIds` will resolve to.
        /// </summary>
        public readonly Pulumi.AwsNative.EntityResolution.IdNamespaceType? Type;
        /// <summary>
        /// The date and time when the IdNamespace was updated
        /// </summary>
        public readonly string? UpdatedAt;

        [OutputConstructor]
        private GetIdNamespaceResult(
            string? createdAt,

            string? description,

            ImmutableArray<Outputs.IdNamespaceIdMappingWorkflowProperties> idMappingWorkflowProperties,

            string? idNamespaceArn,

            ImmutableArray<Outputs.IdNamespaceInputSource> inputSourceConfig,

            string? roleArn,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            Pulumi.AwsNative.EntityResolution.IdNamespaceType? type,

            string? updatedAt)
        {
            CreatedAt = createdAt;
            Description = description;
            IdMappingWorkflowProperties = idMappingWorkflowProperties;
            IdNamespaceArn = idNamespaceArn;
            InputSourceConfig = inputSourceConfig;
            RoleArn = roleArn;
            Tags = tags;
            Type = type;
            UpdatedAt = updatedAt;
        }
    }
}
