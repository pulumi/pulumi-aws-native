// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift.Inputs
{

    /// <summary>
    /// Specifies container groups that this instance will hold. You must specify exactly one replica group. Optionally, you may specify exactly one daemon group. You can't change this property after you create the fleet.
    /// </summary>
    public sealed class FleetContainerGroupsConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("connectionPortRange", required: true)]
        public Input<Inputs.FleetConnectionPortRangeArgs> ConnectionPortRange { get; set; } = null!;

        [Input("containerGroupDefinitionNames", required: true)]
        private InputList<string>? _containerGroupDefinitionNames;

        /// <summary>
        /// The names of the container group definitions that will be created in an instance. You must specify exactly one REPLICA container group. You have the option to also specify one DAEMON container group.
        /// </summary>
        public InputList<string> ContainerGroupDefinitionNames
        {
            get => _containerGroupDefinitionNames ?? (_containerGroupDefinitionNames = new InputList<string>());
            set => _containerGroupDefinitionNames = value;
        }

        [Input("containerGroupsPerInstance")]
        public Input<Inputs.FleetContainerGroupsPerInstanceArgs>? ContainerGroupsPerInstance { get; set; }

        public FleetContainerGroupsConfigurationArgs()
        {
        }
        public static new FleetContainerGroupsConfigurationArgs Empty => new FleetContainerGroupsConfigurationArgs();
    }
}