// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScaling.Outputs
{

    [OutputType]
    public sealed class AutoScalingGroupInstanceMaintenancePolicy
    {
        public readonly int? MaxHealthyPercentage;
        public readonly int? MinHealthyPercentage;

        [OutputConstructor]
        private AutoScalingGroupInstanceMaintenancePolicy(
            int? maxHealthyPercentage,

            int? minHealthyPercentage)
        {
            MaxHealthyPercentage = maxHealthyPercentage;
            MinHealthyPercentage = minHealthyPercentage;
        }
    }
}