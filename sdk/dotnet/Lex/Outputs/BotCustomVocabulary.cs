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
    public sealed class BotCustomVocabulary
    {
        /// <summary>
        /// Specifies a list of words that you expect to be used during a conversation with your bot.
        /// </summary>
        public readonly ImmutableArray<Outputs.BotCustomVocabularyItem> CustomVocabularyItems;

        [OutputConstructor]
        private BotCustomVocabulary(ImmutableArray<Outputs.BotCustomVocabularyItem> customVocabularyItems)
        {
            CustomVocabularyItems = customVocabularyItems;
        }
    }
}
