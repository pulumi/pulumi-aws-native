// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Outputs
{

    [OutputType]
    public sealed class BotQnAIntentConfigurationDataSourceConfigurationPropertiesBedrockKnowledgeStoreConfigurationProperties
    {
        public readonly string? BedrockKnowledgeBaseArn;
        public readonly Outputs.BotQnAIntentConfigurationDataSourceConfigurationPropertiesBedrockKnowledgeStoreConfigurationPropertiesBkbExactResponseFieldsProperties? BkbExactResponseFields;
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
