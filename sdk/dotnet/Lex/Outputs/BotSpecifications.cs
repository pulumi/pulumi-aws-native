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
    public sealed class BotSpecifications
    {
        public readonly string SlotTypeId;
        public readonly Outputs.BotSubSlotValueElicitationSetting ValueElicitationSetting;

        [OutputConstructor]
        private BotSpecifications(
            string slotTypeId,

            Outputs.BotSubSlotValueElicitationSetting valueElicitationSetting)
        {
            SlotTypeId = slotTypeId;
            ValueElicitationSetting = valueElicitationSetting;
        }
    }
}
