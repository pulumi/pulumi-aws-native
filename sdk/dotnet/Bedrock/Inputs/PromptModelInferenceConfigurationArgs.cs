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
    /// Prompt model inference configuration
    /// </summary>
    public sealed class PromptModelInferenceConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum length of output
        /// </summary>
        [Input("maxTokens")]
        public Input<double>? MaxTokens { get; set; }

        [Input("stopSequences")]
        private InputList<string>? _stopSequences;

        /// <summary>
        /// List of stop sequences
        /// </summary>
        public InputList<string> StopSequences
        {
            get => _stopSequences ?? (_stopSequences = new InputList<string>());
            set => _stopSequences = value;
        }

        /// <summary>
        /// Controls randomness, higher values increase diversity
        /// </summary>
        [Input("temperature")]
        public Input<double>? Temperature { get; set; }

        /// <summary>
        /// Cumulative probability cutoff for token selection
        /// </summary>
        [Input("topP")]
        public Input<double>? TopP { get; set; }

        public PromptModelInferenceConfigurationArgs()
        {
        }
        public static new PromptModelInferenceConfigurationArgs Empty => new PromptModelInferenceConfigurationArgs();
    }
}
