// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2.Outputs
{

    [OutputType]
    public sealed class TargetGroupTargetDescription
    {
        /// <summary>
        /// An Availability Zone or all. This determines whether the target receives traffic from the load balancer nodes in the specified Availability Zone or from all enabled Availability Zones for the load balancer.
        /// </summary>
        public readonly string? AvailabilityZone;
        /// <summary>
        /// The ID of the target. If the target type of the target group is instance, specify an instance ID. If the target type is ip, specify an IP address. If the target type is lambda, specify the ARN of the Lambda function. If the target type is alb, specify the ARN of the Application Load Balancer target. 
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The port on which the target is listening. If the target group protocol is GENEVE, the supported port is 6081. If the target type is alb, the targeted Application Load Balancer must have at least one listener whose port matches the target group port. Not used if the target is a Lambda function.
        /// </summary>
        public readonly int? Port;

        [OutputConstructor]
        private TargetGroupTargetDescription(
            string? availabilityZone,

            string id,

            int? port)
        {
            AvailabilityZone = availabilityZone;
            Id = id;
            Port = port;
        }
    }
}
