// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScaling.Outputs
{

    [OutputType]
    public sealed class ScalingPolicyTargetTrackingMetricDataQuery
    {
        public readonly string? Expression;
        public readonly string Id;
        public readonly string? Label;
        public readonly Outputs.ScalingPolicyTargetTrackingMetricStat? MetricStat;
        public readonly bool? ReturnData;

        [OutputConstructor]
        private ScalingPolicyTargetTrackingMetricDataQuery(
            string? expression,

            string id,

            string? label,

            Outputs.ScalingPolicyTargetTrackingMetricStat? metricStat,

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