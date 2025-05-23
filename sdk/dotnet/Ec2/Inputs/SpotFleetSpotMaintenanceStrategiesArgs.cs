// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    public sealed class SpotFleetSpotMaintenanceStrategiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Spot Instance replacement strategy to use when Amazon EC2 emits a signal that your Spot Instance is at an elevated risk of being interrupted. For more information, see [Capacity rebalancing](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet-capacity-rebalance.html) in the *Amazon EC2 User Guide* .
        /// </summary>
        [Input("capacityRebalance")]
        public Input<Inputs.SpotFleetSpotCapacityRebalanceArgs>? CapacityRebalance { get; set; }

        public SpotFleetSpotMaintenanceStrategiesArgs()
        {
        }
        public static new SpotFleetSpotMaintenanceStrategiesArgs Empty => new SpotFleetSpotMaintenanceStrategiesArgs();
    }
}
