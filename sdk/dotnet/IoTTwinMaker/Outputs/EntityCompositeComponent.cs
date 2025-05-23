// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTTwinMaker.Outputs
{

    [OutputType]
    public sealed class EntityCompositeComponent
    {
        /// <summary>
        /// The name of the component.
        /// </summary>
        public readonly string? ComponentName;
        /// <summary>
        /// The path of the component.
        /// </summary>
        public readonly string? ComponentPath;
        /// <summary>
        /// The ID of the component type.
        /// </summary>
        public readonly string? ComponentTypeId;
        /// <summary>
        /// The description of the component.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// An object that maps strings to the properties to set in the component type. Each string in the mapping must be unique to this object.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.EntityProperty>? Properties;
        /// <summary>
        /// An object that maps strings to the property groups to set in the component type. Each string in the mapping must be unique to this object.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.EntityPropertyGroup>? PropertyGroups;
        /// <summary>
        /// The current status of the component.
        /// </summary>
        public readonly Outputs.EntityStatus? Status;

        [OutputConstructor]
        private EntityCompositeComponent(
            string? componentName,

            string? componentPath,

            string? componentTypeId,

            string? description,

            ImmutableDictionary<string, Outputs.EntityProperty>? properties,

            ImmutableDictionary<string, Outputs.EntityPropertyGroup>? propertyGroups,

            Outputs.EntityStatus? status)
        {
            ComponentName = componentName;
            ComponentPath = componentPath;
            ComponentTypeId = componentTypeId;
            Description = description;
            Properties = properties;
            PropertyGroups = propertyGroups;
            Status = status;
        }
    }
}
