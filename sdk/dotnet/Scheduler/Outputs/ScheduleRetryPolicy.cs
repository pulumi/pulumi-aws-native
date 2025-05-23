// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Scheduler.Outputs
{

    /// <summary>
    /// A RetryPolicy object that includes information about the retry policy settings.
    /// </summary>
    [OutputType]
    public sealed class ScheduleRetryPolicy
    {
        /// <summary>
        /// The maximum amount of time, in seconds, to continue to make retry attempts.
        /// </summary>
        public readonly double? MaximumEventAgeInSeconds;
        /// <summary>
        /// The maximum number of retry attempts to make before the request fails. Retry attempts with exponential backoff continue until either the maximum number of attempts is made or until the duration of the MaximumEventAgeInSeconds is reached.
        /// </summary>
        public readonly double? MaximumRetryAttempts;

        [OutputConstructor]
        private ScheduleRetryPolicy(
            double? maximumEventAgeInSeconds,

            double? maximumRetryAttempts)
        {
            MaximumEventAgeInSeconds = maximumEventAgeInSeconds;
            MaximumRetryAttempts = maximumRetryAttempts;
        }
    }
}
