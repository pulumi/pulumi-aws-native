// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Detective
{
    /// <summary>
    /// Resource schema for AWS::Detective::MemberInvitation
    /// </summary>
    [AwsNativeResourceType("aws-native:detective:MemberInvitation")]
    public partial class MemberInvitation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// When set to true, invitation emails are not sent to the member accounts. Member accounts must still accept the invitation before they are added to the behavior graph. Updating this field has no effect.
        /// </summary>
        [Output("disableEmailNotification")]
        public Output<bool?> DisableEmailNotification { get; private set; } = null!;

        /// <summary>
        /// The ARN of the graph to which the member account will be invited
        /// </summary>
        [Output("graphArn")]
        public Output<string> GraphArn { get; private set; } = null!;

        /// <summary>
        /// The root email address for the account to be invited, for validation. Updating this field has no effect.
        /// </summary>
        [Output("memberEmailAddress")]
        public Output<string> MemberEmailAddress { get; private set; } = null!;

        /// <summary>
        /// The AWS account ID to be invited to join the graph as a member
        /// </summary>
        [Output("memberId")]
        public Output<string> MemberId { get; private set; } = null!;

        /// <summary>
        /// A message to be included in the email invitation sent to the invited account. Updating this field has no effect.
        /// </summary>
        [Output("message")]
        public Output<string?> Message { get; private set; } = null!;


        /// <summary>
        /// Create a MemberInvitation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MemberInvitation(string name, MemberInvitationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:detective:MemberInvitation", name, args ?? new MemberInvitationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MemberInvitation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:detective:MemberInvitation", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "graphArn",
                    "memberId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MemberInvitation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MemberInvitation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new MemberInvitation(name, id, options);
        }
    }

    public sealed class MemberInvitationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When set to true, invitation emails are not sent to the member accounts. Member accounts must still accept the invitation before they are added to the behavior graph. Updating this field has no effect.
        /// </summary>
        [Input("disableEmailNotification")]
        public Input<bool>? DisableEmailNotification { get; set; }

        /// <summary>
        /// The ARN of the graph to which the member account will be invited
        /// </summary>
        [Input("graphArn", required: true)]
        public Input<string> GraphArn { get; set; } = null!;

        /// <summary>
        /// The root email address for the account to be invited, for validation. Updating this field has no effect.
        /// </summary>
        [Input("memberEmailAddress", required: true)]
        public Input<string> MemberEmailAddress { get; set; } = null!;

        /// <summary>
        /// The AWS account ID to be invited to join the graph as a member
        /// </summary>
        [Input("memberId", required: true)]
        public Input<string> MemberId { get; set; } = null!;

        /// <summary>
        /// A message to be included in the email invitation sent to the invited account. Updating this field has no effect.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        public MemberInvitationArgs()
        {
        }
        public static new MemberInvitationArgs Empty => new MemberInvitationArgs();
    }
}
