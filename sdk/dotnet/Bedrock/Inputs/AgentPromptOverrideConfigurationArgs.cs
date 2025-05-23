// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Inputs
{

    /// <summary>
    /// Configuration for prompt override.
    /// </summary>
    public sealed class AgentPromptOverrideConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of a Lambda.
        /// </summary>
        [Input("overrideLambda")]
        public Input<string>? OverrideLambda { get; set; }

        [Input("promptConfigurations", required: true)]
        private InputList<Inputs.AgentPromptConfigurationArgs>? _promptConfigurations;

        /// <summary>
        /// List of BasePromptConfiguration
        /// </summary>
        public InputList<Inputs.AgentPromptConfigurationArgs> PromptConfigurations
        {
            get => _promptConfigurations ?? (_promptConfigurations = new InputList<Inputs.AgentPromptConfigurationArgs>());
            set => _promptConfigurations = value;
        }

        public AgentPromptOverrideConfigurationArgs()
        {
        }
        public static new AgentPromptOverrideConfigurationArgs Empty => new AgentPromptOverrideConfigurationArgs();
    }
}
