// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.B2bi.Outputs
{

    [OutputType]
    public sealed class PartnershipX12OutboundEdiHeaders
    {
        public readonly Outputs.PartnershipX12Delimiters? Delimiters;
        public readonly Outputs.PartnershipX12FunctionalGroupHeaders? FunctionalGroupHeaders;
        public readonly Outputs.PartnershipX12InterchangeControlHeaders? InterchangeControlHeaders;
        public readonly bool? ValidateEdi;

        [OutputConstructor]
        private PartnershipX12OutboundEdiHeaders(
            Outputs.PartnershipX12Delimiters? delimiters,

            Outputs.PartnershipX12FunctionalGroupHeaders? functionalGroupHeaders,

            Outputs.PartnershipX12InterchangeControlHeaders? interchangeControlHeaders,

            bool? validateEdi)
        {
            Delimiters = delimiters;
            FunctionalGroupHeaders = functionalGroupHeaders;
            InterchangeControlHeaders = interchangeControlHeaders;
            ValidateEdi = validateEdi;
        }
    }
}