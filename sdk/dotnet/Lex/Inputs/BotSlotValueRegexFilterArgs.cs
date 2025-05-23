// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Inputs
{

    public sealed class BotSlotValueRegexFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A regular expression used to validate the value of a slot.
        /// 
        /// Use a standard regular expression. Amazon Lex supports the following characters in the regular expression:
        /// 
        /// - A-Z, a-z
        /// - 0-9
        /// - Unicode characters ("\⁠u&lt;Unicode&gt;")
        /// 
        /// Represent Unicode characters with four digits, for example "\⁠u0041" or "\⁠u005A".
        /// 
        /// The following regular expression operators are not supported:
        /// 
        /// - Infinite repeaters: *, +, or {x,} with no upper bound.
        /// - Wild card (.)
        /// </summary>
        [Input("pattern", required: true)]
        public Input<string> Pattern { get; set; } = null!;

        public BotSlotValueRegexFilterArgs()
        {
        }
        public static new BotSlotValueRegexFilterArgs Empty => new BotSlotValueRegexFilterArgs();
    }
}
