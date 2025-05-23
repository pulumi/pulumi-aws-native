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
    /// The deployment circuit breaker can only be used for services using the rolling update (``ECS``) deployment type.
    ///   The *deployment circuit breaker* determines whether a service deployment will fail if the service can't reach a steady state. If it is turned on, a service deployment will transition to a failed state and stop launching new tasks. You can also configure Amazon ECS to roll back your service to the last completed deployment after a failure. For more information, see [Rolling update](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-ecs.html) in the *Amazon Elastic Container Service Developer Guide*.
    ///  For more information about API failure reasons, see [API failure reasons](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/api_failures_messages.html) in the *Amazon Elastic Container Service Developer Guide*.
    /// </summary>
    public sealed class ServiceDeploymentCircuitBreakerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines whether to use the deployment circuit breaker logic for the service.
        /// </summary>
        [Input("enable", required: true)]
        public Input<bool> Enable { get; set; } = null!;

        /// <summary>
        /// Determines whether to configure Amazon ECS to roll back the service if a service deployment fails. If rollback is on, when a service deployment fails, the service is rolled back to the last deployment that completed successfully.
        /// </summary>
        [Input("rollback", required: true)]
        public Input<bool> Rollback { get; set; } = null!;

        public ServiceDeploymentCircuitBreakerArgs()
        {
        }
        public static new ServiceDeploymentCircuitBreakerArgs Empty => new ServiceDeploymentCircuitBreakerArgs();
    }
}
