// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Outputs
{

    /// <summary>
    /// The attributes for the instance types.
    /// </summary>
    [OutputType]
    public sealed class LaunchTemplateInstanceRequirements
    {
        public readonly Outputs.LaunchTemplateAcceleratorCount? AcceleratorCount;
        /// <summary>
        /// Indicates whether instance types must have accelerators by specific manufacturers.
        /// </summary>
        public readonly ImmutableArray<string> AcceleratorManufacturers;
        /// <summary>
        /// The accelerators that must be on the instance type.
        /// </summary>
        public readonly ImmutableArray<string> AcceleratorNames;
        public readonly Outputs.LaunchTemplateAcceleratorTotalMemoryMiB? AcceleratorTotalMemoryMiB;
        /// <summary>
        /// The accelerator types that must be on the instance type.
        /// </summary>
        public readonly ImmutableArray<string> AcceleratorTypes;
        /// <summary>
        /// The instance types to apply your specified attributes against.
        /// </summary>
        public readonly ImmutableArray<string> AllowedInstanceTypes;
        /// <summary>
        /// Indicates whether bare metal instance types must be included, excluded, or required.
        /// </summary>
        public readonly string? BareMetal;
        public readonly Outputs.LaunchTemplateBaselineEbsBandwidthMbps? BaselineEbsBandwidthMbps;
        public readonly string? BurstablePerformance;
        /// <summary>
        /// The CPU manufacturers to include.
        /// </summary>
        public readonly ImmutableArray<string> CpuManufacturers;
        /// <summary>
        /// The instance types to exclude.
        /// </summary>
        public readonly ImmutableArray<string> ExcludedInstanceTypes;
        /// <summary>
        /// Indicates whether current or previous generation instance types are included.
        /// </summary>
        public readonly ImmutableArray<string> InstanceGenerations;
        /// <summary>
        /// The user data to make available to the instance.
        /// </summary>
        public readonly string? LocalStorage;
        /// <summary>
        /// The type of local storage that is required.
        /// </summary>
        public readonly ImmutableArray<string> LocalStorageTypes;
        /// <summary>
        /// The price protection threshold for Spot Instances.
        /// </summary>
        public readonly int? MaxSpotPriceAsPercentageOfOptimalOnDemandPrice;
        public readonly Outputs.LaunchTemplateMemoryGiBPerVCpu? MemoryGiBPerVCpu;
        public readonly Outputs.LaunchTemplateMemoryMiB? MemoryMiB;
        public readonly Outputs.LaunchTemplateNetworkBandwidthGbps? NetworkBandwidthGbps;
        public readonly Outputs.LaunchTemplateNetworkInterfaceCount? NetworkInterfaceCount;
        /// <summary>
        /// The price protection threshold for On-Demand Instances.
        /// </summary>
        public readonly int? OnDemandMaxPricePercentageOverLowestPrice;
        /// <summary>
        /// Indicates whether instance types must support hibernation for On-Demand Instances.
        /// </summary>
        public readonly bool? RequireHibernateSupport;
        /// <summary>
        /// The price protection threshold for Spot Instances.
        /// </summary>
        public readonly int? SpotMaxPricePercentageOverLowestPrice;
        public readonly Outputs.LaunchTemplateTotalLocalStorageGb? TotalLocalStorageGb;
        public readonly Outputs.LaunchTemplateVCpuCount? VCpuCount;

        [OutputConstructor]
        private LaunchTemplateInstanceRequirements(
            Outputs.LaunchTemplateAcceleratorCount? acceleratorCount,

            ImmutableArray<string> acceleratorManufacturers,

            ImmutableArray<string> acceleratorNames,

            Outputs.LaunchTemplateAcceleratorTotalMemoryMiB? acceleratorTotalMemoryMiB,

            ImmutableArray<string> acceleratorTypes,

            ImmutableArray<string> allowedInstanceTypes,

            string? bareMetal,

            Outputs.LaunchTemplateBaselineEbsBandwidthMbps? baselineEbsBandwidthMbps,

            string? burstablePerformance,

            ImmutableArray<string> cpuManufacturers,

            ImmutableArray<string> excludedInstanceTypes,

            ImmutableArray<string> instanceGenerations,

            string? localStorage,

            ImmutableArray<string> localStorageTypes,

            int? maxSpotPriceAsPercentageOfOptimalOnDemandPrice,

            Outputs.LaunchTemplateMemoryGiBPerVCpu? memoryGiBPerVCpu,

            Outputs.LaunchTemplateMemoryMiB? memoryMiB,

            Outputs.LaunchTemplateNetworkBandwidthGbps? networkBandwidthGbps,

            Outputs.LaunchTemplateNetworkInterfaceCount? networkInterfaceCount,

            int? onDemandMaxPricePercentageOverLowestPrice,

            bool? requireHibernateSupport,

            int? spotMaxPricePercentageOverLowestPrice,

            Outputs.LaunchTemplateTotalLocalStorageGb? totalLocalStorageGb,

            Outputs.LaunchTemplateVCpuCount? vCpuCount)
        {
            AcceleratorCount = acceleratorCount;
            AcceleratorManufacturers = acceleratorManufacturers;
            AcceleratorNames = acceleratorNames;
            AcceleratorTotalMemoryMiB = acceleratorTotalMemoryMiB;
            AcceleratorTypes = acceleratorTypes;
            AllowedInstanceTypes = allowedInstanceTypes;
            BareMetal = bareMetal;
            BaselineEbsBandwidthMbps = baselineEbsBandwidthMbps;
            BurstablePerformance = burstablePerformance;
            CpuManufacturers = cpuManufacturers;
            ExcludedInstanceTypes = excludedInstanceTypes;
            InstanceGenerations = instanceGenerations;
            LocalStorage = localStorage;
            LocalStorageTypes = localStorageTypes;
            MaxSpotPriceAsPercentageOfOptimalOnDemandPrice = maxSpotPriceAsPercentageOfOptimalOnDemandPrice;
            MemoryGiBPerVCpu = memoryGiBPerVCpu;
            MemoryMiB = memoryMiB;
            NetworkBandwidthGbps = networkBandwidthGbps;
            NetworkInterfaceCount = networkInterfaceCount;
            OnDemandMaxPricePercentageOverLowestPrice = onDemandMaxPricePercentageOverLowestPrice;
            RequireHibernateSupport = requireHibernateSupport;
            SpotMaxPricePercentageOverLowestPrice = spotMaxPricePercentageOverLowestPrice;
            TotalLocalStorageGb = totalLocalStorageGb;
            VCpuCount = vCpuCount;
        }
    }
}