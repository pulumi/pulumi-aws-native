// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Outputs
{

    /// <summary>
    /// A locale in the bot, which contains the intents and slot types that the bot uses in conversations with users in the specified language and locale.
    /// </summary>
    [OutputType]
    public sealed class BotLocale
    {
        public readonly Outputs.BotCustomVocabulary? CustomVocabulary;
        public readonly string? Description;
        /// <summary>
        /// List of intents
        /// </summary>
        public readonly ImmutableArray<Outputs.BotIntent> Intents;
        public readonly string LocaleId;
        public readonly double NluConfidenceThreshold;
        /// <summary>
        /// List of SlotTypes
        /// </summary>
        public readonly ImmutableArray<Outputs.BotSlotType> SlotTypes;
        public readonly Outputs.BotVoiceSettings? VoiceSettings;

        [OutputConstructor]
        private BotLocale(
            Outputs.BotCustomVocabulary? customVocabulary,

            string? description,

            ImmutableArray<Outputs.BotIntent> intents,

            string localeId,

            double nluConfidenceThreshold,

            ImmutableArray<Outputs.BotSlotType> slotTypes,

            Outputs.BotVoiceSettings? voiceSettings)
        {
            CustomVocabulary = customVocabulary;
            Description = description;
            Intents = intents;
            LocaleId = localeId;
            NluConfidenceThreshold = nluConfidenceThreshold;
            SlotTypes = slotTypes;
            VoiceSettings = voiceSettings;
        }
    }
}