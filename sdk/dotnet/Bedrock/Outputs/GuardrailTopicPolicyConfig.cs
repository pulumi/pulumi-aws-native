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
    /// Topic policy config for a guardrail.
    /// </summary>
    [OutputType]
    public sealed class GuardrailTopicPolicyConfig
    {
        /// <summary>
        /// List of topic configs in topic policy.
        /// </summary>
        public readonly ImmutableArray<Outputs.GuardrailTopicConfig> TopicsConfig;

        [OutputConstructor]
        private GuardrailTopicPolicyConfig(ImmutableArray<Outputs.GuardrailTopicConfig> topicsConfig)
        {
            TopicsConfig = topicsConfig;
        }
    }
}
