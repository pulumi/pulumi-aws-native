// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk
{
    public static class GetCluster
    {
        /// <summary>
        /// Resource Type definition for AWS::MSK::Cluster
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("aws-native:msk:getCluster", args ?? new GetClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::MSK::Cluster
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClusterResult>("aws-native:msk:getCluster", args ?? new GetClusterInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::MSK::Cluster
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetClusterResult>("aws-native:msk:getCluster", args ?? new GetClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClusterArgs : global::Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetClusterArgs()
        {
        }
        public static new GetClusterArgs Empty => new GetClusterArgs();
    }

    public sealed class GetClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetClusterInvokeArgs()
        {
        }
        public static new GetClusterInvokeArgs Empty => new GetClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        public readonly string? Arn;
        public readonly Outputs.ClusterBrokerNodeGroupInfo? BrokerNodeGroupInfo;
        public readonly Outputs.ClusterClientAuthentication? ClientAuthentication;
        public readonly Outputs.ClusterConfigurationInfo? ConfigurationInfo;
        /// <summary>
        /// The current version of the MSK cluster
        /// </summary>
        public readonly string? CurrentVersion;
        public readonly Outputs.ClusterEncryptionInfo? EncryptionInfo;
        public readonly Pulumi.AwsNative.Msk.ClusterEnhancedMonitoring? EnhancedMonitoring;
        public readonly string? KafkaVersion;
        public readonly Outputs.ClusterLoggingInfo? LoggingInfo;
        public readonly int? NumberOfBrokerNodes;
        public readonly Outputs.ClusterOpenMonitoring? OpenMonitoring;
        public readonly Pulumi.AwsNative.Msk.ClusterStorageMode? StorageMode;
        /// <summary>
        /// A key-value pair to associate with a resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private GetClusterResult(
            string? arn,

            Outputs.ClusterBrokerNodeGroupInfo? brokerNodeGroupInfo,

            Outputs.ClusterClientAuthentication? clientAuthentication,

            Outputs.ClusterConfigurationInfo? configurationInfo,

            string? currentVersion,

            Outputs.ClusterEncryptionInfo? encryptionInfo,

            Pulumi.AwsNative.Msk.ClusterEnhancedMonitoring? enhancedMonitoring,

            string? kafkaVersion,

            Outputs.ClusterLoggingInfo? loggingInfo,

            int? numberOfBrokerNodes,

            Outputs.ClusterOpenMonitoring? openMonitoring,

            Pulumi.AwsNative.Msk.ClusterStorageMode? storageMode,

            ImmutableDictionary<string, string>? tags)
        {
            Arn = arn;
            BrokerNodeGroupInfo = brokerNodeGroupInfo;
            ClientAuthentication = clientAuthentication;
            ConfigurationInfo = configurationInfo;
            CurrentVersion = currentVersion;
            EncryptionInfo = encryptionInfo;
            EnhancedMonitoring = enhancedMonitoring;
            KafkaVersion = kafkaVersion;
            LoggingInfo = loggingInfo;
            NumberOfBrokerNodes = numberOfBrokerNodes;
            OpenMonitoring = openMonitoring;
            StorageMode = storageMode;
            Tags = tags;
        }
    }
}
