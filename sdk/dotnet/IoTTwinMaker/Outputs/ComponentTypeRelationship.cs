// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTTwinMaker.Outputs
{

    /// <summary>
    /// The type of the relationship.
    /// </summary>
    [OutputType]
    public sealed class ComponentTypeRelationship
    {
        /// <summary>
        /// The type of the relationship.
        /// </summary>
        public readonly string? RelationshipType;
        /// <summary>
        /// The ID of the target component type associated with this relationship.
        /// </summary>
        public readonly string? TargetComponentTypeId;

        [OutputConstructor]
        private ComponentTypeRelationship(
            string? relationshipType,

            string? targetComponentTypeId)
        {
            RelationshipType = relationshipType;
            TargetComponentTypeId = targetComponentTypeId;
        }
    }
}