// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Outputs
{

    /// <summary>
    /// Configuration for prompt override.
    /// </summary>
    [OutputType]
    public sealed class AgentPromptOverrideConfiguration
    {
        /// <summary>
        /// ARN of a Lambda.
        /// </summary>
        public readonly string? OverrideLambda;
        /// <summary>
        /// List of BasePromptConfiguration
        /// </summary>
        public readonly ImmutableArray<Outputs.AgentPromptConfiguration> PromptConfigurations;

        [OutputConstructor]
        private AgentPromptOverrideConfiguration(
            string? overrideLambda,

            ImmutableArray<Outputs.AgentPromptConfiguration> promptConfigurations)
        {
            OverrideLambda = overrideLambda;
            PromptConfigurations = promptConfigurations;
        }
    }
}