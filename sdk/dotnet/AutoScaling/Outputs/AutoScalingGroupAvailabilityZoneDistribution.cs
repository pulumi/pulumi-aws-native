// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScaling.Outputs
{

    /// <summary>
    /// ``AvailabilityZoneDistribution`` is a property of the [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html) resource.
    /// </summary>
    [OutputType]
    public sealed class AutoScalingGroupAvailabilityZoneDistribution
    {
        /// <summary>
        /// If launches fail in an Availability Zone, the following strategies are available. The default is ``balanced-best-effort``. 
        ///   +  ``balanced-only`` - If launches fail in an Availability Zone, Auto Scaling will continue to attempt to launch in the unhealthy zone to preserve a balanced distribution.
        ///   +  ``balanced-best-effort`` - If launches fail in an Availability Zone, Auto Scaling will attempt to launch in another healthy Availability Zone instead.
        /// </summary>
        public readonly Pulumi.AwsNative.AutoScaling.AutoScalingGroupAvailabilityZoneDistributionCapacityDistributionStrategy? CapacityDistributionStrategy;

        [OutputConstructor]
        private AutoScalingGroupAvailabilityZoneDistribution(Pulumi.AwsNative.AutoScaling.AutoScalingGroupAvailabilityZoneDistributionCapacityDistributionStrategy? capacityDistributionStrategy)
        {
            CapacityDistributionStrategy = capacityDistributionStrategy;
        }
    }
}
