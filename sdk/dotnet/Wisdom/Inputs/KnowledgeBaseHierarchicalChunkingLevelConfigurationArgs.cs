// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Wisdom.Inputs
{

    public sealed class KnowledgeBaseHierarchicalChunkingLevelConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum number of tokens that a chunk can contain in this layer.
        /// </summary>
        [Input("maxTokens", required: true)]
        public Input<double> MaxTokens { get; set; } = null!;

        public KnowledgeBaseHierarchicalChunkingLevelConfigurationArgs()
        {
        }
        public static new KnowledgeBaseHierarchicalChunkingLevelConfigurationArgs Empty => new KnowledgeBaseHierarchicalChunkingLevelConfigurationArgs();
    }
}
