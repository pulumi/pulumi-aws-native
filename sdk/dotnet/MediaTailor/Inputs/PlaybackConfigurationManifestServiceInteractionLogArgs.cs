// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaTailor.Inputs
{

    public sealed class PlaybackConfigurationManifestServiceInteractionLogArgs : global::Pulumi.ResourceArgs
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

        public PlaybackConfigurationManifestServiceInteractionLogArgs()
        {
        }
        public static new PlaybackConfigurationManifestServiceInteractionLogArgs Empty => new PlaybackConfigurationManifestServiceInteractionLogArgs();
    }
}
