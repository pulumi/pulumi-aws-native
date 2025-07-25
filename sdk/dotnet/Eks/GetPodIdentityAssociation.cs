// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Eks
{
    public static class GetPodIdentityAssociation
    {
        /// <summary>
        /// An object representing an Amazon EKS PodIdentityAssociation.
        /// </summary>
        public static Task<GetPodIdentityAssociationResult> InvokeAsync(GetPodIdentityAssociationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPodIdentityAssociationResult>("aws-native:eks:getPodIdentityAssociation", args ?? new GetPodIdentityAssociationArgs(), options.WithDefaults());

        /// <summary>
        /// An object representing an Amazon EKS PodIdentityAssociation.
        /// </summary>
        public static Output<GetPodIdentityAssociationResult> Invoke(GetPodIdentityAssociationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPodIdentityAssociationResult>("aws-native:eks:getPodIdentityAssociation", args ?? new GetPodIdentityAssociationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// An object representing an Amazon EKS PodIdentityAssociation.
        /// </summary>
        public static Output<GetPodIdentityAssociationResult> Invoke(GetPodIdentityAssociationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetPodIdentityAssociationResult>("aws-native:eks:getPodIdentityAssociation", args ?? new GetPodIdentityAssociationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPodIdentityAssociationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the pod identity association.
        /// </summary>
        [Input("associationArn", required: true)]
        public string AssociationArn { get; set; } = null!;

        public GetPodIdentityAssociationArgs()
        {
        }
        public static new GetPodIdentityAssociationArgs Empty => new GetPodIdentityAssociationArgs();
    }

    public sealed class GetPodIdentityAssociationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the pod identity association.
        /// </summary>
        [Input("associationArn", required: true)]
        public Input<string> AssociationArn { get; set; } = null!;

        public GetPodIdentityAssociationInvokeArgs()
        {
        }
        public static new GetPodIdentityAssociationInvokeArgs Empty => new GetPodIdentityAssociationInvokeArgs();
    }


    [OutputType]
    public sealed class GetPodIdentityAssociationResult
    {
        /// <summary>
        /// The ARN of the pod identity association.
        /// </summary>
        public readonly string? AssociationArn;
        /// <summary>
        /// The ID of the pod identity association.
        /// </summary>
        public readonly string? AssociationId;
        /// <summary>
        /// The Disable Session Tags of the pod identity association.
        /// </summary>
        public readonly bool? DisableSessionTags;
        /// <summary>
        /// The External Id of the pod identity association.
        /// </summary>
        public readonly string? ExternalId;
        /// <summary>
        /// The IAM role ARN that the pod identity association is created for.
        /// </summary>
        public readonly string? RoleArn;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// The Target Role Arn of the pod identity association.
        /// </summary>
        public readonly string? TargetRoleArn;

        [OutputConstructor]
        private GetPodIdentityAssociationResult(
            string? associationArn,

            string? associationId,

            bool? disableSessionTags,

            string? externalId,

            string? roleArn,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            string? targetRoleArn)
        {
            AssociationArn = associationArn;
            AssociationId = associationId;
            DisableSessionTags = disableSessionTags;
            ExternalId = externalId;
            RoleArn = roleArn;
            Tags = tags;
            TargetRoleArn = targetRoleArn;
        }
    }
}
