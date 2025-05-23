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
    /// Prompt flow node configuration
    /// </summary>
    [OutputType]
    public sealed class FlowPromptFlowNodeConfiguration
    {
        public readonly Outputs.FlowGuardrailConfiguration? GuardrailConfiguration;
        public readonly Union<Outputs.FlowPromptFlowNodeSourceConfiguration0Properties, Outputs.FlowPromptFlowNodeSourceConfiguration1Properties> SourceConfiguration;

        [OutputConstructor]
        private FlowPromptFlowNodeConfiguration(
            Outputs.FlowGuardrailConfiguration? guardrailConfiguration,

            Union<Outputs.FlowPromptFlowNodeSourceConfiguration0Properties, Outputs.FlowPromptFlowNodeSourceConfiguration1Properties> sourceConfiguration)
        {
            GuardrailConfiguration = guardrailConfiguration;
            SourceConfiguration = sourceConfiguration;
        }
    }
}
