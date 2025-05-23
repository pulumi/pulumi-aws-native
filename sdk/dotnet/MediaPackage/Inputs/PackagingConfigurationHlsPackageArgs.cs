// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackage.Inputs
{

    /// <summary>
    /// An HTTP Live Streaming (HLS) packaging configuration.
    /// </summary>
    public sealed class PackagingConfigurationHlsPackageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Parameters for encrypting content.
        /// </summary>
        [Input("encryption")]
        public Input<Inputs.PackagingConfigurationHlsEncryptionArgs>? Encryption { get; set; }

        [Input("hlsManifests", required: true)]
        private InputList<Inputs.PackagingConfigurationHlsManifestArgs>? _hlsManifests;

        /// <summary>
        /// A list of HLS manifest configurations.
        /// </summary>
        public InputList<Inputs.PackagingConfigurationHlsManifestArgs> HlsManifests
        {
            get => _hlsManifests ?? (_hlsManifests = new InputList<Inputs.PackagingConfigurationHlsManifestArgs>());
            set => _hlsManifests = value;
        }

        /// <summary>
        /// When enabled, MediaPackage passes through digital video broadcasting (DVB) subtitles into the output.
        /// </summary>
        [Input("includeDvbSubtitles")]
        public Input<bool>? IncludeDvbSubtitles { get; set; }

        /// <summary>
        /// Duration (in seconds) of each fragment. Actual fragments are rounded to the nearest multiple of the source fragment duration.
        /// </summary>
        [Input("segmentDurationSeconds")]
        public Input<int>? SegmentDurationSeconds { get; set; }

        /// <summary>
        /// When enabled, audio streams will be placed in rendition groups in the output.
        /// </summary>
        [Input("useAudioRenditionGroup")]
        public Input<bool>? UseAudioRenditionGroup { get; set; }

        public PackagingConfigurationHlsPackageArgs()
        {
        }
        public static new PackagingConfigurationHlsPackageArgs Empty => new PackagingConfigurationHlsPackageArgs();
    }
}
