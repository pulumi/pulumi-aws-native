// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2
{
    public static class GetTargetGroup
    {
        /// <summary>
        /// Resource Type definition for AWS::ElasticLoadBalancingV2::TargetGroup
        /// </summary>
        public static Task<GetTargetGroupResult> InvokeAsync(GetTargetGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTargetGroupResult>("aws-native:elasticloadbalancingv2:getTargetGroup", args ?? new GetTargetGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ElasticLoadBalancingV2::TargetGroup
        /// </summary>
        public static Output<GetTargetGroupResult> Invoke(GetTargetGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTargetGroupResult>("aws-native:elasticloadbalancingv2:getTargetGroup", args ?? new GetTargetGroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ElasticLoadBalancingV2::TargetGroup
        /// </summary>
        public static Output<GetTargetGroupResult> Invoke(GetTargetGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetTargetGroupResult>("aws-native:elasticloadbalancingv2:getTargetGroup", args ?? new GetTargetGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTargetGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the Target Group
        /// </summary>
        [Input("targetGroupArn", required: true)]
        public string TargetGroupArn { get; set; } = null!;

        public GetTargetGroupArgs()
        {
        }
        public static new GetTargetGroupArgs Empty => new GetTargetGroupArgs();
    }

    public sealed class GetTargetGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the Target Group
        /// </summary>
        [Input("targetGroupArn", required: true)]
        public Input<string> TargetGroupArn { get; set; } = null!;

        public GetTargetGroupInvokeArgs()
        {
        }
        public static new GetTargetGroupInvokeArgs Empty => new GetTargetGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetTargetGroupResult
    {
        /// <summary>
        /// Indicates whether health checks are enabled. If the target type is lambda, health checks are disabled by default but can be enabled. If the target type is instance, ip, or alb, health checks are always enabled and cannot be disabled.
        /// </summary>
        public readonly bool? HealthCheckEnabled;
        /// <summary>
        /// The approximate amount of time, in seconds, between health checks of an individual target.
        /// </summary>
        public readonly int? HealthCheckIntervalSeconds;
        /// <summary>
        /// [HTTP/HTTPS health checks] The destination for health checks on the targets. [HTTP1 or HTTP2 protocol version] The ping path. The default is /. [GRPC protocol version] The path of a custom health check method with the format /package.service/method. The default is /AWS.ALB/healthcheck.
        /// </summary>
        public readonly string? HealthCheckPath;
        /// <summary>
        /// The port the load balancer uses when performing health checks on targets. 
        /// </summary>
        public readonly string? HealthCheckPort;
        /// <summary>
        /// The protocol the load balancer uses when performing health checks on targets. 
        /// </summary>
        public readonly string? HealthCheckProtocol;
        /// <summary>
        /// The amount of time, in seconds, during which no response from a target means a failed health check.
        /// </summary>
        public readonly int? HealthCheckTimeoutSeconds;
        /// <summary>
        /// The number of consecutive health checks successes required before considering an unhealthy target healthy. 
        /// </summary>
        public readonly int? HealthyThresholdCount;
        /// <summary>
        /// The Amazon Resource Names (ARNs) of the load balancers that route traffic to this target group.
        /// </summary>
        public readonly ImmutableArray<string> LoadBalancerArns;
        /// <summary>
        /// [HTTP/HTTPS health checks] The HTTP or gRPC codes to use when checking for a successful response from a target.
        /// </summary>
        public readonly Outputs.TargetGroupMatcher? Matcher;
        /// <summary>
        /// The tags.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// The ARN of the Target Group
        /// </summary>
        public readonly string? TargetGroupArn;
        /// <summary>
        /// The attributes.
        /// </summary>
        public readonly ImmutableArray<Outputs.TargetGroupAttribute> TargetGroupAttributes;
        /// <summary>
        /// The full name of the target group.
        /// </summary>
        public readonly string? TargetGroupFullName;
        /// <summary>
        /// The name of the target group.
        /// </summary>
        public readonly string? TargetGroupName;
        /// <summary>
        /// The targets.
        /// </summary>
        public readonly ImmutableArray<Outputs.TargetGroupTargetDescription> Targets;
        /// <summary>
        /// The number of consecutive health check failures required before considering a target unhealthy.
        /// </summary>
        public readonly int? UnhealthyThresholdCount;

        [OutputConstructor]
        private GetTargetGroupResult(
            bool? healthCheckEnabled,

            int? healthCheckIntervalSeconds,

            string? healthCheckPath,

            string? healthCheckPort,

            string? healthCheckProtocol,

            int? healthCheckTimeoutSeconds,

            int? healthyThresholdCount,

            ImmutableArray<string> loadBalancerArns,

            Outputs.TargetGroupMatcher? matcher,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            string? targetGroupArn,

            ImmutableArray<Outputs.TargetGroupAttribute> targetGroupAttributes,

            string? targetGroupFullName,

            string? targetGroupName,

            ImmutableArray<Outputs.TargetGroupTargetDescription> targets,

            int? unhealthyThresholdCount)
        {
            HealthCheckEnabled = healthCheckEnabled;
            HealthCheckIntervalSeconds = healthCheckIntervalSeconds;
            HealthCheckPath = healthCheckPath;
            HealthCheckPort = healthCheckPort;
            HealthCheckProtocol = healthCheckProtocol;
            HealthCheckTimeoutSeconds = healthCheckTimeoutSeconds;
            HealthyThresholdCount = healthyThresholdCount;
            LoadBalancerArns = loadBalancerArns;
            Matcher = matcher;
            Tags = tags;
            TargetGroupArn = targetGroupArn;
            TargetGroupAttributes = targetGroupAttributes;
            TargetGroupFullName = targetGroupFullName;
            TargetGroupName = targetGroupName;
            Targets = targets;
            UnhealthyThresholdCount = unhealthyThresholdCount;
        }
    }
}
