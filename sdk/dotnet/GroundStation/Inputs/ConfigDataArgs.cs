// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GroundStation.Inputs
{

    public sealed class ConfigDataArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Provides information for an antenna downlink config object. Antenna downlink config objects are used to provide parameters for downlinks where no demodulation or decoding is performed by Ground Station (RF over IP downlinks).
        /// </summary>
        [Input("antennaDownlinkConfig")]
        public Input<Inputs.ConfigAntennaDownlinkConfigArgs>? AntennaDownlinkConfig { get; set; }

        /// <summary>
        /// Provides information for a downlink demod decode config object. Downlink demod decode config objects are used to provide parameters for downlinks where the Ground Station service will demodulate and decode the downlinked data.
        /// </summary>
        [Input("antennaDownlinkDemodDecodeConfig")]
        public Input<Inputs.ConfigAntennaDownlinkDemodDecodeConfigArgs>? AntennaDownlinkDemodDecodeConfig { get; set; }

        /// <summary>
        /// Provides information for an uplink config object. Uplink config objects are used to provide parameters for uplink contacts.
        /// </summary>
        [Input("antennaUplinkConfig")]
        public Input<Inputs.ConfigAntennaUplinkConfigArgs>? AntennaUplinkConfig { get; set; }

        /// <summary>
        /// Provides information for a dataflow endpoint config object. Dataflow endpoint config objects are used to provide parameters about which IP endpoint(s) to use during a contact. Dataflow endpoints are where Ground Station sends data during a downlink contact and where Ground Station receives data to send to the satellite during an uplink contact.
        /// </summary>
        [Input("dataflowEndpointConfig")]
        public Input<Inputs.ConfigDataflowEndpointConfigArgs>? DataflowEndpointConfig { get; set; }

        /// <summary>
        /// Provides information for an S3 recording config object. S3 recording config objects are used to provide parameters for S3 recording during downlink contacts.
        /// </summary>
        [Input("s3RecordingConfig")]
        public Input<Inputs.ConfigS3RecordingConfigArgs>? S3RecordingConfig { get; set; }

        /// <summary>
        /// Provides information for a tracking config object. Tracking config objects are used to provide parameters about how to track the satellite through the sky during a contact.
        /// </summary>
        [Input("trackingConfig")]
        public Input<Inputs.ConfigTrackingConfigArgs>? TrackingConfig { get; set; }

        /// <summary>
        /// Provides information for an uplink echo config object. Uplink echo config objects are used to provide parameters for uplink echo during uplink contacts.
        /// </summary>
        [Input("uplinkEchoConfig")]
        public Input<Inputs.ConfigUplinkEchoConfigArgs>? UplinkEchoConfig { get; set; }

        public ConfigDataArgs()
        {
        }
        public static new ConfigDataArgs Empty => new ConfigDataArgs();
    }
}
