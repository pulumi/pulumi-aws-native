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
    /// A slot is a variable needed to fulfill an intent, where an intent can require zero or more slots.
    /// </summary>
    [OutputType]
    public sealed class BotSlot
    {
        public readonly string? Description;
        public readonly Outputs.BotMultipleValuesSetting? MultipleValuesSetting;
        public readonly string Name;
        public readonly Outputs.BotObfuscationSetting? ObfuscationSetting;
        public readonly string SlotTypeName;
        public readonly Outputs.BotSlotValueElicitationSetting ValueElicitationSetting;

        [OutputConstructor]
        private BotSlot(
            string? description,

            Outputs.BotMultipleValuesSetting? multipleValuesSetting,

            string name,

            Outputs.BotObfuscationSetting? obfuscationSetting,

            string slotTypeName,

            Outputs.BotSlotValueElicitationSetting valueElicitationSetting)
        {
            Description = description;
            MultipleValuesSetting = multipleValuesSetting;
            Name = name;
            ObfuscationSetting = obfuscationSetting;
            SlotTypeName = slotTypeName;
            ValueElicitationSetting = valueElicitationSetting;
        }
    }
}