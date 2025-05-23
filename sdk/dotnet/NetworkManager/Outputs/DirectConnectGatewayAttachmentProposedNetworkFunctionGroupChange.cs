// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkManager.Outputs
{

    /// <summary>
    /// The attachment to move from one network function group to another.
    /// </summary>
    [OutputType]
    public sealed class DirectConnectGatewayAttachmentProposedNetworkFunctionGroupChange
    {
        /// <summary>
        /// The rule number in the policy document that applies to this change.
        /// </summary>
        public readonly int? AttachmentPolicyRuleNumber;
        /// <summary>
        /// The name of the network function group to change.
        /// </summary>
        public readonly string? NetworkFunctionGroupName;
        /// <summary>
        /// The key-value tags that changed for the network function group.
        /// </summary>
        public readonly ImmutableArray<Outputs.DirectConnectGatewayAttachmentTag> Tags;

        [OutputConstructor]
        private DirectConnectGatewayAttachmentProposedNetworkFunctionGroupChange(
            int? attachmentPolicyRuleNumber,

            string? networkFunctionGroupName,

            ImmutableArray<Outputs.DirectConnectGatewayAttachmentTag> tags)
        {
            AttachmentPolicyRuleNumber = attachmentPolicyRuleNumber;
            NetworkFunctionGroupName = networkFunctionGroupName;
            Tags = tags;
        }
    }
}
