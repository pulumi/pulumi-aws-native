// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Emr
{
    public static class GetWalWorkspace
    {
        /// <summary>
        /// Resource schema for AWS::EMR::WALWorkspace Type
        /// </summary>
        public static Task<GetWalWorkspaceResult> InvokeAsync(GetWalWorkspaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWalWorkspaceResult>("aws-native:emr:getWalWorkspace", args ?? new GetWalWorkspaceArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::EMR::WALWorkspace Type
        /// </summary>
        public static Output<GetWalWorkspaceResult> Invoke(GetWalWorkspaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWalWorkspaceResult>("aws-native:emr:getWalWorkspace", args ?? new GetWalWorkspaceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::EMR::WALWorkspace Type
        /// </summary>
        public static Output<GetWalWorkspaceResult> Invoke(GetWalWorkspaceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetWalWorkspaceResult>("aws-native:emr:getWalWorkspace", args ?? new GetWalWorkspaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWalWorkspaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the emrwal container
        /// </summary>
        [Input("walWorkspaceName", required: true)]
        public string WalWorkspaceName { get; set; } = null!;

        public GetWalWorkspaceArgs()
        {
        }
        public static new GetWalWorkspaceArgs Empty => new GetWalWorkspaceArgs();
    }

    public sealed class GetWalWorkspaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the emrwal container
        /// </summary>
        [Input("walWorkspaceName", required: true)]
        public Input<string> WalWorkspaceName { get; set; } = null!;

        public GetWalWorkspaceInvokeArgs()
        {
        }
        public static new GetWalWorkspaceInvokeArgs Empty => new GetWalWorkspaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetWalWorkspaceResult
    {
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetWalWorkspaceResult(ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            Tags = tags;
        }
    }
}
