// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Inputs
{

    /// <summary>
    /// The start time or end time for an hours of operation.
    /// </summary>
    public sealed class HoursOfOperationTimeSliceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The hours.
        /// </summary>
        [Input("hours", required: true)]
        public Input<int> Hours { get; set; } = null!;

        /// <summary>
        /// The minutes.
        /// </summary>
        [Input("minutes", required: true)]
        public Input<int> Minutes { get; set; } = null!;

        public HoursOfOperationTimeSliceArgs()
        {
        }
        public static new HoursOfOperationTimeSliceArgs Empty => new HoursOfOperationTimeSliceArgs();
    }
}