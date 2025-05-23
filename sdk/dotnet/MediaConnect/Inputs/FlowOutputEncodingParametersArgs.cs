// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaConnect.Inputs
{

    /// <summary>
    /// A collection of parameters that determine how MediaConnect will convert the content. These fields only apply to outputs on flows that have a CDI source.
    /// </summary>
    public sealed class FlowOutputEncodingParametersArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A value that is used to calculate compression for an output. The bitrate of the output is calculated as follows: Output bitrate = (1 / compressionFactor) * (source bitrate) This property only applies to outputs that use the ST 2110 JPEG XS protocol, with a flow source that uses the CDI protocol. Valid values are in the range of 3.0 to 10.0, inclusive.
        /// </summary>
        [Input("compressionFactor", required: true)]
        public Input<double> CompressionFactor { get; set; } = null!;

        /// <summary>
        /// A setting on the encoder that drives compression settings. This property only applies to video media streams associated with outputs that use the ST 2110 JPEG XS protocol, with a flow source that uses the CDI protocol.
        /// </summary>
        [Input("encoderProfile")]
        public Input<Pulumi.AwsNative.MediaConnect.FlowOutputEncodingParametersEncoderProfile>? EncoderProfile { get; set; }

        public FlowOutputEncodingParametersArgs()
        {
        }
        public static new FlowOutputEncodingParametersArgs Empty => new FlowOutputEncodingParametersArgs();
    }
}
