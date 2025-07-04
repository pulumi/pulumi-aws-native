// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RedshiftServerless
{
    public static class GetSnapshot
    {
        /// <summary>
        /// Resource Type definition for AWS::RedshiftServerless::Snapshot Resource Type.
        /// </summary>
        public static Task<GetSnapshotResult> InvokeAsync(GetSnapshotArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSnapshotResult>("aws-native:redshiftserverless:getSnapshot", args ?? new GetSnapshotArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::RedshiftServerless::Snapshot Resource Type.
        /// </summary>
        public static Output<GetSnapshotResult> Invoke(GetSnapshotInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSnapshotResult>("aws-native:redshiftserverless:getSnapshot", args ?? new GetSnapshotInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::RedshiftServerless::Snapshot Resource Type.
        /// </summary>
        public static Output<GetSnapshotResult> Invoke(GetSnapshotInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSnapshotResult>("aws-native:redshiftserverless:getSnapshot", args ?? new GetSnapshotInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSnapshotArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the snapshot.
        /// </summary>
        [Input("snapshotName", required: true)]
        public string SnapshotName { get; set; } = null!;

        public GetSnapshotArgs()
        {
        }
        public static new GetSnapshotArgs Empty => new GetSnapshotArgs();
    }

    public sealed class GetSnapshotInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the snapshot.
        /// </summary>
        [Input("snapshotName", required: true)]
        public Input<string> SnapshotName { get; set; } = null!;

        public GetSnapshotInvokeArgs()
        {
        }
        public static new GetSnapshotInvokeArgs Empty => new GetSnapshotInvokeArgs();
    }


    [OutputType]
    public sealed class GetSnapshotResult
    {
        /// <summary>
        /// The owner account of the snapshot.
        /// </summary>
        public readonly string? OwnerAccount;
        /// <summary>
        /// The retention period of the snapshot.
        /// </summary>
        public readonly int? RetentionPeriod;
        /// <summary>
        /// Definition for snapshot resource
        /// </summary>
        public readonly Outputs.Snapshot? SnapshotValue;

        [OutputConstructor]
        private GetSnapshotResult(
            string? ownerAccount,

            int? retentionPeriod,

            Outputs.Snapshot? snapshot)
        {
            OwnerAccount = ownerAccount;
            RetentionPeriod = retentionPeriod;
            SnapshotValue = snapshot;
        }
    }
}
