// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GroundStation.Inputs
{

    public sealed class ConfigAntennaDownlinkDemodDecodeConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines how the RF signal will be decoded.
        /// </summary>
        [Input("decodeConfig")]
        public Input<Inputs.ConfigDecodeConfigArgs>? DecodeConfig { get; set; }

        /// <summary>
        /// Defines how the RF signal will be demodulated.
        /// </summary>
        [Input("demodulationConfig")]
        public Input<Inputs.ConfigDemodulationConfigArgs>? DemodulationConfig { get; set; }

        /// <summary>
        /// Defines the spectrum configuration.
        /// </summary>
        [Input("spectrumConfig")]
        public Input<Inputs.ConfigSpectrumConfigArgs>? SpectrumConfig { get; set; }

        public ConfigAntennaDownlinkDemodDecodeConfigArgs()
        {
        }
        public static new ConfigAntennaDownlinkDemodDecodeConfigArgs Empty => new ConfigAntennaDownlinkDemodDecodeConfigArgs();
    }
}
