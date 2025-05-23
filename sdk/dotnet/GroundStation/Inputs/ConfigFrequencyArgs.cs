// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GroundStation.Inputs
{

    public sealed class ConfigFrequencyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The units of the frequency.
        /// </summary>
        [Input("units")]
        public Input<Pulumi.AwsNative.GroundStation.ConfigFrequencyUnits>? Units { get; set; }

        /// <summary>
        /// The value of the frequency. Valid values are between 2200 to 2300 MHz and 7750 to 8400 MHz for downlink and 2025 to 2120 MHz for uplink.
        /// </summary>
        [Input("value")]
        public Input<double>? Value { get; set; }

        public ConfigFrequencyArgs()
        {
        }
        public static new ConfigFrequencyArgs Empty => new ConfigFrequencyArgs();
    }
}
