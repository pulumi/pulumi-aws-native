// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Outputs
{

    [OutputType]
    public sealed class FlowAggregationConfig
    {
        /// <summary>
        /// Specifies whether Amazon AppFlow aggregates the flow records into a single file, or leave them unaggregated.
        /// </summary>
        public readonly Pulumi.AwsNative.AppFlow.FlowAggregationType? AggregationType;
        /// <summary>
        /// The desired file size, in MB, for each output file that Amazon AppFlow writes to the flow destination. For each file, Amazon AppFlow attempts to achieve the size that you specify. The actual file sizes might differ from this target based on the number and size of the records that each file contains.
        /// </summary>
        public readonly int? TargetFileSize;

        [OutputConstructor]
        private FlowAggregationConfig(
            Pulumi.AwsNative.AppFlow.FlowAggregationType? aggregationType,

            int? targetFileSize)
        {
            AggregationType = aggregationType;
            TargetFileSize = targetFileSize;
        }
    }
}
