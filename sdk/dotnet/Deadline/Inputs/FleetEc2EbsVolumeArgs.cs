// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Deadline.Inputs
{

    public sealed class FleetEc2EbsVolumeArgs : global::Pulumi.ResourceArgs
    {
        [Input("iops")]
        public Input<int>? Iops { get; set; }

        [Input("sizeGiB")]
        public Input<int>? SizeGiB { get; set; }

        [Input("throughputMiB")]
        public Input<int>? ThroughputMiB { get; set; }

        public FleetEc2EbsVolumeArgs()
        {
        }
        public static new FleetEc2EbsVolumeArgs Empty => new FleetEc2EbsVolumeArgs();
    }
}
