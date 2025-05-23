// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    public static class GetSnapshotBlockPublicAccess
    {
        /// <summary>
        /// Resource Type definition for AWS::EC2::SnapshotBlockPublicAccess
        /// </summary>
        public static Task<GetSnapshotBlockPublicAccessResult> InvokeAsync(GetSnapshotBlockPublicAccessArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSnapshotBlockPublicAccessResult>("aws-native:ec2:getSnapshotBlockPublicAccess", args ?? new GetSnapshotBlockPublicAccessArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::EC2::SnapshotBlockPublicAccess
        /// </summary>
        public static Output<GetSnapshotBlockPublicAccessResult> Invoke(GetSnapshotBlockPublicAccessInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSnapshotBlockPublicAccessResult>("aws-native:ec2:getSnapshotBlockPublicAccess", args ?? new GetSnapshotBlockPublicAccessInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::EC2::SnapshotBlockPublicAccess
        /// </summary>
        public static Output<GetSnapshotBlockPublicAccessResult> Invoke(GetSnapshotBlockPublicAccessInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSnapshotBlockPublicAccessResult>("aws-native:ec2:getSnapshotBlockPublicAccess", args ?? new GetSnapshotBlockPublicAccessInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSnapshotBlockPublicAccessArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier for the specified AWS account.
        /// </summary>
        [Input("accountId", required: true)]
        public string AccountId { get; set; } = null!;

        public GetSnapshotBlockPublicAccessArgs()
        {
        }
        public static new GetSnapshotBlockPublicAccessArgs Empty => new GetSnapshotBlockPublicAccessArgs();
    }

    public sealed class GetSnapshotBlockPublicAccessInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier for the specified AWS account.
        /// </summary>
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        public GetSnapshotBlockPublicAccessInvokeArgs()
        {
        }
        public static new GetSnapshotBlockPublicAccessInvokeArgs Empty => new GetSnapshotBlockPublicAccessInvokeArgs();
    }


    [OutputType]
    public sealed class GetSnapshotBlockPublicAccessResult
    {
        /// <summary>
        /// The identifier for the specified AWS account.
        /// </summary>
        public readonly string? AccountId;
        /// <summary>
        /// The state of EBS Snapshot Block Public Access.
        /// </summary>
        public readonly Pulumi.AwsNative.Ec2.SnapshotBlockPublicAccessState? State;

        [OutputConstructor]
        private GetSnapshotBlockPublicAccessResult(
            string? accountId,

            Pulumi.AwsNative.Ec2.SnapshotBlockPublicAccessState? state)
        {
            AccountId = accountId;
            State = state;
        }
    }
}
