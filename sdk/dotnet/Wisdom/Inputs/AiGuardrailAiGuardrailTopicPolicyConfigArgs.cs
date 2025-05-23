// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Wisdom.Inputs
{

    /// <summary>
    /// Topic policy config for a guardrail.
    /// </summary>
    public sealed class AiGuardrailAiGuardrailTopicPolicyConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("topicsConfig", required: true)]
        private InputList<Inputs.AiGuardrailGuardrailTopicConfigArgs>? _topicsConfig;

        /// <summary>
        /// List of topic configs in topic policy.
        /// </summary>
        public InputList<Inputs.AiGuardrailGuardrailTopicConfigArgs> TopicsConfig
        {
            get => _topicsConfig ?? (_topicsConfig = new InputList<Inputs.AiGuardrailGuardrailTopicConfigArgs>());
            set => _topicsConfig = value;
        }

        public AiGuardrailAiGuardrailTopicPolicyConfigArgs()
        {
        }
        public static new AiGuardrailAiGuardrailTopicPolicyConfigArgs Empty => new AiGuardrailAiGuardrailTopicPolicyConfigArgs();
    }
}
