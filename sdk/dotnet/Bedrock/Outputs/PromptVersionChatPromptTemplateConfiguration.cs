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
    /// Configuration for chat prompt template
    /// </summary>
    [OutputType]
    public sealed class PromptVersionChatPromptTemplateConfiguration
    {
        /// <summary>
        /// List of input variables
        /// </summary>
        public readonly ImmutableArray<Outputs.PromptVersionPromptInputVariable> InputVariables;
        /// <summary>
        /// List of messages for chat prompt template
        /// </summary>
        public readonly ImmutableArray<Outputs.PromptVersionMessage> Messages;
        /// <summary>
        /// Configuration for chat prompt template
        /// </summary>
        public readonly ImmutableArray<Union<Outputs.PromptVersionSystemContentBlock0Properties, Outputs.PromptVersionSystemContentBlock1Properties>> System;
        public readonly Outputs.PromptVersionToolConfiguration? ToolConfiguration;

        [OutputConstructor]
        private PromptVersionChatPromptTemplateConfiguration(
            ImmutableArray<Outputs.PromptVersionPromptInputVariable> inputVariables,

            ImmutableArray<Outputs.PromptVersionMessage> messages,

            ImmutableArray<Union<Outputs.PromptVersionSystemContentBlock0Properties, Outputs.PromptVersionSystemContentBlock1Properties>> system,

            Outputs.PromptVersionToolConfiguration? toolConfiguration)
        {
            InputVariables = inputVariables;
            Messages = messages;
            System = system;
            ToolConfiguration = toolConfiguration;
        }
    }
}
