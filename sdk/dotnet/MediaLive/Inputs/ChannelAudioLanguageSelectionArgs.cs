// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class ChannelAudioLanguageSelectionArgs : global::Pulumi.ResourceArgs
    {
        [Input("languageCode")]
        public Input<string>? LanguageCode { get; set; }

        [Input("languageSelectionPolicy")]
        public Input<string>? LanguageSelectionPolicy { get; set; }

        public ChannelAudioLanguageSelectionArgs()
        {
        }
        public static new ChannelAudioLanguageSelectionArgs Empty => new ChannelAudioLanguageSelectionArgs();
    }
}