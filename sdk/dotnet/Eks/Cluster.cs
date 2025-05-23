// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Eks
{
    /// <summary>
    /// An object representing an Amazon EKS cluster.
    /// </summary>
    [AwsNativeResourceType("aws-native:eks:Cluster")]
    public partial class Cluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The access configuration for the cluster.
        /// </summary>
        [Output("accessConfig")]
        public Output<Outputs.ClusterAccessConfig?> AccessConfig { get; private set; } = null!;

        /// <summary>
        /// The ARN of the cluster, such as arn:aws:eks:us-west-2:666666666666:cluster/prod.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The unique ID given to your cluster.
        /// </summary>
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// Set this value to false to avoid creating the default networking add-ons when the cluster is created.
        /// </summary>
        [Output("bootstrapSelfManagedAddons")]
        public Output<bool?> BootstrapSelfManagedAddons { get; private set; } = null!;

        /// <summary>
        /// The certificate-authority-data for your cluster.
        /// </summary>
        [Output("certificateAuthorityData")]
        public Output<string> CertificateAuthorityData { get; private set; } = null!;

        /// <summary>
        /// The cluster security group that was created by Amazon EKS for the cluster. Managed node groups use this security group for control plane to data plane communication.
        /// </summary>
        [Output("clusterSecurityGroupId")]
        public Output<string> ClusterSecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// Indicates the current configuration of the compute capability on your EKS Auto Mode cluster. For example, if the capability is enabled or disabled. If the compute capability is enabled, EKS Auto Mode will create and delete EC2 Managed Instances in your AWS account. For more information, see EKS Auto Mode compute capability in the *Amazon EKS User Guide* .
        /// </summary>
        [Output("computeConfig")]
        public Output<Outputs.ClusterComputeConfig?> ComputeConfig { get; private set; } = null!;

        /// <summary>
        /// The encryption configuration for the cluster.
        /// </summary>
        [Output("encryptionConfig")]
        public Output<ImmutableArray<Outputs.ClusterEncryptionConfig>> EncryptionConfig { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) or alias of the customer master key (CMK).
        /// </summary>
        [Output("encryptionConfigKeyArn")]
        public Output<string> EncryptionConfigKeyArn { get; private set; } = null!;

        /// <summary>
        /// The endpoint for your Kubernetes API server, such as https://5E1D0CEXAMPLEA591B746AFC5AB30262.yl4.us-west-2.eks.amazonaws.com.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// Force cluster version update
        /// </summary>
        [Output("force")]
        public Output<bool?> Force { get; private set; } = null!;

        /// <summary>
        /// The Kubernetes network configuration for the cluster.
        /// </summary>
        [Output("kubernetesNetworkConfig")]
        public Output<Outputs.ClusterKubernetesNetworkConfig?> KubernetesNetworkConfig { get; private set; } = null!;

        /// <summary>
        /// The logging configuration for your cluster.
        /// </summary>
        [Output("logging")]
        public Output<Outputs.Logging?> Logging { get; private set; } = null!;

        /// <summary>
        /// The unique name to give to your cluster.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The issuer URL for the cluster's OIDC identity provider, such as https://oidc.eks.us-west-2.amazonaws.com/id/EXAMPLED539D4633E53DE1B716D3041E. If you need to remove https:// from this output value, you can include the following code in your template.
        /// </summary>
        [Output("openIdConnectIssuerUrl")]
        public Output<string> OpenIdConnectIssuerUrl { get; private set; } = null!;

        /// <summary>
        /// An object representing the configuration of your local Amazon EKS cluster on an AWS Outpost. This object isn't available for clusters on the AWS cloud.
        /// </summary>
        [Output("outpostConfig")]
        public Output<Outputs.ClusterOutpostConfig?> OutpostConfig { get; private set; } = null!;

        /// <summary>
        /// The configuration in the cluster for EKS Hybrid Nodes. You can add, change, or remove this configuration after the cluster is created.
        /// </summary>
        [Output("remoteNetworkConfig")]
        public Output<Outputs.ClusterRemoteNetworkConfig?> RemoteNetworkConfig { get; private set; } = null!;

        /// <summary>
        /// The VPC configuration that's used by the cluster control plane. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the *Amazon EKS User Guide* . You must specify at least two subnets. You can specify up to five security groups, but we recommend that you use a dedicated security group for your cluster control plane.
        /// </summary>
        [Output("resourcesVpcConfig")]
        public Output<Outputs.ClusterResourcesVpcConfig> ResourcesVpcConfig { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        /// <summary>
        /// Indicates the current configuration of the block storage capability on your EKS Auto Mode cluster. For example, if the capability is enabled or disabled. If the block storage capability is enabled, EKS Auto Mode will create and delete EBS volumes in your AWS account. For more information, see EKS Auto Mode block storage capability in the *Amazon EKS User Guide* .
        /// </summary>
        [Output("storageConfig")]
        public Output<Outputs.ClusterStorageConfig?> StorageConfig { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// This value indicates if extended support is enabled or disabled for the cluster.
        /// 
        /// [Learn more about EKS Extended Support in the *Amazon EKS User Guide* .](https://docs.aws.amazon.com/eks/latest/userguide/extended-support-control.html)
        /// </summary>
        [Output("upgradePolicy")]
        public Output<Outputs.ClusterUpgradePolicy?> UpgradePolicy { get; private set; } = null!;

        /// <summary>
        /// The desired Kubernetes version for your cluster. If you don't specify a value here, the latest version available in Amazon EKS is used.
        /// </summary>
        [Output("version")]
        public Output<string?> Version { get; private set; } = null!;

        /// <summary>
        /// The configuration for zonal shift for the cluster.
        /// </summary>
        [Output("zonalShiftConfig")]
        public Output<Outputs.ClusterZonalShiftConfig?> ZonalShiftConfig { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("aws-native:eks:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:eks:Cluster", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "accessConfig.bootstrapClusterCreatorAdminPermissions",
                    "bootstrapSelfManagedAddons",
                    "encryptionConfig[*]",
                    "kubernetesNetworkConfig.ipFamily",
                    "kubernetesNetworkConfig.serviceIpv4Cidr",
                    "name",
                    "outpostConfig",
                    "roleArn",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, options);
        }
    }

    public sealed class ClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access configuration for the cluster.
        /// </summary>
        [Input("accessConfig")]
        public Input<Inputs.ClusterAccessConfigArgs>? AccessConfig { get; set; }

        /// <summary>
        /// Set this value to false to avoid creating the default networking add-ons when the cluster is created.
        /// </summary>
        [Input("bootstrapSelfManagedAddons")]
        public Input<bool>? BootstrapSelfManagedAddons { get; set; }

        /// <summary>
        /// Indicates the current configuration of the compute capability on your EKS Auto Mode cluster. For example, if the capability is enabled or disabled. If the compute capability is enabled, EKS Auto Mode will create and delete EC2 Managed Instances in your AWS account. For more information, see EKS Auto Mode compute capability in the *Amazon EKS User Guide* .
        /// </summary>
        [Input("computeConfig")]
        public Input<Inputs.ClusterComputeConfigArgs>? ComputeConfig { get; set; }

        [Input("encryptionConfig")]
        private InputList<Inputs.ClusterEncryptionConfigArgs>? _encryptionConfig;

        /// <summary>
        /// The encryption configuration for the cluster.
        /// </summary>
        public InputList<Inputs.ClusterEncryptionConfigArgs> EncryptionConfig
        {
            get => _encryptionConfig ?? (_encryptionConfig = new InputList<Inputs.ClusterEncryptionConfigArgs>());
            set => _encryptionConfig = value;
        }

        /// <summary>
        /// Force cluster version update
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// The Kubernetes network configuration for the cluster.
        /// </summary>
        [Input("kubernetesNetworkConfig")]
        public Input<Inputs.ClusterKubernetesNetworkConfigArgs>? KubernetesNetworkConfig { get; set; }

        /// <summary>
        /// The logging configuration for your cluster.
        /// </summary>
        [Input("logging")]
        public Input<Inputs.LoggingArgs>? Logging { get; set; }

        /// <summary>
        /// The unique name to give to your cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// An object representing the configuration of your local Amazon EKS cluster on an AWS Outpost. This object isn't available for clusters on the AWS cloud.
        /// </summary>
        [Input("outpostConfig")]
        public Input<Inputs.ClusterOutpostConfigArgs>? OutpostConfig { get; set; }

        /// <summary>
        /// The configuration in the cluster for EKS Hybrid Nodes. You can add, change, or remove this configuration after the cluster is created.
        /// </summary>
        [Input("remoteNetworkConfig")]
        public Input<Inputs.ClusterRemoteNetworkConfigArgs>? RemoteNetworkConfig { get; set; }

        /// <summary>
        /// The VPC configuration that's used by the cluster control plane. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the *Amazon EKS User Guide* . You must specify at least two subnets. You can specify up to five security groups, but we recommend that you use a dedicated security group for your cluster control plane.
        /// </summary>
        [Input("resourcesVpcConfig", required: true)]
        public Input<Inputs.ClusterResourcesVpcConfigArgs> ResourcesVpcConfig { get; set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// Indicates the current configuration of the block storage capability on your EKS Auto Mode cluster. For example, if the capability is enabled or disabled. If the block storage capability is enabled, EKS Auto Mode will create and delete EBS volumes in your AWS account. For more information, see EKS Auto Mode block storage capability in the *Amazon EKS User Guide* .
        /// </summary>
        [Input("storageConfig")]
        public Input<Inputs.ClusterStorageConfigArgs>? StorageConfig { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// This value indicates if extended support is enabled or disabled for the cluster.
        /// 
        /// [Learn more about EKS Extended Support in the *Amazon EKS User Guide* .](https://docs.aws.amazon.com/eks/latest/userguide/extended-support-control.html)
        /// </summary>
        [Input("upgradePolicy")]
        public Input<Inputs.ClusterUpgradePolicyArgs>? UpgradePolicy { get; set; }

        /// <summary>
        /// The desired Kubernetes version for your cluster. If you don't specify a value here, the latest version available in Amazon EKS is used.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The configuration for zonal shift for the cluster.
        /// </summary>
        [Input("zonalShiftConfig")]
        public Input<Inputs.ClusterZonalShiftConfigArgs>? ZonalShiftConfig { get; set; }

        public ClusterArgs()
        {
        }
        public static new ClusterArgs Empty => new ClusterArgs();
    }
}
