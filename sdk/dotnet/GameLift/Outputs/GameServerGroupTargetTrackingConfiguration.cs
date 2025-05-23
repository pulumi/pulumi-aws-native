// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift.Outputs
{

    /// <summary>
    /// Settings for a target-based scaling policy applied to Auto Scaling group.
    /// </summary>
    [OutputType]
    public sealed class GameServerGroupTargetTrackingConfiguration
    {
        /// <summary>
        /// Desired value to use with a game server group target-based scaling policy.
        /// </summary>
        public readonly double TargetValue;

        [OutputConstructor]
        private GameServerGroupTargetTrackingConfiguration(double targetValue)
        {
            TargetValue = targetValue;
        }
    }
}
