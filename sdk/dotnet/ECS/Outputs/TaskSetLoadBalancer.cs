// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ECS.Outputs
{

    /// <summary>
    /// A load balancer object representing the load balancer to use with the task set. The supported load balancer types are either an Application Load Balancer or a Network Load Balancer. 
    /// </summary>
    [OutputType]
    public sealed class TaskSetLoadBalancer
    {
        /// <summary>
        /// The name of the container (as it appears in a container definition) to associate with the load balancer.
        /// </summary>
        public readonly string? ContainerName;
        /// <summary>
        /// The port on the container to associate with the load balancer. This port must correspond to a containerPort in the task definition the tasks in the service are using. For tasks that use the EC2 launch type, the container instance they are launched on must allow ingress traffic on the hostPort of the port mapping.
        /// </summary>
        public readonly int? ContainerPort;
        /// <summary>
        /// The name of the load balancer to associate with the Amazon ECS service or task set. A load balancer name is only specified when using a Classic Load Balancer. If you are using an Application Load Balancer or a Network Load Balancer this should be omitted.
        /// </summary>
        public readonly string? LoadBalancerName;
        /// <summary>
        /// The full Amazon Resource Name (ARN) of the Elastic Load Balancing target group or groups associated with a service or task set. A target group ARN is only specified when using an Application Load Balancer or Network Load Balancer. If you are using a Classic Load Balancer this should be omitted. For services using the ECS deployment controller, you can specify one or multiple target groups. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/register-multiple-targetgroups.html in the Amazon Elastic Container Service Developer Guide. For services using the CODE_DEPLOY deployment controller, you are required to define two target groups for the load balancer. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-bluegreen.html in the Amazon Elastic Container Service Developer Guide. If your service's task definition uses the awsvpc network mode (which is required for the Fargate launch type), you must choose ip as the target type, not instance, when creating your target groups because tasks that use the awsvpc network mode are associated with an elastic network interface, not an Amazon EC2 instance.
        /// </summary>
        public readonly string? TargetGroupArn;

        [OutputConstructor]
        private TaskSetLoadBalancer(
            string? containerName,

            int? containerPort,

            string? loadBalancerName,

            string? targetGroupArn)
        {
            ContainerName = containerName;
            ContainerPort = containerPort;
            LoadBalancerName = loadBalancerName;
            TargetGroupArn = targetGroupArn;
        }
    }
}