// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NimbleStudio.Inputs
{

    public sealed class LaunchProfileVolumeConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("iops")]
        public Input<double>? Iops { get; set; }

        [Input("size")]
        public Input<double>? Size { get; set; }

        [Input("throughput")]
        public Input<double>? Throughput { get; set; }

        public LaunchProfileVolumeConfigurationArgs()
        {
        }
        public static new LaunchProfileVolumeConfigurationArgs Empty => new LaunchProfileVolumeConfigurationArgs();
    }
}
