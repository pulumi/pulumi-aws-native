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
    /// Model inference configuration
    /// </summary>
    [OutputType]
    public sealed class PromptVersionPromptInferenceConfigurationProperties
    {
        public readonly Outputs.PromptVersionPromptModelInferenceConfiguration Text;

        [OutputConstructor]
        private PromptVersionPromptInferenceConfigurationProperties(Outputs.PromptVersionPromptModelInferenceConfiguration text)
        {
            Text = text;
        }
    }
}