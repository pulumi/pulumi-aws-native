// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Deadline.Inputs
{

    public sealed class FleetAcceleratorTotalMemoryMiBRangeArgs : global::Pulumi.ResourceArgs
    {
        [Input("max")]
        public Input<int>? Max { get; set; }

        [Input("min", required: true)]
        public Input<int> Min { get; set; } = null!;

        public FleetAcceleratorTotalMemoryMiBRangeArgs()
        {
        }
        public static new FleetAcceleratorTotalMemoryMiBRangeArgs Empty => new FleetAcceleratorTotalMemoryMiBRangeArgs();
    }
}
