// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FSx.Outputs
{

    [OutputType]
    public sealed class FileSystemOpenZFSConfiguration
    {
        public readonly int? AutomaticBackupRetentionDays;
        public readonly bool? CopyTagsToBackups;
        public readonly bool? CopyTagsToVolumes;
        public readonly string? DailyAutomaticBackupStartTime;
        public readonly string DeploymentType;
        public readonly Outputs.FileSystemDiskIopsConfiguration? DiskIopsConfiguration;
        public readonly ImmutableArray<string> Options;
        public readonly Outputs.FileSystemRootVolumeConfiguration? RootVolumeConfiguration;
        public readonly int? ThroughputCapacity;
        public readonly string? WeeklyMaintenanceStartTime;

        [OutputConstructor]
        private FileSystemOpenZFSConfiguration(
            int? automaticBackupRetentionDays,

            bool? copyTagsToBackups,

            bool? copyTagsToVolumes,

            string? dailyAutomaticBackupStartTime,

            string deploymentType,

            Outputs.FileSystemDiskIopsConfiguration? diskIopsConfiguration,

            ImmutableArray<string> options,

            Outputs.FileSystemRootVolumeConfiguration? rootVolumeConfiguration,

            int? throughputCapacity,

            string? weeklyMaintenanceStartTime)
        {
            AutomaticBackupRetentionDays = automaticBackupRetentionDays;
            CopyTagsToBackups = copyTagsToBackups;
            CopyTagsToVolumes = copyTagsToVolumes;
            DailyAutomaticBackupStartTime = dailyAutomaticBackupStartTime;
            DeploymentType = deploymentType;
            DiskIopsConfiguration = diskIopsConfiguration;
            Options = options;
            RootVolumeConfiguration = rootVolumeConfiguration;
            ThroughputCapacity = throughputCapacity;
            WeeklyMaintenanceStartTime = weeklyMaintenanceStartTime;
        }
    }
}