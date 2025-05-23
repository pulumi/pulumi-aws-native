// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Outputs
{

    /// <summary>
    /// Information needed to set the timer.
    /// </summary>
    [OutputType]
    public sealed class DetectorModelSetTimer
    {
        /// <summary>
        /// The duration of the timer, in seconds. You can use a string expression that includes numbers, variables (``$variable.&lt;variable-name&gt;``), and input values (``$input.&lt;input-name&gt;.&lt;path-to-datum&gt;``) as the duration. The range of the duration is 1-31622400 seconds. To ensure accuracy, the minimum duration is 60 seconds. The evaluated result of the duration is rounded down to the nearest whole number.
        /// </summary>
        public readonly string? DurationExpression;
        /// <summary>
        /// The number of seconds until the timer expires. The minimum value is 60 seconds to ensure accuracy. The maximum value is 31622400 seconds.
        /// </summary>
        public readonly int? Seconds;
        /// <summary>
        /// The name of the timer.
        /// </summary>
        public readonly string TimerName;

        [OutputConstructor]
        private DetectorModelSetTimer(
            string? durationExpression,

            int? seconds,

            string timerName)
        {
            DurationExpression = durationExpression;
            Seconds = seconds;
            TimerName = timerName;
        }
    }
}
