// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    /// <summary>
    /// A key-value pair to associate with a resource.
    /// </summary>
    public sealed class MultiplexSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum video buffer delay in milliseconds.
        /// </summary>
        [Input("maximumVideoBufferDelayMilliseconds")]
        public Input<int>? MaximumVideoBufferDelayMilliseconds { get; set; }

        /// <summary>
        /// Transport stream bit rate.
        /// </summary>
        [Input("transportStreamBitrate", required: true)]
        public Input<int> TransportStreamBitrate { get; set; } = null!;

        /// <summary>
        /// Transport stream ID.
        /// </summary>
        [Input("transportStreamId", required: true)]
        public Input<int> TransportStreamId { get; set; } = null!;

        /// <summary>
        /// Transport stream reserved bit rate.
        /// </summary>
        [Input("transportStreamReservedBitrate")]
        public Input<int>? TransportStreamReservedBitrate { get; set; }

        public MultiplexSettingsArgs()
        {
        }
        public static new MultiplexSettingsArgs Empty => new MultiplexSettingsArgs();
    }
}
