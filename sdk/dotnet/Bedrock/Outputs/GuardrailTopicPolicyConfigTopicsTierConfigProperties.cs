// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Outputs
{

    /// <summary>
    /// Guardrail tier config for topic policy
    /// </summary>
    [OutputType]
    public sealed class GuardrailTopicPolicyConfigTopicsTierConfigProperties
    {
        public readonly Pulumi.AwsNative.Bedrock.GuardrailTopicsTierName TierName;

        [OutputConstructor]
        private GuardrailTopicPolicyConfigTopicsTierConfigProperties(Pulumi.AwsNative.Bedrock.GuardrailTopicsTierName tierName)
        {
            TierName = tierName;
        }
    }
}
