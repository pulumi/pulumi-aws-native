// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTWireless.Inputs
{

    public sealed class WirelessDeviceOtaaV11Args : global::Pulumi.ResourceArgs
    {
        [Input("appKey", required: true)]
        public Input<string> AppKey { get; set; } = null!;

        [Input("joinEui", required: true)]
        public Input<string> JoinEui { get; set; } = null!;

        [Input("nwkKey", required: true)]
        public Input<string> NwkKey { get; set; } = null!;

        public WirelessDeviceOtaaV11Args()
        {
        }
        public static new WirelessDeviceOtaaV11Args Empty => new WirelessDeviceOtaaV11Args();
    }
}