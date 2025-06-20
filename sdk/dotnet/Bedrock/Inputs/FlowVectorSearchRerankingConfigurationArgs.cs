// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Inputs
{

    public sealed class FlowVectorSearchRerankingConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("bedrockRerankingConfiguration")]
        public Input<Inputs.FlowVectorSearchBedrockRerankingConfigurationArgs>? BedrockRerankingConfiguration { get; set; }

        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.Bedrock.FlowVectorSearchRerankingConfigurationType> Type { get; set; } = null!;

        public FlowVectorSearchRerankingConfigurationArgs()
        {
        }
        public static new FlowVectorSearchRerankingConfigurationArgs Empty => new FlowVectorSearchRerankingConfigurationArgs();
    }
}
