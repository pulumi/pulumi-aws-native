// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Iam
{
    public static class GetManagedPolicy
    {
        /// <summary>
        /// Resource Type definition for AWS::IAM::ManagedPolicy
        /// </summary>
        public static Task<GetManagedPolicyResult> InvokeAsync(GetManagedPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetManagedPolicyResult>("aws-native:iam:getManagedPolicy", args ?? new GetManagedPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::IAM::ManagedPolicy
        /// </summary>
        public static Output<GetManagedPolicyResult> Invoke(GetManagedPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetManagedPolicyResult>("aws-native:iam:getManagedPolicy", args ?? new GetManagedPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetManagedPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the managed policy
        /// </summary>
        [Input("policyArn", required: true)]
        public string PolicyArn { get; set; } = null!;

        public GetManagedPolicyArgs()
        {
        }
        public static new GetManagedPolicyArgs Empty => new GetManagedPolicyArgs();
    }

    public sealed class GetManagedPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the managed policy
        /// </summary>
        [Input("policyArn", required: true)]
        public Input<string> PolicyArn { get; set; } = null!;

        public GetManagedPolicyInvokeArgs()
        {
        }
        public static new GetManagedPolicyInvokeArgs Empty => new GetManagedPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetManagedPolicyResult
    {
        /// <summary>
        /// The number of entities (users, groups, and roles) that the policy is attached to.
        /// </summary>
        public readonly int? AttachmentCount;
        /// <summary>
        /// The date and time, in ISO 8601 date-time format, when the policy was created.
        /// </summary>
        public readonly string? CreateDate;
        /// <summary>
        /// The identifier for the version of the policy that is set as the default version.
        /// </summary>
        public readonly string? DefaultVersionId;
        /// <summary>
        /// The name (friendly name, not ARN) of the group to attach the policy to.
        /// </summary>
        public readonly ImmutableArray<string> Groups;
        /// <summary>
        /// Specifies whether the policy can be attached to an IAM user, group, or role.
        /// </summary>
        public readonly bool? IsAttachable;
        /// <summary>
        /// The number of entities (users and roles) for which the policy is used to set the permissions boundary.
        /// </summary>
        public readonly int? PermissionsBoundaryUsageCount;
        /// <summary>
        /// Amazon Resource Name (ARN) of the managed policy
        /// </summary>
        public readonly string? PolicyArn;
        /// <summary>
        /// The JSON policy document that you want to use as the content for the new policy.
        /// </summary>
        public readonly object? PolicyDocument;
        /// <summary>
        /// The stable and unique string identifying the policy.
        /// </summary>
        public readonly string? PolicyId;
        /// <summary>
        /// The name (friendly name, not ARN) of the role to attach the policy to.
        /// </summary>
        public readonly ImmutableArray<string> Roles;
        /// <summary>
        /// The date and time, in ISO 8601 date-time format, when the policy was last updated.
        /// </summary>
        public readonly string? UpdateDate;
        /// <summary>
        /// The name (friendly name, not ARN) of the IAM user to attach the policy to.
        /// </summary>
        public readonly ImmutableArray<string> Users;

        [OutputConstructor]
        private GetManagedPolicyResult(
            int? attachmentCount,

            string? createDate,

            string? defaultVersionId,

            ImmutableArray<string> groups,

            bool? isAttachable,

            int? permissionsBoundaryUsageCount,

            string? policyArn,

            object? policyDocument,

            string? policyId,

            ImmutableArray<string> roles,

            string? updateDate,

            ImmutableArray<string> users)
        {
            AttachmentCount = attachmentCount;
            CreateDate = createDate;
            DefaultVersionId = defaultVersionId;
            Groups = groups;
            IsAttachable = isAttachable;
            PermissionsBoundaryUsageCount = permissionsBoundaryUsageCount;
            PolicyArn = policyArn;
            PolicyDocument = policyDocument;
            PolicyId = policyId;
            Roles = roles;
            UpdateDate = updateDate;
            Users = users;
        }
    }
}