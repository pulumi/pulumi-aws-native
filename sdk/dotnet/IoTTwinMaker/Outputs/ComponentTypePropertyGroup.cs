// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTTwinMaker.Outputs
{

    /// <summary>
    /// An object that sets information about a property group.
    /// </summary>
    [OutputType]
    public sealed class ComponentTypePropertyGroup
    {
        /// <summary>
        /// The type of property group.
        /// </summary>
        public readonly Pulumi.AwsNative.IoTTwinMaker.ComponentTypePropertyGroupGroupType? GroupType;
        /// <summary>
        /// The list of property names in the property group.
        /// </summary>
        public readonly ImmutableArray<string> PropertyNames;

        [OutputConstructor]
        private ComponentTypePropertyGroup(
            Pulumi.AwsNative.IoTTwinMaker.ComponentTypePropertyGroupGroupType? groupType,

            ImmutableArray<string> propertyNames)
        {
            GroupType = groupType;
            PropertyNames = propertyNames;
        }
    }
}
