// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Outputs
{

    /// <summary>
    /// On premises settings which will have the interface network mappings and default Output logical interface
    /// </summary>
    [OutputType]
    public sealed class ClusterNetworkSettings
    {
        /// <summary>
        /// Default value if the customer does not define it in channel Output API
        /// </summary>
        public readonly string? DefaultRoute;
        /// <summary>
        /// Network mappings for the cluster
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterInterfaceMapping> InterfaceMappings;

        [OutputConstructor]
        private ClusterNetworkSettings(
            string? defaultRoute,

            ImmutableArray<Outputs.ClusterInterfaceMapping> interfaceMappings)
        {
            DefaultRoute = defaultRoute;
            InterfaceMappings = interfaceMappings;
        }
    }
}