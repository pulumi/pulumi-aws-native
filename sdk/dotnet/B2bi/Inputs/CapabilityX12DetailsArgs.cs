// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.B2bi.Inputs
{

    public sealed class CapabilityX12DetailsArgs : global::Pulumi.ResourceArgs
    {
        [Input("transactionSet")]
        public Input<Pulumi.AwsNative.B2bi.CapabilityX12TransactionSet>? TransactionSet { get; set; }

        [Input("version")]
        public Input<Pulumi.AwsNative.B2bi.CapabilityX12Version>? Version { get; set; }

        public CapabilityX12DetailsArgs()
        {
        }
        public static new CapabilityX12DetailsArgs Empty => new CapabilityX12DetailsArgs();
    }
}
