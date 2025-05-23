// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lightsail.Inputs
{

    /// <summary>
    /// Describes the settings of a public endpoint for an Amazon Lightsail container service.
    /// </summary>
    public sealed class ContainerPublicEndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the container for the endpoint.
        /// </summary>
        [Input("containerName")]
        public Input<string>? ContainerName { get; set; }

        /// <summary>
        /// The port of the container to which traffic is forwarded to.
        /// </summary>
        [Input("containerPort")]
        public Input<int>? ContainerPort { get; set; }

        /// <summary>
        /// An object that describes the health check configuration of the container.
        /// </summary>
        [Input("healthCheckConfig")]
        public Input<Inputs.ContainerHealthCheckConfigArgs>? HealthCheckConfig { get; set; }

        public ContainerPublicEndpointArgs()
        {
        }
        public static new ContainerPublicEndpointArgs Empty => new ContainerPublicEndpointArgs();
    }
}
