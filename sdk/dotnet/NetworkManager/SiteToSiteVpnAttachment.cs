// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkManager
{
    /// <summary>
    /// AWS::NetworkManager::SiteToSiteVpnAttachment Resource Type definition.
    /// </summary>
    [AwsNativeResourceType("aws-native:networkmanager:SiteToSiteVpnAttachment")]
    public partial class SiteToSiteVpnAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the attachment.
        /// </summary>
        [Output("attachmentId")]
        public Output<string> AttachmentId { get; private set; } = null!;

        /// <summary>
        /// The policy rule number associated with the attachment.
        /// </summary>
        [Output("attachmentPolicyRuleNumber")]
        public Output<int> AttachmentPolicyRuleNumber { get; private set; } = null!;

        /// <summary>
        /// The type of attachment.
        /// </summary>
        [Output("attachmentType")]
        public Output<string> AttachmentType { get; private set; } = null!;

        /// <summary>
        /// The ARN of a core network for the VPC attachment.
        /// </summary>
        [Output("coreNetworkArn")]
        public Output<string> CoreNetworkArn { get; private set; } = null!;

        /// <summary>
        /// The ID of a core network where you're creating a site-to-site VPN attachment.
        /// </summary>
        [Output("coreNetworkId")]
        public Output<string> CoreNetworkId { get; private set; } = null!;

        /// <summary>
        /// Creation time of the attachment.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The Region where the edge is located.
        /// </summary>
        [Output("edgeLocation")]
        public Output<string> EdgeLocation { get; private set; } = null!;

        /// <summary>
        /// The name of the network function group attachment.
        /// </summary>
        [Output("networkFunctionGroupName")]
        public Output<string?> NetworkFunctionGroupName { get; private set; } = null!;

        /// <summary>
        /// Owner account of the attachment.
        /// </summary>
        [Output("ownerAccountId")]
        public Output<string> OwnerAccountId { get; private set; } = null!;

        /// <summary>
        /// The attachment to move from one network function group to another.
        /// </summary>
        [Output("proposedNetworkFunctionGroupChange")]
        public Output<Outputs.SiteToSiteVpnAttachmentProposedNetworkFunctionGroupChange?> ProposedNetworkFunctionGroupChange { get; private set; } = null!;

        /// <summary>
        /// The attachment to move from one segment to another.
        /// </summary>
        [Output("proposedSegmentChange")]
        public Output<Outputs.SiteToSiteVpnAttachmentProposedSegmentChange?> ProposedSegmentChange { get; private set; } = null!;

        /// <summary>
        /// The ARN of the Resource.
        /// </summary>
        [Output("resourceArn")]
        public Output<string> ResourceArn { get; private set; } = null!;

        /// <summary>
        /// The name of the segment that attachment is in.
        /// </summary>
        [Output("segmentName")]
        public Output<string> SegmentName { get; private set; } = null!;

        /// <summary>
        /// The state of the attachment.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Tags for the attachment.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Last update time of the attachment.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The ARN of the site-to-site VPN attachment.
        /// </summary>
        [Output("vpnConnectionArn")]
        public Output<string> VpnConnectionArn { get; private set; } = null!;


        /// <summary>
        /// Create a SiteToSiteVpnAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SiteToSiteVpnAttachment(string name, SiteToSiteVpnAttachmentArgs args, CustomResourceOptions? options = null)
            : base("aws-native:networkmanager:SiteToSiteVpnAttachment", name, args ?? new SiteToSiteVpnAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SiteToSiteVpnAttachment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:networkmanager:SiteToSiteVpnAttachment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "coreNetworkId",
                    "vpnConnectionArn",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SiteToSiteVpnAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SiteToSiteVpnAttachment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SiteToSiteVpnAttachment(name, id, options);
        }
    }

    public sealed class SiteToSiteVpnAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of a core network where you're creating a site-to-site VPN attachment.
        /// </summary>
        [Input("coreNetworkId", required: true)]
        public Input<string> CoreNetworkId { get; set; } = null!;

        /// <summary>
        /// The name of the network function group attachment.
        /// </summary>
        [Input("networkFunctionGroupName")]
        public Input<string>? NetworkFunctionGroupName { get; set; }

        /// <summary>
        /// The attachment to move from one network function group to another.
        /// </summary>
        [Input("proposedNetworkFunctionGroupChange")]
        public Input<Inputs.SiteToSiteVpnAttachmentProposedNetworkFunctionGroupChangeArgs>? ProposedNetworkFunctionGroupChange { get; set; }

        /// <summary>
        /// The attachment to move from one segment to another.
        /// </summary>
        [Input("proposedSegmentChange")]
        public Input<Inputs.SiteToSiteVpnAttachmentProposedSegmentChangeArgs>? ProposedSegmentChange { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// Tags for the attachment.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The ARN of the site-to-site VPN attachment.
        /// </summary>
        [Input("vpnConnectionArn", required: true)]
        public Input<string> VpnConnectionArn { get; set; } = null!;

        public SiteToSiteVpnAttachmentArgs()
        {
        }
        public static new SiteToSiteVpnAttachmentArgs Empty => new SiteToSiteVpnAttachmentArgs();
    }
}
