// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkManager
{
    public static class GetVpcAttachment
    {
        /// <summary>
        /// AWS::NetworkManager::VpcAttachment Resoruce Type
        /// </summary>
        public static Task<GetVpcAttachmentResult> InvokeAsync(GetVpcAttachmentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcAttachmentResult>("aws-native:networkmanager:getVpcAttachment", args ?? new GetVpcAttachmentArgs(), options.WithDefaults());

        /// <summary>
        /// AWS::NetworkManager::VpcAttachment Resoruce Type
        /// </summary>
        public static Output<GetVpcAttachmentResult> Invoke(GetVpcAttachmentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcAttachmentResult>("aws-native:networkmanager:getVpcAttachment", args ?? new GetVpcAttachmentInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// AWS::NetworkManager::VpcAttachment Resoruce Type
        /// </summary>
        public static Output<GetVpcAttachmentResult> Invoke(GetVpcAttachmentInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcAttachmentResult>("aws-native:networkmanager:getVpcAttachment", args ?? new GetVpcAttachmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcAttachmentArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Id of the attachment.
        /// </summary>
        [Input("attachmentId", required: true)]
        public string AttachmentId { get; set; } = null!;

        public GetVpcAttachmentArgs()
        {
        }
        public static new GetVpcAttachmentArgs Empty => new GetVpcAttachmentArgs();
    }

    public sealed class GetVpcAttachmentInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Id of the attachment.
        /// </summary>
        [Input("attachmentId", required: true)]
        public Input<string> AttachmentId { get; set; } = null!;

        public GetVpcAttachmentInvokeArgs()
        {
        }
        public static new GetVpcAttachmentInvokeArgs Empty => new GetVpcAttachmentInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcAttachmentResult
    {
        /// <summary>
        /// Id of the attachment.
        /// </summary>
        public readonly string? AttachmentId;
        /// <summary>
        /// The policy rule number associated with the attachment.
        /// </summary>
        public readonly int? AttachmentPolicyRuleNumber;
        /// <summary>
        /// Attachment type.
        /// </summary>
        public readonly string? AttachmentType;
        /// <summary>
        /// The ARN of a core network for the VPC attachment.
        /// </summary>
        public readonly string? CoreNetworkArn;
        /// <summary>
        /// Creation time of the attachment.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// The Region where the edge is located.
        /// </summary>
        public readonly string? EdgeLocation;
        /// <summary>
        /// The name of the network function group attachment.
        /// </summary>
        public readonly string? NetworkFunctionGroupName;
        /// <summary>
        /// Vpc options of the attachment.
        /// </summary>
        public readonly Outputs.VpcAttachmentVpcOptions? Options;
        /// <summary>
        /// Owner account of the attachment.
        /// </summary>
        public readonly string? OwnerAccountId;
        /// <summary>
        /// The attachment to move from one network function group to another.
        /// </summary>
        public readonly Outputs.VpcAttachmentProposedNetworkFunctionGroupChange? ProposedNetworkFunctionGroupChange;
        /// <summary>
        /// The attachment to move from one segment to another.
        /// </summary>
        public readonly Outputs.VpcAttachmentProposedSegmentChange? ProposedSegmentChange;
        /// <summary>
        /// The ARN of the Resource.
        /// </summary>
        public readonly string? ResourceArn;
        /// <summary>
        /// The name of the segment attachment..
        /// </summary>
        public readonly string? SegmentName;
        /// <summary>
        /// State of the attachment.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// Subnet Arn list
        /// </summary>
        public readonly ImmutableArray<string> SubnetArns;
        /// <summary>
        /// Tags for the attachment.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// Last update time of the attachment.
        /// </summary>
        public readonly string? UpdatedAt;

        [OutputConstructor]
        private GetVpcAttachmentResult(
            string? attachmentId,

            int? attachmentPolicyRuleNumber,

            string? attachmentType,

            string? coreNetworkArn,

            string? createdAt,

            string? edgeLocation,

            string? networkFunctionGroupName,

            Outputs.VpcAttachmentVpcOptions? options,

            string? ownerAccountId,

            Outputs.VpcAttachmentProposedNetworkFunctionGroupChange? proposedNetworkFunctionGroupChange,

            Outputs.VpcAttachmentProposedSegmentChange? proposedSegmentChange,

            string? resourceArn,

            string? segmentName,

            string? state,

            ImmutableArray<string> subnetArns,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            string? updatedAt)
        {
            AttachmentId = attachmentId;
            AttachmentPolicyRuleNumber = attachmentPolicyRuleNumber;
            AttachmentType = attachmentType;
            CoreNetworkArn = coreNetworkArn;
            CreatedAt = createdAt;
            EdgeLocation = edgeLocation;
            NetworkFunctionGroupName = networkFunctionGroupName;
            Options = options;
            OwnerAccountId = ownerAccountId;
            ProposedNetworkFunctionGroupChange = proposedNetworkFunctionGroupChange;
            ProposedSegmentChange = proposedSegmentChange;
            ResourceArn = resourceArn;
            SegmentName = segmentName;
            State = state;
            SubnetArns = subnetArns;
            Tags = tags;
            UpdatedAt = updatedAt;
        }
    }
}
