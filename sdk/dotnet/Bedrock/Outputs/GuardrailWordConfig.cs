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
    /// A custom word config.
    /// </summary>
    [OutputType]
    public sealed class GuardrailWordConfig
    {
        public readonly Pulumi.AwsNative.Bedrock.GuardrailWordAction? InputAction;
        public readonly bool? InputEnabled;
        public readonly Pulumi.AwsNative.Bedrock.GuardrailWordAction? OutputAction;
        public readonly bool? OutputEnabled;
        /// <summary>
        /// The custom word text.
        /// </summary>
        public readonly string Text;

        [OutputConstructor]
        private GuardrailWordConfig(
            Pulumi.AwsNative.Bedrock.GuardrailWordAction? inputAction,

            bool? inputEnabled,

            Pulumi.AwsNative.Bedrock.GuardrailWordAction? outputAction,

            bool? outputEnabled,

            string text)
        {
            InputAction = inputAction;
            InputEnabled = inputEnabled;
            OutputAction = outputAction;
            OutputEnabled = outputEnabled;
            Text = text;
        }
    }
}
