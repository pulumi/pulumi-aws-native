// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTWireless.Inputs
{

    public sealed class WirelessDeviceLoRaWanDeviceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ABP device object for LoRaWAN specification v1.0.x.
        /// </summary>
        [Input("abpV10x")]
        public Input<Inputs.WirelessDeviceAbpV10xArgs>? AbpV10x { get; set; }

        /// <summary>
        /// ABP device object for create APIs for v1.1.
        /// </summary>
        [Input("abpV11")]
        public Input<Inputs.WirelessDeviceAbpV11Args>? AbpV11 { get; set; }

        /// <summary>
        /// The DevEUI value.
        /// </summary>
        [Input("devEui")]
        public Input<string>? DevEui { get; set; }

        /// <summary>
        /// The ID of the device profile for the new wireless device.
        /// </summary>
        [Input("deviceProfileId")]
        public Input<string>? DeviceProfileId { get; set; }

        /// <summary>
        /// List of FPort assigned for different LoRaWAN application packages to use.
        /// </summary>
        [Input("fPorts")]
        public Input<Inputs.WirelessDeviceFPortsArgs>? FPorts { get; set; }

        /// <summary>
        /// OTAA device object for create APIs for v1.0.x
        /// </summary>
        [Input("otaaV10x")]
        public Input<Inputs.WirelessDeviceOtaaV10xArgs>? OtaaV10x { get; set; }

        /// <summary>
        /// OTAA device object for v1.1 for create APIs.
        /// </summary>
        [Input("otaaV11")]
        public Input<Inputs.WirelessDeviceOtaaV11Args>? OtaaV11 { get; set; }

        /// <summary>
        /// The ID of the service profile.
        /// </summary>
        [Input("serviceProfileId")]
        public Input<string>? ServiceProfileId { get; set; }

        public WirelessDeviceLoRaWanDeviceArgs()
        {
        }
        public static new WirelessDeviceLoRaWanDeviceArgs Empty => new WirelessDeviceLoRaWanDeviceArgs();
    }
}
