// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Outputs
{

    /// <summary>
    /// Contains details about the configuration of a Amazon Bedrock knowledge base.
    /// </summary>
    [OutputType]
    public sealed class BotQnAIntentConfigurationDataSourceConfigurationPropertiesBedrockKnowledgeStoreConfigurationProperties
    {
        /// <summary>
        /// The base ARN of the knowledge base used.
        /// </summary>
        public readonly string? BedrockKnowledgeBaseArn;
        /// <summary>
        /// Contains the names of the fields used for an exact response to the user.
        /// </summary>
        public readonly Outputs.BotQnAIntentConfigurationDataSourceConfigurationPropertiesBedrockKnowledgeStoreConfigurationPropertiesBkbExactResponseFieldsProperties? BkbExactResponseFields;
        /// <summary>
        /// Specifies whether to return an exact response, or to return an answer generated by the model, using the fields you specify from the database.
        /// </summary>
        public readonly bool? ExactResponse;

        [OutputConstructor]
        private BotQnAIntentConfigurationDataSourceConfigurationPropertiesBedrockKnowledgeStoreConfigurationProperties(
            string? bedrockKnowledgeBaseArn,

            Outputs.BotQnAIntentConfigurationDataSourceConfigurationPropertiesBedrockKnowledgeStoreConfigurationPropertiesBkbExactResponseFieldsProperties? bkbExactResponseFields,

            bool? exactResponse)
        {
            BedrockKnowledgeBaseArn = bedrockKnowledgeBaseArn;
            BkbExactResponseFields = bkbExactResponseFields;
            ExactResponse = exactResponse;
        }
    }
}
