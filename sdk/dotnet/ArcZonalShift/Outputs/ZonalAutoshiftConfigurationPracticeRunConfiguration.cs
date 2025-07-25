// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ArcZonalShift.Outputs
{

    [OutputType]
    public sealed class ZonalAutoshiftConfigurationPracticeRunConfiguration
    {
        /// <summary>
        /// An array of one or more dates that you can specify when AWS does not start practice runs for a resource. Dates are in UTC.
        /// 
        /// Specify blocked dates in the format `YYYY-MM-DD` , separated by spaces.
        /// </summary>
        public readonly ImmutableArray<string> BlockedDates;
        /// <summary>
        /// An array of one or more days and times that you can specify when ARC does not start practice runs for a resource. Days and times are in UTC.
        /// 
        /// Specify blocked windows in the format `DAY:HH:MM-DAY:HH:MM` , separated by spaces. For example, `MON:18:30-MON:19:30 TUE:18:30-TUE:19:30` .
        /// 
        /// &gt; Blocked windows have to start and end on the same day. Windows that span multiple days aren't supported.
        /// </summary>
        public readonly ImmutableArray<string> BlockedWindows;
        /// <summary>
        /// An optional alarm that you can specify that blocks practice runs when the alarm is in an `ALARM` state. When a blocking alarm goes into an `ALARM` state, it prevents practice runs from being started, and ends practice runs that are in progress.
        /// </summary>
        public readonly ImmutableArray<Outputs.ZonalAutoshiftConfigurationControlCondition> BlockingAlarms;
        /// <summary>
        /// The alarm that you specify to monitor the health of your application during practice runs. When the outcome alarm goes into an `ALARM` state, the practice run is ended and the outcome is set to `FAILED` .
        /// </summary>
        public readonly ImmutableArray<Outputs.ZonalAutoshiftConfigurationControlCondition> OutcomeAlarms;

        [OutputConstructor]
        private ZonalAutoshiftConfigurationPracticeRunConfiguration(
            ImmutableArray<string> blockedDates,

            ImmutableArray<string> blockedWindows,

            ImmutableArray<Outputs.ZonalAutoshiftConfigurationControlCondition> blockingAlarms,

            ImmutableArray<Outputs.ZonalAutoshiftConfigurationControlCondition> outcomeAlarms)
        {
            BlockedDates = blockedDates;
            BlockedWindows = blockedWindows;
            BlockingAlarms = blockingAlarms;
            OutcomeAlarms = outcomeAlarms;
        }
    }
}
