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
    /// A mapping of Bedrock Knowledge Base fields to OpenSearch Serverless field names
    /// </summary>
    [OutputType]
    public sealed class KnowledgeBaseOpenSearchServerlessFieldMapping
    {
        /// <summary>
        /// The name of the field in which Amazon Bedrock stores metadata about the vector store.
        /// </summary>
        public readonly string MetadataField;
        /// <summary>
        /// The name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.
        /// </summary>
        public readonly string TextField;
        /// <summary>
        /// The name of the field in which Amazon Bedrock stores the vector embeddings for your data sources.
        /// </summary>
        public readonly string VectorField;

        [OutputConstructor]
        private KnowledgeBaseOpenSearchServerlessFieldMapping(
            string metadataField,

            string textField,

            string vectorField)
        {
            MetadataField = metadataField;
            TextField = textField;
            VectorField = vectorField;
        }
    }
}