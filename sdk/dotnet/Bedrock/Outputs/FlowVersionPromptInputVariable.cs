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
    /// Input variable
    /// </summary>
    [OutputType]
    public sealed class FlowVersionPromptInputVariable
    {
        /// <summary>
        /// Name for an input variable
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private FlowVersionPromptInputVariable(string? name)
        {
            Name = name;
        }
    }
}
