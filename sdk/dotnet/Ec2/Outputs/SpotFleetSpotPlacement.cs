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
    public sealed class SpotFleetSpotPlacement
    {
        /// <summary>
        /// The Availability Zone.
        /// 
        /// To specify multiple Availability Zones, separate them using commas; for example, "us-west-2a, us-west-2b".
        /// </summary>
        public readonly string? AvailabilityZone;
        /// <summary>
        /// The name of the placement group.
        /// </summary>
        public readonly string? GroupName;
        /// <summary>
        /// The tenancy of the instance (if the instance is running in a VPC). An instance with a tenancy of `dedicated` runs on single-tenant hardware. The `host` tenancy is not supported for Spot Instances.
        /// </summary>
        public readonly Pulumi.AwsNative.Ec2.SpotFleetSpotPlacementTenancy? Tenancy;

        [OutputConstructor]
        private SpotFleetSpotPlacement(
            string? availabilityZone,

            string? groupName,

            Pulumi.AwsNative.Ec2.SpotFleetSpotPlacementTenancy? tenancy)
        {
            AvailabilityZone = availabilityZone;
            GroupName = groupName;
            Tenancy = tenancy;
        }
    }
}
