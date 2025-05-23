// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudTrail.Outputs
{

    /// <summary>
    /// Configures the automatic refresh schedule for the dashboard. Includes the frequency unit (DAYS or HOURS) and value, as well as the status (ENABLED or DISABLED) of the refresh schedule.
    /// </summary>
    [OutputType]
    public sealed class DashboardRefreshSchedule
    {
        /// <summary>
        /// The frequency at which you want the dashboard refreshed.
        /// </summary>
        public readonly Outputs.DashboardRefreshScheduleFrequencyProperties? Frequency;
        /// <summary>
        /// The status of the schedule. Supported values are ENABLED and DISABLED.
        /// </summary>
        public readonly Pulumi.AwsNative.CloudTrail.DashboardRefreshScheduleStatus? Status;
        /// <summary>
        /// StartTime of the automatic schedule refresh.
        /// </summary>
        public readonly string? TimeOfDay;

        [OutputConstructor]
        private DashboardRefreshSchedule(
            Outputs.DashboardRefreshScheduleFrequencyProperties? frequency,

            Pulumi.AwsNative.CloudTrail.DashboardRefreshScheduleStatus? status,

            string? timeOfDay)
        {
            Frequency = frequency;
            Status = status;
            TimeOfDay = timeOfDay;
        }
    }
}
