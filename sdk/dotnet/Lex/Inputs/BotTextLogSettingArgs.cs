// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Inputs
{

    public sealed class BotTextLogSettingArgs : global::Pulumi.ResourceArgs
    {
        [Input("destination", required: true)]
        public Input<Inputs.BotTextLogDestinationArgs> Destination { get; set; } = null!;

        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        public BotTextLogSettingArgs()
        {
        }
        public static new BotTextLogSettingArgs Empty => new BotTextLogSettingArgs();
    }
}
