// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Outputs
{

    [OutputType]
    public sealed class CapacityReservationFleetInstanceTypeSpecification
    {
        /// <summary>
        /// The Availability Zone in which the Capacity Reservation Fleet reserves the capacity. A Capacity Reservation Fleet can't span Availability Zones. All instance type specifications that you specify for the Fleet must use the same Availability Zone.
        /// </summary>
        public readonly string? AvailabilityZone;
        /// <summary>
        /// The ID of the Availability Zone in which the Capacity Reservation Fleet reserves the capacity. A Capacity Reservation Fleet can't span Availability Zones. All instance type specifications that you specify for the Fleet must use the same Availability Zone.
        /// </summary>
        public readonly string? AvailabilityZoneId;
        /// <summary>
        /// Indicates whether the Capacity Reservation Fleet supports EBS-optimized instances types. This optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal I/O performance. This optimization isn't available with all instance types. Additional usage charges apply when using EBS-optimized instance types.
        /// </summary>
        public readonly bool? EbsOptimized;
        /// <summary>
        /// The type of operating system for which the Capacity Reservation Fleet reserves capacity.
        /// </summary>
        public readonly string? InstancePlatform;
        /// <summary>
        /// The instance type for which the Capacity Reservation Fleet reserves capacity.
        /// </summary>
        public readonly string? InstanceType;
        /// <summary>
        /// The priority to assign to the instance type. This value is used to determine which of the instance types specified for the Fleet should be prioritized for use. A lower value indicates a high priority. For more information, see [Instance type priority](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/crfleet-concepts.html#instance-priority) in the *Amazon EC2 User Guide* .
        /// </summary>
        public readonly int? Priority;
        /// <summary>
        /// The number of capacity units provided by the specified instance type. This value, together with the total target capacity that you specify for the Fleet determine the number of instances for which the Fleet reserves capacity. Both values are based on units that make sense for your workload. For more information, see [Total target capacity](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/crfleet-concepts.html#target-capacity) in the Amazon EC2 User Guide.
        /// 
        /// Valid Range: Minimum value of `0.001` . Maximum value of `99.999` .
        /// </summary>
        public readonly double? Weight;

        [OutputConstructor]
        private CapacityReservationFleetInstanceTypeSpecification(
            string? availabilityZone,

            string? availabilityZoneId,

            bool? ebsOptimized,

            string? instancePlatform,

            string? instanceType,

            int? priority,

            double? weight)
        {
            AvailabilityZone = availabilityZone;
            AvailabilityZoneId = availabilityZoneId;
            EbsOptimized = ebsOptimized;
            InstancePlatform = instancePlatform;
            InstanceType = instanceType;
            Priority = priority;
            Weight = weight;
        }
    }
}
