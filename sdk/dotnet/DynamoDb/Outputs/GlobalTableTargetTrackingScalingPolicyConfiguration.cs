// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DynamoDb.Outputs
{

    [OutputType]
    public sealed class GlobalTableTargetTrackingScalingPolicyConfiguration
    {
        public readonly bool? DisableScaleIn;
        public readonly int? ScaleInCooldown;
        public readonly int? ScaleOutCooldown;
        public readonly double TargetValue;

        [OutputConstructor]
        private GlobalTableTargetTrackingScalingPolicyConfiguration(
            bool? disableScaleIn,

            int? scaleInCooldown,

            int? scaleOutCooldown,

            double targetValue)
        {
            DisableScaleIn = disableScaleIn;
            ScaleInCooldown = scaleInCooldown;
            ScaleOutCooldown = scaleOutCooldown;
            TargetValue = targetValue;
        }
    }
}