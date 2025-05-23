// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackage.Outputs
{

    /// <summary>
    /// A Common Media Application Format (CMAF) packaging configuration.
    /// </summary>
    [OutputType]
    public sealed class OriginEndpointCmafPackage
    {
        /// <summary>
        /// Parameters for encrypting content.
        /// </summary>
        public readonly Outputs.OriginEndpointCmafEncryption? Encryption;
        /// <summary>
        /// A list of HLS manifest configurations
        /// </summary>
        public readonly ImmutableArray<Outputs.OriginEndpointHlsManifest> HlsManifests;
        /// <summary>
        /// Duration (in seconds) of each segment. Actual segments will be rounded to the nearest multiple of the source segment duration.
        /// </summary>
        public readonly int? SegmentDurationSeconds;
        /// <summary>
        /// An optional custom string that is prepended to the name of each segment. If not specified, it defaults to the ChannelId.
        /// </summary>
        public readonly string? SegmentPrefix;
        /// <summary>
        /// Limitations for outputs from the endpoint, based on the video bitrate.
        /// </summary>
        public readonly Outputs.OriginEndpointStreamSelection? StreamSelection;

        [OutputConstructor]
        private OriginEndpointCmafPackage(
            Outputs.OriginEndpointCmafEncryption? encryption,

            ImmutableArray<Outputs.OriginEndpointHlsManifest> hlsManifests,

            int? segmentDurationSeconds,

            string? segmentPrefix,

            Outputs.OriginEndpointStreamSelection? streamSelection)
        {
            Encryption = encryption;
            HlsManifests = hlsManifests;
            SegmentDurationSeconds = segmentDurationSeconds;
            SegmentPrefix = segmentPrefix;
            StreamSelection = streamSelection;
        }
    }
}
