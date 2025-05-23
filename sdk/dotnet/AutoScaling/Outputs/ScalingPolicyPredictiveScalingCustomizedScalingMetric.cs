// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScaling.Outputs
{

    [OutputType]
    public sealed class ScalingPolicyPredictiveScalingCustomizedScalingMetric
    {
        /// <summary>
        /// One or more metric data queries to provide the data points for a scaling metric. Use multiple metric data queries only if you are performing a math expression on returned data.
        /// </summary>
        public readonly ImmutableArray<Outputs.ScalingPolicyMetricDataQuery> MetricDataQueries;

        [OutputConstructor]
        private ScalingPolicyPredictiveScalingCustomizedScalingMetric(ImmutableArray<Outputs.ScalingPolicyMetricDataQuery> metricDataQueries)
        {
            MetricDataQueries = metricDataQueries;
        }
    }
}
