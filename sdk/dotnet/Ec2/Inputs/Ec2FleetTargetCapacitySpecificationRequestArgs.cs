// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    public sealed class Ec2FleetTargetCapacitySpecificationRequestArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default target capacity type.
        /// </summary>
        [Input("defaultTargetCapacityType")]
        public Input<Pulumi.AwsNative.Ec2.Ec2FleetTargetCapacitySpecificationRequestDefaultTargetCapacityType>? DefaultTargetCapacityType { get; set; }

        /// <summary>
        /// The number of On-Demand units to request.
        /// </summary>
        [Input("onDemandTargetCapacity")]
        public Input<int>? OnDemandTargetCapacity { get; set; }

        /// <summary>
        /// The number of Spot units to request.
        /// </summary>
        [Input("spotTargetCapacity")]
        public Input<int>? SpotTargetCapacity { get; set; }

        /// <summary>
        /// The unit for the target capacity. You can specify this parameter only when using attributed-based instance type selection.
        /// 
        /// Default: `units` (the number of instances)
        /// </summary>
        [Input("targetCapacityUnitType")]
        public Input<Pulumi.AwsNative.Ec2.Ec2FleetTargetCapacitySpecificationRequestTargetCapacityUnitType>? TargetCapacityUnitType { get; set; }

        /// <summary>
        /// The number of units to request, filled using the default target capacity type.
        /// </summary>
        [Input("totalTargetCapacity", required: true)]
        public Input<int> TotalTargetCapacity { get; set; } = null!;

        public Ec2FleetTargetCapacitySpecificationRequestArgs()
        {
        }
        public static new Ec2FleetTargetCapacitySpecificationRequestArgs Empty => new Ec2FleetTargetCapacitySpecificationRequestArgs();
    }
}
