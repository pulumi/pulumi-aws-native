// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    public sealed class EC2FleetOnDemandOptionsRequestArgs : global::Pulumi.ResourceArgs
    {
        [Input("allocationStrategy")]
        public Input<string>? AllocationStrategy { get; set; }

        [Input("capacityReservationOptions")]
        public Input<Inputs.EC2FleetCapacityReservationOptionsRequestArgs>? CapacityReservationOptions { get; set; }

        [Input("maxTotalPrice")]
        public Input<string>? MaxTotalPrice { get; set; }

        [Input("minTargetCapacity")]
        public Input<int>? MinTargetCapacity { get; set; }

        [Input("singleAvailabilityZone")]
        public Input<bool>? SingleAvailabilityZone { get; set; }

        [Input("singleInstanceType")]
        public Input<bool>? SingleInstanceType { get; set; }

        public EC2FleetOnDemandOptionsRequestArgs()
        {
        }
        public static new EC2FleetOnDemandOptionsRequestArgs Empty => new EC2FleetOnDemandOptionsRequestArgs();
    }
}