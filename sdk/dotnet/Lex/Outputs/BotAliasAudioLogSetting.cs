// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Outputs
{

    /// <summary>
    /// Settings for logging audio of conversations between Amazon Lex and a user. You specify whether to log audio and the Amazon S3 bucket where the audio file is stored.
    /// </summary>
    [OutputType]
    public sealed class BotAliasAudioLogSetting
    {
        public readonly Outputs.BotAliasAudioLogDestination Destination;
        public readonly bool Enabled;

        [OutputConstructor]
        private BotAliasAudioLogSetting(
            Outputs.BotAliasAudioLogDestination destination,

            bool enabled)
        {
            Destination = destination;
            Enabled = enabled;
        }
    }
}
