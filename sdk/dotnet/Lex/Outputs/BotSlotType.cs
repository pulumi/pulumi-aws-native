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
    /// A custom, extended built-in or a grammar slot type.
    /// </summary>
    [OutputType]
    public sealed class BotSlotType
    {
        public readonly string? Description;
        public readonly Outputs.BotExternalSourceSetting? ExternalSourceSetting;
        public readonly string Name;
        public readonly string? ParentSlotTypeSignature;
        public readonly ImmutableArray<Outputs.BotSlotTypeValue> SlotTypeValues;
        public readonly Outputs.BotSlotValueSelectionSetting? ValueSelectionSetting;

        [OutputConstructor]
        private BotSlotType(
            string? description,

            Outputs.BotExternalSourceSetting? externalSourceSetting,

            string name,

            string? parentSlotTypeSignature,

            ImmutableArray<Outputs.BotSlotTypeValue> slotTypeValues,

            Outputs.BotSlotValueSelectionSetting? valueSelectionSetting)
        {
            Description = description;
            ExternalSourceSetting = externalSourceSetting;
            Name = name;
            ParentSlotTypeSignature = parentSlotTypeSignature;
            SlotTypeValues = slotTypeValues;
            ValueSelectionSetting = valueSelectionSetting;
        }
    }
}