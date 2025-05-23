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
    /// Configurations for when you choose hierarchical chunking. If you set the chunkingStrategy as NONE, exclude this field.
    /// </summary>
    [OutputType]
    public sealed class DataSourceHierarchicalChunkingConfiguration
    {
        /// <summary>
        /// Token settings for each layer.
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSourceHierarchicalChunkingLevelConfiguration> LevelConfigurations;
        /// <summary>
        /// The number of tokens to repeat across chunks in the same layer.
        /// </summary>
        public readonly int OverlapTokens;

        [OutputConstructor]
        private DataSourceHierarchicalChunkingConfiguration(
            ImmutableArray<Outputs.DataSourceHierarchicalChunkingLevelConfiguration> levelConfigurations,

            int overlapTokens)
        {
            LevelConfigurations = levelConfigurations;
            OverlapTokens = overlapTokens;
        }
    }
}
