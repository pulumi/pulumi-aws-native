// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    public sealed class Ec2FleetAcceleratorTotalMemoryMiBRequestArgs : global::Pulumi.ResourceArgs
    {
        [Input("max")]
        public Input<int>? Max { get; set; }

        [Input("min")]
        public Input<int>? Min { get; set; }

        public Ec2FleetAcceleratorTotalMemoryMiBRequestArgs()
        {
        }
        public static new Ec2FleetAcceleratorTotalMemoryMiBRequestArgs Empty => new Ec2FleetAcceleratorTotalMemoryMiBRequestArgs();
    }
}