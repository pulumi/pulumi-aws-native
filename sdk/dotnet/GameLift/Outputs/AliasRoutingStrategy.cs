// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift.Outputs
{

    [OutputType]
    public sealed class AliasRoutingStrategy
    {
        /// <summary>
        /// A unique identifier for a fleet that the alias points to. If you specify SIMPLE for the Type property, you must specify this property.
        /// </summary>
        public readonly string? FleetId;
        /// <summary>
        /// The message text to be used with a terminal routing strategy. If you specify TERMINAL for the Type property, you must specify this property.
        /// </summary>
        public readonly string? Message;
        /// <summary>
        /// Simple routing strategy. The alias resolves to one specific fleet. Use this type when routing to active fleets.
        /// </summary>
        public readonly Pulumi.AwsNative.GameLift.AliasRoutingStrategyType Type;

        [OutputConstructor]
        private AliasRoutingStrategy(
            string? fleetId,

            string? message,

            Pulumi.AwsNative.GameLift.AliasRoutingStrategyType type)
        {
            FleetId = fleetId;
            Message = message;
            Type = type;
        }
    }
}