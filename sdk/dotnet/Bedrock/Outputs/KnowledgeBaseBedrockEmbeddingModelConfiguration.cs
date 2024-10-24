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
    /// The vector configuration details for the Bedrock embeddings model.
    /// </summary>
    [OutputType]
    public sealed class KnowledgeBaseBedrockEmbeddingModelConfiguration
    {
        /// <summary>
        /// The dimensions details for the vector configuration used on the Bedrock embeddings model.
        /// </summary>
        public readonly int? Dimensions;

        [OutputConstructor]
        private KnowledgeBaseBedrockEmbeddingModelConfiguration(int? dimensions)
        {
            Dimensions = dimensions;
        }
    }
}
