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
    /// &lt;p&gt;The day scheduled for refresh.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class RefreshScheduleMapScheduleFrequencyPropertiesRefreshOnDayProperties
    {
        /// <summary>
        /// &lt;p&gt;The Day Of Month for scheduled refresh.&lt;/p&gt;
        /// </summary>
        public readonly string? DayOfMonth;
        public readonly Pulumi.AwsNative.QuickSight.RefreshScheduleMapScheduleFrequencyPropertiesRefreshOnDayPropertiesDayOfWeek? DayOfWeek;

        [OutputConstructor]
        private RefreshScheduleMapScheduleFrequencyPropertiesRefreshOnDayProperties(
            string? dayOfMonth,

            Pulumi.AwsNative.QuickSight.RefreshScheduleMapScheduleFrequencyPropertiesRefreshOnDayPropertiesDayOfWeek? dayOfWeek)
        {
            DayOfMonth = dayOfMonth;
            DayOfWeek = dayOfWeek;
        }
    }
}