// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaTailor.Inputs
{

    /// <summary>
    /// The configuration for pre-roll ad insertion.
    /// </summary>
    public sealed class PlaybackConfigurationLivePreRollConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL for the ad decision server (ADS) for pre-roll ads. This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing, you can provide a static VAST URL. The maximum length is 25,000 characters.
        /// </summary>
        [Input("adDecisionServerUrl")]
        public Input<string>? AdDecisionServerUrl { get; set; }

        /// <summary>
        /// The maximum allowed duration for the pre-roll ad avail. AWS Elemental MediaTailor won't play pre-roll ads to exceed this duration, regardless of the total duration of ads that the ADS returns.
        /// </summary>
        [Input("maxDurationSeconds")]
        public Input<int>? MaxDurationSeconds { get; set; }

        public PlaybackConfigurationLivePreRollConfigurationArgs()
        {
        }
        public static new PlaybackConfigurationLivePreRollConfigurationArgs Empty => new PlaybackConfigurationLivePreRollConfigurationArgs();
    }
}
