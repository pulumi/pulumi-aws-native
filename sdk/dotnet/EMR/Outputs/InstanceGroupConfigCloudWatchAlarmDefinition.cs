// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EMR.Outputs
{

    [OutputType]
    public sealed class InstanceGroupConfigCloudWatchAlarmDefinition
    {
        public readonly string ComparisonOperator;
        public readonly ImmutableArray<Outputs.InstanceGroupConfigMetricDimension> Dimensions;
        public readonly int? EvaluationPeriods;
        public readonly string MetricName;
        public readonly string? Namespace;
        public readonly int Period;
        public readonly string? Statistic;
        public readonly double Threshold;
        public readonly string? Unit;

        [OutputConstructor]
        private InstanceGroupConfigCloudWatchAlarmDefinition(
            string comparisonOperator,

            ImmutableArray<Outputs.InstanceGroupConfigMetricDimension> dimensions,

            int? evaluationPeriods,

            string metricName,

            string? @namespace,

            int period,

            string? statistic,

            double threshold,

            string? unit)
        {
            ComparisonOperator = comparisonOperator;
            Dimensions = dimensions;
            EvaluationPeriods = evaluationPeriods;
            MetricName = metricName;
            Namespace = @namespace;
            Period = period;
            Statistic = statistic;
            Threshold = threshold;
            Unit = unit;
        }
    }
}