// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaTailor.Outputs
{

    /// <summary>
    /// &lt;p&gt;HLS playlist configuration parameters.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class ChannelHlsPlaylistSettings
    {
        /// <summary>
        /// &lt;p&gt;Determines the type of SCTE 35 tags to use in ad markup. Specify &lt;code&gt;DATERANGE&lt;/code&gt; to use &lt;code&gt;DATERANGE&lt;/code&gt; tags (for live or VOD content). Specify &lt;code&gt;SCTE35_ENHANCED&lt;/code&gt; to use &lt;code&gt;EXT-X-CUE-OUT&lt;/code&gt; and &lt;code&gt;EXT-X-CUE-IN&lt;/code&gt; tags (for VOD content only).&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.MediaTailor.ChannelAdMarkupType> AdMarkupType;
        /// <summary>
        /// &lt;p&gt;The total duration (in seconds) of each manifest. Minimum value: &lt;code&gt;30&lt;/code&gt; seconds. Maximum value: &lt;code&gt;3600&lt;/code&gt; seconds.&lt;/p&gt;
        /// </summary>
        public readonly double? ManifestWindowSeconds;

        [OutputConstructor]
        private ChannelHlsPlaylistSettings(
            ImmutableArray<Pulumi.AwsNative.MediaTailor.ChannelAdMarkupType> adMarkupType,

            double? manifestWindowSeconds)
        {
            AdMarkupType = adMarkupType;
            ManifestWindowSeconds = manifestWindowSeconds;
        }
    }
}
