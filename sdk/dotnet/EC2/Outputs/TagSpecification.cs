// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Outputs
{

    /// <summary>
    /// Specifies the tags to apply to a resource when the resource is created for the launch template.
    /// </summary>
    [OutputType]
    public sealed class TagSpecification
    {
        /// <summary>
        /// The type of resource to tag.
        /// </summary>
        public readonly string? ResourceType;
        /// <summary>
        /// The tags for the resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.LaunchTemplateTag> Tags;

        [OutputConstructor]
        private TagSpecification(
            string? resourceType,

            ImmutableArray<Outputs.LaunchTemplateTag> tags)
        {
            ResourceType = resourceType;
            Tags = tags;
        }
    }
}