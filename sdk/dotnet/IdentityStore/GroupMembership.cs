// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IdentityStore
{
    /// <summary>
    /// Resource Type Definition for AWS:IdentityStore::GroupMembership
    /// </summary>
    [AwsNativeResourceType("aws-native:identitystore:GroupMembership")]
    public partial class GroupMembership : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The unique identifier for a group in the identity store.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// The globally unique identifier for the identity store.
        /// </summary>
        [Output("identityStoreId")]
        public Output<string> IdentityStoreId { get; private set; } = null!;

        /// <summary>
        /// An object containing the identifier of a group member.
        /// </summary>
        [Output("memberId")]
        public Output<Outputs.GroupMembershipMemberId> MemberId { get; private set; } = null!;

        /// <summary>
        /// The identifier for a GroupMembership in the identity store.
        /// </summary>
        [Output("membershipId")]
        public Output<string> MembershipId { get; private set; } = null!;


        /// <summary>
        /// Create a GroupMembership resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupMembership(string name, GroupMembershipArgs args, CustomResourceOptions? options = null)
            : base("aws-native:identitystore:GroupMembership", name, args ?? new GroupMembershipArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupMembership(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:identitystore:GroupMembership", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GroupMembership resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupMembership Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GroupMembership(name, id, options);
        }
    }

    public sealed class GroupMembershipArgs : global::Pulumi.ResourceArgs
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

        /// <summary>
        /// An object containing the identifier of a group member.
        /// </summary>
        [Input("memberId", required: true)]
        public Input<Inputs.GroupMembershipMemberIdArgs> MemberId { get; set; } = null!;

        public GroupMembershipArgs()
        {
        }
        public static new GroupMembershipArgs Empty => new GroupMembershipArgs();
    }
}