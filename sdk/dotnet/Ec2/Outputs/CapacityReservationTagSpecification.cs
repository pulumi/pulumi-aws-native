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
    public sealed class CapacityReservationTagSpecification
    {
        /// <summary>
        /// The type of resource to tag. Specify `capacity-reservation` .
        /// </summary>
        public readonly string? ResourceType;
        /// <summary>
        /// The tags to apply to the resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.CapacityReservationTag> Tags;

        [OutputConstructor]
        private CapacityReservationTagSpecification(
            string? resourceType,

            ImmutableArray<Outputs.CapacityReservationTag> tags)
        {
            ResourceType = resourceType;
            Tags = tags;
        }
    }
}
