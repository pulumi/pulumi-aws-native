// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.VpcLattice.Outputs
{

    [OutputType]
    public sealed class TargetGroupConfig
    {
        /// <summary>
        /// The health check configuration. Not supported if the target group type is `LAMBDA` or `ALB` .
        /// </summary>
        public readonly Outputs.TargetGroupHealthCheckConfig? HealthCheck;
        /// <summary>
        /// The type of IP address used for the target group. Supported only if the target group type is `IP` . The default is `IPV4` .
        /// </summary>
        public readonly Pulumi.AwsNative.VpcLattice.TargetGroupConfigIpAddressType? IpAddressType;
        /// <summary>
        /// The version of the event structure that your Lambda function receives. Supported only if the target group type is `LAMBDA` . The default is `V1` .
        /// </summary>
        public readonly Pulumi.AwsNative.VpcLattice.TargetGroupConfigLambdaEventStructureVersion? LambdaEventStructureVersion;
        /// <summary>
        /// The port on which the targets are listening. For HTTP, the default is 80. For HTTPS, the default is 443. Not supported if the target group type is `LAMBDA` .
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// The protocol to use for routing traffic to the targets. The default is the protocol of the target group. Not supported if the target group type is `LAMBDA` .
        /// </summary>
        public readonly Pulumi.AwsNative.VpcLattice.TargetGroupConfigProtocol? Protocol;
        /// <summary>
        /// The protocol version. The default is `HTTP1` . Not supported if the target group type is `LAMBDA` .
        /// </summary>
        public readonly Pulumi.AwsNative.VpcLattice.TargetGroupConfigProtocolVersion? ProtocolVersion;
        /// <summary>
        /// The ID of the VPC. Not supported if the target group type is `LAMBDA` .
        /// </summary>
        public readonly string? VpcIdentifier;

        [OutputConstructor]
        private TargetGroupConfig(
            Outputs.TargetGroupHealthCheckConfig? healthCheck,

            Pulumi.AwsNative.VpcLattice.TargetGroupConfigIpAddressType? ipAddressType,

            Pulumi.AwsNative.VpcLattice.TargetGroupConfigLambdaEventStructureVersion? lambdaEventStructureVersion,

            int? port,

            Pulumi.AwsNative.VpcLattice.TargetGroupConfigProtocol? protocol,

            Pulumi.AwsNative.VpcLattice.TargetGroupConfigProtocolVersion? protocolVersion,

            string? vpcIdentifier)
        {
            HealthCheck = healthCheck;
            IpAddressType = ipAddressType;
            LambdaEventStructureVersion = lambdaEventStructureVersion;
            Port = port;
            Protocol = protocol;
            ProtocolVersion = protocolVersion;
            VpcIdentifier = vpcIdentifier;
        }
    }
}
