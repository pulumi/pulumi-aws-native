// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.B2bi.Inputs
{

    public sealed class PartnershipX12EnvelopeArgs : global::Pulumi.ResourceArgs
    {
        [Input("common")]
        public Input<Inputs.PartnershipX12OutboundEdiHeadersArgs>? Common { get; set; }

        public PartnershipX12EnvelopeArgs()
        {
        }
        public static new PartnershipX12EnvelopeArgs Empty => new PartnershipX12EnvelopeArgs();
    }
}