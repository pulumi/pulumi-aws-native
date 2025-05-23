// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Outputs
{

    /// <summary>
    /// Specifies a start time and duration for a scheduled Job.
    /// </summary>
    [OutputType]
    public sealed class JobTemplateMaintenanceWindow
    {
        /// <summary>
        /// Displays the duration of the next maintenance window.
        /// </summary>
        public readonly int? DurationInMinutes;
        /// <summary>
        /// Displays the start time of the next maintenance window.
        /// </summary>
        public readonly string? StartTime;

        [OutputConstructor]
        private JobTemplateMaintenanceWindow(
            int? durationInMinutes,

            string? startTime)
        {
            DurationInMinutes = durationInMinutes;
            StartTime = startTime;
        }
    }
}
