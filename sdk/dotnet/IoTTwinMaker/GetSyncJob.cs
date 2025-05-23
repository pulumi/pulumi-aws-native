// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTTwinMaker
{
    public static class GetSyncJob
    {
        /// <summary>
        /// Resource schema for AWS::IoTTwinMaker::SyncJob
        /// </summary>
        public static Task<GetSyncJobResult> InvokeAsync(GetSyncJobArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSyncJobResult>("aws-native:iottwinmaker:getSyncJob", args ?? new GetSyncJobArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::IoTTwinMaker::SyncJob
        /// </summary>
        public static Output<GetSyncJobResult> Invoke(GetSyncJobInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSyncJobResult>("aws-native:iottwinmaker:getSyncJob", args ?? new GetSyncJobInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::IoTTwinMaker::SyncJob
        /// </summary>
        public static Output<GetSyncJobResult> Invoke(GetSyncJobInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSyncJobResult>("aws-native:iottwinmaker:getSyncJob", args ?? new GetSyncJobInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSyncJobArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The source of the SyncJob.
        /// </summary>
        [Input("syncSource", required: true)]
        public string SyncSource { get; set; } = null!;

        /// <summary>
        /// The ID of the workspace.
        /// </summary>
        [Input("workspaceId", required: true)]
        public string WorkspaceId { get; set; } = null!;

        public GetSyncJobArgs()
        {
        }
        public static new GetSyncJobArgs Empty => new GetSyncJobArgs();
    }

    public sealed class GetSyncJobInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The source of the SyncJob.
        /// </summary>
        [Input("syncSource", required: true)]
        public Input<string> SyncSource { get; set; } = null!;

        /// <summary>
        /// The ID of the workspace.
        /// </summary>
        [Input("workspaceId", required: true)]
        public Input<string> WorkspaceId { get; set; } = null!;

        public GetSyncJobInvokeArgs()
        {
        }
        public static new GetSyncJobInvokeArgs Empty => new GetSyncJobInvokeArgs();
    }


    [OutputType]
    public sealed class GetSyncJobResult
    {
        /// <summary>
        /// The ARN of the SyncJob.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The date and time when the sync job was created.
        /// </summary>
        public readonly string? CreationDateTime;
        /// <summary>
        /// The state of SyncJob.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// The date and time when the sync job was updated.
        /// </summary>
        public readonly string? UpdateDateTime;

        [OutputConstructor]
        private GetSyncJobResult(
            string? arn,

            string? creationDateTime,

            string? state,

            string? updateDateTime)
        {
            Arn = arn;
            CreationDateTime = creationDateTime;
            State = state;
            UpdateDateTime = updateDateTime;
        }
    }
}
