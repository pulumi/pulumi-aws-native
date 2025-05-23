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
    /// Chat prompt Message
    /// </summary>
    [OutputType]
    public sealed class PromptVersionMessage
    {
        /// <summary>
        /// List of Content Blocks
        /// </summary>
        public readonly ImmutableArray<Union<Outputs.PromptVersionContentBlock0Properties, Outputs.PromptVersionContentBlock1Properties>> Content;
        public readonly Pulumi.AwsNative.Bedrock.PromptVersionConversationRole Role;

        [OutputConstructor]
        private PromptVersionMessage(
            ImmutableArray<Union<Outputs.PromptVersionContentBlock0Properties, Outputs.PromptVersionContentBlock1Properties>> content,

            Pulumi.AwsNative.Bedrock.PromptVersionConversationRole role)
        {
            Content = content;
            Role = role;
        }
    }
}
