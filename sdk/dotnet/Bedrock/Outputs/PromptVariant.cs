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
    /// Prompt variant
    /// </summary>
    [OutputType]
    public sealed class PromptVariant
    {
        /// <summary>
        /// Contains inference configurations for the prompt variant.
        /// </summary>
        public readonly Outputs.PromptInferenceConfigurationProperties? InferenceConfiguration;
        /// <summary>
        /// ARN or name of a Bedrock model.
        /// </summary>
        public readonly string? ModelId;
        /// <summary>
        /// Name for a variant.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Contains configurations for the prompt template.
        /// </summary>
        public readonly Outputs.PromptTemplateConfigurationProperties? TemplateConfiguration;
        /// <summary>
        /// The type of prompt template to use.
        /// </summary>
        public readonly Pulumi.AwsNative.Bedrock.PromptTemplateType TemplateType;

        [OutputConstructor]
        private PromptVariant(
            Outputs.PromptInferenceConfigurationProperties? inferenceConfiguration,

            string? modelId,

            string name,

            Outputs.PromptTemplateConfigurationProperties? templateConfiguration,

            Pulumi.AwsNative.Bedrock.PromptTemplateType templateType)
        {
            InferenceConfiguration = inferenceConfiguration;
            ModelId = modelId;
            Name = name;
            TemplateConfiguration = templateConfiguration;
            TemplateType = templateType;
        }
    }
}