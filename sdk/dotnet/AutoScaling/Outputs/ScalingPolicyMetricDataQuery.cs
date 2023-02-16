// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScaling.Outputs
{

    [OutputType]
    public sealed class ScalingPolicyMetricDataQuery
    {
        public readonly string? Expression;
        public readonly string Id;
        public readonly string? Label;
        public readonly Outputs.ScalingPolicyMetricStat? MetricStat;
        public readonly bool? ReturnData;

        [OutputConstructor]
        private ScalingPolicyMetricDataQuery(
            string? expression,

            string id,

            string? label,

            Outputs.ScalingPolicyMetricStat? metricStat,

            bool? returnData)
        {
            Expression = expression;
            Id = id;
            Label = label;
            MetricStat = metricStat;
            ReturnData = returnData;
        }
    }
}