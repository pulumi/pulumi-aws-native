// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Outputs
{

    /// <summary>
    /// The start time or end time for an hours of operation.
    /// </summary>
    [OutputType]
    public sealed class HoursOfOperationTimeSlice
    {
        /// <summary>
        /// The hours.
        /// </summary>
        public readonly int Hours;
        /// <summary>
        /// The minutes.
        /// </summary>
        public readonly int Minutes;

        [OutputConstructor]
        private HoursOfOperationTimeSlice(
            int hours,

            int minutes)
        {
            Hours = hours;
            Minutes = minutes;
        }
    }
}