// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Timestream.Outputs
{

    /// <summary>
    /// The retention duration of the memory store and the magnetic store.
    /// </summary>
    [OutputType]
    public sealed class RetentionPropertiesProperties
    {
        /// <summary>
        /// The duration for which data must be stored in the magnetic store.
        /// </summary>
        public readonly string? MagneticStoreRetentionPeriodInDays;
        /// <summary>
        /// The duration for which data must be stored in the memory store.
        /// </summary>
        public readonly string? MemoryStoreRetentionPeriodInHours;

        [OutputConstructor]
        private RetentionPropertiesProperties(
            string? magneticStoreRetentionPeriodInDays,

            string? memoryStoreRetentionPeriodInHours)
        {
            MagneticStoreRetentionPeriodInDays = magneticStoreRetentionPeriodInDays;
            MemoryStoreRetentionPeriodInHours = memoryStoreRetentionPeriodInHours;
        }
    }
}
