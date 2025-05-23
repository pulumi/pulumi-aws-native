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
    /// Describes the health check configuration of an Amazon Lightsail container service.
    /// </summary>
    public sealed class ContainerHealthCheckConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of consecutive health checks successes required before moving the container to the Healthy state. The default value is 2.
        /// </summary>
        [Input("healthyThreshold")]
        public Input<int>? HealthyThreshold { get; set; }

        /// <summary>
        /// The approximate interval, in seconds, between health checks of an individual container. You can specify between 5 and 300 seconds. The default value is 5.
        /// </summary>
        [Input("intervalSeconds")]
        public Input<int>? IntervalSeconds { get; set; }

        /// <summary>
        /// The path on the container on which to perform the health check. The default value is /.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The HTTP codes to use when checking for a successful response from a container. You can specify values between 200 and 499. You can specify multiple values (for example, 200,202) or a range of values (for example, 200-299).
        /// </summary>
        [Input("successCodes")]
        public Input<string>? SuccessCodes { get; set; }

        /// <summary>
        /// The amount of time, in seconds, during which no response means a failed health check. You can specify between 2 and 60 seconds. The default value is 2.
        /// </summary>
        [Input("timeoutSeconds")]
        public Input<int>? TimeoutSeconds { get; set; }

        /// <summary>
        /// The number of consecutive health check failures required before moving the container to the Unhealthy state. The default value is 2.
        /// </summary>
        [Input("unhealthyThreshold")]
        public Input<int>? UnhealthyThreshold { get; set; }

        public ContainerHealthCheckConfigArgs()
        {
        }
        public static new ContainerHealthCheckConfigArgs Empty => new ContainerHealthCheckConfigArgs();
    }
}
