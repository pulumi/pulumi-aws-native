// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.B2bi.Outputs
{

    [OutputType]
    public sealed class PartnershipOutboundEdiOptionsProperties
    {
        public readonly Outputs.PartnershipX12Envelope X12;

        [OutputConstructor]
        private PartnershipOutboundEdiOptionsProperties(Outputs.PartnershipX12Envelope x12)
        {
            X12 = x12;
        }
    }
}
