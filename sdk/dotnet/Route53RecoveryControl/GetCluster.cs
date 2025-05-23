// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53RecoveryControl
{
    public static class GetCluster
    {
        /// <summary>
        /// AWS Route53 Recovery Control Cluster resource schema
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("aws-native:route53recoverycontrol:getCluster", args ?? new GetClusterArgs(), options.WithDefaults());

        /// <summary>
        /// AWS Route53 Recovery Control Cluster resource schema
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClusterResult>("aws-native:route53recoverycontrol:getCluster", args ?? new GetClusterInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// AWS Route53 Recovery Control Cluster resource schema
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetClusterResult>("aws-native:route53recoverycontrol:getCluster", args ?? new GetClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClusterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the cluster.
        /// </summary>
        [Input("clusterArn", required: true)]
        public string ClusterArn { get; set; } = null!;

        public GetClusterArgs()
        {
        }
        public static new GetClusterArgs Empty => new GetClusterArgs();
    }

    public sealed class GetClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the cluster.
        /// </summary>
        [Input("clusterArn", required: true)]
        public Input<string> ClusterArn { get; set; } = null!;

        public GetClusterInvokeArgs()
        {
        }
        public static new GetClusterInvokeArgs Empty => new GetClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the cluster.
        /// </summary>
        public readonly string? ClusterArn;
        /// <summary>
        /// Endpoints for the cluster.
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterEndpoint> ClusterEndpoints;
        /// <summary>
        /// Cluster supports IPv4 endpoints and Dual-stack IPv4 and IPv6 endpoints. NetworkType can be IPV4 or DUALSTACK.
        /// </summary>
        public readonly Pulumi.AwsNative.Route53RecoveryControl.ClusterNetworkType? NetworkType;
        /// <summary>
        /// Deployment status of a resource. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.
        /// </summary>
        public readonly Pulumi.AwsNative.Route53RecoveryControl.ClusterStatus? Status;

        [OutputConstructor]
        private GetClusterResult(
            string? clusterArn,

            ImmutableArray<Outputs.ClusterEndpoint> clusterEndpoints,

            Pulumi.AwsNative.Route53RecoveryControl.ClusterNetworkType? networkType,

            Pulumi.AwsNative.Route53RecoveryControl.ClusterStatus? status)
        {
            ClusterArn = clusterArn;
            ClusterEndpoints = clusterEndpoints;
            NetworkType = networkType;
            Status = status;
        }
    }
}
