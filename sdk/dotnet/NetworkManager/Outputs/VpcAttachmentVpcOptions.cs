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
    /// Vpc options of the attachment.
    /// </summary>
    [OutputType]
    public sealed class VpcAttachmentVpcOptions
    {
        /// <summary>
        /// Indicates whether to enable ApplianceModeSupport Support for Vpc Attachment. Valid Values: true | false
        /// </summary>
        public readonly bool? ApplianceModeSupport;
        /// <summary>
        /// Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable
        /// </summary>
        public readonly bool? Ipv6Support;

        [OutputConstructor]
        private VpcAttachmentVpcOptions(
            bool? applianceModeSupport,

            bool? ipv6Support)
        {
            ApplianceModeSupport = applianceModeSupport;
            Ipv6Support = ipv6Support;
        }
    }
}
