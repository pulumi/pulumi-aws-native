// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CustomerProfiles.Inputs
{

    /// <summary>
    /// The day and time when do you want to start the Identity Resolution Job every week.
    /// </summary>
    public sealed class DomainJobScheduleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The day when the Identity Resolution Job should run every week.
        /// </summary>
        [Input("dayOfTheWeek", required: true)]
        public Input<Pulumi.AwsNative.CustomerProfiles.DomainJobScheduleDayOfTheWeek> DayOfTheWeek { get; set; } = null!;

        /// <summary>
        /// The time when the Identity Resolution Job should run every week.
        /// </summary>
        [Input("time", required: true)]
        public Input<string> Time { get; set; } = null!;

        public DomainJobScheduleArgs()
        {
        }
        public static new DomainJobScheduleArgs Empty => new DomainJobScheduleArgs();
    }
}