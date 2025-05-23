// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CustomerProfiles.Outputs
{

    [OutputType]
    public sealed class IntegrationScheduledTriggerProperties
    {
        /// <summary>
        /// Specifies whether a scheduled flow has an incremental data transfer or a complete data transfer for each flow run.
        /// </summary>
        public readonly Pulumi.AwsNative.CustomerProfiles.IntegrationScheduledTriggerPropertiesDataPullMode? DataPullMode;
        /// <summary>
        /// Specifies the date range for the records to import from the connector in the first flow run.
        /// </summary>
        public readonly double? FirstExecutionFrom;
        /// <summary>
        /// Specifies the scheduled end time for a scheduled-trigger flow.
        /// </summary>
        public readonly double? ScheduleEndTime;
        /// <summary>
        /// The scheduling expression that determines the rate at which the schedule will run, for example rate (5 minutes).
        /// </summary>
        public readonly string ScheduleExpression;
        /// <summary>
        /// Specifies the optional offset that is added to the time interval for a schedule-triggered flow.
        /// </summary>
        public readonly int? ScheduleOffset;
        /// <summary>
        /// Specifies the scheduled start time for a scheduled-trigger flow. The value must be a date/time value in EPOCH format.
        /// </summary>
        public readonly double? ScheduleStartTime;
        /// <summary>
        /// Specifies the time zone used when referring to the date and time of a scheduled-triggered flow, such as America/New_York.
        /// </summary>
        public readonly string? Timezone;

        [OutputConstructor]
        private IntegrationScheduledTriggerProperties(
            Pulumi.AwsNative.CustomerProfiles.IntegrationScheduledTriggerPropertiesDataPullMode? dataPullMode,

            double? firstExecutionFrom,

            double? scheduleEndTime,

            string scheduleExpression,

            int? scheduleOffset,

            double? scheduleStartTime,

            string? timezone)
        {
            DataPullMode = dataPullMode;
            FirstExecutionFrom = firstExecutionFrom;
            ScheduleEndTime = scheduleEndTime;
            ScheduleExpression = scheduleExpression;
            ScheduleOffset = scheduleOffset;
            ScheduleStartTime = scheduleStartTime;
            Timezone = timezone;
        }
    }
}
