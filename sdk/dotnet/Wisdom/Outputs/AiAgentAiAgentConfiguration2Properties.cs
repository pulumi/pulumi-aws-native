// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Wisdom.Outputs
{

    [OutputType]
    public sealed class AiAgentAiAgentConfiguration2Properties
    {
        public readonly Outputs.AiAgentSelfServiceAiAgentConfiguration SelfServiceAiAgentConfiguration;

        [OutputConstructor]
        private AiAgentAiAgentConfiguration2Properties(Outputs.AiAgentSelfServiceAiAgentConfiguration selfServiceAiAgentConfiguration)
        {
            SelfServiceAiAgentConfiguration = selfServiceAiAgentConfiguration;
        }
    }
}
