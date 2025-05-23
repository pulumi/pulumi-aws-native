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
    /// A custom word config.
    /// </summary>
    public sealed class GuardrailWordConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("inputAction")]
        public Input<Pulumi.AwsNative.Bedrock.GuardrailWordAction>? InputAction { get; set; }

        [Input("inputEnabled")]
        public Input<bool>? InputEnabled { get; set; }

        [Input("outputAction")]
        public Input<Pulumi.AwsNative.Bedrock.GuardrailWordAction>? OutputAction { get; set; }

        [Input("outputEnabled")]
        public Input<bool>? OutputEnabled { get; set; }

        /// <summary>
        /// The custom word text.
        /// </summary>
        [Input("text", required: true)]
        public Input<string> Text { get; set; } = null!;

        public GuardrailWordConfigArgs()
        {
        }
        public static new GuardrailWordConfigArgs Empty => new GuardrailWordConfigArgs();
    }
}
