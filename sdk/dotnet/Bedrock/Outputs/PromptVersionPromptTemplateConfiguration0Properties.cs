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
    /// Prompt template configuration
    /// </summary>
    [OutputType]
    public sealed class PromptVersionPromptTemplateConfiguration0Properties
    {
        public readonly Outputs.PromptVersionTextPromptTemplateConfiguration Text;

        [OutputConstructor]
        private PromptVersionPromptTemplateConfiguration0Properties(Outputs.PromptVersionTextPromptTemplateConfiguration text)
        {
            Text = text;
        }
    }
}
