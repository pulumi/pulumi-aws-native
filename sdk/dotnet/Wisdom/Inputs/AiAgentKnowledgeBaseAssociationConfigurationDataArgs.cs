// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Wisdom.Inputs
{

    public sealed class AiAgentKnowledgeBaseAssociationConfigurationDataArgs : global::Pulumi.ResourceArgs
    {
        [Input("contentTagFilter")]
        public object? ContentTagFilter { get; set; }

        [Input("maxResults")]
        public Input<double>? MaxResults { get; set; }

        [Input("overrideKnowledgeBaseSearchType")]
        public Input<Pulumi.AwsNative.Wisdom.AiAgentKnowledgeBaseSearchType>? OverrideKnowledgeBaseSearchType { get; set; }

        public AiAgentKnowledgeBaseAssociationConfigurationDataArgs()
        {
        }
        public static new AiAgentKnowledgeBaseAssociationConfigurationDataArgs Empty => new AiAgentKnowledgeBaseAssociationConfigurationDataArgs();
    }
}
