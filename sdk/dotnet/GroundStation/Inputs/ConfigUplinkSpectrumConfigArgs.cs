// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GroundStation.Inputs
{

    public sealed class ConfigUplinkSpectrumConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("centerFrequency")]
        public Input<Inputs.ConfigFrequencyArgs>? CenterFrequency { get; set; }

        [Input("polarization")]
        public Input<Pulumi.AwsNative.GroundStation.ConfigPolarization>? Polarization { get; set; }

        public ConfigUplinkSpectrumConfigArgs()
        {
        }
        public static new ConfigUplinkSpectrumConfigArgs Empty => new ConfigUplinkSpectrumConfigArgs();
    }
}