// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Inputs
{

    public sealed class BotSampleUtteranceArgs : global::Pulumi.ResourceArgs
    {
        [Input("utterance", required: true)]
        public Input<string> Utterance { get; set; } = null!;

        public BotSampleUtteranceArgs()
        {
        }
        public static new BotSampleUtteranceArgs Empty => new BotSampleUtteranceArgs();
    }
}
