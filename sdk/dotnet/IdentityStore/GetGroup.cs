// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IdentityStore
{
    public static class GetGroup
    {
        /// <summary>
        /// Resource Type definition for AWS::IdentityStore::Group
        /// </summary>
        public static Task<GetGroupResult> InvokeAsync(GetGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupResult>("aws-native:identitystore:getGroup", args ?? new GetGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::IdentityStore::Group
        /// </summary>
        public static Output<GetGroupResult> Invoke(GetGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupResult>("aws-native:identitystore:getGroup", args ?? new GetGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique identifier for a group in the identity store.
        /// </summary>
        [Input("groupId", required: true)]
        public string GroupId { get; set; } = null!;

        /// <summary>
        /// The globally unique identifier for the identity store.
        /// </summary>
        [Input("identityStoreId", required: true)]
        public string IdentityStoreId { get; set; } = null!;

        public GetGroupArgs()
        {
        }
        public static new GetGroupArgs Empty => new GetGroupArgs();
    }

    public sealed class GetGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique identifier for a group in the identity store.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// The globally unique identifier for the identity store.
        /// </summary>
        [Input("identityStoreId", required: true)]
        public Input<string> IdentityStoreId { get; set; } = null!;

        public GetGroupInvokeArgs()
        {
        }
        public static new GetGroupInvokeArgs Empty => new GetGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetGroupResult
    {
        /// <summary>
        /// A string containing the description of the group.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// A string containing the name of the group. This value is commonly displayed when the group is referenced.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// The unique identifier for a group in the identity store.
        /// </summary>
        public readonly string? GroupId;

        [OutputConstructor]
        private GetGroupResult(
            string? description,

            string? displayName,

            string? groupId)
        {
            Description = description;
            DisplayName = displayName;
            GroupId = groupId;
        }
    }
}