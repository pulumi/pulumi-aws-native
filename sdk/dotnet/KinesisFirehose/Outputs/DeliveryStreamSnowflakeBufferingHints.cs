// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Outputs
{

    [OutputType]
    public sealed class DeliveryStreamSnowflakeBufferingHints
    {
        /// <summary>
        /// Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination. The default value is 0.
        /// </summary>
        public readonly int? IntervalInSeconds;
        /// <summary>
        /// Buffer incoming data to the specified size, in MBs, before delivering it to the destination. The default value is 1.
        /// </summary>
        public readonly int? SizeInMbs;

        [OutputConstructor]
        private DeliveryStreamSnowflakeBufferingHints(
            int? intervalInSeconds,

            int? sizeInMbs)
        {
            IntervalInSeconds = intervalInSeconds;
            SizeInMbs = sizeInMbs;
        }
    }
}