// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Wisdom.Inputs
{

    public sealed class KnowledgeBaseVectorIngestionConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Details about how to chunk the documents in the data source. A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried.
        /// </summary>
        [Input("chunkingConfiguration")]
        public Input<Inputs.KnowledgeBaseVectorIngestionConfigurationChunkingConfigurationPropertiesArgs>? ChunkingConfiguration { get; set; }

        /// <summary>
        /// A custom parser for data source documents.
        /// </summary>
        [Input("parsingConfiguration")]
        public Input<Inputs.KnowledgeBaseVectorIngestionConfigurationParsingConfigurationPropertiesArgs>? ParsingConfiguration { get; set; }

        public KnowledgeBaseVectorIngestionConfigurationArgs()
        {
        }
        public static new KnowledgeBaseVectorIngestionConfigurationArgs Empty => new KnowledgeBaseVectorIngestionConfigurationArgs();
    }
}
