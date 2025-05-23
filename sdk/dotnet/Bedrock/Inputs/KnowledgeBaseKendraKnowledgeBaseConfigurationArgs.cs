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
    /// Configurations for a Kendra knowledge base
    /// </summary>
    public sealed class KnowledgeBaseKendraKnowledgeBaseConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the Amazon Kendra index.
        /// </summary>
        [Input("kendraIndexArn", required: true)]
        public Input<string> KendraIndexArn { get; set; } = null!;

        public KnowledgeBaseKendraKnowledgeBaseConfigurationArgs()
        {
        }
        public static new KnowledgeBaseKendraKnowledgeBaseConfigurationArgs Empty => new KnowledgeBaseKendraKnowledgeBaseConfigurationArgs();
    }
}
