// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUiBuilder.Outputs
{

    [OutputType]
    public sealed class ComponentChild
    {
        public readonly ImmutableArray<Outputs.ComponentChild> Children;
        public readonly string ComponentType;
        public readonly Outputs.ComponentEvents? Events;
        public readonly string Name;
        public readonly Outputs.ComponentProperties Properties;

        [OutputConstructor]
        private ComponentChild(
            ImmutableArray<Outputs.ComponentChild> children,

            string componentType,

            Outputs.ComponentEvents? events,

            string name,

            Outputs.ComponentProperties properties)
        {
            Children = children;
            ComponentType = componentType;
            Events = events;
            Name = name;
            Properties = properties;
        }
    }
}