// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Outputs
{

    /// <summary>
    /// Information required to reset the timer. The timer is reset to the previously evaluated result of the duration. The duration expression isn't reevaluated when you reset the timer.
    /// </summary>
    [OutputType]
    public sealed class DetectorModelResetTimer
    {
        /// <summary>
        /// The name of the timer to reset.
        /// </summary>
        public readonly string TimerName;

        [OutputConstructor]
        private DetectorModelResetTimer(string timerName)
        {
            TimerName = timerName;
        }
    }
}