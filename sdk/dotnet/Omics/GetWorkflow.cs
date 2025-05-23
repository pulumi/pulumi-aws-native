// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Omics
{
    public static class GetWorkflow
    {
        /// <summary>
        /// Definition of AWS::Omics::Workflow Resource Type
        /// </summary>
        public static Task<GetWorkflowResult> InvokeAsync(GetWorkflowArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWorkflowResult>("aws-native:omics:getWorkflow", args ?? new GetWorkflowArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::Omics::Workflow Resource Type
        /// </summary>
        public static Output<GetWorkflowResult> Invoke(GetWorkflowInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWorkflowResult>("aws-native:omics:getWorkflow", args ?? new GetWorkflowInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::Omics::Workflow Resource Type
        /// </summary>
        public static Output<GetWorkflowResult> Invoke(GetWorkflowInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetWorkflowResult>("aws-native:omics:getWorkflow", args ?? new GetWorkflowInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWorkflowArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The workflow's ID.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetWorkflowArgs()
        {
        }
        public static new GetWorkflowArgs Empty => new GetWorkflowArgs();
    }

    public sealed class GetWorkflowInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The workflow's ID.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetWorkflowInvokeArgs()
        {
        }
        public static new GetWorkflowInvokeArgs Empty => new GetWorkflowInvokeArgs();
    }


    [OutputType]
    public sealed class GetWorkflowResult
    {
        /// <summary>
        /// The ARN for the workflow.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// When the workflow was created.
        /// </summary>
        public readonly string? CreationTime;
        /// <summary>
        /// The parameter's description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The workflow's ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The workflow's name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The workflow's status.
        /// </summary>
        public readonly Pulumi.AwsNative.Omics.WorkflowStatus? Status;
        public readonly Pulumi.AwsNative.Omics.WorkflowStorageType? StorageType;
        /// <summary>
        /// Tags for the workflow.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The workflow's type.
        /// </summary>
        public readonly Pulumi.AwsNative.Omics.WorkflowType? Type;
        public readonly string? Uuid;

        [OutputConstructor]
        private GetWorkflowResult(
            string? arn,

            string? creationTime,

            string? description,

            string? id,

            string? name,

            Pulumi.AwsNative.Omics.WorkflowStatus? status,

            Pulumi.AwsNative.Omics.WorkflowStorageType? storageType,

            ImmutableDictionary<string, string>? tags,

            Pulumi.AwsNative.Omics.WorkflowType? type,

            string? uuid)
        {
            Arn = arn;
            CreationTime = creationTime;
            Description = description;
            Id = id;
            Name = name;
            Status = status;
            StorageType = storageType;
            Tags = tags;
            Type = type;
            Uuid = uuid;
        }
    }
}
