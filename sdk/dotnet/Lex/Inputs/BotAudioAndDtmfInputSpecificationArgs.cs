// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Inputs
{

    public sealed class BotAudioAndDtmfInputSpecificationArgs : global::Pulumi.ResourceArgs
    {
        [Input("audioSpecification")]
        public Input<Inputs.BotAudioSpecificationArgs>? AudioSpecification { get; set; }

        [Input("dtmfSpecification")]
        public Input<Inputs.BotDtmfSpecificationArgs>? DtmfSpecification { get; set; }

        [Input("startTimeoutMs", required: true)]
        public Input<int> StartTimeoutMs { get; set; } = null!;

        public BotAudioAndDtmfInputSpecificationArgs()
        {
        }
        public static new BotAudioAndDtmfInputSpecificationArgs Empty => new BotAudioAndDtmfInputSpecificationArgs();
    }
}
