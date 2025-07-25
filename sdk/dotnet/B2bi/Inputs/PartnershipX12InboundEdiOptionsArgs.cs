// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.B2bi.Inputs
{

    public sealed class PartnershipX12InboundEdiOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies acknowledgment options for inbound X12 EDI files. These options control how functional and technical acknowledgments are handled.
        /// </summary>
        [Input("acknowledgmentOptions")]
        public Input<Inputs.PartnershipX12AcknowledgmentOptionsArgs>? AcknowledgmentOptions { get; set; }

        public PartnershipX12InboundEdiOptionsArgs()
        {
        }
        public static new PartnershipX12InboundEdiOptionsArgs Empty => new PartnershipX12InboundEdiOptionsArgs();
    }
}
