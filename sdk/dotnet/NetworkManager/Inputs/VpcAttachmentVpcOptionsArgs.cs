// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkManager.Inputs
{

    /// <summary>
    /// Vpc options of the attachment.
    /// </summary>
    public sealed class VpcAttachmentVpcOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether to enable ApplianceModeSupport Support for Vpc Attachment. Valid Values: true | false
        /// </summary>
        [Input("applianceModeSupport")]
        public Input<bool>? ApplianceModeSupport { get; set; }

        /// <summary>
        /// Indicates whether to enable private DNS Support for Vpc Attachment. Valid Values: true | false
        /// </summary>
        [Input("dnsSupport")]
        public Input<bool>? DnsSupport { get; set; }

        /// <summary>
        /// Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable
        /// </summary>
        [Input("ipv6Support")]
        public Input<bool>? Ipv6Support { get; set; }

        /// <summary>
        /// Indicates whether to enable Security Group Referencing Support for Vpc Attachment. Valid Values: true | false
        /// </summary>
        [Input("securityGroupReferencingSupport")]
        public Input<bool>? SecurityGroupReferencingSupport { get; set; }

        public VpcAttachmentVpcOptionsArgs()
        {
        }
        public static new VpcAttachmentVpcOptionsArgs Empty => new VpcAttachmentVpcOptionsArgs();
    }
}
