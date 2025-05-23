// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Outputs
{

    /// <summary>
    /// Contains information about the hours of operation override.
    /// </summary>
    [OutputType]
    public sealed class HoursOfOperationOverrideConfig
    {
        /// <summary>
        /// The day that the hours of operation override applies to.
        /// </summary>
        public readonly Pulumi.AwsNative.Connect.HoursOfOperationOverrideConfigDay Day;
        /// <summary>
        /// The new end time that your contact center closes for the overriden days.
        /// </summary>
        public readonly Outputs.HoursOfOperationOverrideTimeSlice EndTime;
        /// <summary>
        /// The new start time that your contact center opens for the overriden days.
        /// </summary>
        public readonly Outputs.HoursOfOperationOverrideTimeSlice StartTime;

        [OutputConstructor]
        private HoursOfOperationOverrideConfig(
            Pulumi.AwsNative.Connect.HoursOfOperationOverrideConfigDay day,

            Outputs.HoursOfOperationOverrideTimeSlice endTime,

            Outputs.HoursOfOperationOverrideTimeSlice startTime)
        {
            Day = day;
            EndTime = endTime;
            StartTime = startTime;
        }
    }
}
