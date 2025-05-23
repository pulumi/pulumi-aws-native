// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Inputs
{

    public sealed class BotGrammarSlotTypeSettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The source of the grammar used to create the slot type.
        /// </summary>
        [Input("source")]
        public Input<Inputs.BotGrammarSlotTypeSourceArgs>? Source { get; set; }

        public BotGrammarSlotTypeSettingArgs()
        {
        }
        public static new BotGrammarSlotTypeSettingArgs Empty => new BotGrammarSlotTypeSettingArgs();
    }
}
