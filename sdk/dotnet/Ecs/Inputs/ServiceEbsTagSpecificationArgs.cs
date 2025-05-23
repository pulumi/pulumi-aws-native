// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Inputs
{

    /// <summary>
    /// The tag specifications of an Amazon EBS volume.
    /// </summary>
    public sealed class ServiceEbsTagSpecificationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines whether to propagate the tags from the task definition to 
        /// the Amazon EBS volume. Tags can only propagate to a ``SERVICE`` specified in 
        /// ``ServiceVolumeConfiguration``. If no value is specified, the tags aren't 
        /// propagated.
        /// </summary>
        [Input("propagateTags")]
        public Input<Pulumi.AwsNative.Ecs.ServiceEbsTagSpecificationPropagateTags>? PropagateTags { get; set; }

        /// <summary>
        /// The type of volume resource.
        /// </summary>
        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.ServiceTagArgs>? _tags;

        /// <summary>
        /// The tags applied to this Amazon EBS volume. ``AmazonECSCreated`` and ``AmazonECSManaged`` are reserved tags that can't be used.
        /// </summary>
        public InputList<Inputs.ServiceTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ServiceTagArgs>());
            set => _tags = value;
        }

        public ServiceEbsTagSpecificationArgs()
        {
        }
        public static new ServiceEbsTagSpecificationArgs Empty => new ServiceEbsTagSpecificationArgs();
    }
}
