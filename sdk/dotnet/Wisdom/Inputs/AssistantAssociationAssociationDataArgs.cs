// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Wisdom.Inputs
{

    public sealed class AssistantAssociationAssociationDataArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identifier of the knowledge base.
        /// </summary>
        [Input("knowledgeBaseId", required: true)]
        public Input<string> KnowledgeBaseId { get; set; } = null!;

        public AssistantAssociationAssociationDataArgs()
        {
        }
        public static new AssistantAssociationAssociationDataArgs Empty => new AssistantAssociationAssociationDataArgs();
    }
}
