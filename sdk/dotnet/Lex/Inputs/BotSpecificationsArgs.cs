// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Inputs
{

    public sealed class BotSpecificationsArgs : global::Pulumi.ResourceArgs
    {
        [Input("slotTypeId", required: true)]
        public Input<string> SlotTypeId { get; set; } = null!;

        [Input("valueElicitationSetting", required: true)]
        public Input<Inputs.BotSubSlotValueElicitationSettingArgs> ValueElicitationSetting { get; set; } = null!;

        public BotSpecificationsArgs()
        {
        }
        public static new BotSpecificationsArgs Empty => new BotSpecificationsArgs();
    }
}
