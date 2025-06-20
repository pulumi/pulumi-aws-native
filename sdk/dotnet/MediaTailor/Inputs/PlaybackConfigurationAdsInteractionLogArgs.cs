// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaTailor.Inputs
{

    public sealed class PlaybackConfigurationAdsInteractionLogArgs : global::Pulumi.ResourceArgs
    {
        [Input("excludeEventTypes")]
        private InputList<string>? _excludeEventTypes;

        /// <summary>
        /// Indicates that MediaTailor won't emit the selected events in the logs for playback sessions that are initialized with this configuration.
        /// </summary>
        public InputList<string> ExcludeEventTypes
        {
            get => _excludeEventTypes ?? (_excludeEventTypes = new InputList<string>());
            set => _excludeEventTypes = value;
        }

        [Input("publishOptInEventTypes")]
        private InputList<string>? _publishOptInEventTypes;

        /// <summary>
        /// Indicates that MediaTailor emits RAW_ADS_RESPONSE logs for playback sessions that are initialized with this configuration.
        /// </summary>
        public InputList<string> PublishOptInEventTypes
        {
            get => _publishOptInEventTypes ?? (_publishOptInEventTypes = new InputList<string>());
            set => _publishOptInEventTypes = value;
        }

        public PlaybackConfigurationAdsInteractionLogArgs()
        {
        }
        public static new PlaybackConfigurationAdsInteractionLogArgs Empty => new PlaybackConfigurationAdsInteractionLogArgs();
    }
}
