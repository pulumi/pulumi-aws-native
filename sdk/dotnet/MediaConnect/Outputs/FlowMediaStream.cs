// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaConnect.Outputs
{

    /// <summary>
    /// A single track or stream of media that contains video, audio, or ancillary data. After you add a media stream to a flow, you can associate it with sources and outputs on that flow, as long as they use the CDI protocol or the ST 2110 JPEG XS protocol. Each source or output can consist of one or many media streams.
    /// </summary>
    [OutputType]
    public sealed class FlowMediaStream
    {
        /// <summary>
        /// Attributes that are related to the media stream.
        /// </summary>
        public readonly Outputs.FlowMediaStreamAttributes? Attributes;
        /// <summary>
        /// The sample rate for the stream. This value in measured in kHz.
        /// </summary>
        public readonly int? ClockRate;
        /// <summary>
        /// A description that can help you quickly identify what your media stream is used for.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The format type number (sometimes referred to as RTP payload type) of the media stream. MediaConnect assigns this value to the media stream. For ST 2110 JPEG XS outputs, you need to provide this value to the receiver.
        /// </summary>
        public readonly int? Fmt;
        /// <summary>
        /// A unique identifier for the media stream.
        /// </summary>
        public readonly int MediaStreamId;
        /// <summary>
        /// A name that helps you distinguish one media stream from another.
        /// </summary>
        public readonly string MediaStreamName;
        /// <summary>
        /// The type of media stream.
        /// </summary>
        public readonly Pulumi.AwsNative.MediaConnect.FlowMediaStreamMediaStreamType MediaStreamType;
        /// <summary>
        /// The resolution of the video.
        /// </summary>
        public readonly Pulumi.AwsNative.MediaConnect.FlowMediaStreamVideoFormat? VideoFormat;

        [OutputConstructor]
        private FlowMediaStream(
            Outputs.FlowMediaStreamAttributes? attributes,

            int? clockRate,

            string? description,

            int? fmt,

            int mediaStreamId,

            string mediaStreamName,

            Pulumi.AwsNative.MediaConnect.FlowMediaStreamMediaStreamType mediaStreamType,

            Pulumi.AwsNative.MediaConnect.FlowMediaStreamVideoFormat? videoFormat)
        {
            Attributes = attributes;
            ClockRate = clockRate;
            Description = description;
            Fmt = fmt;
            MediaStreamId = mediaStreamId;
            MediaStreamName = mediaStreamName;
            MediaStreamType = mediaStreamType;
            VideoFormat = videoFormat;
        }
    }
}