// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    public sealed class SpotFleetLaunchTemplateOverridesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Availability Zone in which to launch the instances.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// The instance requirements. When you specify instance requirements, Amazon EC2 will identify instance types with the provided requirements, and then use your On-Demand and Spot allocation strategies to launch instances from these instance types, in the same way as when you specify a list of instance types.
        /// 
        /// &gt; If you specify `InstanceRequirements` , you can't specify `InstanceType` .
        /// </summary>
        [Input("instanceRequirements")]
        public Input<Inputs.SpotFleetInstanceRequirementsRequestArgs>? InstanceRequirements { get; set; }

        /// <summary>
        /// The instance type.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The priority for the launch template override. The highest priority is launched first.
        /// 
        /// If `OnDemandAllocationStrategy` is set to `prioritized` , Spot Fleet uses priority to determine which launch template override to use first in fulfilling On-Demand capacity.
        /// 
        /// If the Spot `AllocationStrategy` is set to `capacityOptimizedPrioritized` , Spot Fleet uses priority on a best-effort basis to determine which launch template override to use in fulfilling Spot capacity, but optimizes for capacity first.
        /// 
        /// Valid values are whole numbers starting at `0` . The lower the number, the higher the priority. If no number is set, the launch template override has the lowest priority. You can set the same priority for different launch template overrides.
        /// </summary>
        [Input("priority")]
        public Input<double>? Priority { get; set; }

        /// <summary>
        /// The maximum price per unit hour that you are willing to pay for a Spot Instance. We do not recommend using this parameter because it can lead to increased interruptions. If you do not specify this parameter, you will pay the current Spot price.
        /// 
        /// &gt; If you specify a maximum price, your instances will be interrupted more frequently than if you do not specify this parameter.
        /// </summary>
        [Input("spotPrice")]
        public Input<string>? SpotPrice { get; set; }

        /// <summary>
        /// The ID of the subnet in which to launch the instances.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// The number of units provided by the specified instance type. These are the same units that you chose to set the target capacity in terms of instances, or a performance characteristic such as vCPUs, memory, or I/O.
        /// 
        /// If the target capacity divided by this value is not a whole number, Amazon EC2 rounds the number of instances to the next whole number. If this value is not specified, the default is 1.
        /// 
        /// &gt; When specifying weights, the price used in the `lowestPrice` and `priceCapacityOptimized` allocation strategies is per *unit* hour (where the instance price is divided by the specified weight). However, if all the specified weights are above the requested `TargetCapacity` , resulting in only 1 instance being launched, the price used is per *instance* hour.
        /// </summary>
        [Input("weightedCapacity")]
        public Input<double>? WeightedCapacity { get; set; }

        public SpotFleetLaunchTemplateOverridesArgs()
        {
        }
        public static new SpotFleetLaunchTemplateOverridesArgs Empty => new SpotFleetLaunchTemplateOverridesArgs();
    }
}
