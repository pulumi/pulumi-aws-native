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
    /// Content policy config for a guardrail.
    /// </summary>
    [OutputType]
    public sealed class GuardrailContentPolicyConfig
    {
        /// <summary>
        /// List of content filter configs in content policy.
        /// </summary>
        public readonly ImmutableArray<Outputs.GuardrailContentFilterConfig> FiltersConfig;

        [OutputConstructor]
        private GuardrailContentPolicyConfig(ImmutableArray<Outputs.GuardrailContentFilterConfig> filtersConfig)
        {
            FiltersConfig = filtersConfig;
        }
    }
}
