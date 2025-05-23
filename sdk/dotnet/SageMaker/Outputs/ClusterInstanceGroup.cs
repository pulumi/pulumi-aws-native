// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// Details of an instance group in a SageMaker HyperPod cluster.
    /// </summary>
    [OutputType]
    public sealed class ClusterInstanceGroup
    {
        /// <summary>
        /// The number of instances that are currently in the instance group of a SageMaker HyperPod cluster.
        /// </summary>
        public readonly int? CurrentCount;
        public readonly string ExecutionRole;
        /// <summary>
        /// The number of instances you specified to add to the instance group of a SageMaker HyperPod cluster.
        /// </summary>
        public readonly int InstanceCount;
        public readonly string InstanceGroupName;
        public readonly ImmutableArray<Outputs.ClusterInstanceStorageConfig> InstanceStorageConfigs;
        public readonly string InstanceType;
        public readonly Outputs.ClusterLifeCycleConfig LifeCycleConfig;
        public readonly ImmutableArray<Pulumi.AwsNative.SageMaker.ClusterDeepHealthCheckType> OnStartDeepHealthChecks;
        public readonly Outputs.ClusterVpcConfig? OverrideVpcConfig;
        /// <summary>
        /// The number you specified to TreadsPerCore in CreateCluster for enabling or disabling multithreading. For instance types that support multithreading, you can specify 1 for disabling multithreading and 2 for enabling multithreading.
        /// </summary>
        public readonly int? ThreadsPerCore;

        [OutputConstructor]
        private ClusterInstanceGroup(
            int? currentCount,

            string executionRole,

            int instanceCount,

            string instanceGroupName,

            ImmutableArray<Outputs.ClusterInstanceStorageConfig> instanceStorageConfigs,

            string instanceType,

            Outputs.ClusterLifeCycleConfig lifeCycleConfig,

            ImmutableArray<Pulumi.AwsNative.SageMaker.ClusterDeepHealthCheckType> onStartDeepHealthChecks,

            Outputs.ClusterVpcConfig? overrideVpcConfig,

            int? threadsPerCore)
        {
            CurrentCount = currentCount;
            ExecutionRole = executionRole;
            InstanceCount = instanceCount;
            InstanceGroupName = instanceGroupName;
            InstanceStorageConfigs = instanceStorageConfigs;
            InstanceType = instanceType;
            LifeCycleConfig = lifeCycleConfig;
            OnStartDeepHealthChecks = onStartDeepHealthChecks;
            OverrideVpcConfig = overrideVpcConfig;
            ThreadsPerCore = threadsPerCore;
        }
    }
}
