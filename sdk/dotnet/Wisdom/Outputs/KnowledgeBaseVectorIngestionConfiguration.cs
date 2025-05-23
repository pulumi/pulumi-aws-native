// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Wisdom.Outputs
{

    [OutputType]
    public sealed class KnowledgeBaseVectorIngestionConfiguration
    {
        /// <summary>
        /// Details about how to chunk the documents in the data source. A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried.
        /// </summary>
        public readonly Outputs.KnowledgeBaseVectorIngestionConfigurationChunkingConfigurationProperties? ChunkingConfiguration;
        /// <summary>
        /// A custom parser for data source documents.
        /// </summary>
        public readonly Outputs.KnowledgeBaseVectorIngestionConfigurationParsingConfigurationProperties? ParsingConfiguration;

        [OutputConstructor]
        private KnowledgeBaseVectorIngestionConfiguration(
            Outputs.KnowledgeBaseVectorIngestionConfigurationChunkingConfigurationProperties? chunkingConfiguration,

            Outputs.KnowledgeBaseVectorIngestionConfigurationParsingConfigurationProperties? parsingConfiguration)
        {
            ChunkingConfiguration = chunkingConfiguration;
            ParsingConfiguration = parsingConfiguration;
        }
    }
}
