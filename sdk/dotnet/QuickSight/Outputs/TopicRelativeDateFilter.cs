// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TopicRelativeDateFilter
    {
        public readonly Outputs.TopicSingularFilterConstant? Constant;
        public readonly Pulumi.AwsNative.QuickSight.TopicRelativeDateFilterFunction? RelativeDateFilterFunction;
        public readonly Pulumi.AwsNative.QuickSight.TopicTimeGranularity? TimeGranularity;

        [OutputConstructor]
        private TopicRelativeDateFilter(
            Outputs.TopicSingularFilterConstant? constant,

            Pulumi.AwsNative.QuickSight.TopicRelativeDateFilterFunction? relativeDateFilterFunction,

            Pulumi.AwsNative.QuickSight.TopicTimeGranularity? timeGranularity)
        {
            Constant = constant;
            RelativeDateFilterFunction = relativeDateFilterFunction;
            TimeGranularity = timeGranularity;
        }
    }
}