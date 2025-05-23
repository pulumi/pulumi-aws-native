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
    /// Configurations for when you choose semantic chunking. If you set the chunkingStrategy as NONE, exclude this field.
    /// </summary>
    public sealed class DataSourceSemanticChunkingConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The dissimilarity threshold for splitting chunks.
        /// </summary>
        [Input("breakpointPercentileThreshold", required: true)]
        public Input<int> BreakpointPercentileThreshold { get; set; } = null!;

        /// <summary>
        /// The buffer size.
        /// </summary>
        [Input("bufferSize", required: true)]
        public Input<int> BufferSize { get; set; } = null!;

        /// <summary>
        /// The maximum number of tokens that a chunk can contain.
        /// </summary>
        [Input("maxTokens", required: true)]
        public Input<int> MaxTokens { get; set; } = null!;

        public DataSourceSemanticChunkingConfigurationArgs()
        {
        }
        public static new DataSourceSemanticChunkingConfigurationArgs Empty => new DataSourceSemanticChunkingConfigurationArgs();
    }
}
