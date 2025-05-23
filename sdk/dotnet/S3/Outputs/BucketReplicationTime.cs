// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Outputs
{

    /// <summary>
    /// A container specifying S3 Replication Time Control (S3 RTC) related information, including whether S3 RTC is enabled and the time when all objects and operations on objects must be replicated. Must be specified together with a ``Metrics`` block.
    /// </summary>
    [OutputType]
    public sealed class BucketReplicationTime
    {
        /// <summary>
        /// Specifies whether the replication time is enabled.
        /// </summary>
        public readonly Pulumi.AwsNative.S3.BucketReplicationTimeStatus Status;
        /// <summary>
        /// A container specifying the time by which replication should be complete for all objects and operations on objects.
        /// </summary>
        public readonly Outputs.BucketReplicationTimeValue Time;

        [OutputConstructor]
        private BucketReplicationTime(
            Pulumi.AwsNative.S3.BucketReplicationTimeStatus status,

            Outputs.BucketReplicationTimeValue time)
        {
            Status = status;
            Time = time;
        }
    }
}
