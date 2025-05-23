// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Inputs
{

    public sealed class BotExternalSourceSettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Settings required for a slot type based on a grammar that you provide.
        /// </summary>
        [Input("grammarSlotTypeSetting")]
        public Input<Inputs.BotGrammarSlotTypeSettingArgs>? GrammarSlotTypeSetting { get; set; }

        public BotExternalSourceSettingArgs()
        {
        }
        public static new BotExternalSourceSettingArgs Empty => new BotExternalSourceSettingArgs();
    }
}
