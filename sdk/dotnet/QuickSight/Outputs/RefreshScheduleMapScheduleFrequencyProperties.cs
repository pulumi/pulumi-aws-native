// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// &lt;p&gt;Information about the schedule frequency.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class RefreshScheduleMapScheduleFrequencyProperties
    {
        public readonly Pulumi.AwsNative.QuickSight.RefreshScheduleMapScheduleFrequencyPropertiesInterval? Interval;
        /// <summary>
        /// &lt;p&gt;The day scheduled for refresh.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.RefreshScheduleMapScheduleFrequencyPropertiesRefreshOnDayProperties? RefreshOnDay;
        /// <summary>
        /// &lt;p&gt;The time of the day for scheduled refresh.&lt;/p&gt;
        /// </summary>
        public readonly string? TimeOfTheDay;
        /// <summary>
        /// &lt;p&gt;The timezone for scheduled refresh.&lt;/p&gt;
        /// </summary>
        public readonly string? TimeZone;

        [OutputConstructor]
        private RefreshScheduleMapScheduleFrequencyProperties(
            Pulumi.AwsNative.QuickSight.RefreshScheduleMapScheduleFrequencyPropertiesInterval? interval,

            Outputs.RefreshScheduleMapScheduleFrequencyPropertiesRefreshOnDayProperties? refreshOnDay,

            string? timeOfTheDay,

            string? timeZone)
        {
            Interval = interval;
            RefreshOnDay = refreshOnDay;
            TimeOfTheDay = timeOfTheDay;
            TimeZone = timeZone;
        }
    }
}