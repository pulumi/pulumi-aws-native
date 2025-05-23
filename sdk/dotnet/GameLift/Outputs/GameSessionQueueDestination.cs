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
    /// A fleet or alias designated in a game session queue.
    /// </summary>
    [OutputType]
    public sealed class GameSessionQueueDestination
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) that is assigned to fleet or fleet alias. ARNs, which include a fleet ID or alias ID and a Region name, provide a unique identifier across all Regions.
        /// </summary>
        public readonly string? DestinationArn;

        [OutputConstructor]
        private GameSessionQueueDestination(string? destinationArn)
        {
            DestinationArn = destinationArn;
        }
    }
}
