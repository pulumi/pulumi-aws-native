// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ImageBuilder.Outputs
{

    /// <summary>
    /// The Windows faster-launching configuration to use for AMI distribution.
    /// </summary>
    [OutputType]
    public sealed class DistributionConfigurationFastLaunchConfiguration
    {
        /// <summary>
        /// The owner account ID for the fast-launch enabled Windows AMI.
        /// </summary>
        public readonly string? AccountId;
        /// <summary>
        /// A Boolean that represents the current state of faster launching for the Windows AMI. Set to true to start using Windows faster launching, or false to stop using it.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The launch template that the fast-launch enabled Windows AMI uses when it launches Windows instances to create pre-provisioned snapshots.
        /// </summary>
        public readonly Outputs.DistributionConfigurationFastLaunchLaunchTemplateSpecification? LaunchTemplate;
        /// <summary>
        /// The maximum number of parallel instances that are launched for creating resources.
        /// </summary>
        public readonly int? MaxParallelLaunches;
        /// <summary>
        /// Configuration settings for managing the number of snapshots that are created from pre-provisioned instances for the Windows AMI when faster launching is enabled.
        /// </summary>
        public readonly Outputs.DistributionConfigurationFastLaunchSnapshotConfiguration? SnapshotConfiguration;

        [OutputConstructor]
        private DistributionConfigurationFastLaunchConfiguration(
            string? accountId,

            bool? enabled,

            Outputs.DistributionConfigurationFastLaunchLaunchTemplateSpecification? launchTemplate,

            int? maxParallelLaunches,

            Outputs.DistributionConfigurationFastLaunchSnapshotConfiguration? snapshotConfiguration)
        {
            AccountId = accountId;
            Enabled = enabled;
            LaunchTemplate = launchTemplate;
            MaxParallelLaunches = maxParallelLaunches;
            SnapshotConfiguration = snapshotConfiguration;
        }
    }
}
