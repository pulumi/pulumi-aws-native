// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkManager.Outputs
{

    [OutputType]
    public sealed class CoreNetworkSegment
    {
        public readonly ImmutableArray<string> EdgeLocations;
        /// <summary>
        /// Name of segment
        /// </summary>
        public readonly string? Name;
        public readonly ImmutableArray<string> SharedSegments;

        [OutputConstructor]
        private CoreNetworkSegment(
            ImmutableArray<string> edgeLocations,

            string? name,

            ImmutableArray<string> sharedSegments)
        {
            EdgeLocations = edgeLocations;
            Name = name;
            SharedSegments = sharedSegments;
        }
    }
}