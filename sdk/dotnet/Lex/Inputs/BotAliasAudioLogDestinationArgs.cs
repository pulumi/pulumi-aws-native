// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Inputs
{

    /// <summary>
    /// The location of audio log files collected when conversation logging is enabled for a bot.
    /// </summary>
    public sealed class BotAliasAudioLogDestinationArgs : global::Pulumi.ResourceArgs
    {
        [Input("s3Bucket", required: true)]
        public Input<Inputs.BotAliasS3BucketLogDestinationArgs> S3Bucket { get; set; } = null!;

        public BotAliasAudioLogDestinationArgs()
        {
        }
        public static new BotAliasAudioLogDestinationArgs Empty => new BotAliasAudioLogDestinationArgs();
    }
}
