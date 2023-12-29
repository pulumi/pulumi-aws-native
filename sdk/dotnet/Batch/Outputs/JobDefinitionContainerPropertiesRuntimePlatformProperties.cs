// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Outputs
{

    [OutputType]
    public sealed class JobDefinitionContainerPropertiesRuntimePlatformProperties
    {
        public readonly string? CpuArchitecture;
        public readonly string? OperatingSystemFamily;

        [OutputConstructor]
        private JobDefinitionContainerPropertiesRuntimePlatformProperties(
            string? cpuArchitecture,

            string? operatingSystemFamily)
        {
            CpuArchitecture = cpuArchitecture;
            OperatingSystemFamily = operatingSystemFamily;
        }
    }
}