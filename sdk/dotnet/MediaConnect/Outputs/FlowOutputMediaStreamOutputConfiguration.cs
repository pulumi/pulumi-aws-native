// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaConnect.Outputs
{

    /// <summary>
    /// The media stream that is associated with the output, and the parameters for that association.
    /// </summary>
    [OutputType]
    public sealed class FlowOutputMediaStreamOutputConfiguration
    {
        /// <summary>
        /// The media streams that you want to associate with the output.
        /// </summary>
        public readonly ImmutableArray<Outputs.FlowOutputDestinationConfiguration> DestinationConfigurations;
        /// <summary>
        /// The format that will be used to encode the data. For ancillary data streams, set the encoding name to smpte291. For audio streams, set the encoding name to pcm. For video streams on sources or outputs that use the CDI protocol, set the encoding name to raw. For video streams on sources or outputs that use the ST 2110 JPEG XS protocol, set the encoding name to jxsv.
        /// </summary>
        public readonly Pulumi.AwsNative.MediaConnect.FlowOutputMediaStreamOutputConfigurationEncodingName EncodingName;
        /// <summary>
        /// A collection of parameters that determine how MediaConnect will convert the content. These fields only apply to outputs on flows that have a CDI source.
        /// </summary>
        public readonly Outputs.FlowOutputEncodingParameters? EncodingParameters;
        /// <summary>
        /// A name that helps you distinguish one media stream from another.
        /// </summary>
        public readonly string MediaStreamName;

        [OutputConstructor]
        private FlowOutputMediaStreamOutputConfiguration(
            ImmutableArray<Outputs.FlowOutputDestinationConfiguration> destinationConfigurations,

            Pulumi.AwsNative.MediaConnect.FlowOutputMediaStreamOutputConfigurationEncodingName encodingName,

            Outputs.FlowOutputEncodingParameters? encodingParameters,

            string mediaStreamName)
        {
            DestinationConfigurations = destinationConfigurations;
            EncodingName = encodingName;
            EncodingParameters = encodingParameters;
            MediaStreamName = mediaStreamName;
        }
    }
}
