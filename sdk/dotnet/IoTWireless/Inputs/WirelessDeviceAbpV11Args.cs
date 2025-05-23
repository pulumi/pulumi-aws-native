// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTWireless.Inputs
{

    public sealed class WirelessDeviceAbpV11Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The DevAddr value.
        /// </summary>
        [Input("devAddr", required: true)]
        public Input<string> DevAddr { get; set; } = null!;

        /// <summary>
        /// Session keys for ABP v1.1.
        /// </summary>
        [Input("sessionKeys", required: true)]
        public Input<Inputs.WirelessDeviceSessionKeysAbpV11Args> SessionKeys { get; set; } = null!;

        public WirelessDeviceAbpV11Args()
        {
        }
        public static new WirelessDeviceAbpV11Args Empty => new WirelessDeviceAbpV11Args();
    }
}
