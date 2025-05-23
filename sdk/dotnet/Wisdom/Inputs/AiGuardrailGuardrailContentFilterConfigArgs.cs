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
    /// Content filter config in content policy.
    /// </summary>
    public sealed class AiGuardrailGuardrailContentFilterConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The strength of the input for the guardrail content filter.
        /// </summary>
        [Input("inputStrength", required: true)]
        public Input<Pulumi.AwsNative.Wisdom.AiGuardrailGuardrailFilterStrength> InputStrength { get; set; } = null!;

        /// <summary>
        /// The output strength of the guardrail content filter.
        /// </summary>
        [Input("outputStrength", required: true)]
        public Input<Pulumi.AwsNative.Wisdom.AiGuardrailGuardrailFilterStrength> OutputStrength { get; set; } = null!;

        /// <summary>
        /// The type of the guardrail content filter.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.Wisdom.AiGuardrailGuardrailContentFilterType> Type { get; set; } = null!;

        public AiGuardrailGuardrailContentFilterConfigArgs()
        {
        }
        public static new AiGuardrailGuardrailContentFilterConfigArgs Empty => new AiGuardrailGuardrailContentFilterConfigArgs();
    }
}
