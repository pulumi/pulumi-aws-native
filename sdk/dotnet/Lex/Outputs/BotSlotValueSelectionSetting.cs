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
    /// Contains settings used by Amazon Lex to select a slot value.
    /// </summary>
    [OutputType]
    public sealed class BotSlotValueSelectionSetting
    {
        public readonly Outputs.BotAdvancedRecognitionSetting? AdvancedRecognitionSetting;
        public readonly Outputs.BotSlotValueRegexFilter? RegexFilter;
        public readonly Pulumi.AwsNative.Lex.BotSlotValueResolutionStrategy ResolutionStrategy;

        [OutputConstructor]
        private BotSlotValueSelectionSetting(
            Outputs.BotAdvancedRecognitionSetting? advancedRecognitionSetting,

            Outputs.BotSlotValueRegexFilter? regexFilter,

            Pulumi.AwsNative.Lex.BotSlotValueResolutionStrategy resolutionStrategy)
        {
            AdvancedRecognitionSetting = advancedRecognitionSetting;
            RegexFilter = regexFilter;
            ResolutionStrategy = resolutionStrategy;
        }
    }
}