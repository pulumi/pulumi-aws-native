// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ImageBuilder.Inputs
{

    /// <summary>
    /// The placement options
    /// </summary>
    public sealed class InfrastructureConfigurationPlacementArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AvailabilityZone
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// HostId
        /// </summary>
        [Input("hostId")]
        public Input<string>? HostId { get; set; }

        /// <summary>
        /// HostResourceGroupArn
        /// </summary>
        [Input("hostResourceGroupArn")]
        public Input<string>? HostResourceGroupArn { get; set; }

        /// <summary>
        /// Tenancy
        /// </summary>
        [Input("tenancy")]
        public Input<Pulumi.AwsNative.ImageBuilder.InfrastructureConfigurationPlacementTenancy>? Tenancy { get; set; }

        public InfrastructureConfigurationPlacementArgs()
        {
        }
        public static new InfrastructureConfigurationPlacementArgs Empty => new InfrastructureConfigurationPlacementArgs();
    }
}
