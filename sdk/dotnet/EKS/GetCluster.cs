// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EKS
{
    public static class GetCluster
    {
        /// <summary>
        /// An object representing an Amazon EKS cluster.
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("aws-native:eks:getCluster", args ?? new GetClusterArgs(), options.WithDefaults());

        /// <summary>
        /// An object representing an Amazon EKS cluster.
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClusterResult>("aws-native:eks:getCluster", args ?? new GetClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClusterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique name to give to your cluster.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetClusterArgs()
        {
        }
        public static new GetClusterArgs Empty => new GetClusterArgs();
    }

    public sealed class GetClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique name to give to your cluster.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetClusterInvokeArgs()
        {
        }
        public static new GetClusterInvokeArgs Empty => new GetClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        /// <summary>
        /// The ARN of the cluster, such as arn:aws:eks:us-west-2:666666666666:cluster/prod.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The certificate-authority-data for your cluster.
        /// </summary>
        public readonly string? CertificateAuthorityData;
        /// <summary>
        /// The cluster security group that was created by Amazon EKS for the cluster. Managed node groups use this security group for control plane to data plane communication.
        /// </summary>
        public readonly string? ClusterSecurityGroupId;
        /// <summary>
        /// Amazon Resource Name (ARN) or alias of the customer master key (CMK).
        /// </summary>
        public readonly string? EncryptionConfigKeyArn;
        /// <summary>
        /// The endpoint for your Kubernetes API server, such as https://5E1D0CEXAMPLEA591B746AFC5AB30262.yl4.us-west-2.eks.amazonaws.com.
        /// </summary>
        public readonly string? Endpoint;
        /// <summary>
        /// The unique ID given to your cluster.
        /// </summary>
        public readonly string? Id;
        public readonly Outputs.ClusterLogging? Logging;
        /// <summary>
        /// The issuer URL for the cluster's OIDC identity provider, such as https://oidc.eks.us-west-2.amazonaws.com/id/EXAMPLED539D4633E53DE1B716D3041E. If you need to remove https:// from this output value, you can include the following code in your template.
        /// </summary>
        public readonly string? OpenIdConnectIssuerUrl;
        public readonly Outputs.ClusterResourcesVpcConfig? ResourcesVpcConfig;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterTag> Tags;
        /// <summary>
        /// The desired Kubernetes version for your cluster. If you don't specify a value here, the latest version available in Amazon EKS is used.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private GetClusterResult(
            string? arn,

            string? certificateAuthorityData,

            string? clusterSecurityGroupId,

            string? encryptionConfigKeyArn,

            string? endpoint,

            string? id,

            Outputs.ClusterLogging? logging,

            string? openIdConnectIssuerUrl,

            Outputs.ClusterResourcesVpcConfig? resourcesVpcConfig,

            ImmutableArray<Outputs.ClusterTag> tags,

            string? version)
        {
            Arn = arn;
            CertificateAuthorityData = certificateAuthorityData;
            ClusterSecurityGroupId = clusterSecurityGroupId;
            EncryptionConfigKeyArn = encryptionConfigKeyArn;
            Endpoint = endpoint;
            Id = id;
            Logging = logging;
            OpenIdConnectIssuerUrl = openIdConnectIssuerUrl;
            ResourcesVpcConfig = resourcesVpcConfig;
            Tags = tags;
            Version = version;
        }
    }
}