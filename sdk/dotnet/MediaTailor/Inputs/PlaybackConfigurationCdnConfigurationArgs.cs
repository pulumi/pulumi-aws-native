// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaTailor.Inputs
{

    /// <summary>
    /// The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
    /// </summary>
    public sealed class PlaybackConfigurationCdnConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A non-default content delivery network (CDN) to serve ad segments. By default, AWS Elemental MediaTailor uses Amazon CloudFront with default cache settings as its CDN for ad segments. To set up an alternate CDN, create a rule in your CDN for the origin ads.mediatailor.&amp;lt;region&gt;.amazonaws.com. Then specify the rule's name in this AdSegmentUrlPrefix. When AWS Elemental MediaTailor serves a manifest, it reports your CDN as the source for ad segments.
        /// </summary>
        [Input("adSegmentUrlPrefix")]
        public Input<string>? AdSegmentUrlPrefix { get; set; }

        /// <summary>
        /// A content delivery network (CDN) to cache content segments, so that content requests don't always have to go to the origin server. First, create a rule in your CDN for the content segment origin server. Then specify the rule's name in this ContentSegmentUrlPrefix. When AWS Elemental MediaTailor serves a manifest, it reports your CDN as the source for content segments.
        /// </summary>
        [Input("contentSegmentUrlPrefix")]
        public Input<string>? ContentSegmentUrlPrefix { get; set; }

        public PlaybackConfigurationCdnConfigurationArgs()
        {
        }
        public static new PlaybackConfigurationCdnConfigurationArgs Empty => new PlaybackConfigurationCdnConfigurationArgs();
    }
}