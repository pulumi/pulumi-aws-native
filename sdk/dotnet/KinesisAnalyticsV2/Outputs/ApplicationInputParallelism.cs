// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2.Outputs
{

    /// <summary>
    /// For a SQL-based Kinesis Data Analytics application, describes the number of in-application streams to create for a given streaming source.
    /// </summary>
    [OutputType]
    public sealed class ApplicationInputParallelism
    {
        /// <summary>
        /// The number of in-application streams to create.
        /// </summary>
        public readonly int? Count;

        [OutputConstructor]
        private ApplicationInputParallelism(int? count)
        {
            Count = count;
        }
    }
}