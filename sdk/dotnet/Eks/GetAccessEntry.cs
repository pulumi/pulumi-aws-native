// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Eks
{
    public static class GetAccessEntry
    {
        /// <summary>
        /// An object representing an Amazon EKS AccessEntry.
        /// </summary>
        public static Task<GetAccessEntryResult> InvokeAsync(GetAccessEntryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccessEntryResult>("aws-native:eks:getAccessEntry", args ?? new GetAccessEntryArgs(), options.WithDefaults());

        /// <summary>
        /// An object representing an Amazon EKS AccessEntry.
        /// </summary>
        public static Output<GetAccessEntryResult> Invoke(GetAccessEntryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccessEntryResult>("aws-native:eks:getAccessEntry", args ?? new GetAccessEntryInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// An object representing an Amazon EKS AccessEntry.
        /// </summary>
        public static Output<GetAccessEntryResult> Invoke(GetAccessEntryInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccessEntryResult>("aws-native:eks:getAccessEntry", args ?? new GetAccessEntryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccessEntryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The cluster that the access entry is created for.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// The principal ARN that the access entry is created for.
        /// </summary>
        [Input("principalArn", required: true)]
        public string PrincipalArn { get; set; } = null!;

        public GetAccessEntryArgs()
        {
        }
        public static new GetAccessEntryArgs Empty => new GetAccessEntryArgs();
    }

    public sealed class GetAccessEntryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The cluster that the access entry is created for.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// The principal ARN that the access entry is created for.
        /// </summary>
        [Input("principalArn", required: true)]
        public Input<string> PrincipalArn { get; set; } = null!;

        public GetAccessEntryInvokeArgs()
        {
        }
        public static new GetAccessEntryInvokeArgs Empty => new GetAccessEntryInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccessEntryResult
    {
        /// <summary>
        /// The ARN of the access entry.
        /// </summary>
        public readonly string? AccessEntryArn;
        /// <summary>
        /// An array of access policies that are associated with the access entry.
        /// </summary>
        public readonly ImmutableArray<Outputs.AccessEntryAccessPolicy> AccessPolicies;
        /// <summary>
        /// The Kubernetes groups that the access entry is associated with.
        /// </summary>
        public readonly ImmutableArray<string> KubernetesGroups;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// The Kubernetes user that the access entry is associated with.
        /// </summary>
        public readonly string? Username;

        [OutputConstructor]
        private GetAccessEntryResult(
            string? accessEntryArn,

            ImmutableArray<Outputs.AccessEntryAccessPolicy> accessPolicies,

            ImmutableArray<string> kubernetesGroups,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            string? username)
        {
            AccessEntryArn = accessEntryArn;
            AccessPolicies = accessPolicies;
            KubernetesGroups = kubernetesGroups;
            Tags = tags;
            Username = username;
        }
    }
}
