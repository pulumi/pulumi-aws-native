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
    /// CachePointBlock
    /// </summary>
    public sealed class PromptCachePointBlockArgs : global::Pulumi.ResourceArgs
    {
        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.Bedrock.PromptCachePointType> Type { get; set; } = null!;

        public PromptCachePointBlockArgs()
        {
        }
        public static new PromptCachePointBlockArgs Empty => new PromptCachePointBlockArgs();
    }
}
