// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Inputs
{

    /// <summary>
    /// The vector store service in which the knowledge base is stored.
    /// </summary>
    public sealed class KnowledgeBaseStorageConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("opensearchServerlessConfiguration")]
        public Input<Inputs.KnowledgeBaseOpenSearchServerlessConfigurationArgs>? OpensearchServerlessConfiguration { get; set; }

        [Input("pineconeConfiguration")]
        public Input<Inputs.KnowledgeBasePineconeConfigurationArgs>? PineconeConfiguration { get; set; }

        [Input("rdsConfiguration")]
        public Input<Inputs.KnowledgeBaseRdsConfigurationArgs>? RdsConfiguration { get; set; }

        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.Bedrock.KnowledgeBaseStorageType> Type { get; set; } = null!;

        public KnowledgeBaseStorageConfigurationArgs()
        {
        }
        public static new KnowledgeBaseStorageConfigurationArgs Empty => new KnowledgeBaseStorageConfigurationArgs();
    }
}