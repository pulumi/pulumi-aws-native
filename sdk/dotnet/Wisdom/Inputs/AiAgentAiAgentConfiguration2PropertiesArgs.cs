// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Wisdom.Inputs
{

    public sealed class AiAgentAiAgentConfiguration2PropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("selfServiceAiAgentConfiguration", required: true)]
        public Input<Inputs.AiAgentSelfServiceAiAgentConfigurationArgs> SelfServiceAiAgentConfiguration { get; set; } = null!;

        public AiAgentAiAgentConfiguration2PropertiesArgs()
        {
        }
        public static new AiAgentAiAgentConfiguration2PropertiesArgs Empty => new AiAgentAiAgentConfiguration2PropertiesArgs();
    }
}
