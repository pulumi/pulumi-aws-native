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
    /// CachePointBlock
    /// </summary>
    [OutputType]
    public sealed class PromptVersionCachePointBlock
    {
        public readonly Pulumi.AwsNative.Bedrock.PromptVersionCachePointType Type;

        [OutputConstructor]
        private PromptVersionCachePointBlock(Pulumi.AwsNative.Bedrock.PromptVersionCachePointType type)
        {
            Type = type;
        }
    }
}
